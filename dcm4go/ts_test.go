package dcm4go

import (
	"fmt"
	"testing"
)

func TestFindTransferSyntax(t *testing.T) {
	transferSyntax, err := findTransferSyntax(ImplicitVRLittleEndianUID)
	if err != nil {
		t.Error(err)
	}
	if transferSyntax != ImplicitVRLittleEndianTS {
		t.Error(fmt.Errorf("transfer syntax was found but was not implicit VR little endian"))
	}
}

func TestNotFindTransferSyntax(t *testing.T) {
	_, err := findTransferSyntax("1.2.3.4")
	if err := testErrEquals(err, ErrUnrecognizedTransferSyntax); err != nil {
		t.Error(err)
	}
}
