// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains methods used by the requestor of
// an association.

package dcm4go

import (
	"fmt"
	"io"
	"log"
	"net"
)

// RequestorAssoc is a type of Assoc, used by requestors of associations.
type RequestorAssoc struct {
	Assoc
}

// requestAssoc is used to request an association.
func requestAssoc(
	conn net.Conn,
	localAE *AE,
	remoteAE *AE,
	capabilities []*PresContext,
	opts *AssocOpts,
) (*RequestorAssoc, error) {

	// create a pdu reader and pdu writer
	pduReader := newPDUReader(conn)
	pduWriter := newPDUWriter(conn)

	// put together an association request pdu
	assocRQPDU := newAssocRQPDU(remoteAE.Title, localAE.Title, capabilities)
	log.Printf("assocRQPDU is %v", assocRQPDU)

	// write the pdu
	if err := writeAssocRQPDU(pduWriter, assocRQPDU); err != nil {
		return nil, err
	}

	// read the response
	pdu, err := pduReader.nextPDU()
	if err != nil {
		return nil, err
	}
	log.Printf("read pdu, %v", pdu)

	// is this an abort request?
	if pdu.pduType == aAbortPDU {
		return nil, onAbort(pduReader)
	}

	// if this an associate reuject?  if so, return error
	if pdu.pduType == aAssociateRJPDU {
		assocRJPDU, err := readAssocRJPDU(pduReader)
		if err != nil {
			return nil, err
		}
		log.Printf("received a associate rejection pdu, %v", assocRJPDU)
		return nil, fmt.Errorf("associate request rejeced, %w", ErrAssociateRequestRejected)
	}

	// is this not an associate accept?
	if pdu.pduType != aAssociateACPDU {
		return nil, onUnexpectedPDU(pduReader, pdu)
	}

	assocACPDU, err := readAssocACPDU(pduReader)
	if err != nil {
		return nil, err
	}
	log.Printf("received a associate acceptance pdu, %v", assocACPDU)

	// create an association from the response
	assoc := &RequestorAssoc{
		Assoc{
			conn:       conn,
			pduReader:  pduReader,
			pduWriter:  pduWriter,
			assocRQPDU: assocRQPDU,
			assocACPDU: assocACPDU,
		},
	}
	log.Printf("created association from %v to %v", assoc.CallingAETitle(), assoc.CalledAETitle())

	// return the association
	return assoc, nil
}

// RequestRelease requests release from an association
func (assoc *Assoc) RequestRelease() error {

	// write a request release pdu
	if err := writeReleaseRQPDU(assoc.pduWriter); err != nil {
		return err
	}
	log.Printf("wrote a release request\n")

	// read the response
	pdu, err := assoc.pduReader.nextPDU()
	if err != nil {
		return err
	}
	log.Printf("pdu is %v\n", pdu)

	// is this an abort request?
	if pdu.pduType == aAbortPDU {
		return onAbort(assoc.pduReader)
	}

	// is this not the pdu we are expecting?
	if pdu.pduType != aReleaseRPPDU {
		return onUnexpectedPDU(assoc.pduReader, pdu)
	}

	releaseRPPDU, err := readReleaseRPPDU(assoc.pduReader)
	if err != nil {
		return err
	}
	log.Printf("received a release response pdu, %v", releaseRPPDU)

	// all is well
	return nil
}

// WriteRequest writes a request
func (assoc *RequestorAssoc) WriteRequest(message *Message) error {
	return assoc.writeMessage(message)
}

// ReadResponse reads a response
func (assoc *RequestorAssoc) ReadResponse() (*Message, error) {

	// read the next PDU
	pdu, err := assoc.pduReader.nextPDU()
	if err != nil {
		return nil, err
	}

	// is this an abort request?  if so, just return EOF
	if pdu.pduType == aAbortPDU {
		return nil, onAbort(assoc.pduReader)
	}

	// is this not a data transfer request?
	if pdu.pduType != pDataTFPDU {
		return nil, onUnexpectedPDU(assoc.pduReader, pdu)
	}

	return assoc.readMessage(false)
}

// Echo sends a DICOM C-Echo request
func (assoc *RequestorAssoc) Echo() error {

	// create a verification request
	request, err := NewCEchoRequest(assoc)
	if err != nil {
		return err
	}

	// write the verification request
	if err := assoc.WriteRequest(request); err != nil {
		return err
	}

	// read the response
	response, err := assoc.ReadResponse()
	if err != nil {
		return err
	}

	// get the status
	status, err := response.Command().asShort(StatusTag, 0)
	if err != nil {
		return err
	}

	if status != 0 {
		return fmt.Errorf("status was %d, not success", status)
	}

	return nil
}

// Store sends a DICOM C-Store request
func (assoc *RequestorAssoc) Store(reader io.Reader) error {

	// read the group two attributes
	groupTwo, err := ReadGroupTwo(reader, 0)
	if err != nil {
		return err
	}

	// get the sop class uid of the stored object
	sopClassUID, err := groupTwo.AsString(MediaStorageSOPClassUIDTag, 0)
	if err != nil {
		return err
	}

	// get the sop class instance UID of the stored object
	sopInstanceUID, err := groupTwo.AsString(MediaStorageSOPInstanceUIDTag, 0)
	if err != nil {
		return err
	}

	// get the transfer syntax used to store the file
	transferSyntaxUID, err := groupTwo.AsString(TransferSyntaxUIDTag, 0)
	if err != nil {
		return err
	}

	// create a c-store request
	request, err := NewCStoreRequest(assoc, sopClassUID, sopInstanceUID, transferSyntaxUID, reader)
	if err != nil {
		return err
	}

	// write the request, with data coming from the reader of the rest of the file
	if err := assoc.WriteRequest(request); err != nil {
		return err
	}

	// read the response
	response, err := assoc.ReadResponse()
	if err != nil {
		return err
	}

	// get the status
	status, err := response.Command().asShort(StatusTag, 0)
	if err != nil {
		return err
	}

	if status != 0 {
		return fmt.Errorf("status was %d, not success", status)
	}

	// otherwise, all is well
	return nil
}
