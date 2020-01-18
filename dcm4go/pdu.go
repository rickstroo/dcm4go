package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PDU represents a DICOM protocol data unit (i.e. PDU)
type PDU struct {
	pduType   byte
	pduLength uint32
}

const (
	aAssociateRQPDU = 0x01
	aAssociateACPDU = 0x02
	aAssociateRJPDU = 0x03
	pDataTFPDU      = 0x04
	aReleaseRQPDU   = 0x05
	aReleaseRPPDU   = 0x06
	aAbortPDU       = 0x07
)

// String returns a string representation of a PDU
func (pdu *PDU) String() string {
	return fmt.Sprintf("pduType: %v, pduLength: %v", pdu.pduType, pdu.pduLength)
}

// readPDU reads a PDU from a reader
func readPDU(reader io.Reader) (*PDU, error) {

	// get the pdu type
	pduType, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// skip a byte, as per the standard
	if err := skipByte(reader); err != nil {
		return nil, err
	}

	// get the length, PDU lengths are always big endian
	pduLength, err := readLong(reader, binary.BigEndian)
	if err != nil {
		return nil, err
	}

	return &PDU{pduType, pduLength}, nil
}

func writePDU(writer io.Writer, pdu *PDU) error {
	if err := writeByte(writer, pdu.pduType); err != nil {
		return err
	}
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}
	if err := writeLong(writer, pdu.pduLength, binary.BigEndian); err != nil {
		return err
	}
	return nil
}
