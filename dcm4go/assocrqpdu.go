package dcm4go

import (
	"io"
)

// An AssocRQPDU represents an associate request PDU
type AssocRQPDU struct {
	*AssocACRQPDU
}

// newAssocRQPDU creates a new association request PDU
func newAssocRQPDU(calledAETitle string, callingAETitle string, presContexts []*PresContext) *AssocRQPDU {
	return &AssocRQPDU{
		&AssocACRQPDU{
			0x01,                      // protocol version, as per the standard
			calledAETitle,             // title of the called, as per the standard
			callingAETitle,            // title of the caller, as per the standard
			ApplicationContextNameUID, // app context name, as per the standard
			presContexts,              // pres context list
			&UserInfo{
				16378,                     // max length received, need to figure out why dcm4che uses this number
				ImplementationClassUID,    // implementation class uid, need to get a root, borrowing dcm4che for now
				ImplementationVersionName, // implementation class name
				0,                         // max num ops invoked
				0,                         // max num ops performed
			},
		},
	}
}

// readAssocRQPDU reads an associate request
func readAssocRQPDU(reader io.Reader) (*AssocRQPDU, error) {

	// read the association request
	assocACRQPDU, err := readAssocACRQPDU(reader, rqPresContextItemType)
	if err != nil {
		return nil, err
	}

	// construct and return an association request pdu
	return &AssocRQPDU{assocACRQPDU}, nil
}

// writeAssocRQPDU writes an associate request
func writeAssocRQPDU(writer io.Writer, assocRQPDU *AssocRQPDU) error {
	return writeAssocACRQPDU(writer, assocRQPDU.AssocACRQPDU, aAssociateRQPDU, rqPresContextItemType)
}
