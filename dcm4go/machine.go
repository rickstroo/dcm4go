// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"io"
	"log"
	"net"
	"time"
)

// HandleConnection is a quick and dirty entrance into the library
func HandleConnection(conn net.Conn, aeTitle string) error {
	sp := &serviceProvider{}
	machine, err := startMachineForServiceProvider(conn, sp)
	if err != nil {
		return err
	}
	machine.run()
	machine.stop()
	return nil
}

type serviceProvider struct {
}

func (sp *serviceProvider) onAssociateRQ(pdu *pdu) (*pdu, error) {
	return nil, nil
}

type serviceUser struct {
}

// the machine
type machine struct {
	state         *state
	pduInputChan  chan *pdu
	pduOutputChan chan *pdu
	artim         *time.Timer // time has a channel built in
	// funcChan      chan func()
	eventChan chan *event

	sp *serviceProvider
	su *serviceUser
}

func startMachine(conn net.Conn) (*machine, error) {
	pduInputChan := make(chan *pdu, 1)
	pduOutputChan := make(chan *pdu, 1)
	eventChan := make(chan *event, 1)
	pduReader := &pduReader1{reader: conn, pduInputChan: pduInputChan}
	pduWriter := &pduWriter1{writer: conn, pduOutputChan: pduOutputChan}
	go pduReader.run()
	go pduWriter.run()
	machine := &machine{
		pduInputChan:  pduInputChan,
		pduOutputChan: pduOutputChan,
		eventChan:     eventChan,
	}
	return machine, nil
}

func startMachineForServiceProvider(conn net.Conn, sp *serviceProvider) (*machine, error) {
	machine, err := startMachine(conn)
	if err != nil {
		return nil, err
	}
	machine.sp = sp
	machine.state = sta2
	return machine, nil
}

func (machine *machine) stop() {
	// stop the timer
	machine.artim.Stop()
	// close the pdu output channel to stop the pdu writer
	close(machine.pduOutputChan)
	//	close(machine.pduInputChan)
	close(machine.eventChan)
}

func (machine *machine) getEvent() *event {
	machine.artim = time.NewTimer(10 * time.Second)
	for {
		log.Printf("waiting for event")
		select {

		case pdu, ok := <-machine.pduInputChan:
			if !ok {
				log.Printf("pdu input channel closed, machine will stop")
				return nil
			}
			log.Printf("machine received pdu, %v", pdu)

			// // just for fun, we're going to abort after receiving a single pdu
			// abortPDU := &abortPDU{
			// 	source: sourceServiceProviderInitiatedAbort,
			// 	reason: reasonNotSpecified,
			// }
			// pdu, err := createAbortPDU(abortPDU)
			// if err != nil {
			// 	log.Printf("error while creating abort pdu, err is %v", err)
			// 	return
			// }
			// machine.pduOutputChan <- pdu

			switch pdu.typ {
			case aAssociateRQPDU:
				ev6.pdu = pdu
				//machine.eventChan <- ev6
				return ev6
			}

		case <-machine.artim.C:
			log.Printf("timer went off, resetting")
			machine.artim.Reset(10 * time.Second)

			//	machine.eventChan <- ev18
			return ev18
		}
	}
}

// these pdu readers and pdu writers are kinda nice.
// they really simplify the management of pdus if you can
// run them as a separate thread.
// perhaps this state machine is going to result in something useful.
// the trick is really to learn how to use channels.

type pduReader1 struct {
	reader       io.Reader
	pduInputChan chan *pdu
}

func (pduReader *pduReader1) run() {
	for {
		log.Printf("pdu reader waiting to read a pdu")
		pdu, err := readPDU(pduReader.reader)
		if err != nil {
			if err != io.EOF {
				log.Printf("error while reading pdu, error is %v", err)
				break
			}
			log.Printf("reached EOF while reading pdus")
			break
		}
		log.Printf("read pdu, %v", pdu)
		pduReader.pduInputChan <- pdu
	}
	// close the channel to let the machine know there are no more pdus
	close(pduReader.pduInputChan)
}

type pduWriter1 struct {
	writer        io.Writer
	pduOutputChan chan *pdu
}

func (pduWriter *pduWriter1) run() {
	for {
		log.Printf("pdu writer waiting to get a pdu to write")
		select {
		case pdu, ok := <-pduWriter.pduOutputChan:
			if !ok {
				log.Printf("pdu output channel closed, pdu writer will stop")
				return
			}
			log.Printf("pdu writer received pdu, %v", pdu)
			if err := writePDU(pduWriter.writer, pdu); err != nil {
				log.Printf("error while writing pdu, pdu is %v, error is %v", pdu, err)
			} else {
				log.Printf("wrote pdu, %v", pdu)
			}
		}
	}
}

// // Machine implements the DICOM Upper Layer Protocol TCP/IP State Machine
// // It is a little complicated because the standard mixes the state machines
// // for the acceptor and requestor of associations.  I'm going to try building
// // the machine as per the standard.  Then, I expect I'll build a couple of
// // wrappers, one for the Acceptor and one for the Requestor, that provide
// // the transport connections, the handlers and an appropriate start state for
// // each.
// type Machine struct {
// 	addr                      string
// 	listener                  net.Listener
// 	conn                      net.Conn
// 	state                     *State
// 	acceptor                  bool
// 	serviceUserInitiatedAbort bool
// 	aeTitle                   string
// 	capabilities              *Capabilities
// 	assocACPDU                *assocPDU
// }
//
// func (machine *Machine) startTimer() {
// }
//
// func (machine *Machine) stopTimer() {
// }
//
// func (machine *Machine) run() error {
// 	for {
// 		log.Printf("entering state %s", machine.state.name)
// 		if err := machine.state.action(machine); err != nil {
// 			if err != io.EOF {
// 				return err
// 			}
// 			break
// 		}
// 	}
// 	return nil
// }
//

type event struct {
	name string
	defn string
	pdu  *pdu
}

var ev1 = &event{name: "ev1", defn: "A-ASSOCIATE Request (local user)"}
var ev2 = &event{name: "ev2", defn: "Transport Connection Confirm (local transport service)"}
var ev3 = &event{name: "ev3", defn: "A-ASSOCIATE-AC PDU (received on transport connection)"}
var ev4 = &event{name: "ev4", defn: "A-ASSOCIATE-RJ PDU (received on transport connection)"}
var ev5 = &event{name: "ev5", defn: "Transport Connection Indication (local transport service)"}
var ev6 = &event{name: "ev6", defn: "A-ASSOCIATE-RQ PDU (received on transport connection)"}
var ev7 = &event{name: "ev7", defn: "A-ASSOCIATE response primitive (accept)"}
var ev8 = &event{name: "ev8", defn: "A-ASSOCIATE response primitive (reject)"}
var ev9 = &event{name: "ev9", defn: "P-DATA request primitive"}
var ev10 = &event{name: "ev10", defn: "P-DATA-TF PDU"}
var ev11 = &event{name: "ev11", defn: "A-RELEASE Request Primitive"}
var ev12 = &event{name: "ev12", defn: "A-RELEASE-RQ PDU (received on open transport connection)"}
var ev13 = &event{name: "ev13", defn: "A-RELEASE-RP PDU (received on transport connection)"}
var ev14 = &event{name: "ev14", defn: "A-RELEASE Response primitive"}
var ev15 = &event{name: "ev15", defn: "A-ABORT Request primitive"}
var ev16 = &event{name: "ev16", defn: "A-ABORT PDU (received on open transport connection)"}
var ev17 = &event{name: "ev17", defn: "Transport connection closed indication (local transport service)"}
var ev18 = &event{name: "ev18", defn: "ARTIM timer expired (Associate reject/release timer)"}
var ev19 = &event{name: "ev19", defn: "Unrecognized or invalid PDU received"}

type state struct {
	name string
	defn string
}

var sta1 = &state{name: "sta1", defn: "Idle"}
var sta2 = &state{name: "sta2", defn: "Transport connection open (Awaiting A-ASSOCIATE-RQ PDU)"}
var sta3 = &state{name: "sta3", defn: "Awaiting local A-ASSOCIATE response primitive (from local user)"}
var sta4 = &state{name: "sta4", defn: "Awaiting transport connection opening to complete (from local transport service)"}
var sta5 = &state{name: "sta5", defn: "Awaiting A-ASSOCIATE-AC or A-ASSOCIATE-RJ PDU"}
var sta6 = &state{name: "sta6", defn: "Association established and ready for data transfer"}
var sta7 = &state{name: "sta7", defn: "Awaiting A-RELEASE-RP PDU"}
var sta8 = &state{name: "sta8", defn: "Awaiting local A-RELEASE response primitive (from local user)"}
var sta9 = &state{name: "sta9", defn: "Release collision requestor side; awaiting A-RELEASE response (from local user)"}
var sta10 = &state{name: "sta10", defn: "Release collision acceptor side; awaiting A-RELEASE-RP PDU"}
var sta11 = &state{name: "sta11", defn: "Release collision requestor side; awaiting A-RELEASE-RP PDU"}
var sta12 = &state{name: "sta12", defn: "Release collision acceptor side, awaitint A-RELEASE reasponse primitive (from local user)"}
var sta13 = &state{name: "sta13", defn: "Awaiting Transport Connection Close Indication (Association no longer exists)"}

type action func(*machine, *event) *state

var actions = map[*state]map[*event]action{
	sta1:  {ev1: ae1},
	sta2:  {ev6: ae6},
	sta3:  {},
	sta4:  {},
	sta5:  {},
	sta6:  {},
	sta7:  {},
	sta8:  {},
	sta9:  {},
	sta10: {},
	sta11: {},
	sta12: {},
	sta13: {},
}

func (machine *machine) step() *state {
	event := machine.getEvent()
	if event == nil {
		return nil
	}
	action := actions[machine.state][event]
	state := action(machine, event)
	return state
}

func (machine *machine) run() {
	for {
		state := machine.step()
		if state == nil {
			break
		}
		machine.state = state
		if machine.state == sta1 {
			break
		}
	}
}

func ae1(machine *machine, event *event) *state {
	return nil
}

func (machine *machine) startTimer() {
	machine.artim.Reset(10 * time.Second)
}

func (machine *machine) stopTimer() {
	machine.artim.Stop()

}
func (machine *machine) abort() *state {
	return sta13
}

func ae6(machine *machine, event *event) *state {
	log.Printf("this is where i call the service provider to negotiate the association")
	machine.stopTimer()
	pdu, err := machine.sp.onAssociateRQ(event.pdu)
	if err != nil {
		return machine.abort()
	}
	if pdu.typ == aAssociateACPDU {
		machine.pduOutputChan <- pdu
		machine.startTimer()
		return sta3
	}
	if pdu.typ == aAssociateRJPDU {
		machine.pduOutputChan <- pdu
		machine.startTimer()
		return sta13
	}
	return machine.abort()
}

//
// func st1(machine *Machine) error {
// 	if machine.acceptor {
// 		return machine.ae5()
// 	}
// 	return machine.ae1()
// }
//
// func st2(machine *Machine) error {
//
// 	// wait for pdu
// 	pdu, err := readPDU(machine.conn)
// 	if err != nil {
// 		return nil
// 	}
// }

//
// 	// if it is an A-ASSOCIATE-RQ PDU, handle it
// 	if pdu.typ == aAssociateRQPDU {
// 		return machine.ae6(pdu)
// 	}
//
// 	// otherwise, send abort
// 	return machine.aa1()
// }
//
// // Awaiting local A-ASSOCIATE response primitive (from local user)
// func st3(machine *Machine) error {
//
// 	// hmm, in ae 6, just before this, i am supposed to tell the local user
// 	// that we accepted an association, and then we go to this state
// 	// to wait for the local user to respond.  i wonder how that is best
// 	// implemented?  in go, we could use channels to communicate between
// 	// the local user and the state machine.  is that really how this is
// 	// all supposed to work?
//
// 	return machine.ae7()
// }
//
// func st4(machine *Machine) error {
// 	return nil
// }
//
// // this is the start sate for a requestor
// func st5(machine *Machine) error {
//
// 	// wait for pdu
// 	pdu, err := readPDU(machine.conn)
// 	if err != nil {
// 		return err
// 	}
//
// 	// if A-ASSOCIATE-AC, accept
// 	if pdu.typ == aAssociateACPDU {
// 		return machine.ae3(pdu)
// 	}
//
// 	// if A-ASSOCIATE-RJ, reject
// 	if pdu.typ == aAssociateRJPDU {
// 		return machine.ae4(pdu)
// 	}
//
// 	// otherwise, abort
// 	return machine.aa8(sourceServiceProviderInitiatedAbort, reasonUnexpectedPDU)
// }
//
// func st6(machine *Machine) error {
//
// 	// hmm, if we are a requestor, here's where we wait
// 	// for a P-DATA request primitive from the local user
// 	// and we go to dt1/sta6
// 	// or we get an abort request form the local user and we go to aa1
// 	// or the connection is closed and we go to aa4/sta1
// 	// otherwise, we are an acceptor and we wait for a PDU
// 	// similar for a release request from local user, we go to ar1/sta7
//
// 	// wait for PDU
// 	pdu, err := readPDU(machine.conn)
// 	if err != nil {
// 		return err
// 	}
//
// 	// if its a release request, we're good
// 	if pdu.typ == aReleaseRQPDU {
// 		return machine.ar2(pdu)
// 	}
//
// 	// if its an abort request, we're good
// 	if pdu.typ == aAbortPDU {
// 		return machine.aa3(pdu)
// 	}
//
// 	// in this state, we can also receive a request
// 	// from the local user to abort
//
// 	// if P-DATA-TF PDU, we're good
// 	if pdu.typ == pDataTFPDU {
// 		return machine.dt2(pdu)
// 	}
//
// 	// otherwise, abort
// 	return machine.aa8(sourceServiceProviderInitiatedAbort, reasonUnexpectedPDU)
// }
//
// // this is where a user initiated release starts
// func st7(machine *Machine) error {
// 	return nil
// }
//
// func st8(machine *Machine) error {
// 	return nil
// }
//
// func st9(machine *Machine) error {
// 	return nil
// }
//
// func st10(machine *Machine) error {
// 	return nil
// }
//
// func st11(machine *Machine) error {
// 	return nil
// }
//
// func st12(machine *Machine) error {
// 	return nil
// }
//
// func st13(machine *Machine) error {
// 	return io.EOF
// }
//
// func (machine *Machine) ae1() error {
// 	// issue connection request
// 	// go to state 4
// 	machine.state = sta4
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ae2() error {
// 	// send A-ASSOCIATE-RQ PDU
// 	// go to state 5
// 	machine.state = sta5
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ae3(pdu *pdu) error {
// 	// call the associate accept handler
// 	// go to state 4
// 	machine.state = sta4
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ae4(pdu *pdu) error {
// 	// call the associate reject handler
// 	// close the connection
// 	// go to state 1
// 	machine.state = sta1
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ae5() error {
// 	// issue transport connection response primitive (what does that mean?)
// 	conn, err := machine.listener.Accept()
// 	if err != nil {
// 		return err
// 	}
// 	machine.conn = conn
// 	// start ARTIM timer
// 	// go to state 2
// 	machine.state = sta2
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ae6(pdu *pdu) error {
//
// 	// read the associate request
// 	assocRQPDU, err := readAssocRQPDU(pdu)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("assocRQPDU is %v\n", assocRQPDU)
//
// 	// stop timer
// 	machine.stopTimer()
//
// 	// build an ae
// 	ae := NewAE(machine.aeTitle)
//
// 	// attempt to negotiate an association
// 	assocACPDU, assocRJPDU, err := negotiateAssoc(assocRQPDU, ae, machine.capabilities)
// 	if err != nil {
// 		return err
// 	}
// 	if assocACPDU == nil {
// 		log.Printf("assocACPDU is nil")
// 	} else {
// 		log.Printf("assocACPDU is %v\n", assocACPDU)
// 	}
// 	if assocRJPDU == nil {
// 		log.Printf("assocRJPDU is nil")
// 	} else {
// 		log.Printf("assocRJPDU is %v\n", assocRJPDU)
// 	}
//
// 	// was association rejected
// 	if assocRJPDU != nil {
//
// 		// wonder why i write the reject pdu in this action
// 		// but i wait to write the accept pdu in another action?
//
// 		// write the associate reject pdu
// 		if err := assocRJPDU.writeTo(machine.conn); err != nil {
// 			return err
// 		}
//
// 		// state the timer
// 		machine.startTimer()
//
// 		// goto state 6
// 		machine.state = sta6
//
// 		// all is well
// 		return nil
// 	}
//
// 	// issue A-ASSOCIATE indication primitive
// 	// not sure what this means
// 	// perhaps its a call to the DIMSE layer telling it we have an association
//
// 	// remember the associate AC PDU
// 	machine.assocACPDU = assocACPDU
//
// 	// go to state 3
// 	machine.state = sta3
//
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ae7() error {
//
// 	// otherwise, write the associate accept pdu
// 	if err := machine.assocACPDU.writeTo(machine.conn); err != nil {
// 		return err
// 	}
//
// 	// go to state 6
// 	machine.state = sta6
//
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ae8() error {
// 	return nil
// }
//
// func (machine *Machine) dt1() error {
// 	return nil
// }
//
// func (machine *Machine) dt2(pdu *pdu) error {
// 	return nil
// }
//
// func (machine *Machine) ar1() error {
// 	// send A-RELEASE-RQ PDU
// 	// next state is sta7
// 	return nil
// }
//
// func (machine *Machine) ar2(pdu *pdu) error {
// 	// issue A-RELEASE indication primitive
// 	// next state is sta8
// 	return nil
// }
//
// func (machine *Machine) ar3() error {
// 	// issue A-RELEASE confirmation primitive
// 	// close transport connection
// 	// next state is sta1
// 	return nil
// }
//
// func (machine *Machine) ar4() error {
// 	// issue A-RELEASE-RQ PDU
// 	// start ARTIM timer
// 	machine.startTimer()
// 	// next state is sta13
// 	machine.state = sta13
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ar5() error {
// 	// stop ARTIM timer
// 	machine.stopTimer()
// 	// next state is sta1
// 	machine.state = sta1
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ar6() error {
// 	// issue P-DATA indication
// 	// next state is sta7
// 	machine.state = sta7
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ar7() error {
// 	// issue P-DATA-TF PDU
// 	// next state is sta8
// 	machine.state = sta8
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ar8() error {
// 	// issue A-RELEASE indication (release condition)
// 	// if association requestor, next state is sta9
// 	// otherwise, next state is sta10
// 	if machine.acceptor {
// 		machine.state = sta10
// 	} else {
// 		machine.state = sta12
// 	}
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ar9() error {
// 	// send A-RELEASE-RP PDU
// 	// next state is sta11
// 	machine.state = sta11
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) ar10() error {
// 	// issue A-RELEASE confirmation primitive
// 	// next state is sta12
// 	machine.state = sta12
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) aa1() error {
// 	// send A-ABORT PDU (service-user source) and start (or restart if already started) ARTIM timer
// 	machine.startTimer()
// 	// next state is sta13
// 	machine.state = sta13
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) aa2() error {
// 	// start timer
// 	machine.startTimer()
// 	// close connection
// 	if err := machine.conn.Close(); err != nil {
// 		return err
// 	}
// 	// go to state 1
// 	machine.state = sta1
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) aa3(pdu *pdu) error {
// 	// if (service-user initiated abort)
// 	// issue A-ABORT indication and close transport connection
// 	// otherwise (service-provider initiated abort)
// 	// issue A-P-ABORT indication and close transport connection
// 	if machine.serviceUserInitiatedAbort {
// 		// issue A-ABORT indication
// 	} else { // if service provider initiated abort
// 		// issue A-P-ABORT indication and close transport connection
// 	}
// 	// close transport connection
// 	if err := machine.conn.Close(); err != nil {
// 		return err
// 	}
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) aa4() error {
// 	// issue A-P-ABORT indication primitive
// 	// go to state 1
// 	machine.state = sta1
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) aa5() error {
// 	// stop the timer
// 	machine.stopTimer()
// 	// go to state 1
// 	machine.state = sta1
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) aa6(pdu *pdu) error {
// 	// go to state 13
// 	machine.state = sta13
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) aa7(source byte, reason byte) error {
// 	// send A-ABORT PDU
// 	pdu := &abortPDU{source: source, reason: reason}
// 	if err := pdu.writeTo(machine.conn); err != nil {
// 		return err
// 	}
// 	// go to state 13
// 	machine.state = sta13
// 	// all is well
// 	return nil
// }
//
// func (machine *Machine) aa8(source byte, reason byte) error {
// 	// send A-ABORT PDU (service provider source)
// 	pdu := &abortPDU{source: source, reason: reason}
// 	if err := pdu.writeTo(machine.conn); err != nil {
// 		return err
// 	}
// 	// issue an A-P-ABORT indication (what does that mean, call a handler?)
// 	// start timer
// 	machine.startTimer()
// 	// all is well
// 	return nil
// }
