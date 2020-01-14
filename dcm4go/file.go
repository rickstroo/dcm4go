package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
)

// ReadFile reads a DICOM object from a reader of a Part 10 source
func ReadFile(reader io.Reader, bulkDataThreshold uint32) (*Object, error) {

	// create a counting reader
	countingReader := newCountingReader(reader)

	// skip the preamble
	if err := skipBytes(countingReader, 128); err != nil {
		return nil, err
	}

	// read the prefix
	prefix, err := readText(countingReader, 4)
	if err != nil {
		return nil, err
	}

	// check the prefix
	if prefix != "DICM" {
		return nil, fmt.Errorf("found '%s': '%w'", prefix, ErrIllegalPrefix)
	}

	// create a decoder
	decoder := newDecoder(bulkDataThreshold)

	// read the group 2 length attribute
	groupTwoLength, err := decoder.readAttribute(countingReader, true, binary.LittleEndian)
	if err != nil {
		return nil, err
	}

	// check that it is the attribute that we are expecting
	if groupTwoLength.tag != FileMetaInformationGroupLengthTag {
		return nil, ErrUnexpectedAttribute
	}

	// calculate the length of group two
	groupTwoLengthValue, err := groupTwoLength.asLong(0)
	if err != nil {
		return nil, err
	}

	// create a limit reader for the remainder of the group two attributes
	limitCountingReader := newLimitedCountingReader(countingReader, int64(groupTwoLengthValue))

	// read the remainder of the group two attribute
	groupTwo, err := decoder.readObject(limitCountingReader, true, binary.LittleEndian)
	if err != nil {
		return nil, err
	}

	// need to find the transfer syntax uid
	transferSyntaxUID, err := groupTwo.asString(TransferSyntaxUIDTag, 0)
	if err != nil {
		return nil, err
	}

	// figure out the vr and endian of the remainder of the object
	transferSyntax, err := findTransferSyntax(transferSyntaxUID)
	if err != nil {
		return nil, err
	}

	// read the remainder of the attributes from the file using the provided transfer syntax
	otherGroups, err := decoder.readObject(countingReader, transferSyntax.explicitVR, transferSyntax.byteOrder)
	if err != nil {
		return nil, err
	}

	// concatenate the groups
	groupTwo.addAll(otherGroups)

	// return the groups
	return groupTwo, nil
}
