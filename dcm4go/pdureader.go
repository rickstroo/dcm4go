// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"bytes"
	"io"
)

// A pduReader is used to read the contents of a raw pdu,
// much like a bytes.Reader or bytes.Buffer is used to read
// the contents of a raw bytes buffer.  In addition, the
// pduReader has the ability to read the next pdu, as data
// can be distributed across multiple pdus.  In this way,
// a pdu reader acts like a multiple reader.
type pduReader struct {
	reader     io.Reader     // underlying reader
	pdu        *pdu          // current pdu
	byteReader *bytes.Reader // reader for current pdu
}

// newPDUReader constructs a new pduReader
func newPDUReader(reader io.Reader) *pduReader {
	return &pduReader{reader: reader}
}

// nextPDU reads the next pdu from the underlying reader
func (pduReader *pduReader) nextPDU() (*pdu, error) {

	// read the pdu
	pdu, err := readPDU(pduReader.reader)
	if err != nil {
		return nil, err
	}

	// remember the pdu
	pduReader.pdu = pdu

	// create a reader
	pduReader.byteReader = bytes.NewReader(pdu.buf)

	// return the pdu so that we can inspect the type
	// perhaps we should just return the type
	return pdu, nil
}

// Read implements the Reader interface
func (pduReader *pduReader) Read(buf []byte) (int, error) {
	return pduReader.byteReader.Read(buf)
}
