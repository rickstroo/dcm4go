package dcm4go

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAttributeToJSON(t *testing.T) {
	attribute := &Attribute{FileMetaInformationGroupLengthTag, "UL", 8, 0, []uint32{0x01, 0x02}}
	json, err := json.MarshalIndent(attribute, "", "  ")
	if err != nil {
		t.Error(err)
	}
	if err := testStringEquals("{\"tag\":\"00020000\",\"vr\":\"UL\",\"len\":8,\"off\":0,\"value\":[1,2]}", string(json)); err != nil {
		t.Error(err)
	}
	fmt.Printf("attribute is %v\n", string(json))
}
