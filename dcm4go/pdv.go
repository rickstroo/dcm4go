package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PDV represents a DICOM protocol data value (i.e. PDV)
type PDV struct {
	pdvLength   uint32
	pcID        byte
	mch         byte
	limitReader io.Reader
}

// Read implements the Reader interface
func (pdv *PDV) Read(buf []byte) (int, error) {
	return pdv.limitReader.Read(buf)
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

	// set up the pdv reader
	// use a limited reader for the length of the PDV
	// actually, we set it to the length less two to make up for the pcid and mch
	// we should really fix that so that the pdv length field reflects that
	// otherwise we have to remember that everywhere
	// in the end, decided to make the PDV a reader so that we could encapsulate
	// the length logic here.  also, makes it easy to just pass the PDV around
	// in other parts of the code instead of passing the PDV and PDV reader
	// separately.
	limitReader := io.LimitReader(reader, int64(pdvLength-2))

	// construct and return a PDV
	return &PDV{pdvLength, pcID, mch, limitReader}, nil
}

func (pdv *PDV) isCommand() bool {
	return pdv.mch&0x01 == 0x01
}

func (pdv *PDV) isLast() bool {
	return pdv.mch&0x02 == 0x02
}

func writePDV(writer io.Writer, pdv *PDV) error {
	if err := writeLong(writer, pdv.pdvLength, binary.BigEndian); err != nil {
		return err
	}
	if err := writeByte(writer, pdv.pcID); err != nil {
		return err
	}
	if err := writeByte(writer, pdv.mch); err != nil {
		return err
	}
	return nil
}
