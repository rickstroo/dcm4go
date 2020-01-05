package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// ReadFile reads a DICOM object from a file
func ReadFile(path string) (*Object, *Object, error) {

	// open the file, which returns a reader, defer a close
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	// make sure we close the file upon exit
	defer file.Close()

	// create a decoder
	decoder := newDecoder()

	// read the preamble
	var preamble [128]byte
	err = decoder.readFully(file, preamble[:])
	if err != nil {
		return nil, nil, err
	}

	// read the prefix
	var prefix [4]byte
	err = decoder.readFully(file, prefix[:])
	if err != nil {
		return nil, nil, err
	}

	// check the prefix
	if string(prefix[:]) != "DICM" {
		return nil, nil, fmt.Errorf("unrecognized prefix, '%s'", prefix)
	}

	// read the group 2 length attribute
	groupTwoLength, err := decoder.readAttribute(file, true, binary.LittleEndian)
	if err != nil {
		return nil, nil, err
	}

	// check that it is the attribute that we are expecting
	if groupTwoLength.tag != FileMetaInformationGroupLengthTag {
		return nil, nil, fmt.Errorf("unexpected first attribute in file, was expecting %s, found %s", toString(FileMetaInformationGroupLengthTag), toString(groupTwoLength.tag))
	}

	// calculate the length of group two
	groupTwoLengthValue, err := groupTwoLength.asLong(0)
	if err != nil {
		return nil, nil, err
	}

	// create a limit reader for the remainder of the group two attributes
	limitReader := io.LimitReader(file, int64(groupTwoLengthValue))

	// read the remainder of the group two attribute
	groupTwo, err := decoder.readObject(limitReader, true, binary.LittleEndian)
	if err != nil {
		return nil, nil, err
	}

	// need to find the transfer syntax uid
	transferSyntaxUID, err := groupTwo.asString(TransferSyntaxUIDTag, 0)
	if err != nil {
		return nil, nil, err
	}

	// figure out the vr and endian of the remainder of the object
	transferSyntax, err := findTransferSyntax(transferSyntaxUID)
	if err != nil {
		return nil, nil, err
	}

	// read the remainder of the attributes from the file
	otherGroups, err := decoder.readObject(file, transferSyntax.explicitVR, transferSyntax.byteOrder)
	if err != nil {
		return nil, nil, err
	}

	return groupTwo, otherGroups, nil
}
