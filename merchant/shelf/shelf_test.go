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

func TestShelfMarshal(t *testing.T) {
	expectBytes := []byte(`{
		"shelf_name":"测试货架",
		"shelf_banner":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2ibrWQn8zWFUh1YznsMV0XEiavFfLzDWYyvQOBBszXlMaiabGWzz5B2KhNn2IDemHa3iarmCyribYlZYyw/0",
		"shelf_info":{
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
	modules := make([]Module, 5)
	modules[0].InitToModule1(50, 2)
	modules[1].InitToModule2([]int64{49, 50, 51})
	modules[2].InitToModule3(52, "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5Jm64z4I0TTicv0TjN7Vl9bykUUibYKIOjicAwIt6Oy0Y6a1Rjp5Tos8tg/0")
	modules[3].InitToModule4([]Group{
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
	})
	modules[4].InitToModule5([]int64{43, 44, 45, 46}, "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0")
	_shelf.Info.ModuleInfos = modules

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
