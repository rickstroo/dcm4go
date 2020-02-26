package dcm4go

import (
	"io"
)

// An AssocRQPDU represents an associate request PDU
type assocRQPDU struct {
	*assocACRQPDU
}

// newAssocRQPDU creates a new association request PDU
func newAssocRQPDU(calledAETitle string, callingAETitle string, capabilities *Capabilities) *assocRQPDU {

	// add presentation context ids
	// ignore what was provided if anything
	// ids need to be odd and increasing in order
	pcs := make([]*pc, 0, 5)
	for i, capability := range capabilities.capabilities {
		pc := &pc{
			id:               byte(i*2 + 1),
			abstractSyntax:   capability.abstractSyntax,
			transferSyntaxes: capability.transferSyntaxes,
		}
		pcs = append(pcs, pc)
	}

	return &assocRQPDU{
		&assocACRQPDU{
			0x01,                      // protocol version, as per the standard
			calledAETitle,             // title of the called, as per the standard
			callingAETitle,            // title of the caller, as per the standard
			ApplicationContextNameUID, // app context name, as per the standard
			pcs,                       // pres context list
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

// readAssocRQPDU reads an associate request
func readAssocRQPDU(reader io.Reader) (*assocRQPDU, error) {

	// read the association request
	assocACRQPDU, err := readAssocACRQPDU(reader, rqPCItemType)
	if err != nil {
		return nil, err
	}

	// construct and return an association request pdu
	return &assocRQPDU{assocACRQPDU}, nil
}

// writeAssocRQPDU writes an associate request
func writeAssocRQPDU(writer io.Writer, assocRQPDU *assocRQPDU) error {
	return writeAssocACRQPDU(writer, assocRQPDU.assocACRQPDU, aAssociateRQPDU, rqPCItemType)
}
