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
func (encoder *Encoder) writeObject(writer io.Writer, object *Object, transferSyntax *TransferSyntax) error {
	for _, attribute := range object.attributes {
		if err := encoder.writeAttribute(writer, attribute, transferSyntax); err != nil {
			return err
		}
	}
	return nil
}

// writeObjectWithGroupLength writes an object to a writer, along with the
// group length for that group.  checks that all attributes are for that group
func (encoder *Encoder) writeObjectWithGroupLength(writer io.Writer, group uint16, object *Object, transferSyntax *TransferSyntax) error {

	// create a buffer to write the temporary object to
	buf := new(bytes.Buffer)

	for _, attribute := range object.attributes {

		// check that this attribute is in this group
		if toGroup(attribute.tag) != group {
			return fmt.Errorf("while writing object with group length, found attribute %s that is not in group %d", tagToString(attribute.tag), group)
		}

		if err := encoder.writeAttribute(buf, attribute, transferSyntax); err != nil {
			return err
		}
	}

	// create an attribute for the group length
	attribute := &Attribute{toTag(group, 0x00), "UL", []uint32{uint32(buf.Len())}}

	// write the attribute to the underlying writer
	if err := encoder.writeAttribute(writer, attribute, transferSyntax); err != nil {
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
func (encoder *Encoder) writeAttribute(writer io.Writer, attribute *Attribute, transferSyntax *TransferSyntax) error {

	// write tag
	if err := encoder.writeTag(writer, attribute.tag, transferSyntax.byteOrder); err != nil {
		return err
	}

	// write vr
	if err := encoder.writeVR(writer, attribute, transferSyntax.explicitVR); err != nil {
		return err
	}

	// create a buffer to write the value into
	buf := new(bytes.Buffer)

	// write value into buffer
	if err := encoder.writeValue(buf, attribute, transferSyntax); err != nil {
		return err
	}

	// write length
	if err := encoder.writeLength(writer, attribute, uint32(buf.Len()), transferSyntax.explicitVR, transferSyntax.byteOrder); err != nil {
		return err
	}

	// write the value
	if err := writeBytes(writer, buf.Bytes()); err != nil {
		return err
	}

	return nil
}

func (encoder *Encoder) writeTag(writer io.Writer, tag uint32, byteOrder binary.ByteOrder) error {
	if err := writeShort(writer, toGroup(tag), byteOrder); err != nil {
		return err
	}
	if err := writeShort(writer, toElement(tag), byteOrder); err != nil {
		return err
	}
	return nil
}

func (encoder *Encoder) writeVR(writer io.Writer, attribute *Attribute, explicitVR bool) error {
	if explicitVR {
		if err := writeString(writer, attribute.vr); err != nil {
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

func (encoder *Encoder) writeValue(writer io.Writer, attribute *Attribute, transferSyntax *TransferSyntax) error {

	switch attribute.vr {
	// these VRs support multiple text values
	case "AE", "AS", "CS", "DA", "DS", "DT", "IS", "LO", "PN", "SH", "TM", "UC":
		return encoder.writePaddedTexts(writer, attribute.value.([]string))
	case "AT":
		return encoder.writeTags(writer, attribute.value.([]uint32), transferSyntax.byteOrder)
	case "FD", "OD":
		return encoder.writeDoubles(writer, attribute.value.([]float64), transferSyntax.byteOrder)
	case "FL", "OF":
		return encoder.writeFloats(writer, attribute.value.([]float32), transferSyntax.byteOrder)
		// these VRs support single text values
	case "LT", "ST", "UT", "UR":
		return encoder.writePaddedText(writer, attribute.value.(string))
	case "OB", "OL", "OV", "OW", "UN":
		return encoder.writePixelData(writer, attribute.value, transferSyntax.byteOrder)
	case "SL", "UL":
		return encoder.writeLongs(writer, attribute.value.([]uint32), transferSyntax.byteOrder)
	case "SQ":
		return encoder.writeSequence(writer, attribute.value.(*Sequence), transferSyntax)
	case "SS", "US":
		return encoder.writeShorts(writer, attribute.value.([]uint16), transferSyntax.byteOrder)
	case "SV", "UV":
		return encoder.writeVeryLongs(writer, attribute.value.([]uint64), transferSyntax.byteOrder)
	case "UI":
		return encoder.writePaddedUIDs(writer, attribute.value.([]string))
	}

	// hmm, didn't recognize the VR
	return ErrUnrecognizedVR
}

func (encoder *Encoder) writePaddedTexts(writer io.Writer, texts []string) error {
	text := flattenStrings(texts)
	if err := encoder.writePaddedText(writer, text); err != nil {
		return err
	}
	return nil
}

// flattenStrings converts an array of strings into a single string with separataor character
func flattenStrings(texts []string) string {
	t := ""
	for _, text := range texts {
		t += text + "\\"
	}
	t = strings.TrimSuffix(t, "\\")
	return t
}

func (encoder *Encoder) writePaddedText(writer io.Writer, text string) error {
	paddedText := padString(text, string(byte(' ')))
	if err := writeString(writer, paddedText); err != nil {
		return err
	}
	return nil
}

// padString returns a string that is padded
func padString(str string, padding string) string {
	if isOdd(len(str)) {
		return str + padding
	}
	return str
}

// isOdd returns true if num is odd
func isOdd(num int) bool {
	return num&0x01 != 0
}

func (encoder *Encoder) writePaddedUIDs(writer io.Writer, uids []string) error {
	uid := flattenStrings(uids)
	if err := encoder.writePaddedUID(writer, uid); err != nil {
		return err
	}
	return nil
}

func (encoder *Encoder) writePaddedUID(writer io.Writer, uid string) error {
	paddedUID := padString(uid, string(byte(0x00)))
	if err := writeString(writer, paddedUID); err != nil {
		return err
	}
	return nil
}

// writes tags
func (encoder *Encoder) writeTags(writer io.Writer, tags []uint32, byteOrder binary.ByteOrder) error {
	for _, tag := range tags {
		if err := encoder.writeTag(writer, tag, byteOrder); err != nil {
			return err
		}
	}
	return nil
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

func (encoder *Encoder) writeVeryLongs(writer io.Writer, veryLongs []uint64, byteOrder binary.ByteOrder) error {
	for _, veryLong := range veryLongs {
		if err := writeVeryLong(writer, veryLong, byteOrder); err != nil {
			return err
		}
	}
	return nil
}

func (encoder *Encoder) writeFloats(writer io.Writer, floats []float32, byteOrder binary.ByteOrder) error {
	for _, float := range floats {
		if err := writeFloat(writer, float, byteOrder); err != nil {
			return err
		}
	}
	return nil
}

func (encoder *Encoder) writeDoubles(writer io.Writer, doubles []float64, byteOrder binary.ByteOrder) error {
	for _, double := range doubles {
		if err := writeDouble(writer, double, byteOrder); err != nil {
			return err
		}
	}
	return nil
}

func (encoder *Encoder) writeSequence(writer io.Writer, sequence *Sequence, transferSyntax *TransferSyntax) error {
	return fmt.Errorf("encoder.writeSequence not implemented")
}

func (encoder *Encoder) writePixelData(writer io.Writer, pixelData interface{}, byteOrder binary.ByteOrder) error {
	return fmt.Errorf("encoder.writePixelData not implemented")
}
