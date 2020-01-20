package dcm4go

import (
	"fmt"
	"io"
)

// PDataReader is used to manage the internal state of a reader that reads
// the stream of bytes from one or more PDVs that can be contained in
// one or more PDUs.
type PDataReader struct {
	reader    io.Reader // the underlying reader
	pduReader io.Reader // the reader for the PDU
	pdvReader io.Reader // the reader for the PDV
	lastPDV   *PDV      // the last PDV that was read
}

// newPDataReader constructs and initializes a PDataReader
func newPDataReader(reader io.Reader, pdu *PDU) (*PDataReader, error) {

	// set up the pdu reader
	// use a limited reader for the length of the PDU
	pduReader := io.LimitReader(reader, int64(pdu.pduLength))

	// read the first PDV
	pdv, err := readPDV(pduReader)
	if err != nil {
		return nil, err
	}

	// set up the pdv reader
	// use a limited reader for the length of the PDV
	// actually, we set it to the length less two to make up for the pcid and mch
	// we should really fix that so that the pdv length field reflects that
	// otherwise we have to remember that everywhere
	pdvReader := io.LimitReader(reader, int64(pdv.pdvLength-2))

	// construct a reader and return
	return &PDataReader{reader, pduReader, pdvReader, pdv}, nil
}

// Read implements the Reader interface
func (pDataReader *PDataReader) Read(buf []byte) (int, error) {

	// attempt to read some bytes
	num, err := pDataReader.pdvReader.Read(buf)

	// if we didn't ready an byte and if an error occured,
	// we need lots of logic to handle it
	if num == 0 && err != nil {

		// if an error occured and it was not end of file, return it
		if err != io.EOF {
			return num, err
		}

		// otherwise, check to see if this is the last PDV,
		// and if it is, return EOF
		if pDataReader.lastPDV.isLast() {
			return num, io.EOF
		}

		// it's not the last, so we read another pdv
		if err := pDataReader.nextPDV(); err != nil {
			return num, err
		}

		// try the read again
		// we can use the buf that was passed originally because
		// we checked earlier that we didn't read any bytes
		return pDataReader.Read(buf)
	}

	// return the number of bytes read
	return num, err
}

// read the next PDV
func (pDataReader *PDataReader) nextPDV() error {

	// it's not the last, so we read another pdv
	pdv, err := readPDV(pDataReader.pduReader)

	// again, need some logic to handle an error at this point
	if err != nil {

		// if the error is not eof, return it
		if err != io.EOF {
			return err
		}

		// otherwise, it means that we've reached the end of the PDU
		// and we need to read another one
		pdu, err := readPDU(pDataReader.reader)

		// not expecting any errors at this point
		if err != nil {
			return err
		}

		// check that it is data PDU
		if pdu.pduType != pDataTFPDU {
			return fmt.Errorf("expecting a pdu of type %d, read a pdu of type %d", pDataTFPDU, pdu.pduType)
		}

		// reset the pdu reader
		pDataReader.pduReader = io.LimitReader(pDataReader.reader, int64(pdu.pduLength))

		// try again
		return pDataReader.nextPDV()
	}

	// check that the presentation context ids match
	if pdv.pcID != pDataReader.lastPDV.pcID {
		return fmt.Errorf(
			"presentation context id for next pdv, %d, does not match presentation id for last pdv, %d",
			pdv.pcID,
			pDataReader.lastPDV.pcID)
	}

	// check that the command or data match the last pdv
	if pdv.isCommand() {
		if !pDataReader.lastPDV.isCommand() {
			return fmt.Errorf(
				"next PDV is a command while last PDV was a data set")
		}
	} else {
		if pDataReader.lastPDV.isCommand() {
			return fmt.Errorf(
				"next PDV is a data set while last PDV was a command")
		}
	}

	// the new pdv is now the last pdv
	pDataReader.lastPDV = pdv

	// create a new limit reader
	pDataReader.pdvReader = io.LimitReader(pDataReader.reader, int64(pdv.pdvLength-2))

	// all is well
	return nil
}
