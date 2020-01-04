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
func NewObject() *Object {
	return &Object{list.New()}
}

// Add adds an attribute to an object
func (object *Object) Add(attribute *Attribute) {
	object.attributes.PushBack(attribute)
}

// String returns a string representation of an object
func (object *Object) String() string {
	s := ""
	for item := object.attributes.Front(); item != nil; item = item.Next() {
		s += item.Value.(*Attribute).String()
	}
	return s
}

// Find looks for an attribute in an object
func (object *Object) Find(group uint16, element uint16) (*Attribute, error) {
	for item := object.attributes.Front(); item != nil; item = item.Next() {
		attribute := item.Value.(*Attribute)
		if attribute.group == group && attribute.element == element {
			return attribute, nil
		}
	}
	return nil, fmt.Errorf("unable to find attribute with group 0x%04x and element 0x%04x", group, element)
}

// AsLong returns attribute value as a long
func (object *Object) AsLong(group uint16, element uint16, index int) (uint32, error) {
	attribute, err := object.Find(group, element)
	if err != nil {
		return 0, err
	}
	return attribute.AsLong(index)
}

// AsString returns attribute value as a string
func (object *Object) AsString(group uint16, element uint16, index int) (string, error) {
	attribute, err := object.Find(group, element)
	if err != nil {
		return "", err
	}
	return attribute.AsString(index)
}
