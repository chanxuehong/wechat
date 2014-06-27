package group

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestGroupProductModifyRequestNewAndAppend(t *testing.T) {
	expectBytes := []byte(`{
		"group_id":28,
		"product":[
			{
				"product_id":"pDF3iY-CgqlAL3k8Ilz-6sj0UYpk",
				"mod_action":1
			},
			{
				"product_id":"pDF3iY-RewlAL3k8Ilz-6sjsepp9",
				"mod_action":0
			}
		]
	}`)

	request := NewGroupProductModifyRequest(
		28,
		[]string{"pDF3iY-CgqlAL3k8Ilz-6sj0UYpk"},
		[]string{"pDF3iY-RewlAL3k8Ilz-6sjsepp9"},
	)

	b, err := json.Marshal(request)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", request, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", request, b, want)
		}
	}

	request = NewGroupProductModifyRequest(28, nil, nil)
	request.AddProduct("pDF3iY-CgqlAL3k8Ilz-6sj0UYpk")
	request.DeleteProduct("pDF3iY-RewlAL3k8Ilz-6sjsepp9")

	b, err = json.Marshal(request)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", request, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", request, b, want)
		}
	}
}
