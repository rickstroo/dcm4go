package dcm4go

import (
	"encoding/base64"
	"fmt"
)

func lengthToString(length uint32) string {
	if length == UndefinedLength {
		return "<undefined>"
	}
	return fmt.Sprintf("%d", length)
}

func attributeToString(attribute *Attribute, prefix string) string {
	s := fmt.Sprintf("%stag=%s vr=%s off=%d len=%s", prefix, tagToString(attribute.tag), attribute.vr, attribute.offset, lengthToString(attribute.length))
	if attribute.length == 0 {
		return s + "\n"
	}
	switch attribute.vr {
	case "AE", "AS", "CS", "DA", "DT", "LO", "SH", "TM", "UC", "UI", "UR", "LT", "ST", "UT", "PN":
		s += fmt.Sprintf(" value=%q\n", attribute.value)
	case "DS", "IS", "FD", "OD", "FL", "OF", "SS", "US", "SL", "UL", "SV", "UV":
		s += fmt.Sprintf(" value=%v\n", attribute.value)
	case "SQ":
		s += "\n" + sequenceToString(attribute.value.(*Sequence), prefix)
	case "AT", "OL", "OV", "OW", "UN":
		s += "\n"
	case "OB":
		switch v := attribute.value.(type) {
		case []byte:
			buf := attribute.value.([]byte)
			if len(buf) > 0 {
				s += fmt.Sprintf(" value=[%s]\n", base64.StdEncoding.EncodeToString(buf))
			} else {
				s += "\n"
			}
		case *Encapsulated:
			s += "\n" + encapsulatedToString(v, prefix)
		default:
			s += "\n"
		}
	default:
		s += "\n"
	}
	return s
}

func sequenceToString(sequence *Sequence, prefix string) string {
	s := ""
	for i, item := 1, sequence.objects.Front(); item != nil; i, item = i+1, item.Next() {
		s += objectToString(item.Value.(*Object), fmt.Sprintf("%sitem#%d>", prefix, i))
	}
	return s
}

func objectToString(object *Object, prefix string) string {
	s := ""
	for item := object.attributes.Front(); item != nil; item = item.Next() {
		s += attributeToString(item.Value.(*Attribute), prefix)
	}
	return s
}

func encapsulatedToString(encapsulated *Encapsulated, prefix string) string {
	s := ""
	for i, item := 1, encapsulated.fragments.Front(); item != nil; i, item = i+1, item.Next() {
		s += fragmentToString(item.Value.(*Fragment), fmt.Sprintf("%sfrag#%d>", prefix, i))
	}
	return s
}

func fragmentToString(fragment *Fragment, prefix string) string {
	return fmt.Sprintf("%soff=%d,len=%d\n", prefix, fragment.offset, fragment.length)
}
