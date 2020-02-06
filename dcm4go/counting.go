package dcm4go

import "io"

// CountingReader wraps a reader and counts the bytes read
type CountingReader struct {
	reader    io.Reader
	bytesRead int
}

// newCountingReader returns a new counting reader
func newCountingReader(reader io.Reader) *CountingReader {
	return &CountingReader{reader: reader, bytesRead: 0}
}

// Read calls underlying reader and counts the bytes
func (countingReader *CountingReader) Read(buf []byte) (int, error) {
	num, err := countingReader.reader.Read(buf)
	countingReader.bytesRead += num
	return num, err
}

// BytesRead returns the number of bytes read
func (countingReader *CountingReader) BytesRead() int {
	return countingReader.bytesRead
}

// limitCountingReader returns a CountingReader that reads from the underlying
// countingReader and stops with EOF after read length bytes.  The returned
// CountingReader has a starting bytesRead equal to the current bytesRead of
// the undelying countingReader.  Since the returned CountingReader reads from the underlying
// countingReader, the underlying countingReader's bytesRead will be updated as well.
func limitCountingReader(countingReader *CountingReader, length int64) *CountingReader {
	return &CountingReader{reader: io.LimitReader(countingReader, length), bytesRead: countingReader.bytesRead}
}
