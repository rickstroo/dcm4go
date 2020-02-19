package dcm4go

import (
	"io"
)

const (
	sourceServiceUserInitiatedAbort     = 0x00
	sourceServiceProviderInitiatedAbort = 0x02
)

const (
	reasonNotSpecified              = 0x00
	reasonUnrecognizedPDU           = 0x01
	reasonUnexpectedPDU             = 0x02
	reasonUnrecognizedPDUParameter  = 0x04
	reasonUnexpectedPDUParameter    = 0x05
	reasonInvalidPDUParamaeterValue = 0x06
)

// An AbortPDU represents a PDU used to abort associations
type AbortPDU struct {
	source byte // the initiator of the abort
	reason byte // the reason for the abort
}

// ReadAbortPDU reads an AbortPDU
func readAbortPDU(reader io.Reader) (*AbortPDU, error) {

	// read the abort pdu
	buf, err := readBytes(reader, 4)
	if err != nil {
		return nil, err
	}

	// construct the abort pdu
	// 1st and 2nd bytes are reserved
	// 3rd byte is the source
	// 3th byte is the reason
	abortPDU := &AbortPDU{
		source: buf[2],
		reason: buf[3],
	}

	return abortPDU, nil
}

// Write writes an AbortPDU to a writer
func (abortPDU *AbortPDU) Write(writer io.Writer) error {

	// construct the abort pdu
	buf := []byte{
		0x00,            // reserved
		0x00,            // reserved
		abortPDU.source, // the source
		abortPDU.reason, // the reason
	}

	// construct the base pdu
	pdu := &PDU{
		pduType:   aAbortPDU,        // the type
		pduLength: uint32(len(buf)), // the length
	}

	// write the base pdu
	if err := pdu.Write(writer); err != nil {
		return err
	}

	// write the abort pdu
	if err := writeBytes(writer, buf[:]); err != nil {
		return err
	}

	return nil
}
