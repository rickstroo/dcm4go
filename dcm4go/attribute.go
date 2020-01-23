package dcm4go

// Attribute contains all the properties of a DICOM attribute
type Attribute struct {
	tag    uint32
	vr     string
	length uint32
	value  interface{}
}

// String returns attribute as a string
func (attribute *Attribute) String() string {
	return attributeToString(attribute, "")
}

// simple check for out of bounds
func checkIndex(index int, length int) error {
	if index < 0 || index >= length {
		return ErrIndexOutOfBounds
	}
	return nil
}

// asShort returns attribute value as a short
func (attribute *Attribute) asShort(index int) (uint16, error) {
	shorts, ok := attribute.value.([]uint16)
	if !ok {
		return 0, ErrWrongType
	}
	if err := checkIndex(index, len(shorts)); err != nil {
		return 0, err
	}
	return shorts[index], nil
}

// asLong returns attribute value as a long
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

// asString returns attribute value as a string
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
