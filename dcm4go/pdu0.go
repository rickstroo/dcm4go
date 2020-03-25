// Copyright 2020 Rick Stroobosscher.  All rights reserved.

// This source file contains the definition and structures and methods used
// to abort associations.

package dcm4go

// import (
// 	"bytes"
// 	"encoding/binary"
// 	"fmt"
// 	"io"
// )
//
// // PDU is used to define the structure and methods of a Protocol Data Unit.
// // PDUs are used to encode and exchange messages between application entities.
// type PDU struct {
// 	typ byte
// 	buf []byte
// }
//
// // ReadPDU reads a PDU
// func ReadPDU(reader io.Reader) (*PDU, error) {
// 	typ, err := readByte(reader)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := skipByte(reader); err != nil {
// 		return nil, err
// 	}
// 	len, err := readLong(reader, binary.BigEndian)
// 	if err != nil {
// 		return nil, err
// 	}
// 	buf, err := readBytes(reader, len)
// 	if err != nil {
// 		return nil, err
// 	}
// 	pdu := &PDU{
// 		typ: typ,
// 		buf: buf,
// 	}
// 	return pdu, nil
// }
//
// // PDUReader creates a reader for a PDU.
// func PDUReader(pdu *PDU) io.Reader {
// 	return bytes.NewReader(pdu.buf)
// }
//
// // ParsePDU reads and parses a PDU
// func ParsePDU(reader io.Reader) (interface{}, error) {
// 	pdu, err := ReadPDU(reader)
// 	if err != nil {
// 		return nil, err
// 	}
// 	pduReader := PDUReader(pdu)
// 	switch pdu.typ {
// 	case aAbortPDU:
// 		return ReadAAbort(pduReader)
// 	case aAssociateRQPDU:
// 		return ReadAAssociateRQ(pduReader)
// 	case aAssociateACPDU:
// 		return ReadAAssociateAC(pduReader)
// 	case aAssociateRJPDU:
// 		return ReadAAssociateRJ(pduReader)
// 	case aReleaseRQPDU:
// 		return ReadAReleaseRQ(pduReader)
// 	case aReleaseRPPDU:
// 		return ReadAReleaseRP(pduReader)
// 	}
// 	return nil, fmt.Errorf("unrecognized PDU type, %v", pdu.typ)
// }
//
// // WritePDU writes a PDU
// func WritePDU(writer io.Writer, pdu *PDU) error {
// 	if err := writeByte(writer, pdu.typ); err != nil {
// 		return err
// 	}
// 	if err := writeByte(writer, 0x00); err != nil {
// 		return err
// 	}
// 	if err := writeLong(writer, uint32(len(pdu.buf)), binary.BigEndian); err != nil {
// 		return err
// 	}
// 	if err := writeBytes(writer, pdu.buf); err != nil {
// 		return err
// 	}
// 	return nil
// }
