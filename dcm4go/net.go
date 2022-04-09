// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"io"
	"log"
	"net"
)

func readPDUs(reader io.Reader, pdus chan *pdu, done chan bool) {
	for {
		pdu, err := readPDU(reader)
		if err != nil {
			if err != io.EOF {
				log.Printf("readPDUs, error occured while attempting to read pdu, err is %v", err)
			}
			break
		}
		log.Printf("readPDUs, read pdu, pdu is %v", pdu)
		pdus <- pdu
	}
	log.Printf("readPDUs, finished reading")
	done <- true
}

func decodePDUs(pdus chan *pdu, events chan *event0, done chan bool) {
	for pdu := range pdus {
		event := encodeEvent(pdu)
		log.Printf("decodePDUs, decoded pdu and created event, event is %v", event)
		events <- event
	}
	log.Printf("decodePDUs, finished decoding PDUs")
	done <- true
}

type event0 struct {
	typ byte
	pdu *pdu
}

// String returns a string representation of an event
func (event *event0) String() string {
	if event == nil {
		return "nil"
	}
	return fmt.Sprintf("{typ:%v,pdu:%v}", event.typ, event.pdu)
}

func encodeEvent(pdu *pdu) *event0 {
	switch pdu.typ {
	case aAssociateACPDU:
		return &event0{typ: 3, pdu: pdu}
	case aAssociateRJPDU:
		return &event0{typ: 4, pdu: pdu}
	case aAssociateRQPDU:
		return &event0{typ: 6, pdu: pdu}
	case pDataTFPDU:
		return &event0{typ: 10, pdu: pdu}
	case aReleaseRQPDU:
		return &event0{typ: 12, pdu: pdu}
	case aReleaseRPPDU:
		return &event0{typ: 13, pdu: pdu}
	case aAbortPDU:
		return &event0{typ: 16, pdu: pdu}
	}
	return &event0{typ: 19, pdu: pdu}
}

func writePDUs(writer io.Writer, pdus chan *pdu, done chan bool) {
	for pdu := range pdus {
		if err := writePDU(writer, pdu); err != nil {
			log.Printf("writePDUs, error occured while attempting to write pdu, err is %v", err)
		}
		log.Printf("writePDUs, wrote pdu, pdu is %v", pdu)
	}
	log.Printf("writePDUs, finished writing")
	done <- true
}

type machine0 struct {
	events chan *event0
	done   chan bool
}

func (machine *machine0) run() {
	for event := range machine.events {
		log.Printf("machine.run(), received event, event is %v", event)
		if event.typ == 17 {
			log.Printf("machine.run(), finished reading events")
			break
		}
	}
	log.Printf("machine.run(), finished receiving")
	machine.done <- true
}

// Start starts the system.  We build a pipeline of goroutines from back
// to front, starting with the pdu writer and finishing with the pdu reader.
// Then, we wait for them to finish in the reverse order in which we
// started them, allowing the pipeline to drain.
func Start(conn net.Conn) error {

	// create channels for reading and writing pdus
	pduInputChan := make(chan *pdu, 1)
	pduOutputChan := make(chan *pdu, 1)
	eventChan := make(chan *event0, 1)

	//create and start a pdu writer
	pduWriterDone := make(chan bool)
	go writePDUs(conn, pduOutputChan, pduWriterDone)

	// create and start a state machine
	machineDone := make(chan bool)
	machine := &machine0{events: eventChan, done: machineDone}
	go machine.run()

	// create and start the encoder
	encoderDone := make(chan bool)
	go decodePDUs(pduInputChan, eventChan, encoderDone)

	// create and start a pdu reader
	pduReaderDone := make(chan bool)
	go readPDUs(conn, pduInputChan, pduReaderDone)

	// wait for the pdu reader to finish
	<-pduReaderDone

	// close the pdu input channel, this will signal the encoder to finish
	close(pduInputChan)

	// wait for the encoder to finish
	<-encoderDone

	// close the event channel, this will signal the machine to finish
	close(eventChan)

	// wait for the machine to finish
	<-machineDone

	// close the pdu writer channel,this will signal the pdu writer to finish
	close(pduOutputChan)

	// wait for the pdu writer to finish
	<-pduWriterDone

	return nil
}
