# 微信公众平台（订阅号、服务号） golang SDK

## 整体架构图

![架构图](https://github.com/chanxuehong/wechat/blob/master/mp/mp.png)

## 示例

### 一个 URL 监听一个公众号的消息
```Go
package main

import (
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"
	"io"
	"log"
	"net/http"
)

// 非法请求处理函数
func InvalidRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
	io.WriteString(w, err.Error())
	log.Println(err.Error())
}

// 文本消息的处理
func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg)
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

### 一个 URL 监听多个公众号的消息
```Go
package main

import (
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"
	"io"
	"log"
	"net/http"
)

// 非法请求处理函数
func InvalidRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
	io.WriteString(w, err.Error())
	log.Println(err.Error())
}

// 文本消息的处理
func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg)
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	//mp.WriteRawResponse(w, r, resp) // 明文模式
	mp.WriteAESResponse(w, r, resp) // 安全模式
}

func main() {
	aesKey1, err := util.AESKeyDecode("encodedAESKey1") // 这里 encodedAESKey1 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux1 := mp.NewMessageServeMux()
	messageServeMux1.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: wechatId1, token1, appId1
	wechatServer1 := mp.NewDefaultWechatServer("wechatId1", "token1", "appId1", aesKey1, messageServeMux1)

	aesKey2, err := util.AESKeyDecode("encodedAESKey2") // 这里 encodedAESKey2 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux2 := mp.NewMessageServeMux()
	messageServeMux2.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: wechatId2, token2, appId2
	wechatServer2 := mp.NewDefaultWechatServer("wechatId2", "token2", "appId2", aesKey2, messageServeMux2)

	var multiWechatServerFrontend mp.MultiWechatServerFrontend
	multiWechatServerFrontend.SetInvalidRequestHandler(mp.InvalidRequestHandlerFunc(InvalidRequestHandler))
	multiWechatServerFrontend.SetWechatServer("wechat1", wechatServer1) // 需要相应设置回调 url 的参数
	multiWechatServerFrontend.SetWechatServer("wechat2", wechatServer2) // 需要相应设置回调 url 的参数

	// 如果你在微信后台设置的回调地址是
	//   http://xxx.yyy.zzz/wechat
	// 那么可以这么注册 http.Handler
	http.Handle("/wechat", &multiWechatServerFrontend)
	http.ListenAndServe(":80", nil)
}
```