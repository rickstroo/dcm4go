package file

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/rickstroo/dcm4go/core"
)

// ReadFile reads a DICOM object from a file
func ReadFile(path string) (*core.Object, *core.Object, error) {

	// open the file, which returns a reader, defer a close
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	// make sure we close the file
	defer file.Close()

	// create a decoder, it counts bytes
	decoder := core.NewDecoder()

	// read the preamble
	_, _, err = decoder.Read(file, 128)
	if err != nil {
		return nil, nil, err
	}

	// read the prefix
	_, prefix, err := decoder.Read(file, 4)
	if err != nil {
		return nil, nil, err
	}

	// check the prefix
	if string(prefix) != "DICM" {
		return nil, nil, fmt.Errorf("unrecognized prefix, '%s'", prefix)
	}

	// read the group 2 length attribute
	groupTwoLength, err := decoder.ReadAttribute(file, true, binary.LittleEndian)
	if err != nil {
		return nil, nil, err
	}

	// calculate the length of group two
	groupTwoLengthValue, err := groupTwoLength.AsLong(0)
	if err != nil {
		return nil, nil, err
	}

	// create a limit reader for the remainder of the group two attributes
	limitReader := io.LimitReader(file, int64(groupTwoLengthValue))

	// read the remainder of the group two attribute
	groupTwo, err := decoder.ReadObject(limitReader, true, binary.LittleEndian)
	if err != nil {
		return nil, nil, err
	}

	// need to find the transfer syntax uid
	transferSyntaxUID, err := groupTwo.AsString(0x0002, 0x0010, 0)
	if err != nil {
		return nil, nil, err
	}

	// figure out the vr and endian of the remainder of the object
	transferSyntax, err := findTransferSyntax(transferSyntaxUID)
	if err != nil {
		return nil, nil, err
	}

	// read the remainder of the attributes from the file
	otherGroups, err := decoder.ReadObject(file, transferSyntax.explicitVR, transferSyntax.byteOrder)
	if err != nil {
		return nil, nil, err
	}

	return groupTwo, otherGroups, nil
}
