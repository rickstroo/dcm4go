package dcm4go

import (
	"encoding/binary"
	"io"
)

// ReleaseRPPDU represents a release response PDU
type ReleaseRPPDU struct{}

// readReleaseRPPDU reads an readReleaseRPPDU from a reader
func readReleaseRPPDU(reader io.Reader) error {

	// read and ignore the long
	if _, err := readLong(reader, binary.BigEndian); err != nil {
		return err
	}

	return nil
}

// Write writes a release response PDU
func (releaseRPPDU *ReleaseRPPDU) Write(writer io.Writer) error {

	// construct the abort pdu
	buf := []byte{
		0x00, // reserved
		0x00, // reserved
		0x00, // reserved
		0x00, // reserved
	}

	// construct the base pdu
	pdu := &PDU{
		pduType:   aReleaseRPPDU,    // the type
		pduLength: uint32(len(buf)), // the length
	}

	// write the base pdu
	if err := pdu.Write(writer); err != nil {
		return err
	}

	// write the release pdu
	if err := writeBytes(writer, buf[:]); err != nil {
		return err
	}

	return nil
}
