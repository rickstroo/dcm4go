package dcm4go

import (
	"bytes"
	"fmt"
	"testing"
)

func initEncoderTest() (*Encoder, error) {
	return newEncoder(), nil
}

func TestNewEncoder(t *testing.T) {
	if _, err := initEncoderTest(); err != nil {
		t.Error(err)
	}
}

func TestEncodeLongAttributeImplicitVRLittleEndianTS(t *testing.T) {
	encoder, err := initEncoderTest()
	if err != nil {
		t.Error(err)
	}
	attribute := &attribute{FileMetaInformationGroupLengthTag, "UL", []uint32{0x04}}
	buf := new(bytes.Buffer)
	if err := encoder.writeAttribute(buf, attribute, ImplicitVRLittleEndianTS); err != nil {
		t.Error(err)
	}
	if err := equal(buf.Bytes(), []byte{2, 0, 0, 0, 4, 0, 0, 0, 4, 0, 0, 0}); err != nil {
		t.Error(err)
	}
}

func TestEncodeLongAttributeExplicitVRLittleEndianTS(t *testing.T) {
	encoder, err := initEncoderTest()
	if err != nil {
		t.Error(err)
	}
	attribute := &attribute{FileMetaInformationGroupLengthTag, "UL", []uint32{0x04}}
	buf := new(bytes.Buffer)
	if err := encoder.writeAttribute(buf, attribute, ExplicitVRLittleEndianTS); err != nil {
		t.Error(err)
	}
	if err := equal(buf.Bytes(), []byte{2, 0, 0, 0, 85, 76, 4, 0, 4, 0, 0, 0}); err != nil {
		t.Error(err)
	}
}

func TestEncodeLongAttributeExplicitVRBigEndianTS(t *testing.T) {
	encoder, err := initEncoderTest()
	if err != nil {
		t.Error(err)
	}
	attribute := &attribute{FileMetaInformationGroupLengthTag, "UL", []uint32{0x04}}
	buf := new(bytes.Buffer)
	if err := encoder.writeAttribute(buf, attribute, ExplicitVRBigEndianTS); err != nil {
		t.Error(err)
	}
	if err := equal(buf.Bytes(), []byte{0, 2, 0, 0, 85, 76, 0, 4, 0, 0, 0, 4}); err != nil {
		t.Error(err)
	}
}

func equal(a, b []byte) error {
	if !bytes.Equal(a, b) {
		return fmt.Errorf("while comparing byte streams, found %v while expecting %v", a, b)
	}
	return nil
}
