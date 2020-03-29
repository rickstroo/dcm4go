// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import "net"

// Start starts the system
func Start(conn net.Conn) error {

	//create and start a pdu writer
	pduWriterDone := make(chan bool)
	pduWriterChan := make(chan *pdu, 1)
	pduWriter := &pduWriter0{writer: conn, done: pduWriterDone, pduChan: pduWriterChan}
	go pduWriter.run()

	// create and start a pdu reader
	pduReaderDone := make(chan bool)
	pduReaderChan := make(chan *pdu, 1)
	pduReader := &pduReader0{reader: conn, done: pduReaderDone, pduChan: pduReaderChan}
	go pduReader.run()

	// wait for the pdu reader to finish
	<-pduReaderDone

	// close the pdu writer channel, and wait for it to finish
	close(pduWriterChan)
	<-pduWriterDone

	return nil
}
