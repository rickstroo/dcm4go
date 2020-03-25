// Copyright 2020 Rick Stroobosscher.  All rights reserved.

package main

func main() {}

// import "log"
//
// // private api and implementation
//
// type actor struct {
// 	ac chan func()   // the action channel
// 	qc chan struct{} // the quit channel
// }
//
// func (a *actor) loop() {
// 	log.Printf("about to enter loop")
// 	for {
// 		log.Printf("waiting for event")
// 		select {
// 		case f := <-a.ac:
// 			log.Printf("received event, about to call function")
// 			f()
// 			log.Printf("called function")
// 		case <-a.qc:
// 			log.Printf("received quit, about to return")
// 			return
// 		}
// 	}
// }
//
// func (a *actor) consumeEvent(e *Event) (int, error) {
// 	log.Printf("consumeEvent running")
// 	return 1, nil
// }
//
// // func (a *actor) handleRequest(r *Request) (*Response, error) {
// // 	return nil, nil
// // }
//
// // public api
//
// // Actor handles events and requests
// type Actor interface {
// 	Start()
// 	SendEvent(*Event) int
// 	//	HandleRequest(*Request) (*Response, error)
// 	Stop()
// }
//
// // Event is an event
// type Event struct{}
//
// // // Request is a request
// // type Request struct{}
// //
// // // Response is a response
// // type Response struct{}
//
// func (a *actor) SendEvent(e *Event) (n int) {
// 	log.Printf("about to send event")
// 	a.ac <- func() {
// 		n, _ = a.consumeEvent(e)
// 		return
// 	}
// 	log.Printf("returned from sending event")
// 	return
// }
//
// // func (a *actor) HandleRequest(rq *Request) (rp *Response, err error) {
// // 	done := make(chan struct{})
// // 	a.ac <- func() {
// // 		defer close(done)
// // 		rp, err = a.handleRequest(rq)
// // 	}
// // 	return
// // }
//
// func (a *actor) Start() {
// 	go a.loop()
// }
//
// func (a *actor) Stop() {
// 	close(a.qc)
// }
//
// // NewActor creates an actor
// func NewActor() Actor {
// 	a := &actor{
// 		ac: make(chan func()),
// 		qc: make(chan struct{}),
// 	}
// 	return a
// }
//
// func main() {
// 	actor := NewActor()
// 	actor.Start()
// 	log.Printf("about to call actor.SendEvent")
// 	n := actor.SendEvent(&Event{})
// 	log.Printf("returned from actor.SendEvent, value is %v", n)
// 	actor.Stop()
// }
