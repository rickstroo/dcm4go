package dcm4go

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestNewDecoder(t *testing.T) {
	decoder := newDecoder(1024)
	if decoder.bulkDataThreshold != 1024 {
		t.Errorf("decoder.bulkDataThreshold: was %d, expected %d", decoder.bulkDataThreshold, 1024)
	}
}

func TestReadShortLittleEndian(t *testing.T) {
	buf := []byte{0x12, 0x34}
	reader := newCountingReader(bytes.NewReader(buf))
	decoder := newDecoder(1024)
	short, err := decoder.readShort(reader, binary.LittleEndian)
	if err != nil {
		t.Error(err)
	}
	if short != 0x3412 {
		t.Errorf("decoder.readShort: expected 0x%04X, was 0x%04X", 0x3412, short)
	}
}

func TestReadShortBigEndian(t *testing.T) {
	buf := []byte{0x12, 0x34}
	reader := newCountingReader(bytes.NewReader(buf))
	decoder := newDecoder(1024)
	short, err := decoder.readShort(reader, binary.BigEndian)
	if err != nil {
		t.Error(err)
	}
	if short != 0x1234 {
		t.Errorf("decoder.readShort: expected 0x%04X, was 0x%04X", 0x1234, short)
	}
}
