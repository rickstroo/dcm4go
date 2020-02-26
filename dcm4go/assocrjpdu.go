package dcm4go

import (
	"fmt"
	"io"
)

const (
	resultRejectedPermanent = 0x01
	resultRejectedTransient = 0x02
)

const (
	sourceServiceUser                                = 0x01
	sourceServiceProviderACSERelatedFunction         = 0x02
	sourceServiceProviderPresentationRelatedFunction = 0x04
)

const (
	reasonServiceUserNoReasonGiven                      = 0x01
	reasonServiceUserApplicationContextNameNotSupported = 0x02
	reasonServiceUserCallingAETitleNotRecognized        = 0x03
	reasonServiceUserCalledAETitleNotRecognized         = 0x07

	reasonServiceProviderACSERelatedFunctionNoReasonGiven               = 0x01
	reasonServiceProviderACSERelatedFunctionProtocolVersionNotSupported = 0x02

	reasonServiceProviderPresentationRelatedFunctionTemporaryCongestion = 0x01
	reasonServiceProviderPresentationRelatedFunctionLocalLimitExceeded  = 0x02
)

// assocRJPDU is an associate reject PDU
type assocRJPDU struct {
	result byte
	source byte
	reason byte
}

func (assocRJPDU *assocRJPDU) String() string {
	return fmt.Sprintf(
		"{result:%d,source:%d,reason:%d}",
		assocRJPDU.result,
		assocRJPDU.source,
		assocRJPDU.reason,
	)
}

func readAssocRJPDU(reader io.Reader) (*assocRJPDU, error) {

	// read the rest of the pdu
	buf, err := readBytes(reader, 4)
	if err != nil {
		return nil, err
	}

	// read the result, source and reason
	result := buf[1]
	source := buf[2]
	reason := buf[3]

	// build the pdu
	assocRJPDU := &assocRJPDU{
		result: result,
		source: source,
		reason: reason,
	}

	// return the pdu
	return assocRJPDU, nil
}

// Write writes an associate reject PDU
func (assocRJPDU *assocRJPDU) Write(writer io.Writer) error {

	// construct the abort pdu
	buf := []byte{
		0x00,              // reserved
		assocRJPDU.result, // result
		assocRJPDU.source, // source
		assocRJPDU.reason, // reason
	}

	// construct the base pdu
	pdu := &pdu{
		pduType:   aAssociateRJPDU,  // the type
		pduLength: uint32(len(buf)), // the length
	}

	// write the base pdu
	if err := pdu.Write(writer); err != nil {
		return err
	}

	// write the release pdu
	if err := writeBytes(writer, buf[:]); err != nil {
		return err
	}

	return nil
}
