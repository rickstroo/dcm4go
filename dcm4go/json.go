package dcm4go

import "fmt"

func lengthToJSON(length uint32) string {
	if length == UndefinedLength {
		return "<undefined>"
	}
	return fmt.Sprintf("%d", length)
}

func tagToJSON(tag uint32) string {
	return fmt.Sprintf("(%08X)", tag)
}

func attributeToJSON(attribute *Attribute, prefix string) string {
	s := fmt.Sprintf("%stag=%s vr=%s off=%d len=%s", prefix, tagToJSON(attribute.tag), attribute.vr, attribute.offset, lengthToJSON(attribute.length))
	switch attribute.vr {
	case "AE", "AS", "CS", "DA", "DT", "LO", "SH", "TM", "UC", "UI", "UR", "LT", "ST", "UT", "PN":
		s += fmt.Sprintf(" value=%q\n", attribute.value)
	case "DS", "IS", "FD", "OD", "FL", "OF", "SS", "US", "SL", "UL", "SV", "UV":
		s += fmt.Sprintf(" value=%v\n", attribute.value)
	case "SQ":
		s += "\n" + sequenceToJSON(attribute.value.(*Sequence), prefix)
	case "AT", "OL", "OV", "OW", "UN":
		s += "\n"
	case "OB":
		switch v := attribute.value.(type) {
		case []byte:
			s += "\n"
		case *Encapsulated:
			s += "\n" + encapsulatedToJSON(v, prefix)
		}
	default:
	}
	return s
}

func sequenceToJSON(sequence *Sequence, prefix string) string {
	s := ""
	for i, item := 1, sequence.objects.Front(); item != nil; i, item = i+1, item.Next() {
		s += objectToJSON(item.Value.(*Object), fmt.Sprintf("%sitem#%d>", prefix, i))
	}
	return s
}

func objectToJSON(object *Object, prefix string) string {
	s := ""
	for item := object.attributes.Front(); item != nil; item = item.Next() {
		s += attributeToJSON(item.Value.(*Attribute), prefix)
	}
	return s
}

func encapsulatedToJSON(encapsulated *Encapsulated, prefix string) string {
	s := ""
	for i, item := 1, encapsulated.fragments.Front(); item != nil; i, item = i+1, item.Next() {
		s += fragmentToJSON(item.Value.(*Fragment), fmt.Sprintf("%sfrag#%d>", prefix, i))
	}
	return s
}

func fragmentToJSON(fragment *Fragment, prefix string) string {
	return fmt.Sprintf("%soffset=%d,length=%d\n", prefix, fragment.offset, fragment.length)
}
