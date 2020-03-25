// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"bytes"
	"fmt"
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

// and now, let's do the pDataReader
// its a little different because we are going to use it to read PDVs
// perhaps we'll make this more symetric at some point in the future

type pDataReader interface {
	nextPDV() (*pdv, error)
}

type pDataNetworkReader struct {
	pduReader *pduReader
}

func (reader *pDataNetworkReader) nextPDV() (*pdv, error) {
	pdv, err := readPDV(reader.pduReader)
	if err != nil {
		return nil, err
	}

	//	fmt.Printf("after readPDV, err is %v\n", err)

	// again, need some logic to handle an error at this point
	if err != nil {

		// if the error is not eof, return it
		if err != io.EOF {

			//			fmt.Printf("hmm, err was not EOF, that's a problem\n")

			return nil, err
		}

		//		fmt.Printf("need to read another pdu\n")

		// otherwise, it means that we've reached the end of the PDU
		// and we need to read another one from the underlying reader
		pdu, err := reader.pduReader.nextPDU()

		//		fmt.Printf("after readPDU, err is %v\n", err)

		// not expecting any errors at this point
		if err != nil {
			//		fmt.Printf("hmm, err was %v, that's a problem\n", err)
			return nil, err
		}

		//	fmt.Printf("after readPDU, pdu is %v\n", pdu)

		// check that it is data PDU
		if pdu.typ != pDataTFPDU {
			return nil, fmt.Errorf("expecting a pdu of type %d, read a pdu of type %d", pDataTFPDU, pdu.typ)
		}

		// remember the pdu that we've read
		// nope, we don't need to remember that any more as
		// the pdu reader remembers that for us
		//pDataReader.pdu = pdu

		//	fmt.Printf("and we will try the read again")

		// try again
		return reader.nextPDV()
	}

	return pdv, nil
}

type pDataMachineReader struct {
	sp *serviceProvider
}

func (reader *pDataMachineReader) nextPDV() (*pdv, error) {
	pdv, ok := <-reader.sp.pdvInputChan
	if !ok {
		return nil, fmt.Errorf("error while reading from channel")
	}
	return pdv, nil
}
