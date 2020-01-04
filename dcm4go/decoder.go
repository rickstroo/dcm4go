package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"
)

// Decoder decodes a stream of bytes as a DICOM object.
// It also keeps track of the number of bytes read so
// that encoders can use that information to generate
// offset information as required.  It seems like that
// information ought to go to into a derived class, but
// it is easier to just include it in this classs for now.
type Decoder struct {
	bytesRead uint32
}

// NewDecoder creates a new Decoder
func newDecoder() *Decoder {
	return &Decoder{0}
}

// ReadObject reads a DICOM object from a reader
func (decoder *Decoder) readObject(reader io.Reader, explicitVR bool, byteOrder binary.ByteOrder) (*Object, error) {

	object := newObject()

	for {
		attribute, err := decoder.readAttribute(reader, explicitVR, byteOrder)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		object.add(attribute)
	}

	return object, nil
}

// ReadAttribute reads a DICOM attribute from a reader
func (decoder *Decoder) readAttribute(reader io.Reader, explicitVR bool, byteOrder binary.ByteOrder) (*Attribute, error) {

	group, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	element, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	// item delimiter tag
	if group == 0xFFFE && element == 0xE00D {

		// need to consume the length
		_, err := decoder.readLong(reader, byteOrder)
		if err != nil {
			return nil, err
		}

		// let the call know that there are no more attributes
		return nil, io.EOF
	}

	vr, err := decoder.readVR(reader, group, element, explicitVR)
	if err != nil {
		return nil, err
	}

	length, err := decoder.readLength(reader, explicitVR, byteOrder, vr)
	if err != nil {
		return nil, err
	}

	// remember the offset of this attribute's value
	offset := decoder.bytesRead

	value, err := decoder.readValue(reader, explicitVR, byteOrder, vr, length)
	if err != nil {
		return nil, err
	}

	return &Attribute{group, element, vr, length, offset, value}, nil
}

// Read reads bytes into a buffer
func (decoder *Decoder) readFully(reader io.Reader, buf []byte) error {
	num, err := io.ReadFull(reader, buf)
	decoder.bytesRead += uint32(num)
	if err != nil {
		return err
	}
	return nil
}

// reads an unsigned short
func (decoder *Decoder) readShort(reader io.Reader, byteOrder binary.ByteOrder) (uint16, error) {
	var buf [2]byte
	err := decoder.readFully(reader, buf[:])
	if err != nil {
		return 0, err
	}
	value := byteOrder.Uint16(buf[:])
	return value, nil
}

// reads an unsigned long
func (decoder *Decoder) readLong(reader io.Reader, byteOrder binary.ByteOrder) (uint32, error) {
	var buf [4]byte
	err := decoder.readFully(reader, buf[:])
	if err != nil {
		return 0, err
	}
	value := byteOrder.Uint32(buf[:])
	return value, nil
}

// reads the vr of an attribute
func (decoder *Decoder) readVR(reader io.Reader, group uint16, element uint16, explicitVR bool) (string, error) {
	if explicitVR {
		var buf [2]byte
		err := decoder.readFully(reader, buf[:])
		if err != nil {
			return "", err
		}
		return string(buf[:]), nil
	}
	vr, err := findVR(group, element)
	if err != nil {
		return "", err
	}
	return vr, nil
}

// finds a vr in a dictionary
func findVR(group uint16, element uint16) (string, error) {
	vr, ok := vrs[uint32(group<<8)|uint32(element)]
	if !ok {
		return "", fmt.Errorf("unable to find vr for tag %04x%04x", group, element)
	}
	return vr, nil
}

// reads the length of an attribute
func (decoder *Decoder) readLength(reader io.Reader, explicitVR bool, byteOrder binary.ByteOrder, vr string) (uint32, error) {

	if explicitVR {

		// if explicit vr and short length, read the length as a short
		if isShortLength(vr) {
			length, err := decoder.readShort(reader, byteOrder)
			if err != nil {
				return 0, err
			}
			return uint32(length), nil
		}

		// if explicit vr but not short length, need to skip 2 bytes before reading length as a long
		var buf [2]byte
		err := decoder.readFully(reader, buf[:])
		if err != nil {
			return 0, err
		}
	}

	// if not explicit vr or explicit vr but not short length, read length as a long
	length, err := decoder.readLong(reader, byteOrder)
	if err != nil {
		return 0, err
	}

	return length, nil
}

// determines if the attribute has a short or long length
func isShortLength(vr string) bool {
	switch vr {
	case "AE", "AS", "AT", "CS", "DA", "DS", "DT", "FL", "FD", "IS", "LO", "LT", "PN", "SH", "SL", "SS", "ST", "TM", "UI", "UL", "US":
		return true
	}
	return false
}

// readValue reads the value of an attribute
func (decoder *Decoder) readValue(reader io.Reader, explicitVR bool, byteOrder binary.ByteOrder, vr string, length uint32) (interface{}, error) {
	switch vr {
	// not all of these VRs interpret the backslash as a separator
	case "AE", "AS", "CS", "DA", "DS", "DT", "LO", "SH", "TM", "UC", "UR":
		return decoder.readTexts(reader, length)
	case "AT":
		return decoder.readTags(reader, length, byteOrder)
	case "FD", "OD":
		return decoder.readDoubles(reader, length, byteOrder)
	case "FL", "OF":
		return decoder.readFloats(reader, length, byteOrder)
	case "IS":
		return decoder.readTexts(reader, length)
	case "LT", "ST", "UT":
		return decoder.readText(reader, length)
	case "PN":
		return decoder.readTexts(reader, length)
	case "OB":
		if length == 0xFFFFFFFF {
			return decoder.readFragments(reader, byteOrder)
		}
		return decoder.readBytes(reader, length)
	case "OL", "OV", "OW":
		return decoder.readBytes(reader, length)
	case "SL", "UL":
		return decoder.readLongs(reader, length, byteOrder)
	case "SQ":
		return decoder.readSequence(reader, length, explicitVR, byteOrder)
	case "SS", "US":
		return decoder.readShorts(reader, length, byteOrder)
	case "SV", "UV":
		return decoder.readVeryLongs(reader, length, byteOrder)
	case "UI":
		return decoder.readUIDs(reader, length)
	case "UN":
		return decoder.readBytes(reader, length)
	}
	return nil, fmt.Errorf("unrecognized vr, '%s'", vr)
}

// reads unsigned shorts
func (decoder *Decoder) readShorts(reader io.Reader, length uint32, byteOrder binary.ByteOrder) ([]uint16, error) {
	shorts := make([]uint16, length/2)
	for i := 0; i < len(shorts); i++ {
		short, err := decoder.readShort(reader, byteOrder)
		if err != nil {
			return nil, err
		}
		shorts[i] = short
	}
	return shorts, nil
}

// reads unsigned longs
func (decoder *Decoder) readLongs(reader io.Reader, length uint32, byteOrder binary.ByteOrder) ([]uint32, error) {
	longs := make([]uint32, length/4)
	for i := 0; i < len(longs); i++ {
		long, err := decoder.readLong(reader, byteOrder)
		if err != nil {
			return nil, err
		}
		longs[i] = long
	}
	return longs, nil
}

// reads an unsigned very long
func (decoder *Decoder) readVeryLong(reader io.Reader, byteOrder binary.ByteOrder) (uint64, error) {
	var buf [8]byte
	err := decoder.readFully(reader, buf[:])
	if err != nil {
		return 0, err
	}
	value := byteOrder.Uint64(buf[:])
	return value, nil
}

// reads unsigned very longs
func (decoder *Decoder) readVeryLongs(reader io.Reader, length uint32, byteOrder binary.ByteOrder) ([]uint64, error) {
	veryLongs := make([]uint64, length/8)
	for i := 0; i < len(veryLongs); i++ {
		veryLong, err := decoder.readVeryLong(reader, byteOrder)
		if err != nil {
			return nil, err
		}
		veryLongs[i] = veryLong
	}
	return veryLongs, nil
}

// reads a float
func (decoder *Decoder) readFloat(reader io.Reader, byteOrder binary.ByteOrder) (float32, error) {
	var buf [4]byte
	err := decoder.readFully(reader, buf[:])
	if err != nil {
		return 0, err
	}
	value := math.Float32frombits(byteOrder.Uint32(buf[:]))
	return value, nil
}

// reads floats
func (decoder *Decoder) readFloats(reader io.Reader, length uint32, byteOrder binary.ByteOrder) ([]float32, error) {
	floats := make([]float32, length/4)
	for i := 0; i < len(floats); i++ {
		float, err := decoder.readFloat(reader, byteOrder)
		if err != nil {
			return nil, err
		}
		floats[i] = float
	}
	return floats, nil
}

// reads a double
func (decoder *Decoder) readDouble(reader io.Reader, byteOrder binary.ByteOrder) (float64, error) {
	var buf [8]byte
	err := decoder.readFully(reader, buf[:])
	if err != nil {
		return 0, err
	}
	value := math.Float64frombits(byteOrder.Uint64(buf[:]))
	return value, nil
}

// reads doubles
func (decoder *Decoder) readDoubles(reader io.Reader, length uint32, byteOrder binary.ByteOrder) ([]float64, error) {
	doubles := make([]float64, length/8)
	for i := 0; i < len(doubles); i++ {
		double, err := decoder.readDouble(reader, byteOrder)
		if err != nil {
			return nil, err
		}
		doubles[i] = double
	}
	return doubles, nil
}

// parseUID parses a UID from an attribute value
func parseUID(value []byte) string {
	if len(value) > 0 && value[len(value)-1] == 0x00 {
		return string(value[:len(value)-1])
	}
	return string(value)
}

// readUIDs reads UIDs from a reader
func (decoder *Decoder) readUIDs(reader io.Reader, length uint32) ([]string, error) {
	buf := make([]byte, int(length))
	err := decoder.readFully(reader, buf)
	if err != nil {
		return nil, err
	}
	return strings.Split(parseUID(buf), "\\"), nil
}

// parseText parses text from an attribute value
func parseText(value []byte) string {
	if len(value) > 0 && value[len(value)-1] == byte(' ') {
		return string(value[:len(value)-1])
	}
	return string(value)
}

// readTexts read texts from a reader
func (decoder *Decoder) readTexts(reader io.Reader, length uint32) ([]string, error) {
	buf := make([]byte, int(length))
	err := decoder.readFully(reader, buf)
	if err != nil {
		return nil, err
	}
	return strings.Split(parseText(buf), "\\"), nil
}

// readText read text from a reader
func (decoder *Decoder) readText(reader io.Reader, length uint32) (string, error) {
	buf := make([]byte, int(length))
	err := decoder.readFully(reader, buf)
	if err != nil {
		return "", err
	}
	return parseText(buf), nil
}

// reads bytes
func (decoder *Decoder) readBytes(reader io.Reader, length uint32) ([]byte, error) {
	buf := make([]byte, int(length))
	err := decoder.readFully(reader, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (decoder *Decoder) readSequence(reader io.Reader, length uint32, explicitVR bool, byteOrder binary.ByteOrder) (*Sequence, error) {

	if length == 0xFFFFFFFF {
		return decoder.readSequenceItems(reader, explicitVR, byteOrder)
	}

	return decoder.readSequenceItems(io.LimitReader(reader, int64(length)), explicitVR, byteOrder)
}

func (decoder *Decoder) readSequenceItems(reader io.Reader, explicitVR bool, byteOrder binary.ByteOrder) (*Sequence, error) {

	// create a sequence
	sequence := newSequence()

	// read the sequence items
	for {

		object, err := decoder.readSequenceItem(reader, explicitVR, byteOrder)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// add the object to the sequence
		sequence.add(object)
	}

	return sequence, nil
}

func (decoder *Decoder) readSequenceItem(reader io.Reader, explicitVR bool, byteOrder binary.ByteOrder) (*Object, error) {

	group, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	element, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	length, err := decoder.readLong(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	// sequence delimitation item
	if group == 0xFFFE && element == 0xE0DD {
		return nil, io.EOF
	}

	// item tag
	if group != 0xFFFE || element != 0xE000 {
		return nil, fmt.Errorf("expecting item tag at beginning of sequence item, found (0x%04x,0x%04x) instead", group, element)
	}

	if length == 0xFFFFFFFF {
		object, err := decoder.readObject(reader, explicitVR, byteOrder)
		if err != nil {
			return nil, err
		}
		return object, nil
	}

	object, err := decoder.readObject(io.LimitReader(reader, int64(length)), explicitVR, byteOrder)
	if err != nil {
		return nil, err
	}
	return object, nil

}

func (decoder *Decoder) readFragments(reader io.Reader, byteOrder binary.ByteOrder) ([][]byte, error) {

	// create a sequence
	fragments := make([][]byte, 0)

	// read the sequence items
	for {

		bytes, err := decoder.readFragment(reader, byteOrder)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// add the object to the sequence
		fragments = append(fragments, bytes)
	}

	return fragments, nil
}

func (decoder *Decoder) readFragment(reader io.Reader, byteOrder binary.ByteOrder) ([]byte, error) {

	group, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	element, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	length, err := decoder.readLong(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	// fragment delimitation item
	if group == 0xFFFE && element == 0xE0DD {
		return nil, io.EOF
	}

	// item tag
	if group != 0xFFFE || element != 0xE000 {
		return nil, fmt.Errorf("expecting item tag at beginning of fragment, found (0x%04x,0x%04x) instead", group, element)
	}

	bytes, err := decoder.readBytes(reader, length)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (decoder *Decoder) readTag(reader io.Reader, byteOrder binary.ByteOrder) (*Tag, error) {
	group, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return nil, err
	}
	element, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return nil, err
	}
	return toTag(group, element), nil
}

func (decoder *Decoder) readTags(reader io.Reader, length uint32, byteOrder binary.ByteOrder) ([]*Tag, error) {
	tags := make([]*Tag, length/4)
	for i := 0; i < len(tags); i++ {
		tag, err := decoder.readTag(reader, byteOrder)
		if err != nil {
			return nil, err
		}
		tags[i] = tag
	}
	return tags, nil
}
