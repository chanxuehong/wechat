# 微信企业号 golang SDK

## 整体架构图

![架构图](https://github.com/chanxuehong/wechat/blob/master/corp/corp.png)

## 示例

### 一个 URL 监听一个企业号应用的消息
```Go
package main

import (
	"github.com/chanxuehong/wechat/corp"
	"github.com/chanxuehong/wechat/corp/message/request"
	"github.com/chanxuehong/wechat/corp/message/response"
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
func TextMessageHandler(w http.ResponseWriter, r *corp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg)
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	corp.WriteResponse(w, r, resp)
}

func main() {
	aesKey, err := util.AESKeyDecode("encodedAESKey") // 这里 encodedAESKey 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux := corp.NewMessageServeMux()
	messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: corpId, agentId, token
	agentServer := corp.NewDefaultAgentServer("corpId", 0 /* agentId */, "token", aesKey, messageServeMux)

	agentServerFrontend := corp.NewAgentServerFrontend(agentServer, corp.InvalidRequestHandlerFunc(InvalidRequestHandler))

	// 如果你在微信后台设置的回调地址是
	//   http://xxx.yyy.zzz/agent
	// 那么可以这么注册 http.Handler
	http.Handle("/agent", agentServerFrontend)
	http.ListenAndServe(":80", nil)
}
```

### 一个 URL 监听多个企业号应用的消息
```Go
package main

import (
	"github.com/chanxuehong/wechat/corp"
	"github.com/chanxuehong/wechat/corp/message/request"
	"github.com/chanxuehong/wechat/corp/message/response"
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
func TextMessageHandler(w http.ResponseWriter, r *corp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg)
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	corp.WriteResponse(w, r, resp)
}

func main() {
	aesKey1, err := util.AESKeyDecode("encodedAESKey1") // 这里 encodedAESKey1 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux1 := corp.NewMessageServeMux()
	messageServeMux1.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: corpId1, agentId1, token1
	agentServer1 := corp.NewDefaultAgentServer("corpId1", 1 /* agentId1 */, "token1", aesKey1, messageServeMux1)

	aesKey2, err := util.AESKeyDecode("encodedAESKey2") // 这里 encodedAESKey2 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux2 := corp.NewMessageServeMux()
	messageServeMux2.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: corpId2, agentId2, token2
	agentServer2 := corp.NewDefaultAgentServer("corpId2", 2 /* agentId2 */, "token2", aesKey2, messageServeMux2)

	var multiAgentServerFrontend corp.MultiAgentServerFrontend
	multiAgentServerFrontend.SetInvalidRequestHandler(corp.InvalidRequestHandlerFunc(InvalidRequestHandler))
	multiAgentServerFrontend.SetAgentServer("agent1", agentServer1) // 需要相应设置回调 url 的参数
	multiAgentServerFrontend.SetAgentServer("agent2", agentServer2) // 需要相应设置回调 url 的参数

	// 如果你在微信后台设置的回调地址是
	//   http://xxx.yyy.zzz/agent
	// 那么可以这么注册 http.Handler
	http.Handle("/agent", &multiAgentServerFrontend)
	http.ListenAndServe(":80", nil)
}
```