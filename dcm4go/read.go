package dcm4go

import (
	"encoding/binary"
	"io"
	"io/ioutil"
	"math"
)

// reads bytes
func readBytes(reader io.Reader, length uint32) ([]byte, error) {
	buf := make([]byte, length)
	if _, err := io.ReadFull(reader, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// skips bytes
func skipBytes(reader io.Reader, length uint32) error {
	_, err := io.CopyN(ioutil.Discard, reader, int64(length))
	return err
}

// reads a byte
func readByte(reader io.Reader) (byte, error) {
	var buf [1]byte
	if _, err := io.ReadFull(reader, buf[:]); err != nil {
		return 0, err
	}
	return buf[0], nil
}

// skips a byte
func skipByte(reader io.Reader) error {
	_, err := readByte(reader)
	return err
}

// reads an unsigned short
func readShort(reader io.Reader, byteOrder binary.ByteOrder) (uint16, error) {
	var buf [2]byte
	if _, err := io.ReadFull(reader, buf[:]); err != nil {
		return 0, err
	}
	return byteOrder.Uint16(buf[:]), nil
}

// reads an unsigned long
func readLong(reader io.Reader, byteOrder binary.ByteOrder) (uint32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(reader, buf[:]); err != nil {
		return 0, err
	}
	return byteOrder.Uint32(buf[:]), nil
}

// reads an unsigned very long
func readVeryLong(reader io.Reader, byteOrder binary.ByteOrder) (uint64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(reader, buf[:]); err != nil {
		return 0, err
	}
	return byteOrder.Uint64(buf[:]), nil
}

// reads a float
func readFloat(reader io.Reader, byteOrder binary.ByteOrder) (float32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(reader, buf[:]); err != nil {
		return 0, err
	}
	return math.Float32frombits(byteOrder.Uint32(buf[:])), nil
}

// reads a double
func readDouble(reader io.Reader, byteOrder binary.ByteOrder) (float64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(reader, buf[:]); err != nil {
		return 0, err
	}
	return math.Float64frombits(byteOrder.Uint64(buf[:])), nil
}

func readUID(reader io.Reader, length uint32) (string, error) {
	buf, err := readBytes(reader, length)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func readText(reader io.Reader, length uint32) (string, error) {
	buf, err := readBytes(reader, length)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
