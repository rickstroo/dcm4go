package dcm4go

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func attributeToString(attribute *attribute, prefix string) string {
	s := fmt.Sprintf("{%stag:0x%08X,vr:%s", prefix, attribute.tag, attribute.vr)
	if attribute.value != nil {
		switch attribute.vr {
		case "AE", "AS", "CS", "DA", "DT", "LO", "SH", "TM", "UC", "UI", "UR", "LT", "ST", "UT", "PN":
			s += fmt.Sprintf(",val:%q", attribute.value)
		case "DS", "IS", "FD", "OD", "FL", "OF", "SS", "US", "SL", "UL", "SV", "UV":
			s += fmt.Sprintf(",val:%v", attribute.value)
		case "SQ":
			s += sequenceToString(attribute.value.(*sequence), prefix)
		case "AT":
			s += fmt.Sprintf(",val:%v", attribute.value)
		case "OB", "OL", "OV", "OW", "UN":
			switch v := attribute.value.(type) {
			case []byte:
				buf := attribute.value.([]byte)
				if len(buf) > 0 && len(buf) < 1024 {
					s += fmt.Sprintf(",value:[%s]", base64.StdEncoding.EncodeToString(buf))
				}
			case *encapsulated:
				s += encapsulatedToString(v, prefix)
			}
		}
	}
	return s + "}"
}

func sequenceToString(sequence *sequence, prefix string) string {
	s := "["
	for i, object := range sequence.objects {
		s += objectToString(object, fmt.Sprintf("%sitem#%d>", prefix, i+1)) + ","
	}
	s = strings.TrimSuffix(s, ",")
	return s + "]"
}

func objectToString(object *Object, prefix string) string {
	s := "["
	for _, attribute := range object.attributes {
		s += attributeToString(attribute, prefix) + ","
	}
	s = strings.TrimSuffix(s, ",")
	return s + "]"
}

func encapsulatedToString(encapsulated *encapsulated, prefix string) string {
	s := ""
	for i, fragment := range encapsulated.fragments {
		s += fragmentToString(fragment, fmt.Sprintf("%sfrag#%d>", prefix, i+1))
	}
	return s
}

func fragmentToString(fragment *fragment, prefix string) string {
	return fmt.Sprintf("{%s,off:%d,len:%d}", prefix, fragment.offset, fragment.length)
}
