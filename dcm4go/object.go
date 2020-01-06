package dcm4go

import (
	"container/list"
	"fmt"
)

// Object cotains all attributes of a DICOM object
type Object struct {
	attributes *list.List
}

// NewObject creates and initializes a new object
func newObject() *Object {
	return &Object{list.New()}
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
	object.attributes.PushBack(attribute)
}

// String returns a string representation of an object
func (object *Object) String() string {
	return objectToString(object, "")
}

// Find looks for an attribute in an object
func (object *Object) find(tag uint32) (*Attribute, error) {
	for item := object.attributes.Front(); item != nil; item = item.Next() {
		attribute := item.Value.(*Attribute)
		if attribute.tag == tag {
			return attribute, nil
		}
	}
	return nil, fmt.Errorf("unable to find attribute with tag %s", tagToString(tag))
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
