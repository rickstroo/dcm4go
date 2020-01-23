package dcm4go

import (
	"bytes"
	"encoding/binary"
	"io"
)

// create a new association acceptance PDU
func newAssocACPDU(assocRQPDU *AssocACRQPDU) *AssocACRQPDU {
	return &AssocACRQPDU{
		0x01,                       // protocol version, as per the standard
		assocRQPDU.calledAETitle,   // copy from the request, as per the standard
		assocRQPDU.callingAETitle,  // copy from the request, as per the standard
		"1.2.840.10008.3.1.1.1",    // app context name, as per the standard
		make([]*PresContext, 0, 5), // empty pres context list
		&UserInfo{
			16378,             // max length received, need to figure out why dcm4che uses this number
			"1.2.40.0.13.1.3", // implementation class uid, need to get a root, borrowing dcm4che for now
			"dcm4go-1.0",      // implementation class name
			0,                 // max num ops invoked
			0,                 // max num ops performed
		},
	}
}

func writeAssocACPDU(writer io.Writer, assocACPDU *AssocACRQPDU) error {

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
	if err := writeShort(byteWriter, assocACPDU.protocol, binary.BigEndian); err != nil {
		return err
	}

	// write a short zero
	if err := writeShort(byteWriter, 0x00, binary.BigEndian); err != nil {
		return err
	}

	// write the called ae title
	if err := writeText(byteWriter, assocACPDU.calledAETitle); err != nil {
		return err
	}

	// write the calling ae title
	if err := writeText(byteWriter, assocACPDU.callingAETitle); err != nil {
		return err
	}

	// write thirty two zeroes, zero is the initial value for arrays, so this works
	var zeros [32]byte
	if err := writeBytes(byteWriter, zeros[:]); err != nil {
		return err
	}

	// write the variable items
	if err := writeVariableItems(byteWriter, assocACPDU); err != nil {
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

func writeVariableItems(writer io.Writer, assocACPDU *AssocACRQPDU) error {

	if err := writeAppContextName(writer, assocACPDU); err != nil {
		return err
	}

	if err := writeACPresContexts(writer, assocACPDU); err != nil {
		return err
	}

	if err := writeUserInfo(writer, assocACPDU.userInfo); err != nil {
		return err
	}

	return nil
}

func writeAppContextName(writer io.Writer, assocACPDU *AssocACRQPDU) error {

	// write item type
	if err := writeByte(writer, 0x10); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}
	// write the length of the application context name
	if err := writeShort(writer, uint16(len(assocACPDU.appContextName)), binary.BigEndian); err != nil {
		return err
	}

	// write the application context name
	if err := writeUID(writer, assocACPDU.appContextName); err != nil {
		return err
	}

	// all is well
	return nil
}

func writeACPresContexts(writer io.Writer, assocACPDU *AssocACRQPDU) error {

	// for each of the accepted presentation contexts
	for _, presContext := range assocACPDU.presContexts {

		// write it
		if err := writeACPresContext(writer, presContext); err != nil {
			return err
		}
	}

	// all is well
	return nil
}

func writeACPresContext(writer io.Writer, presContext *PresContext) error {

	// write item type
	if err := writeByte(writer, 0x21); err != nil {
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
	if err := writeUID(writer, transferSyntax); err != nil {
		return err
	}

	// all is well
	return nil
}
