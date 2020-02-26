package dcm4go

import (
	"io"
)

// An AssocRQPDU represents an associate request PDU
type AssocRQPDU struct {
	*AssocACRQPDU
}

// newAssocRQPDU creates a new association request PDU
func newAssocRQPDU(calledAETitle string, callingAETitle string, capabilities *Capabilities) *AssocRQPDU {

	// add presentation context ids
	// ignore what was provided if anything
	// ids need to be odd and increasing in order
	pcs := make([]*pc, 0, 5)
	for i, capability := range capabilities.capabilities {
		pc := &pc{
			id:               byte(i*2 + 1),
			abstractSyntax:   capability.AbstractSyntax,
			transferSyntaxes: capability.TransferSyntaxes,
		}
		pcs = append(pcs, pc)
	}

	return &AssocRQPDU{
		&AssocACRQPDU{
			0x01,                      // protocol version, as per the standard
			calledAETitle,             // title of the called, as per the standard
			callingAETitle,            // title of the caller, as per the standard
			ApplicationContextNameUID, // app context name, as per the standard
			pcs,                       // pres context list
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

// readAssocRQPDU reads an associate request
func readAssocRQPDU(reader io.Reader) (*AssocRQPDU, error) {

	// read the association request
	assocACRQPDU, err := readAssocACRQPDU(reader, rqPCItemType)
	if err != nil {
		return nil, err
	}

	// construct and return an association request pdu
	return &AssocRQPDU{assocACRQPDU}, nil
}

// writeAssocRQPDU writes an associate request
func writeAssocRQPDU(writer io.Writer, assocRQPDU *AssocRQPDU) error {
	return writeAssocACRQPDU(writer, assocRQPDU.AssocACRQPDU, aAssociateRQPDU, rqPCItemType)
}
