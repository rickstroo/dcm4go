package dcm4go

import (
	"fmt"
	"io"
	"net"
)

// AcceptorAssoc is an association used by acceptors of associations
type AcceptorAssoc struct {
	Assoc
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
	fmt.Printf("pdu is %v\n", pdu)

	// is this an association request?
	if pdu.pduType == aAssociateRQPDU {
		assocRQPDU, err := readAssocRQPDU(pdu)
		if err != nil {
			return nil, err
		}
		fmt.Printf("assocRQPDU is %v\n", assocRQPDU)

		assocACPDU, err := negotiateAssoc(assocRQPDU, ae, handlers)
		if err != nil {
			return nil, err
		}
		fmt.Printf("assocACPDU is %v\n", assocACPDU)

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

		return assoc, nil
	}

	return nil, fmt.Errorf("unrecognized pdu type: %d", pdu.pduType)
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

// Serve reads and services requests
func (assoc *AcceptorAssoc) Serve() error {

	for {
		// read a pdu
		pdu, err := readPDU(assoc.conn)
		if err != nil {
			return err
		}
		fmt.Printf("pdu is %v\n", pdu)

		// is this an association release request?  if so, write response and return EOF
		if pdu.pduType == aReleaseRQPDU {
			if err := readReleaseRQPDU(pdu); err != nil {
				return err
			}
			if err := writeReleaseRPPDU(assoc.conn); err != nil {
				return err
			}
			return io.EOF
		}

		// is this an abort request?  if so, just return EOF
		if pdu.pduType == aAbortPDU {
			return io.EOF
		}

		// is this not a data transfer request?
		if pdu.pduType != pDataTFPDU {
			return fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
		}

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

		// find the persentation context by id
		pc, err := assoc.findAcceptedPresContextByPCID(pcID)
		if err != nil {
			return err
		}

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

		// that's it, get another pdu
	}

	// we never get to this point
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
