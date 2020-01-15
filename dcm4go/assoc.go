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
func AcceptAssoc(conn net.Conn) (*Assoc, error) {

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

		assocACPDU, err := negotiateAssoc(assocRQPDU)
		if err != nil {
			return nil, err
		}

		if err := writeAssocACPDU(conn, assocACPDU); err != nil {
			return nil, err
		}

		return &Assoc{conn, assocRQPDU, assocACPDU}, nil
	}

	return nil, fmt.Errorf("unrecognized pdu type: %d", pdu.pduType)
}

func negotiateAssoc(assocRQPDU *AssocRQPDU) (*AssocACPDU, error) {
	return nil, fmt.Errorf("negotiateAssoc: not implemented")
}

// Close closes an association
func (assoc *Assoc) Close() error {
	return fmt.Errorf("Assoc.Close: not implemented")
}
