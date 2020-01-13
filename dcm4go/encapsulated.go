package dcm4go

// Encapsulated contains an ordered list of fragments
type Encapsulated struct {
	fragments []*Fragment
}

func (encapsulated *Encapsulated) add(fragment *Fragment) {
	encapsulated.fragments = append(encapsulated.fragments, fragment)
}

func newEncapsulated() *Encapsulated {
	return &Encapsulated{make([]*Fragment, 0, 10)}
}
