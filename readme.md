# 微信公众平台 golang SDK

Version: 0.5.2

## 简介

目前完全实现的功能是被动消息的接收和处理（因为我的公众平台只有这个基本接口，订阅号、没有认证）；
其他部分的实现都是参考微信官方的 API文档（个人认为不是很规范，也许和实际不能匹配），欢迎大家测试 和 fork。

代码还在继续添加中，欢迎大家 push issues。
联系方式：chanxuehong@gmail.com

## 入门

wechat 主要分为 Client 和 Server 两个部分。

Client 实现的是主动发送请求的功能，比如发送客服消息，群发消息，创建自定义菜单......
Client 是并发安全的，在你的应用中一般只用常驻一个 Client 对象就可以了。

Server 实现的是处理被动接收的消息的功能，微信服务器推送过来的普通消息 和 事件推送消息都是 Server 处理的。
Server 实现了 http.Handler 接口，所以一般的应用就是实例化一个 Server 的实例，然后注册到特定的 pattern 上：
```Go
http.Handle("/path", ServerEntity)
```

## 安装

通过执行下列语句就可以完成安装

	go get github.com/chanxuehong/wechat

## 示例

### Server示例：被动处理文本消息

```Go
package main

import (
	"github.com/chanxuehong/wechat"
	"github.com/chanxuehong/wechat/message"
	"net/http"
)

const wechatToken = "yourToken" // 你的微信平台 token

// 处理用户发送过来的 文本消息
func TextRequestHandler(w http.ResponseWriter, r *http.Request, rqst *message.Request) {
	//TODO: 增加你的代码
}

func main() {
	wechatServer := wechat.NewServer(wechatToken)

	// 自定义 文本消息 处理函数，当然你也可以定义别的函数
	wechatServer.SetTextRequestHandler(TextRequestHandler)

	http.Handle("/", wechatServer)

	http.ListenAndServe(":80", nil) // 启动接收微信数据服务器
}
```

#### 自定义处理函数

处理函数的定义可以使用下面的形式

```Go
// 非法的请求（包括不是微信服务器发送过来的和签名认证不通过的）处理函数
type InvalidRequestHandlerFunc func(http.ResponseWriter, *http.Request, error)
// 目前 SDK 不能识别的请求处理函数
type UnknownRequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.Request)
// 正常的从微信服务器推送过来的请求处理函数，都可以自定义。SDK提供了下面的自定义点：
type RequestHandlerFunc func(http.ResponseWriter, *http.Request, *message.Request)

/*
	默认的处理函数是什么都不做。在下面这些 hook 点可以设置自定义函数。
	
	NOTE: 这些函数调用一定要在 http.Handle("/", wechatServer) 之前完成！
	
    func (s *Server) SetClickEventRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetImageRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetInvalidRequestHandler(handler InvalidRequestHandlerFunc)
    func (s *Server) SetLinkRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetLocationEventRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetLocationRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetMasssendjobfinishEventRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetScanEventRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetSubscribeEventByScanRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetSubscribeEventRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetTextRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetUnknownRequestHandler(handler UnknownRequestHandlerFunc)
    func (s *Server) SetUnsubscribeEventRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetVideoRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetViewEventRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetVoiceRecognitionRequestHandler(handler RequestHandlerFunc)
    func (s *Server) SetVoiceRequestHandler(handler RequestHandlerFunc)
*/
```

### Client示例：创建一个临时的二维码

```Go
package main

import (
	"fmt"
	"github.com/chanxuehong/wechat"
)

func main() {
	c := wechat.NewClient("appid", "appsecret")

	qrcode, err := c.QRCodeCreate(100, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(qrcode)
}
```