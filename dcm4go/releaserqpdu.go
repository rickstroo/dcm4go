package dcm4go

import (
	"encoding/binary"
	"io"
)

// readReleaseRQPDU reads an ReleaseRQPDU from a reader
func readReleaseRQPDU(reader io.Reader) error {

	// read and ignore the long
	if _, err := readLong(reader, binary.BigEndian); err != nil {
		return err
	}

	return nil
}

// writeReleaseRQPDU writes a release request PDU to a writer
func writeReleaseRQPDU(writer io.Writer) error {

	// construct a pdu
	pdu := &pdu{aReleaseRQPDU, 0x04, nil}

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
