package core

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

func attributeToString(attribute *Attribute, prefix string) string {
	s := fmt.Sprintf("%stag=(%04X,%04X) vr=%s off=%d len=%d", prefix, attribute.group, attribute.element, attribute.vr, attribute.offset, attribute.length)
	switch attribute.vr {
	case "AE", "AS", "CS", "DA", "DT", "LO", "SH", "TM", "UC", "UI", "UR", "LT", "ST", "UT", "PN":
		s += fmt.Sprintf(" value=%q\n", attribute.value)
	case "DS", "IS", "FD", "OD", "FL", "OF", "SS", "US", "SL", "UL", "SV", "UV":
		s += fmt.Sprintf(" value=%v\n", attribute.value)
	case "SQ":
		s += "\n" + sequenceToString(attribute.value.(*Sequence), prefix)
	case "AT", "OB", "OL", "OV", "OW", "UN":
		s += "\n"
	default:
	}
	return s
}

func sequenceToString(sequence *Sequence, prefix string) string {
	s := ""
	for i, item := 1, sequence.objects.Front(); item != nil; i, item = i+1, item.Next() {
		s += itemToString(item.Value.(*Object), fmt.Sprintf("%s%d>", prefix, i))
	}
	return s
}

func itemToString(object *Object, prefix string) string {
	s := ""
	for item := object.attributes.Front(); item != nil; item = item.Next() {
		s += attributeToString(item.Value.(*Attribute), prefix)
	}
	return s
}

// AsLong returns attribute value as a long
func (attribute *Attribute) AsLong(index int) (uint32, error) {
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
func (attribute *Attribute) AsString(index int) (string, error) {
	strings, ok := attribute.value.([]string)
	if !ok {
		return "", fmt.Errorf("attribute was not of types strings")
	}
	if index < 0 || index >= len(strings) {
		return "", fmt.Errorf("index %d out of bounds, range is 0..%d", index, len(strings)-1)
	}
	return strings[index], nil
}
