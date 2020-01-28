package dcm4go

import (
	"io"
)

// newAssocRQPDU creates a new association request PDU
func newAssocRQPDU(calledAETitle string, callingAETitle string, presContexts []*PresContext) *AssocACRQPDU {
	return &AssocACRQPDU{
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
	}
}

// readAssocRQPDU reads an associate request
func readAssocRQPDU(reader io.Reader) (*AssocACRQPDU, error) {

	// read the association request
	assocRQPDU, err := readAssocACRQPDU(reader, rqPresContextItemType)
	if err != nil {
		return nil, err
	}

	// construct and return an association request pdu
	return assocRQPDU, nil
}

// writeAssocRQPDU writes an associate request
func writeAssocRQPDU(writer io.Writer, assocRQPDU *AssocACRQPDU) error {
	return writeAssocACRQPDU(writer, assocRQPDU, aAssociateRQPDU, rqPresContextItemType)
}
