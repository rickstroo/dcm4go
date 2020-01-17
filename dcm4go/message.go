package dcm4go

import (
	"fmt"
	"io"
)

const (
	cEchoRQ  = 0x0030
	cEchoRSP = 0x8030
)

// Message presents the requests and responses that are passed between AEs
type Message struct {
	command *Object
	data    *Object
}

func (message *Message) Command() *Object {
	return message.command
}

func (message *Message) Data() *Object {
	return message.data
}

func noDataSet(commandDataSetType uint16) bool {
	return commandDataSetType == 0x0101
}

func readMessage(reader io.Reader, assoc *Assoc) (*Message, error) {

	command, err := readCommand(reader, assoc)
	if err != nil {
		return nil, err
	}

	commandDataSet, err := command.asShort(CommandDataSetTypeTag, 0)
	if err != nil {
		return nil, err
	}

	if noDataSet(commandDataSet) {
		return &Message{command, nil}, nil
	}

	data, err := readData(reader, assoc)
	if err != nil {
		return nil, err
	}

	return &Message{command, data}, nil
}

func findAcceptedTransferSyntax(assoc *Assoc, pcid byte) (*TransferSyntax, error) {
	for _, presContext := range assoc.assocACPDU.presContexts {
		if presContext.id == pcid {
			return findTransferSyntax(presContext.transferSyntax)
		}
	}
	return nil, fmt.Errorf("no supported transfer syntax found for presentation context id %d", pcid)
}

func readCommand(reader io.Reader, assoc *Assoc) (*Object, error) {

	// read the initial PDV
	pdv, err := readPDV(reader)
	if err != nil {
		return nil, err
	}

	// check that this is a command pdv
	if !pdv.isCommand() {
		return nil, fmt.Errorf("not a command pdv")
	}

	// check that this is the last pdv
	// later we will implement support for multiple pdu pdvs
	if !pdv.isLast() {
		return nil, fmt.Errorf("not the last fragment")
	}

	// create a reader for the rest of the pdv, less two bytes for the pcid and msh
	limitReader := io.LimitReader(reader, int64(pdv.pdvLength)-2)

	// create a counting reader
	countingReader := newCountingReader(limitReader)

	// create a decoder to read the data
	decoder := newDecoder(0)

	// set the transfer syntax for now, will find it later
	transferSyntax, err := findAcceptedTransferSyntax(assoc, pdv.pcID)
	if err != nil {
		return nil, err
	}

	// read the data, assuming explicit VR and big endian for now
	command, err := decoder.readObject(countingReader, transferSyntax.explicitVR, transferSyntax.byteOrder)
	if err != nil {
		return nil, err
	}

	fmt.Printf("command is %q\n", command)

	// return the command
	return command, nil
}

func readData(reader io.Reader, assoc *Assoc) (*Object, error) {

	// read the initial PDV
	pdv, err := readPDV(reader)
	if err != nil {
		return nil, err
	}

	// check that this is a data pdv
	if pdv.isCommand() {
		return nil, fmt.Errorf("not a command pdv")
	}

	// check that this is the last pdv
	// later we will implement support for multiple pdu pdvs
	if !pdv.isLast() {
		return nil, fmt.Errorf("not the last fragment")
	}

	// create a reader for the rest of the pdv, less two bytes for the pcid and msh
	limitReader := io.LimitReader(reader, int64(pdv.pdvLength)-2)

	// create a counting reader
	countingReader := newCountingReader(limitReader)

	// create a decoder to read the data
	decoder := newDecoder(0)

	// set the transfer syntax for now, will find it later
	transferSyntax, err := findTransferSyntax("1.2.840.10008.1.2.1")
	if err != nil {
		return nil, err
	}

	// read the data, assuming explicit VR and big endian for now
	data, err := decoder.readObject(countingReader, transferSyntax.explicitVR, transferSyntax.byteOrder)
	if err != nil {
		return nil, err
	}

	// return the data
	return data, nil
}
