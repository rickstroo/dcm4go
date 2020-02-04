package dcm4go

import (
	"fmt"
	"net"
)

// AcceptorAssoc is an association used by acceptors of associations
type AcceptorAssoc struct {
	Assoc
}

// AcceptAssoc accepts an association
func AcceptAssoc(conn net.Conn, ae *AE, capabilities []*Capability) (*AcceptorAssoc, error) {

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

		assocACPDU, err := negotiateAssoc(assocRQPDU, ae, capabilities)
		if err != nil {
			return nil, err
		}
		fmt.Printf("assocACPDU is %v\n", assocACPDU)

		if err := writeAssocACPDU(conn, assocACPDU); err != nil {
			return nil, err
		}

		assoc := &AcceptorAssoc{
			Assoc{
				conn,
				ae,
				assocRQPDU,
				assocACPDU,
			},
		}

		return assoc, nil
	}

	return nil, fmt.Errorf("unrecognized pdu type: %d", pdu.pduType)
}

// negotiateAssoc determines what requested presentation contexts
// are accepted based on the presentation contexts that are supported
// by the ae
func negotiateAssoc(assocRQPDU *AssocRQPDU, ae *AE, capabilities []*Capability) (*AssocACPDU, error) {

	// initialize the association accept pdu
	assocACPDU := newAssocACPDU(assocRQPDU)

	// negotiate each of the presentation contexts
	for _, rqPresContext := range assocRQPDU.presContexts {
		acPresContext, err := negotiatePresContext(rqPresContext, capabilities)
		if err != nil {
			return nil, err
		}
		assocACPDU.presContexts = append(assocACPDU.presContexts, acPresContext)
	}

	return assocACPDU, nil
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

// ReadRequest reads a request from the association
func (assoc *AcceptorAssoc) ReadRequest() (*Message, error) {
	return assoc.ReadRequestOrResponse()
}

// WriteResponse writes a response to the association
func (assoc *AcceptorAssoc) WriteResponse(message *Message) error {
	return assoc.WriteRequestOrResponse(message)
}
