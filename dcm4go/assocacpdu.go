package dcm4go

import (
	"io"
)

// An AssocACPDU represents an associate accept PDU
type AssocACPDU struct {
	*AssocACRQPDU
}

// create an associate accept PDU from an associate request PDU
func newAssocACPDU(assocRQPDU *AssocRQPDU) *AssocACPDU {
	return &AssocACPDU{
		&AssocACRQPDU{
			0x01,                      // protocol version, as per the standard
			assocRQPDU.calledAETitle,  // copy from the request, as per the standard
			assocRQPDU.callingAETitle, // copy from the request, as per the standard
			ApplicationContextNameUID, // app context name, as per the standard
			make([]*PC, 0, 5),         // empty pres context list
			&UserInfo{
				16378,                     // max length received, need to figure out why dcm4che uses this number
				ImplementationClassUID,    // implementation class uid, we have our own now
				ImplementationVersionName, // implementation version name
				0,                         // max num ops invoked
				0,                         // max num ops performed
			},
		},
	}
}

// readAssocACPDU reads an associate accept
func readAssocACPDU(reader io.Reader) (*AssocACPDU, error) {

	// read the associate request
	assocACRQPDU, err := readAssocACRQPDU(reader, acPCItemType)
	if err != nil {
		return nil, err
	}

	// construct and return an association request pdu
	return &AssocACPDU{assocACRQPDU}, nil
}

// Write writes an associate accept PDU
func (assocACPDU *AssocACPDU) Write(writer io.Writer) error {
	return writeAssocACRQPDU(writer, assocACPDU.AssocACRQPDU, aAssociateACPDU, acPCItemType)
}

func writeAssocACPDU(writer io.Writer, assocACPDU *AssocACPDU) error {
	return assocACPDU.Write(writer)
}
