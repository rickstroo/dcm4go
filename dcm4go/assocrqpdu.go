package dcm4go

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// AssocRQPDU borrows from the AssocACRQPDU
type AssocRQPDU struct {
	AssocACRQPDU
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
