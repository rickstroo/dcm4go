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
func acceptAssoc(conn net.Conn, ae *AE, capabilities []*Capability) (*AcceptorAssoc, error) {

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
		log.Printf("assocACPDU is %v\n", assocACPDU)
	}
	if assocRJPDU != nil {
		log.Printf("assocRJPDU is %v\n", assocRJPDU)
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
func negotiateAssoc(assocRQPDU *AssocRQPDU, ae *AE, capabilities []*Capability) (*AssocACPDU, *AssocRJPDU, error) {

	// reject if the called ae title does not match the given ae title
	calledAETitle := strings.TrimSpace(assocRQPDU.calledAETitle)
	if calledAETitle != ae.Title() {
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
func negotiatePresContext(rqPresContext *PresContext, capabilities []*Capability) (*PresContext, error) {

	// look for a capability for this abstract syntax
	capability, found := findCapability(rqPresContext.abstractSyntax, capabilities)

	// if we don't find one, return a failure for this requested presentation context
	if !found {
		acPresContext := &PresContext{
			rqPresContext.id,             // the id
			"",                           // no abstract syntax
			nil,                          // no transfer syntaxes
			pcAbstractSyntaxNotSupported, // failure
		}
		return acPresContext, nil
	}

	// if we found one, now we look for a matching transfer syntax
	for _, rqTansferSyntax := range rqPresContext.transferSyntaxes {
		if contains(capability.TransferSyntaxes, rqTansferSyntax) {
			acPresContext := &PresContext{
				rqPresContext.id,          // the id
				"",                        // no abstract syntax
				[]string{rqTansferSyntax}, // the transfer syntax
				pcAcceptance,              // success
			}
			return acPresContext, nil
		}
	}

	// we didn't find a matching transfer syntax, so return a failed acceptance presentation context
	acPresContext := &PresContext{
		rqPresContext.id,               // the id
		"",                             // no abstract syntax
		nil,                            // no transfer syntaxes
		pcTransferSyntaxesNotSupported, // failure
	}

	// return the accepted presentation context
	return acPresContext, nil
}

// findCapability searches for a capability for an abstract syntax
func findCapability(abstractSyntax string, capabilities []*Capability) (*Capability, bool) {
	for _, capability := range capabilities {
		if abstractSyntax == capability.AbstractSyntax {
			return capability, true
		}
	}
	return nil, false
}

// contains looks for a string in a set of strings
func contains(ses []string, t string) bool {
	for _, s := range ses {
		if s == t {
			return true
		}
	}
	return false
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
		if err := readReleaseRQPDU(pdu); err != nil {
			return nil, err
		}

		releaseRPPDU := &ReleaseRPPDU{}
		if err := releaseRPPDU.Write(assoc.conn); err != nil {
			return nil, err
		}

		return nil, io.EOF
	}

	// is this an abort request?  if so, just return EOF
	if pdu.pduType == aAbortPDU {

		log.Printf("received an abort pdu")

		abortPDU, err := readAbortPDU(assoc.pduReader)
		if err != nil {
			return nil, err
		}
		log.Printf("read abort pdu, %v", abortPDU)

		// return eof
		return nil, io.EOF
	}

	// is this not a data transfer request?
	if pdu.pduType != pDataTFPDU {
		return nil, ErrUnexpectedPDU
	}

	// read the message
	return assoc.readMessage(false)
}

// WriteResponse writes a response
func (assoc *AcceptorAssoc) WriteResponse(message *Message) error {
	return assoc.writeMessage(message)
}
