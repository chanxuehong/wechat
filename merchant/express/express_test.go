// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package express

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestGroupProductModifyRequestNewAndAppend(t *testing.T) {
	expectBytes := []byte(`{
		"Name":"testexpress",
		"Assumer":0,
		"Valuation":0,
		"TopFee":[
			{
				"Type":10000027,
				"Normal":{
					"StartStandards":1,
					"StartFees":2,
					"AddStandards":3,
					"AddFees":1
				},
				"Custom":[
					{
						"StartStandards":1,
						"StartFees":100,
						"AddStandards":1,
						"AddFees":3,
						"DestCountry":"中国",
						"DestProvince":"广东省",
						"DestCity":"广州市"
					}
				]
			},
			{
				"Type":10000028,
				"Normal":{
					"StartStandards":1,
					"StartFees":3,
					"AddStandards":3,
					"AddFees":2
				},
				"Custom":[
					{
						"StartStandards":1,
						"StartFees":10,
						"AddStandards":1,
						"AddFees":30,
						"DestCountry":"中国",
						"DestProvince":"广东省",
						"DestCity":"广州市"
					}
				]
			},
			{
				"Type":10000029,
				"Normal":{
					"StartStandards":1,
					"StartFees":4,
					"AddStandards":3,
					"AddFees":3
				},
				"Custom":[
					{
						"StartStandards":1,
						"StartFees":8,
						"AddStandards":2,
						"AddFees":11,
						"DestCountry":"中国",
						"DestProvince":"广东省",
						"DestCity":"广州市"
					}
				]
			}
		]
	}`)

	tpl := DeliveryTemplate{
		Name:      "testexpress",
		Assumer:   ASSUMER_BUYER,
		Valuation: VALUATION_BY_ITEM,
		TopFees: []*TopFee{
			&TopFee{
				ExpressId: 10000027,
				Normal: TopFeeNormal{
					StartStandards: 1,
					StartFees:      2,
					AddStandards:   3,
					AddFees:        1,
				},
				Customs: []*TopFeeCustom{
					&TopFeeCustom{
						TopFeeNormal: TopFeeNormal{
							StartStandards: 1,
							StartFees:      100,
							AddStandards:   1,
							AddFees:        3,
						},
						DestCountry:  "中国",
						DestProvince: "广东省",
						DestCity:     "广州市",
					},
				},
			},
			&TopFee{
				ExpressId: 10000028,
				Normal: TopFeeNormal{
					StartStandards: 1,
					StartFees:      3,
					AddStandards:   3,
					AddFees:        2,
				},
				Customs: []*TopFeeCustom{
					&TopFeeCustom{
						TopFeeNormal: TopFeeNormal{
							StartStandards: 1,
							StartFees:      10,
							AddStandards:   1,
							AddFees:        30,
						},
						DestCountry:  "中国",
						DestProvince: "广东省",
						DestCity:     "广州市",
					},
				},
			},
			&TopFee{
				ExpressId: 10000029,
				Normal: TopFeeNormal{
					StartStandards: 1,
					StartFees:      4,
					AddStandards:   3,
					AddFees:        3,
				},
				Customs: []*TopFeeCustom{
					&TopFeeCustom{
						TopFeeNormal: TopFeeNormal{
							StartStandards: 1,
							StartFees:      8,
							AddStandards:   2,
							AddFees:        11,
						},
						DestCountry:  "中国",
						DestProvince: "广东省",
						DestCity:     "广州市",
					},
				},
			},
		},
	}

	b, err := json.Marshal(tpl)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", tpl, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", tpl, b, want)
		}
	}
}
