package dcm4go

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
)

func initDecoderTest(buf []byte) (*CountingReader, *Decoder) {
	return newCountingReader(bytes.NewReader(buf)), newDecoder(1024)
}

func TestNewDecoder(t *testing.T) {
	_, decoder := initDecoderTest([]byte{})
	if decoder.bulkDataThreshold != 1024 {
		t.Errorf("expected %d, was %d", 1024, decoder.bulkDataThreshold)
	}
}

func TestReadShortsLittleEndian(t *testing.T) {
	reader, decoder := initDecoderTest([]byte{0x12, 0x34, 0x56, 0x78})
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
	reader, decoder := initDecoderTest([]byte{0x12, 0x34, 0x56, 0x78})
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
	reader, decoder := initDecoderTest([]byte{0x12})
	_, err := decoder.readShorts(reader, 4, binary.LittleEndian)
	if err := testErrEquals(err, io.ErrUnexpectedEOF); err != nil {
		t.Error(err)
	}
}
