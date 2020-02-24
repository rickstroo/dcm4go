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
	localAETitle string,
	remoteAETitle string,
	capabilities []*Capability,
	opts *AssocOpts,
) (*RequestorAssoc, error) {

	// put together an association request pdu
	assocRQPDU := newAssocRQPDU(remoteAETitle, localAETitle, capabilities)
	log.Printf("assocRQPDU is %v", assocRQPDU)

	// write the pdu
	if err := writeAssocRQPDU(conn, assocRQPDU); err != nil {
		return nil, err
	}

	// read the response
	pdu, err := readPDU(conn)
	if err != nil {
		return nil, err
	}
	log.Printf("pdu is %v", pdu)

	// is this an abort request?  if so, return error
	if pdu.pduType == aAbortPDU {
		log.Printf("received an abort pdu")
		abortPDU, err := readAbortPDU(pdu)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("associate request aborted, %v", abortPDU)
	}

	// if this an associate reuject?  if so, return error
	if pdu.pduType == aAssociateRJPDU {
		log.Printf("received a rejection pdu")
		assocRJPDU, err := readAssocRJPDU(pdu)
		if err != nil {
			return nil, err
		}
		log.Printf("assocRJPDU is %v", assocRJPDU)

		return nil, fmt.Errorf("associate request rejected, %s", assocRJPDU)
	}

	// is this not an associate accept?  if not, return error
	if pdu.pduType != aAssociateACPDU {
		return nil, fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
	}

	assocACPDU, err := readAssocACPDU(pdu)
	if err != nil {
		return nil, err
	}
	log.Printf("assocACPDU is %v", assocACPDU)

	// create an association from the response
	assoc := &RequestorAssoc{
		Assoc{
			conn:       conn,
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
	if err := writeReleaseRQPDU(assoc.conn); err != nil {
		return err
	}
	log.Printf("wrote a release request\n")

	// read the response
	pdu, err := readPDU(assoc.conn)
	if err != nil {
		return err
	}
	log.Printf("pdu is %v\n", pdu)

	// is this an abort request?  if so, just return EOF
	if pdu.pduType == aAbortPDU {

		log.Printf("received an abort pdu")

		abortPDU, err := readAbortPDU(pdu)
		if err != nil {
			return err
		}
		log.Printf("read abort pdu, %v", abortPDU)

		// all is well
		return nil
	}

	if pdu.pduType != aReleaseRPPDU {
		return fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
	}

	log.Printf("received a release response pdu")
	releaseRPPDU, err := readReleaseRPPDU(pdu)
	if err != nil {
		return err
	}
	log.Printf("read release response pdu, %v", releaseRPPDU)

	// all is well
	return nil
}

// // WriteRequest writes a request
// func (assoc *RequestorAssoc) WriteRequest(
// 	pcID byte,
// 	command *Object,
// 	data *Object,
// 	reader io.Reader,
// ) error {
//
// 	// write the command
// 	if err := assoc.WriteCommand(pcID, command); err != nil {
// 		return err
// 	}
//
// 	// if there is data to be written, write it
// 	if data != nil {
// 		if err := assoc.WriteData(pcID, data); err != nil {
// 			return err
// 		}
// 	}
//
// 	// if there is data to be copied, copy it
// 	if reader != nil {
// 		num, err := assoc.CopyDataFrom(pcID, reader)
// 		if err != nil {
// 			return err
// 		}
// 		log.Printf("copied %d bytes", num)
// 	}
//
// 	// return success
// 	return nil
// }

// // ReadResponse reads a response
// func (assoc *RequestorAssoc) ReadResponse() (*PresContext, *Object, error) {
//
// 	// read the command and the presentation context id
// 	pcID, command, err := assoc.ReadCommand()
// 	if err != nil {
// 		return nil, nil, err
// 	}
//
// 	// look for the presentation context
// 	presContext, err := assoc.findAcceptedPresContextByPCID(pcID)
// 	if err != nil {
// 		return nil, nil, err
// 	}
//
// 	// return the presentation context and the command
// 	return presContext, command, nil
// }

// Echo sends a DICOM C-Echo request
func (assoc *RequestorAssoc) Echo() error {

	// find the accepted presentation context for this abstract syntax and any transfer syntax
	presContext, err := assoc.findAcceptedPresContextByCapability(VerificationUID, "*")
	if err != nil {
		return err
	}

	// create a verification request
	request := NewCEchoRequest()

	// write the verification request
	if err := assoc.WriteCommand(presContext.ID(), request); err != nil {
		return err
	}

	// read the response
	_, response, err := assoc.ReadCommand()
	if err != nil {
		return err
	}

	// get the status
	status, err := response.asShort(StatusTag, 0)
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

	// find the accepted presentation context for this transfer syntax
	presContext, err := assoc.findAcceptedPresContextByCapability(sopClassUID, transferSyntaxUID)
	if err != nil {
		return err
	}

	// create a c-store request
	command := NewCStoreRequest(sopClassUID, sopInstanceUID)

	// write the request, with data coming from the reader of the rest of the file
	if err := assoc.WriteCommand(presContext.ID(), command); err != nil {
		return err
	}

	// copy the data from the reader to the association
	num, err := assoc.CopyDataFrom(presContext.ID(), reader)
	if err != nil {
		return err
	}
	log.Printf("wrote %d bytes to the association", num)

	// read the response
	_, response, err := assoc.ReadCommand()
	if err != nil {
		return err
	}

	// get the status
	status, err := response.asShort(StatusTag, 0)
	if err != nil {
		return err
	}

	if status != 0 {
		return fmt.Errorf("status was %d, not success", status)
	}

	// otherwise, all is well
	return nil
}
