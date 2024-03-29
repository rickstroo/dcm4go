// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

// This package is used to read and write pdus.
// While this implementation may not be the most efficient, I have chosen
// readability over efficiency.  In truth, most of the code here is used
// to support the reading and writing of PDUs that are used during
// association negotiation or association release, which is an insignificant
// cost compared to the sending and receiving of data.  As long as the data
// transfers are efficient, we should be alright.

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
)

// define the value pdu types
const (
	aAssociateRQPDU = 0x01
	aAssociateACPDU = 0x02
	aAssociateRJPDU = 0x03
	pDataTFPDU      = 0x04
	aReleaseRQPDU   = 0x05
	aReleaseRPPDU   = 0x06
	aAbortPDU       = 0x07
)

// pdu represents a DICOM protocol data unit (i.e. PDU)
type pdu struct {
	typ byte
	buf []byte
}

// String returns a string representation of a PDU
func (pdu *pdu) String() string {
	if pdu == nil {
		return "nil"
	}
	return fmt.Sprintf("{typ:%v,buf:%v}", pdu.typ, pdu.buf)
}

// readPDU reads a PDU from a reader
func readPDU(reader io.Reader) (*pdu, error) {

	// get the pdu type
	typ, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// skip a byte, as per the standard
	if err := skipByte(reader); err != nil {
		return nil, err
	}

	// read the length, pdu lengths are always big endian
	len, err := readLong(reader, binary.BigEndian)
	if err != nil {
		return nil, err
	}

	// read the contents of the pdu.
	// we may want to revisit this in the future
	// it could be more efficient not to read the bytes
	// at this point, and use a limited reader to read
	// the bytes when required.
	// on the other hand, if we read the contents of the pdu now, we could
	// use worker threads to read the pdus from the network and then pass
	// them up to higher levels of logic for processing.  this might be
	// a fun project to try at some point.
	// this decision is perhaps one of the most contentious and significant
	// in my design in think.
	// having all the bytes read at this point may also make for more efficent
	// conversion of bytes into integers, floats and strings.
	// that being said, the place that we only need to think about efficiency
	// is the handling of the data pdvs that are contained in pdus.  for the
	// most part, those data pdvs are going to be copied from the network
	// to disk, and then back again.
	buf, err := readBytes(reader, len)
	if err != nil {
		return nil, err
	}

	// construct a PDU
	pdu := &pdu{
		typ: typ,
		buf: buf,
	}

	// return the pdu
	return pdu, nil
}

// writePDU writes a PDU to a writer
func writePDU(writer io.Writer, pdu *pdu) error {

	// write the type
	if err := writeByte(writer, pdu.typ); err != nil {
		return err
	}

	// write a zero byte, as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}

	// write the length
	if err := writeLong(writer, uint32(len(pdu.buf)), binary.BigEndian); err != nil {
		return err
	}

	// write the bytes
	if err := writeBytes(writer, pdu.buf); err != nil {
		return err
	}

	// return success
	return nil
}

// writeTo writes a pdu
func (pdu *pdu) writeTo(writer io.Writer) error {
	return writePDU(writer, pdu)
}

// define the sources of an abort
const (
	sourceServiceUserInitiatedAbort     = 0x00
	sourceServiceProviderInitiatedAbort = 0x02
)

// define the reasons for an abort
const (
	reasonNotSpecified              = 0x00
	reasonUnrecognizedPDU           = 0x01
	reasonUnexpectedPDU             = 0x02
	reasonUnrecognizedPDUParameter  = 0x04
	reasonUnexpectedPDUParameter    = 0x05
	reasonInvalidPDUParamaeterValue = 0x06
)

// an abortPDU represents a PDU used to abort associations
type abortPDU struct {
	source byte // the initiator of the abort
	reason byte // the reason for the abort
}

// readAbortPDU reads an abortPDU after the base PDU has been read
func readAbortPDU(reader io.Reader) (*abortPDU, error) {

	// skip two bytes, as per the standard
	if err := skipBytes(reader, 2); err != nil {
		return nil, err
	}

	// read the source
	source, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// read the reason
	reason, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// construct the abort pdu
	abortPDU := &abortPDU{
		source: source,
		reason: reason,
	}

	return abortPDU, nil
}

func decodeAbortPDU(pdu *pdu) (*abortPDU, error) {
	return readAbortPDU(bytes.NewReader(pdu.buf))
}

// writeAbortPDU writes an AbortPDU to a writer
func writeAbortPDU(writer io.Writer, abortPDU *abortPDU) error {

	// create a pdu
	pdu, err := createAbortPDU(abortPDU)
	if err != nil {
		return err
	}

	// write the pdu
	if err := writePDU(writer, pdu); err != nil {
		return err
	}

	// return success
	return nil
}

// createAbortPDU creates an abort pdu
func createAbortPDU(abortPDU *abortPDU) (*pdu, error) {

	// create a byte writer
	byteWriter := new(bytes.Buffer)

	// write a zero byte, as per the standard
	if err := writeByte(byteWriter, 0x00); err != nil {
		return nil, err
	}

	// write a zero byte, as per the standard
	if err := writeByte(byteWriter, 0x00); err != nil {
		return nil, err
	}

	// write the source
	if err := writeByte(byteWriter, abortPDU.source); err != nil {
		return nil, err
	}

	// write the reason
	if err := writeByte(byteWriter, abortPDU.reason); err != nil {
		return nil, err
	}

	// create a pdu
	pdu := &pdu{
		typ: aAbortPDU,
		buf: byteWriter.Bytes(),
	}

	// return the podu and success
	return pdu, nil
}

// writeTo writes an abort pdu
func (abortPDU *abortPDU) writeTo(writer io.Writer) error {
	return writeAbortPDU(writer, abortPDU)
}

// define item types
const (
	appContextItemType     = 0x10
	rqPCItemType           = 0x20
	acPCItemType           = 0x21
	abstractSyntaxItemType = 0x30
	transferSyntaxItemType = 0x40
	userInfoItemType       = 0x50
)

// define sub item types
const (
	maxLengthItemType       = 0x51
	implClassUIDItemType    = 0x52
	maxNumOpsItemType       = 0x53
	implVersionNameItemType = 0x55
)

// assocPDU represents an association request or accept PDU
type assocPDU struct {
	protocol       uint16
	calledAETitle  string
	callingAETitle string
	appContextName string
	pcs            []*pc
	userInfo       *userInfo
}

// String returns a string representation of a assocPDU
func (assocPDU *assocPDU) String() string {
	return fmt.Sprintf(
		"{protocol:%v,calledAETitle:%q,callingAETitle:%q,appContextName:%q,pcs:%s,userInfo:%s}",
		assocPDU.protocol,
		assocPDU.calledAETitle,
		assocPDU.callingAETitle,
		assocPDU.appContextName,
		assocPDU.pcs,
		assocPDU.userInfo,
	)
}

// newAssocPDU creates a new association request or accept PDU
func newAssocPDU(calledAETitle string, callingAETitle string, capabilities *Capabilities) *assocPDU {

	// initialze the presentation contexts
	pcs := make([]*pc, 0, 5)

	// if capabilities were provided, add them to the presentation contexts
	if capabilities != nil {
		for i, capability := range capabilities.capabilities {
			pc := &pc{
				id:               byte(i*2 + 1), // odd numbers increasing order
				abstractSyntax:   capability.abstractSyntax,
				transferSyntaxes: capability.transferSyntaxes,
			}
			pcs = append(pcs, pc)
		}
	}

	// create a default associate request or accept pdu
	assocPDU := &assocPDU{
		ProtocolVersion,           // protocol version, as per the standard
		calledAETitle,             // title of the called, as per the standard
		callingAETitle,            // title of the caller, as per the standard
		ApplicationContextNameUID, // app context name, as per the standard
		pcs,                       // pres context list
		defaultUserInfo(),         // default user info
	}

	// return the pdu
	return assocPDU
}

// readAssocPDU reads an association accept or request from the reader
func readAssocPDU(reader io.Reader, pcItemType byte) (*assocPDU, error) {

	// read the protocol
	protocol, err := readShort(reader, binary.BigEndian)
	if err != nil {
		return nil, err
	}

	// skip two bytes, as per the standard
	if err := skipBytes(reader, 2); err != nil {
		return nil, err
	}

	// read the called AE title
	calledAETitle, err := readText(reader, 16)
	if err != nil {
		return nil, err
	}

	// read the calling AE title
	callingAETitle, err := readText(reader, 16)
	if err != nil {
		return nil, err
	}

	// skip thirty two bytes as per the standard
	if err := skipBytes(reader, 32); err != nil {
		return nil, err
	}

	// create an association accept or request pdu
	assocPDU := &assocPDU{
		protocol:       protocol,
		calledAETitle:  strings.TrimSpace(calledAETitle),
		callingAETitle: strings.TrimSpace(callingAETitle),
		pcs:            make([]*pc, 0, 5),
	}

	for {

		// read an item
		itemType, err := readByte(reader)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		// skip a byte, as per the standard
		if err := skipByte(reader); err != nil {
			return nil, err
		}

		// read the length
		length, err := readShort(reader, binary.BigEndian)
		if err != nil {
			return nil, err
		}

		// create a limited reader for the sub item
		limitedReader := io.LimitReader(reader, int64(length))

		if itemType == appContextItemType { // application context name

			// read the application context name
			assocPDU.appContextName, err = readUID(limitedReader, uint32(length))
			if err != nil {
				return nil, err
			}

		} else if itemType == pcItemType { // presentation context item type

			// read the presentation context
			pc, err := readPC(limitedReader, itemType)
			if err != nil {
				return nil, err
			}

			// add it to the list of presentation contexts
			assocPDU.pcs = append(assocPDU.pcs, pc)

		} else if itemType == 0x050 { // user info

			// read the user info sub items
			assocPDU.userInfo, err = readUserInfo(limitedReader)
			if err != nil {
				return nil, err
			}

		} else {
			return nil, fmt.Errorf("unrecognized item type: 0x%02X", itemType)
		}
	}

	// return the associate request or accept pdu
	return assocPDU, nil
}

// writeAssocPDU writes an associate PDU to a writer
func writeAssocPDU(writer io.Writer, assocPDU *assocPDU, pduType byte, pcItemType byte) error {

	// create a pdu
	pdu, err := createAssocPDU(assocPDU, pduType, pcItemType)
	if err != nil {
		return err
	}

	// write the pdu
	if err := writePDU(writer, pdu); err != nil {
		return err
	}

	// return success
	return nil
}

// createAssocPDU creates an associate PDU
func createAssocPDU(assocPDU *assocPDU, pduType byte, pcItemType byte) (*pdu, error) {

	// create a byte array output stream so we can calculate the length of the rest of the PDU
	byteWriter := new(bytes.Buffer)

	// write the protocol version
	if err := writeShort(byteWriter, assocPDU.protocol, binary.BigEndian); err != nil {
		return nil, err
	}

	// write a short zero
	if err := writeShort(byteWriter, 0x00, binary.BigEndian); err != nil {
		return nil, err
	}

	// write the called ae title
	if err := writeString(byteWriter, fmt.Sprintf("%-16s", assocPDU.calledAETitle)); err != nil {
		return nil, err
	}

	// write the calling ae title
	if err := writeString(byteWriter, fmt.Sprintf("%-16s", assocPDU.callingAETitle)); err != nil {
		return nil, err
	}

	// write thirty two zeroes, zero is the initial value for arrays, so this works
	var zeros [32]byte
	if err := writeBytes(byteWriter, zeros[:]); err != nil {
		return nil, err
	}

	// write the variable items
	if err := assocPDU.writeVariableItems(byteWriter, pcItemType); err != nil {
		return nil, err
	}

	// create a pdu
	pdu := &pdu{
		typ: pduType,
		buf: byteWriter.Bytes(),
	}

	// return the pdu and success
	return pdu, nil
}

// writeVariableItems writes the application context name, the presentation contexts and user info
func (assocPDU *assocPDU) writeVariableItems(writer io.Writer, pcItemType byte) error {

	// write the application context name
	if err := writeAppContextName(writer, assocPDU.appContextName); err != nil {
		return err
	}

	// write the presentation contexts
	if err := writePCs(writer, assocPDU.pcs, pcItemType); err != nil {
		return err
	}

	// write the user info
	if err := writeUserInfo(writer, assocPDU.userInfo); err != nil {
		return err
	}

	return nil
}

// writeAppContextName writes the application context name
func writeAppContextName(writer io.Writer, appContextName string) error {

	// write item type
	if err := writeByte(writer, appContextItemType); err != nil {
		return err
	}

	// write a zero as per the standard
	if err := writeByte(writer, 0x00); err != nil {
		return err
	}
	// write the length of the application context name
	if err := writeShort(writer, uint16(len(appContextName)), binary.BigEndian); err != nil {
		return err
	}

	// write the application context name
	if err := writeString(writer, appContextName); err != nil {
		return err
	}

	// all is well
	return nil
}

// newAssocRQPDU creates a new association request PDU
func newAssocRQPDU(calledAETitle string, callingAETitle string, capabilities *Capabilities) *assocPDU {
	return newAssocPDU(calledAETitle, callingAETitle, capabilities)
}

// readAssocRQPDU reads an associate request
func readAssocRQPDU(reader io.Reader) (*assocPDU, error) {
	return readAssocPDU(reader, rqPCItemType)
}

// writeAssocRQPDU writes an associate request
func writeAssocRQPDU(writer io.Writer, assocRQPDU *assocPDU) error {
	return writeAssocPDU(writer, assocRQPDU, aAssociateRQPDU, rqPCItemType)
}

// createAssocRQPDU creates an associate request
func createAssocRQPDU(assocRQPDU *assocPDU) (*pdu, error) {
	return createAssocPDU(assocRQPDU, aAssociateRQPDU, rqPCItemType)
}

// newAssocACPDU creates an associate accept PDU from an associate request PDU
func newAssocACPDU(assocRQPDU *assocPDU) *assocPDU {
	return newAssocPDU(assocRQPDU.calledAETitle, assocRQPDU.callingAETitle, nil)
}

// readAssocACPDU reads an associate accept
func readAssocACPDU(reader io.Reader) (*assocPDU, error) {
	return readAssocPDU(reader, acPCItemType)
}

// writeAssocACPDU writes an associate accept PDU
func writeAssocACPDU(writer io.Writer, assocACPDU *assocPDU) error {
	return writeAssocPDU(writer, assocACPDU, aAssociateACPDU, acPCItemType)
}

// createAssocACPDU creates an associate accept PDU
func createAssocACPDU(assocACPDU *assocPDU) (*pdu, error) {
	return createAssocPDU(assocACPDU, aAssociateACPDU, acPCItemType)
}

// define the results of association rejection
const (
	resultRejectedPermanent = 0x01
	resultRejectedTransient = 0x02
)

// define the sources of the results of association rejection
const (
	sourceServiceUser                                = 0x01
	sourceServiceProviderACSERelatedFunction         = 0x02
	sourceServiceProviderPresentationRelatedFunction = 0x04
)

// define ther reasons for the results of association rejection
const (
	reasonServiceUserNoReasonGiven                      = 0x01
	reasonServiceUserApplicationContextNameNotSupported = 0x02
	reasonServiceUserCallingAETitleNotRecognized        = 0x03
	reasonServiceUserCalledAETitleNotRecognized         = 0x07

	reasonServiceProviderACSERelatedFunctionNoReasonGiven               = 0x01
	reasonServiceProviderACSERelatedFunctionProtocolVersionNotSupported = 0x02

	reasonServiceProviderPresentationRelatedFunctionTemporaryCongestion = 0x01
	reasonServiceProviderPresentationRelatedFunctionLocalLimitExceeded  = 0x02
)

// assocRJPDU is an associate reject PDU
type assocRJPDU struct {
	result byte
	source byte
	reason byte
}

// String returns a string representation of an association reject PDU
func (assocRJPDU *assocRJPDU) String() string {
	return fmt.Sprintf(
		"{result:%d,source:%d,reason:%d}",
		assocRJPDU.result,
		assocRJPDU.source,
		assocRJPDU.reason,
	)
}

// readAssocRJDPU reads an association reject PDU
func readAssocRJPDU(reader io.Reader) (*assocRJPDU, error) {

	// skip a byte, as per the standard
	if err := skipByte(reader); err != nil {
		return nil, err
	}

	// read the result
	result, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// read the source
	source, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// read the reason
	reason, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// build the pdu
	assocRJPDU := &assocRJPDU{
		result: result,
		source: source,
		reason: reason,
	}

	// return the pdu
	return assocRJPDU, nil
}

// writeAssocRJPDU writes an associate reject pdu to a writer
func writeAssocRJPDU(writer io.Writer, assocRJPDU *assocRJPDU) error {

	// create a pdu
	pdu, err := createAssocRJPDU(assocRJPDU)
	if err != nil {
		return err
	}

	// write the pdu
	if err := writePDU(writer, pdu); err != nil {
		return err
	}

	// return success
	return nil
}

// createAssocRJPDU writes an associate reject PDU
func createAssocRJPDU(assocRJPDU *assocRJPDU) (*pdu, error) {

	// create a byte writer
	byteWriter := new(bytes.Buffer)

	// write a reserved byte
	if err := writeByte(byteWriter, 0x00); err != nil {
		return nil, err
	}

	// write the result
	if err := writeByte(byteWriter, assocRJPDU.result); err != nil {
		return nil, err
	}

	// write the source
	if err := writeByte(byteWriter, assocRJPDU.source); err != nil {
		return nil, err
	}

	// write the reason
	if err := writeByte(byteWriter, assocRJPDU.reason); err != nil {
		return nil, err
	}

	// create a pdu
	pdu := &pdu{
		typ: aAssociateRJPDU,
		buf: byteWriter.Bytes(),
	}

	// return the pdu and success
	return pdu, nil
}

// readReleasePDU reads an release request pdu
func readReleasePDU(reader io.Reader) error {

	// skip the long, as per the standard
	_, err := readLong(reader, binary.BigEndian)
	if err != nil {
		return err
	}

	// return success
	return nil
}

// writeReleasePDU writes a release request PDU to a writer
func writeReleasePDU(writer io.Writer, pduType byte) error {

	// create a pdu
	pdu, err := createReleasePDU(pduType)
	if err != nil {
		return err
	}

	// write the pdu
	if err := writePDU(writer, pdu); err != nil {
		return err
	}

	// return success
	return nil
}

func createReleasePDU(pduType byte) (*pdu, error) {

	// create a byte writer
	byteWriter := new(bytes.Buffer)

	// write a long
	if err := writeLong(byteWriter, 0x00, binary.BigEndian); err != nil {
		return nil, err
	}

	// create a pdu
	pdu := &pdu{
		typ: pduType,
		buf: byteWriter.Bytes(),
	}

	// return success
	return pdu, nil
}

// readReleaseRQPDU reads a release request PDU
func readReleaseRQPDU(reader io.Reader) error {
	return readReleasePDU(reader)
}

// writeReleaseRQPDU writes an release request PDU
func writeReleaseRQPDU(writer io.Writer) error {
	return writeReleasePDU(writer, aReleaseRQPDU)
}

func createReleaseRQPDU() (*pdu, error) {
	return createReleasePDU(aReleaseRQPDU)
}

// readReleaseRPPDU reads an associate request
func readReleaseRPPDU(reader io.Reader) error {
	return readReleasePDU(reader)
}

// writeReleaseRPPDU writes an release request PDU
func writeReleaseRPPDU(writer io.Writer) error {
	return writeReleasePDU(writer, aReleaseRPPDU)
}
func createReleaseRPPDU() (*pdu, error) {
	return createReleasePDU(aReleaseRPPDU)
}

// a dataTFPDU repesents a data transfer pdu
// not using this yet, but may do so in the future
type dataTFPDU struct {
	pdvs     []*pdv
	pdvIndex int
}

// readDataTFPDU parses a data transfer pdu
func readDataTFPDU(reader io.Reader) (*dataTFPDU, error) {

	// initialize a list of pdvs
	pdvs := make([]*pdv, 0, 1)

	for {

		// read a pdv until there are no more
		// would like to make this more efficient, as the bulk
		// of the pdv is really just a slice into the pdu.
		pdv, err := readPDV(reader)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		// append the pdv to the list of pdvs
		pdvs = append(pdvs, pdv)
	}

	// construct a data transfer pdu and initialize the index
	dataTFPDU := &dataTFPDU{
		pdvs:     pdvs,
		pdvIndex: 0,
	}

	// return the data transfer pdu and success
	return dataTFPDU, nil
}

// writeDataTFPDU writers a data transfer pdu
func writeDataTFPDU(writer io.Writer, dataTFPDU *dataTFPDU) error {

	// create a byte writer
	byteWriter := new(bytes.Buffer)

	// write the pdvs
	for _, pdv := range dataTFPDU.pdvs {
		if err := writePDV(byteWriter, pdv); err != nil {
			return err
		}
	}

	// create a pdu
	pdu := &pdu{
		typ: pDataTFPDU,
		buf: byteWriter.Bytes(),
	}

	// write the pdu
	if err := writePDU(writer, pdu); err != nil {
		return err
	}

	// return success
	return nil
}

// writeTo writes a data transfer PDU
func (dataTFPDU *dataTFPDU) writeTo(writer io.Writer) error {
	return writeDataTFPDU(writer, dataTFPDU)
}

// nextPDV returns the next PDV from a data transfer PDU
func (dataTFPDU *dataTFPDU) nextPDV() *pdv {

	// if we are at the end, return nil
	if dataTFPDU.pdvIndex >= len(dataTFPDU.pdvs) {
		return nil
	}

	// get the next pdv and increment the index
	pdv := dataTFPDU.pdvs[dataTFPDU.pdvIndex]
	dataTFPDU.pdvIndex++

	// return the pdv
	return pdv
}
