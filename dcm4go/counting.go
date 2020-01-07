package dcm4go

import "io"

// CounterReader extends the Reader interface to include counting of bytes read
type CounterReader interface {
	io.Reader
	BytesRead() int
}

// CountingReader wraps a reader and counts the bytes read
type CountingReader struct {
	bytesRead int
	reader    io.Reader
}

// newCountingReader returns a new counting reader
func newCountingReader(reader io.Reader) CounterReader {
	return &CountingReader{0, reader}
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

// LimitedCountingReader wraps a limited reader and provides access to bytes read
type LimitedCountingReader struct {
	countingReader CounterReader
	limitedReader  io.Reader
}

// newLimitedReader returns a new limited counting reader
func newLimitedCountingReader(countingReader CounterReader, length int64) CounterReader {
	return &LimitedCountingReader{countingReader, io.LimitReader(countingReader, length)}
}

// Read reads bytes from the underling limited reader
func (limitedCountingReader *LimitedCountingReader) Read(buf []byte) (int, error) {
	return limitedCountingReader.limitedReader.Read(buf)
}

// BytesRead returns the number of bytes read by the underlying counting reader
func (limitedCountingReader *LimitedCountingReader) BytesRead() int {
	return limitedCountingReader.countingReader.BytesRead()
}
