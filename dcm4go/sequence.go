package core

import "container/list"

// Sequence contains an ordered list of objects
type Sequence struct {
	objects *list.List
}

func (sequence *Sequence) add(object *Object) {
	sequence.objects.PushBack(object)
}

func newSequence() *Sequence {
	return &Sequence{list.New()}
}
