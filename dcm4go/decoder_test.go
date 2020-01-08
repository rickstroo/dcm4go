package dcm4go

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"testing"
)

func initTest(buf []byte) (CounterReader, *Decoder) {
	return newCountingReader(bytes.NewReader(buf)), newDecoder(1024)
}

func TestNewDecoder(t *testing.T) {
	_, decoder := initTest([]byte{})
	if decoder.bulkDataThreshold != 1024 {
		t.Errorf("expected %d, was %d", 1024, decoder.bulkDataThreshold)
	}
}

func testByteEquals(a, b byte) error {
	if a != b {
		return fmt.Errorf("expected 0x%02X, was 0x%02X", a, b)
	}
	return nil
}

func TestReadFully(t *testing.T) {
	reader, decoder := initTest([]byte{0x12, 0x34, 0x56, 0x78})
	var buf [4]byte
	if err := decoder.readFully(reader, buf[:]); err != nil {
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
	reader, decoder := initTest([]byte{0x12, 0x34})
	short, err := decoder.readShort(reader, binary.LittleEndian)
	if err != nil {
		t.Error(err)
	}
	if err := testShortEquals(short, 0x3412); err != nil {
		t.Error(err)
	}
}

func TestReadShortBigEndian(t *testing.T) {
	reader, decoder := initTest([]byte{0x12, 0x34})
	short, err := decoder.readShort(reader, binary.BigEndian)
	if err != nil {
		t.Error(err)
	}
	if err := testShortEquals(short, 0x1234); err != nil {
		t.Error(err)
	}
}

func TestReadShortUnexpectedEOF(t *testing.T) {
	reader, decoder := initTest([]byte{0x12})
	_, err := decoder.readShort(reader, binary.BigEndian)
	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected io.ErrUnexpectedEOF, was %v", err)
	}
}

func TestReadShortsLittleEndian(t *testing.T) {
	reader, decoder := initTest([]byte{0x12, 0x34, 0x56, 0x78})
	shorts, err := decoder.readShorts(reader, 4, binary.LittleEndian)
	if err != nil {
		t.Error(err)
	}
	if err := testShortEquals(shorts[0], 0x3412); err != nil {
		t.Error(err)
	}
	if err := testShortEquals(shorts[1], 0x7856); err != nil {
		t.Error(err)
	}
}

func TestReadShortsBigEndian(t *testing.T) {
	reader, decoder := initTest([]byte{0x12, 0x34, 0x56, 0x78})
	shorts, err := decoder.readShorts(reader, 4, binary.BigEndian)
	if err != nil {
		t.Error(err)
	}
	if err := testShortEquals(shorts[0], 0x1234); err != nil {
		t.Error(err)
	}
	if err := testShortEquals(shorts[1], 0x5678); err != nil {
		t.Error(err)
	}
}

func TestReadShortsUnexpectedEOF(t *testing.T) {
	reader, decoder := initTest([]byte{0x12})
	_, err := decoder.readShorts(reader, 4, binary.LittleEndian)
	if err != io.ErrUnexpectedEOF {
		t.Errorf("expected io.ErrUnexpectedEOF, was %v", err)
	}
}
