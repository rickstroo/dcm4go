package dcm4go

import (
	"io"
	"log"
	"net"
)

// AcceptorAssoc is an association used by acceptors of associations
type AcceptorAssoc struct {
	Assoc
}

// String prints a string representation of an acceptor association
func (assoc *AcceptorAssoc) String() string {
	return assoc.Assoc.String()
}

// AcceptAssoc accepts an association
func AcceptAssoc(conn net.Conn, ae *AE, handlers []Handler) (*AcceptorAssoc, error) {

	// this should really be handled as a state machine
	// will think about doing that later
	// for now, want to focus on getting the data transfer
	// mechanisms working

	// read a pdu
	pdu, err := readPDU(conn)
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
		if err := abortPDU.Write(conn); err != nil {
			return nil, err
		}

		// let the caller know why we were not able to negotiate an association
		return nil, ErrUnexpectedPDU
	}

	// read the associate request
	assocRQPDU, err := readAssocRQPDU(pdu)
	if err != nil {
		return nil, err
	}
	log.Printf("assocRQPDU is %v\n", assocRQPDU)

	// attempt to negotiate an association
	assocACPDU, err := negotiateAssoc(assocRQPDU, ae, handlers)
	if err != nil {
		return nil, err
	}
	log.Printf("assocACPDU is %v\n", assocACPDU)

	// hmm, this might be a rejection, need to handle that as well

	if err := writeAssocACPDU(conn, assocACPDU); err != nil {
		return nil, err
	}

	assoc := &AcceptorAssoc{
		Assoc{
			conn:       conn,
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
func negotiateAssoc(assocRQPDU *AssocRQPDU, ae *AE, handlers []Handler) (*AssocACPDU, error) {

	// initialize the association accept pdu
	assocACPDU := newAssocACPDU(assocRQPDU)

	// negotiate each of the presentation contexts
	for _, rqPresContext := range assocRQPDU.presContexts {
		acPresContext, err := negotiatePresContext(rqPresContext, handlers)
		if err != nil {
			return nil, err
		}
		assocACPDU.presContexts = append(assocACPDU.presContexts, acPresContext)
	}

	return assocACPDU, nil
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
func (assoc *AcceptorAssoc) ReadRequest() (*Message, error) {
	return assoc.ReadRequestOrResponse()
}

// Serve reads and services a single request
func (assoc *AcceptorAssoc) Serve() error {

	// read a pdu
	pdu, err := readPDU(assoc.conn)
	if err != nil {
		return err
	}
	log.Printf("pdu is %v\n", pdu)

	// is this an association release request?  if so, write response and return EOF
	if pdu.pduType == aReleaseRQPDU {

		log.Printf("received release request, attempting to release association\n")

		if err := readReleaseRQPDU(pdu); err != nil {
			return err
		}

		// construct a release response pdu
		releaseRPPDU := &ReleaseRPPDU{}
		if err := releaseRPPDU.Write(assoc.conn); err != nil {
			return err
		}

		// return EOF to indicate that the association is completed
		return io.EOF
	}

	// is this an abort request? if so, simply return EOF
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
		if err := abortPDU.Write(assoc.conn); err != nil {
			return err
		}

		// let the caller know why we were not able to negotiate an association
		return ErrUnexpectedPDU
	}

	log.Printf("attempting to accept data transfer\n")

	// create a reader for the command
	commandReader, err := newPDataReader(assoc.Conn(), pdu, true)
	if err != nil {
		return err
	}

	// get the presentation context id from the reader
	pcID := commandReader.pdv.pcID

	// read the command
	command, err := readCommand(commandReader, &assoc.Assoc)
	if err != nil {
		return err
	}
	log.Printf("command is %v\n", command)

	// find the persentation context by id
	pc, err := assoc.findAcceptedPresContextByPCID(pcID)
	if err != nil {
		return err
	}
	log.Printf("pc is %v\n", pc)

	// get the command data set
	commandDataSet, err := command.asShort(CommandDataSetTypeTag, 0)
	if err != nil {
		return err
	}

	// get a data reader if required
	dataReader, err := getDataReader(commandDataSet, &assoc.Assoc, pdu)
	if err != nil {
		return err
	}

	// call the handler for the command
	if err := pc.handler.HandleRequest(&assoc.Assoc, pc, command, dataReader); err != nil {
		return err
	}

	// all is well
	return nil
}

func getDataReader(commandDataSet uint16, assoc *Assoc, pdu *PDU) (*PDataReader, error) {

	// check to see if data is present
	if isDataSetPresent(commandDataSet) {

		// create a reader for the data
		dataReader, err := newPDataReader(assoc.Conn(), pdu, false)
		if err != nil {
			return nil, err
		}

		// return the data reader
		return dataReader, nil
	}

	// return nothing, as no data reader is required
	return nil, nil
}
