package dcm4go

import "fmt"

// Attribute contains all the properties of a DICOM attribute
type Attribute struct {
	group   uint16
	element uint16
	vr      string
	length  uint32
	offset  uint32
	value   interface{}
}

// String returns attribute as a string
func (attribute *Attribute) String() string {
	return attributeToString(attribute, "")
}

// AsLong returns attribute value as a long
func (attribute *Attribute) asLong(index int) (uint32, error) {
	longs, ok := attribute.value.([]uint32)
	if !ok {
		return 0, fmt.Errorf("attribute was not of types longs")
	}
	if index < 0 || index >= len(longs) {
		return 0, fmt.Errorf("index %d out of bounds, range is 0..%d", index, len(longs)-1)
	}
	return longs[index], nil
}

// AsString returns attribute value as a string
func (attribute *Attribute) asString(index int) (string, error) {
	strings, ok := attribute.value.([]string)
	if !ok {
		return "", fmt.Errorf("attribute was not of types strings")
	}
	if index < 0 || index >= len(strings) {
		return "", fmt.Errorf("index %d out of bounds, range is 0..%d", index, len(strings)-1)
	}
	return strings[index], nil
}
