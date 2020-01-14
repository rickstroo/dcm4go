package dcm4go

// Attribute contains all the properties of a DICOM attribute
type Attribute struct {
	tag    uint32
	vr     string
	length uint32
	offset uint32
	value  interface{}
}

// String returns attribute as a string
func (attribute *Attribute) String() string {
	return attributeToString(attribute, "")
}

// // MarshalJSON returns attribute as JSON
// func (attribute *Attribute) MarshalJSON() ([]byte, error) {
// 	//	if attribute.length == 0 {
// 	//		return []byte("\"\""), nil
// 	//	}
// 	value, err := prepare(attribute)
// 	if err != nil {
// 		return nil, err
// 	}
// 	switch v := attribute.value.(type) {
// 	case []byte:
// 		if v != nil {
// 			return json.Marshal(&struct {
// 				Value interface{} `json:"InlineBinary"`
// 				VR    string      `json:"vr"`
// 			}{
// 				Value: v,
// 				VR:    attribute.vr,
// 			})
// 		}
// 		return json.Marshal(&struct {
// 			Value string `json:"BulkDataURI"`
// 			VR    string `json:"vr"`
// 		}{
// 			Value: "TODO", // fmt.Sprintf("file:%s?offset=%d&length=%d", "path", attribute.offset, attribute.length),
// 			VR:    attribute.vr,
// 		})
// 	}
// 	return json.Marshal(&struct {
// 		Value interface{} `json:"Value"`
// 		VR    string      `json:"vr"`
// 	}{
// 		Value: value,
// 		VR:    attribute.vr,
// 	})
// }
//
// // prepares the attribute for JSON marshalling as required
// func prepare(attribute *Attribute) (interface{}, error) {
// 	switch attribute.vr {
// 	case "DS":
// 		// the standard says to format DS values as floats, not strings
// 		return convertStringsToFloats(attribute.value)
// 	case "IS":
// 		// the standards says to format IS values as ints, not strings
// 		return convertStringsToInts(attribute.value)
// 	case "PN":
// 		return convertStringsToAlphabetics(attribute.value)
// 	case "OB":
// 		switch attribute.value.(type) {
// 		case []byte:
// 			// do nothing, json converts byte slice to base64
// 		case []*Fragment:
// 			// todo
// 		}
// 	}
// 	return attribute.value, nil
// }
//
// // converts a slice of strings to floats
// func convertStringsToFloats(value interface{}) ([]float64, error) {
// 	strings, ok := value.([]string)
// 	if !ok {
// 		return nil, fmt.Errorf("value not of type []string")
// 	}
// 	floats := make([]float64, 0, len(strings))
// 	for _, string := range strings {
// 		float, err := strconv.ParseFloat(string, 64)
// 		if err != nil {
// 			fmt.Printf("while parsing float, caught %v, will skip\n", err)
// 		} else {
// 			floats = append(floats, float)
// 		}
// 	}
// 	return floats, nil
// }
//
// // converts a slice of strings to ints
// func convertStringsToInts(value interface{}) ([]int64, error) {
// 	strings, ok := value.([]string)
// 	if !ok {
// 		return nil, fmt.Errorf("value not of type []string")
// 	}
// 	ints := make([]int64, len(strings))
// 	for i, string := range strings {
// 		int, err := strconv.ParseInt(string, 10, 64)
// 		if err != nil {
// 			return nil, err
// 		}
// 		ints[i] = int
// 	}
// 	return ints, nil
// }
//
// // Alphabetic as per the standard
// type Alphabetic struct {
// 	Alphabetic string
// }
//
// // converts a slice of strings to alphabetics, as per the standard
// func convertStringsToAlphabetics(value interface{}) ([]*Alphabetic, error) {
// 	strings, ok := value.([]string)
// 	if !ok {
// 		return nil, fmt.Errorf("value not of type []string")
// 	}
// 	alphabetics := make([]*Alphabetic, len(strings))
// 	for i, string := range strings {
// 		alphabetics[i] = &Alphabetic{string}
// 	}
// 	return alphabetics, nil
// }

// simple check for out of bounds
func checkIndex(index int, length int) error {
	if index < 0 || index >= length {
		return ErrIndexOutOfBounds
	}
	return nil
}

// AsLong returns attribute value as a long
func (attribute *Attribute) asLong(index int) (uint32, error) {
	longs, ok := attribute.value.([]uint32)
	if !ok {
		return 0, ErrWrongType
	}
	if err := checkIndex(index, len(longs)); err != nil {
		return 0, err
	}
	return longs[index], nil
}

// AsString returns attribute value as a string
func (attribute *Attribute) asString(index int) (string, error) {
	strings, ok := attribute.value.([]string)
	if !ok {
		return "", ErrWrongType
	}
	if err := checkIndex(index, len(strings)); err != nil {
		return "", err
	}
	return strings[index], nil
}
