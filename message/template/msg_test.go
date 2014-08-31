// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package template

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestMarshalAndNewFunc(t *testing.T) {
	// 纯文本模板消息==============================================================

	expectBytes := []byte(`{
	    "touser":"OPENID",
	    "template_id":"aygtGTLdrjHJP7Bu4EdkptNfYaeFKi98ygn2kitCJ6fAfdmN88naVvX6V5uIV5x0",
	    "data":{
	        "Goods":"苹果",
	        "Unit_price":"RMB 20.13",
	        "Quantity":"5",
	        "Total":"RMB 100.65",
	        "Source":{
	            "Shop":"Jas屌丝商店",
	            "Recommend":"5颗星"
	        }
	    }
	}`)

	var data struct {
		Goods      string
		Unit_price string
		Quantity   string
		Total      string
		Source     struct {
			Shop      string
			Recommend string
		}
	}
	data.Goods = "苹果"
	data.Unit_price = "RMB 20.13"
	data.Quantity = "5"
	data.Total = "RMB 100.65"
	data.Source.Shop = "Jas屌丝商店"
	data.Source.Recommend = "5颗星"

	msg := NewMsg("OPENID", "aygtGTLdrjHJP7Bu4EdkptNfYaeFKi98ygn2kitCJ6fAfdmN88naVvX6V5uIV5x0", &data)

	b, err := json.Marshal(msg)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", msg, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", msg, b, want)
		}
	}

	// 链接的模板消息==============================================================

	expectBytes = []byte(`{
	    "touser":"OPENID",
	    "template_id":"ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
	    "data":{
	        "User":{
	            "value":"黄先生",
	            "color":"#173177"
	        },
	        "Date":{
	            "value":"06月07日 19时24分",
	            "color":"#173177"
	        },
	        "CardNumber":{
	            "value":"0426",
	            "color":"#173177"
	        },
	        "Type":{
	            "value":"消费",
	            "color":"#173177"
	        },
	        "Money":{
	            "value":"人民币260.00元",
	            "color":"#173177"
	        },
	        "DeadTime":{
	            "value":"06月07日19时24分",
	            "color":"#173177"
	        },
	        "Left":{
	            "value":"6504.09",
	            "color":"#173177"
	        }
	    },
		"url":"http://weixin.qq.com/download",
	    "topcolor":"#FF0000"
	}`)

	type DateValue struct {
		Value string `json:"value"`
		Color string `json:"color"`
	}
	var data1 struct {
		User       DateValue
		Date       DateValue
		CardNumber DateValue
		Type       DateValue
		Money      DateValue
		DeadTime   DateValue
		Left       DateValue
	}
	data1.User.Value = "黄先生"
	data1.User.Color = "#173177"
	data1.Date.Value = "06月07日 19时24分"
	data1.Date.Color = "#173177"
	data1.CardNumber.Value = "0426"
	data1.CardNumber.Color = "#173177"
	data1.Type.Value = "消费"
	data1.Type.Color = "#173177"
	data1.Money.Value = "人民币260.00元"
	data1.Money.Color = "#173177"
	data1.DeadTime.Value = "06月07日19时24分"
	data1.DeadTime.Color = "#173177"
	data1.Left.Value = "6504.09"
	data1.Left.Color = "#173177"

	msg1 := NewMsgWithLink("OPENID", "ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
		&data1, "http://weixin.qq.com/download", "#FF0000")

	b, err = json.Marshal(msg1)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", msg1, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", msg1, b, want)
		}
	}
}
