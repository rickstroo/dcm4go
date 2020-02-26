package dcm4go

import "io"

// countingReader wraps a reader and counts the bytes read
type countingReader struct {
	reader    io.Reader
	bytesRead int
}

// newCountingReader returns a new counting reader
func newCountingReader(reader io.Reader) *countingReader {
	return &countingReader{reader: reader, bytesRead: 0}
}

// Read calls underlying reader and counts the bytes
func (countingReader *countingReader) Read(buf []byte) (int, error) {
	num, err := countingReader.reader.Read(buf)
	countingReader.bytesRead += num
	return num, err
}

// limitCountingReader returns a CountingReader that reads from the underlying
// countingReader and stops with EOF after read length bytes.  The returned
// CountingReader has a starting bytesRead equal to the current bytesRead of
// the undelying countingReader.  Since the returned CountingReader reads from the underlying
// countingReader, the underlying countingReader's bytesRead will be updated as well.
func limitCountingReader(reader *countingReader, length int64) *countingReader {
	return &countingReader{reader: io.LimitReader(reader, length), bytesRead: reader.bytesRead}
}
