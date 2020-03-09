// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains the definition and common methods of an Assoc.

package dcm4go

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// Assoc represents a DICOM association
type Assoc struct {
	conn       net.Conn   // the connection used to exchange information
	pduReader  *pduReader // a reader of pdus
	pduWriter  *pduWriter // a writer of pdus
	ae         *AE        // the ae requesting or accepting the association
	assocRQPDU *assocPDU  // the associate request
	assocACPDU *assocPDU  // the associate response (if accepted)
}

// AssocOpts impact the behaviour of a Assoc.
type AssocOpts struct {
	WriteTimeOut time.Duration // a zero value means no write timeout
	ReadTimeOut  time.Duration // a zero value means no read timeout
	MaxBufLen    int           // a zero value defaults to 16K
}

// String returns a string representation of an association
func (assoc *Assoc) String() string {
	return fmt.Sprintf(
		"conn:{local:%v,remote:%v},ae:%v,assocRQPDU:%v,assocACPDU:%v",
		assoc.conn.LocalAddr(),
		assoc.conn.RemoteAddr(),
		assoc.ae,
		assoc.assocRQPDU,
		assoc.assocACPDU,
	)
}

// Conn returns the connection
func (assoc *Assoc) Conn() net.Conn {
	return assoc.conn
}

// CalledAETitle returns called ae title from the association request
func (assoc *Assoc) CalledAETitle() string {
	return strings.TrimSpace(assoc.assocRQPDU.calledAETitle)
}

// CallingAETitle returns calling ae title from the association request
func (assoc *Assoc) CallingAETitle() string {
	return strings.TrimSpace(assoc.assocRQPDU.callingAETitle)
}

// findAcceptedPCByCapability searches for a presentation context
// that was accepted for an abstract syntax and transfer syntax.
func (assoc *Assoc) findAcceptedPCByCapability(abstractSyntax string, transferSyntax string) (*pc, error) {

	// find the abstract syntax from the requested presentation contexts, there may be more than one
	for _, rqpc := range assoc.assocRQPDU.pcs {
		if rqpc.abstractSyntax == abstractSyntax {
			// now, look for the accepted presentation context for the same pcID that was requested
			for _, acpc := range assoc.assocACPDU.pcs {
				// if it's for the same id, and for the same transfer syntax id, and it was accepted
				if acpc.id == rqpc.id &&
					(transferSyntax == "*" || acpc.transferSyntaxes[0] == transferSyntax) &&
					acpc.result == pcAcceptance {
					return acpc, nil
				}
			}
		}
	}

	// we didn't find anything
	return nil, fmt.Errorf(
		"unable to find accepted presentation context for abstract syntax %q and transfer syntax %q",
		abstractSyntax,
		transferSyntax,
	)
}

// findAcceptedPCByPCID searches for a presentation context
// that was accepted for a presentation context id.
func (assoc *Assoc) findAcceptedPCByPCID(pcid byte) (*pc, error) {
	for _, pc := range assoc.assocACPDU.pcs {
		// find the accepted presentation context for the presentation context id
		if pc.id == pcid && pc.result == pcAcceptance {
			return pc, nil
		}
	}

	// we didn't find anything
	return nil, fmt.Errorf("unable to find accepted presentation context for presentation context id %d", pcid)
}

// findAcceptedTransferSyntaxByPCID finds the transfer syntax for the presentation
// context that was accepted for a presentation context id
func (assoc *Assoc) findAcceptedTransferSyntaxByPCID(pcid byte) (*transferSyntax, error) {
	pc, err := assoc.findAcceptedPCByPCID(pcid)
	if err != nil {
		return nil, err
	}
	transferSyntax, err := findTransferSyntax(pc.transferSyntaxes[0])
	if err != nil {
		return nil, err
	}
	return transferSyntax, nil
}

func (assoc *Assoc) writeMessage(message *Message) error {
	return writeMessage(assoc, message)
}

func (assoc *Assoc) readMessage(shouldReadData bool) (*Message, error) {
	return readMessage(assoc, shouldReadData)
}

func onAbort(reader io.Reader) error {
	abortPDU, err := readAbortPDU(reader)
	if err != nil {
		return err
	}
	log.Printf("received an abort pdu, %v", abortPDU)
	return fmt.Errorf("associate request aborted, %w", ErrAssociationAborted)
}

func onUnexpectedPDU(reader io.Reader, pdu *pdu) error {
	log.Printf("received unexpected pdu type, %v", pdu)
	return fmt.Errorf("unexpected pdu type, %d, %w", pdu.typ, ErrUnexpectedPDU)
}

// acceptAssoc accepts an association
func acceptAssoc(conn net.Conn, ae *AE, capabilities *Capabilities) (*Assoc, error) {

	// I've decided not to implement a state machine.
	// I've looked at a number of implementations and it looks
	// to me like a state machine makes it really hard to follow
	// all the logic.  So, in the spirit of writing easy to
	// read programs, I will implement the logic of the state
	// machine in the AcceptAssoc and RequestAssoc structs.

	// create a pdu reader and pdu writer
	pduReader := newPDUReader(conn)
	pduWriter := newPDUWriter(conn)

	// read a pdu
	pdu, err := pduReader.nextPDU()
	if err != nil {
		return nil, err
	}
	log.Printf("pdu is %v\n", pdu)

	// if abort, we simply exit
	if pdu.typ == aAbortPDU {
		return nil, onAbort(pduReader)
	}

	// if anything other than an associate request, we abort
	if pdu.typ != aAssociateRQPDU {
		return nil, onUnexpectedPDU(pduReader, pdu)
	}

	// read the associate request
	assocRQPDU, err := readAssocRQPDU(pduReader)
	if err != nil {
		return nil, err
	}
	log.Printf("assocRQPDU is %v\n", assocRQPDU)

	// attempt to negotiate an association
	assocACPDU, assocRJPDU, err := negotiateAssoc(assocRQPDU, ae, capabilities)
	if err != nil {
		return nil, err
	}
	if assocACPDU != nil {
		log.Printf("accepted associate request, assocACPDU is %v\n", assocACPDU)
	}
	if assocRJPDU != nil {
		log.Printf("rejected associate request, assocRJPDU is %v\n", assocRJPDU)
	}

	// was association rejected
	if assocRJPDU != nil {

		// write the associate reject pdu
		if err := writeAssocRJPDU(pduWriter, assocRJPDU); err != nil {
			return nil, err
		}
		// let the caller know that the associate request was rejected
		return nil, ErrAssociateRequestRejected
	}

	// otherwise, write the associate accept pdu
	if err := writeAssocACPDU(pduWriter, assocACPDU); err != nil {
		return nil, err
	}

	// construct an association
	assoc := &Assoc{
		conn:       conn,
		pduReader:  pduReader,
		pduWriter:  pduWriter,
		ae:         ae,
		assocRQPDU: assocRQPDU,
		assocACPDU: assocACPDU,
	}
	log.Printf("assoc is %v\n", assoc)

	// return the association to the caller
	return assoc, nil
}

// negotiateAssoc determines what requested presentation contexts
// are accepted based on the presentation contexts that are supported
// by the ae
func negotiateAssoc(assocRQPDU *assocPDU, ae *AE, capabilities *Capabilities) (*assocPDU, *assocRJPDU, error) {

	// reject if the called ae title does not match the given ae title
	calledAETitle := strings.TrimSpace(assocRQPDU.calledAETitle)
	if calledAETitle != ae.Title {
		// create and return an associate reject pdu
		assocRJPDU := &assocRJPDU{
			result: resultRejectedPermanent,
			source: sourceServiceProviderACSERelatedFunction,
			reason: reasonServiceUserCalledAETitleNotRecognized,
		}
		return nil, assocRJPDU, nil
	}

	// initialize the association accept pdu
	assocACPDU := newAssocACPDU(assocRQPDU)

	// negotiate each of the presentation contexts
	for _, rqPresContext := range assocRQPDU.pcs {
		pc, err := negotiatePresContext(rqPresContext, capabilities)
		if err != nil {
			return nil, nil, err
		}
		assocACPDU.pcs = append(assocACPDU.pcs, pc)
	}

	return assocACPDU, nil, nil
}

// negotiationPresContext negotiates a single presentation context
func negotiatePresContext(rqpc *pc, capabilities *Capabilities) (*pc, error) {

	// look for a capability for this abstract syntax
	capability, found := findAbstractSyntaxCapability(rqpc.abstractSyntax, capabilities)

	// if we don't find one, return a failure for this requested presentation context
	if !found {
		pc := &pc{
			id:     rqpc.id,                      // the id
			result: pcAbstractSyntaxNotSupported, // reason for failure
		}

		// return the failed presentation context
		return pc, nil
	}

	// look for a matching transfer syntax
	transferSyntax, found := findTransferSyntaxCapability(rqpc.transferSyntaxes, capability)

	// if we didn't find one, return failure
	if !found {
		pc := &pc{
			id:     rqpc.id,                        // the id
			result: pcTransferSyntaxesNotSupported, // reason for failure
		}

		// return the failed presentation context
		return pc, nil
	}

	// found one, create an accepted presentation context
	pc := &pc{
		id:               rqpc.id,                  // the id
		transferSyntaxes: []string{transferSyntax}, // the transfer syntax
		result:           pcAcceptance,             // acceptance
	}

	// return the accepted presentation context
	return pc, nil
}

// findAbstractSyntaxCapability searches for a capability for an abstract syntax
func findAbstractSyntaxCapability(rqAbstractSyntax string, capabilities *Capabilities) (*Capability, bool) {
	for _, capability := range capabilities.capabilities {
		if rqAbstractSyntax == capability.abstractSyntax {
			return capability, true
		}
	}
	return nil, false
}

// findTransferSyntaxCapability searches for a capability for a transfer syntax
func findTransferSyntaxCapability(rqTransferSyntaxes []string, capability *Capability) (string, bool) {
	for _, rqTransferSyntax := range rqTransferSyntaxes {
		for _, transferSyntax := range capability.transferSyntaxes {
			if rqTransferSyntax == transferSyntax {
				return rqTransferSyntax, true
			}
		}
	}
	return "", false
}

// ReadRequest read a request
func (assoc *Assoc) ReadRequest() (*Message, error) {

	// read the next PDU
	pdu, err := assoc.pduReader.nextPDU()
	if err != nil {
		return nil, err
	}

	// is this an association release request?  if so, write response and return EOF
	if pdu.typ == aReleaseRQPDU {
		return nil, assoc.onRelease()
	}

	// is this an abort request?  if so, just return EOF
	if pdu.typ == aAbortPDU {
		return nil, onAbort(assoc.pduReader)
	}

	// is this not a data transfer request?
	if pdu.typ != pDataTFPDU {
		return nil, onUnexpectedPDU(assoc.pduReader, pdu)
	}

	// read the message
	return assoc.readMessage(false)
}

func (assoc *Assoc) onRelease() error {
	if err := readReleaseRQPDU(assoc.pduReader); err != nil {
		return err
	}

	if err := writeReleaseRPPDU(assoc.pduWriter); err != nil {
		return err
	}

	return io.EOF
}

// WriteResponse writes a response
func (assoc *Assoc) WriteResponse(message *Message) error {
	return assoc.writeMessage(message)
}

// requestAssoc is used to request an association.
func requestAssoc(
	conn net.Conn,
	localAE *AE,
	remoteAE *AE,
	capabilities *Capabilities,
	opts *AssocOpts,
) (*Assoc, error) {

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
	if pdu.typ == aAbortPDU {
		return nil, onAbort(pduReader)
	}

	// if this an associate reuject?  if so, return error
	if pdu.typ == aAssociateRJPDU {
		assocRJPDU, err := readAssocRJPDU(pduReader)
		if err != nil {
			return nil, err
		}
		log.Printf("received a associate rejection pdu, %v", assocRJPDU)
		return nil, fmt.Errorf("associate request rejeced, %w", ErrAssociateRequestRejected)
	}

	// is this not an associate accept?
	if pdu.typ != aAssociateACPDU {
		return nil, onUnexpectedPDU(pduReader, pdu)
	}

	assocACPDU, err := readAssocACPDU(pduReader)
	if err != nil {
		return nil, err
	}
	log.Printf("received a associate acceptance pdu, %v", assocACPDU)

	// create an association from the response
	assoc := &Assoc{
		conn:       conn,
		pduReader:  pduReader,
		pduWriter:  pduWriter,
		assocRQPDU: assocRQPDU,
		assocACPDU: assocACPDU,
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
	if pdu.typ == aAbortPDU {
		return onAbort(assoc.pduReader)
	}

	// is this not the pdu we are expecting?
	if pdu.typ != aReleaseRPPDU {
		return onUnexpectedPDU(assoc.pduReader, pdu)
	}

	if err := readReleaseRPPDU(assoc.pduReader); err != nil {
		return err
	}
	log.Printf("received a release response pdu")

	// all is well
	return nil
}

// WriteRequest writes a request
func (assoc *Assoc) WriteRequest(message *Message) error {
	return assoc.writeMessage(message)
}

// ReadResponse reads a response
func (assoc *Assoc) ReadResponse() (*Message, error) {

	// read the next PDU
	pdu, err := assoc.pduReader.nextPDU()
	if err != nil {
		return nil, err
	}

	// is this an abort request?  if so, just return EOF
	if pdu.typ == aAbortPDU {
		return nil, onAbort(assoc.pduReader)
	}

	// is this not a data transfer request?
	if pdu.typ != pDataTFPDU {
		return nil, onUnexpectedPDU(assoc.pduReader, pdu)
	}

	return assoc.readMessage(false)
}

// Echo sends a DICOM C-Echo request
func (assoc *Assoc) Echo() error {

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
func (assoc *Assoc) Store(reader io.Reader) error {

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
