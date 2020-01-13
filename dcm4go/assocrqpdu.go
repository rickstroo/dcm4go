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
	presContexts   []*PresContext
	userInfo       *UserInfo
}

// PresContext represents a presentation context
type PresContext struct {
	id               byte
	abstractSyntax   string
	transferSyntaxes []string
}

// UserInfo represents user information
type UserInfo struct {
	maxLenReceived     uint32
	implClassUID       string
	implVersionName    string
	maxNumOpsInvoked   uint16
	maxNumOpsPerformed uint16
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
func (presContext *PresContext) String() string {
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

// String returns a string representation of a UserInfo
func (userInfo *UserInfo) String() string {
	return fmt.Sprintf(
		"{maxLenReceived:%v,implClassUID:%q,implVersionName:%q,maxNumOpsInvoked:%v,maxNumOpsPerformed:%v}",
		userInfo.maxLenReceived,
		userInfo.implClassUID,
		userInfo.implVersionName,
		userInfo.maxNumOpsInvoked,
		userInfo.maxNumOpsPerformed)
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

			id, err := readByte(limitedReader)
			if err != nil {
				return nil, err
			}

			if err := skipBytes(limitedReader, 3); err != nil {
				return nil, err
			}

			abstractSyntax := ""

			transferSyntaxes := make([]string, 0)

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

				if subItemType == 0x30 {

					// abstract syntax

					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					abstractSyntax, err = readUID(limitedReader, uint32(length))
					if err != nil {
						return nil, err
					}

				} else if subItemType == 0x40 {

					// transfer syntax

					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					// read the transfer syntax
					transferSyntax, err := readUID(limitedReader, uint32(length))
					if err != nil {
						return nil, err
					}

					// add it to the list
					transferSyntaxes = append(transferSyntaxes, transferSyntax)

				} else {

					// unrecognized item
					return nil, fmt.Errorf("unrecognized presentation context sub item type: 0x%02X", subItemType)

				}

			}

			// create the presentation context
			presContext := &PresContext{id, abstractSyntax, transferSyntaxes}

			// add it to the list
			presContexts = append(presContexts, presContext)

		} else if itemType == 0x050 {

			// create a limited reader for the user info
			limitedReader := io.LimitReader(reader, int64(length))

			// initialze a user info
			userInfo = &UserInfo{}

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

				if subItemType == 0x51 { // maximum length

					// read the length
					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					// check it
					if length != 0x04 {
						return nil, fmt.Errorf("expected length to be 0x04, was 0x%04X", length)
					}

					// read the maximum length received
					maxLenReceived, err := readLong(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}
					userInfo.maxLenReceived = maxLenReceived

				} else if subItemType == 0x52 { // implementation class UID

					// read the length
					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					// read the implementation class UID
					implClassUID, err := readUID(limitedReader, uint32(length))
					if err != nil {
						return nil, err
					}
					userInfo.implClassUID = implClassUID

				} else if subItemType == 0x53 { // maximum number operations

					// read the length
					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					// check it
					if length != 0x04 {
						return nil, fmt.Errorf("expected length to be 0x04, was 0x%04X", length)
					}

					// read the maximum number of operations invoked
					maxNumOpsInvoked, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}
					userInfo.maxNumOpsInvoked = maxNumOpsInvoked

					// read the maximum number of operations performed
					maxNumOpsPerformed, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}
					userInfo.maxNumOpsPerformed = maxNumOpsPerformed

				} else if subItemType == 0x55 { // implementation version name

					// read the length
					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					// read the implementation version name
					implVersionName, err := readText(limitedReader, uint32(length))
					if err != nil {
						return nil, err
					}
					userInfo.implVersionName = implVersionName

				} else {

					// unrecognized item
					fmt.Printf("ignoring unrecognized user info sub item type: 0x%02X\n", subItemType)

					// read the length
					length, err := readShort(limitedReader, binary.BigEndian)
					if err != nil {
						return nil, err
					}

					// skip the bytes
					if err := skipBytes(limitedReader, uint32(length)); err != nil {
						return nil, err
					}
				}

			}

		} else {

			// unrecognized item
			return nil, fmt.Errorf("unrecognized item type: 0x%02X", itemType)
		}

	}

	return &AssocRQPDU{protocol, calledAETitle, callingAETitle, appContextName, presContexts, userInfo}, nil
}
