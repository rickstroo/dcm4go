package dcm4go

import (
	"io"
)

// AssocRQPDU borrows from the AssocACRQPDU
type AssocRQPDU struct {
	AssocACRQPDU
}

// newAssocRQPDU creates a new association request PDU
func newAssocRQPDU(calledAETitle string, callingAETitle string, presContexts []*PresContext) *AssocRQPDU {
	return &AssocRQPDU{
		AssocACRQPDU{
			0x01,                    // protocol version, as per the standard
			calledAETitle,           // title of the called, as per the standard
			callingAETitle,          // title of the caller, as per the standard
			"1.2.840.10008.3.1.1.1", // app context name, as per the standard
			presContexts,            // pres context list
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

// readAssocRQPDU reads an AssocRQPDU from a reader
func readAssocRQPDU(reader io.Reader) (*AssocRQPDU, error) {

	// read the association request
	assocACRQPDU, err := readAssocACRQPDU(reader, rqPresContextItemType)
	if err != nil {
		return nil, err
	}

	// construct and return an association request pdu
	return &AssocRQPDU{*assocACRQPDU}, nil
}

func writeAssocRQPDU(writer io.Writer, assocRQPDU *AssocRQPDU) error {
	return writeAssocACRQPDU(writer, &assocRQPDU.AssocACRQPDU, rqPresContextItemType)
}
