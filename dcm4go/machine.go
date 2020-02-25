// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package dcm4go

import (
	"io"
	"log"
	"net"
)

// Machine implements the DICOM Upper Layer Protocol TCP/IP State Machine
// It is a little complicated because the standard mixes the state machines
// for the acceptor and requestor of associations.  I'm going to try building
// the machine as per the standard.  Then, I expect I'll build a couple of
// wrappers, one for the Acceptor and one for the Requestor, that provide
// the transport connections, the handlers and an appropriate start state for
// each.
type Machine struct {
	addr                      string
	listener                  net.Listener
	conn                      net.Conn
	state                     *State
	acceptor                  bool
	serviceUserInitiatedAbort bool
	aeTitle                   string
	pcs                       []*PC
	assocACPDU                *AssocACPDU
}

func (machine *Machine) startTimer() {
}

func (machine *Machine) stopTimer() {
}

func (machine *Machine) run() error {
	for {
		log.Printf("entering state %s", machine.state.name)
		if err := machine.state.action(machine); err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}
	return nil
}

// State of the state machine
type State struct {
	name        string
	description string
	action      func(*Machine) error
}

// sta1 is the idle or start state for
var sta1 = &State{
	name:        "sta1",
	description: "Idle",
	action:      st1,
}

// sta2 is the idle or start state for
var sta2 = &State{
	name:        "sta2",
	description: "Transport connection open (Awaiting A-ASSOCIATE-RQ PDU)",
	action:      st2,
}

// sta3 is the idle or start state for
var sta3 = &State{
	name:        "sta3",
	description: "Awaiting local A-ASSOCIATE response primitive (from local user)",
	action:      st3,
}

// sta4 is the idle or start state for
var sta4 = &State{
	name:        "sta4",
	description: "Awaiting transport connection opening to complete (from local transport service)",
	action:      st4,
}

// sta5 is the idle or start state for
var sta5 = &State{
	name:        "sta5",
	description: "Awaiting A-ASSOCIATE-AC or A-ASSOCIATE-RJ PDU",
	action:      st5,
}

// sta6 is the idle or start state for
var sta6 = &State{
	name:        "sta6",
	description: "Association established and ready for data transfer",
	action:      st6,
}

// sta7 is the idle or start state for
var sta7 = &State{
	name:        "sta7",
	description: "Awaiting A-RELEASE-RP PDU",
	action:      st7,
}

// sta8 is the idle or start state for
var sta8 = &State{
	name:        "sta8",
	description: "Awaiting local A-RELEASE response primitive (from local user)",
	action:      st8,
}

// sta9 is the idle or start state for
var sta9 = &State{
	name:        "sta9",
	description: "Release collision requestor side; awaiting A-RELEASE response (from local user)",
	action:      st9,
}

// sta10 is the idle or start state for
var sta10 = &State{
	name:        "sta10",
	description: "Release collision acceptor side; awaiting A-RELEASE-RP PDU",
	action:      st10,
}

// sta11 is the idle or start state for
var sta11 = &State{
	name:        "sta11",
	description: "Release collision requestor side; awaiting A-RELEASE-RP PDU",
	action:      st11,
}

// sta12 is the idle or start state for
var sta12 = &State{
	name:        "sta12",
	description: "Release collision acceptor side, awaitint A-RELEASE reasponse primitive (from local user)",
	action:      st12,
}

// sta13 is the idle or start state for
var sta13 = &State{
	name:        "sta13",
	description: "Awaiting Transport Connection Close Indication (Association no longer exists)",
	action:      st13,
}

func st1(machine *Machine) error {
	if machine.acceptor {
		return machine.ae5()
	}
	return machine.ae1()
}

func st2(machine *Machine) error {

	// wait for pdu
	pdu, err := readPDU(machine.conn)
	if err != nil {
		return nil
	}

	// if it is an A-ASSOCIATE-RQ PDU, handle it
	if pdu.pduType == aAssociateRQPDU {
		return machine.ae6(pdu)
	}

	// otherwise, send abort
	return machine.aa1()
}

// Awaiting local A-ASSOCIATE response primitive (from local user)
func st3(machine *Machine) error {

	// hmm, in ae 6, just before this, i am supposed to tell the local user
	// that we accepted an association, and then we go to this state
	// to wait for the local user to respond.  i wonder how that is best
	// implemented?  in go, we could use channels to communicate between
	// the local user and the state machine.  is that really how this is
	// all supposed to work?

	return machine.ae7()
}

func st4(machine *Machine) error {
	return nil
}

// this is the start sate for a requestor
func st5(machine *Machine) error {

	// wait for pdu
	pdu, err := readPDU(machine.conn)
	if err != nil {
		return err
	}

	// if A-ASSOCIATE-AC, accept
	if pdu.pduType == aAssociateACPDU {
		return machine.ae3(pdu)
	}

	// if A-ASSOCIATE-RJ, reject
	if pdu.pduType == aAssociateRJPDU {
		return machine.ae4(pdu)
	}

	// otherwise, abort
	return machine.aa8(sourceServiceProviderInitiatedAbort, reasonUnexpectedPDU)
}

func st6(machine *Machine) error {

	// hmm, if we are a requestor, here's where we wait
	// for a P-DATA request primitive from the local user
	// and we go to dt1/sta6
	// or we get an abort request form the local user and we go to aa1
	// or the connection is closed and we go to aa4/sta1
	// otherwise, we are an acceptor and we wait for a PDU
	// similar for a release request from local user, we go to ar1/sta7

	// wait for PDU
	pdu, err := readPDU(machine.conn)
	if err != nil {
		return err
	}

	// if its a release request, we're good
	if pdu.pduType == aReleaseRQPDU {
		return machine.ar2(pdu)
	}

	// if its an abort request, we're good
	if pdu.pduType == aAbortPDU {
		return machine.aa3(pdu)
	}

	// in this state, we can also receive a request
	// from the local user to abort

	// if P-DATA-TF PDU, we're good
	if pdu.pduType == pDataTFPDU {
		return machine.dt2(pdu)
	}

	// otherwise, abort
	return machine.aa8(sourceServiceProviderInitiatedAbort, reasonUnexpectedPDU)
}

// this is where a user initiated release starts
func st7(machine *Machine) error {
	return nil
}

func st8(machine *Machine) error {
	return nil
}

func st9(machine *Machine) error {
	return nil
}

func st10(machine *Machine) error {
	return nil
}

func st11(machine *Machine) error {
	return nil
}

func st12(machine *Machine) error {
	return nil
}

func st13(machine *Machine) error {
	return io.EOF
}

func (machine *Machine) ae1() error {
	// issue connection request
	// go to state 4
	machine.state = sta4
	// all is well
	return nil
}

func (machine *Machine) ae2() error {
	// send A-ASSOCIATE-RQ PDU
	// go to state 5
	machine.state = sta5
	// all is well
	return nil
}

func (machine *Machine) ae3(pdu *pdu) error {
	// call the associate accept handler
	// go to state 4
	machine.state = sta4
	// all is well
	return nil
}

func (machine *Machine) ae4(pdu *pdu) error {
	// call the associate reject handler
	// close the connection
	// go to state 1
	machine.state = sta1
	// all is well
	return nil
}

func (machine *Machine) ae5() error {
	// issue transport connection response primitive (what does that mean?)
	conn, err := machine.listener.Accept()
	if err != nil {
		return err
	}
	machine.conn = conn
	// start ARTIM timer
	// go to state 2
	machine.state = sta2
	// all is well
	return nil
}

func (machine *Machine) ae6(pdu *pdu) error {

	// read the associate request
	assocRQPDU, err := readAssocRQPDU(pdu)
	if err != nil {
		return err
	}
	log.Printf("assocRQPDU is %v\n", assocRQPDU)

	// stop timer
	machine.stopTimer()

	// build an ae
	ae := NewAE(machine.aeTitle)

	// attempt to negotiate an association
	assocACPDU, assocRJPDU, err := negotiateAssoc(assocRQPDU, ae, machine.pcs)
	if err != nil {
		return err
	}
	if assocACPDU == nil {
		log.Printf("assocACPDU is nil")
	} else {
		log.Printf("assocACPDU is %v\n", assocACPDU)
	}
	if assocRJPDU == nil {
		log.Printf("assocRJPDU is nil")
	} else {
		log.Printf("assocRJPDU is %v\n", assocRJPDU)
	}

	// was association rejected
	if assocRJPDU != nil {

		// wonder why i write the reject pdu in this action
		// but i wait to write the accept pdu in another action?

		// write the associate reject pdu
		if err := assocRJPDU.Write(machine.conn); err != nil {
			return err
		}

		// state the timer
		machine.startTimer()

		// goto state 6
		machine.state = sta6

		// all is well
		return nil
	}

	// issue A-ASSOCIATE indication primitive
	// not sure what this means
	// perhaps its a call to the DIMSE layer telling it we have an association

	// remember the associate AC PDU
	machine.assocACPDU = assocACPDU

	// go to state 3
	machine.state = sta3

	// all is well
	return nil
}

func (machine *Machine) ae7() error {

	// otherwise, write the associate accept pdu
	if err := machine.assocACPDU.Write(machine.conn); err != nil {
		return err
	}

	// go to state 6
	machine.state = sta6

	// all is well
	return nil
}

func (machine *Machine) ae8() error {
	return nil
}

func (machine *Machine) dt1() error {
	return nil
}

func (machine *Machine) dt2(pdu *pdu) error {
	return nil
}

func (machine *Machine) ar1() error {
	// send A-RELEASE-RQ PDU
	// next state is sta7
	return nil
}

func (machine *Machine) ar2(pdu *pdu) error {
	// issue A-RELEASE indication primitive
	// next state is sta8
	return nil
}

func (machine *Machine) ar3() error {
	// issue A-RELEASE confirmation primitive
	// close transport connection
	// next state is sta1
	return nil
}

func (machine *Machine) ar4() error {
	// issue A-RELEASE-RQ PDU
	// start ARTIM timer
	machine.startTimer()
	// next state is sta13
	machine.state = sta13
	// all is well
	return nil
}

func (machine *Machine) ar5() error {
	// stop ARTIM timer
	machine.stopTimer()
	// next state is sta1
	machine.state = sta1
	// all is well
	return nil
}

func (machine *Machine) ar6() error {
	// issue P-DATA indication
	// next state is sta7
	machine.state = sta7
	// all is well
	return nil
}

func (machine *Machine) ar7() error {
	// issue P-DATA-TF PDU
	// next state is sta8
	machine.state = sta8
	// all is well
	return nil
}

func (machine *Machine) ar8() error {
	// issue A-RELEASE indication (release condition)
	// if association requestor, next state is sta9
	// otherwise, next state is sta10
	if machine.acceptor {
		machine.state = sta10
	} else {
		machine.state = sta12
	}
	// all is well
	return nil
}

func (machine *Machine) ar9() error {
	// send A-RELEASE-RP PDU
	// next state is sta11
	machine.state = sta11
	// all is well
	return nil
}

func (machine *Machine) ar10() error {
	// issue A-RELEASE confirmation primitive
	// next state is sta12
	machine.state = sta12
	// all is well
	return nil
}

func (machine *Machine) aa1() error {
	// send A-ABORT PDU (service-user source) and start (or restart if already started) ARTIM timer
	machine.startTimer()
	// next state is sta13
	machine.state = sta13
	// all is well
	return nil
}

func (machine *Machine) aa2() error {
	// start timer
	machine.startTimer()
	// close connection
	if err := machine.conn.Close(); err != nil {
		return err
	}
	// go to state 1
	machine.state = sta1
	// all is well
	return nil
}

func (machine *Machine) aa3(pdu *pdu) error {
	// if (service-user initiated abort)
	// issue A-ABORT indication and close transport connection
	// otherwise (service-provider initiated abort)
	// issue A-P-ABORT indication and close transport connection
	if machine.serviceUserInitiatedAbort {
		// issue A-ABORT indication
	} else { // if service provider initiated abort
		// issue A-P-ABORT indication and close transport connection
	}
	// close transport connection
	if err := machine.conn.Close(); err != nil {
		return err
	}
	// all is well
	return nil
}

func (machine *Machine) aa4() error {
	// issue A-P-ABORT indication primitive
	// go to state 1
	machine.state = sta1
	// all is well
	return nil
}

func (machine *Machine) aa5() error {
	// stop the timer
	machine.stopTimer()
	// go to state 1
	machine.state = sta1
	// all is well
	return nil
}

func (machine *Machine) aa6(pdu *pdu) error {
	// ignore the PDU (just skip the rest of the bytes)
	skipBytes(pdu, pdu.pduLength)
	// go to state 13
	machine.state = sta13
	// all is well
	return nil
}

func (machine *Machine) aa7(source byte, reason byte) error {
	// send A-ABORT PDU
	pdu := &AbortPDU{source: source, reason: reason}
	if err := pdu.Write(machine.conn); err != nil {
		return err
	}
	// go to state 13
	machine.state = sta13
	// all is well
	return nil
}

func (machine *Machine) aa8(source byte, reason byte) error {
	// send A-ABORT PDU (service provider source)
	pdu := &AbortPDU{source: source, reason: reason}
	if err := pdu.Write(machine.conn); err != nil {
		return err
	}
	// issue an A-P-ABORT indication (what does that mean, call a handler?)
	// start timer
	machine.startTimer()
	// all is well
	return nil
}
