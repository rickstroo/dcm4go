package dcm4go

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
)

// item types and sub item types
const (
	appContextItemType      = 0x10
	rqPCItemType            = 0x20
	acPCItemType            = 0x21
	abstractSyntaxItemType  = 0x30
	transferSyntaxItemType  = 0x40
	userInfoItemType        = 0x50
	maxLengthItemType       = 0x51
	implClassUIDItemType    = 0x52
	maxNumOpsItemType       = 0x53
	implVersionNameItemType = 0x55
)

// assocACRQPDU presents an Association Accept and an Association Request PDU
type assocACRQPDU struct {
	protocol       uint16
	calledAETitle  string
	callingAETitle string
	appContextName string
	pcs            []*pc
	userInfo       *UserInfo
}

// String returns a string representation of a AssocACRQPDU
func (pdu *assocACRQPDU) String() string {
	return fmt.Sprintf(
		"{protocol:%v,calledAET:%q,callingAET:%q,appContextName:%q,pcs:%s,userInfo:%s}",
		pdu.protocol,
		strings.TrimSpace(pdu.calledAETitle),
		strings.TrimSpace(pdu.callingAETitle),
		pdu.appContextName,
		pdu.pcs,
		pdu.userInfo)
}

// readAssocRACQPDU reads an association accept or request from the reader
func readAssocACRQPDU(reader io.Reader, pcItemType byte) (*assocACRQPDU, error) {

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
	pcs := make([]*pc, 0, 5)

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

		if itemType == appContextItemType { // application context name

			// read the application context name
			appContextName, err = readUID(reader, uint32(length))
			if err != nil {
				return nil, err
			}

		} else if itemType == pcItemType { // presentation context item type

			// create a limited reader for the requested presentation contextx
			limitedReader := io.LimitReader(reader, int64(length))

			// read the presentation context
			pc, err := readPC(limitedReader, itemType)
			if err != nil {
				return nil, err
			}

			// add it to the list of requested presentation contexts
			pcs = append(pcs, pc)

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
	return &assocACRQPDU{
			protocol,
			calledAETitle,
			callingAETitle,
			appContextName,
			pcs,
			userInfo,
		},
		nil
}

// writeAssocACRQPDU writes an associate request or accept
func writeAssocACRQPDU(writer io.Writer, assocACRQPDU *assocACRQPDU, pduType byte, pcItemType byte) error {

	// write pdu type
	if err := writeByte(writer, pduType); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// create a byte array output stream so we can calculate the length of the rest of the PDU
	byteWriter := new(bytes.Buffer)

	// write the protocol version
	if err := writeShort(byteWriter, assocACRQPDU.protocol, binary.BigEndian); err != nil {
		return err
	}

	// write a short zero
	if err := writeShort(byteWriter, 0x00, binary.BigEndian); err != nil {
		return err
	}

	// write the called ae title
	if err := writeString(byteWriter, padAETitle(assocACRQPDU.calledAETitle)); err != nil {
		return err
	}

	// write the calling ae title
	if err := writeString(byteWriter, padAETitle(assocACRQPDU.callingAETitle)); err != nil {
		return err
	}

	// write thirty two zeroes, zero is the initial value for arrays, so this works
	var zeros [32]byte
	if err := writeBytes(byteWriter, zeros[:]); err != nil {
		return err
	}

	// write the variable items
	if err := writeVariableItems(byteWriter, assocACRQPDU, pcItemType); err != nil {
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

func padAETitle(aeTitle string) string {
	return fmt.Sprintf("%-16s", aeTitle)
}

func writeVariableItems(writer io.Writer, assocACRQPDU *assocACRQPDU, itemType byte) error {

	if err := writeAppContextName(writer, assocACRQPDU.appContextName); err != nil {
		return err
	}

	if err := writePCs(writer, assocACRQPDU.pcs, itemType); err != nil {
		return err
	}

	if err := writeUserInfo(writer, assocACRQPDU.userInfo); err != nil {
		return err
	}

	return nil
}

func writeAppContextName(writer io.Writer, appContextName string) error {

	// write item type
	if err := writeByte(writer, appContextItemType); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}
	// write the length of the application context name
	if err := writeShort(writer, uint16(len(appContextName)), binary.BigEndian); err != nil {
		return err
	}

	// write the application context name
	if err := writeString(writer, appContextName); err != nil {
		return err
	}

	// all is well
	return nil
}
