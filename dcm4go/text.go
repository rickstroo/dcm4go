package dcm4go

import (
	"encoding/base64"
	"fmt"
)

func attributeToText(attribute *Attribute, prefix string) string {
	s := fmt.Sprintf("%s(%04X,%04X) %s ", prefix, toGroup(attribute.tag), toElement(attribute.tag), attribute.vr)
	if attribute.value != nil {
		switch attribute.vr {
		case "AE", "AS", "CS", "DA", "DS", "DT", "IS", "LO", "SH", "TM", "UC", "UI", "UR", "LT", "ST", "UT", "PN":
			s += fmt.Sprintf("%q", attribute.value)
		case "FD", "OD", "FL", "OF", "SS", "US", "SL", "UL", "SV", "UV":
			s += fmt.Sprintf("%v", attribute.value)
		case "SQ":
			s += sequenceToText(attribute.value.(*Sequence), prefix)
		case "AT":
			s += fmt.Sprintf("%v", attribute.value)
		case "OB", "OL", "OV", "OW", "UN":
			switch v := attribute.value.(type) {
			case []byte:
				buf := attribute.value.([]byte)
				if len(buf) > 0 && len(buf) < 1024 {
					s += fmt.Sprintf("[%s]", base64.StdEncoding.EncodeToString(buf))
				}
			case *Encapsulated:
				s += encapsulatedToText(v, prefix)
			}
		}
	}
	return s
}

func sequenceToText(sequence *Sequence, prefix string) string {
	s := ""
	for i, object := range sequence.objects {
		s += objectToText(object, fmt.Sprintf("%s%d>", prefix, i+1))
	}
	return s
}

func objectToText(object *Object, prefix string) string {
	s := ""
	for _, attribute := range object.attributes {
		s += "\n" + attributeToText(attribute, prefix)
	}
	return s
}

func encapsulatedToText(encapsulated *Encapsulated, prefix string) string {
	s := ""
	for i, fragment := range encapsulated.fragments {
		s += "\n" + fragmentToText(fragment, fmt.Sprintf("%s%d>", prefix, i+1))
	}
	return s
}

func fragmentToText(fragment *Fragment, prefix string) string {
	return fmt.Sprintf("%s %d,%d", prefix, fragment.offset, fragment.length)
}
