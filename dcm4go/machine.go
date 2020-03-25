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
		m.nr = &networkReader{r: m.conn, dc: m.deedChan}
		go m.nr.run()
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
	nr            *networkReader
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
				d := &deed{e: evt19}
				log.Printf("created deed %s", d)
				nr.dc <- d
				break
			}
			log.Printf("reached EOF while reading pdus")
			break
		}
		log.Printf("read pdu %s", p)
		d := makePDUDeed(p)
		log.Printf("created deed %s", d)
		nr.dc <- d
	}
}
