package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// ReadFile reads a DICOM object from a file
func ReadFile(path string, bulkDataThreshold uint32) (*Object, *Object, error) {

	// open the file, which returns a reader, defer a close
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	// make sure we close the file upon exit
	defer file.Close()

	// read the file
	return Read(file, bulkDataThreshold)
}

// Read reads a DICOM object from a reader
func Read(reader io.Reader, bulkDataThreshold uint32) (*Object, *Object, error) {

	// create a counting reader
	countingReader := newCountingReader(reader)

	// create a decoder
	decoder := newDecoder(bulkDataThreshold)

	// read the preamble
	var preamble [128]byte
	if err := decoder.readFully(countingReader, preamble[:]); err != nil {
		return nil, nil, err
	}

	// read the prefix
	var prefix [4]byte
	if err := decoder.readFully(countingReader, prefix[:]); err != nil {
		return nil, nil, err
	}

	// check the prefix
	if string(prefix[:]) != "DICM" {
		return nil, nil, fmt.Errorf("unrecognized prefix, '%s'", prefix)
	}

	// read the group 2 length attribute
	groupTwoLength, err := decoder.readAttribute(countingReader, true, binary.LittleEndian)
	if err != nil {
		return nil, nil, err
	}

	// check that it is the attribute that we are expecting
	if groupTwoLength.tag != FileMetaInformationGroupLengthTag {
		return nil, nil, fmt.Errorf("unexpected first attribute in file, was expecting %s, found %s", tagToString(FileMetaInformationGroupLengthTag), tagToString(groupTwoLength.tag))
	}

	// calculate the length of group two
	groupTwoLengthValue, err := groupTwoLength.asLong(0)
	if err != nil {
		return nil, nil, err
	}

	// create a limit reader for the remainder of the group two attributes
	limitCountingReader := newLimitedCountingReader(countingReader, int64(groupTwoLengthValue))

	// read the remainder of the group two attribute
	groupTwo, err := decoder.readObject(limitCountingReader, true, binary.LittleEndian)
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
	otherGroups, err := decoder.readObject(countingReader, transferSyntax.explicitVR, transferSyntax.byteOrder)
	if err != nil {
		return nil, nil, err
	}

	return groupTwo, otherGroups, nil
}
