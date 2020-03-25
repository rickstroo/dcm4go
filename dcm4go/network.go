package dcm4go

import (
	"io"
	"log"
)

type networkReader struct {
	r  io.Reader
	dc chan *deed
}

func makePDUDeed(p *pdu) *deed {
	switch p.typ {
	case aAssociateRQPDU:
		return &deed{e: evt6, p: p}
	case aAssociateACPDU:
		return &deed{e: evt3, p: p}
	case aAssociateRJPDU:
		return &deed{e: evt4, p: p}
	case pDataTFPDU:
		return &deed{e: evt10, p: p}
	case aReleaseRQPDU:
		return &deed{e: evt12, p: p}
	case aReleaseRPPDU:
		return &deed{e: evt13, p: p}
	case aAbortPDU:
		return &deed{e: evt16, p: p}
	}
	return &deed{e: evt19, p: p}
}

func (nr *networkReader) run() {
	for {
		log.Printf("waiting to read pdu")
		p, err := readPDU(nr.r)
		if err != nil {
			if err != io.EOF {
				log.Printf("error while attempting to read pdu, error is %v", err)
				d := &deed{e: evt19} // unrecognized or invalid pdu received
				log.Printf("created deed %s", d)
				nr.dc <- d
				break
			}
			log.Printf("reached EOF while reading pdus")
			d := &deed{e: evt17} // transport closed
			log.Printf("created deed %s", d)
			nr.dc <- d
			break
		}
		log.Printf("read pdu %s", p)
		d := makePDUDeed(p)
		log.Printf("created deed %s", d)
		nr.dc <- d
	}
}

type networkWriter struct {
	w             io.Writer
	pduOutputChan chan *pdu
}

func (nw *networkWriter) run() {
	for {
		log.Printf("waiting to receieve pdu")
		p, ok := <-nw.pduOutputChan
		if !ok {
			log.Printf("something went wrong when reading from pdu channel")
			return
		}
		log.Printf("received pdu %s", p)
		if err := writePDU(nw.w, p); err != nil {
			log.Printf("error while writing pdu, err is %v", err)
			return
		}
	}
}
