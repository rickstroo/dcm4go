package dcm4go

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// AssocACRQPDU presents an Association Accept and an Association Request PDU
type AssocACRQPDU struct {
	protocol       uint16
	calledAETitle  string
	callingAETitle string
	appContextName string
	presContexts   []*PresContext
	userInfo       *UserInfo
}

// String returns a string representation of a AssocACRQPDU
func (pdu *AssocACRQPDU) String() string {
	return fmt.Sprintf(
		"{protocol:%v,calledAET:%q,callingAET:%q,appContextName:%q,presContexts:%s,userInfo:%s}",
		pdu.protocol,
		strings.TrimSpace(pdu.calledAETitle),
		strings.TrimSpace(pdu.callingAETitle),
		pdu.appContextName,
		pdu.presContexts,
		pdu.userInfo)
}

func writeVariableItems(writer io.Writer, appContextName string, presContexts []*PresContext, itemType byte, userInfo *UserInfo) error {

	if err := writeAppContextName(writer, appContextName); err != nil {
		return err
	}

	if err := writePresContexts(writer, presContexts, itemType); err != nil {
		return err
	}

	if err := writeUserInfo(writer, userInfo); err != nil {
		return err
	}

	return nil
}

func writeAppContextName(writer io.Writer, appContextName string) error {

	// write item type
	if err := writeByte(writer, 0x10); err != nil {
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

func writePresContexts(writer io.Writer, presContexts []*PresContext, itemType byte) error {

	// for each of the  presentation contexts
	for _, presContext := range presContexts {

		// write it
		if err := writePresContext(writer, presContext, itemType); err != nil {
			return err
		}
	}

	// all is well
	return nil
}

func writePresContext(writer io.Writer, presContext *PresContext, itemType byte) error {

	// write item type
	if err := writeByte(writer, itemType); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// create a byte array output stream so we can calculate the length of the rest of the PDU
	byteWriter := new(bytes.Buffer)

	// write the presentation context id
	if err := writeByte(byteWriter, presContext.id); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(byteWriter, 0x00); err != nil {
		return err
	}

	// write the result
	if err := writeByte(byteWriter, presContext.result); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(byteWriter, 0x00); err != nil {
		return err
	}

	// write the transfer syntax
	if err := writeTransferSyntax(byteWriter, presContext.transferSyntaxes[0]); err != nil {
		return err
	}

	// write the length to the original writer
	if err := writeShort(writer, uint16(byteWriter.Len()), binary.BigEndian); err != nil {
		return err
	}

	// write the byte array to the original writer
	if err := writeBytes(writer, byteWriter.Bytes()); err != nil {
		return err

	}

	// all is good
	return nil
}

func writeTransferSyntax(writer io.Writer, transferSyntax string) error {

	// write sub item type
	if err := writeByte(writer, 0x40); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// write the length of the application context name
	if err := writeShort(writer, uint16(len(transferSyntax)), binary.BigEndian); err != nil {
		return err
	}

	// write the application context name
	if err := writeString(writer, transferSyntax); err != nil {
		return err
	}

	// all is well
	return nil
}
