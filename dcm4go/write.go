package dcm4go

import (
	"encoding/binary"
	"io"
	"math"
)

// writes bytes
func writeBytes(writer io.Writer, buf []byte) error {
	if _, err := writer.Write(buf); err != nil {
		return err
	}
	return nil
}

// writes a byte
func writeByte(writer io.Writer, b byte) error {
	var buf [1]byte
	buf[0] = b
	if err := writeBytes(writer, buf[:]); err != nil {
		return err
	}
	return nil
}

// writes an unsigned short
func writeShort(writer io.Writer, short uint16, byteOrder binary.ByteOrder) error {
	var buf [2]byte
	byteOrder.PutUint16(buf[:], short)
	if err := writeBytes(writer, buf[:]); err != nil {
		return err
	}
	return nil
}

// writes an unsigned long
func writeLong(writer io.Writer, long uint32, byteOrder binary.ByteOrder) error {
	var buf [4]byte
	byteOrder.PutUint32(buf[:], long)
	if err := writeBytes(writer, buf[:]); err != nil {
		return err
	}
	return nil
}

// writes an unsigned very long
func writeVeryLong(writer io.Writer, veryLong uint64, byteOrder binary.ByteOrder) error {
	var buf [8]byte
	byteOrder.PutUint64(buf[:], veryLong)
	if err := writeBytes(writer, buf[:]); err != nil {
		return err
	}
	return nil
}

//return math.Float64frombits(byteOrder.Uint64(buf[:])), nil

// writes a float
func writeFloat(writer io.Writer, float float32, byteOrder binary.ByteOrder) error {
	var buf [4]byte
	byteOrder.PutUint32(buf[:], math.Float32bits(float))
	if err := writeBytes(writer, buf[:]); err != nil {
		return err
	}
	return nil
}

// writes a double
func writeDouble(writer io.Writer, double float64, byteOrder binary.ByteOrder) error {
	var buf [8]byte
	byteOrder.PutUint64(buf[:], math.Float64bits(double))
	if err := writeBytes(writer, buf[:]); err != nil {
		return err
	}
	return nil
}

// writeText writes a single text
func writeText(writer io.Writer, text string) error {
	buf := []byte(text)
	if err := writeBytes(writer, buf); err != nil {
		return err
	}
	return nil
}

// writeUID writes a single UID
func writeUID(writer io.Writer, text string) error {
	buf := []byte(text)
	if err := writeBytes(writer, buf); err != nil {
		return err
	}
	return nil
}
