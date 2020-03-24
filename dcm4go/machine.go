// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type event struct {
	n string // name
	d string // definition
}

func (e *event) String() string { return fmt.Sprintf("{n:%q,d:%q}", e.n, e.d) }

var evt1 = &event{n: "ev1", d: "A-ASSOCIATE request (local user)"}
var evt2 = &event{n: "ev2", d: "Connection established (for service user)"}
var evt3 = &event{n: "ev3", d: "A-ASSOCIATE-AC PDU (received on transport connection)"}
var evt4 = &event{n: "ev4", d: "A-ASSOCIATE-RJ PDU (received on transport connection)"}
var evt5 = &event{n: "ev5", d: "Connection accepted (for service provider)"}
var evt6 = &event{n: "ev6", d: "A-ASSOCIATE-RQ PDU (received on transport connection)"}
var evt7 = &event{n: "ev7", d: "A-ASSOCIATE response primitive (accept)"}
var evt8 = &event{n: "ev8", d: "A-ASSOCIATE response primitive (reject)"}
var evt9 = &event{n: "ev9", d: "P-DATA request primitive"}
var evt10 = &event{n: "ev10", d: "P-DATA-TF PDU (received on transport connection)"}
var evt11 = &event{n: "ev11", d: "A-RELEASE request primitive"}
var evt12 = &event{n: "ev12", d: "A-RELEASE-RQ PDU (received on transport connection)"}
var evt13 = &event{n: "ev13", d: "A-RELEASE-RP PDU (received on transport connection)"}
var evt14 = &event{n: "ev14", d: "A-RELEASE response primitive"}
var evt15 = &event{n: "ev15", d: "A-ABORT request primitive"}
var evt16 = &event{n: "ev16", d: "A-ABORT PDU (received on transport connection)"}
var evt17 = &event{n: "ev17", d: "Transport connection closed indication (local transport service)"}
var evt18 = &event{n: "ev18", d: "ARTIM timer expired (associate reject/release timer)"}
var evt19 = &event{n: "ev19", d: "Unrecognized or invalid PDU received"}

type state struct {
	n string // name
	d string // definition
}

func (s *state) String() string { return fmt.Sprintf("{n:%q,d:%q}", s.n, s.d) }

var sta1 = &state{n: "sta1", d: "Idle"}
var sta2 = &state{n: "sta2", d: "Transport connection open (Awaiting A-ASSOCIATE-RQ PDU)"}
var sta3 = &state{n: "sta3", d: "Awaiting local A-ASSOCIATE response primitive (from local user)"}
var sta4 = &state{n: "sta4", d: "Awaiting transport connection opening to complete (from local transport service)"}
var sta5 = &state{n: "sta5", d: "Awaiting A-ASSOCIATE-AC or A-ASSOCIATE-RJ PDU"}
var sta6 = &state{n: "sta6", d: "Association established and ready for data transfer"}
var sta7 = &state{n: "sta7", d: "Awaiting A-RELEASE-RP PDU"}
var sta8 = &state{n: "sta8", d: "Awaiting local A-RELEASE response primitive (from local user)"}
var sta9 = &state{n: "sta9", d: "Release collision requestor side; awaiting A-RELEASE response (from local user)"}
var sta10 = &state{n: "sta10", d: "Release collision acceptor side; awaiting A-RELEASE-RP PDU"}
var sta11 = &state{n: "sta11", d: "Release collision requestor side; awaiting A-RELEASE-RP PDU"}
var sta12 = &state{n: "sta12", d: "Release collision acceptor side, awaiting A-RELEASE reasponse primitive (from local user)"}
var sta13 = &state{n: "sta13", d: "Awaiting Transport Connection Close Indication (Association no longer exists)"}

type action struct {
	n string                       // name
	d string                       // definition
	f func(*machine, *deed) *state // function
}

func (a *action) String() string { return fmt.Sprintf("{n:%q,d:%q}", a.n, a.d) }

var ae1 = &action{n: "ae1", d: "Issue Transport Connect request primitive to local transport service.  Next state is Sta4.",
	f: func(m *machine, d *deed) *state {
		return sta4
	},
}

var ae2 = &action{n: "ae2", d: "Send A-ASSOCIATE-RQ-PDU.  Next state is Sta5.",
	f: func(m *machine, e *deed) *state {
		return sta5
	},
}

var ae3 = &action{n: "ae3", d: "Issue A-ASSOCIATE confirmation (accept) primitive.  Next state is Sta6.",
	f: func(m *machine, e *deed) *state {
		return sta6
	},
}

var ae4 = &action{n: "ae4", d: "Issue A-ASSOCIATE confirmation (reject) primitive and close transport connection.  Next state is Sta1.",
	f: func(m *machine, e *deed) *state {
		if m.conn != nil {
			m.conn.Close()
		}
		return sta1
	},
}

var ae5 = &action{n: "ae5", d: "Issue Transport Connect response primitive. Start ARTIM timer.  Next state is Sta2.",
	f: func(m *machine, d *deed) *state {
		m.conn = d.c
		m.pduReader = &pduReader1{r: m.conn, dc: m.deedChan}
		go m.pduReader.run()
		m.startTimer()
		return sta2
	},
}

func sendAssocReject(writer io.Writer, result byte, source byte, reason byte) {
	assocRJPDU := &assocRJPDU{result: result, source: source, reason: reason}
	pdu, e := createAssocRJPDU(assocRJPDU)
	if e != nil {
		log.Printf("error occured while attempting to create associate reject pdu, err is %v", e)
		return
	}
	if e := writePDU(writer, pdu); e != nil {
		log.Printf("error occured while attempting to write associate reject pdu, err is %v", e)
		return
	}
}

var ae6 = &action{n: "ae6", d: "Stop ARTIM timer. If A-ASSOCIATE-RQ acceptable by service provider, issue A-ASSOCIATE indication primitive and next state is Sta3.  Otherwise, issue A-ASSOCIATE-RJ-PDU, start ARTIM timer and next state is Sta13.",
	f: func(m *machine, d *deed) *state {
		m.stopTimer()
		p, e := m.sp.onAssociateRQ(d.p)
		if e != nil {
			log.Printf("error occured while attempting to negotiate association, error %v", e)
			sendAssocReject(m.conn, resultRejectedPermanent, sourceServiceProviderACSERelatedFunction, reasonServiceProviderACSERelatedFunctionNoReasonGiven)
			m.startTimer()
			return sta13
		}
		if p.typ == aAssociateACPDU {
			m.deedChan <- &deed{e: evt7, p: p}
			return sta3
		}
		if p.typ == aAssociateRJPDU {
			m.deedChan <- &deed{e: evt8, p: p}
			return sta3
		}
		log.Printf("unrecognized pdu from associate negotiation, %v", p)
		sendAssocReject(m.conn, resultRejectedPermanent, sourceServiceProviderACSERelatedFunction, reasonServiceProviderACSERelatedFunctionNoReasonGiven)
		m.startTimer()
		return sta13
	},
}

var ae7 = &action{n: "ae7", d: "Send A-ASSOCIATE-AC PDU.  Next state is Sta6.",
	f: func(m *machine, d *deed) *state {
		err := writePDU(m.conn, d.p)
		if err != nil {
			log.Printf("error while writing A-ASSOCIATE-AC pdu, error is %v", err)
			return sta13
		}
		return sta6
	},
}

var ae8 = &action{n: "ae8", d: "Send A-ASSOCIATE-RJ PDU.  Next state is Sta13.",
	f: func(m *machine, d *deed) *state {
		err := writePDU(m.conn, d.p)
		if err != nil {
			log.Printf("error while writing A-ASSOCIATE-RJ pdu, error is %v", err)
			return sta13
		}
		return sta13
	},
}

var dt1 = &action{n: "dt1", d: "Send P-DATA-TF PDU.  Next state is Sta6.",
	f: func(m *machine, d *deed) *state {
		err := writePDU(m.conn, d.p)
		if err != nil {
			log.Printf("error while writing P-DATA-TF pdu, error is %v", err)
			return sta13
		}
		return sta6
	},
}

var dt2 = &action{n: "dt2", d: "Send P-DATA indication primitive.  Next state is Sta6.",
	f: func(m *machine, d *deed) *state {
		if err := m.sp.onDataTF(d.p); err != nil {
			log.Printf("err is %v", err)
		}
		return sta6
	},
}

var ar1 = &action{n: "ar1", d: "Send A-RELEASE-RQ PDU.  Next state is Sta7.", f: func(m *machine, d *deed) *state { return sta7 }}
var ar2 = &action{n: "ar2", d: "Issue A-RELEASE indication primitive.  Next state is Sta8.",
	f: func(m *machine, d *deed) *state {
		m.sp.onReleaseRQ(d.p, false)
		return sta8
	},
}
var ar3 = &action{n: "ar3", d: "Issue A-RELEASE confirmation primitive, and close transport connection.  Next state is Sta1.", f: func(m *machine, d *deed) *state { return sta1 }}
var ar4 = &action{n: "ar4", d: "Issue A-RELEASE-RP PDU and start ARTIM timer.  Next state is Sta13.",
	f: func(m *machine, d *deed) *state {
		p, e := createReleaseRPPDU()
		if e != nil {
			log.Printf("error while creating A-RELEASE-RP PDU, error is %v", e)
			return sta13
		}
		if e := writePDU(m.conn, p); e != nil {
			log.Printf("error while writing A-RELEASE-RP PDU, error is %v", e)
			return sta13
		}
		m.startTimer()
		return sta13
	},
}
var ar5 = &action{n: "ar5", d: "Stop ARTIM timer.  Next state is Sta1.",
	f: func(m *machine, de *deed) *state {
		m.stopTimer()
		return sta1
	},
}
var ar6 = &action{n: "ar6", d: "Issue P-DATA indication.  Next state is Sta7.", f: func(m *machine, d *deed) *state { return sta7 }}

var ar7 = &action{n: "ar7", d: "Issue P-DATA-TF PDU.  Next state is Sta8.",
	f: func(m *machine, d *deed) *state {
		return sta8
	},
}

var ar8 = &action{n: "ar8", d: "Issue A-RELEASE indication (release collision).  If association-requestor (service user), next state is Sta9, other next state is Sta10.",
	f: func(m *machine, d *deed) *state {
		m.sp.onReleaseRQ(d.p, true)
		if m.isServiceUser {
			return sta9
		}
		return sta10
	},
}

var ar9 = &action{n: "ar9", d: "Send A-RELEASE-RP PDU.  Next state is Sta11.",
	f: func(m *machine, d *deed) *state {
		p, e := createReleaseRPPDU()
		if e != nil {
			log.Printf("error while creating A-RELEASE-RP PDU, error is %v", e)
			return sta13
		}
		if e := writePDU(m.conn, p); e != nil {
			log.Printf("error while writing A-RELEASE-RP PDU, error is %v", e)
			return sta13
		}
		return sta11
	},
}

var ar10 = &action{n: "ar10", d: "Issue A-RELEASE confirmation primitive.  Next state is Sta12.",
	f: func(m *machine, d *deed) *state {
		return sta12
	},
}

var aa1 = &action{n: "aa1", d: "Send A-ABORT PDU (service-user source) and start (or restart if already started) ARTIM timer.  Next state is Sta13.",
	f: func(m *machine, d *deed) *state {
		return sta13
	},
}

var aa2 = &action{n: "aa2", d: "Stop ARTIM timer if running.  Close transport connection.  Next state is Sta1.",
	f: func(m *machine, d *deed) *state {
		m.stopTimer()
		m.conn.Close()
		return sta1
	},
}
var aa3 = &action{n: "aa3", d: "", f: func(m *machine, d *deed) *state { return sta1 }}
var aa4 = &action{n: "aa4", d: "", f: func(m *machine, d *deed) *state { return sta1 }}
var aa5 = &action{n: "aa5", d: "", f: func(m *machine, d *deed) *state { return sta1 }}
var aa6 = &action{n: "aa6", d: "", f: func(m *machine, d *deed) *state { return sta13 }}
var aa7 = &action{n: "aa7", d: "", f: func(m *machine, d *deed) *state { return sta13 }}
var aa8 = &action{n: "aa8", d: "", f: func(m *machine, d *deed) *state { return sta13 }}

var actions = map[*event]map[*state]*action{
	evt1:  {sta1: ae1},
	evt2:  {sta4: ae2},
	evt3:  {sta2: aa1, sta3: aa8, sta5: ae3, sta6: aa8, sta7: aa8, sta8: aa8, sta9: aa8, sta10: aa8, sta11: aa8, sta12: aa8, sta13: aa6},
	evt4:  {sta2: aa1, sta3: aa8, sta5: ae4, sta6: aa8, sta7: aa8, sta8: aa8, sta9: aa8, sta10: aa8, sta11: aa8, sta12: aa8, sta13: aa6},
	evt5:  {sta1: ae5},
	evt6:  {sta2: ae6, sta3: aa8, sta5: aa8, sta6: aa8, sta7: aa8, sta8: aa8, sta9: aa8, sta10: aa8, sta11: aa8, sta12: aa8, sta13: aa7},
	evt7:  {sta3: ae7},
	evt8:  {sta3: ae8},
	evt9:  {sta6: dt1, sta8: ar7},
	evt10: {sta2: aa1, sta3: aa8, sta8: aa3, sta6: dt2, sta7: ar6, sta8: aa8, sta9: aa8, sta10: aa8, sta11: aa8, sta12: aa8, sta13: aa6},
	evt11: {sta6: ar1},
	evt12: {sta2: aa1, sta3: aa8, sta5: aa8, sta6: ar2, sta7: ar8, sta8: aa8, sta9: aa8, sta10: aa8, sta11: aa8, sta12: aa8, sta13: aa6},
	evt13: {sta2: aa1, sta3: aa8, sta5: aa8, sta6: aa8, sta7: ar3, sta8: aa8, sta9: aa8, sta10: ar10, sta11: ar3, sta12: aa8, sta13: aa6},
	evt14: {sta8: ar4, sta9: ar9, sta12: ar4},
	evt15: {sta3: aa1, sta4: aa2, sta5: aa1, sta6: aa1, sta7: aa1, sta8: aa1, sta9: aa1, sta10: aa1, sta11: aa1, sta12: aa1},
	evt16: {sta2: aa2, sta3: aa3, sta5: aa3, sta6: aa3, sta7: aa3, sta8: aa3, sta9: aa3, sta10: aa3, sta11: aa3, sta12: aa3, sta13: aa2},
	evt17: {sta2: aa5, sta3: aa4, sta4: aa4, sta5: aa4, sta6: aa4, sta7: aa4, sta8: aa4, sta9: aa4, sta10: aa4, sta11: aa4, sta12: aa4, sta13: ar5},
	evt18: {sta2: aa2, sta13: aa2},
	evt19: {sta2: aa1, sta3: aa8, sta5: aa8, sta6: aa8, sta7: ar8, sta8: aa8, sta9: aa8, sta10: aa8, sta11: aa8, sta12: aa8, sta13: aa7},
}

type deed struct {
	e *event
	c net.Conn
	p *pdu
}

func (d *deed) String() string {
	s := fmt.Sprintf("{e:%s", d.e)
	if d.c != nil {
		s += fmt.Sprintf(",c:%s", d.c.LocalAddr())
	}
	if d.p != nil {
		s += fmt.Sprintf(",p:%s", d.p)
	}
	s += fmt.Sprintf("}")
	return s
}

type machine struct {
	sp            *serviceProvider
	conn          net.Conn
	state         *state
	isServiceUser bool
	deedChan      chan *deed
	pduReader     *pduReader1
	artim         *time.Timer
}

func (m *machine) startTimer() {
	m.artim = time.AfterFunc(time.Duration(10)*time.Second,
		func() {
			m.deedChan <- &deed{e: evt18}
		},
	)
	log.Printf("started timer")
}

func (m *machine) stopTimer() {
	if m.artim != nil {
		m.artim.Stop()
		log.Printf("stopped timer")
	}
}

func findAction(e *event, s *state) *action {
	a, ok := actions[e][s]
	if !ok {
		log.Printf("no action found for event %s and state %s", e, s)
		return nil
	}
	return a
}

func (m *machine) run() error {
	log.Printf("first state is %s", m.state)
	for {
		d := <-m.deedChan
		log.Printf("read deed %s", d)
		a := findAction(d.e, m.state)
		log.Printf("found action %s", a)
		s := a.f(m, d)
		log.Printf("next state is %s", s)
		// break when we get back to the idle state
		if s == sta1 {
			break
		}
		// otherwise, update the state
		m.state = s
	}
	return nil
}

// StartMachineForServiceProvider starts the machine for a service provider
func StartMachineForServiceProvider(conn net.Conn, aeTitle string, capabilities *Capabilities) error {
	sp := newServiceProvider(aeTitle, capabilities)
	m := &machine{
		sp:            sp,                  // the service provider
		isServiceUser: false,               // is a service provider
		deedChan:      make(chan *deed, 1), // a channel for events
		state:         sta1,                // the idle state
	}
	sp.machine = m
	// the first event is a connection event
	m.deedChan <- &deed{e: evt5, c: conn}
	// run the machine
	if err := m.run(); err != nil {
		return err
	}
	return nil
}

type pduReader1 struct {
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

func (pr *pduReader1) run() {
	for {
		log.Printf("waiting to read pdu")
		p, err := readPDU(pr.r)
		if err != nil {
			if err != io.EOF {
				log.Printf("error while attempting to read pdu, error is %v", err)
				d := &deed{e: evt19}
				log.Printf("created deed %s", d)
				pr.dc <- d
				break
			}
			log.Printf("reached EOF while reading pdus")
			break
		}
		log.Printf("read pdu %s", p)
		d := makePDUDeed(p)
		log.Printf("created deed %s", d)
		pr.dc <- d
	}
}

// // StartMachineForServiceProvider starts the state machine for a service provider.
// func StartMachineForServiceProvider(conn net.Conn, aeTitle string) {
// 	machine := &machine{
// 		state: sta5,
// 	}
// 	machine.run()
// }
//
// type serviceProvider struct {
// }
//
// func (serviceProvider *serviceProvider) onAssociateRQ(pdu *pdu) (*pdu, error) {
// 	return nil, nil
// }
//
// type serviceUser struct {
// }
//
// // the machine
// type machine struct {
// 	state           stateType
// 	pduInputChan    chan *pdu
// 	pduOutputChan   chan *pdu
// 	artim           *time.Timer
// 	eventChan       chan *event
// 	serviceProvider *serviceProvider
// 	serviceUser     *serviceUser
// }
//
// func startMachine(conn net.Conn) *machine {
// 	pduInputChan := make(chan *pdu, 1)
// 	pduOutputChan := make(chan *pdu, 1)
// 	eventChan := make(chan *event, 1)
// 	pduReader := &pduReader1{reader: conn, pduInputChan: pduInputChan}
// 	pduWriter := &pduWriter1{writer: conn, pduOutputChan: pduOutputChan}
// 	go pduReader.run()
// 	go pduWriter.run()
// 	machine := &machine{
// 		pduInputChan:  pduInputChan,
// 		pduOutputChan: pduOutputChan,
// 		eventChan:     eventChan,
// 	}
// 	return machine
// }
//
// func startMachineForServiceProvider(conn net.Conn, serviceProvider *serviceProvider) *machine {
// 	machine := startMachine(conn)
// 	machine.serviceProvider = serviceProvider
// 	machine.state = sta2
// 	return machine
// }
//
// func (machine *machine) stop() {
// 	// stop the timer
// 	machine.artim.Stop()
// 	// close the pdu output channel to stop the pdu writer
// 	close(machine.pduOutputChan)
// 	//	close(machine.pduInputChan)
// 	close(machine.eventChan)
// }
//
// // func (machine *machine) wait() (*event, *pdu, error) {
// // 	eventChan := make(chan *event, 1)
// // 	pduChan := make(chan *pdu, 1)
// //
// // }
//
// func (machine *machine) getEvent() *event {
// 	machine.artim = time.NewTimer(10 * time.Second)
// 	for {
// 		log.Printf("waiting for event")
// 		select {
//
// 		case pdu, ok := <-machine.pduInputChan:
// 			if !ok {
// 				log.Printf("pdu input channel closed, machine will stop")
// 				return nil
// 			}
// 			log.Printf("machine received pdu, %v", pdu)
//
// 			// // just for fun, we're going to abort after receiving a single pdu
// 			// abortPDU := &abortPDU{
// 			// 	source: sourceServiceProviderInitiatedAbort,
// 			// 	reason: reasonNotSpecified,
// 			// }
// 			// pdu, err := createAbortPDU(abortPDU)
// 			// if err != nil {
// 			// 	log.Printf("error while creating abort pdu, err is %v", err)
// 			// 	return
// 			// }
// 			// machine.pduOutputChan <- pdu
//
// 			switch pdu.typ {
// 			case aAssociateRQPDU:
// 				ev6.pdu = pdu
// 				//machine.eventChan <- ev6
// 				return ev6
// 			}
//
// 		case <-machine.artim.C:
// 			log.Printf("timer went off, resetting")
// 			machine.artim.Reset(10 * time.Second)
//
// 			//	machine.eventChan <- ev18
// 			return ev18
// 		}
// 	}
// }
//
// // these pdu readers and pdu writers are kinda nice.
// // they really simplify the management of pdus if you can
// // run them as a separate thread.
// // perhaps this state machine is going to result in something useful.
// // the trick is really to learn how to use channels.
//
// type pduReader1 struct {
// 	reader       io.Reader
// 	pduInputChan chan *pdu
// }
//
// func (pduReader *pduReader1) run() {
// 	for {
// 		log.Printf("pdu reader waiting to read a pdu")
// 		pdu, err := readPDU(pduReader.reader)
// 		if err != nil {
// 			if err != io.EOF {
// 				log.Printf("error while reading pdu, error is %v", err)
// 				break
// 			}
// 			log.Printf("reached EOF while reading pdus")
// 			break
// 		}
// 		log.Printf("read pdu, %v", pdu)
// 		pduReader.pduInputChan <- pdu
// 	}
// 	// close the channel to let the machine know there are no more pdus
// 	close(pduReader.pduInputChan)
// }
//
// type pduWriter1 struct {
// 	writer        io.Writer
// 	pduOutputChan chan *pdu
// }
//
// func (pduWriter *pduWriter1) run() {
// 	for {
// 		log.Printf("pdu writer waiting to get a pdu to write")
// 		select {
// 		case pdu, ok := <-pduWriter.pduOutputChan:
// 			if !ok {
// 				log.Printf("pdu output channel closed, pdu writer will stop")
// 				return
// 			}
// 			log.Printf("pdu writer received pdu, %v", pdu)
// 			if err := writePDU(pduWriter.writer, pdu); err != nil {
// 				log.Printf("error while writing pdu, pdu is %v, error is %v", pdu, err)
// 			} else {
// 				log.Printf("wrote pdu, %v", pdu)
// 			}
// 		}
// 	}
// }
//
// // // Machine implements the DICOM Upper Layer Protocol TCP/IP State Machine
// // // It is a little complicated because the standard mixes the state machines
// // // for the acceptor and requestor of associations.  I'm going to try building
// // // the machine as per the standard.  Then, I expect I'll build a couple of
// // // wrappers, one for the Acceptor and one for the Requestor, that provide
// // // the transport connections, the handlers and an appropriate start state for
// // // each.
// // type Machine struct {
// // 	addr                      string
// // 	listener                  net.Listener
// // 	conn                      net.Conn
// // 	state                     *State
// // 	acceptor                  bool
// // 	serviceUserInitiatedAbort bool
// // 	aeTitle                   string
// // 	capabilities              *Capabilities
// // 	assocACPDU                *assocPDU
// // }
// //
// // func (machine *Machine) startTimer() {
// // }
// //
// // func (machine *Machine) stopTimer() {
// // }
// //
// // func (machine *Machine) run() error {
// // 	for {
// // 		log.Printf("entering state %s", machine.state.name)
// // 		if err := machine.state.action(machine); err != nil {
// // 			if err != io.EOF {
// // 				return err
// // 			}
// // 			break
// // 		}
// // 	}
// // 	return nil
// // }
// //
//
// // type eventType byte
// //
// // const (
// // 	evt01 = eventType(1)
// // 	evt02 = eventType(2)
// // 	evt03 = eventType(3)
// // 	evt04 = eventType(4)
// // 	evt05 = eventType(5)
// // 	evt06 = eventType(6)
// // 	evt07 = eventType(7)
// // 	evt08 = eventType(8)
// // 	evt09 = eventType(9)
// // 	evt10 = eventType(10)
// // 	evt11 = eventType(11)
// // 	evt12 = eventType(12)
// // 	evt13 = eventType(13)
// // 	evt14 = eventType(14)
// // 	evt15 = eventType(15)
// // 	evt16 = eventType(16)
// // 	evt17 = eventType(17)
// // 	evt18 = eventType(18)
// // 	evt19 = eventType(19)
// // )
// //
//
// func (machine *machine) step() stateType {
// 	event := machine.getEvent()
// 	if event == nil {
// 		return nil
// 	}
// 	action, ok := actions[event.eventType][machine.state]
// 	if !ok {
// 		log.Printf("hmm, did not find an action")
// 		return nil
// 	}
// 	state := action(machine, event)
// 	return state
// }
//
// func (machine *machine) run() {
// 	for {
// 		state := machine.step()
// 		if state == nil {
// 			break
// 		}
// 		machine.state = state
// 		if machine.state == sta1 {
// 			break
// 		}
// 	}
// }
//
// func (machine *machine) startTimer() {
// 	machine.artim.Reset(10 * time.Second)
// }
//
// func (machine *machine) stopTimer() {
// 	machine.artim.Stop()
//
// }
// func (machine *machine) abort() stateType {
// 	return sta13
// }
//
// func ae1(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ae2(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ae3(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ae4(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ae5(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ae6(machine *machine, event *event) stateType {
// 	log.Printf("this is where i call the service provider to negotiate the association")
// 	machine.stopTimer()
// 	pdu, err := machine.sp.onAssociateRQ(event.pdu)
// 	if err != nil {
// 		return machine.abort()
// 	}
// 	if pdu.typ == aAssociateACPDU {
// 		machine.pduOutputChan <- pdu
// 		machine.startTimer()
// 		return sta3
// 	}
// 	if pdu.typ == aAssociateRJPDU {
// 		machine.pduOutputChan <- pdu
// 		machine.startTimer()
// 		return sta13
// 	}
// 	return machine.abort()
// }
//
// func ae7(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ae8(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func dt1(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func dt2(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar1(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar2(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar3(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar4(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar5(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar6(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar7(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar8(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar9(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func ar10(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func aa1(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func aa2(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func aa3(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func aa4(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func aa5(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func aa6(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func aa7(machine *machine, event *event) stateType {
// 	return nil
// }
//
// func aa8(machine *machine, event *event) stateType {
// 	return nil
// }
//
// //
// // func st1(machine *Machine) error {
// // 	if machine.acceptor {
// // 		return machine.ae5()
// // 	}
// // 	return machine.ae1()
// // }
// //
// // func st2(machine *Machine) error {
// //
// // 	// wait for pdu
// // 	pdu, err := readPDU(machine.conn)
// // 	if err != nil {
// // 		return nil
// // 	}
// // }
//
// //
// // 	// if it is an A-ASSOCIATE-RQ PDU, handle it
// // 	if pdu.typ == aAssociateRQPDU {
// // 		return machine.ae6(pdu)
// // 	}
// //
// // 	// otherwise, send abort
// // 	return machine.aa1()
// // }
// //
// // // Awaiting local A-ASSOCIATE response primitive (from local user)
// // func st3(machine *Machine) error {
// //
// // 	// hmm, in ae 6, just before this, i am supposed to tell the local user
// // 	// that we accepted an association, and then we go to this state
// // 	// to wait for the local user to respond.  i wonder how that is best
// // 	// implemented?  in go, we could use channels to communicate between
// // 	// the local user and the state machine.  is that really how this is
// // 	// all supposed to work?
// //
// // 	return machine.ae7()
// // }
// //
// // func st4(machine *Machine) error {
// // 	return nil
// // }
// //
// // // this is the start sate for a requestor
// // func st5(machine *Machine) error {
// //
// // 	// wait for pdu
// // 	pdu, err := readPDU(machine.conn)
// // 	if err != nil {
// // 		return err
// // 	}
// //
// // 	// if A-ASSOCIATE-AC, accept
// // 	if pdu.typ == aAssociateACPDU {
// // 		return machine.ae3(pdu)
// // 	}
// //
// // 	// if A-ASSOCIATE-RJ, reject
// // 	if pdu.typ == aAssociateRJPDU {
// // 		return machine.ae4(pdu)
// // 	}
// //
// // 	// otherwise, abort
// // 	return machine.aa8(sourceServiceProviderInitiatedAbort, reasonUnexpectedPDU)
// // }
// //
// // func st6(machine *Machine) error {
// //
// // 	// hmm, if we are a requestor, here's where we wait
// // 	// for a P-DATA request primitive from the local user
// // 	// and we go to dt1/sta6
// // 	// or we get an abort request form the local user and we go to aa1
// // 	// or the connection is closed and we go to aa4/sta1
// // 	// otherwise, we are an acceptor and we wait for a PDU
// // 	// similar for a release request from local user, we go to ar1/sta7
// //
// // 	// wait for PDU
// // 	pdu, err := readPDU(machine.conn)
// // 	if err != nil {
// // 		return err
// // 	}
// //
// // 	// if its a release request, we're good
// // 	if pdu.typ == aReleaseRQPDU {
// // 		return machine.ar2(pdu)
// // 	}
// //
// // 	// if its an abort request, we're good
// // 	if pdu.typ == aAbortPDU {
// // 		return machine.aa3(pdu)
// // 	}
// //
// // 	// in this state, we can also receive a request
// // 	// from the local user to abort
// //
// // 	// if P-DATA-TF PDU, we're good
// // 	if pdu.typ == pDataTFPDU {
// // 		return machine.dt2(pdu)
// // 	}
// //
// // 	// otherwise, abort
// // 	return machine.aa8(sourceServiceProviderInitiatedAbort, reasonUnexpectedPDU)
// // }
// //
// // // this is where a user initiated release starts
// // func st7(machine *Machine) error {
// // 	return nil
// // }
// //
// // func st8(machine *Machine) error {
// // 	return nil
// // }
// //
// // func st9(machine *Machine) error {
// // 	return nil
// // }
// //
// // func st10(machine *Machine) error {
// // 	return nil
// // }
// //
// // func st11(machine *Machine) error {
// // 	return nil
// // }
// //
// // func st12(machine *Machine) error {
// // 	return nil
// // }
// //
// // func st13(machine *Machine) error {
// // 	return io.EOF
// // }
// //
// // func (machine *Machine) ae1() error {
// // 	// issue connection request
// // 	// go to state 4
// // 	machine.state = sta4
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ae2() error {
// // 	// send A-ASSOCIATE-RQ PDU
// // 	// go to state 5
// // 	machine.state = sta5
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ae3(pdu *pdu) error {
// // 	// call the associate accept handler
// // 	// go to state 4
// // 	machine.state = sta4
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ae4(pdu *pdu) error {
// // 	// call the associate reject handler
// // 	// close the connection
// // 	// go to state 1
// // 	machine.state = sta1
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ae5() error {
// // 	// issue transport connection response primitive (what does that mean?)
// // 	conn, err := machine.listener.Accept()
// // 	if err != nil {
// // 		return err
// // 	}
// // 	machine.conn = conn
// // 	// start ARTIM timer
// // 	// go to state 2
// // 	machine.state = sta2
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ae6(pdu *pdu) error {
// //
// // 	// read the associate request
// // 	assocRQPDU, err := readAssocRQPDU(pdu)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	log.Printf("assocRQPDU is %v\n", assocRQPDU)
// //
// // 	// stop timer
// // 	machine.stopTimer()
// //
// // 	// build an ae
// // 	ae := NewAE(machine.aeTitle)
// //
// // 	// attempt to negotiate an association
// // 	assocACPDU, assocRJPDU, err := negotiateAssoc(assocRQPDU, ae, machine.capabilities)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	if assocACPDU == nil {
// // 		log.Printf("assocACPDU is nil")
// // 	} else {
// // 		log.Printf("assocACPDU is %v\n", assocACPDU)
// // 	}
// // 	if assocRJPDU == nil {
// // 		log.Printf("assocRJPDU is nil")
// // 	} else {
// // 		log.Printf("assocRJPDU is %v\n", assocRJPDU)
// // 	}
// //
// // 	// was association rejected
// // 	if assocRJPDU != nil {
// //
// // 		// wonder why i write the reject pdu in this action
// // 		// but i wait to write the accept pdu in another action?
// //
// // 		// write the associate reject pdu
// // 		if err := assocRJPDU.writeTo(machine.conn); err != nil {
// // 			return err
// // 		}
// //
// // 		// state the timer
// // 		machine.startTimer()
// //
// // 		// goto state 6
// // 		machine.state = sta6
// //
// // 		// all is well
// // 		return nil
// // 	}
// //
// // 	// issue A-ASSOCIATE indication primitive
// // 	// not sure what this means
// // 	// perhaps its a call to the DIMSE layer telling it we have an association
// //
// // 	// remember the associate AC PDU
// // 	machine.assocACPDU = assocACPDU
// //
// // 	// go to state 3
// // 	machine.state = sta3
// //
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ae7() error {
// //
// // 	// otherwise, write the associate accept pdu
// // 	if err := machine.assocACPDU.writeTo(machine.conn); err != nil {
// // 		return err
// // 	}
// //
// // 	// go to state 6
// // 	machine.state = sta6
// //
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ae8() error {
// // 	return nil
// // }
// //
// // func (machine *Machine) dt1() error {
// // 	return nil
// // }
// //
// // func (machine *Machine) dt2(pdu *pdu) error {
// // 	return nil
// // }
// //
// // func (machine *Machine) ar1() error {
// // 	// send A-RELEASE-RQ PDU
// // 	// next state is sta7
// // 	return nil
// // }
// //
// // func (machine *Machine) ar2(pdu *pdu) error {
// // 	// issue A-RELEASE indication primitive
// // 	// next state is sta8
// // 	return nil
// // }
// //
// // func (machine *Machine) ar3() error {
// // 	// issue A-RELEASE confirmation primitive
// // 	// close transport connection
// // 	// next state is sta1
// // 	return nil
// // }
// //
// // func (machine *Machine) ar4() error {
// // 	// issue A-RELEASE-RQ PDU
// // 	// start ARTIM timer
// // 	machine.startTimer()
// // 	// next state is sta13
// // 	machine.state = sta13
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ar5() error {
// // 	// stop ARTIM timer
// // 	machine.stopTimer()
// // 	// next state is sta1
// // 	machine.state = sta1
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ar6() error {
// // 	// issue P-DATA indication
// // 	// next state is sta7
// // 	machine.state = sta7
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ar7() error {
// // 	// issue P-DATA-TF PDU
// // 	// next state is sta8
// // 	machine.state = sta8
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ar8() error {
// // 	// issue A-RELEASE indication (release condition)
// // 	// if association requestor, next state is sta9
// // 	// otherwise, next state is sta10
// // 	if machine.acceptor {
// // 		machine.state = sta10
// // 	} else {
// // 		machine.state = sta12
// // 	}
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ar9() error {
// // 	// send A-RELEASE-RP PDU
// // 	// next state is sta11
// // 	machine.state = sta11
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) ar10() error {
// // 	// issue A-RELEASE confirmation primitive
// // 	// next state is sta12
// // 	machine.state = sta12
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) aa1() error {
// // 	// send A-ABORT PDU (service-user source) and start (or restart if already started) ARTIM timer
// // 	machine.startTimer()
// // 	// next state is sta13
// // 	machine.state = sta13
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) aa2() error {
// // 	// start timer
// // 	machine.startTimer()
// // 	// close connection
// // 	if err := machine.conn.Close(); err != nil {
// // 		return err
// // 	}
// // 	// go to state 1
// // 	machine.state = sta1
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) aa3(pdu *pdu) error {
// // 	// if (service-user initiated abort)
// // 	// issue A-ABORT indication and close transport connection
// // 	// otherwise (service-provider initiated abort)
// // 	// issue A-P-ABORT indication and close transport connection
// // 	if machine.serviceUserInitiatedAbort {
// // 		// issue A-ABORT indication
// // 	} else { // if service provider initiated abort
// // 		// issue A-P-ABORT indication and close transport connection
// // 	}
// // 	// close transport connection
// // 	if err := machine.conn.Close(); err != nil {
// // 		return err
// // 	}
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) aa4() error {
// // 	// issue A-P-ABORT indication primitive
// // 	// go to state 1
// // 	machine.state = sta1
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) aa5() error {
// // 	// stop the timer
// // 	machine.stopTimer()
// // 	// go to state 1
// // 	machine.state = sta1
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) aa6(pdu *pdu) error {
// // 	// go to state 13
// // 	machine.state = sta13
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) aa7(source byte, reason byte) error {
// // 	// send A-ABORT PDU
// // 	pdu := &abortPDU{source: source, reason: reason}
// // 	if err := pdu.writeTo(machine.conn); err != nil {
// // 		return err
// // 	}
// // 	// go to state 13
// // 	machine.state = sta13
// // 	// all is well
// // 	return nil
// // }
// //
// // func (machine *Machine) aa8(source byte, reason byte) error {
// // 	// send A-ABORT PDU (service provider source)
// // 	pdu := &abortPDU{source: source, reason: reason}
// // 	if err := pdu.writeTo(machine.conn); err != nil {
// // 		return err
// // 	}
// // 	// issue an A-P-ABORT indication (what does that mean, call a handler?)
// // 	// start timer
// // 	machine.startTimer()
// // 	// all is well
// // 	return nil
// // }
