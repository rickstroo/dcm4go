package dcm4go

import (
	"encoding/binary"
	"io"
)

type pdu1 struct {
	typ byte
	buf []byte
}

func readPDU1(reader io.Reader) (*pdu1, error) {
	typ, err := readByte(reader)
	if err != nil {
		return nil, err
	}
	if err := skipByte(reader); err != nil {
		return nil, err
	}
	len, err := readLong(reader, binary.BigEndian)
	if err != nil {
		return nil, err
	}
	buf, err := readBytes(reader, len)
	if err != nil {
		return nil, err
	}
	pdu := &pdu1{
		typ: typ,
		buf: buf,
	}
	return pdu, nil
}

func writePDU1(writer io.Writer, pdu *pdu1) error {
	// if err := writeByte(writer, pdu.typ); err != nil {
	// 	return err
	// }
	// if err := writeByte(writer, 0x00); err != nil {
	// 	return err
	// }
	// if err := writeLong(writer, uint32(len(pdu.buf)), binary.BigEndian); err != nil {
	// 	return err
	// }
	// if err := writeBytes(writer, pdu.buf); err != nil {
	// 	return err
	// }
	var buf [6]byte
	buf[0] = pdu.typ
	buf[1] = 0x00
	binary.BigEndian.PutUint32(buf[2:6], uint32(len(pdu.buf)))
	if _, err := writer.Write(buf[:]); err != nil {
		return err
	}
	if _, err := writer.Write(pdu.buf); err != nil {
		return err
	}
	return nil
}

// func pduReader1(pdu *pdu1) io.Reader {
// 	return bytes.NewReader(pdu.buf)
// }
//
// func pduWriter1() *bytes.Buffer {
// 	return bytes.NewBuffer(make([]byte, 256))
// }

type abort1 struct {
	source byte
	reason byte
}

func decodeAbort1(pdu *pdu1) (*abort1, error) {
	// reader := pduReader1(pdu)
	// if err := skipByte(reader); err != nil {
	// 	return nil, err
	// }
	// if err := skipByte(reader); err != nil {
	// 	return nil, err
	// }
	// source, err := readByte(reader)
	// if err != nil {
	// 	return nil, err
	// }
	// reason, err := readByte(reader)
	// if err != nil {
	// 	return nil, err
	// }
	source := pdu.buf[2]
	reason := pdu.buf[3]
	abort := &abort1{
		source: source,
		reason: reason,
	}
	return abort, nil
}

func encodeAbort1(abort *abort1) (*pdu1, error) {
	// writer := pduWriter1()
	// if err := writeByte(writer, 0x00); err != nil {
	// 	return nil, err
	// }
	// if err := writeByte(writer, 0x00); err != nil {
	// 	return nil, err
	// }
	// if err := writeByte(writer, abort.source); err != nil {
	// 	return nil, err
	// }
	// if err := writeByte(writer, abort.reason); err != nil {
	// 	return nil, err
	// }
	// pdu := &pdu1{
	// 	typ: 0x07,
	// 	buf: writer.Bytes(),
	// }
	var buf [4]byte
	buf[0] = 0x00
	buf[1] = 0x00
	buf[2] = abort.source
	buf[3] = abort.reason
	pdu := &pdu1{
		typ: 0x07,
		buf: buf[:],
	}
	return pdu, nil
}
