package dcm4go

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type serviceProvider struct {
	aeTitle      string
	capabilities *Capabilities
	pdvInputChan chan *pdv
}

func newServiceProvider(aeTitle string, capabilities *Capabilities) *serviceProvider {
	sp := &serviceProvider{
		aeTitle:      aeTitle,
		capabilities: capabilities,
		pdvInputChan: make(chan *pdv, 1),
	}
	mr := messageReader{
		sp: sp,
	}
	go mr.run()
	return sp
}

func (sp *serviceProvider) onAssociateRQ(p *pdu) (*pdu, error) {

	// make a reader
	pr := bytes.NewReader(p.buf)

	// read the associate request
	assocRQPDU, err := readAssocRQPDU(pr)
	if err != nil {
		return nil, err
	}
	log.Printf("assocRQPDU is %v\n", assocRQPDU)

	// attempt to negotiate an association
	assocACPDU, assocRJPDU, err := negotiateAssoc(assocRQPDU, sp.aeTitle, sp.capabilities)
	if err != nil {
		return nil, err
	}
	if assocACPDU != nil {
		log.Printf("accepted associate request, assocACPDU is %v\n", assocACPDU)
		pdu, err := createAssocACPDU(assocACPDU)
		if err != nil {
			return nil, err
		}
		return pdu, nil
	}
	if assocRJPDU != nil {
		log.Printf("rejected associate request, assocRJPDU is %v\n", assocRJPDU)
		pdu, err := createAssocRJPDU(assocRJPDU)
		if err != nil {
			return nil, err
		}
		return pdu, nil
	}

	return nil, fmt.Errorf("didn't accept or reject association request, hmm, that's weird")
}

func (sp *serviceProvider) onDataTF(p *pdu) error {

	// make a reader
	pr := bytes.NewReader(p.buf)

	// read the data transfer
	dataTFPDU, err := readDataTFPDU(pr)
	if err != nil {
		return err
	}

	// parse out the pdvs
	for _, pdv := range dataTFPDU.pdvs {
		log.Printf("sp putting pdv on pdv input channel, %v", pdv)
		sp.pdvInputChan <- pdv
	}

	// return success
	return nil
}

type messageReader struct {
	sp *serviceProvider
}

func (mr *messageReader) run() {

	for {

		// create a reader for the command
		// start by reading from the pdu that was provided
		log.Printf("message reader waiting for pdv reader")
		commandReader, err := newPDVReader1(mr.sp, true)
		if err != nil {
			log.Printf("while creating pdv reader, error was %v", err)
			continue
		}

		// get the presentation context id from the reader
		pcID := commandReader.pdv.pcID

		// read the command from the pdu
		command, err := readCommand(commandReader)
		if err != nil {
			log.Printf("while reading command, error was %v", err)
			continue
		}

		log.Printf("pc id is %v", pcID)
		log.Printf("command is %v", command)
	}
}

// pdvReader is used to manage the internal state of a reader that reads
// the stream of bytes from one or more PDVs that can be contained in
// one or more PDUs.
type pdvReader1 struct {
	sp         *serviceProvider // the underlying pdu reader
	pdv        *pdv             // the PDV that we are reading from
	byteReader *bytes.Reader    // a reader for the bytes of the PDV
	isCommand  bool             // are we reading a command or a data set?
}

// newPDataReader constructs and initializes a PDataReader
// the pdu reader is used to read additional pdus if required
// isCommand indicates whether we are reading a command or data
func newPDVReader1(sp *serviceProvider, isCommand bool) (*pdvReader1, error) {

	// read the first PDV
	pdv, ok := <-sp.pdvInputChan
	if !ok {
		return nil, fmt.Errorf("error while reading from channel")
	}

	// create a reader
	byteReader := bytes.NewReader(pdv.buf)

	// check that the command or data match the last pdv
	if err := checkCommand(isCommand, pdv); err != nil {
		return nil, err
	}

	// construct a reader
	pdvReader := &pdvReader1{
		sp:         sp,
		pdv:        pdv,
		byteReader: byteReader,
		isCommand:  isCommand,
	}

	// return the pdv reader and success
	return pdvReader, nil
}

// Read implements the Reader interface
func (pr *pdvReader1) Read(buf []byte) (int, error) {

	// attempt to read some bytes
	num, err := pr.byteReader.Read(buf)

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
		if pr.pdv.isLast() {
			//			fmt.Printf("reached EOF on last PDV, so we are done\n")
			return num, io.EOF
		}

		//		fmt.Printf("reached EOF on PDV, not last, so we need to read another PDV\n")

		// it's not the last, so we read another pdv
		if err := pr.nextPDV(); err != nil {
			return num, err
		}

		// try the read again
		// we can use the buf that was passed originally because
		// we checked earlier that we didn't read any bytes
		return pr.Read(buf)
	}

	// return the number of bytes read
	return num, err
}

// read the next PDV
func (pr *pdvReader1) nextPDV() error {

	// it's not the last, so we read another pdv
	pdv, ok := <-pr.sp.pdvInputChan
	if !ok {
		return fmt.Errorf("error while reading from channel")
	}

	// check that the presentation context ids match
	if pdv.pcID != pr.pdv.pcID {
		return fmt.Errorf(
			"presentation context id for next pdv, %d, does not match presentation id for last pdv, %d",
			pdv.pcID,
			pr.pdv.pcID)
	}

	// check that the command or data match the last pdv
	if err := checkCommand1(pr.isCommand, pdv); err != nil {
		return err
	}

	// the new pdv is now the last pdv
	pr.pdv = pdv

	// all is well
	return nil
}

func checkCommand1(isCommand bool, pdv *pdv) error {

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
