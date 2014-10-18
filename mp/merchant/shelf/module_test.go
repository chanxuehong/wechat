// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package shelf

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestMarshalAndInitFunc(t *testing.T) {
	var expectBytes []byte

	// 控件1 ====================================================================

	expectBytes = []byte(`{
		"eid":1,
    	"group_info":{
        	"group_id":50,
        	"filter":{
            	"count":4
        	}
    	}
	}`)

	var md1 Module
	md1.InitToModule1(50, 4)

	b, err := json.Marshal(md1)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", md1, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", md1, b, want)
		}
	}

	// 控件2 ====================================================================

	expectBytes = []byte(`{
		"eid":2,
		"group_infos":{
			"groups":[
				{
					"group_id":49
				}, 
				{
					"group_id":50
				}, 
				{
					"group_id":51
				}
			]
		}
	}`)

	var md2 Module
	md2.InitToModule2([]int64{49, 50, 51})

	b, err = json.Marshal(md2)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", md2, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", md2, b, want)
		}
	}

	// 控件3 ====================================================================

	expectBytes = []byte(`{
		"eid":3,
		"group_info":{
			"group_id":52, 
			"img":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5Jm64z4I0TTicv0TjN7Vl9bykUUibYKIOjicAwIt6Oy0Y6a1Rjp5Tos8tg/0"
		}
	}`)

	var md3 Module
	md3.InitToModule3(52, "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5Jm64z4I0TTicv0TjN7Vl9bykUUibYKIOjicAwIt6Oy0Y6a1Rjp5Tos8tg/0")

	b, err = json.Marshal(md3)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", md3, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", md3, b, want)
		}
	}

	// 控件4 ====================================================================

	expectBytes = []byte(`{
		"eid":4,
		"group_infos":{
			"groups":[
				{
					"group_id":49, 
					"img":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"
				}, 
				{
					"group_id":50, 
					"img":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5G1kdy3ViblHrR54gbCmbiaMnl5HpLGm5JFeENyO9FEZAy6mPypEpLibLA/0"
				}, 
				{
					"group_id":52, 
					"img":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"
				}
			]
		}
	}`)

	var md4 Module
	md4.InitToModule4([]Group{
		{
			49,
			"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0",
		},
		{
			50,
			"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5G1kdy3ViblHrR54gbCmbiaMnl5HpLGm5JFeENyO9FEZAy6mPypEpLibLA/0",
		},
		{
			52,
			"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0",
		},
	})

	b, err = json.Marshal(md4)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", md4, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", md4, b, want)
		}
	}

	// 控件5 ====================================================================

	expectBytes = []byte(`{
		"eid":5,
		"group_infos":{
			"groups":[
				{
					"group_id":43
				}, 
				{
					"group_id":44
				}, 
				{
					"group_id":45
				}, 
				{
					"group_id":46
				}
			], 
			"img_background":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"
		}
	}`)

	var md5 Module
	md5.InitToModule5(
		[]int64{43, 44, 45, 46},
		"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0",
	)

	b, err = json.Marshal(md5)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", md5, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", md5, b, want)
		}
	}
}
