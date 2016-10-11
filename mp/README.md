### 回调请求的一般处理逻辑（一个回调地址处理一个公众号的消息和事件）
```Go
package main

import (
	"log"
	"net/http"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
)

const (
	wxAppId     = "appid"
	wxAppSecret = "appsecret"

	wxOriId         = "oriid"
	wxToken         = "token"
	wxEncodedAESKey = "aeskey"
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler core.Handler
	msgServer  *core.Server
)

func init() {
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	msgHandler = mux
	msgServer = core.NewServer(wxOriId, wxAppId, wxToken, wxEncodedAESKey, msgHandler, nil)
}

func textMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)

	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	//ctx.RawResponse(resp) // 明文回复
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func init() {
	http.HandleFunc("/wx_callback", wxCallbackHandler)
}

// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func wxCallbackHandler(w http.ResponseWriter, r *http.Request) {
	msgServer.ServeHTTP(w, r, nil)
}

func main() {
	log.Println(http.ListenAndServe(":80", nil))
}
```

### 公众号api调用的一般处理逻辑
```Go
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat.v2/mp/base"
	"github.com/chanxuehong/wechat.v2/mp/core"
)

const (
	wxAppId     = "appid"
	wxAppSecret = "appsecret"

	wxOriId         = "oriid"
	wxToken         = "token"
	wxEncodedAESKey = "aeskey"
)

var (
	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(wxAppId, wxAppSecret, "public",nil)
	wechatClient      *core.Client           = core.NewClient(accessTokenServer, nil)
)

func main() {
	fmt.Println(base.GetCallbackIP(wechatClient))
}
```

### 企业号demo
```
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat.v2/mp/core"
)

const (
	CORPID     = "corpid"
	CORPSECRET = "corpsecret"
)

var (
	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(CORPID, CORPSECRET, "corp", nil)
	wechatClient      *core.Client           = core.NewClient(accessTokenServer, nil)
)

func main() {
	currentToken, _ := accessTokenServer.Token()
	fmt.Println(currentToken)
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="

	// 发送的文本消息
	// 	{
	//    "touser": "@all",
	//    "msgtype": "text",
	//    "agentid": 2,
	//    "text": {
	//        "content": "Holiday Request For Pony(http://xxxxx)"
	//    },
	//    "safe":1
	// }
	type MsgText struct {
		Content string `json:"content"`
	}
	type Reqjson struct {
		Touser  string  `json:"touser"`
		Msgtype string  `json:"msgtype"`
		Agentid int64   `json:"agentid"`
		Text    MsgText `json:"text"`
		Safe    int64   `json:"safe"`
	}

	reqjson := &Reqjson{
		Touser:  "@all",
		Msgtype: "text",
		Agentid: 2,
		Text:    MsgText{Content: "aaaaaaa"},
		Safe:    0,
	}

	// 接收消息
	//  {
	//    "errcode": 0,
	//    "errmsg": "ok",
	//    "invaliduser": "UserID1",
	//    "invalidparty":"PartyID1",
	//    "invalidtag":"TagID1"
	// }
	type msgRes struct {
		core.Error
		Errcode      int64  `json:"errcode"`
		Errmsg       string `json:"errmsg"`
		Invaliduser  string `json:"invaliduser"`
		Invalidparty string `json:"invalidparty"`
		Invalidtag   string `json:"invalidtag"`
	}
	var resjson = &msgRes{}
	err := wechatClient.PostJSON(url, reqjson, resjson)
	fmt.Println(resjson, err)
}
```