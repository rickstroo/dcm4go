package dcm4go

import "fmt"

func toTag(group uint16, element uint16) uint32 {
	return uint32(group)<<16 | uint32(element)
}

func toGroup(tag uint32) uint16 {
	return uint16(tag >> 16)
}

func toElement(tag uint32) uint16 {
	return uint16(tag)
}

func tagToString(tag uint32) string {
	return fmt.Sprintf("(%04X,%04X)", toGroup(tag), toElement(tag))
}
