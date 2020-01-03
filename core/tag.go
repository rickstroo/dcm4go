package core

// Tag uniquely identifies a DICOM attribute
type Tag struct {
	group   uint16
	element uint16
}

func toTag(group uint16, element uint16) *Tag {
	return &Tag{group, element}
}
