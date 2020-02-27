// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// pdu represents a DICOM protocol data unit (i.e. PDU)
type pdu struct {
	pduType     byte
	pduLength   uint32
	limitReader io.Reader
	buffer      bytes.Buffer
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

// Read implements the Reader interface
func (pdu *pdu) Read(buf []byte) (int, error) {
	return pdu.limitReader.Read(buf)
}

// Write implements the Writer interface
func (pdu *pdu) Write(buf []byte) (int, error) {
	return pdu.buffer.Write(buf)
}

// String returns a string representation of a PDU
func (pdu *pdu) String() string {
	return fmt.Sprintf("{pduType:%v,pduLength:%v}", pdu.pduType, pdu.pduLength)
}

// readPDU reads a PDU from a reader
func readPDU(reader io.Reader) (*pdu, error) {

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

	// set up the a reader for the bytes of the pdu
	limitReader := io.LimitReader(reader, int64(pduLength))

	// construct a PDU
	pdu := &pdu{
		pduType:     pduType,
		pduLength:   pduLength,
		limitReader: limitReader,
	}

	// return the pdu
	return pdu, nil
}

func writePDU(writer io.Writer, pdu *pdu) error {
	if err := writeByte(writer, pdu.pduType); err != nil {
		return err
	}
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}
	if err := writeLong(writer, (uint32)(pdu.buffer.Len()), binary.BigEndian); err != nil {
		return err
	}
	if err := writeBytes(writer, pdu.buffer.Bytes()); err != nil {
		return err
	}
	return nil
}
