// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"io"
)

// Message presents the requests and responses that are passed between AEs
type Message struct {
	pcID        byte
	command     *Object
	data        *Object
	pDataReader *PDataReader
}

var messageID uint16

func nextMessageID() uint16 {
	messageID := messageID + 1
	return messageID
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

// PCID returns the presentation context id of the message
func (message *Message) PCID() byte {
	return message.pcID
}

// Command returns the command portion of the message
func (message *Message) Command() *Object {
	return message.command
}

// Data returns the data portion of the message
func (message *Message) Data() *Object {
	return message.data
}

// PDataReader returns the data reader
func (message *Message) PDataReader() *PDataReader {
	return message.pDataReader
}

func isDataSetPresent(commandDataSetType uint16) bool {
	return commandDataSetType != 0x0101
}

func readMessage(
	reader io.Reader,
	assoc *Assoc,
	pdu *PDU,
	shouldReadData bool,
) (
	*Message,
	error,
) {

	// create a reader for the command
	commandReader, err := newPDataReader(reader, pdu, true)
	if err != nil {
		return nil, err
	}

	// get the presentation context id from the reader
	pcID := commandReader.pdv.pcID

	// read the command from the pdu
	command, err := readCommand(commandReader, assoc)
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
		pDataReader, err := newPDataReader(reader, pdu, false)
		if err != nil {
			return nil, err
		}

		// should we read the data, or pass the data reader back to the caller?
		if shouldReadData {

			// read the data
			data, err := readData(pDataReader, assoc, pcID)
			if err != nil {
				return nil, err
			}

			// add the data to the message
			message.data = data

		} else {

			// otherwise, add the data reader to the message
			message.pDataReader = pDataReader
		}
	}

	// return the message
	return message, nil
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
	transferSyntax, err := assoc.findAcceptedTransferSyntaxByPCID(pcID)
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

	if err := writeCommand(writer, assoc, message.pcID, message.Command()); err != nil {
		return err
	}

	if message.Data() != nil {

		transferSyntax, err := assoc.findAcceptedTransferSyntaxByPCID(message.pcID)
		if err != nil {
			return err
		}

		if err := writeData(writer, assoc, message.pcID, message.Command(), transferSyntax); err != nil {
			return err
		}
	}

	if message.PDataReader() != nil {
		if err := copyDataFromReader(writer, message.pcID, assoc.assocRQPDU.userInfo.maxLenReceived, message.PDataReader()); err != nil {
			return err
		}
	}
	return nil
}

// writeCommand writes the command portion of the message
func writeCommand(writer io.Writer, assoc *Assoc, pcID byte, command *Object) error {

	// create a writer to write the data to
	pDataWriter := newPDataWriter(writer, pcID, true, assoc.assocRQPDU.userInfo.maxLenReceived)

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the command with group length
	if err := encoder.writeObjectWithGroupLength(pDataWriter, 0x0000, command, ImplicitVRLittleEndianTS); err != nil {
		return err
	}

	// flush to the underlying writer
	// passing true means we are done writing this object
	if err := pDataWriter.Flush(true); err != nil {
		return err
	}

	// all is well
	return nil
}

// // writeData writes the data portion of the message
func writeData(writer io.Writer, assoc *Assoc, pcID byte, data *Object, transferSyntax *TransferSyntax) error {

	// create a writer to write the data to
	pDataWriter := newPDataWriter(writer, pcID, false, assoc.assocRQPDU.userInfo.maxLenReceived)

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the command to the buffer
	if err := encoder.writeObject(pDataWriter, data, transferSyntax); err != nil {
		return err
	}

	// flush to the underlying writer
	// passing true means we are done writing this object
	if err := pDataWriter.Flush(true); err != nil {
		return err
	}

	// all is well
	return nil
}

// copyDataFromReader copies the data from a reader to a stream of PDUs and PDVs
func copyDataFromReader(
	writer io.Writer,
	pcID byte,
	maxBufLen uint32,
	reader io.Reader,
) error {

	// create a pdatawriter to write the data to
	// it knows how to create pdus and
	// since it implements a writer, we can use a copy method
	pDataWriter := newPDataWriter(
		writer,    // write to the association connection
		pcID,      // pc id for each pdv
		false,     // false means we are writing data
		maxBufLen, // max length of each pdu
	)

	// copy the data
	if _, err := io.Copy(pDataWriter, reader); err != nil {
		return err
	}

	// flush the data writer, true means we are done writing this object
	if err := pDataWriter.Flush(true); err != nil {
		return err
	}

	// return success
	return nil
}
