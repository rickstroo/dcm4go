// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import "io"

type pduReader struct {
	reader io.Reader // underlying reader
	pdu    *pdu      // current pdu
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
	// return the pdu so that we can inspect the type
	// perhaps we should just return the type
	return pdu, nil
}

// Read implements the Reader interface
func (pduReader *pduReader) Read(buf []byte) (int, error) {
	return pduReader.pdu.Read(buf)
}
