## 简介

封装微信服务器推送到回调 URL 的消息(事件)处理 Handler.

## 示例

```golang
package main

import (
	"github.com/chanxuehong/wechat/message/passive/request"
	"github.com/chanxuehong/wechat/message/passive/response"
	"github.com/chanxuehong/wechat/server"
	"net/http"
)

// 自己实现一个 server.MsgHandler
type CustomMsgHandler struct {
	server.DefaultMsgHandler // 提供了默认实现
}

// 自定义文本消息处理函数, 覆盖默认的实现
func (handler *CustomMsgHandler) TextMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64) {
	// 示例代码

	w.Header().Set("Content-Type", "application/xml; charset=utf-8") // 可选
	resp := response.NewText(msg.FromUserName, msg.ToUserName, "已收到: "+msg.Content)
	if err := handler.WriteText(w, resp); err != nil {
		// TODO: 增加错误处理代码
	}
}

func init() {
	// TODO: 获取必要数据的代码

	var MsgHandler CustomMsgHandler
	MsgHandler.Token = "Token"

	var HttpHandler server.HttpHandler
	HttpHandler.MsgHandler = &MsgHandler

	// 注册这个 handler 到回调 URL 上
	// 比如你在公众平台后台注册的回调地址是 http://abc.xxx.com/weixin，那么可以这样注册
	http.Handle("/weixin", HttpHandler)
}

func main() {
	http.ListenAndServe(":80", nil)
}
```