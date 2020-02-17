package dcm4go

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// A Requestor is used to negotiate an associate request, to issue requests
// on the resulting association, and eventually, to release the association.
//
// I've been debating about whether the need for a Requestor (and Acceptor)
// is really necessary.  All of these methods could be part of Assoc (which
// represents an association).  While that is true, an association is used
// by two types of users, a requestor and an acceptor, and largely, the
// methods that those users use are different.  So, if we implemented all
// of these methods as part of Assoc, we would need to check at run time what
// type of user was calling the method, or worse, we would just let the method
// execute and allow the other side of the association to return an error.
// I thought it would better to split these users up into two different
// classes, each with their own methods, so these types of errors could be
// caught at compile time, which is always better.
type Requestor struct {
	ae    *AE      // the AE for this Requestor
	conn  net.Conn // the connection for the association
	assoc *Assoc   // the association
}

// NewRequestor creates a new Requestor.
//
// I still don't have a good idea of when to provide a constructor and
// when to allow a caller to initialize an object.  One good reason is when
// an object requires complex initialization.  Perhaps another reason is when
// one wants to hide implementation details.
//
// I've also read a couple of intersting articles about how to initialize
// objects.  One article talks about passing in initializes as a set of
// methods.  I need to find a link for that.  Another article talks about
// passing in an Opts object that contains a list of options.  It avoids
// the need for complex and long calling signatures.  I've chosen
// to go with the Opts object, not necessarily for initializaion, but for
// calling methods.
func NewRequestor(ae *AE) *Requestor {
	return &Requestor{ae: ae}
}

// AE returns the ae of this requestor
func (requestor *Requestor) AE() *AE {
	return requestor.ae
}

// Assoc returns the association of this requestor
func (requestor *Requestor) Assoc() *Assoc {
	return requestor.assoc
}

// RequestAssoc is used to send an associate request.
//
// I've been debating about whether to have the Requestor be responsible
// for managing the connection of having the connection passed to the
// Requestor.  The DICOM standard says that an association and connection
// have a one-to-one relationship, so it makes sense to manage them together.
// Software design suggests that we might want to separate those, so that
// we could substitute other types of connections, say for testing purposes.
// In the end, I've decided to have the Requestor manage the connection
// because it satisfies the standard and it makes it easier to write
// applications correctly.  If we ever want to support other types of
// connections in the future, perhaps we can initialize a Requestor with
// a connection factory.
func (requestor *Requestor) RequestAssoc(remoteAddr string, capabilities []*Capability, opts *AssocOpts) error {

	// parse the remote address
	remoteAETitle, remoteHostPort, err := requestor.parseAddr(remoteAddr)
	if err != nil {
		return err
	}

	// connect to the remote
	conn, err := net.Dial("tcp", remoteHostPort)
	if err != nil {
		return err
	}
	log.Printf("opened connection from %v to %v\n", conn.LocalAddr(), conn.RemoteAddr())

	// remember the connection
	requestor.conn = conn

	// put together an association request pdu
	assocRQPDU := newAssocRQPDU(remoteAETitle, requestor.ae.AETitle, capabilities)
	log.Printf("assocRQPDU is %v", assocRQPDU)

	// write the pdu
	if err := writeAssocRQPDU(conn, assocRQPDU); err != nil {
		return err
	}

	// read the response
	pdu, err := readPDU(conn)
	if err != nil {
		return err
	}
	log.Printf("pdu is %v\n", pdu)

	// is this an abort request?  if so, return error
	if pdu.pduType == aAbortPDU {
		log.Printf("received an abort\n")
		abortPDU, err := readAbortPDU(pdu)
		if err != nil {
			return err
		}
		return fmt.Errorf("associate request aborted, %v", abortPDU)
	}

	// if this an associate reuject?  if so, return error
	if pdu.pduType == aAssociateRJPDU {
		fmt.Printf("received a rejection\n")
		assocRJPDU, err := readAssocRJPDU(pdu)
		if err != nil {
			return err
		}
		log.Printf("assocRJPDU is %v\n", assocRJPDU)

		return fmt.Errorf("associate request rejected, %s", assocRJPDU)
	}

	// is this not an associate accept?  if not, return error
	if pdu.pduType != aAssociateACPDU {
		return fmt.Errorf("unexpected pdu type, %d", pdu.pduType)
	}

	assocACPDU, err := readAssocACPDU(pdu)
	if err != nil {
		return err
	}
	log.Printf("assocACPDU is %v\n", assocACPDU)

	// create an association from the response
	assoc := &Assoc{
		conn:       conn,
		ae:         requestor.ae,
		assocRQPDU: assocRQPDU,
		assocACPDU: assocACPDU,
	}

	// remember the association
	requestor.assoc = assoc

	// return success
	return nil
}

// parseAddr parses an address of the form 'ae@host:port' and returns the
// 'ae' and 'host:port' parts separately
func (requestor *Requestor) parseAddr(addr string) (string, string, error) {
	s := strings.Split(addr, "@")
	if len(s) != 2 {
		return "", "", fmt.Errorf("expected address of form 'ae@host:port', found '%v'", addr)
	}
	return s[0], s[1], nil
}

// SendRequest is used to send a request and receive responses.
func (requestor *Requestor) SendRequest(request *Request) ([]*Response, error) {
	return nil, fmt.Errorf("Requestor.SendRequest(): not implemented")
}

// ReleaseAssoc releases the association and closes the connection
func (requestor *Requestor) ReleaseAssoc() error {

	// release the association
	if requestor.assoc != nil {
		if err := requestor.assoc.RequestRelease(); err != nil {
			log.Printf("while releasing association, caught error %v", err)
		}
	}

	// close the connection
	if requestor.conn != nil {
		if err := requestor.conn.Close(); err != nil {
			log.Printf("while closing connection, caught error %v", err)
		}

		requestor.conn = nil
	}

	// return success
	return nil
}

// Abort aborts the association.
func (requestor *Requestor) Abort() error {
	return fmt.Errorf("Requestor.Abort(): not implemented")
}

// Echo sends a DICOM C-Echo request
func (requestor *Requestor) Echo() error {

	// create a verification request
	pc, request, err := NewCEchoRequest(requestor.assoc)
	if err != nil {
		return err
	}

	// write the verification request
	if err := requestor.assoc.writeCommand(pc, request); err != nil {
		return err
	}

	// read the response
	response, err := requestor.assoc.ReadRequestOrResponse()
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
func (requestor *Requestor) Store(reader io.Reader) error {

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
	pc, request, err := NewCStoreRequest(requestor.assoc, sopClassUID, sopInstanceUID, transferSyntaxUID)
	if err != nil {
		return err
	}

	// write the request, but no data
	if err := requestor.assoc.writeCommand(pc, request); err != nil {
		return err
	}

	// create a pdatawriter to copy the data to
	// it knows how to create pdus and pdvs as required
	// since it implements a writer, we can then simply copy the data
	pDataWriter := newPDataWriter(
		requestor.assoc.conn, // the writer is the association connection
		pc.id,                // write using the same presentation context id as in the request
		false,                // false means we are writing data
		requestor.assoc.assocRQPDU.userInfo.maxLenReceived, // the max length of each PDU written
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
	response, err := requestor.assoc.ReadRequestOrResponse()
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
