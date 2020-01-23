package dcm4go

import (
	"testing"
)

func newAttribute(tag uint32) *Attribute {
	return &Attribute{tag, "", 0, nil}
}

func TestAddAttribute(t *testing.T) {
	object := newObject()
	object.add(newAttribute(0x00080001))
	if _, err := object.find(0x00080001); err != nil {
		t.Error(err)
	}
}

func TestNotFindAttribute(t *testing.T) {
	object := newObject()
	_, err := object.find(0x00080001)
	if err := testErrEquals(err, ErrAttributeNotFound); err != nil {
		t.Error(err)
	}
}

func TestToString(t *testing.T) {
	object := newObject()
	object.add(newAttribute(0x00080001))
	if err := testStringEquals("[{tag:0x00080001,vr:,len:0}]", object.String()); err != nil {
		t.Error(err)
	}
}
