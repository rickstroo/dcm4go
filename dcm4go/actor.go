package dcm4go

// import "log"
//
// // Actor is an actor
// type Actor struct {
// 	actionChan chan *Action
// }
//
// // Action is an Action
// type Action struct {
// 	message string
// }
//
// // StartActor starts an Actor
// func StartActor() *Actor {
// 	ch := make(chan *Action, 1)
// 	actor := &Actor{}
// 	go actor.actorLoop(ch)
// 	return actor
// }
//
// // DoSomething sends a message telling the actor to do something
// func (actor *Actor) DoSomething() {
// 	actor.actionChan <- &Action{message: "Hello World!"}
// }
//
// func (actor *Actor) actorLoop(ch <-chan *Action) {
// 	for {
// 		action := <-ch
// 		log.Printf("action is %v", action)
// 	}
// }
