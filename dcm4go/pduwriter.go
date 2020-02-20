package dcm4go

import "io"

type pduWriter struct {
	writer io.Writer // underlying writer
}

func newPDUWriter(writer io.Writer) *pduWriter {
	return &pduWriter{writer: writer}
}

func (pduWriter *pduWriter) Write(buf []byte) (int, error) {
	return pduWriter.writer.Write(buf)
}
