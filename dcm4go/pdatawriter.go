package dcm4go

import (
	"bytes"
	"fmt"
	"io"
)

// PDataWriter is used to write DICOM data using a series of PDUs and PDVs
type PDataWriter struct {
	writer    io.Writer     // the underlying writer
	pcID      byte          // the presentation context id
	isCommand bool          // is this a command or a data set?
	maxLen    uint32        // the maximum length of a PDU that we will write
	buf       *bytes.Buffer // a buffer to stage into
}

// newPDataWriter constructs a new PDataWriter
func newPDataWriter(writer io.Writer, pcID byte, isCommand bool, maxLen uint32) *PDataWriter {
	return &PDataWriter{
		writer,
		pcID,
		isCommand,
		maxLen,
		bytes.NewBuffer(make([]byte, 0, maxLen-6)), // need to reserve 6 bytes for the PDV
	}
}

// Write implements the Writer inferface
func (pDataWriter *PDataWriter) Write(buf []byte) (int, error) {
	fmt.Printf("len(buf) is %d, pDataWriter.buf.Len() is %d, pDataWriter.buf.Cap() is %d\n", len(buf), pDataWriter.buf.Len(), pDataWriter.buf.Cap())

	// calculate how much is remaining
	remaining := pDataWriter.buf.Cap() - pDataWriter.buf.Len()

	// if there is more to write than remaining capacity,
	// write what we can, flush, and then write the remainder

	if len(buf) > remaining {

		fmt.Printf("not enough capacity, will write %d bytes and then flush\n", remaining)

		// write what we can
		num, err := pDataWriter.Write(buf[:remaining])
		if err != nil {
			return num, err
		}

		// flush with more to come
		if err := pDataWriter.Flush(false); err != nil {
			return num, err
		}

		fmt.Printf("and now, we will write the remainder\n")

		// write the remaidner
		return pDataWriter.Write(buf[remaining:])
	}

	// otherwise, just write bytes
	return pDataWriter.buf.Write(buf)
}

// Flush writes a PDV and PDU
func (pDataWriter *PDataWriter) Flush(isLast bool) error {

	// create a PDV
	pdv := &PDV{}
	pdv.pcID = pDataWriter.pcID
	if pDataWriter.isCommand {
		pdv.mch = pdv.mch | 0x01
	}
	if isLast {
		pdv.mch = pdv.mch | 0x2
	}
	pdv.pdvLength = uint32(pDataWriter.buf.Len() + 2) // need to add two bytes for the pcID and mch
	fmt.Printf("pdv to write is is %v\n", pdv)

	// create a pdu
	pdu := &PDU{}
	pdu.pduType = pDataTFPDU
	pdu.pduLength = uint32(pdv.pdvLength + 4) // need to add four bytes for the PDV length
	fmt.Printf("pdu to write is %v\n", pdu)

	// we always write a pdv and pdu
	// while it is possible pack multiple pdvs into a single pdu
	// that requires some addition logic that i don't think benefits
	// us all that greatly

	// write the pdu header
	if err := writePDU(pDataWriter.writer, pdu); err != nil {
		return err
	}

	// write the pdv header
	if err := writePDV(pDataWriter.writer, pdv); err != nil {
		return err
	}

	// write the bytes of the buffer
	if err := writeBytes(pDataWriter.writer, pDataWriter.buf.Bytes()); err != nil {
		return err
	}

	// reset the buffer
	pDataWriter.buf.Reset()

	// all is well
	return nil
}
