package dcm4go

import (
	"fmt"
	"io"
)

const (
	// CEchoRQ is command field value for C-Echo request
	CEchoRQ = 0x0030
	// CEchoRSP is command field value for C-Recho response
	CEchoRSP = 0x8030
	// CStoreRQ is command field value for C-Store request
	CStoreRQ = 0x0001
	// CStoreRSP is command field value for C-Store response
	CStoreRSP = 0x8001
)

// Message presents the requests and responses that are passed between AEs
type Message struct {
	pcID    byte
	command *Object
	data    *Object
}

func (message *Message) String() string {
	s := fmt.Sprintf("{pcID:%v", message.pcID)
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
	s += "}"
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

// PCID returns the presentation context id of the message
func (message *Message) PCID() byte {
	return message.pcID
}

func isDataSetPresent(commandDataSetType uint16) bool {
	return commandDataSetType != 0x0101
}

func readMessage(reader io.Reader, assoc *Assoc, pdu *PDU) (*Message, error) {

	// create a reader for the command
	commandReader, err := newPDataReader(reader, pdu, true)
	if err != nil {
		return nil, err
	}

	// get the presentation context id from the reader
	pcID := commandReader.pdv.pcID

	// read the command
	command, err := readCommand(commandReader, assoc)
	if err != nil {
		return nil, err
	}

	// get the command data set
	commandDataSet, err := command.asShort(CommandDataSetTypeTag, 0)
	if err != nil {
		return nil, err
	}

	if isDataSetPresent(commandDataSet) {

		// create a reader for the data
		dataReader, err := newPDataReader(reader, pdu, false)
		if err != nil {
			return nil, err
		}

		// find the affected sop class uid
		affectedSOPClassUID, err := command.asString(AffectedSOPClassUIDTag, 0)
		if err != nil {
			return nil, err
		}

		// if there is a handler for this command, call it
		commandHandler, ok := assoc.ae.commandHandlers[affectedSOPClassUID]
		if ok && (commandHandler != nil) {

			// call the handler, which will return the data, which may be potentially nil
			// if the handler decides to consume the data itself
			data, err := commandHandler.HandleCommand(assoc, pcID, command, dataReader)
			if err != nil {
				return nil, err
			}

			// return the request with command and data
			return &Message{pcID, command, data}, nil
		}

		// if no handler, read the data ourselves
		data, err := readData(dataReader, assoc, pcID)
		if err != nil {
			return nil, err
		}

		// return the request with command and data
		return &Message{pcID, command, data}, nil
	}

	// return the request with command and no data
	return &Message{pcID, command, nil}, nil
}

func findAcceptedTransferSyntax(assoc *Assoc, pcid byte) (*TransferSyntax, error) {
	for _, presContext := range assoc.assocACPDU.presContexts {
		if presContext.id == pcid {
			return findTransferSyntax(presContext.transferSyntaxes[0])
		}
	}
	return nil, fmt.Errorf("no supported transfer syntax found for presentation context id %d", pcid)
}

func readCommand(reader io.Reader, assoc *Assoc) (*Object, error) {

	// create a counting reader
	countingReader := newCountingReader(reader)

	// create a decoder to read the data
	decoder := newDecoder(0)

	// read the data, assuming explicit VR and big endian for now
	command, err := decoder.readObject(countingReader, ImplicitVRLittleEndianTS)
	if err != nil {
		return nil, err
	}

	// return the command and transfer syntax used to read the command
	return command, nil
}

func readData(reader io.Reader, assoc *Assoc, pcID byte) (*Object, error) {

	// create a counting reader
	countingReader := newCountingReader(reader)

	// create a decoder to read the data
	decoder := newDecoder(1024)

	// find the negotiated transfer syntax for the data
	transferSyntax, err := findAcceptedTransferSyntax(assoc, pcID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("transfer syntax for request data is %v\n", transferSyntax)

	// read the data, assuming the negotiated transfer syntax
	data, err := decoder.readObject(countingReader, transferSyntax)
	if err != nil {
		return nil, err
	}

	// return the data
	return data, nil
}

// WriteMessage writes the message
func writeMessage(writer io.Writer, assoc *Assoc, message *Message) error {

	// create a writer to write the data to
	pDataWriter := newPDataWriter(writer, message.pcID, true, assoc.assocRQPDU.userInfo.maxLenReceived)

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the command to the buffer
	if err := encoder.writeObjectWithGroupLength(pDataWriter, 0x0000, message.Command(), ImplicitVRLittleEndianTS); err != nil {
		return err
	}

	// flush to the underlying writer
	// passing true means we are done writing this object
	if err := pDataWriter.Flush(true); err != nil {
		return err
	}

	return nil
}
