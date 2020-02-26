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

// Assoc represents a DICOM association
type Assoc struct {
	conn       net.Conn
	pduReader  *pduReader
	pduWriter  *pduWriter
	ae         *AE
	assocRQPDU *assocRQPDU
	assocACPDU *assocACPDU
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
		"conn:{local:%v,remote:%v},ae:%v,assocRQPDU:%v,assocACPDU:%v",
		assoc.conn.LocalAddr(),
		assoc.conn.RemoteAddr(),
		assoc.ae,
		assoc.assocRQPDU,
		assoc.assocACPDU,
	)
}

// Conn returns the connection
func (assoc *Assoc) Conn() net.Conn {
	return assoc.conn
}

// CalledAETitle returns called ae title from the association request
func (assoc *Assoc) CalledAETitle() string {
	return strings.TrimSpace(assoc.assocRQPDU.calledAETitle)
}

// CallingAETitle returns calling ae title from the association request
func (assoc *Assoc) CallingAETitle() string {
	return strings.TrimSpace(assoc.assocRQPDU.callingAETitle)
}

// findAcceptedPCByCapability searches for a presentation context
// that was accepted for an abstract syntax and transfer syntax.
func (assoc *Assoc) findAcceptedPCByCapability(abstractSyntax string, transferSyntax string) (*pc, error) {

	// find the abstract syntax from the requested presentation contexts, there may be more than one
	for _, rqpc := range assoc.assocRQPDU.pcs {
		if rqpc.abstractSyntax == abstractSyntax {
			// now, look for the accepted presentation context for the same pcID that was requested
			for _, acpc := range assoc.assocACPDU.pcs {
				// if it's for the same id, and for the same transfer syntax id, and it was accepted
				if acpc.id == rqpc.id &&
					(transferSyntax == "*" || acpc.transferSyntaxes[0] == transferSyntax) &&
					acpc.result == pcAcceptance {
					return acpc, nil
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

// findAcceptedPCByPCID searches for a presentation context
// that was accepted for a presentation context id.
func (assoc *Assoc) findAcceptedPCByPCID(pcid byte) (*pc, error) {
	for _, pc := range assoc.assocACPDU.pcs {
		// find the accepted presentation context for the presentation context id
		if pc.id == pcid && pc.result == pcAcceptance {
			return pc, nil
		}
	}

	// we didn't find anything
	return nil, fmt.Errorf("unable to find accepted presentation context for presentation context id %d", pcid)
}

// findAcceptedTransferSyntaxByPCID finds the transfer syntax for the presentation
// context that was accepted for a presentation context id
func (assoc *Assoc) findAcceptedTransferSyntaxByPCID(pcid byte) (*TransferSyntax, error) {
	pc, err := assoc.findAcceptedPCByPCID(pcid)
	if err != nil {
		return nil, err
	}
	transferSyntax, err := findTransferSyntax(pc.transferSyntaxes[0])
	if err != nil {
		return nil, err
	}
	return transferSyntax, nil
}

func (assoc *Assoc) writeMessage(message *Message) error {
	return writeMessage(assoc, message)
}

func (assoc *Assoc) readMessage(shouldReadData bool) (*Message, error) {
	return readMessage(assoc, shouldReadData)
}

func onAbort(reader io.Reader) error {
	abortPDU, err := readAbortPDU(reader)
	if err != nil {
		return err
	}
	log.Printf("received an abort pdu, %v", abortPDU)
	return fmt.Errorf("associate request aborted, %w", ErrAssociationAborted)
}

func onUnexpectedPDU(reader io.Reader, pdu *pdu) error {
	log.Printf("received unexpected pdu type, %v", pdu)
	return fmt.Errorf("unexpected pdu type, %d, %w", pdu.pduType, ErrUnexpectedPDU)
}
