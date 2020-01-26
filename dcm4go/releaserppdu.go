package dcm4go

import (
	"encoding/binary"
	"io"
)

// readReleaseRPPDU reads an readReleaseRPPDU from a reader
func readReleaseRPPDU(reader io.Reader) error {

	// read and ignore the long
	if _, err := readLong(reader, binary.BigEndian); err != nil {
		return err
	}

	return nil
}

// writeReleaseRPPDU writes a release response PDU to a writer
func writeReleaseRPPDU(writer io.Writer) error {

	// construct a pdu
	pdu := &PDU{aReleaseRPPDU, 0x04, nil}

	// write the pdu header
	if err := writePDU(writer, pdu); err != nil {
		return err
	}

	// write a long zero
	if err := writeLong(writer, 0x00, binary.BigEndian); err != nil {
		return err
	}

	return nil
}
