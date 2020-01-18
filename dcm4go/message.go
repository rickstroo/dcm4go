package dcm4go

import (
	"bytes"
	"fmt"
	"io"
)

const (
	// CEchoRQ is command field value for C-Echo request
	CEchoRQ = 0x0030
	// CEchoRSP is command field value for C-Recho response
	CEchoRSP = 0x8030
)

// Message presents the requests and responses that are passed between AEs
type Message struct {
	pcID    byte
	command *Object
	data    *Object
}

func (message *Message) String() string {
	s := fmt.Sprintf("pcID:%d", message.pcID)
	if message.command != nil {
		s += fmt.Sprintf(",command:%s", message.command)
	} else {
		s += fmt.Sprintf(",command:nil")
	}
	if message.data != nil {
		s += fmt.Sprintf(",data:%s", message.data)
	} else {
		s += fmt.Sprintf(",data:nil")
	}
	return s
}

// Command returns the command portion of the message
func (message *Message) Command() *Object {
	return message.command
}

// Data returns the data portion of the message
func (message *Message) Data() *Object {
	return message.data
}

func isDataSetPresent(commandDataSetType uint16) bool {
	return commandDataSetType != 0x0101
}

func readMessage(reader io.Reader, assoc *Assoc) (*Message, error) {

	pcid, command, err := readCommand(reader, assoc)
	if err != nil {
		return nil, err
	}

	commandDataSet, err := command.asShort(CommandDataSetTypeTag, 0)
	if err != nil {
		return nil, err
	}

	if isDataSetPresent(commandDataSet) {
		return nil, fmt.Errorf("readMessage: read of data set not implemented")
	}

	return &Message{pcid, command, nil}, nil
}

func findAcceptedTransferSyntax(assoc *Assoc, pcid byte) (*TransferSyntax, error) {
	for _, presContext := range assoc.assocACPDU.presContexts {
		if presContext.id == pcid {
			return findTransferSyntax(presContext.transferSyntax)
		}
	}
	return nil, fmt.Errorf("no supported transfer syntax found for presentation context id %d", pcid)
}

func readCommand(reader io.Reader, assoc *Assoc) (byte, *Object, error) {

	// read the initial PDV
	pdv, err := readPDV(reader)
	if err != nil {
		return 0, nil, err
	}

	// check that this is a command pdv
	if !pdv.isCommand() {
		return 0, nil, fmt.Errorf("not a command pdv")
	}

	// check that this is the last pdv
	// later we will implement support for multiple pdu pdvs
	if !pdv.isLast() {
		return 0, nil, fmt.Errorf("not the last fragment")
	}

	// create a reader for the rest of the pdv, less two bytes for the pcid and msh
	limitReader := io.LimitReader(reader, int64(pdv.pdvLength)-2)

	// create a counting reader
	countingReader := newCountingReader(limitReader)

	// create a decoder to read the data
	decoder := newDecoder(0)

	// find the transfer syntax
	transferSyntax, err := findAcceptedTransferSyntax(assoc, pdv.pcID)
	if err != nil {
		return 0, nil, err
	}

	// read the data, assuming explicit VR and big endian for now
	command, err := decoder.readObject(countingReader, transferSyntax.explicitVR, transferSyntax.byteOrder)
	if err != nil {
		return 0, nil, err
	}

	// return the command and transfer syntax used to read the command
	return pdv.pcID, command, nil
}

// NewCEchoResponse constructs a C-Echo response message based on the C-Echo request message
func NewCEchoResponse(request *Message) (*Message, error) {
	pcID := request.pcID
	response := &Message{pcID, newObject(), nil}
	response.Command().addUID(AffectedSOPClassUIDTag, VerificationUID)
	response.Command().addShort(CommandFieldTag, "US", CEchoRSP)
	messageID, err := request.Command().AsShort(MessageIDTag, 0)
	if err != nil {
		return nil, err
	}
	response.Command().addShort(MessageIDBeingRespondedToTag, "US", messageID)
	response.Command().addShort(CommandDataSetTypeTag, "US", 0x0101)
	response.Command().addShort(StatusTag, "US", 0x00)
	return response, nil
}

// WriteMessage writes the message
func writeMessage(writer io.Writer, assoc *Assoc, message *Message) error {

	// find the transfer syntax
	transferSyntax, err := findAcceptedTransferSyntax(assoc, message.pcID)
	if err != nil {
		return err
	}

	// create a buffer to write the command object to
	buf := new(bytes.Buffer)

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the object to the buffer
	if err := encoder.writeObject(buf, message.Command(), transferSyntax.explicitVR, transferSyntax.byteOrder); err != nil {
		return err
	}

	// create a PDV
	pdv := &PDV{}
	pdv.pcID = message.pcID               // same pc id as in the request
	pdv.mch = 0x3                         // last and command
	pdv.pdvLength = uint32(buf.Len() + 2) // need to add two bytes for the pcID and mch

	// create a pdu
	pdu := &PDU{}
	pdu.pduType = pDataTFPDU
	pdu.pduLength = uint32(buf.Len() + 6) // need to add two bytes for the pcID and mch and 4 bytes for the PDV header

	// write the pdu header
	if err := writePDU(writer, pdu); err != nil {
		return err
	}

	// write the pdv header
	if err := writePDV(writer, pdv); err != nil {
		return err
	}

	// write the bytes containing the object
	if err := writeBytes(writer, buf.Bytes()); err != nil {
		return err
	}

	return nil
}
