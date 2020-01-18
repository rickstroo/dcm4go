package dcm4go

import (
	"encoding/binary"
	"io"
)

// readReleaseRQPDU reads an ReleaseRQPDU from a reader
func readReleaseRQPDU(reader io.Reader) error {

	// read and ignore the long
	if _, err := readLong(reader, binary.BigEndian); err != nil {
		return err
	}

	return nil
}
