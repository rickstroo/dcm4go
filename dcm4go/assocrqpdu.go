package dcm4go

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// AssocRQPDU borrows from the AssocACRQPDU
type AssocRQPDU struct {
	AssocACRQPDU
}

// newAssocRQPDU creates a new association request PDU
func newAssocRQPDU(calledAETitle string, callingAETitle string, presContexts []*PresContext) *AssocRQPDU {
	return &AssocRQPDU{
		AssocACRQPDU{
			0x01,                    // protocol version, as per the standard
			calledAETitle,           // title of the called, as per the standard
			callingAETitle,          // title of the caller, as per the standard
			"1.2.840.10008.3.1.1.1", // app context name, as per the standard
			presContexts,            // pres context list
			&UserInfo{
				16378,             // max length received, need to figure out why dcm4che uses this number
				"1.2.40.0.13.1.3", // implementation class uid, need to get a root, borrowing dcm4che for now
				"dcm4go-1.0",      // implementation class name
				0,                 // max num ops invoked
				0,                 // max num ops performed
			},
		},
	}
}

// readAssocRQPDU reads an AssocRQPDU from a reader
func readAssocRQPDU(reader io.Reader) (*AssocRQPDU, error) {

	// read the protocol
	protocol, err := readShort(reader, binary.BigEndian)
	if err != nil {
		return nil, err
	}

	// skip two bytes, as per the standard
	if err := skipBytes(reader, 2); err != nil {
		return nil, err
	}

	// read the called AE title
	calledAETitle, err := readText(reader, 16)
	if err != nil {
		return nil, err
	}

	// read the calling AE title
	callingAETitle, err := readText(reader, 16)
	if err != nil {
		return nil, err
	}

	// skip thirty two bytes as per the standard
	if err := skipBytes(reader, 32); err != nil {
		return nil, err
	}

	// initialize the application context name
	var appContextName string

	// initialize a list of presentation contexts
	presContexts := make([]*PresContext, 0, 5)

	// initialize the user info
	var userInfo *UserInfo

	for {

		// read an item
		itemType, err := readByte(reader)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		// skip a byte, as per the standard
		if err := skipByte(reader); err != nil {
			return nil, err
		}

		// read the length
		length, err := readShort(reader, binary.BigEndian)
		if err != nil {
			return nil, err
		}

		if itemType == 0x10 { // application context name

			// read the application context name
			appContextName, err = readUID(reader, uint32(length))
			if err != nil {
				return nil, err
			}

		} else if itemType == 0x020 { // presentation context

			// create a limited reader for the requested presentation contextx
			limitedReader := io.LimitReader(reader, int64(length))

			// read the presentation context id
			id, err := readByte(limitedReader)
			if err != nil {
				return nil, err
			}

			// skip a byte, as per the standrd
			if err := skipBytes(limitedReader, 3); err != nil {
				return nil, err
			}

			// initialize the abstract syntax name
			var abstractSyntax string

			// initialize the list of transfer syntaxes
			transferSyntaxes := make([]string, 0, 5)

			// read the abstract syntax and transfer syntax items
			for {

				// read a sub item
				subItemType, err := readByte(limitedReader)
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return nil, err
				}

				// skip a byte, as per the standard
				if err := skipByte(limitedReader); err != nil {
					return nil, err
				}

				// read the length
				length, err := readShort(limitedReader, binary.BigEndian)
				if err != nil {
					return nil, err
				}

				if subItemType == 0x30 { // abstract syntax

					// read the uid
					uid, err := readUID(limitedReader, uint32(length))
					if err != nil {
						return nil, err
					}

					// assign it to the abstract syntax
					abstractSyntax = uid

				} else if subItemType == 0x40 { // transfer syntax

					// read the transfer syntax
					uid, err := readUID(limitedReader, uint32(length))
					if err != nil {
						return nil, err
					}

					// add it to the list of transfer syntaxes
					transferSyntaxes = append(transferSyntaxes, uid)

				} else { // unrecgonized item
					return nil, fmt.Errorf("unrecognized presentation context sub item type: 0x%02X", subItemType)
				}
			}

			// create the presentation context
			presContext := &PresContext{
				id,               // the pc id
				abstractSyntax,   // the abstract syntax
				transferSyntaxes, // the transfer syntaxes
				0x00,             // no reason
			}

			// add it to the list of requested presentation contexts
			presContexts = append(presContexts, presContext)

		} else if itemType == 0x050 { // user info

			// create a limited reader for the user info
			limitedReader := io.LimitReader(reader, int64(length))

			// read the user info sub items
			userInfo, err = readUserInfo(limitedReader)
			if err != nil {
				return nil, err
			}

		} else {
			// unrecognized item
			return nil, fmt.Errorf("unrecognized item type: 0x%02X", itemType)
		}
	}

	// construct and return an association request pdu
	return &AssocRQPDU{
			AssocACRQPDU{
				protocol,
				calledAETitle,
				callingAETitle,
				appContextName,
				presContexts,
				userInfo},
		},
		nil
}

func writeAssocRQPDU(writer io.Writer, assocRQPDU *AssocRQPDU) error {

	// write pdu type
	if err := writeByte(writer, 0x02); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// create a byte array output stream so we can calculate the length of the rest of the PDU
	byteWriter := new(bytes.Buffer)

	// write the protocol version
	if err := writeShort(byteWriter, assocRQPDU.protocol, binary.BigEndian); err != nil {
		return err
	}

	// write a short zero
	if err := writeShort(byteWriter, 0x00, binary.BigEndian); err != nil {
		return err
	}

	// write the called ae title
	if err := writeString(byteWriter, assocRQPDU.calledAETitle); err != nil {
		return err
	}

	// write the calling ae title
	if err := writeString(byteWriter, assocRQPDU.callingAETitle); err != nil {
		return err
	}

	// write thirty two zeroes, zero is the initial value for arrays, so this works
	var zeros [32]byte
	if err := writeBytes(byteWriter, zeros[:]); err != nil {
		return err
	}

	// write the variable items
	if err := writeVariableItems(byteWriter, assocRQPDU.appContextName, assocRQPDU.presContexts, 0x20, assocRQPDU.userInfo); err != nil {
		return err
	}

	// write the length to the original writer
	if err := writeLong(writer, uint32(byteWriter.Len()), binary.BigEndian); err != nil {
		return err
	}

	// write the byte array to the original writer
	if err := writeBytes(writer, byteWriter.Bytes()); err != nil {
		return err

	}

	// all is good
	return nil
}
