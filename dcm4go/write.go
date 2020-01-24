package dcm4go

import (
	"encoding/binary"
	"io"
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
