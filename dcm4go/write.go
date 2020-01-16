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

// // reads an unsigned very long
// func readVeryLong(reader io.Reader, byteOrder binary.ByteOrder) (uint64, error) {
// 	var buf [8]byte
// 	if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 		return 0, err
// 	}
// 	return byteOrder.Uint64(buf[:]), nil
// }
//
// // reads a float
// func readFloat(reader io.Reader, byteOrder binary.ByteOrder) (float32, error) {
// 	var buf [4]byte
// 	if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 		return 0, err
// 	}
// 	return math.Float32frombits(byteOrder.Uint32(buf[:])), nil
// }
//
// // reads a double
// func readDouble(reader io.Reader, byteOrder binary.ByteOrder) (float64, error) {
// 	var buf [8]byte
// 	if _, err := io.ReadFull(reader, buf[:]); err != nil {
// 		return 0, err
// 	}
// 	return math.Float64frombits(byteOrder.Uint64(buf[:])), nil
// }
//
// // readUID reads a single UID from a reader
// func readUID(reader io.Reader, length uint32) (string, error) {
// 	buf, err := readBytes(reader, length)
// 	if err != nil {
// 		return "", err
// 	}
// 	return removeUIDPadding(buf), nil
// }
//
// // removeUIDPadding removes the padding from the UID if any
// func removeUIDPadding(buf []byte) string {
// 	if len(buf) > 0 && buf[len(buf)-1] == 0x00 {
// 		return string(buf[:len(buf)-1])
// 	}
// 	return string(buf)
// }
//

// writeText writes a single text
func writeText(writer io.Writer, text string) error {
	buf := []byte(text)
	if err := writeBytes(writer, buf); err != nil {
		return err
	}
	return nil
}

// isOdd returns true if num is odd
func isOdd(num int) bool {
	return num&0x01 != 0
}

// addTextPadding adds padding if required to make the text length even
func addTextPadding(text string) string {
	if isOdd(len(text)) {
		return text + " "
	}
	return text
}

// writeText writes a single text, padding if required
func writePaddedText(writer io.Writer, text string) error {
	if err := writeText(writer, addTextPadding(text)); err != nil {
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

//
// // removeTextPadding removes the padding from the text if any
// func removeTextPadding(buf []byte) string {
// 	if len(buf) > 0 && buf[len(buf)-1] == byte(' ') {
// 		return string(buf[:len(buf)-1])
// 	}
// 	return string(buf)
// }
//
// // checks if odd
// func isOdd(num int) bool {
// 	return num&1 != 0
// }
