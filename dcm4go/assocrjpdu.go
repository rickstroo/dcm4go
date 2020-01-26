package dcm4go

import (
	"fmt"
	"io"
)

// AssocRJPDU is an associate reject PDU
type AssocRJPDU struct {
	result byte
	source byte
	reason byte
}

func (assocRJPDU *AssocRJPDU) String() string {
	return fmt.Sprintf(
		"{result:%d,source:%d,reason:%d}",
		assocRJPDU.result,
		assocRJPDU.source,
		assocRJPDU.reason,
	)
}

func readAssocRJPDU(reader io.Reader) (*AssocRJPDU, error) {

	// skip a byte as per the standard
	if err := skipByte(reader); err != nil {
		return nil, err
	}

	// read the result
	result, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// read the source
	source, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// read the reason
	reason, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	assocRJPDU := &AssocRJPDU{
		result: result,
		source: source,
		reason: reason,
	}

	return assocRJPDU, nil
}
