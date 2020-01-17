package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PDV represents a DICOM protocol data value (i.e. PDV)
type PDV struct {
	pdvLength uint32
	pcID      byte
	mch       byte
}

// String returns a string representation of a PDU
func (pdv *PDV) String() string {
	return fmt.Sprintf("pdvLength: %v, pcID: %v, mch: 0x%1X", pdv.pdvLength, pdv.pcID, pdv.mch)
}

// readPDV reads a PDV from a reader
func readPDV(reader io.Reader) (*PDV, error) {

	// read the pdv length
	pdvLength, err := readLong(reader, binary.BigEndian)
	if err != nil {
		return nil, err
	}

	// read the presentation context id
	pcID, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	// read the message control header
	mch, err := readByte(reader)
	if err != nil {
		return nil, err
	}

	return &PDV{pdvLength, pcID, mch}, nil
}

func (pdv *PDV) isCommand() bool {
	return pdv.mch&0x01 == 0x01
}

func (pdv *PDV) isLast() bool {
	return pdv.mch&0x02 == 0x02
}
