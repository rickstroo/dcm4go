// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"io"
	"log"
	"net"
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

// AE returns the ae of this requestor
func (requestor *Requestor) AE() *AE {
	return requestor.ae
}

// Conn returns the connection of this requestor
func (requestor *Requestor) Conn() net.Conn {
	return requestor.conn
}

// Assoc returns the association of this requestor
func (requestor *Requestor) Assoc() *Assoc {
	return requestor.assoc
}

// I've been debating about whether to have the Requestor be responsible
// for managing the connection of having the connection passed to the
// Requestor.  The DICOM standard says that an association and connection
// have a one-to-one relationship, so it makes sense to manage them together.
// Good design suggests that we might want to separate those, so that
// we could substitute other types of connections, say for testing purposes.
// In the end, I've decided to have the Requestor manage the connection
// because it satisfies the standard and it makes it easier to write
// applications correctly.  If we ever want to support other types of
// connections in the future, perhaps we can initialize a Requestor with
// a connection factory.

// requestAssoc is used to send an associate request.
func requestAssoc(localAE *AE, remoteAE *AE, capabilities []*Capability, opts *AssocOpts) (*Requestor, error) {

	// connect to the remote
	conn, err := net.Dial("tcp", remoteAE.Host()+":"+remoteAE.Port())
	if err != nil {
		return nil, err
	}
	log.Printf("opened connection from %v to %v", conn.LocalAddr(), conn.RemoteAddr())

	// put together an association request pdu
	assocRQPDU := newAssocRQPDU(remoteAE.AETitle(), localAE.AETitle(), capabilities)
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
		ae:         localAE,
		assocRQPDU: assocRQPDU,
		assocACPDU: assocACPDU,
	}
	log.Printf("created association from %v to %v", assoc.CallingAETitle(), assoc.CalledAETitle())

	// create a requesetor
	requestor := &Requestor{
		ae:    localAE,
		conn:  conn,
		assoc: assoc,
	}

	// return the requestor
	return requestor, nil
}

// ReleaseAssoc releases the association and closes the connection
func (requestor *Requestor) ReleaseAssoc() error {

	// release the association
	if requestor.assoc != nil {
		if err := requestor.assoc.RequestRelease(); err != nil {
			log.Printf("while releasing association, caught error %v", err)
		}
		log.Printf("released association from %v to %v", requestor.assoc.CallingAETitle(), requestor.assoc.CalledAETitle())

		requestor.assoc = nil
	}

	// close the connection
	if requestor.conn != nil {
		if err := requestor.conn.Close(); err != nil {
			log.Printf("while closing connection, caught error %v", err)
		}
		log.Printf("closed connection from %v to %v", requestor.conn.LocalAddr(), requestor.conn.RemoteAddr())

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

	// find the accepted presentation context for this transfer syntax
	presContex, err := requestor.assoc.findAcceptedPresContextByAbstractSyntax(VerificationUID)
	if err != nil {
		return err
	}

	// create a verification request
	request, err := newCEchoRequest(requestor.assoc)
	if err != nil {
		return err
	}

	// write the verification request
	if err := requestor.assoc.writeCommand(presContex, request); err != nil {
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

	// find the accepted presentation context for this transfer syntax
	presContex, err := requestor.assoc.findAcceptedPresContextByAbstractSyntax(transferSyntaxUID)
	if err != nil {
		return err
	}

	// create a group zero object
	request, err := newCStoreRequest(requestor.assoc, sopClassUID, sopInstanceUID)
	if err != nil {
		return err
	}

	// write the request, but no data
	if err := requestor.assoc.writeCommand(presContex, request); err != nil {
		return err
	}

	// create a pdatawriter to copy the data to
	// it knows how to create pdus and pdvs as required
	// since it implements a writer, we can then simply copy the data
	pDataWriter := newPDataWriter(
		requestor.assoc.conn, // the writer is the association connection
		presContex.id,        // write using the same presentation context id as in the request
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
