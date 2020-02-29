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
	typ byte
	buf *bytes.Buffer
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
	return pdu.buf.Read(buf)
}

// Write implements the Writer interface
func (pdu *pdu) Write(buf []byte) (int, error) {
	// if the write buffer has not been initialized, create one with 16K capacity
	if pdu.buf == nil {
		pdu.buf = bytes.NewBuffer(make([]byte, 0, 16*1024))
	}
	return pdu.buf.Write(buf)
}

// String returns a string representation of a PDU
func (pdu *pdu) String() string {
	return fmt.Sprintf("{typ:%v,buf:%v}", pdu.typ, pdu.buf)
}

// readPDU reads a PDU from a reader
func readPDU(reader io.Reader) (*pdu, error) {

	// get the pdu type
	typ, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// skip a byte, as per the standard
	if err := skipByte(reader); err != nil {
		return nil, err
	}

	// read the length, PDU lengths are always big endian
	length, err := readLong(reader, binary.BigEndian)
	if err != nil {
		return nil, err
	}

	// read the bytes
	buf, err := readBytes(reader, length)
	if err != nil {
		return nil, err
	}

	// construct a PDU
	pdu := &pdu{
		typ: typ,
		buf: bytes.NewBuffer(buf),
	}

	// return the pdu
	return pdu, nil
}

func writePDU(writer io.Writer, pdu *pdu) error {

	// write the type
	if err := writeByte(writer, pdu.typ); err != nil {
		return err
	}

	// skip a byte, as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// write the length
	if err := writeLong(writer, uint32(pdu.buf.Len()), binary.BigEndian); err != nil {
		return err
	}

	// write the bytes
	if err := writeBytes(writer, pdu.buf.Bytes()); err != nil {
		return err
	}

	// return success
	return nil
}
