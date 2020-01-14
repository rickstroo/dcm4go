package dcm4go

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Attribute contains all the properties of a DICOM attribute
type Attribute struct {
	tag    uint32
	vr     string
	length uint32
	offset uint32
	value  interface{}
}

// String returns attribute as a string
func (attribute *Attribute) String() string {
	return attributeToString(attribute, "")
}

// MarshalJSON returns attribute as JSON
func (attribute *Attribute) MarshalJSON() ([]byte, error) {
	value, err := prepare(attribute)
	if err != nil {
		return nil, err
	}
	return json.Marshal(&struct {
		Tag    string      `json:"tag"`
		VR     string      `json:"vr"`
		Length uint32      `json:"len"`
		Offset uint32      `json:"off"`
		Value  interface{} `json:"value"`
	}{
		Tag:    fmt.Sprintf("%08X", attribute.tag),
		VR:     attribute.vr,
		Length: attribute.length,
		Offset: attribute.offset,
		Value:  value,
	})
}

// prepares the attribute for JSON marshalling as required
func prepare(attribute *Attribute) (interface{}, error) {
	switch attribute.vr {
	case "DS":
		// the standard says to format DS values as floats, not strings
		return convertStringsToFloats(attribute.value)
	case "IS":
		// the standards says to format IS values as ints, not strings
		return convertStringsToInts(attribute.value)
	}
	return attribute.value, nil
}

// converts a slice of strings to floats
func convertStringsToFloats(value interface{}) ([]float64, error) {
	strings, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("value not of type []string")
	}
	floats := make([]float64, len(strings))
	for i, string := range strings {
		float, err := strconv.ParseFloat(string, 64)
		if err != nil {
			return nil, err
		}
		floats[i] = float
	}
	return floats, nil
}

// converts a slice of strings to ints
func convertStringsToInts(value interface{}) ([]int64, error) {
	strings, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("value not of type []string")
	}
	ints := make([]int64, len(strings))
	for i, string := range strings {
		int, err := strconv.ParseInt(string, 10, 64)
		if err != nil {
			return nil, err
		}
		ints[i] = int
	}
	return ints, nil
}

// simple check for out of bounds
func checkIndex(index int, length int) error {
	if index < 0 || index >= length {
		return ErrIndexOutOfBounds
	}
	return nil
}

// AsLong returns attribute value as a long
func (attribute *Attribute) asLong(index int) (uint32, error) {
	longs, ok := attribute.value.([]uint32)
	if !ok {
		return 0, ErrWrongType
	}
	if err := checkIndex(index, len(longs)); err != nil {
		return 0, err
	}
	return longs[index], nil
}

// AsString returns attribute value as a string
func (attribute *Attribute) asString(index int) (string, error) {
	strings, ok := attribute.value.([]string)
	if !ok {
		return "", ErrWrongType
	}
	if err := checkIndex(index, len(strings)); err != nil {
		return "", err
	}
	return strings[index], nil
}
