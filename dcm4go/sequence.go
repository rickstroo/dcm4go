package dcm4go

// Sequence contains an ordered list of objects
type Sequence struct {
	objects []*Object
}

func (sequence *Sequence) add(object *Object) {
	sequence.objects = append(sequence.objects, object)
}

func newSequence() *Sequence {
	return &Sequence{make([]*Object, 0, 10)}
}
