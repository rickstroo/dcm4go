package dcm4go

import (
	"encoding/binary"
	"fmt"
)

// TransferSyntax describes the how an object is encoded
type TransferSyntax struct {
	uid        string
	explicitVR bool
	byteOrder  binary.ByteOrder
	name       string
}

// ImplicitVRLittleEndianTS returns the implicit VR little endian transfer syntax
func ImplicitVRLittleEndianTS() *TransferSyntax {
	return &TransferSyntax{ImplicitVRLittleEndianUID, false, binary.LittleEndian, "Implicit VR Little Endian"}
}

// FindTransferSyntax figures out the explicit vr and byte ByteOrder
func findTransferSyntax(transferSyntaxUID string) (*TransferSyntax, error) {
	switch transferSyntaxUID {
	case "1.2.840.10008.1.2":
		return ImplicitVRLittleEndianTS(), nil
	case "1.2.840.10008.1.2.1":
		return &TransferSyntax{ExplicitVRLittleEndianUID, true, binary.LittleEndian, "Explicit VR Little Endian"}, nil
	case "1.2.840.10008.1.2.1.99":
		return &TransferSyntax{DeflatedExplicitVRLittleEndianUID, true, binary.LittleEndian, "Deflated Explicit VR Little Endian"}, nil
	case "1.2.840.10008.1.2.2":
		return &TransferSyntax{ExplicitVRBigEndianUID, true, binary.BigEndian, "Explicit VR Big Endian"}, nil
	case "1.2.840.10008.1.2.4.90":
		return &TransferSyntax{JPEG2000ImageCompressionLosslessOnlyUID, true, binary.LittleEndian, "JPEG 2000 Image Compression (Lossless Only)"}, nil
	case "1.2.840.10008.1.2.4.91":
		return &TransferSyntax{JPEG2000ImageCompressUID, true, binary.LittleEndian, "JPEG 2000 Image Compression"}, nil
	}
	return nil, fmt.Errorf("transfer syntax '%s': %w", transferSyntaxUID, ErrUnrecognizedTransferSyntax)
}
