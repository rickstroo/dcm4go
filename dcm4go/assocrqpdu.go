package dcm4go

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
)

// AssocRQPDU presents an Association Request PDU
type AssocRQPDU struct {
	protocol       uint16
	calledAETitle  string
	callingAETitle string
	appContextName string
	presContexts   []*RQPresContext
	userInfo       *UserInfo
}

// RQPresContext represents a presentation context
type RQPresContext struct {
	id               byte
	abstractSyntax   string
	transferSyntaxes []string
}

// String returns a string representation of a AssocRQPDU
func (pdu *AssocRQPDU) String() string {

	// format the simple attribute of the pdu
	s := fmt.Sprintf(
		"{protocol:%v,calledAET:%q,callingAET:%q,appContextName:%q,presContexts:[",
		pdu.protocol,
		strings.TrimSpace(pdu.calledAETitle),
		strings.TrimSpace(pdu.callingAETitle),
		pdu.appContextName)

	// print the list of presentation contexts
	for _, presContext := range pdu.presContexts {
		s += fmt.Sprintf("%s,", presContext)
	}

	// trim the trailing comma separating presentation contexts and add the enclosing bracket
	s = strings.TrimSuffix(s, ",") + "]"

	// add the user info
	s += fmt.Sprintf(",userInfo:%v", pdu.userInfo)

	// return the fully constructed string including a closing parenthesis
	return s + "}"
}

// String returns a string representation of a UserInfo
func (presContext *RQPresContext) String() string {
	s := fmt.Sprintf(
		"{id:%d,abstractSyntax:%q,transferSyntaxes:[",
		presContext.id,
		presContext.abstractSyntax)

	// print the list of transfer syntaxes
	for _, transferSyntax := range presContext.transferSyntaxes {
		s += fmt.Sprintf("%q,", transferSyntax)
	}

	// trim the trailing comma separating transfer syntaxes and add the enclosing bracket
	s = strings.TrimSuffix(s, ",") + "]"

	// return the fully constructed including a closing parenthesis for this presentation context
	return s + "}"
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
	presContexts := make([]*RQPresContext, 0, 5)

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

		if itemType == 0x10 { // application context name

			// read the length
			length, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}

			// read the application context name
			appContextName, err = readUID(reader, uint32(length))
			if err != nil {
				return nil, err
			}

		} else if itemType == 0x020 { // presentation context

			// read the length
			length, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}

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

				if subItemType == 0x30 { // abstract syntax

					// read the length
					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					// read the uid
					uid, err := readUID(limitedReader, uint32(length))
					if err != nil {
						return nil, err
					}

					// assign it to the abstract syntax
					abstractSyntax = uid

				} else if subItemType == 0x40 { // transfer syntax

					// read the length
					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

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
			presContext := &RQPresContext{id, abstractSyntax, transferSyntaxes}

			// add it to the list of requested presentation contexts
			presContexts = append(presContexts, presContext)

		} else if itemType == 0x050 { // user info

			// read the length
			length, err := readShort(reader, binary.BigEndian)
			if err != nil {
				return nil, err
			}

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

	return &AssocRQPDU{protocol, calledAETitle, callingAETitle, appContextName, presContexts, userInfo}, nil
}
