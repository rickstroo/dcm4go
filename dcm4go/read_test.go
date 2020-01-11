package dcm4go

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
)

func initReadTest(buf []byte) io.Reader {
	return bytes.NewReader(buf)
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

func TestReadByte(t *testing.T) {
	reader := initReadTest([]byte{0x12})
	b, err := readByte(reader)
	if err != nil {
		t.Error(err)
	}
	if err := testByteEquals(b, 0x12); err != nil {
		t.Error(err)
	}
}

func TestReadByteUnexpectedEOF(t *testing.T) {
	reader := initReadTest([]byte{})
	_, err := readByte(reader)
	if err := testErrEquals(err, io.EOF); err != nil {
		t.Error(err)
	}
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
	if err := testErrEquals(err, io.ErrUnexpectedEOF); err != nil {
		t.Error(err)
	}
}
