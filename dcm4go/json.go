package dcm4go

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func encapsulatedToJSON(path string, encapsulated *Encapsulated) string {
	s := ",\"DataFragment\":["
	for _, fragment := range encapsulated.fragments {
		s += fragmentToJSON(path, fragment) + ","
	}
	return strings.TrimSuffix(s, ",") + "]"
}

func fragmentToJSON(path string, fragment *Fragment) string {
	return fmt.Sprintf("{\"BulkDataURI\":\"file:%s?offset=%d&length=%d\"}", path, fragment.offset, fragment.length)
}

func valuesToJSON(attribute *Attribute, format string) string {
	s := ""
	switch values := attribute.value.(type) {
	case []string:
		for _, value := range values {
			s += fmt.Sprintf(format+",", value)
		}
	case []uint16:
		for _, value := range values {
			s += fmt.Sprintf(format+",", value)
		}
	case []uint32:
		for _, value := range values {
			s += fmt.Sprintf(format+",", value)
		}
	case []uint64:
		for _, value := range values {
			s += fmt.Sprintf(format+",", value)
		}
	case []float32:
		for _, value := range values {
			s += fmt.Sprintf(format+",", value)
		}
	case []float64:
		for _, value := range values {
			s += fmt.Sprintf(format+",", value)
		}
	case [][2]uint16:
		for _, value := range values {
			s += fmt.Sprintf(format+",", value[0], value[1])
		}
	}
	return strings.TrimSuffix(s, ",")
}

func cleanUpNumberStrings(ins []string) []string {
	outs := make([]string, len(ins))
	for i := 0; i < len(ins); i++ {
		s := ins[i]
		// trim the leading + sign if any, JSON doesn't like it
		s = strings.TrimPrefix(s, "+")
		// add a leading zero, JSON doesn't like leading decimals
		if s[0] == '.' {
			s = "0" + s
		}
		outs[i] = s
	}
	return outs
}

func attributeToJSON(path string, attribute *Attribute) string {
	s := fmt.Sprintf("\"%08X\":{\"vr\":\"%s\"", attribute.tag, attribute.vr)
	if attribute.length > 0 {
		switch attribute.vr {
		case "AE", "AS", "CS", "DA", "DT", "LO", "SH", "TM", "UC", "UI", "UR":
			s += fmt.Sprintf(",\"Value\":[%s]", valuesToJSON(attribute, "%q"))
		case "AT":
			s += fmt.Sprintf(",\"Value\":[%s]", valuesToJSON(attribute, "\"%04x%04x\""))
		case "FD", "OD", "FL", "OF":
			s += fmt.Sprintf(",\"Value\":[%s]", valuesToJSON(attribute, "%f"))
		case "DS", "IS":
			// clean up strings representing numbers
			attribute.value = cleanUpNumberStrings(attribute.value.([]string))
			s += fmt.Sprintf(",\"Value\":[%s]", valuesToJSON(attribute, "%s"))
		case "LT", "ST", "UT":
			s += fmt.Sprintf(",\"Value\":[%q]", attribute.value.(string))
		case "PN":
			s += fmt.Sprintf(",\"Value\":[%s]", valuesToJSON(attribute, "{\"Alphabetic\":%q}"))
		case "SS", "US", "SL", "UL", "SV", "UV":
			s += fmt.Sprintf(",\"Value\":[%s]", valuesToJSON(attribute, "%d"))
		case "OB", "OL", "OV", "OW", "UN":
			s += pixelDataToJSON(path, attribute)
		case "SQ":
			s += fmt.Sprintf(",\"Value\":[%s]", sequenceToJSON(path, attribute.value.(*Sequence)))
		}
	}
	s += "}"
	return s
}

func pixelDataToJSON(path string, attribute *Attribute) string {
	switch v := attribute.value.(type) {
	case *Encapsulated:
		return encapsulatedToJSON(path, v)
	case []byte:
		if v == nil {
			return fmt.Sprintf(",\"BulkDataURI\":\"file:%s?offset=%d&length=%d\"", path, attribute.offset, attribute.length)
		}
		return fmt.Sprintf(",\"InlineBinary\":\"%s\"", base64.StdEncoding.EncodeToString(v))
	default:
		return ""
	}
}

func sequenceToJSON(path string, sequence *Sequence) string {
	s := ""
	for _, object := range sequence.objects {
		s += ObjectToJSON(path, object) + ","
	}
	return strings.TrimSuffix(s, ",")
}

// ObjectToJSON prints the objects as JSON
func ObjectToJSON(path string, objects ...*Object) string {
	s := ""
	for _, object := range objects {
		for _, attribute := range object.attributes {
			// group lengths are not to be encoded in JSON representation
			if toElement(attribute.tag) != 0x0000 {
				s += attributeToJSON(path, attribute) + ","
			}
		}
	}
	return "{" + strings.TrimSuffix(s, ",") + "}"
}
