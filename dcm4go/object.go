package dcm4go

// Object cotains all attributes of a DICOM object
type Object struct {
	attributes []*Attribute
}

// // MarshalJSON returns sequence as JSON
// func (object *Object) MarshalJSON() ([]byte, error) {
// 	// return json.Marshal(&struct {
// 	// 	Attributes []*Attribute `json:"attributes"`
// 	// }{
// 	// 	Attributes: object.attributes,
// 	// })
// 	type a struct {
// 		VR    string
// 		Value interface{}
// 	}
// 	m := make(map[string]*Attribute)
// 	for _, attribute := range object.attributes {
// 		// standards not to output group length attributes
// 		if toElement(attribute.tag) != 0x0000 {
// 			m[fmt.Sprintf("%08X", attribute.tag)] = attribute
// 		}
// 	}
// 	return json.Marshal(m)
// }

// NewObject creates and initializes a new object
func newObject() *Object {
	return &Object{make([]*Attribute, 0, 100)}
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
func (object *Object) add(attribute *Attribute) {
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
func (object *Object) find(tag uint32) (*Attribute, error) {
	for _, attribute := range object.attributes {
		if attribute.tag == tag {
			return attribute, nil
		}
	}
	return nil, ErrAttributeNotFound
}

// AsLong returns attribute value as a long
func (object *Object) asLong(tag uint32, index int) (uint32, error) {
	attribute, err := object.find(tag)
	if err != nil {
		return 0, err
	}
	return attribute.asLong(index)
}

// AsString returns attribute value as a string
func (object *Object) asString(tag uint32, index int) (string, error) {
	attribute, err := object.find(tag)
	if err != nil {
		return "", err
	}
	return attribute.asString(index)
}
