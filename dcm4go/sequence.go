package dcm4go

// Sequence contains an ordered list of objects
type Sequence struct {
	objects []*Object
}

// // MarshalJSON returns sequence as JSON
// func (sequence *Sequence) MarshalJSON() ([]byte, error) {
// 	// return json.Marshal(&struct {
// 	// 	Objects []*Object `json:"objects"`
// 	// }{
// 	// 	Objects: sequence.objects,
// 	// })
// 	return json.Marshal(sequence.objects)
// }

func (sequence *Sequence) add(object *Object) {
	sequence.objects = append(sequence.objects, object)
}

func newSequence() *Sequence {
	return &Sequence{make([]*Object, 0, 10)}
}
