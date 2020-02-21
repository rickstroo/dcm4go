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

// associate negotiation results for presentation contexts
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
	pduReader  *pduReader
	pduWriter  *pduWriter
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
		message, err := readMessage(assoc, true)
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

// WriteResponse writes a response
func (assoc *Assoc) WriteResponse(presContext *PresContext, command *Object, data *Object) error {
	return assoc.writeMessage(presContext, command, data, nil)
}

// writeMessage writes a message.  a message can be a request or
// a response.  the command is required.  the data and
// reader are optional.  it's assumed that only one of the data or
// reader will be passed, but we don't enforce that.  perhaps there are
// some interesting situations where you want to write some constructed
// data, and then follow that up with data copied from a reader.
func (assoc *Assoc) writeMessage(
	presContext *PresContext,
	command *Object,
	data *Object,
	reader io.Reader,
) error {

	// write the command, always using implicit vr little endian ts
	if err := assoc.writeObject(presContext, command, true, ImplicitVRLittleEndianTS); err != nil {
		return err
	}

	// write the data if present
	if data != nil {

		// find the transfer syntax
		transferSyntax, err := assoc.findAcceptedTransferSyntaxByPCID(presContext.id)
		if err != nil {
			return err
		}

		// write the data, using the transfer syntax negotiated for this pc
		if err := assoc.writeObject(presContext, data, false, transferSyntax); err != nil {
			return err
		}
	}

	// copy data from the reader if present
	if reader != nil {

		// copy the data
		if err := assoc.copyDataFromReader(presContext, reader); err != nil {
			return err
		}
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

	// create a pdatawriter to write the object to
	pdvWriter := newPDVWriter(
		assoc.conn,                               // write to the association connection
		presContext.id,                           // pc id for each pdv
		isCommand,                                // is command or data
		assoc.assocRQPDU.userInfo.maxLenReceived, // max length of pdu
	)

	// create an encoder for writing objects
	encoder := newEncoder()

	// write the command to the buffer
	if err := encoder.writeObject(pdvWriter, object, transferSyntax); err != nil {
		return err
	}

	// flush the data writer, true means we are done writing this object
	if err := pdvWriter.Flush(true); err != nil {
		return err
	}

	// all is well
	return nil
}

// copyDataFromReader copies the data from a reader to a stream of PDUs and PDVs
func (assoc *Assoc) copyDataFromReader(
	presContext *PresContext,
	reader io.Reader,
) error {

	// create a pdatawriter to write the data to
	// it knows how to create pdus and
	// since it implements a writer, we can use a copy method
	pdvWriter := newPDVWriter(
		assoc.conn,                               // write to the association connection
		presContext.id,                           // pc id for each pdv
		false,                                    // false means we are writing data
		assoc.assocRQPDU.userInfo.maxLenReceived, // max length of each pdu
	)

	// copy the data
	if _, err := io.Copy(pdvWriter, reader); err != nil {
		return err
	}

	// flush the data writer, true means we are done writing this object
	if err := pdvWriter.Flush(true); err != nil {
		return err
	}

	// return success
	return nil
}

// // readMessage reads a message.  a message can be a request or a response.
// func (assoc *Assoc) readMessage(
// 	presContext *PresContext,
// 	readData bool, // if true, and data present, read the data into an object
// ) (
// 	*Object, // command
// 	*Object, // data read into an object
// 	*PDataReader, // data available from a reader
// 	error,
// ) {
//
// 		// create a reader for the command
// 		commandReader, err := newPDataReader(reader, pdu, true)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		// get the presentation context id from the reader
// 		pcID := commandReader.pdv.pcID
//
// 		// read the command from the pdu
// 		command, err := readCommand(commandReader, assoc)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		// get the command data set
// 		commandDataSet, err := command.asShort(CommandDataSetTypeTag, 0)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		// create a message
// 		message := &Message{
// 			pcID:    pcID,
// 			command: command,
// 		}
//
// 		if isDataSetPresent(commandDataSet) {
//
// 			// create a reader for the data
// 			pDataReader, err := newPDataReader(reader, pdu, false)
// 			if err != nil {
// 				return nil, err
// 			}
//
// 			// should we read the data, or pass the data reader back to the caller?
// 			if shouldReadData {
//
// 				// read the data
// 				data, err := readData(pDataReader, assoc, pcID)
// 				if err != nil {
// 					return nil, err
// 				}
//
// 				// add the data to the message
// 				message.data = data
//
// 			} else {
//
// 				// otherwise, add the data reader to the message
// 				message.pDataReader = pDataReader
// 			}
// 		}
//
// 		// return the message
// 		return message, nil
// 	return nil, nil, nil, fmt.Errorf("Assoc.readMessage(): not implemented")
// }
