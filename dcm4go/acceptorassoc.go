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

// AcceptAssoc accepts an association
func AcceptAssoc(conn net.Conn, ae *AE, handlers []Handler) (*AcceptorAssoc, error) {

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
		return nil, io.EOF
	}

	// if anything other than an associate request, we abort
	if pdu.pduType != aAssociateRQPDU {

		log.Printf("unexpected pdu type, %d\n", pdu.pduType)

		// construct an abort pdu
		abortPDU := &AbortPDU{
			source: sourceServiceProviderInitiatedAbort, // the provider is initiating the abort
			reason: reasonUnexpectedPDU,                 // didn't expect this pdu
		}

		// attempt to write it
		if err := abortPDU.Write(pduWriter); err != nil {
			return nil, err
		}

		// let the caller know why we were not able to negotiate an association
		return nil, ErrUnexpectedPDU
	}

	// read the associate request
	assocRQPDU, err := readAssocRQPDU(pduReader)
	if err != nil {
		return nil, err
	}
	log.Printf("assocRQPDU is %v\n", assocRQPDU)

	// attempt to negotiate an association
	assocACPDU, assocRJPDU, err := negotiateAssoc(assocRQPDU, ae, handlers)
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
func negotiateAssoc(assocRQPDU *AssocRQPDU, ae *AE, handlers []Handler) (*AssocACPDU, *AssocRJPDU, error) {

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
		acPresContext, err := negotiatePresContext(rqPresContext, handlers)
		if err != nil {
			return nil, nil, err
		}
		assocACPDU.presContexts = append(assocACPDU.presContexts, acPresContext)
	}

	return assocACPDU, nil, nil
}

// negotiationPresContext negotiates a single presentation context
func negotiatePresContext(rqPresContext *PresContext, handlers []Handler) (*PresContext, error) {

	// look for a capability for this abstract syntax
	handler, capability, found := findCapability(rqPresContext.abstractSyntax, handlers)

	// if we don't find one, return a failure for this requested presentation context
	if !found {
		acPresContext := &PresContext{
			rqPresContext.id,             // the id
			"",                           // no abstract syntax
			nil,                          // no transfer syntaxes
			pcAbstractSyntaxNotSupported, // failure
			nil,                          // no handler
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
				handler,                   // the handler
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
		nil,                            // no handler
	}

	// return the accepted presentation context
	return acPresContext, nil
}

// findCapability searches for a capability for an abstract syntax
func findCapability(abstractSyntax string, handlers []Handler) (Handler, *Capability, bool) {
	for _, handler := range handlers {
		for _, capability := range handler.Capabilities() {
			if abstractSyntax == capability.AbstractSyntax {
				return handler, capability, true
			}
		}
	}
	return nil, nil, false
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

// ReadRequest reads a request from the association
func (assoc *Assoc) ReadRequest() (*Message, error) {
	return assoc.ReadRequestOrResponse()
}

// Serve reads and services a single request
func (assoc *Assoc) Serve(handler Handler) error {

	// read a pdu
	pdu, err := assoc.pduReader.nextPDU()
	if err != nil {
		return err
	}
	log.Printf("pdu is %v\n", pdu)

	// is this an association release request?
	/// if so, write release response and return EOF
	if pdu.pduType == aReleaseRQPDU {

		log.Printf("received release request, attempting to release association\n")

		if err := readReleaseRQPDU(assoc.pduReader); err != nil {
			return err
		}

		// construct a release response pdu
		releaseRPPDU := &ReleaseRPPDU{}
		if err := releaseRPPDU.Write(assoc.pduWriter); err != nil {
			return err
		}

		// return EOF to indicate that the association is released
		return io.EOF
	}

	// is this an abort request?  if so, simply return EOF
	if pdu.pduType == aAbortPDU {
		log.Printf("received abort request, aborting association\n")
		return io.EOF
	}

	// if anything other than an data transfer request, we abort
	if pdu.pduType != pDataTFPDU {

		log.Printf("unexpected pdu type, %d\n", pdu.pduType)

		// construct an abort pdu
		abortPDU := &AbortPDU{
			source: sourceServiceProviderInitiatedAbort, // the provider is initiating the abort
			reason: reasonUnexpectedPDU,                 // didn't expect this pdu
		}

		// attempt to write it
		if err := abortPDU.Write(assoc.pduWriter); err != nil {
			return err
		}

		// let the caller know why we were not able to negotiate an association
		return ErrUnexpectedPDU
	}

	log.Printf("attempting to accept data transfer\n")

	// create a reader for the command
	commandReader, err := newPDVReader(assoc.pduReader, true)
	if err != nil {
		return err
	}

	// get the presentation context id from the reader
	pcID := commandReader.pdv.pcID

	// read the command
	command, err := readCommand(commandReader)
	if err != nil {
		return err
	}
	log.Printf("command is %v\n", command)

	// find the persentation context by id
	presContext, err := assoc.findAcceptedPresContextByPCID(pcID)
	if err != nil {
		return err
	}
	log.Printf("presContext is %v\n", presContext)

	// get the command data set
	commandDataSet, err := command.asShort(CommandDataSetTypeTag, 0)
	if err != nil {
		return err
	}

	// get a data reader if required
	dataReader, err := getDataReader(commandDataSet, assoc)
	if err != nil {
		return err
	}

	if handler != nil {
		handler.HandleRequest(assoc, presContext, command, dataReader)
	} else {
		// call the handler for the command
		if err := presContext.handler.HandleRequest(assoc, presContext, command, dataReader); err != nil {
			return err
		}
	}

	// all is well
	return nil
}

func getDataReader(commandDataSet uint16, assoc *Assoc) (*pdvReader, error) {

	// check to see if data is present
	if isDataSetPresent(commandDataSet) {

		// create a reader for the data
		dataReader, err := newPDVReader(assoc.pduReader, false)
		if err != nil {
			return nil, err
		}

		// return the data reader
		return dataReader, nil
	}

	// return nothing, as no data reader is required
	return nil, nil
}
