package dcm4go

import (
	"io"
)

// An assocACPDU represents an associate accept PDU
type assocACPDU struct {
	*assocACRQPDU
}

// create an associate accept PDU from an associate request PDU
func newAssocACPDU(assocRQPDU *assocRQPDU) *assocACPDU {
	return &assocACPDU{
		&assocACRQPDU{
			0x01,                      // protocol version, as per the standard
			assocRQPDU.calledAETitle,  // copy from the request, as per the standard
			assocRQPDU.callingAETitle, // copy from the request, as per the standard
			ApplicationContextNameUID, // app context name, as per the standard
			make([]*pc, 0, 5),         // empty pres context list
			&userInfo{
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
func readAssocACPDU(reader io.Reader) (*assocACPDU, error) {

	// read the associate request
	assocACRQPDU, err := readAssocACRQPDU(reader, acPCItemType)
	if err != nil {
		return nil, err
	}

	// construct and return an association request pdu
	return &assocACPDU{assocACRQPDU}, nil
}

// Write writes an associate accept PDU
func (assocACPDU *assocACPDU) Write(writer io.Writer) error {
	return writeAssocACRQPDU(writer, assocACPDU.assocACRQPDU, aAssociateACPDU, acPCItemType)
}

func writeAssocACPDU(writer io.Writer, assocACPDU *assocACPDU) error {
	return assocACPDU.Write(writer)
}
