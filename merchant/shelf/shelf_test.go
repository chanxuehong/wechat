package shelf

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestShelfMarshal(t *testing.T) {
	expectBytes := []byte(`{
		"shelf_name":"测试货架",
		"shelf_banner":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2ibrWQn8zWFUh1YznsMV0XEiavFfLzDWYyvQOBBszXlMaiabGWzz5B2KhNn2IDemHa3iarmCyribYlZYyw/0",
		"shelf_data":{
			"module_infos":[
				{
					"eid":1,
					"group_info":{
						"group_id":50,
						"filter":{
							"count":2
						}
					}
				},
				{
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
				},
				{
					"eid":3,
					"group_info":{
						"group_id":52,
						"img":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5Jm64z4I0TTicv0TjN7Vl9bykUUibYKIOjicAwIt6Oy0Y6a1Rjp5Tos8tg/0"
					}
				},
				{
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
				},
				{
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
				}
			]
		}
	}`)

	_shelf := Shelf{
		Name:   "测试货架",
		Banner: "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2ibrWQn8zWFUh1YznsMV0XEiavFfLzDWYyvQOBBszXlMaiabGWzz5B2KhNn2IDemHa3iarmCyribYlZYyw/0",
	}
	_shelf.Info.ModuleInfos = []*Module{
		NewModule1(50, 2),
		NewModule2([]int64{49, 50, 51}),
		NewModule3(52, "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5Jm64z4I0TTicv0TjN7Vl9bykUUibYKIOjicAwIt6Oy0Y6a1Rjp5Tos8tg/0"),
		NewModule4([]Group{
			Group{
				GroupId: 49,
				Image:   "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0",
			},
			Group{
				GroupId: 50,
				Image:   "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5G1kdy3ViblHrR54gbCmbiaMnl5HpLGm5JFeENyO9FEZAy6mPypEpLibLA/0",
			},
			Group{
				GroupId: 52,
				Image:   "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0",
			},
		}),
		NewModule5([]int64{43, 44, 45, 46}, "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"),
	}

	b, err := json.Marshal(_shelf)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", _shelf, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", _shelf, b, want)
		}
	}
}
