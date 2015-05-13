# 微信公众平台（订阅号、服务号） golang SDK

## 整体架构图

![架构图](https://github.com/chanxuehong/wechat/blob/master/mp/mp.png)

## 示例

### 主動調用微信 api，mp 子包裏面的 Client 基本都是這樣的調用方式
```Go
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/menu"
)

var AccessTokenServer = mp.NewDefaultAccessTokenServer("appId", "appSecret", nil) // 一個應用只能有一個實例
var WechatClient = mp.NewWechatClient(AccessTokenServer, nil)

func main() {
	var mn menu.Menu
	mn.Buttons = make([]menu.Button, 3)
	mn.Buttons[0].SetAsClickButton("今日歌曲", "V1001_TODAY_MUSIC")
	mn.Buttons[1].SetAsViewButton("视频", "http://v.qq.com/")

	var subButtons = make([]menu.Button, 2)
	subButtons[0].SetAsViewButton("搜索", "http://www.soso.com/")
	subButtons[1].SetAsClickButton("赞一下我们", "V1001_GOOD")

	mn.Buttons[2].SetAsSubMenuButton("子菜单", subButtons)

	clt := menu.Client{WechatClient: WechatClient}
	if err := clt.CreateMenu(mn); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}
```

### 被動接收消息（事件）推送，一个 URL 监听一个公众号的消息
```Go
package main

import (
	"log"
	"net/http"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"
)

// 非法请求的 Handler
func InvalidRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err.Error())
}

// 文本消息的 Handler
func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg) // 可以省略...
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	//mp.WriteRawResponse(w, r, resp) // 明文模式
	mp.WriteAESResponse(w, r, resp) // 安全模式
}

func main() {
	aesKey, err := util.AESKeyDecode("encodedAESKey") // 这里 encodedAESKey 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux := mp.NewMessageServeMux()
	messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: wechatId, token, appId
	wechatServer := mp.NewDefaultWechatServer("wechatId", "token", "appId", aesKey, messageServeMux)

	wechatServerFrontend := mp.NewWechatServerFrontend(wechatServer, mp.InvalidRequestHandlerFunc(InvalidRequestHandler))

	// 如果你在微信后台设置的回调地址是
	//   http://xxx.yyy.zzz/wechat
	// 那么可以这么注册 http.Handler
	http.Handle("/wechat", wechatServerFrontend)
	http.ListenAndServe(":80", nil)
}
```

### 被動接收消息（事件）推送，一个 URL 监听多个公众号的消息
```Go
package main

import (
	"log"
	"net/http"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"
)

// 非法请求处理函数
func InvalidRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err.Error())
}

// 文本消息的处理
func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg) // 可以省略...
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	//mp.WriteRawResponse(w, r, resp) // 明文模式
	mp.WriteAESResponse(w, r, resp) // 安全模式
}

func main() {
	// wechatServer1
	aesKey1, err := util.AESKeyDecode("encodedAESKey1") // 这里 encodedAESKey1 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux1 := mp.NewMessageServeMux()
	messageServeMux1.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: wechatId1, token1, appId1
	wechatServer1 := mp.NewDefaultWechatServer("wechatId1", "token1", "appId1", aesKey1, messageServeMux1)

	// wechatServer2
	aesKey2, err := util.AESKeyDecode("encodedAESKey2") // 这里 encodedAESKey2 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux2 := mp.NewMessageServeMux()
	messageServeMux2.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: wechatId2, token2, appId2
	wechatServer2 := mp.NewDefaultWechatServer("wechatId2", "token2", "appId2", aesKey2, messageServeMux2)

	// multiWechatServerFrontend
	var multiWechatServerFrontend mp.MultiWechatServerFrontend
	multiWechatServerFrontend.SetInvalidRequestHandler(mp.InvalidRequestHandlerFunc(InvalidRequestHandler))
	multiWechatServerFrontend.SetWechatServer("wechat1", wechatServer1) // 回調url上面要加上 wechat_server=wechat1
	multiWechatServerFrontend.SetWechatServer("wechat2", wechatServer2) // 回調url上面要加上 wechat_server=wechat2

	// 如果你在微信后台设置的回调地址是
	//   http://xxx.yyy.zzz/wechat
	// 那么可以这么注册 http.Handler
	http.Handle("/wechat", &multiWechatServerFrontend)
	http.ListenAndServe(":80", nil)
}
```