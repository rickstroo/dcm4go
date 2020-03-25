package dcm4go

import (
	"bytes"
	"fmt"
	"log"
)

type serviceProvider struct {
	machine      *machine
	assoc        *Assoc
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

		// create an association from the response and save it in the sp
		// only need the assoc rq and assoc ac
		assoc := &Assoc{
			assocRQPDU: assocRQPDU,
			assocACPDU: assocACPDU,
		}
		sp.assoc = assoc

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

func (sp *serviceProvider) onReleaseRQ(p *pdu, c bool) error {
	d := &deed{e: evt14}
	sp.machine.deedChan <- d
	return nil
}

func (sp *serviceProvider) onCEcho(request *Message) error {
	log.Printf("received DICOM C-Echo request, request is %v", request)

	// create a response
	response, err := NewCEchoResponse(request)
	if err != nil {
		return err
	}

	// write the response
	if err := sp.WriteResponse(response); err != nil {
		return err
	}

	// all is well
	return nil

}

func (sp *serviceProvider) WriteResponse(response *Message) error {
	return sp.writeMessage(response)
}

func (sp *serviceProvider) writeMessage(message *Message) error {
	// create a writer for the machine
	pDataWriter := &pDataMachineWriter{machine: sp.machine}
	// write the message
	return writeMessage(sp.assoc, message, pDataWriter)
}

type messageReader struct {
	sp *serviceProvider
}

func (mr *messageReader) run() {

	for {

		// create pDataReader
		pDataReader := &pDataMachineReader{sp: mr.sp}

		message, err := readMessage(mr.sp.assoc, false, pDataReader)
		if err != nil {
			log.Printf("while reading message, caught error %v", err)
			continue
		}

		commandField, err := message.command.AsShort(CommandFieldTag, 0)
		if err != nil {
			log.Printf("error while getting command field, %v", err)
			continue
		}
		switch commandField {
		case CEchoRQ:
			mr.sp.onCEcho(message)
		default:
			log.Printf("unrecognized command, %v", message)
			continue
		}
	}
}
