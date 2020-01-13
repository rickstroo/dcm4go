package dcm4go

import (
	"container/list"
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
	presContexts   *list.List
	userInfo       *UserInfo
}

// String returns a string representation of a AssocRQPDU
func (pdu *AssocRQPDU) String() string {

	// format the simple attribute of the pdu
	s := fmt.Sprintf("protocol: %v, calledAET: %q, callingAET: %q, appContextName: %q, presContexts: [", pdu.protocol, strings.TrimSpace(pdu.calledAETitle), strings.TrimSpace(pdu.callingAETitle), pdu.appContextName)

	// print the list of presentation contexts
	for item := pdu.presContexts.Front(); item != nil; item = item.Next() {
		presContext := item.Value.(*PresContext)
		s += fmt.Sprintf("{id: %d, abstractSyntax: %q, transferSyntaxes: [", presContext.id, presContext.abstractSyntax)

		// print the list of transfer syntaxes
		for item := presContext.transferSyntaxes.Front(); item != nil; item = item.Next() {
			transferSyntax := item.Value.(string)
			s += fmt.Sprintf("%q,", transferSyntax)
		}

		// trim the trailing comma separating transfer syntaxes and add the enclosing bracket
		s = strings.TrimSuffix(s, ",") + "]"

		// add a closing parenthesis and comma for this presentation context
		s += "},"
	}

	// remove the trainling comma separating presentation contexts and add the enclosing bracket
	s = strings.TrimSuffix(s, ",") + "]"

	// return the fully constructed string
	return s
}

// PresContext represents a presentation context
type PresContext struct {
	id               byte
	abstractSyntax   string
	transferSyntaxes *list.List
}

// UserInfo represents user information
type UserInfo struct {
}

// readAssocRQPDU reads an AssocRQPDU from a reader
func readAssocRQPDU(reader io.Reader) (*AssocRQPDU, error) {

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
	presContexts := list.New()

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

		// if application context item
		if itemType == 0x10 {

			// read the name
			appContextName, err = readUID(reader, uint32(length))
			if err != nil {
				return nil, err
			}

			// if presentation context item
		} else if itemType == 0x020 {

			limitedReader := io.LimitReader(reader, int64(length))

			pcid, err := readByte(limitedReader)
			if err != nil {
				return nil, err
			}

			if err := skipBytes(limitedReader, 3); err != nil {
				return nil, err
			}

			abstractSyntax := ""

			transferSyntaxes := list.New()

			// read the abstract syntax and transfer syntax items
			for {

				// read a sub item
				itemType, err := readByte(limitedReader)
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

				if itemType == 0x30 {

					// abstract syntax

					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					abstractSyntax, err = readUID(limitedReader, uint32(length))
					if err != nil {
						return nil, err
					}

				} else if itemType == 0x40 {

					// transfer syntax

					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					transferSyntax, err := readUID(limitedReader, uint32(length))
					if err != nil {
						return nil, err
					}
					transferSyntaxes.PushBack(transferSyntax)

				} else {

					// unrecognized item
					return nil, fmt.Errorf("unrecognized sub item type: 0x%2x", itemType)

				}

			}

			presContext := &PresContext{pcid, abstractSyntax, transferSyntaxes}
			presContexts.PushBack(presContext)

		} else if itemType == 0x050 {

			// just skip the bytes for now
			if err := skipBytes(reader, uint32(length)); err != nil {
				return nil, err
			}

		} else {

			// unrecognized item
			return nil, fmt.Errorf("unrecognized item type: 0x%2x", itemType)
		}

	}

	return &AssocRQPDU{protocol, calledAETitle, callingAETitle, appContextName, presContexts, userInfo}, nil
}
