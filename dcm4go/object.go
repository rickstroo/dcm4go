package dcm4go

// Object cotains all attributes of a DICOM object
type Object struct {
	attributes []*attribute
}

// newObject creates and initializes a new object
func newObject() *Object {
	return &Object{make([]*attribute, 0, 100)}
}

// Attributes returns a reference to the attributes
func (object *Object) Attributes() []*attribute {
	return object.attributes
}

// Add adds an attribute to an object
// TODO: Right now, this works because we are
// parsing files and we know that we will be
// adding attributes in order.  We might
// even get away with this simple implementation
// if we are careful in how we construct our
// objects.  However, for general purposes, we
// will want to add support for adding attributes
// in an unordered manner.  We could do that using
// a list, or we could use a map that we would need
// to sort before using.  I prefer the list, as I think
// that will be more efficient.
func (object *Object) add(attribute *attribute) {
	object.attributes = append(object.attributes, attribute)
}

// Adds all attributes of an object
func (object *Object) addAll(other *Object) {
	object.attributes = append(object.attributes, other.attributes...)
}

// String returns a string representation of an object
func (object *Object) String() string {
	return objectToString(object, "")
}

// Find looks for an attribute in an object
func (object *Object) find(tag uint32) (*attribute, error) {
	for _, attribute := range object.attributes {
		if attribute.tag == tag {
			return attribute, nil
		}
	}
	return nil, ErrAttributeNotFound
}

// asShort returns attribute value as a short
func (object *Object) asShort(tag uint32, index int) (uint16, error) {
	attribute, err := object.find(tag)
	if err != nil {
		return 0, err
	}
	return attribute.asShort(index)
}

// AsShort returns attribute value as a short
func (object *Object) AsShort(tag uint32, index int) (uint16, error) {
	return object.asShort(tag, index)
}

// asLong returns attribute value as a long
func (object *Object) asLong(tag uint32, index int) (uint32, error) {
	attribute, err := object.find(tag)
	if err != nil {
		return 0, err
	}
	return attribute.asLong(index)
}

// asString returns attribute value as a string
func (object *Object) asString(tag uint32, index int) (string, error) {
	attribute, err := object.find(tag)
	if err != nil {
		return "", err
	}
	return attribute.asString(index)
}

// AsString returns attribute value as a string
func (object *Object) AsString(tag uint32, index int) (string, error) {
	return object.asString(tag, index)
}

// addUID adds a UID attribute
func (object *Object) addUID(tag uint32, uid string) {
	attribute := &attribute{tag, "UI", []string{uid}}
	object.add(attribute)
}

// addText adds a text attribute
func (object *Object) addText(tag uint32, vr string, text string) {
	attribute := &attribute{tag, "UI", []string{text}}
	object.add(attribute)
}

// addShort adds a short attribute
func (object *Object) addShort(tag uint32, vr string, value uint16) {
	attribute := &attribute{tag, vr, []uint16{value}}
	object.add(attribute)
}

// addLong adds a long attribute
func (object *Object) addLong(tag uint32, vr string, value uint32) {
	attribute := &attribute{tag, vr, []uint32{value}}
	object.add(attribute)
}
