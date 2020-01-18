package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Encoder decodes a stream of bytes as a DICOM object.
type Encoder struct {
}

// newEncoder creates a new Encoder
func newEncoder() *Encoder {
	return &Encoder{}
}

// writeObject writes an object to a writer
func (encoder *Encoder) writeObject(writer io.Writer, object *Object, explicitVR bool, byteOrder binary.ByteOrder) error {
	for _, attribute := range object.attributes {
		if err := encoder.writeAttribute(writer, attribute, explicitVR, byteOrder); err != nil {
			return err
		}
	}
	return nil
}

// writeAttributes writes an object to a writer
func (encoder *Encoder) writeAttribute(writer io.Writer, attribute *Attribute, explicitVR bool, byteOrder binary.ByteOrder) error {
	return fmt.Errorf("Encoder.writeAttribute: not implemented")
}
