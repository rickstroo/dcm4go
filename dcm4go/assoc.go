// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains the definition and common methods of an Assoc.

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

// func (assoc *Assoc) writeCommand(pc *PresContext, command *Object) error {
// 	// write the command
// 	return assoc.writeObject(pc, command, true, ImplicitVRLittleEndianTS)
// }
//
// func (assoc *Assoc) writeData(pc *PresContext, data *Object) error {
//
// 	// find the transfer syntax
// 	ts, err := assoc.findAcceptedTransferSyntax(pc.id)
// 	if err != nil {
// 		return err
// 	}
//
// 	// write the data
// 	return assoc.writeObject(pc, data, false, ts)
// }

// // WriteResponse writes a response to the association
// func (assoc *Assoc) WriteResponse(pcID byte, command *Object, data *Object) error {
// 	message := &Message{pcID, command, data}
// 	return assoc.WriteRequestOrResponse(message)
// }

// writeCommandOnly writes a request with no data
func (assoc *Assoc) writeCommandOnly(
	presContext *PresContext,
	command *Object,
) error {
	// write the command
	if err := assoc.writeObject(presContext, command, true, ImplicitVRLittleEndianTS); err != nil {
		return err
	}
	// return success
	return nil
}

// writeObject writes an object, command or data
func (assoc *Assoc) writeObject(
	presContext *PresContext,
	object *Object,
	isCommand bool,
	transferSyntax *TransferSyntax,
) error {

	// create a writer to write the data to
	pDataWriter := newPDataWriter(assoc.conn, presContext.id, isCommand, assoc.assocRQPDU.userInfo.maxLenReceived)

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the command to the buffer
	if err := encoder.writeObject(pDataWriter, object, transferSyntax); err != nil {
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

// writeCommandWithData writes a request with data
func (assoc *Assoc) writeCommandWithData(
	presContext *PresContext,
	command *Object,
	data *Object,
) error {
	// write the command
	if err := assoc.writeObject(presContext, command, true, ImplicitVRLittleEndianTS); err != nil {
		return err
	}

	// find the transfer syntax
	transferSyntax, err := assoc.findAcceptedTransferSyntax(presContext.id)
	if err != nil {
		return err
	}

	// write the data
	if err := assoc.writeObject(presContext, data, false, transferSyntax); err != nil {
		return err
	}

	// return success
	return nil
}

// writeCommandWithDataFromReader writes a request with data from a reader
func (assoc *Assoc) writeCommandWithDataFromReader(
	presContext *PresContext,
	command *Object,
	reader io.Reader,
) error {
	// write the command
	if err := assoc.writeObject(presContext, command, true, ImplicitVRLittleEndianTS); err != nil {
		return err
	}

	// copy the data
	if err := assoc.copyDataFromReader(presContext, reader); err != nil {
		return nil
	}

	// return success
	return nil
}

// copyDataFromReader copies the data from a reader to a stream of PDUs and PDVs
func (assoc *Assoc) copyDataFromReader(
	presContext *PresContext,
	reader io.Reader,
) error {
	// create a pdatawriter to copy the data to
	// it knows how to create pdus and pdvs as required
	// since it implements a writer, we can then simply copy the data
	pDataWriter := newPDataWriter(
		assoc.conn,                               // the writer is the association connection
		presContext.id,                           // write using the same presentation context id as in the request
		false,                                    // false means we are writing data
		assoc.assocRQPDU.userInfo.maxLenReceived, // the max length of each PDU written
	)

	// copy the data
	if _, err := io.Copy(pDataWriter, reader); err != nil {
		return err
	}

	// flush the data writer
	// passing true means we are done writing this object
	if err := pDataWriter.Flush(true); err != nil {
		return err
	}

	// return success
	return nil
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
