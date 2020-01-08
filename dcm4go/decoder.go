package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"
)

// Decoder decodes a stream of bytes as a DICOM object.
type Decoder struct {
	bulkDataThreshold uint32
}

// newDecoder creates a new Decoder
func newDecoder(bulkDataThreshold uint32) *Decoder {
	return &Decoder{bulkDataThreshold}
}

// readObject reads a DICOM object from a reader
func (decoder *Decoder) readObject(reader CounterReader, explicitVR bool, byteOrder binary.ByteOrder) (*Object, error) {

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

// readAttribute reads a DICOM attribute from a reader
func (decoder *Decoder) readAttribute(reader CounterReader, explicitVR bool, byteOrder binary.ByteOrder) (*Attribute, error) {

	tag, err := decoder.readTag(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	// item deer tag
	if tag == ItemDelimitationItemTag {

		// need to consume the length
		_, err := decoder.readLong(reader, byteOrder)
		if err != nil {
			return nil, err
		}

		// let the call know that there are no more attributes
		return nil, io.EOF
	}

	vr, err := decoder.readVR(reader, tag, explicitVR)
	if err != nil {
		return nil, err
	}

	length, err := decoder.readLength(reader, explicitVR, byteOrder, vr)
	if err != nil {
		return nil, err
	}

	// remember the offset of this attribute's value
	offset := uint32(reader.BytesRead())

	value, err := decoder.readValue(reader, explicitVR, byteOrder, vr, length)
	if err != nil {
		return nil, err
	}

	return &Attribute{tag, vr, length, offset, value}, nil
}

// reads the vr of an attribute
func (decoder *Decoder) readVR(reader CounterReader, tag uint32, explicitVR bool) (string, error) {
	if explicitVR {
		var buf [2]byte
		err := decoder.readFully(reader, buf[:])
		if err != nil {
			return "", err
		}
		return string(buf[:]), nil
	}
	vr, err := findVR(tag)
	if err != nil {
		return "", err
	}
	return vr, nil
}

// finds a vr in a dictionary
func findVR(tag uint32) (string, error) {
	vr, ok := vrs[tag]
	if !ok {
		return "", fmt.Errorf("unable to find vr for tag %s", tagToString(tag))
	}
	return vr, nil
}

// reads the length of an attribute
func (decoder *Decoder) readLength(reader CounterReader, explicitVR bool, byteOrder binary.ByteOrder, vr string) (uint32, error) {

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

// reads the value of an attribute
func (decoder *Decoder) readValue(reader CounterReader, explicitVR bool, byteOrder binary.ByteOrder, vr string, length uint32) (interface{}, error) {
	switch vr {
	// these VRs support multiple text values
	case "AE", "AS", "CS", "DA", "DS", "DT", "IS", "LO", "PN", "SH", "TM", "UC":
		return decoder.readTexts(reader, length)
	case "AT":
		return decoder.readTags(reader, length, byteOrder)
	case "FD", "OD":
		return decoder.readDoubles(reader, length, byteOrder)
	case "FL", "OF":
		return decoder.readFloats(reader, length, byteOrder)
		// these VRs support single text values
	case "LT", "ST", "UT", "UR":
		return decoder.readText(reader, length)
	case "OB":
		return decoder.readPixelData(reader, length, byteOrder)
	case "OL", "OV", "OW", "UN":
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
	}
	return nil, fmt.Errorf("unrecognized vr, '%s'", vr)
}

// reads bytes into a buffer
func (decoder *Decoder) readFully(reader CounterReader, buf []byte) error {
	_, err := io.ReadFull(reader, buf)
	return err
}

// reads tag
func (decoder *Decoder) readTag(reader CounterReader, byteOrder binary.ByteOrder) (uint32, error) {
	group, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return 0, err
	}
	element, err := decoder.readShort(reader, byteOrder)
	if err != nil {
		return 0, err
	}
	return toTag(group, element), nil
}

// reads tags
func (decoder *Decoder) readTags(reader CounterReader, length uint32, byteOrder binary.ByteOrder) ([]uint32, error) {
	tags := make([]uint32, length/4)
	for i := 0; i < len(tags); i++ {
		tag, err := decoder.readTag(reader, byteOrder)
		if err != nil {
			return nil, err
		}
		tags[i] = tag
	}
	return tags, nil
}

// reads an unsigned short
func (decoder *Decoder) readShort(reader CounterReader, byteOrder binary.ByteOrder) (uint16, error) {
	var short uint16
	if err := binary.Read(reader, byteOrder, &short); err != nil {
		return 0, err
	}
	return short, nil
}

// reads unsigned shorts
func (decoder *Decoder) readShorts(reader CounterReader, length uint32, byteOrder binary.ByteOrder) ([]uint16, error) {
	shorts := make([]uint16, length/2)
	if err := binary.Read(reader, byteOrder, shorts); err != nil {
		return nil, err
	}
	return shorts, nil
}

// reads an unsigned long
func (decoder *Decoder) readLong(reader CounterReader, byteOrder binary.ByteOrder) (uint32, error) {
	var buf [4]byte
	err := decoder.readFully(reader, buf[:])
	if err != nil {
		return 0, err
	}
	value := byteOrder.Uint32(buf[:])
	return value, nil
}

// reads unsigned longs
func (decoder *Decoder) readLongs(reader CounterReader, length uint32, byteOrder binary.ByteOrder) ([]uint32, error) {
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
func (decoder *Decoder) readVeryLong(reader CounterReader, byteOrder binary.ByteOrder) (uint64, error) {
	var buf [8]byte
	err := decoder.readFully(reader, buf[:])
	if err != nil {
		return 0, err
	}
	value := byteOrder.Uint64(buf[:])
	return value, nil
}

// reads unsigned very longs
func (decoder *Decoder) readVeryLongs(reader CounterReader, length uint32, byteOrder binary.ByteOrder) ([]uint64, error) {
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
func (decoder *Decoder) readFloat(reader CounterReader, byteOrder binary.ByteOrder) (float32, error) {
	var buf [4]byte
	err := decoder.readFully(reader, buf[:])
	if err != nil {
		return 0, err
	}
	value := math.Float32frombits(byteOrder.Uint32(buf[:]))
	return value, nil
}

// reads floats
func (decoder *Decoder) readFloats(reader CounterReader, length uint32, byteOrder binary.ByteOrder) ([]float32, error) {
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
func (decoder *Decoder) readDouble(reader CounterReader, byteOrder binary.ByteOrder) (float64, error) {
	var buf [8]byte
	err := decoder.readFully(reader, buf[:])
	if err != nil {
		return 0, err
	}
	value := math.Float64frombits(byteOrder.Uint64(buf[:]))
	return value, nil
}

// reads doubles
func (decoder *Decoder) readDoubles(reader CounterReader, length uint32, byteOrder binary.ByteOrder) ([]float64, error) {
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
func (decoder *Decoder) readUIDs(reader CounterReader, length uint32) ([]string, error) {
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
func (decoder *Decoder) readTexts(reader CounterReader, length uint32) ([]string, error) {
	buf := make([]byte, int(length))
	err := decoder.readFully(reader, buf)
	if err != nil {
		return nil, err
	}
	return strings.Split(parseText(buf), "\\"), nil
}

// readText read text from a reader
func (decoder *Decoder) readText(reader CounterReader, length uint32) (string, error) {
	buf := make([]byte, int(length))
	err := decoder.readFully(reader, buf)
	if err != nil {
		return "", err
	}
	return parseText(buf), nil
}

// reads bytes
func (decoder *Decoder) readBytes(reader CounterReader, length uint32) ([]byte, error) {
	buf := make([]byte, int(length))
	err := decoder.readFully(reader, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// UndefinedLength represents the value for undefined length
const UndefinedLength = 0xFFFFFFFF

// parses and reads a sequence
func (decoder *Decoder) readSequence(reader CounterReader, length uint32, explicitVR bool, byteOrder binary.ByteOrder) (*Sequence, error) {

	// if undefined length, read the sequence  using the provided reader knowing that there will be a deer item
	if length == UndefinedLength {
		return decoder.readSequenceItems(reader, explicitVR, byteOrder)
	}

	// otherwise, read the sequence using a limited reader for the length of sequence
	return decoder.readSequenceItems(newLimitedCountingReader(reader, int64(length)), explicitVR, byteOrder)
}

// reads the items of a sequence
func (decoder *Decoder) readSequenceItems(reader CounterReader, explicitVR bool, byteOrder binary.ByteOrder) (*Sequence, error) {

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

// reads a single item of a sequence
func (decoder *Decoder) readSequenceItem(reader CounterReader, explicitVR bool, byteOrder binary.ByteOrder) (*Object, error) {

	tag, err := decoder.readTag(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	length, err := decoder.readLong(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	// sequence deation item
	if tag == SequenceDelimitationItemTag {
		return nil, io.EOF
	}

	// item tag
	if tag != ItemTag {
		return nil, fmt.Errorf("expecting item tag at beginning of sequence item, found %s instead", tagToString(tag))
	}

	if length == UndefinedLength {
		object, err := decoder.readObject(reader, explicitVR, byteOrder)
		if err != nil {
			return nil, err
		}
		return object, nil
	}

	object, err := decoder.readObject(newLimitedCountingReader(reader, int64(length)), explicitVR, byteOrder)
	if err != nil {
		return nil, err
	}
	return object, nil

}

// parses and reads pixel data, in native or encapsulated formats
func (decoder *Decoder) readPixelData(reader CounterReader, length uint32, byteOrder binary.ByteOrder) (interface{}, error) {

	// if undefined length, read the pixel data  using the provided reader knowing that there will be a deer item
	if length == UndefinedLength {
		return decoder.readEncapsulatedPixelData(reader, byteOrder)
	}

	return decoder.readNativePixelData(reader, length)
}

// read native pixel data
func (decoder *Decoder) readNativePixelData(reader CounterReader, length uint32) ([]byte, error) {

	bytes, err := decoder.readBytes(reader, length)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// read encapsulated pixel data
func (decoder *Decoder) readEncapsulatedPixelData(reader CounterReader, byteOrder binary.ByteOrder) (*Encapsulated, error) {

	// create an encapsualted pixel data object
	encapsulated := newEncapsulated()

	// read the fragments
	for {

		fragment, err := decoder.readFragment(reader, byteOrder)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// add the object to the sequence
		encapsulated.add(fragment)
	}

	return encapsulated, nil
}

// read a fragment
func (decoder *Decoder) readFragment(reader CounterReader, byteOrder binary.ByteOrder) (*Fragment, error) {

	tag, err := decoder.readTag(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	length, err := decoder.readLong(reader, byteOrder)
	if err != nil {
		return nil, err
	}

	// fragment deation item
	if tag == SequenceDelimitationItemTag {
		return nil, io.EOF
	}

	// item tag
	if tag != ItemTag {
		return nil, fmt.Errorf("expecting item tag at beginning of fragment, found %s instead", tagToString(tag))
	}

	// get the offset from the underlying reader before we read the pixel data
	offset := uint32(reader.BytesRead())

	// read the pixel data
	// TODO: skip the pixel data and create a bulk data reference
	bytes, err := decoder.readBytes(reader, length)
	if err != nil {
		return nil, err
	}

	return &Fragment{bytes, offset}, nil
}
