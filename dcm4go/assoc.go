package dcm4go

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// associate negotiation results
const (
	pcAcceptance                   = 0x00
	pcUserRejection                = 0x01
	pcNoReason                     = 0x02
	pcAbstractSyntaxNotSupported   = 0x03
	pcTransferSyntaxesNotSupported = 0x04
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

// negotiateAssoc determines what requested presentation contexts
// are accepted based on the presentation contexts that are supported
// by the ae
func negotiateAssoc(assocRQPDU *AssocRQPDU, ae *AE) (*AssocACPDU, error) {

	// initialize the association accept pdu
	assocACPDU := newAssocACPDU(assocRQPDU)

	// negotiate each of the presentation contexts
	for _, rqPresContext := range assocRQPDU.presContexts {
		acPresContext, err := negotiatePresContext(rqPresContext, ae.presContexts)
		if err != nil {
			return nil, err
		}
		assocACPDU.presContexts = append(assocACPDU.presContexts, acPresContext)
	}

	return assocACPDU, nil
}

// negotiationPresContext negotiates a single presentation context
func negotiatePresContext(rqPresContext *PresContext, spPresContexts []*PresContext) (*PresContext, error) {

	// look for a supported presentation context for this abstract syntax
	spPresContext, found := findSupportedPresContext(rqPresContext.abstractSyntax, spPresContexts)

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
		spTransferSyntax, found := findSupportedTransferSyntax(rqTansferSyntax, spPresContext.transferSyntaxes)
		if found {
			acPresContext := &PresContext{
				rqPresContext.id,           // the id
				"",                         // no abstract syntax
				[]string{spTransferSyntax}, // the transfer syntax
				pcAcceptance,               // success
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

// findSupportedPresContext searches for a supported presentation context
// for an abstract syntax
func findSupportedPresContext(abstractSyntax string, spPresContexts []*PresContext) (*PresContext, bool) {
	for _, spPresContext := range spPresContexts {
		if abstractSyntax == spPresContext.abstractSyntax {
			return spPresContext, true
		}
	}
	return nil, false
}

// findSupportedTransferSyntax looks for a supported transfer syntax
// that matches the requested transfer syntax
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

// findAcceptedPresContextByAbstractSyntax searches for a presentation context
// that was accepted for an abstract syntax.
func (assoc *Assoc) findAcceptedPresContextByAbstractSyntax(abstractSyntax string) (*PresContext, error) {

	// find the abstract syntax from the requested presentation contexts
	for _, rqPresContext := range assoc.assocRQPDU.presContexts {
		if rqPresContext.abstractSyntax == abstractSyntax {
			// now, look for the accepted presentation context for the same pcID that was requested
			for _, acPresContext := range assoc.assocACPDU.presContexts {
				if rqPresContext.id == acPresContext.id {
					// and make sure it was accepted
					if acPresContext.result == pcAcceptance {
						return acPresContext, nil
					}
				}
			}
		}
	}

	// we didn't find anything
	return nil, fmt.Errorf("unable to find accepted presentation context for abstract syntax %q", abstractSyntax)
}

// findAcceptedPresContextByPCID searches for a presentation context
// that was accepted for a presentation context id.
func (assoc *Assoc) findAcceptedPresContextByPCID(pcid byte) (*PresContext, error) {
	for _, acPresContext := range assoc.assocACPDU.presContexts {
		// find the accepted presentation context for the presentation context id
		if acPresContext.id == pcid {
			if acPresContext.result == pcAcceptance {
				return acPresContext, nil
			}
		}
	}
	return nil, fmt.Errorf("unable to find accepted presentation context for presentation context id %d", pcid)
}

// findAcceptedTransferSyntaxByPCID finds the transfer syntax for the presentation
// context that was accepted for a presentation context id
func (assoc *Assoc) findAcceptedTransferSyntaxByPCID(pcid byte) (*TransferSyntax, error) {
	presContext, err := assoc.findAcceptedPresContextByPCID(pcid)
	if err != nil {
		return nil, err
	}
	transferSyntax, err := findTransferSyntax(presContext.transferSyntaxes[0])
	if err != nil {
		return nil, err
	}
	return transferSyntax, nil
}

// ReadRequestOrResponse reads a request from the association
func (assoc *Assoc) ReadRequestOrResponse() (*Message, error) {

	// read a pdu
	pdu, err := readPDU(assoc.conn)
	if err != nil {
		return nil, err
	}
	fmt.Printf("pdu is %v\n", pdu)

	// is this an association release request?  if so, write response and return EOF
	if pdu.pduType == aReleaseRQPDU {
		if err := readReleaseRQPDU(pdu); err != nil {
			return nil, err
		}
		if err := writeReleaseRPPDU(assoc.conn); err != nil {
			return nil, err
		}
		return nil, io.EOF
	}

	// is this an abort request?  if so, just return EOF
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

// ReadRequest reads a request from the association
func (assoc *Assoc) ReadRequest() (*Message, error) {
	return assoc.ReadRequestOrResponse()
}

// ReadResponse reads a response from the association
func (assoc *Assoc) ReadResponse() (*Message, error) {
	return assoc.ReadRequestOrResponse()
}

// WriteResponse writes a response to the association
func (assoc *Assoc) WriteResponse(message *Message) error {
	return writeMessage(assoc.conn, assoc, message)
}

// WriteRequest writes a request to the association
func (assoc *Assoc) WriteRequest(message *Message) error {
	return writeMessage(assoc.conn, assoc, message)
}

// RequestAssoc requests an association
func RequestAssoc(conn net.Conn, ae *AE, calledAETitle string) (*Assoc, error) {

	// put together an association request pdu
	assocRQPDU := newAssocRQPDU(calledAETitle, ae.aeTitle, ae.presContexts)
	fmt.Printf("assocRQPDU is %v", assocRQPDU)

	// write the pdu
	if err := writeAssocRQPDU(conn, assocRQPDU); err != nil {
		return nil, err
	}

	// read the response
	pdu, err := readPDU(conn)
	if err != nil {
		return nil, err
	}
	fmt.Printf("pdu is %v\n", pdu)

	// is this an association accept?
	if pdu.pduType == aAssociateACPDU {
		assocACPDU, err := readAssocACPDU(pdu)
		if err != nil {
			return nil, err
		}
		fmt.Printf("assocACPDU is %v\n", assocACPDU)

		// create an association from the response
		assoc := &Assoc{
			conn,
			ae,
			assocRQPDU,
			assocACPDU,
		}

		// return assoc
		return assoc, nil
	}

	if pdu.pduType == aAssociateRJPDU {
		fmt.Printf("received a rejection\n")
		assocRJPDU, err := readAssocRJPDU(pdu)
		if err != nil {
			return nil, err
		}
		fmt.Printf("assocRJPDU is %v\n", assocRJPDU)

		return nil, fmt.Errorf("associate request rejected, %s", assocRJPDU)
	}

	// is this an abort request?  if so, just return EOF
	if pdu.pduType == aAbortPDU {
		fmt.Printf("received an abort\n")
		return nil, io.EOF
	}

	return nil, fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
}

// RequestRelease requests release from an association
func (assoc *Assoc) RequestRelease() error {

	// write a request release pdu
	if err := writeReleaseRQPDU(assoc.conn); err != nil {
		return err
	}
	fmt.Printf("wrote a release request\n")

	// read the response
	pdu, err := readPDU(assoc.conn)
	if err != nil {
		return err
	}
	fmt.Printf("pdu is %v\n", pdu)

	if pdu.pduType == aReleaseRPPDU {
		fmt.Printf("received a release response\n")
		if err := readReleaseRPPDU(pdu); err != nil {
			return err
		}

		// all is well
		return nil
	}

	// is this an abort request?  if so, just return EOF
	if pdu.pduType == aAbortPDU {
		fmt.Printf("received an abort\n")
		// all is well
		return nil
	}

	return fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
}

// Verify sends a verification request
func (assoc *Assoc) Verify() error {

	// create a verification request
	request, err := NewCEchoRequest(assoc)
	if err != nil {
		return err
	}
	fmt.Printf("c-echo request is %v\n", request)

	// write the verification request
	if err := assoc.WriteRequest(request); err != nil {
		return err
	}

	// read the response
	response, err := assoc.ReadResponse()
	if err != nil {
		return err
	}
	fmt.Printf("c-echo response is %v\n", response)

	// get the status
	status, err := response.command.asShort(StatusTag, 0)
	if err != nil {
		return err
	}

	if status != 0 {
		return fmt.Errorf("status was %d, not success", status)
	}

	return nil
}

// Send sends a store request
func (assoc *Assoc) Send(reader io.Reader) error {

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
	fmt.Printf("sop class uid is %q\n", sopClassUID)

	// get the sop class instance UID of the stored object
	sopInstanceUID, err := groupTwo.AsString(MediaStorageSOPInstanceUIDTag, 0)
	if err != nil {
		return err
	}
	fmt.Printf("sop class uid is %q\n", sopClassUID)

	// get the transfer syntax used to store the file
	transferSyntaxUID, err := groupTwo.AsString(TransferSyntaxUIDTag, 0)
	if err != nil {
		return err
	}
	fmt.Printf("transfer syntax uid is %q\n", transferSyntaxUID)

	// create a group zero object
	request, err := NewCStoreRequest(assoc, sopClassUID, sopInstanceUID, transferSyntaxUID)
	if err != nil {
		return err
	}

	// write the request, but no data
	if err := assoc.WriteRequest(request); err != nil {
		return err
	}

	// grab the pc id of the request
	pcID := request.pcID

	// create a pdatawriter to copy the data to
	// it knows how to create pdus and pdvs as required
	// since it implements a writer, we can then simply copy the data
	pDataWriter := newPDataWriter(assoc.conn, pcID, false, assoc.assocRQPDU.userInfo.maxLenReceived)

	// copy the data
	num, err := io.Copy(pDataWriter, reader)
	if err != nil {
		return err
	}
	fmt.Printf("copied %d bytes\n", num)

	// flush to the underlying writer
	// passing true means we are done writing this object
	if err := pDataWriter.Flush(true); err != nil {
		return err
	}

	// all is well
	return nil
}
