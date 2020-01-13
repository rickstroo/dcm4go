package dcm4go

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

// FindTransferSyntax figures out the explicit vr and byte ByteOrder
func findTransferSyntax(transferSyntaxUID string) (*TransferSyntax, error) {
	switch transferSyntaxUID {
	case "1.2.840.10008.1.2":
		return &TransferSyntax{false, binary.LittleEndian, "Implicit VR Little Endian"}, nil
	case "1.2.840.10008.1.2.1":
		return &TransferSyntax{true, binary.LittleEndian, "Explicit VR Little Endian"}, nil
	case "1.2.840.10008.1.2.1.99":
		return &TransferSyntax{true, binary.LittleEndian, "Deflated Explicit VR Little Endian"}, nil
	case "1.2.840.10008.1.2.2":
		return &TransferSyntax{true, binary.BigEndian, "Explicit VR Big Endian"}, nil
	case "1.2.840.10008.1.2.4.90":
		return &TransferSyntax{true, binary.LittleEndian, "JPEG 2000 Image Compression (Lssless Only)"}, nil
	case "1.2.840.10008.1.2.4.91":
		return &TransferSyntax{true, binary.LittleEndian, "JPEG 2000 Image Compression"}, nil
	}
	return nil, fmt.Errorf("transfer syntax '%s': %w", transferSyntaxUID, ErrUnrecognizedTransferSyntax)
}
