package core

import (
	"encoding/binary"
	"fmt"
)

// TransferSyntax describes the how an object is encoded
type TransferSyntax struct {
	explicitVR bool
	byteOrder  binary.ByteOrder
	name       string
}

// ExplicitVR returns a boolean indicating whether or not this transfer syntax use an explicit VR
func (transferSyntax *TransferSyntax) ExplicitVR() bool {
	return transferSyntax.explicitVR
}

// ByteOrder returns the byte order for this transfer syntax
func (transferSyntax *TransferSyntax) ByteOrder() binary.ByteOrder {
	return transferSyntax.byteOrder
}

// Name returns the name of this transfer syntax
func (transferSyntax *TransferSyntax) Name() string {
	return transferSyntax.name
}

// FindTransferSyntax figures out the explicit vr and byte ByteOrder
func FindTransferSyntax(transferSyntaxUID string) (*TransferSyntax, error) {
	switch transferSyntaxUID {
	case "1.2.840.10008.1.2":
		return &TransferSyntax{false, binary.LittleEndian, "Implicit VR Little Endian"}, nil
	case "1.2.840.10008.1.2.1":
		return &TransferSyntax{true, binary.LittleEndian, "Explicit VR Little Endian"}, nil
	case "1.2.840.10008.1.2.1.99":
		return &TransferSyntax{true, binary.LittleEndian, "Deflated Explicit VR Little Endian"}, nil
	case "1.2.840.10008.1.2.2":
		return &TransferSyntax{true, binary.BigEndian, "Explicit VR Big Endian"}, nil
	case "1.2.840.10008.1.2.4.91":
		return &TransferSyntax{true, binary.LittleEndian, "JPEG 2000 Image Compression"}, nil
	}
	return nil, fmt.Errorf("unrecognized transfer syntax uid, '%s'", transferSyntaxUID)
}
