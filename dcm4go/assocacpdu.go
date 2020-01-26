package dcm4go

import (
	"io"
)

// AssocACPDU borrows from AssocACRQPDU
type AssocACPDU struct {
	AssocACRQPDU
}

// create an associate acceptance PDU from an associate request PDU
func newAssocACPDU(assocRQPDU *AssocRQPDU) *AssocACPDU {
	return &AssocACPDU{
		AssocACRQPDU{
			0x01,                       // protocol version, as per the standard
			assocRQPDU.calledAETitle,   // copy from the request, as per the standard
			assocRQPDU.callingAETitle,  // copy from the request, as per the standard
			"1.2.840.10008.3.1.1.1",    // app context name, as per the standard
			make([]*PresContext, 0, 5), // empty pres context list
			&UserInfo{
				16378,             // max length received, need to figure out why dcm4che uses this number
				"1.2.40.0.13.1.3", // implementation class uid, need to get a root, borrowing dcm4che for now
				"dcm4go-1.0",      // implementation class name
				0,                 // max num ops invoked
				0,                 // max num ops performed
			},
		},
	}
}

// readAssocACPDU reads an AssocACPDU from a reader
func readAssocACPDU(reader io.Reader) (*AssocACPDU, error) {

	// read the association request
	assocACRQPDU, err := readAssocACRQPDU(reader, acPresContextItemType)
	if err != nil {
		return nil, err
	}

	// construct and return an association request pdu
	return &AssocACPDU{*assocACRQPDU}, nil
}

func writeAssocACPDU(writer io.Writer, assocACPDU *AssocACPDU) error {
	return writeAssocACRQPDU(writer, &assocACPDU.AssocACRQPDU, acPresContextItemType)
}
