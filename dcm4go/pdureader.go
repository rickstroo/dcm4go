// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import "io"

type pduReader struct {
	reader io.Reader // underlying reader
	pdu    *PDU      // current pdu
}

// newPDUReader constructs a new pduReader
func newPDUReader(reader io.Reader) *pduReader {
	return &pduReader{reader: reader}
}

// nextPDU reads the next pdu from the underlying reader
func (pduReader *pduReader) nextPDU() (*PDU, error) {
	pdu, err := readPDU(pduReader.reader)
	if err != nil {
		return nil, err
	}
	return pdu, nil
}

// Read implements the Reader interface
func (pduReader *pduReader) Read(buf []byte) (int, error) {
	return pduReader.pdu.Read(buf)
}
