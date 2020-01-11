package dcm4go

import (
	"errors"
	"testing"
)

func TestAttributeToString(t *testing.T) {
	attribute := &Attribute{FileMetaInformationGroupLengthTag, "UL", 4, 0, []uint32{0x04}}
	if err := testStringEquals(attribute.String(), "tag=(0002,0000) vr=UL off=0 len=4 value=[4]\n"); err != nil {
		t.Error(err)
	}
}

func TestAttributeAsLong(t *testing.T) {
	attribute := &Attribute{FileMetaInformationGroupLengthTag, "UL", 4, 0, []uint32{0x04}}
	value, err := attribute.asLong(0)
	if err != nil {
		t.Error(err)
	}
	if err := testLongEquals(value, 0x04); err != nil {
		t.Error(err)
	}
}

func TestAttributeAsLongIndexOutOfBounds(t *testing.T) {
	attribute := &Attribute{FileMetaInformationGroupLengthTag, "UL", 4, 0, []uint32{0x04}}
	_, err := attribute.asLong(1)
	if !errors.Is(err, ErrIndexOutOfBounds) {
		t.Error(err)
	}
}

func TestAttributeAsLongWrongType(t *testing.T) {
	attribute := &Attribute{FileMetaInformationGroupLengthTag, "UL", 4, 0, []uint16{0x04}}
	_, err := attribute.asLong(0)
	if !errors.Is(err, ErrWrongType) {
		t.Error(err)
	}
}

func TestAttributeAsString(t *testing.T) {
	attribute := &Attribute{ModalityTag, "CS", 4, 0, []string{"CT"}}
	value, err := attribute.asString(0)
	if err != nil {
		t.Error(err)
	}
	if err := testStringEquals(value, "CT"); err != nil {
		t.Error(err)
	}
}

func TestAttributeAsStringIndexOutOfBounds(t *testing.T) {
	attribute := &Attribute{ModalityTag, "CS", 4, 0, []string{"CT"}}
	_, err := attribute.asString(1)
	if err := testErrEquals(err, ErrIndexOutOfBounds); err != nil {
		t.Error(err)
	}
}

func TestAttributeAsStringWrongType(t *testing.T) {
	attribute := &Attribute{ModalityTag, "CS", 4, 0, []byte{0x01}}
	_, err := attribute.asString(0)
	if err := testErrEquals(err, ErrWrongType); err != nil {
		t.Error(err)
	}
}