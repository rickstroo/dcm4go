package dcm4go

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"testing"
)

func initReadTest(buf []byte) io.Reader {
	return bytes.NewReader(buf)
}

func testByteEquals(a, b byte) error {
	if a != b {
		return fmt.Errorf("expected 0x%02X, was 0x%02X", a, b)
	}
	return nil
}

func TestReadBytes(t *testing.T) {
	reader := initReadTest([]byte{0x12, 0x34, 0x56, 0x78})
	var buf [4]byte
	if err := readBytes(reader, buf[:]); err != nil {
		t.Error(err)
	}
	if err := testByteEquals(buf[0], 0x12); err != nil {
		t.Error(err)
	}
	if err := testByteEquals(buf[1], 0x34); err != nil {
		t.Error(err)
	}
	if err := testByteEquals(buf[2], 0x56); err != nil {
		t.Error(err)
	}
	if err := testByteEquals(buf[3], 0x78); err != nil {
		t.Error(err)
	}
}

func testShortEquals(a, b uint16) error {
	if a != b {
		return fmt.Errorf("expected 0x%04X, was 0x%04X", a, b)
	}
	return nil
}

func TestReadShortLittleEndian(t *testing.T) {
	reader := initReadTest([]byte{0x12, 0x34})
	short, err := readShort(reader, binary.LittleEndian)
	if err != nil {
		t.Error(err)
	}
	if err := testShortEquals(short, 0x3412); err != nil {
		t.Error(err)
	}
}

func TestReadShortBigEndian(t *testing.T) {
	reader := initReadTest([]byte{0x12, 0x34})
	short, err := readShort(reader, binary.BigEndian)
	if err != nil {
		t.Error(err)
	}
	if err := testShortEquals(short, 0x1234); err != nil {
		t.Error(err)
	}
}

func TestReadShortUnexpectedEOF(t *testing.T) {
	reader := initReadTest([]byte{0x12})
	_, err := readShort(reader, binary.BigEndian)
	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected io.ErrUnexpectedEOF, was %v", err)
	}
}
