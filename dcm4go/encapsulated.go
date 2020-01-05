package dcm4go

import "container/list"

// PixelData contains an ordered list of fragments
type Encapsulated struct {
	fragments *list.List
}

func (encapsulated *Encapsulated) add(fragment *Fragment) {
	encapsulated.fragments.PushBack(fragment)
}

func newEncapsulated() *Encapsulated {
	return &Encapsulated{list.New()}
}
