## 简介

封装微信服务器推送过来的常规消息(事件)处理 Handler.
对于公众号开发模式, 都会要求提供一个 URL 来处理微信服务器推送过来的消息和事件,
这个 package 就是封装推送到这个 URL 上的消息(事件)处理 Handler.

## 使用方法

package 会自动接收消息并做解析，然后回调你在 hook 点上设置的函数；你可以自定义下面这些类型的 hook 函数：

```Go
// 非法请求的处理函数.
//  @err: 具体的错误信息
type InvalidRequestHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

// 未知消息类型的消息处理函数.
//  @msg: 接收到的消息体
type UnknownRequestHandlerFunc func(w http.ResponseWriter, r *http.Request, msg []byte)

// 正常的消息处理函数
type TextRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Text)
type ImageRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Image)
type VoiceRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Voice)
type VideoRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Video)
type LocationRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Location)
type LinkRequestHandlerFunc func(http.ResponseWriter, *http.Request, *request.Link)
type SubscribeEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.SubscribeEvent)
type UnsubscribeEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.UnsubscribeEvent)
type SubscribeByScanEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.SubscribeByScanEvent)
type ScanEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.ScanEvent)
type LocationEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.LocationEvent)
type MenuClickEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuClickEvent)
type MenuViewEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MenuViewEvent)
type MassSendJobFinishEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MassSendJobFinishEvent)
type MerchantOrderEventHandlerFunc func(http.ResponseWriter, *http.Request, *request.MerchantOrderEvent)
```

你可以在 HandlerSetting 里设置自定义的 hook 处理函数（默认的函数是什么都不做），
然后通过这个 HandlerSetting 创建 Handler，把这个 Handler 绑定到你在公众平台中注册的回调 URL 上。

## 示例

```Go
package main

import (
	"github.com/chanxuehong/wechat/message/request"
	"github.com/chanxuehong/wechat/message/response"
	"github.com/chanxuehong/wechat/server"
	"net/http"
)

// 一般一个应用维护一个实例即可
var wechatHandler *server.Handler

func init() {
	// TODO: 获取必要数据的代码

	// 初始化 wechatHandler
	setting := &server.HandlerSetting{
		Token:              "你的公众号的 token",
		TextRequestHandler: TextRequestHandler,
	}
	wechatHandler = server.NewHandler(setting) // 并发安全，一般一个应用只用一个实例即可

	// 注册这个 handler 到回调 URL 上
	// 比如你在公众平台后台注册的回调地址是 http://abc.xxx.com/weixin，那么可以这样注册
	http.Handle("/weixin", wechatHandler)
}

// 自定义文本消息处理函数
func TextRequestHandler(w http.ResponseWriter, r *http.Request, text *request.Text) {
	//TODO: 添加你的代码，下面只是示例代码！

	// 把用户发送过来的文本原样的回复过去
	resp := response.NewText(text.FromUserName, text.ToUserName, text.Content)

	w.Header().Set("Content-Type", "application/xml; charset=utf-8") // 可选
	server.WriteText(w, resp)
}

func main() {
	http.ListenAndServe(":80", nil)
}
```