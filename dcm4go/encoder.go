package dcm4go

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Encoder decodes a stream of bytes as a DICOM object.
type Encoder struct {
}

// newEncoder creates a new Encoder
func newEncoder() *Encoder {
	return &Encoder{}
}

// writeObject writes an object to a writer
func (encoder *Encoder) writeObject(writer io.Writer, object *Object, explicitVR bool, byteOrder binary.ByteOrder) error {
	for _, attribute := range object.attributes {
		if err := encoder.writeAttribute(writer, attribute, explicitVR, byteOrder); err != nil {
			return err
		}
	}
	return nil
}

// writeObjectWithGroupLength writes an object to a writer, along with the
// group length for that group.  checks that all attributes are for that group
func (encoder *Encoder) writeObjectWithGroupLength(writer io.Writer, group uint16, object *Object, explicitVR bool, byteOrder binary.ByteOrder) error {

	// create a buffer to write the temporary object to
	buf := new(bytes.Buffer)

	for _, attribute := range object.attributes {

		// check that this attribute is in this group
		if toGroup(attribute.tag) != group {
			return fmt.Errorf("while writing object with group length, found attribute %s that is not in group %d", tagToString(attribute.tag), group)
		}

		if err := encoder.writeAttribute(buf, attribute, explicitVR, byteOrder); err != nil {
			return err
		}
	}

	// create an attribute for the group length
	attribute := &Attribute{toTag(group, 0x00), "UL", []uint32{uint32(buf.Len())}}

	// write the attribute to the underlying writer
	if err := encoder.writeAttribute(writer, attribute, explicitVR, byteOrder); err != nil {
		return err
	}

	// write the bytes containing the group to the underlying writer
	if err := writeBytes(writer, buf.Bytes()); err != nil {
		return err
	}

	// all is well
	return nil
}

// writeAttributes writes an object to a writer
func (encoder *Encoder) writeAttribute(writer io.Writer, attribute *Attribute, explicitVR bool, byteOrder binary.ByteOrder) error {

	// write tag
	if err := encoder.writeTag(writer, attribute, byteOrder); err != nil {
		return err
	}

	// write vr
	if err := encoder.writeVR(writer, attribute, explicitVR); err != nil {
		return err
	}

	// prepare the value
	length, err := encoder.prepareValue(attribute)
	if err != nil {
		return err
	}

	// write length
	if err := encoder.writeLength(writer, attribute, length, explicitVR, byteOrder); err != nil {
		return err
	}

	// write value
	if err := encoder.writeValue(writer, attribute, byteOrder); err != nil {
		return err
	}

	return nil
}

func (encoder *Encoder) writeTag(writer io.Writer, attribute *Attribute, byteOrder binary.ByteOrder) error {
	if err := writeShort(writer, toGroup(attribute.tag), byteOrder); err != nil {
		return err
	}
	if err := writeShort(writer, toElement(attribute.tag), byteOrder); err != nil {
		return err
	}
	return nil
}

func (encoder *Encoder) writeVR(writer io.Writer, attribute *Attribute, explicitVR bool) error {
	if explicitVR {
		if err := writeText(writer, attribute.vr); err != nil {
			return err
		}
	}
	return nil
}

func (encoder *Encoder) writeLength(writer io.Writer, attribute *Attribute, length uint32, explicitVR bool, byteOrder binary.ByteOrder) error {

	if explicitVR {

		// if explicit vr and short length, write the length as a short
		if isShortLength(attribute.vr) {
			return writeShort(writer, uint16(length), byteOrder)
		}

		// if explicit vr but not short length, need to write a short zero before writing length as a long
		if err := writeShort(writer, 0x00, byteOrder); err != nil {
			return err
		}
	}

	// if not explicit vr or explicit vr but not short length, write length as a long
	if err := writeLong(writer, length, byteOrder); err != nil {
		return err
	}

	// all is well
	return nil
}

func (encoder *Encoder) writeValue(writer io.Writer, attribute *Attribute, byteOrder binary.ByteOrder) error {
	switch value := attribute.value.(type) {
	case []uint16:
		return encoder.writeShorts(writer, value, byteOrder)
	case []uint32:
		return encoder.writeLongs(writer, value, byteOrder)
	case string:
		switch attribute.vr {
		case "UI":
			return writeUID(writer, value)
		default:
			return writeText(writer, value)
		}
	default:
		return fmt.Errorf("Encoder.writeValue: not implemented for value %v", attribute.value)
	}
}

func (encoder *Encoder) writeShorts(writer io.Writer, shorts []uint16, byteOrder binary.ByteOrder) error {
	for _, short := range shorts {
		if err := writeShort(writer, short, byteOrder); err != nil {
			return err
		}
	}
	return nil
}

func (encoder *Encoder) writeLongs(writer io.Writer, longs []uint32, byteOrder binary.ByteOrder) error {
	for _, long := range longs {
		if err := writeLong(writer, long, byteOrder); err != nil {
			return err
		}
	}
	return nil
}

func (encoder *Encoder) prepareValue(attribute *Attribute) (uint32, error) {
	var length int
	switch value := attribute.value.(type) {
	case []uint16:
		length = len(value) * 2
	case []uint32:
		length = len(value) * 4
	case []float32:
		length = len(value) * 4
	case []float64:
		length = len(value) * 8
	case []string:
		switch attribute.vr {
		case "UI":
			s := flattenString(value, byte(0x00))
			length = len(s)
			attribute.value = s
		default:
			s := flattenString(value, byte(' '))
			length = len(s)
			attribute.value = s
		}
	default:
		return 0, fmt.Errorf("Encoder.prepareValue: not implemented for value %v", attribute.value)
	}

	return uint32(length), nil
}

func flattenString(values []string, padding byte) string {
	s := ""
	for _, value := range values {
		s += value + "\\"
	}
	s = strings.TrimSuffix(s, "\\")
	if isOdd(len(s)) {
		s += string(padding)
	}
	return s
}

// isOdd returns true if num is odd
func isOdd(num int) bool {
	return num&0x01 != 0
}
