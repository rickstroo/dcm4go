// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"bytes"
	"fmt"
	"io"
)

// pdvReader is used to manage the internal state of a reader that reads
// the stream of bytes from one or more PDVs that can be contained in
// one or more PDUs.
type pdvReader struct {
	pDataReader pDataReader
	pdv         *pdv          // the PDV that we are reading from
	byteReader  *bytes.Reader // a reader for the bytes of the PDV
	isCommand   bool          // are we reading a command or a data set?
}

// newPDataReader constructs and initializes a PDataReader
// the pdu reader is used to read additional pdus if required
// isCommand indicates whether we are reading a command or data
func newPDVReader(pDataReader pDataReader, isCommand bool) (*pdvReader, error) {

	// read the first PDV
	pdv, err := pDataReader.nextPDV()
	if err != nil {
		return nil, err
	}

	// create a reader
	byteReader := bytes.NewReader(pdv.buf)

	// check that the command or data match the last pdv
	if err := checkCommand(isCommand, pdv); err != nil {
		return nil, err
	}

	// construct a reader
	pdvReader := &pdvReader{
		pdv:         pdv,
		byteReader:  byteReader,
		isCommand:   isCommand,
		pDataReader: pDataReader,
	}

	// return the pdv reader and success
	return pdvReader, nil
}

// Read implements the Reader interface
func (pdvReader *pdvReader) Read(buf []byte) (int, error) {

	// attempt to read some bytes
	num, err := pdvReader.byteReader.Read(buf)

	//	fmt.Printf("after read, num is %d and err is %v\n", num, err)

	// if we didn't ready an byte and if an error occured,
	// we need lots of logic to handle it
	if num == 0 && err != nil {

		// if an error occured and it was not end of file, return it
		if err != io.EOF {
			return num, err
		}

		// otherwise, check to see if this is the last PDV,
		// and if it is, return EOF
		if pdvReader.pdv.isLast() {
			//			fmt.Printf("reached EOF on last PDV, so we are done\n")
			return num, io.EOF
		}

		//		fmt.Printf("reached EOF on PDV, not last, so we need to read another PDV\n")

		// it's not the last, so we read another pdv
		if err := pdvReader.nextPDV(); err != nil {
			return num, err
		}

		// try the read again
		// we can use the buf that was passed originally because
		// we checked earlier that we didn't read any bytes
		return pdvReader.Read(buf)
	}

	// return the number of bytes read
	return num, err
}

// read the next PDV
func (pdvReader *pdvReader) nextPDV() error {

	// it's not the last, so we read another pdv
	pdv, err := pdvReader.pDataReader.nextPDV()
	if err != nil {
		return err
	}

	//	fmt.Printf("after reading PDV, pdv is %v\n", pdv)

	// check that the presentation context ids match
	if pdv.pcID != pdvReader.pdv.pcID {
		return fmt.Errorf(
			"presentation context id for next pdv, %d, does not match presentation id for last pdv, %d",
			pdv.pcID,
			pdvReader.pdv.pcID)
	}

	// check that the command or data match the last pdv
	if err := checkCommand(pdvReader.isCommand, pdv); err != nil {
		return err
	}

	// the new pdv is now the last pdv
	pdvReader.pdv = pdv

	// all is well
	return nil
}

func checkCommand(isCommand bool, pdv *pdv) error {

	// check that pdv type matches what is expected
	if isCommand {
		if !pdv.isCommand() {
			return fmt.Errorf("received data set PDV while expecting a command PDV")
		}
	} else {
		if pdv.isCommand() {
			return fmt.Errorf("received command PDV while expecting a data set PDV")
		}
	}

	// all is well
	return nil
}
