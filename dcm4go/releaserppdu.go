package dcm4go

import (
	"encoding/binary"
	"io"
)

// readReleaseRPPDU reads an releaseRPPDU from a reader
func readReleaseRPPDU(reader io.Reader) error {

	// read and ignore the long
	if _, err := readLong(reader, binary.BigEndian); err != nil {
		return err
	}

	return nil
}

// writeReleaseRPPDU writes a release response PDU
func writeReleaseRPPDU(writer io.Writer) error {

	// construct the release response pdu
	buf := []byte{
		0x00, // reserved
		0x00, // reserved
		0x00, // reserved
		0x00, // reserved
	}

	// construct the base pdu
	pdu := &pdu{typ: aReleaseRPPDU}

	// write the buffer to the pdu
	if err := writeBytes(pdu, buf[:]); err != nil {
		return err
	}

	// write the pdu
	if err := writePDU(writer, pdu); err != nil {
		return err
	}

	return nil
}
