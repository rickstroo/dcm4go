package dcm4go

import (
	"encoding/binary"
	"io"
	"math"
)

// reads bytes
func readBytes(reader io.Reader, buf []byte) error {
	if _, err := io.ReadFull(reader, buf); err != nil {
		return err
	}
	return nil
}

// reads a byte
func readByte(reader io.Reader) (byte, error) {
	var buf [1]byte
	if err := readBytes(reader, buf[:]); err != nil {
		return 0, err
	}
	return buf[0], nil
}

// reads an unsigned short
func readShort(reader io.Reader, byteOrder binary.ByteOrder) (uint16, error) {
	var buf [2]byte
	if err := readBytes(reader, buf[:]); err != nil {
		return 0, err
	}
	return byteOrder.Uint16(buf[:]), nil
}

// reads an unsigned long
func readLong(reader CounterReader, byteOrder binary.ByteOrder) (uint32, error) {
	var buf [4]byte
	if err := readBytes(reader, buf[:]); err != nil {
		return 0, err
	}
	return byteOrder.Uint32(buf[:]), nil
}

// reads an unsigned very long
func readVeryLong(reader CounterReader, byteOrder binary.ByteOrder) (uint64, error) {
	var buf [8]byte
	if err := readBytes(reader, buf[:]); err != nil {
		return 0, err
	}
	return byteOrder.Uint64(buf[:]), nil
}

// reads a float
func readFloat(reader CounterReader, byteOrder binary.ByteOrder) (float32, error) {
	var buf [4]byte
	if err := readBytes(reader, buf[:]); err != nil {
		return 0, err
	}
	return math.Float32frombits(byteOrder.Uint32(buf[:])), nil
}

// reads a double
func readDouble(reader CounterReader, byteOrder binary.ByteOrder) (float64, error) {
	var buf [8]byte
	if err := readBytes(reader, buf[:]); err != nil {
		return 0, err
	}
	return math.Float64frombits(byteOrder.Uint64(buf[:])), nil
}

// readUID reads a single UID from a reader
func readUID(reader CounterReader, length uint32) (string, error) {
	buf := make([]byte, length)
	if err := readBytes(reader, buf[:]); err != nil {
		return "", err
	}
	return removeUIDPadding(buf), nil
}

// removeUIDPadding removes the padding from the UID if any
func removeUIDPadding(value []byte) string {
	if len(value) > 0 && value[len(value)-1] == 0x00 {
		return string(value[:len(value)-1])
	}
	return string(value)
}

// readText reads a single text from a reader
func readText(reader CounterReader, length uint32) (string, error) {
	buf := make([]byte, length)
	if err := readBytes(reader, buf[:]); err != nil {
		return "", err
	}
	return removeTextPadding(buf), nil
}

// removeTextPadding removes the padding from the text if any
func removeTextPadding(value []byte) string {
	if len(value) > 0 && value[len(value)-1] == byte(' ') {
		return string(value[:len(value)-1])
	}
	return string(value)
}
