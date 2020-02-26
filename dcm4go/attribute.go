package dcm4go

// attribute contains all the properties of a DICOM attribute
type attribute struct {
	tag   uint32
	vr    string
	value interface{}
}

// String returns attribute as a string
func (attribute *attribute) String() string {
	return attributeToString(attribute, "")
}

// Text returns attribute as text
func (attribute *attribute) Text() string {
	return attributeToText(attribute, "")
}

// simple check for out of bounds
func checkIndex(index int, length int) error {
	if index < 0 || index >= length {
		return ErrIndexOutOfBounds
	}
	return nil
}

// asShort returns attribute value as a short
func (attribute *attribute) asShort(index int) (uint16, error) {
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
func (attribute *attribute) asLong(index int) (uint32, error) {
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
func (attribute *attribute) asString(index int) (string, error) {
	strings, ok := attribute.value.([]string)
	if !ok {
		return "", ErrWrongType
	}
	if err := checkIndex(index, len(strings)); err != nil {
		return "", err
	}
	return strings[index], nil
}

// Sequence contains an ordered list of objects
type Sequence struct {
	objects []*Object
}

func (sequence *Sequence) add(object *Object) {
	sequence.objects = append(sequence.objects, object)
}

func newSequence() *Sequence {
	return &Sequence{make([]*Object, 0, 10)}
}

// Fragment encapsulates pixel data
type Fragment struct {
	offset uint32
	length uint32
}

// Encapsulated contains an ordered list of fragments
type Encapsulated struct {
	fragments []*Fragment
}

func (encapsulated *Encapsulated) add(fragment *Fragment) {
	encapsulated.fragments = append(encapsulated.fragments, fragment)
}

func newEncapsulated() *Encapsulated {
	return &Encapsulated{make([]*Fragment, 0, 10)}
}
