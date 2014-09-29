## 简介

封装微信服务器推送到回调 URL 的消息(事件)处理 Handler.

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
	// 如果想动态增加/删除 MsgHandler, 请使用 server.CSHttpHandler
	// 如果你只有一个公众号, 也可以直接使用 server.SingleHttpHandler
	var HttpHandler server.NCSHttpHandler
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