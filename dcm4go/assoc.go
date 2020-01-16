package dcm4go

import (
	"fmt"
	"io"
	"net"
)

// Assoc represents a DICOM association
type Assoc struct {
	conn       net.Conn
	assocRQPDU *AssocRQPDU
	assocACPDU *AssocACPDU
}

// AcceptAssoc accepts an association
func AcceptAssoc(conn net.Conn, ae *AE) (*Assoc, error) {

	// read a pdu
	pdu, err := readPDU(conn)
	if err != nil {
		return nil, err
	}
	fmt.Printf("pdu is %v\n", pdu)

	// is this an association request?
	if pdu.pduType == 0x01 {
		assocRQPDU, err := readAssocRQPDU(io.LimitReader(conn, int64(pdu.pduLength)))
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

		return &Assoc{conn, assocRQPDU, assocACPDU}, nil
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

	// implement a simple acceptanace of the presentation contexts
	for _, rqPresContext := range assocRQPDU.presContexts {
		acPresContext := &ACPresContext{
			rqPresContext.id,                  // the id
			0x00,                              // success
			rqPresContext.transferSyntaxes[0], // the first one
		}
		assocACPDU.presContexts = append(assocACPDU.presContexts, acPresContext)
	}

	// return the association accept pdu
	return assocACPDU, nil
}

// Close closes an association
func (assoc *Assoc) Close() error {
	return nil
}
