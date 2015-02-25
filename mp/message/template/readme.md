### 发送模板消息示例
```Go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/template"
)

var TokenServer = mp.NewDefaultTokenServer("appid", "appsecret", nil)

func main() {
	type Item struct {
		Value string `json:"value"`
		Color string `json:"color"`
	}

	//{
	//    "first": {
	//        "value": "恭喜你购买成功！",
	//        "color": "#173177"
	//    },
	//    "keynote1": {
	//        "value": "巧克力",
	//        "color": "#173177"
	//    },
	//    "keynote2": {
	//        "value": "39.8元",
	//        "color": "#173177"
	//    },
	//    "keynote3": {
	//        "value": "2014年9月16日",
	//        "color": "#173177"
	//    },
	//    "remark": {
	//        "value": "欢迎再次购买！",
	//        "color": "#173177"
	//    }
	//}
	var data = struct {
		First    Item `json:"first"`
		KeyNote1 Item `json:"keynote1"`
		KeyNote2 Item `json:"keynote2"`
		KeyNote3 Item `json:"keynote3"`
		Remark   Item `json:"remark"`
	}{
		First: Item{
			Value: "恭喜你购买成功！",
			Color: "#173177",
		},
		KeyNote1: Item{
			Value: "巧克力",
			Color: "#173177",
		},
		KeyNote2: Item{
			Value: "39.8元",
			Color: "#173177",
		},
		KeyNote3: Item{
			Value: "2014年9月16日",
			Color: "#173177",
		},
		Remark: Item{
			Value: "欢迎再次购买！",
			Color: "#173177",
		},
	}

	dataJSONBytes, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := template.TemplateMessage{
		ToUser:      "touser",
		TemplateId:  "template_id",
		URL:         "http://weixin.qq.com/download",
		TopColor:    "#FF0000",
		RawJSONData: dataJSONBytes,
	}

	clt := template.NewClient(TokenServer, nil)
	msgId, err := clt.Send(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("MsgId:", msgId)
}
```
