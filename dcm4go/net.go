// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"io"
	"log"
)

type pduReader0 struct {
	reader  io.Reader
	done    chan bool
	pduChan chan *pdu
}

func (pduReader *pduReader0) run() {
	for {
		pdu, err := readPDU(pduReader.reader)
		if err != nil {
			if err != io.EOF {
				log.Printf("pduReader.run(), error occured while attempting to read pdu, err is %v", err)
			}
			break
		}
		log.Printf("pduReader.run(), read pdu, pdu is %v", pdu)
		pduReader.pduChan <- pdu
	}
	log.Printf("pduReader.run(), finished reading")
	pduReader.done <- true
}

type pduWriter0 struct {
	writer  io.Writer
	done    chan bool
	pduChan chan *pdu
}

func (pduWriter *pduWriter0) run() {
	for pdu := range pduWriter.pduChan {
		if err := writePDU(pduWriter.writer, pdu); err != nil {
			log.Printf("pduWriter.run(), error occured while attempting to write pdu, err is %v", err)
		}
		log.Printf("pduWriter.run(), wrote pdu, pdu is %v", pdu)
	}
	log.Printf("pduWriter.run(), finished writing")
	pduWriter.done <- true
}
