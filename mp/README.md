### 回调请求的一般处理逻辑（一个回调地址处理一个公众号的消息和事件）
```Go
package main

import (
	"log"
	"net/http"

	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
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
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
}

func menuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)
}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
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

	"github.com/chanxuehong/wechat/mp/base"
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	wxAppId     = "appid"
	wxAppSecret = "appsecret"

	wxOriId         = "oriid"
	wxToken         = "token"
	wxEncodedAESKey = "aeskey"
)

var (
	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(wxAppId, wxAppSecret, nil)
	wechatClient      *core.Client           = core.NewClient(accessTokenServer, nil)
)

func main() {
	fmt.Println(base.GetCallbackIP(wechatClient))
}
```