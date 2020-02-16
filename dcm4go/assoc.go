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
	state      int
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
		"conn:{local:%v,remote:%v},ae:%v,assocRQPDU:%v,assocACPDU:%v,state:%v,handlers:%v",
		assoc.conn.LocalAddr(),
		assoc.conn.RemoteAddr(),
		assoc.ae,
		assoc.assocRQPDU,
		assoc.assocACPDU,
		assoc.state,
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
// that was accepted for an abstract syntax.
func (assoc *Assoc) findAcceptedPresContextByAbstractSyntax(abstractSyntax string) (*PresContext, error) {

	// find the abstract syntax from the requested presentation contexts, there may be more than one
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
			// and make sure it was accepted
			if acPresContext.result == pcAcceptance {
				return acPresContext, nil
			}
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
	fmt.Printf("pdu is %v\n", pdu)

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

// Close closes the connection of the association
func (assoc *Assoc) Close() error {
	return assoc.conn.Close()
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
		log.Printf("received an abort\n")
		// all is well
		return nil
	}

	return fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
}

// Echo sends a DICOM C-Echo request
func (assoc *Assoc) Echo() error {

	// create a verification request
	pc, request, err := NewCEchoRequest(assoc)
	if err != nil {
		return err
	}

	// write the verification request
	if err := assoc.writeCommand(pc, request); err != nil {
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

	// create a group zero object
	pc, request, err := NewCStoreRequest(assoc, sopClassUID, sopInstanceUID, transferSyntaxUID)
	if err != nil {
		return err
	}

	// write the request, but no data
	if err := assoc.writeCommand(pc, request); err != nil {
		return err
	}

	// create a pdatawriter to copy the data to
	// it knows how to create pdus and pdvs as required
	// since it implements a writer, we can then simply copy the data
	pDataWriter := newPDataWriter(
		assoc.conn,                               // the writer is the association connection
		pc.id,                                    // write using the same presentation context id as in the request
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
