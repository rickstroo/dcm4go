// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains the definition and common methods of an Assoc.

package dcm4go

import (
	"fmt"
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

// // WriteCommand writes a command to the association
// func (assoc *Assoc) WriteCommand(pcID byte, command *Object) error {
//
// 	// write the command, always using implicit vr little endian ts
// 	if err := assoc.writeObject(pcID, command, true, ImplicitVRLittleEndianTS); err != nil {
// 		return err
// 	}
//
// 	// return success
// 	return nil
// }
//
// // ReadCommand reads a command from the association
// func (assoc *Assoc) ReadCommand() (byte, *Object, error) {
//
// 	// get a command reader
// 	commandReader, err := assoc.getCommandReader()
// 	if err != nil {
// 		return 0, nil, err
// 	}
//
// 	// grab the presentation context id
// 	pcID := commandReader.pdv.pcID
//
// 	// create a counting reader
// 	countingReader := newCountingReader(commandReader)
//
// 	// create a decoder to read the data
// 	decoder := newDecoder(1024)
//
// 	// read the command, assuming the implicit vr little endian ts
// 	command, err := decoder.readObject(countingReader, ImplicitVRLittleEndianTS)
// 	if err != nil {
// 		return 0, nil, err
// 	}
//
// 	// return the presentation context id and command
// 	return pcID, command, nil
// }
//
// // WriteData writes a data set to the association
// func (assoc *Assoc) WriteData(pcID byte, data *Object) error {
//
// 	// find the transfer syntax
// 	transferSyntax, err := assoc.findAcceptedTransferSyntaxByPCID(pcID)
// 	if err != nil {
// 		return err
// 	}
//
// 	// write the data, using the transfer syntax negotiated for this pc
// 	if err := assoc.writeObject(pcID, data, false, transferSyntax); err != nil {
// 		return err
// 	}
//
// 	// return success
// 	return nil
// }
//
// // ReadData reads a data set from the association
// func (assoc *Assoc) ReadData(pcID byte) (*Object, error) {
//
// 	// get a data reader
// 	dataReader, err := assoc.getDataReader(pcID)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// create a counting reader
// 	countingReader := newCountingReader(dataReader)
//
// 	// create a decoder to read the data
// 	decoder := newDecoder(1024)
//
// 	// find the negotiated transfer syntax for the data
// 	transferSyntax, err := assoc.findAcceptedTransferSyntaxByPCID(pcID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("transfer syntax for request data is %v\n", transferSyntax)
//
// 	// read the data, assuming the negotiated transfer syntax
// 	data, err := decoder.readObject(countingReader, transferSyntax)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// return the data
// 	return data, nil
// }
//
// // CopyyDataTo copies the data set from the association to
// // a writer.  It is modeled after the WriteTo method of the
// // WriterTo interface.  It is used to handle the transmission
// // of large data sets without having to read the entire data set
// // into memory.
// func (assoc *Assoc) CopyyDataTo(pcID byte, writer io.Writer) (int64, error) {
//
// 	// get a data reader
// 	dataReader, err := assoc.getDataReader(pcID)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	// copy the data from the association to the writer
// 	num, err := io.Copy(writer, dataReader)
// 	if err != nil {
// 		return num, err
// 	}
//
// 	// return the number of bytes copied
// 	return num, nil
// }
//
// // CopyDataFrom copies the data set from a reader to the association, more
// // specifically, writing it to the other end of the association.  It is
// // modelled after the ReadFrom method of the ReaderFrom interface.  It is
// // used to handle the transmission of large data sets without having to
// // read the entire data set into memory.
// func (assoc *Assoc) CopyDataFrom(pcID byte, reader io.Reader) (int64, error) {
//
// 	// create a pdatawriter to write the data to
// 	// it knows how to create pdus and pdvs
// 	// since it implements a writer, we can use a copy method
// 	pdvWriter := newPDVWriter(
// 		assoc.pduWriter,                          // write to the association connection
// 		pcID,                                     // pc id for each pdv
// 		false,                                    // false means we are writing data
// 		assoc.assocRQPDU.userInfo.maxLenReceived, // max length of each pdu
// 	)
//
// 	// copy the data from the reader to the association
// 	num, err := io.Copy(pdvWriter, reader)
// 	if err != nil {
// 		return num, err
// 	}
//
// 	// flush the data writer, true means we are done writing this object
// 	if err := pdvWriter.Flush(true); err != nil {
// 		return num, err
// 	}
//
// 	// return success
// 	return num, nil
// }
//
// // writeObject writes an object, command or data
// func (assoc *Assoc) writeObject(
// 	pcID byte,
// 	object *Object,
// 	isCommand bool,
// 	transferSyntax *TransferSyntax,
// ) error {
//
// 	// create a pdatawriter to write the object to
// 	pdvWriter := newPDVWriter(
// 		assoc.conn,                               // write to the association connection
// 		pcID,                                     // pc id for each pdv
// 		isCommand,                                // is command or data
// 		assoc.assocRQPDU.userInfo.maxLenReceived, // max length of pdu
// 	)
//
// 	// create an encoder for writing objects
// 	encoder := newEncoder()
//
// 	// write the command to the buffer
// 	if err := encoder.writeObject(pdvWriter, object, transferSyntax); err != nil {
// 		return err
// 	}
//
// 	// flush the data writer, true means we are done writing this object
// 	if err := pdvWriter.Flush(true); err != nil {
// 		return err
// 	}
//
// 	// all is well
// 	return nil
// }
//
// // getComandReader gets a pdv reader for commands
// func (assoc *Assoc) getCommandReader() (*pdvReader, error) {
//
// 	// create a reader for the command
// 	commandReader, err := newPDVReader(assoc.pduReader, true)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// return the command reader
// 	return commandReader, nil
// }
//
// // getDataReader gets a pdv reader for data and validates it
// func (assoc *Assoc) getDataReader(pcID byte) (*pdvReader, error) {
//
// 	// create a reader for the data
// 	dataReader, err := newPDVReader(assoc.pduReader, false)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// check that the presentation context id that was passed, presumable
// 	// the presentation context id used to read the command, is the same
// 	// as the presentation context id being used to read the data set
// 	if pcID != dataReader.pdv.pcID {
// 		return nil, fmt.Errorf(
// 			"presentation context id for data set, %d, is different than for command, %d",
// 			pcID,
// 			dataReader.pdv.pcID,
// 		)
// 	}
//
// 	// check that we are reading data
// 	if dataReader.pdv.isCommand() {
// 		return nil, fmt.Errorf("expecting a data set PDV, but found a command PVD")
// 	}
//
// 	// return the data reader
// 	return dataReader, nil
// }
