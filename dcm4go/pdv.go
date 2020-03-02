package dcm4go

import (
	"encoding/binary"
	"fmt"
	"io"
)

// pdv represents a DICOM protocol data value (i.e. PDV)
type pdv struct {
	pcID byte
	mch  byte
	buf  []byte
}

// String returns a string representation of a PDV
func (pdv *pdv) String() string {
	return fmt.Sprintf("{pcID:%v,mch:0x%1X,buf:%v}", pdv.pcID, pdv.mch, pdv.buf)
}

// readPDV reads a PDV from a reader
func readPDV(reader io.Reader) (*pdv, error) {

	// read the pdv length
	len, err := readLong(reader, binary.BigEndian)
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

	// read the bytes
	// notice that we actually read two bytes less,
	// as we've already read the pc id and the
	// message control header, which were included
	// in the length calculation
	buf, err := readBytes(reader, len-2)
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
	//limitReader := io.LimitReader(reader, int64(pdvLength-2))

	// construct a PDV
	pdv := &pdv{
		pcID: pcID,
		mch:  mch,
		buf:  buf,
	}

	// return the pdv and success
	return pdv, nil
}

func (pdv *pdv) isCommand() bool {
	return pdv.mch&0x01 == 0x01
}

func (pdv *pdv) isLast() bool {
	return pdv.mch&0x02 == 0x02
}

func (pdv *pdv) writeTo(writer io.Writer) error {

	// write the length of the pdv
	// note that it is two plus the length of the buffer
	// as we need to include the pc id and message control header
	// in the calculation
	if err := writeLong(writer, uint32(len(pdv.buf)+2), binary.BigEndian); err != nil {
		return err
	}

	// write the pc id
	if err := writeByte(writer, pdv.pcID); err != nil {
		return err
	}

	// write the message control header
	if err := writeByte(writer, pdv.mch); err != nil {
		return err
	}

	// write the bytes
	if err := writeBytes(writer, pdv.buf); err != nil {
		return err
	}

	// return success
	return nil
}
