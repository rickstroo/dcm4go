package dcm4go

import "fmt"

func attributeToString(attribute *Attribute, prefix string) string {
	s := fmt.Sprintf("%stag=(%04X,%04X) vr=%s off=%d len=%d", prefix, attribute.group, attribute.element, attribute.vr, attribute.offset, attribute.length)
	switch attribute.vr {
	case "AE", "AS", "CS", "DA", "DT", "LO", "SH", "TM", "UC", "UI", "UR", "LT", "ST", "UT", "PN":
		s += fmt.Sprintf(" value=%q\n", attribute.value)
	case "DS", "IS", "FD", "OD", "FL", "OF", "SS", "US", "SL", "UL", "SV", "UV":
		s += fmt.Sprintf(" value=%v\n", attribute.value)
	case "SQ":
		s += "\n" + sequenceToString(attribute.value.(*Sequence), prefix)
	case "AT", "OB", "OL", "OV", "OW", "UN":
		s += "\n"
	default:
	}
	return s
}

func sequenceToString(sequence *Sequence, prefix string) string {
	s := ""
	for i, item := 1, sequence.objects.Front(); item != nil; i, item = i+1, item.Next() {
		s += objectToString(item.Value.(*Object), fmt.Sprintf("%s%d>", prefix, i))
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
