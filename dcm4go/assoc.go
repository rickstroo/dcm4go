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
	state      int
	handlers   []Handler
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

// // the states of the DICOM state machine
// const (
// 	sta1  = 1  // idle
// 	sta2  = 2  // transport connection open (awaiting A-ASSOCIATE-RQ PDU)
// 	sta3  = 3  // awaiting A-ASSOCIATE response primitive from local user
// 	sta4  = 4  // awaiting transport connection opening to complete (from local transport service)
// 	sta5  = 5  // awaiting A-ASSOCIATE-AC or A-ASSOCIATE-RJ PDU
// 	sta6  = 6  // association established and ready for data transfer
// 	sta7  = 7  // awaiting A-RELEASE-RP PDU
// 	sta8  = 8  // awaiting A-RELEASE response primitive (from local user)
// 	sta9  = 9  // release collision requestor side, awaiting A-RELEASE primitive response (from local user)
// 	sta10 = 10 // release collision acceptor side, awaiting A-RELEASE-RP PDU
// 	sta11 = 11 // release collision requestor side, awaiting A-RELEASE-RP PDU
// 	sta12 = 12 // release collision acceptor side, awaiting A-RELEASE response primitive (from local user)
// 	sta13 = 13 // awaiting connection to close (association no longer exists)
// )
//
// // Serve reads and services requests
// func (assoc *Assoc) Serve(conn net.Conn, ae *AE, handlers []Handler) error {
//
// 	// remember the connection
// 	assoc.conn = conn
//
// 	// remember the ae
// 	assoc.ae = ae
//
// 	// remember the handlers
// 	assoc.handlers = handlers
//
// 	// set the state for an acceptor of requests
// 	assoc.state = sta2
//
// 	// call the state machine
// 	return assoc.activate()
// }
//
// func (assoc *Assoc) activate() error {
// 	for {
// 		_, err := assoc.nextPDU()
// 		if err != nil {
// 			return err
// 		}
// 		if assoc.state == sta1 || assoc.state == sta13 {
// 			break
// 		}
// 	}
// 	return nil
// }
//
// func (assoc *Assoc) nextPDU() (*PDU, error) {
// 	// read a pdu
// 	pdu, err := readPDU(assoc.conn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("pdu is %v\n", pdu)
//
// 	switch assoc.state {
// 	case sta1:
// 	case sta2:
// 		if err := assoc.sta2(pdu); err != nil {
// 			return nil, err
// 		}
// 	case sta3:
// 	case sta4:
// 	case sta5:
// 	case sta6:
// 	case sta7:
// 	case sta8:
// 	case sta9:
// 	case sta10:
// 	case sta11:
// 	case sta12:
// 	case sta13:
// 	}
//
// 	return pdu, nil
// }
//
// func (assoc *Assoc) sta1(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta1(): not implemented")
// }
//
// func (assoc *Assoc) sta2(pdu *PDU) error {
// 	switch pdu.pduType {
// 	case aAssociateACPDU, aAssociateRJPDU, pDataTFPDU, aReleaseRQPDU, aReleaseRPPDU:
// 		if err := assoc.aa1(); err != nil {
// 			return err
// 		}
// 	case aAssociateRQPDU:
// 		if err := assoc.ae6(pdu); err != nil {
// 			return err
// 		}
// 	case aAbortPDU:
// 		if err := assoc.aa2(); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
//
// func (assoc *Assoc) sta3(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta3(): not implemented")
// }
//
// func (assoc *Assoc) sta4(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta4(): not implemented")
// }
//
// func (assoc *Assoc) sta5(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta5(): not implemented")
// }
//
// func (assoc *Assoc) sta6(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta6(): not implemented")
// }
//
// func (assoc *Assoc) sta7(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta7(): not implemented")
// }
//
// func (assoc *Assoc) sta8(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta8(): not implemented")
// }
//
// func (assoc *Assoc) sta9(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta9(): not implemented")
// }
//
// func (assoc *Assoc) sta10(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta10(): not implemented")
// }
//
// func (assoc *Assoc) sta11(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta11(): not implemented")
// }
//
// func (assoc *Assoc) sta12(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta12(): not implemented")
// }
//
// func (assoc *Assoc) sta13(pdu *PDU) error {
// 	return fmt.Errorf("Assoc.sta13(): not implemented")
// }
//
// func (assoc *Assoc) aa1() error {
// 	// send abort
// 	// start or restart artim timer
// 	// go to state sta13
// 	assoc.state = sta13
//
// 	// all is well
// 	return nil
// }
//
// func (assoc *Assoc) aa2() error {
// 	// stop the artim timer
// 	// close the connection
// 	// go sta1
// 	assoc.state = sta1
// 	return nil
// }
//
// func (assoc *Assoc) ae6(pdu *PDU) error {
// 	// stop artim timer
// 	// if a-associate-rq acceptable, issue a-associate-ac and go to state sta3
// 	// otherwise, issue a-associate-rj, start timer and go to sta13
//
// 	// read the A-ASSOCIATE-RQ PDU
// 	assocRQPDU, err := readAssocRQPDU(pdu)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("assocRQPDU is %v\n", assocRQPDU)
//
// 	assocACPDU, err := negotiateAssoc(assocRQPDU, assoc.ae, assoc.handlers)
// 	if err != nil {
// 		// need to write an A-ASSOCIATE-RJ here
// 		return err
// 	}
// 	fmt.Printf("assocACPDU is %v\n", assocACPDU)
//
// 	if err := writeAssocACPDU(assoc.conn, assocACPDU); err != nil {
// 		return err
// 	}
//
// 	// remember the associate request and response PDUs
// 	assoc.assocRQPDU = assocRQPDU
// 	assoc.assocACPDU = assocACPDU
//
// 	// go to the next state
// 	assoc.state = sta3
//
// 	// all is well
// 	return nil
// }
