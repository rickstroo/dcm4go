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
	attribute := &Attribute{toTag(group, 0x00), "UL", 4, []uint32{uint32(buf.Len())}}

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

	// write length
	if err := encoder.writeLength(writer, attribute, explicitVR, byteOrder); err != nil {
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

func (encoder *Encoder) writeLength(writer io.Writer, attribute *Attribute, explicitVR bool, byteOrder binary.ByteOrder) error {

	if explicitVR {

		// if explicit vr and short length, write the length as a short
		if isShortLength(attribute.vr) {
			return writeShort(writer, uint16(attribute.length), byteOrder)
		}

		// if explicit vr but not short length, need to write a short zero before writing length as a long
		if err := writeShort(writer, 0x00, byteOrder); err != nil {
			return err
		}
	}

	// if not explicit vr or explicit vr but not short length, write length as a long
	if err := writeLong(writer, attribute.length, byteOrder); err != nil {
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
	case []string:
		switch attribute.vr {
		case "UI":
			return encoder.writeUIDs(writer, value)
		default:
			return fmt.Errorf("Encoder.writeValue: not implemented for value %v and VR %s", attribute.value, attribute.vr)
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

func convertToMultiValueUID(uids []string) string {
	mvUID := ""
	for _, uid := range uids {
		mvUID += uid + "\\"
	}
	mvUID = strings.TrimSuffix(mvUID, "\\")
	if isOdd(len(mvUID)) {
		mvUID += string(byte(0x00))
	}
	return mvUID
}

func (encoder *Encoder) writeUIDs(writer io.Writer, uids []string) error {
	mvUID := convertToMultiValueUID(uids)
	if err := writeUID(writer, mvUID); err != nil {
		return err
	}
	return nil
}
