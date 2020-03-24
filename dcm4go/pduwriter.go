package dcm4go

import (
	"io"
)

type pduWriter struct {
	writer io.Writer // underlying writer
}

func newPDUWriter(writer io.Writer) *pduWriter {
	pduWriter := &pduWriter{
		writer: writer,
	}
	return pduWriter
}

func (pduWriter *pduWriter) Write(buf []byte) (int, error) {
	return pduWriter.writer.Write(buf)
}

// this is a bit of an experiment
// i'm going to use an interface to define a P-DATA-TF pdu writer
// then i'm going to implement a couple of versions of the writer
// one that writes to the state machine and one that writes to a connection
// i'm hoping this allows me to use a single pdv writer for both implementations

type pDataWriter interface {
	writePDU(pdu *pdu) error
}

type pDataNetworkWriter struct {
	writer io.Writer
}

func (writer *pDataNetworkWriter) writePDU(pdu *pdu) error {
	if err := writePDU(writer.writer, pdu); err != nil {
		return err
	}
	return nil
}

type pDataMachineWriter struct {
	machine *machine
}

func (writer *pDataMachineWriter) writePDU(pdu *pdu) error {
	writer.machine.deedChan <- &deed{e: evt9, p: pdu}
	return nil
}
