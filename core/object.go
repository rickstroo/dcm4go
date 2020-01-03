package core

import (
	"container/list"
)

// Object cotains a list of DICOM attributes
type Object struct {
	attributes *list.List
}

// NewObject creates and initializes a new object
func NewObject() *Object {
	return &Object{list.New()}
}
