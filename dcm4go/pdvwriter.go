// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"bytes"
	"io"
)

// pdvWriter is used to write DICOM data using a series of PDUs and PDVs
type pdvWriter struct {
	writer    io.Writer     // the underlying writer
	pcID      byte          // the presentation context id
	isCommand bool          // is this a command or a data set?
	buf       *bytes.Buffer // a buffer to stage into
}

// newPDVWriter constructs a new PDataWriter
func newPDVWriter(writer io.Writer, pcID byte, isCommand bool, maxLen uint32) *pdvWriter {
	return &pdvWriter{
		writer,
		pcID,
		isCommand,
		bytes.NewBuffer(make([]byte, 0, maxLen-6)), // need to reserve 6 bytes for the PDV
	}
}

// Write implements the Writer inferface
func (pdvWriter *pdvWriter) Write(buf []byte) (int, error) {

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
func (pdvWriter *pdvWriter) Flush(isLast bool) error {

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

	// create a data transfer PDU with a single PDV
	// if we wanted to implement packing of multiple PDVs into a single PDU,
	// we would probably write the PDV to the PDU writer and let it take
	// care of the packing, flushing when the PDU length is long enough
	dataTFPDU := newDataTFPDU()
	dataTFPDU.addPDV(pdv)

	// write the pdu
	if err := dataTFPDU.writeTo(pdvWriter.writer); err != nil {
		return err
	}

	// // we always write a pdv and pdu
	// // while it is possible pack multiple pdvs into a single pdu
	// // that requires some addition logic that i don't think benefits
	// // us all that greatly
	//
	// // create a byte writer
	// byteWriter := new(bytes.Buffer)
	//
	// // write the pdv header to the byte writer
	// if err := writePDV(byteWriter, pdv); err != nil {
	// 	return err
	// }
	//
	// // write the bytes of the pdv to the byte writer
	// // we really want to make this more efficient
	// // we don't want to copy the bytes and copy them again
	// // oh well, we'll do it this way for now
	// if err := writeBytes(byteWriter, pdvWriter.buf.Bytes()); err != nil {
	// 	return err
	// }
	//
	// // create a data transfer pdu
	// dataTFPDU := &dataTFPDU{
	// 	buf: byteWriter.Bytes(),
	// }
	//
	// // write the data transfer pdu
	// if err := dataTFPDU.writeTo(pdvWriter.writer); err != nil {
	// 	return err
	// }

	// reset the buffer
	pdvWriter.buf.Reset()

	// all is well
	return nil
}
