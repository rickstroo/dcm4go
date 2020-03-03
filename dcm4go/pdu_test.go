// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"
)

// all the following code was placed here for testing purposes only
// perhaps this should be part of a test or part of the pdu reader

func initPDUTest(buf []byte) *bytes.Buffer {
	return bytes.NewBuffer(buf)
}

func TestPDU(t *testing.T) {
	reader := initPDUTest([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02})
	pdu, err := readPDU(reader)
	if err != nil {
		t.Error(err)
	}
	if err := testByteEquals(0x01, pdu.typ); err != nil {
		t.Error(err)
	}
	if err := testLongEquals(0x02, uint32(len(pdu.buf))); err != nil {
		t.Error(err)
	}
	if err := testByteEquals(0x01, pdu.buf[0]); err != nil {
		t.Error(err)
	}
	if err := testByteEquals(0x02, pdu.buf[1]); err != nil {
		t.Error(err)
	}
}

func TestEOFWhileReadingPDUType(t *testing.T) {
	reader := initPDUTest([]byte{})
	_, err := readPDU(reader)
	if err := testErrEquals(io.EOF, err); err != nil {
		t.Error(err)
	}
}

func TestEOFWhileSkippingByteAfterPDUTYPE(t *testing.T) {
	reader := initPDUTest([]byte{0x01})
	_, err := readPDU(reader)
	if err := testErrEquals(io.EOF, err); err != nil {
		t.Error(err)
	}
}

func TestEOFWhileReadingPDULength(t *testing.T) {
	reader := initPDUTest([]byte{0x01, 0x00})
	_, err := readPDU(reader)
	if err := testErrEquals(io.EOF, err); err != nil {
		t.Error(err)
	}
}

func TestEOFWhileReadingPDUBytes(t *testing.T) {
	reader := initPDUTest([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01})
	_, err := readPDU(reader)
	if err := testErrEquals(io.ErrUnexpectedEOF, err); err != nil {
		t.Error(err)
	}
}

func TestStringPDU(t *testing.T) {
	reader := initPDUTest([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02})
	pdu, err := readPDU(reader)
	if err != nil {
		t.Error(err)
	}
	if err := testStringEquals("{typ:1,buf:[1 2]}", pdu.String()); err != nil {
		t.Error(err)
	}
}

// nextPDU reads the next PDU
func nextPDU(reader io.Reader) (interface{}, error) {

	pdu, err := readPDU(reader)
	if err != nil {
		return nil, err
	}

	byteReader := bytes.NewBuffer(pdu.buf)

	switch pdu.typ {
	case aAbortPDU:
		abortPDU, err := readAbortPDU(byteReader)
		if err != nil {
			return nil, err
		}
		log.Printf("received abort pdu, %v", abortPDU)
		return abortPDU, nil
	case aAssociateRQPDU:
		assocRQPDU, err := readAssocRQPDU(byteReader)
		if err != nil {
			return nil, err
		}
		log.Printf("received associate request pdu, %v", assocRQPDU)
		return assocRQPDU, nil
	case aAssociateACPDU:
		assocACPDU, err := readAssocACPDU(byteReader)
		if err != nil {
			return nil, err
		}
		log.Printf("received associate accept pdu, %v", assocACPDU)
		return assocACPDU, nil
	case aAssociateRJPDU:
		assocRJPDU, err := readAssocRJPDU(byteReader)
		if err != nil {
			return nil, err
		}
		log.Printf("received associate reject pdu, %v", assocRJPDU)
		return assocRJPDU, nil
	case aReleaseRQPDU:
		releaseRQPDU, err := readReleaseRQPDU(byteReader)
		if err != nil {
			return nil, err
		}
		log.Printf("received release request pdu, %v", releaseRQPDU)
		return releaseRQPDU, nil
	case aReleaseRPPDU:
		releaseRPPDU, err := readReleaseRPPDU(byteReader)
		if err != nil {
			return nil, err
		}
		log.Printf("received release response pdu, %v", releaseRPPDU)
		return releaseRPPDU, nil
	case pDataTFPDU:
		dataTFPDU, err := readDataTFPDU(byteReader)
		if err != nil {
			return nil, err
		}
		log.Printf("received data transfer pdu, %v", dataTFPDU)
		return dataTFPDU, nil
	}

	return nil, fmt.Errorf("pdu type not recognized, %v", pdu.typ)
}
