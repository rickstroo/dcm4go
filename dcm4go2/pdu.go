package dcm4go2

import (
	"encoding/binary"
	"io"
)

type pdu struct {
	typ byte
	buf []byte
}

func readPDU(reader io.Reader) (*pdu, error) {
	var hdr [6]byte
	if _, err := io.ReadFull(reader, hdr[:]); err != nil {
		return nil, err
	}
	typ := hdr[0]
	len := binary.BigEndian.Uint32(hdr[2:6])
	buf := make([]byte, len)
	if _, err := io.ReadFull(reader, buf); err != nil {
		return nil, err
	}
	pdu := &pdu{
		typ: typ,
		buf: buf,
	}
	return pdu, nil
}

func writePDU(writer io.Writer, pdu *pdu) error {
	var hdr [6]byte
	hdr[0] = pdu.typ
	hdr[1] = 0x00
	binary.BigEndian.PutUint32(hdr[2:6], uint32(len(pdu.buf)))
	if _, err := writer.Write(hdr[:]); err != nil {
		return err
	}
	if _, err := writer.Write(pdu.buf); err != nil {
		return err
	}
	return nil
}

type abort struct {
	source byte
	reason byte
}

func decodeAbort(pdu *pdu) (*abort, error) {
	source := pdu.buf[2]
	reason := pdu.buf[3]
	abort := &abort{
		source: source,
		reason: reason,
	}
	return abort, nil
}

func encodeAbort(abort *abort) (*pdu, error) {
	var buf [4]byte
	buf[0] = 0x00
	buf[1] = 0x00
	buf[2] = abort.source
	buf[3] = abort.reason
	pdu := &pdu{
		typ: 0x07,
		buf: buf[:],
	}
	return pdu, nil
}

type pdv struct {
	pcid byte
	mch  byte
	buf  []byte
}

type dataTransfer struct {
	pdvs []*pdv
}

func decodeDataTransfer(pdu *pdu) (*dataTransfer, error) {
	pdvs := make([]*pdv, 1)
	i := 0
	for i < len(pdu.buf) {
		len := binary.BigEndian.Uint32(pdu.buf[i : i+4])
		pcid := pdu.buf[i+4]
		mch := pdu.buf[i+5]
		buf := pdu.buf[i+6 : i+6+int(len)-2]
		pdv := &pdv{
			pcid: pcid,
			mch:  mch,
			buf:  buf,
		}
		pdvs = append(pdvs, pdv)
		i += int(len)
	}
	dataTransfer := &dataTransfer{
		pdvs: pdvs,
	}
	return dataTransfer, nil
}

func encodeDataTransfer(dataTransfer *dataTransfer) (*pdu, error) {
	buf := make([]byte, 16*1024)
	for _, pdv := range dataTransfer.pdvs {
		var hdr [6]byte
		binary.BigEndian.PutUint32(hdr[0:4], uint32(len(pdv.buf))+2)
		hdr[4] = pdv.pcid
		hdr[5] = pdv.mch
		buf = append(buf, hdr[:]...)
		buf = append(buf, pdv.buf...)
	}
	pdu := &pdu{
		typ: 0x05,
		buf: buf,
	}
	return pdu, nil
}
