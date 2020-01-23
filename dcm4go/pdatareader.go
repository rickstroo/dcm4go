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
	pdu       *PDU      // the PDU we are reading from
	pdv       *PDV      // the PDV that we are reading from
	isCommand bool      // are we reading a command or a data set?
}

// newPDataReader constructs and initializes a PDataReader
// notice that we don't pass in a pdu, but a pdu reader
// that's because a single pdu reader can be used to read a
// sequence of command PDVs and a sequence of data set PDVs
func newPDataReader(reader io.Reader, pdu *PDU, isCommand bool) (*PDataReader, error) {

	// read the first PDV
	pdv, err := readPDV(pdu)
	if err != nil {
		return nil, err
	}

	// check that the command or data match the last pdv
	if err := checkCommand(isCommand, pdv); err != nil {
		return nil, err
	}

	// construct a reader and return
	return &PDataReader{reader, pdu, pdv, isCommand}, nil
}

// Read implements the Reader interface
func (pDataReader *PDataReader) Read(buf []byte) (int, error) {

	// attempt to read some bytes
	num, err := pDataReader.pdv.Read(buf)

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
		if pDataReader.pdv.isLast() {
			//			fmt.Printf("reached EOF on last PDV, so we are done\n")
			return num, io.EOF
		}

		//		fmt.Printf("reached EOF on PDV, not last, so we need to read another PDV\n")

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
	pdv, err := readPDV(pDataReader.pdu)

	//	fmt.Printf("after readPDV, err is %v\n", err)

	// again, need some logic to handle an error at this point
	if err != nil {

		// if the error is not eof, return it
		if err != io.EOF {

			//			fmt.Printf("hmm, err was not EOF, that's a problem\n")

			return err
		}

		//		fmt.Printf("need to read another pdu\n")

		// otherwise, it means that we've reached the end of the PDU
		// and we need to read another one from the underlying reader
		pdu, err := readPDU(pDataReader.reader)

		//		fmt.Printf("after readPDU, err is %v\n", err)

		// not expecting any errors at this point
		if err != nil {
			//		fmt.Printf("hmm, err was %v, that's a problem\n", err)
			return err
		}

		//	fmt.Printf("after readPDU, pdu is %v\n", pdu)

		// check that it is data PDU
		if pdu.pduType != pDataTFPDU {
			return fmt.Errorf("expecting a pdu of type %d, read a pdu of type %d", pDataTFPDU, pdu.pduType)
		}

		// remember the pdu that we've read
		pDataReader.pdu = pdu

		//	fmt.Printf("and we will try the read again")

		// try again
		return pDataReader.nextPDV()
	}

	//	fmt.Printf("after reading PDV, pdv is %v\n", pdv)

	// check that the presentation context ids match
	if pdv.pcID != pDataReader.pdv.pcID {
		return fmt.Errorf(
			"presentation context id for next pdv, %d, does not match presentation id for last pdv, %d",
			pdv.pcID,
			pDataReader.pdv.pcID)
	}

	// check that the command or data match the last pdv
	if err := checkCommand(pDataReader.isCommand, pdv); err != nil {
		return err
	}

	// the new pdv is now the last pdv
	pDataReader.pdv = pdv

	// all is well
	return nil
}

func checkCommand(isCommand bool, pdv *PDV) error {

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
