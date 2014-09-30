## 简介

封装微信服务器推送到回调 URL 的消息(事件)处理 Handler.

## 注意

这里提供了三种 HttpHandler: SingleHttpHandler, CSMultiHttpHandler, NCSMultiHttpHandler，

正常情况下使用 SingleHttpHandler 即可，即一个回调 URL 只能接受一个公众号的消息（事件），如果需要处理多个公众号的消息（事件），可以调用 net/http.Handle 来动态增加 URL，SingleHttpHandler 对。

如果某些特殊情况下，给你的 URL 只有一个，但是你又想处理多个公众号的消息（事件），我们这里提供了 CSMultiHttpHandler 和 NCSMultiHttpHandler，这两个的区别就是一个并发安全，一个并发不安全，
CSMultiHttpHandler 可以动态的增加 MsgHandler，NCSMultiHttpHandler 只能在初始化的时候增加
MsgHandler，运行中不能动态增加，因为并发不安全！

## 示例

```golang
package main

import (
	"github.com/chanxuehong/wechat/message/passive/request"
	"github.com/chanxuehong/wechat/message/passive/response"
	"github.com/chanxuehong/wechat/server"
	"log"
	"net/http"
	"time"
)

// 实现 server.MsgHandler
type CustomMsgHandler struct {
	server.DefaultMsgHandler // 可选, 不是必须!!! 提供了默认实现
}

// 文本消息处理函数, 覆盖默认的实现
func (handler *CustomMsgHandler) TextMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64) {
	// TODO: 示例代码

	w.Header().Set("Content-Type", "application/xml; charset=utf-8") // 可选

	// 时间戳也可以用传入的参数 timestamp, 即微信服务器请求的 timestamp
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.Content, time.Now().Unix())

	if err := handler.WriteText(w, resp); err != nil {
		// TODO: 错误处理代码
	}
}

// 自定义错误请求处理函数
func CustomInvalidRequestHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	// TODO: 这里只是简单的做下 log
	log.Println(err)
}

func init() {
	var MsgHandler CustomMsgHandler
	MsgHandler.DefaultMsgHandler.Token = "Token" // 这里填上你自己的 Token

	// 这里创建的是非并发安全的 HttpHandler, 所有的配置工作都要在注册到 URL 之前完成,
	// 如果想动态增加/删除 MsgHandler, 请使用 server.CSMultiHttpHandler
	// 如果你只有一个公众号, 也可以直接使用 server.SingleHttpHandler
	var HttpHandler server.NCSMultiHttpHandler
	HttpHandler.SetInvalidRequestHandler(server.InvalidRequestHandlerFunc(CustomInvalidRequestHandlerFunc))
	HttpHandler.SetMsgHandler("WechatMPId", &MsgHandler) // 微信号公众原始ID，在后台中能看到

	// 注册这个 handler 到回调 URL 上
	// 比如你在公众平台后台注册的回调地址是 http://abc.xxx.com/weixin，那么可以这样注册
	http.Handle("/weixin", &HttpHandler)
}

func main() {
	http.ListenAndServe(":80", nil)
}
```