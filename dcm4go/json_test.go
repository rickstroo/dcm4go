package dcm4go

// func TestAttributeToJSON(t *testing.T) {
// 	attribute := &Attribute{0x00020001, "UL", 8, 0, []uint32{0x01, 0x02}}
// 	json, err := json.Marshal(attribute)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if err := testStringEquals("{\"Value\":[1,2],\"vr\":\"UL\"}", string(json)); err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("attribute is %v\n", string(json))
// }
//
// func TestMakeJSON(t *testing.T) {
// 	type a struct {
// 		VR    string
// 		Value interface{}
// 	}
// 	m := make(map[uint32]*a)
// 	m[1] = &a{"UL", []uint32{1, 2}}
// 	m[2] = &a{"SH", []string{"tom", "dick", "harry"}}
// 	json, err := json.Marshal(m)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Printf("json is %s\n", string(json))
// }
