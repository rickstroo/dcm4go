package dcm4go

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// Assoc represents a DICOM association
type Assoc struct {
	conn       net.Conn
	ae         *AE
	assocRQPDU *AssocRQPDU
	assocACPDU *AssocACPDU
}

// CalledAETitle returns called ae title from the association request
func (assoc *Assoc) CalledAETitle() string {
	return strings.TrimSpace(assoc.assocRQPDU.calledAETitle)
}

// CallingAETitle returns calling ae title from the association request
func (assoc *Assoc) CallingAETitle() string {
	return strings.TrimSpace(assoc.assocRQPDU.callingAETitle)
}

// AcceptAssoc accepts an association
func AcceptAssoc(conn net.Conn, ae *AE) (*Assoc, error) {

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

		assocACPDU, err := negotiateAssoc(assocRQPDU, ae)
		if err != nil {
			return nil, err
		}
		fmt.Printf("assocACPDU is %v\n", assocACPDU)

		if err := writeAssocACPDU(conn, assocACPDU); err != nil {
			return nil, err
		}

		return &Assoc{conn, ae, assocRQPDU, assocACPDU}, nil
	}

	return nil, fmt.Errorf("unrecognized pdu type: %d", pdu.pduType)
}

const (
	pcAcceptance                   = 0x00
	pcUserRejection                = 0x01
	pcNoReason                     = 0x02
	pcAbstractSyntaxNotSupported   = 0x03
	pcTransferSyntaxesNotSupported = 0x04
)

func negotiateAssoc(assocRQPDU *AssocRQPDU, ae *AE) (*AssocACPDU, error) {

	// initialize the association accept pdu
	assocACPDU := newAssocACPDU(assocRQPDU)

	// implement a simple acceptanace of the presentation contexts
	for _, rqPresContext := range assocRQPDU.presContexts {
		acPresContext, err := negotiatePresContext(rqPresContext, ae.presContexts)
		if err != nil {
			return nil, err
		}
		assocACPDU.presContexts = append(assocACPDU.presContexts, acPresContext)
	}

	// return the association accept pdu
	return assocACPDU, nil
}

func negotiatePresContext(rqPresContext *RQPresContext, spPresContexts []*SPPresContext) (*ACPresContext, error) {

	// look for a supported presentation context for this abstract syntax
	spPresContext, found := findSupportedPresContext(rqPresContext.abstractSyntax, spPresContexts)

	// if we don't find out, return a failure for this requested presentation context
	if !found {
		acPresContext := &ACPresContext{
			rqPresContext.id,             // the id
			pcAbstractSyntaxNotSupported, // failure
			""}
		return acPresContext, nil
	}

	// if we found one, now we look for a matching transfer syntax
	for _, rqTansferSyntax := range rqPresContext.transferSyntaxes {
		spTransferSyntax, found := findSupportedTransferSyntax(rqTansferSyntax, spPresContext.transferSyntaxes)
		if found {
			acPresContext := &ACPresContext{
				rqPresContext.id, // the id
				pcAcceptance,     // success
				spTransferSyntax}
			return acPresContext, nil
		}
	}

	// we didn't find a matching transfer syntax, so return a failed acceptance presentation context
	acPresContext := &ACPresContext{
		rqPresContext.id,               // the id
		pcTransferSyntaxesNotSupported, // failure
		""}
	return acPresContext, nil
}

func findSupportedPresContext(abstractSyntax string, spPresContexts []*SPPresContext) (*SPPresContext, bool) {
	for _, spPresContext := range spPresContexts {
		if abstractSyntax == spPresContext.abstractSyntax {
			return spPresContext, true
		}
	}
	return nil, false
}

func findSupportedTransferSyntax(rqTransferSyntax string, spTransferSyntaxes []string) (string, bool) {

	// compare against all the supported transfer syntaxes
	for _, spTransferSyntax := range spTransferSyntaxes {
		// if found, return the transfer syntax and true
		if rqTransferSyntax == spTransferSyntax {
			return spTransferSyntax, true
		}
	}

	// we didn't find anything
	return "", false
}

// ReadRequest reads a request from the association
func (assoc *Assoc) ReadRequest(reader io.Reader) (*Message, error) {

	// read a pdu
	pdu, err := readPDU(assoc.conn)
	if err != nil {
		return nil, err
	}
	fmt.Printf("pdu is %v\n", pdu)

	// is this an association release request?
	if pdu.pduType == aReleaseRQPDU {
		if err := readReleaseRQPDU(pdu); err != nil {
			return nil, err
		}
		if err := writeReleaseRPPDU(assoc.conn); err != nil {
			return nil, err
		}
		return nil, io.EOF
	}

	// is this an abort request?
	if pdu.pduType == aAbortPDU {
		return nil, io.EOF
	}

	// is this a data transfer request?
	if pdu.pduType == pDataTFPDU {
		message, err := readMessage(assoc.conn, assoc, pdu)
		if err != nil {
			return nil, err
		}
		return message, nil
	}

	return nil, fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
}

// HandleRequest handles a request by calling the appropriate handler
func (assoc *Assoc) HandleRequest(request *Message) (*Message, error) {

	// find the affected sop class uid
	affectedSOPClassUID, err := request.Command().asString(AffectedSOPClassUIDTag, 0)
	if err != nil {
		return nil, err
	}

	// find the handler
	handler, ok := assoc.ae.requestHandlers[affectedSOPClassUID]
	if !ok {
		return nil, fmt.Errorf("no handler found for SOP Class UID %q", affectedSOPClassUID)
	}
	if handler == nil {
		return nil, fmt.Errorf("nil handler found for SOP Class UID %q", affectedSOPClassUID)
	}

	// call the handler
	response, err := handler.HandleRequest(assoc, request)
	if err != nil {
		return nil, err
	}

	// all is well, return the response
	return response, nil
}

// WriteResponse writes a response to the association
func (assoc *Assoc) WriteResponse(writer io.Writer, message *Message) error {
	return writeMessage(assoc.conn, assoc, message)
}
