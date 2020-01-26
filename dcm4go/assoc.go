package dcm4go

import (
	"fmt"
	"io"
	"net"
	"strings"
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

const (
	pcAcceptance                   = 0x00
	pcUserRejection                = 0x01
	pcNoReason                     = 0x02
	pcAbstractSyntaxNotSupported   = 0x03
	pcTransferSyntaxesNotSupported = 0x04
)

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

func findSupportedPresContext(abstractSyntax string, spPresContexts []*PresContext) (*PresContext, bool) {
	for _, spPresContext := range spPresContexts {
		if abstractSyntax == spPresContext.abstractSyntax {
			return spPresContext, true
		}
	}
	return nil, false
}

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

// ReadRequest reads a request from the association
func (assoc *Assoc) ReadRequest(reader io.Reader) (*Message, error) {

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

// WriteResponse writes a response to the association
func (assoc *Assoc) WriteResponse(writer io.Writer, message *Message) error {
	return writeMessage(assoc.conn, assoc, message)
}

// CreateFileMetaInfo creates the file meta information for a Part 10 file
func (assoc *Assoc) CreateFileMetaInfo(pcID byte, command *Object) (*Object, error) {

	// get the required information from the command
	sopClassUID, err := command.asString(AffectedSOPClassUIDTag, 0)
	if err != nil {
		return nil, err
	}
	sopInstanceUID, err := command.asString(AffectedSOPInstanceUIDTag, 0)
	if err != nil {
		return nil, err
	}

	// find the transfer syntax used to receive the object
	transferSyntax, err := findAcceptedTransferSyntax(assoc, pcID)
	if err != nil {
		return nil, err
	}

	// create the fmi
	fmi := newObject()
	fmi.addShort(FileMetaInformationVersionTag, "US", 0x0100)
	fmi.addUID(MediaStorageSOPClassUIDTag, sopClassUID)
	fmi.addUID(MediaStorageSOPInstanceUIDTag, sopInstanceUID)
	fmi.addUID(TransferSyntaxUIDTag, transferSyntax.uid)
	fmi.addUID(ImplementationClassUIDTag, "1.2.40.0.13.1.3") // borrowed from dcm4che for now
	fmi.addText(ImplementationVersionNameTag, "SH", "dcm4go")
	fmi.addText(SourceApplicationEntityTitleTag, "AE", assoc.ae.aeTitle)
	fmi.addText(SendingApplicationEntityTitleTag, "AE", assoc.CallingAETitle())
	fmi.addText(ReceivingApplicationEntityTitleTag, "AE", assoc.CalledAETitle())

	// return the file meta information
	return fmi, nil
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
	return fmt.Errorf("Assoc.RequestRelease: not implemented")
}

// RequestVerification sends a verification request
func (assoc *Assoc) RequestVerification() error {
	return fmt.Errorf("Assoc.RequestVerification: not implemented")
}
