// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"io"
)

// A Message represents the requests and responses that are
// passed between AEs
type Message struct {
	pcID       byte
	command    *Object
	data       *Object
	dataReader io.Reader
}

var messageID uint16

func nextMessageID() uint16 {
	messageID := messageID + 1
	return messageID
}

// PCID returns the presentation context of the message
func (message *Message) PCID() byte {
	return message.pcID
}

// Command returns the command of the message
func (message *Message) Command() *Object {
	return message.command
}

// Data returns the data of the message
func (message *Message) Data() *Object {
	return message.data
}

// DataReader returns the data reader
func (message *Message) DataReader() io.Reader {
	return message.dataReader
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

func isDataSetPresent(commandDataSetType uint16) bool {
	return commandDataSetType != NoDataSetCode
}

func readMessage(assoc *Assoc, shouldReadData bool, pDataReader pDataReader) (*Message, error) {

	// create a reader for the command
	// start by reading from the pdu that was provided
	commandReader, err := newPDVReader(pDataReader, true)
	if err != nil {
		return nil, err
	}

	// get the presentation context id from the reader
	pcID := commandReader.pdv.pcID

	// read the command from the pdu
	command, err := readCommand(commandReader)
	if err != nil {
		return nil, err
	}

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
		dataReader, err := newPDVReader(pDataReader, false)
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

func readCommand(reader io.Reader) (*Object, error) {

	// read the object
	object, err := readObject(reader, ImplicitVRLittleEndianTS)
	if err != nil {
		return nil, err
	}

	// return the object
	return object, nil
}

func readData(reader io.Reader, assoc *Assoc, pcID byte) (*Object, error) {

	// find the negotiated transfer syntax for the data
	transferSyntax, err := assoc.findAcceptedTransferSyntaxByPCID(pcID)
	if err != nil {
		return nil, err
	}
	fmt.Printf("transfer syntax for request data is %v\n", transferSyntax)

	// read the object
	object, err := readObject(reader, transferSyntax)
	if err != nil {
		return nil, err
	}

	// return the object
	return object, nil
}

func readObject(reader io.Reader, transferSyntax *transferSyntax) (*Object, error) {

	// create a counting reader
	countingReader := newCountingReader(reader)

	// create a decoder to read the data
	decoder := newDecoder(0)

	// read the object, assuming explicit VR and big endian for now
	object, err := decoder.readObject(countingReader, transferSyntax)
	if err != nil {
		return nil, err
	}

	// return the object
	return object, nil
}

// WriteMessage writes the message
func writeMessage(assoc *Assoc, message *Message, pDataWriter pDataWriter) error {

	if err := writeCommand(pDataWriter, message.pcID, assoc.assocRQPDU.userInfo.maxLenReceived, message.command); err != nil {
		return err
	}

	if message.data != nil {

		transferSyntax, err := assoc.findAcceptedTransferSyntaxByPCID(message.pcID)
		if err != nil {
			return err
		}

		if err := writeData(pDataWriter, message.pcID, assoc.assocRQPDU.userInfo.maxLenReceived, message.data, transferSyntax); err != nil {
			return err
		}

	} else if message.dataReader != nil {

		if err := copyDataFromReader(pDataWriter, message.pcID, assoc.assocRQPDU.userInfo.maxLenReceived, message.dataReader); err != nil {
			return err
		}
	}

	// return success
	return nil
}

// writeCommand writes the command portion of the message
func writeCommand(pDataWriter pDataWriter, pcID byte, maxBufLen uint32, command *Object) error {

	// create a writer to write the data to
	pdvWriter := newPDVWriter(pDataWriter, pcID, true, maxBufLen)

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
func writeData(pDataWriter pDataWriter, pcID byte, maxBufLen uint32, data *Object, transferSyntax *transferSyntax) error {

	// create a writer to write the data to
	pdvWriter := newPDVWriter(pDataWriter, pcID, false, maxBufLen)

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
func copyDataFromReader(
	pDataWriter pDataWriter,
	pcID byte,
	maxBufLen uint32,
	reader io.Reader,
) error {

	// create a pdvWriter to write the data to
	// it knows how to create pdus and
	// since it implements a writer, we can use a copy method
	pdvWriter := newPDVWriter(
		pDataWriter, // write to the association connection
		pcID,        // pc id for each pdv
		false,       // false means we are writing data
		maxBufLen,   // max length of each pdu
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
