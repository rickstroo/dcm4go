package dcm4go

import (
	"fmt"
	"io"
)

// ReadFile reads a DICOM object from a reader of a Part 10 source
func ReadFile(reader io.Reader, bulkDataThreshold uint32) (*Object, error) {

	// create a counting reader
	countingReader := newCountingReader(reader)

	// read the group two attributes
	groupTwo, err := ReadGroupTwo(countingReader, bulkDataThreshold)
	if err != nil {
		return nil, err
	}

	// create a decoder
	decoder := newDecoder(bulkDataThreshold)

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
	otherGroups, err := decoder.readObject(countingReader, transferSyntax)
	if err != nil {
		return nil, err
	}

	// concatenate the group two object and othe other groups object
	object := newObject()
	object.addAll(groupTwo)
	object.addAll(otherGroups)

	// return the groups
	return object, nil
}

// ReadGroupTwo reads the group two attributes of a DICOM object from a reader of a Part 10 source
func ReadGroupTwo(reader io.Reader, bulkDataThreshold uint32) (*Object, error) {

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
	groupTwoLength, err := decoder.readAttribute(countingReader, ExplicitVRLittleEndianTS)
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
	groupTwo, err := decoder.readObject(limitCountingReader, ExplicitVRLittleEndianTS)
	if err != nil {
		return nil, err
	}

	// create a new object with the group two length attribute
	// followed by the rest of the group two attributes
	object := newObject()
	object.add(groupTwoLength)
	object.addAll(groupTwo)

	// return the new object as the complete group two
	return object, nil
}

// WriteFile writes a DICOM file
func WriteFile(writer io.Writer, fmi *Object, reader io.Reader) error {

	// write the leading 128 zeroes
	var zeros [128]byte
	if err := writeBytes(writer, zeros[:]); err != nil {
		return err
	}

	// write the DICM prefix
	if err := writeString(writer, "DICM"); err != nil {
		return err
	}

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the fmi
	if err := encoder.writeObjectWithGroupLength(writer, 0x0002, fmi, ImplicitVRLittleEndianTS); err != nil {
		return err
	}

	// copy the data
	_, err := io.Copy(writer, reader)
	if err != nil {
		return err
	}

	// all is well
	return nil
}

// CreateFileMetaInfo creates the file meta information for a Part 10 file
func CreateFileMetaInfo(assoc *Assoc, pcID byte, command *Object) (*Object, error) {

	// get the required information from the command
	sopClassUID, err := command.asString(AffectedSOPClassUIDTag, 0)
	if err != nil {
		return nil, err
	}
	sopInstanceUID, err := command.asString(AffectedSOPInstanceUIDTag, 0)
	if err != nil {
		return nil, err
	}

	// find the transfer syntax used to receive the object
	transferSyntax, err := assoc.findAcceptedTransferSyntaxByPCID(pcID)
	if err != nil {
		return nil, err
	}

	// create the fmi
	fmi := newObject()
	fmi.addShort(FileMetaInformationVersionTag, "US", 0x0100)
	fmi.addUID(MediaStorageSOPClassUIDTag, sopClassUID)
	fmi.addUID(MediaStorageSOPInstanceUIDTag, sopInstanceUID)
	fmi.addUID(TransferSyntaxUIDTag, transferSyntax.uid)
	fmi.addUID(ImplementationClassUIDTag, ImplementationClassUID)
	fmi.addText(ImplementationVersionNameTag, "SH", "dcm4go")
	fmi.addText(SourceApplicationEntityTitleTag, "AE", assoc.ae.aeTitle)
	fmi.addText(SendingApplicationEntityTitleTag, "AE", assoc.CallingAETitle())
	fmi.addText(ReceivingApplicationEntityTitleTag, "AE", assoc.CalledAETitle())

	// return the file meta information
	return fmi, nil
}
