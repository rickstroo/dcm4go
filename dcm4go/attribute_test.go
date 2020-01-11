package dcm4go

import (
	"errors"
	"fmt"
	"testing"
)

func testLongEquals(a, b uint32) error {
	if a != b {
		return fmt.Errorf("expected 0x%04X, was 0x%04X", a, b)
	}
	return nil
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

func testStringEquals(a, b string) error {
	if a != b {
		return fmt.Errorf("expected '%s', was '%s'", a, b)
	}
	return nil
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
