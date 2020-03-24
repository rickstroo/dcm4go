package dcm4go

import (
	"bytes"
	"fmt"
	"io"
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

	if err := sp.writeCommand(message.pcID, sp.assoc.assocRQPDU.userInfo.maxLenReceived, message.command); err != nil {
		return err
	}

	if message.data != nil {

		transferSyntax, err := sp.assoc.findAcceptedTransferSyntaxByPCID(message.pcID)
		if err != nil {
			return err
		}

		if err := sp.writeData(message.pcID, sp.assoc.assocRQPDU.userInfo.maxLenReceived, message.data, transferSyntax); err != nil {
			return err
		}

	} else if message.dataReader != nil {

		if err := sp.copyDataFromReader(message.pcID, sp.assoc.assocRQPDU.userInfo.maxLenReceived, message.dataReader); err != nil {
			return err
		}
	}

	// return success
	return nil
}

// writeCommand writes the command portion of the message
func (sp *serviceProvider) writeCommand(pcID byte, maxBufLen uint32, command *Object) error {

	// create a writer to write the data to
	pdvWriter := newPDVWriter1(sp.machine, pcID, true, maxBufLen)

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the command with group length
	if err := encoder.writeObjectWithGroupLength(pdvWriter, 0x0000, command, ImplicitVRLittleEndianTS); err != nil {
		return err
	}

	// flush to the underlying writer
	// passing true means we are done writing this object
	if err := pdvWriter.Flush(true); err != nil {
		return err
	}

	// all is well
	return nil
}

// // writeData writes the data portion of the message
func (sp *serviceProvider) writeData(pcID byte, maxBufLen uint32, data *Object, transferSyntax *transferSyntax) error {

	// create a writer to write the data to
	pdvWriter := newPDVWriter1(sp.machine, pcID, false, maxBufLen)

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the command to the buffer
	if err := encoder.writeObject(pdvWriter, data, transferSyntax); err != nil {
		return err
	}

	// flush to the underlying writer
	// passing true means we are done writing this object
	if err := pdvWriter.Flush(true); err != nil {
		return err
	}

	// all is well
	return nil
}

// copyDataFromReader copies the data from a reader to a stream of PDUs and PDVs
func (sp *serviceProvider) copyDataFromReader(
	pcID byte,
	maxBufLen uint32,
	reader io.Reader,
) error {

	// create a pdvWriter to write the data to
	// it knows how to create pdus and
	// since it implements a writer, we can use a copy method
	pdvWriter := newPDVWriter1(
		sp.machine, // write to the association connection
		pcID,       // pc id for each pdv
		false,      // false means we are writing data
		maxBufLen,  // max length of each pdu
	)

	// copy the data
	if _, err := io.Copy(pdvWriter, reader); err != nil {
		return err
	}

	// flush the data writer, true means we are done writing this object
	if err := pdvWriter.Flush(true); err != nil {
		return err
	}

	// return success
	return nil
}

// pdvWriter is used to write DICOM data using a series of PDUs and PDVs
type pdvWriter1 struct {
	machine   *machine      // the underlying machine
	pcID      byte          // the presentation context id
	isCommand bool          // is this a command or a data set?
	buf       *bytes.Buffer // a buffer to stage into
}

// newPDVWriter1 constructs a new PDataWriter
func newPDVWriter1(machine *machine, pcID byte, isCommand bool, maxLen uint32) *pdvWriter1 {
	return &pdvWriter1{
		machine,
		pcID,
		isCommand,
		bytes.NewBuffer(make([]byte, 0, maxLen-6)), // need to reserve 6 bytes for the PDV
	}
}

// Write implements the Writer inferface
func (pdvWriter *pdvWriter1) Write(buf []byte) (int, error) {

	// calculate how much is remaining
	remaining := pdvWriter.buf.Cap() - pdvWriter.buf.Len()

	// if there is more to write than remaining capacity,
	// write what we can, flush, and then write the remainder

	if len(buf) > remaining {

		// write what we can
		num, err := pdvWriter.Write(buf[:remaining])
		if err != nil {
			return num, err
		}

		// flush with more to come
		if err := pdvWriter.Flush(false); err != nil {
			return num, err
		}

		// write the remainder
		nextNum, err := pdvWriter.Write(buf[remaining:])

		// return the result, including the sum of the num bytes written
		return num + nextNum, err
	}

	// otherwise, just write bytes and return the result
	return pdvWriter.buf.Write(buf)
}

// Flush writes a PDV and PDU
func (pdvWriter *pdvWriter1) Flush(isLast bool) error {

	// create a PDV
	pdv := &pdv{}
	pdv.pcID = pdvWriter.pcID
	if pdvWriter.isCommand {
		pdv.mch = pdv.mch | 0x01
	}
	if isLast {
		pdv.mch = pdv.mch | 0x2
	}
	pdv.buf = pdvWriter.buf.Bytes()

	// we always write a pdv and pdu
	// while it is possible pack multiple pdvs into a single pdu
	// that requires some addition logic that i don't think benefits
	// us all that greatly

	// create a byte writer
	byteWriter := new(bytes.Buffer)

	// write the pdv header to the byte writer
	if err := writePDV(byteWriter, pdv); err != nil {
		return err
	}

	// create a PDU
	pdu := &pdu{}
	pdu.typ = pDataTFPDU
	pdu.buf = byteWriter.Bytes()

	// // write the PDU
	// if err := writePDU(pdvWriter.writer, pdu); err != nil {
	// 	return err
	// }
	d := &deed{e: evt9, p: pdu}
	pdvWriter.machine.deedChan <- d

	// reset the buffer
	pdvWriter.buf.Reset()

	// all is well
	return nil
}

type messageReader struct {
	sp *serviceProvider
}

func (mr *messageReader) run() {

	for {

		message, err := mr.readMessage(mr.sp.assoc, false)
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

func (mr *messageReader) readMessage(assoc *Assoc, shouldReadData bool) (*Message, error) {

	// create a reader for the command
	// start by reading from the pdu that was provided
	log.Printf("message reader waiting for pdv reader")
	commandReader, err := newPDVReader1(mr.sp, true)
	if err != nil {
		log.Printf("while creating pdv reader, error was %v", err)
		return nil, err
	}

	// get the presentation context id from the reader
	pcID := commandReader.pdv.pcID

	// read the command from the pdu
	command, err := readCommand(commandReader)
	if err != nil {
		log.Printf("while reading command, error was %v", err)
		return nil, err
	}

	log.Printf("pc id is %v", pcID)
	log.Printf("command is %v", command)

	// get the command data set
	commandDataSet, err := command.asShort(CommandDataSetTypeTag, 0)
	if err != nil {
		return nil, err
	}

	// create a message
	message := &Message{
		pcID:    pcID,
		command: command,
	}

	if isDataSetPresent(commandDataSet) {

		// create a reader for the data
		dataReader, err := newPDVReader1(mr.sp, false)
		if err != nil {
			return nil, err
		}

		if shouldReadData {

			// read the data
			data, err := readData(dataReader, assoc, pcID)
			if err != nil {
				return nil, err
			}

			// add the data to the message
			message.data = data

		} else {

			// otherwise, add the data reader to the message
			message.dataReader = dataReader
		}
	}

	// return the message
	return message, nil
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
