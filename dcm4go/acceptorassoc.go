// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains methods used by the acceptor of
// an association.

package dcm4go

import (
	"io"
	"log"
	"net"
	"strings"
)

// AcceptorAssoc is a type of Assoc, used by acceptors of associations.
type AcceptorAssoc struct {
	Assoc
}

// acceptAssoc accepts an association
func acceptAssoc(conn net.Conn, ae *AE, capabilities []*PresContext) (*AcceptorAssoc, error) {

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
	if pdu.pduType == aAbortPDU {
		return nil, onAbort(pduReader)
	}

	// if anything other than an associate request, we abort
	if pdu.pduType != aAssociateRQPDU {
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
		if err := assocRJPDU.Write(pduWriter); err != nil {
			return nil, err
		}
		// let the caller know that the associate request was rejected
		return nil, ErrAssociateRequestRejected
	}

	// otherwise, write the associate accept pdu
	if err := assocACPDU.Write(pduWriter); err != nil {
		return nil, err
	}

	// construct an association
	assoc := &AcceptorAssoc{
		Assoc{
			conn:       conn,
			pduReader:  pduReader,
			pduWriter:  pduWriter,
			ae:         ae,
			assocRQPDU: assocRQPDU,
			assocACPDU: assocACPDU,
		},
	}
	log.Printf("assoc is %v\n", assoc)

	// return the association to the caller
	return assoc, nil
}

// negotiateAssoc determines what requested presentation contexts
// are accepted based on the presentation contexts that are supported
// by the ae
func negotiateAssoc(assocRQPDU *AssocRQPDU, ae *AE, capabilities []*PresContext) (*AssocACPDU, *AssocRJPDU, error) {

	// reject if the called ae title does not match the given ae title
	calledAETitle := strings.TrimSpace(assocRQPDU.calledAETitle)
	if calledAETitle != ae.Title {
		// create and return an associate reject pdu
		assocRJPDU := &AssocRJPDU{
			result: resultRejectedPermanent,
			source: sourceServiceProviderACSERelatedFunction,
			reason: reasonServiceUserCalledAETitleNotRecognized,
		}
		return nil, assocRJPDU, nil
	}

	// initialize the association accept pdu
	assocACPDU := newAssocACPDU(assocRQPDU)

	// negotiate each of the presentation contexts
	for _, rqPresContext := range assocRQPDU.presContexts {
		acPresContext, err := negotiatePresContext(rqPresContext, capabilities)
		if err != nil {
			return nil, nil, err
		}
		assocACPDU.presContexts = append(assocACPDU.presContexts, acPresContext)
	}

	return assocACPDU, nil, nil
}

// negotiationPresContext negotiates a single presentation context
func negotiatePresContext(rqPresContext *PresContext, capabilities []*PresContext) (*PresContext, error) {

	// look for a capability for this abstract syntax
	capability, found := findAbstractSyntaxCapability(rqPresContext.AbstractSyntax, capabilities)

	// if we don't find one, return a failure for this requested presentation context
	if !found {
		presContext := &PresContext{
			ID:     rqPresContext.ID,             // the id
			Result: pcAbstractSyntaxNotSupported, // reason for failure
		}

		// return the failed presentation context
		return presContext, nil
	}

	// look for a matching transfer syntax
	transferSyntax, found := findTransferSyntaxCapability(rqPresContext.TransferSyntaxes, capability)

	// if we didn't find one, return failure
	if !found {
		presContext := &PresContext{
			ID:     rqPresContext.ID,               // the id
			Result: pcTransferSyntaxesNotSupported, // reason for failure
		}

		// return the failed presentation context
		return presContext, nil
	}

	// found one, create an accepted presentation context
	presContext := &PresContext{
		ID:               rqPresContext.ID,         // the id
		TransferSyntaxes: []string{transferSyntax}, // the transfer syntax
		Result:           pcAcceptance,             // acceptance
	}

	// return the accepted presentation context
	return presContext, nil
}

// findAbstractSyntaxCapability searches for a capability for an abstract syntax
func findAbstractSyntaxCapability(rqAbstractSyntax string, capabilities []*PresContext) (*PresContext, bool) {
	for _, capability := range capabilities {
		if rqAbstractSyntax == capability.AbstractSyntax {
			return capability, true
		}
	}
	return nil, false
}

// findTransferSyntaxCapability searches for a capability for a transfer syntax
func findTransferSyntaxCapability(rqTransferSyntaxes []string, capability *PresContext) (string, bool) {
	for _, rqTransferSyntax := range rqTransferSyntaxes {
		for _, transferSyntax := range capability.TransferSyntaxes {
			if rqTransferSyntax == transferSyntax {
				return rqTransferSyntax, true
			}
		}
	}
	return "", false
}

// ReadRequest read a request
func (assoc *AcceptorAssoc) ReadRequest() (*Message, error) {

	// read the next PDU
	pdu, err := assoc.pduReader.nextPDU()
	if err != nil {
		return nil, err
	}

	// is this an association release request?  if so, write response and return EOF
	if pdu.pduType == aReleaseRQPDU {
		return nil, assoc.onRelease()
	}

	// is this an abort request?  if so, just return EOF
	if pdu.pduType == aAbortPDU {
		return nil, onAbort(assoc.pduReader)
	}

	// is this not a data transfer request?
	if pdu.pduType != pDataTFPDU {
		return nil, onUnexpectedPDU(assoc.pduReader, pdu)
	}

	// read the message
	return assoc.readMessage(false)
}

func (assoc *AcceptorAssoc) onRelease() error {
	if err := readReleaseRQPDU(assoc.pduReader); err != nil {
		return err
	}

	releaseRPPDU := &ReleaseRPPDU{}
	if err := releaseRPPDU.Write(assoc.pduWriter); err != nil {
		return err
	}

	return io.EOF
}

// WriteResponse writes a response
func (assoc *AcceptorAssoc) WriteResponse(message *Message) error {
	return assoc.writeMessage(message)
}
