package core

// Attribute is a single element of a DICOM object
type Attribute struct {
	group   uint16
	element uint16
	vr      string
	length  uint32
	value   interface{}
}
