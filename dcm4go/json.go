package dcm4go

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// these functions are used to create JSON representations of DICOM objects
// i had played around with using some of the builting JSON marshalling
// tools, but found i was having to jump through lots of hoops to make
// the output follow the DICOM specification.  one of the articles that i
// found useful is at http://choly.ca/post/go-json-marshalling/

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
	if attribute.value != nil {
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
		return fmt.Sprintf(",\"InlineBinary\":\"%s\"", base64.StdEncoding.EncodeToString(v))
	case *Fragment:
		return fmt.Sprintf(",\"BulkDataURI\":\"file:%s?offset=%d&length=%d\"", path, v.offset, v.length)
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
func ObjectToJSON(path string, object *Object) string {
	s := ""
	for _, attribute := range object.attributes {
		// group lengths are not to be encoded in JSON representation
		if toElement(attribute.tag) != 0x0000 {
			s += attributeToJSON(path, attribute) + ","
		}
	}
	return "{" + strings.TrimSuffix(s, ",") + "}"
}

// ToJSON returns a JSON representation of an object
func ToJSON(object *Object) ([]byte, error) {
	jsonObject, err := prepareJSONObject(object)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(jsonObject, "", " ")
}

// JSONObject is a struct optimized for JSON marshalling
type JSONObject struct {
	Attributes map[string]*JSONAttribute
}

// JSONAttribute is a struct optimized for JSON marshalling
type JSONAttribute struct {
	VR    string
	Value interface{}
}

func prepareJSONObject(object *Object) (*JSONObject, error) {
	jsonObject := &JSONObject{make(map[string]*JSONAttribute)}
	for _, attribute := range object.attributes {
		// as per the standard, group length attributes are ommited from JSON output
		if toElement(attribute.tag) != 0x0000 {
			jsonAttribute, err := prepareJSONAttribute(attribute)
			if err != nil {
				return nil, err
			}
			jsonObject.Attributes[fmt.Sprintf("%08X", attribute.tag)] = jsonAttribute
		}
	}
	return jsonObject, nil
}

func prepareJSONAttribute(attribute *Attribute) (*JSONAttribute, error) {
	jsonAttribute := &JSONAttribute{}
	jsonAttribute.VR = attribute.vr
	jsonValue, err := prepareJSONValue(attribute)
	if err != nil {
		return nil, err
	}
	jsonAttribute.Value = jsonValue
	return jsonAttribute, nil
}

func prepareJSONValue(attribute *Attribute) (interface{}, error) {
	switch attribute.vr {
	case "DS":
		// the standard says to format DS values as floats, not strings
		return convertStringsToFloats(attribute.value)
	case "IS":
		// the standards says to format IS values as ints, not strings
		return convertStringsToInts(attribute.value)
	case "PN":
		return convertStringsToAlphabetics(attribute.value)
	case "OB":
		switch attribute.value.(type) {
		case []byte:
			// do nothing, json converts byte slice to base64
		case []*Fragment:
			// todo
		}
	}
	return attribute.value, nil
}

// converts a slice of strings to floats
func convertStringsToFloats(value interface{}) ([]float64, error) {
	strings, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("value not of type []string")
	}
	floats := make([]float64, 0, len(strings))
	for _, string := range strings {
		float, err := strconv.ParseFloat(string, 64)
		if err != nil {
			fmt.Printf("while parsing float, caught %v, will skip\n", err)
		} else {
			floats = append(floats, float)
		}
	}
	return floats, nil
}

// converts a slice of strings to ints
func convertStringsToInts(value interface{}) ([]int64, error) {
	strings, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("value not of type []string")
	}
	ints := make([]int64, len(strings))
	for i, string := range strings {
		int, err := strconv.ParseInt(string, 10, 64)
		if err != nil {
			return nil, err
		}
		ints[i] = int
	}
	return ints, nil
}

// Alphabetic as per the standard
type Alphabetic struct {
	Alphabetic string
}

// converts a slice of strings to alphabetics, as per the standard
func convertStringsToAlphabetics(value interface{}) ([]*Alphabetic, error) {
	strings, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("value not of type []string")
	}
	alphabetics := make([]*Alphabetic, len(strings))
	for i, string := range strings {
		alphabetics[i] = &Alphabetic{string}
	}
	return alphabetics, nil
}
