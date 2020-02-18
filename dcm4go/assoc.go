// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
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
	handlers   []Handler
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
		"conn:{local:%v,remote:%v},ae:%v,assocRQPDU:%v,assocACPDU:%v,handlers:%v",
		assoc.conn.LocalAddr(),
		assoc.conn.RemoteAddr(),
		assoc.ae,
		assoc.assocRQPDU,
		assoc.assocACPDU,
		assoc.handlers,
	)
}

// Conn returns the connection
func (assoc *Assoc) Conn() net.Conn {
	return assoc.conn
}

// AE returns the AE
func (assoc *Assoc) AE() *AE {
	return assoc.ae
}

// CalledAETitle returns called ae title from the association request
func (assoc *Assoc) CalledAETitle() string {
	return strings.TrimSpace(assoc.assocRQPDU.calledAETitle)
}

// CallingAETitle returns calling ae title from the association request
func (assoc *Assoc) CallingAETitle() string {
	return strings.TrimSpace(assoc.assocRQPDU.callingAETitle)
}

// findAcceptedPresContextByAbstractSyntax searches for a presentation context
// that was accepted for an abstract syntax and transfer syntax.
func (assoc *Assoc) findAcceptedPresContextByCapability(abstractSyntax string, transferSyntax string) (*PresContext, error) {

	// find the abstract syntax from the requested presentation contexts, there may be more than one
	for _, rqPresContext := range assoc.assocRQPDU.presContexts {
		if rqPresContext.abstractSyntax == abstractSyntax {
			// now, look for the accepted presentation context for the same pcID that was requested
			for _, acPresContext := range assoc.assocACPDU.presContexts {
				// if it's for the same id, and for the same transfer syntax id, and it was accepted
				if acPresContext.id == rqPresContext.id &&
					(transferSyntax == "*" || acPresContext.transferSyntaxes[0] == transferSyntax) &&
					acPresContext.result == pcAcceptance {
					return acPresContext, nil
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

// findAcceptedPresContextByPCID searches for a presentation context
// that was accepted for a presentation context id.
func (assoc *Assoc) findAcceptedPresContextByPCID(pcid byte) (*PresContext, error) {
	for _, acPresContext := range assoc.assocACPDU.presContexts {
		// find the accepted presentation context for the presentation context id
		if acPresContext.id == pcid && acPresContext.result == pcAcceptance {
			return acPresContext, nil
		}
	}

	// we didn't find anything
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

// ReadRequestOrResponse reads a request or response from the association
func (assoc *Assoc) ReadRequestOrResponse() (*Message, error) {

	// read a pdu
	pdu, err := readPDU(assoc.conn)
	if err != nil {
		return nil, err
	}
	log.Printf("pdu is %v", pdu)

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

// WriteRequestOrResponse writes a reqest or response from the association
func (assoc *Assoc) WriteRequestOrResponse(message *Message) error {
	return writeMessage(assoc.conn, assoc, message)
}

func (assoc *Assoc) writeCommand(pc *PresContext, command *Object) error {
	// write the command
	return assoc.writeObject(pc, command, true, ImplicitVRLittleEndianTS)
}

func (assoc *Assoc) writeData(pc *PresContext, data *Object) error {

	// find the transfer syntax
	ts, err := assoc.findAcceptedTransferSyntax(pc.id)
	if err != nil {
		return err
	}

	// write the data
	return assoc.writeObject(pc, data, false, ts)
}

func (assoc *Assoc) writeObject(pc *PresContext, object *Object, isCommand bool, ts *TransferSyntax) error {

	// create a writer to write the data to
	pDataWriter := newPDataWriter(assoc.conn, pc.id, isCommand, assoc.assocRQPDU.userInfo.maxLenReceived)

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the command to the buffer
	if err := encoder.writeObject(pDataWriter, object, ts); err != nil {
		return err
	}

	// flush to the underlying writer
	// passing true means we are done writing this object
	if err := pDataWriter.Flush(true); err != nil {
		return err
	}

	// all is well
	return nil
}

// WriteResponse writes a response to the association
func (assoc *Assoc) WriteResponse(pcID byte, command *Object, data *Object) error {
	message := &Message{pcID, command, data}
	return assoc.WriteRequestOrResponse(message)
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

// RequestAssoc is used to request an association.
func RequestAssoc(
	conn net.Conn,
	localAETitle string,
	remoteAETitle string,
	capabilities []*Capability,
	opts *AssocOpts,
) (*Assoc, error) {

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
	assoc := &Assoc{
		conn:       conn,
		assocRQPDU: assocRQPDU,
		assocACPDU: assocACPDU,
	}
	log.Printf("created association from %v to %v", assoc.CallingAETitle(), assoc.CalledAETitle())

	// return the association
	return assoc, nil
}

// ReleaseAssoc releases the association and closes the connection
func (assoc *Assoc) ReleaseAssoc() error {

	// release the association
	if assoc != nil {
		if err := assoc.RequestRelease(); err != nil {
			log.Printf("while releasing association, caught error %v", err)
		}
		log.Printf("released association from %v to %v", assoc.CallingAETitle(), assoc.CalledAETitle())
	}

	// close the connection
	if assoc.conn != nil {
		if err := assoc.conn.Close(); err != nil {
			log.Printf("while closing connection, caught error %v", err)
		}
		log.Printf("closed connection from %v to %v", assoc.conn.LocalAddr(), assoc.conn.RemoteAddr())

		assoc.conn = nil
	}

	// return success
	return nil
}

// Echo sends a DICOM C-Echo request
func (assoc *Assoc) Echo() error {

	// find the accepted presentation context for this abstract syntax and any transfer syntax
	presContex, err := assoc.findAcceptedPresContextByCapability(VerificationUID, "*")
	if err != nil {
		return err
	}

	// create a verification request
	request, err := newCEchoRequest(assoc)
	if err != nil {
		return err
	}

	// write the verification request
	if err := assoc.writeCommand(presContex, request); err != nil {
		return err
	}

	// read the response
	response, err := assoc.ReadRequestOrResponse()
	if err != nil {
		return err
	}

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

	// find the accepted presentation context for this transfer syntax
	presContex, err := assoc.findAcceptedPresContextByCapability(sopClassUID, transferSyntaxUID)
	if err != nil {
		return err
	}

	// create a group zero object
	request, err := newCStoreRequest(assoc, sopClassUID, sopInstanceUID)
	if err != nil {
		return err
	}

	// write the request, but no data
	if err := assoc.writeCommand(presContex, request); err != nil {
		return err
	}

	// create a pdatawriter to copy the data to
	// it knows how to create pdus and pdvs as required
	// since it implements a writer, we can then simply copy the data
	pDataWriter := newPDataWriter(
		assoc.conn,                               // the writer is the association connection
		presContex.id,                            // write using the same presentation context id as in the request
		false,                                    // false means we are writing data
		assoc.assocRQPDU.userInfo.maxLenReceived, // the max length of each PDU written
	)

	// copy the data
	if _, err := io.Copy(pDataWriter, reader); err != nil {
		return err
	}

	// flush the underlying writer
	// passing true means we are done writing this object
	if err := pDataWriter.Flush(true); err != nil {
		return err
	}

	// read the response
	response, err := assoc.ReadRequestOrResponse()
	if err != nil {
		return err
	}

	// get the status
	status, err := response.command.asShort(StatusTag, 0)
	if err != nil {
		return err
	}

	if status != 0 {
		return fmt.Errorf("status was %d, not success", status)
	}

	// otherwise, all is well
	return nil
}

// AcceptAssoc accepts an association
func AcceptAssoc(conn net.Conn, ae *AE, handlers []Handler) (*Assoc, error) {

	// I've decided not to implement a state machine.
	// I've looked at a number of implementations and it looks
	// to me like a state machine makes it really hard to follow
	// all the logic.  So, in the spirit of writing easy to
	// read programs, I will implement the logic of the state
	// machine in the AcceptAssoc and RequestAssoc structs.

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
	assocACPDU, assocRJPDU, err := negotiateAssoc(assocRQPDU, ae, handlers)
	if err != nil {
		return nil, err
	}
	log.Printf("assocACPDU is %v\n", assocACPDU)
	log.Printf("assocRJPDU is %v\n", assocRJPDU)

	// was association rejected
	if assocRJPDU != nil {

		// write the associate reject pdu
		if err := assocRJPDU.Write(conn); err != nil {
			return nil, err
		}
		// let the caller know that the associate request was rejected
		return nil, ErrAssociateRequestRejected
	}

	// otherwise, write the associate accept pdu
	if err := assocACPDU.Write(conn); err != nil {
		return nil, err
	}

	// construct an association
	assoc := &Assoc{
		conn:       conn,
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
func negotiateAssoc(assocRQPDU *AssocRQPDU, ae *AE, handlers []Handler) (*AssocACPDU, *AssocRJPDU, error) {

	// reject if the called ae title does not match the given ae title
	calledAETitle := strings.TrimSpace(assocRQPDU.calledAETitle)
	if calledAETitle != ae.AETitle() {
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
func (assoc *Assoc) Serve() error {

	// read a pdu
	pdu, err := readPDU(assoc.conn)
	if err != nil {
		return err
	}
	log.Printf("pdu is %v\n", pdu)

	// is this an association release request?
	/// if so, write release response and return EOF
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
	command, err := readCommand(commandReader, assoc)
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
	dataReader, err := getDataReader(commandDataSet, assoc, pdu)
	if err != nil {
		return err
	}

	// call the handler for the command
	if err := pc.handler.HandleRequest(assoc, pc, command, dataReader); err != nil {
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
