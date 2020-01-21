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

// findTransferSyntax figures out the explicit vr and byte ByteOrder
func findTransferSyntax(transferSyntaxUID string) (*TransferSyntax, error) {
	transferSyntax, ok := tses[transferSyntaxUID]
	if !ok {
		return nil, fmt.Errorf("transfer syntax '%s': %w", transferSyntaxUID, ErrUnrecognizedTransferSyntax)
	}
	return transferSyntax, nil
}
