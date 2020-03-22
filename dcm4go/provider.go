package dcm4go

import (
	"bytes"
	"fmt"
	"log"
)

type serviceProvider struct {
	aeTitle      string
	capabilities *Capabilities
}

func (sp *serviceProvider) onAssociateRQ(p *pdu) (*pdu, error) {
	pr := bytes.NewReader(p.buf)

	// read the associate request
	assocRQPDU, err := readAssocRQPDU(pr)
	if err != nil {
		return nil, err
	}
	log.Printf("assocRQPDU is %v\n", assocRQPDU)

	// attempt to negotiate an association
	assocACPDU, assocRJPDU, err := negotiateAssoc(assocRQPDU, sp.aeTitle, sp.capabilities)
	if err != nil {
		return nil, err
	}
	if assocACPDU != nil {
		log.Printf("accepted associate request, assocACPDU is %v\n", assocACPDU)
		pdu, err := createAssocACPDU(assocACPDU)
		if err != nil {
			return nil, err
		}
		return pdu, nil
	}
	if assocRJPDU != nil {
		log.Printf("rejected associate request, assocRJPDU is %v\n", assocRJPDU)
		pdu, err := createAssocRJPDU(assocRJPDU)
		if err != nil {
			return nil, err
		}
		return pdu, nil
	}

	return nil, fmt.Errorf("didn't accept or reject association request, hmm, that's weird")
}
