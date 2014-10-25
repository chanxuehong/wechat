## 简介

封装微信服务器推送到回调 URL 的消息(事件)处理 Handler.

## 注意

server 提供了 AgentFrontend，MultiAgentFrontend

正常情况下使用 AgentFrontend 即可，即一个回调 URL 只能接受一个公众号的消息（事件），
如果需要处理多个公众号的消息（事件），可以调用 net/http.Handle 来动态增加 URL<-->AgentFrontend pair。

如果某些特殊情况下，给你的 URL 只有一个，但是你又想处理多个公众号的消息（事件），我们这里提供了
MultiAgentFrontend，这时公众号的回调 URL 就要增加一个查询参数 agentkey（根据需要你也可以
更改这个参数名称，如果更改了同时也要修改 URLQueryAgentKeyName 的值），这样回调 URL 的格式一般为

http://abc.xyz.com/weixin?agentkey=agentkey_value

MultiAgentFrontend 是并发安全的，可以动态增加（删除）Agent，如果你没有动态需求，
只是在初始化的时候配置各种参数，可以自己把 MultiAgentFrontend 里面的带有 rwmutex 的代码
都去掉，高并发没有锁的开销。

## 示例

```golang
package main

import (
	"github.com/chanxuehong/wechat/corp/message/passive/request"
	"github.com/chanxuehong/wechat/corp/message/passive/response"
	"github.com/chanxuehong/wechat/corp/server"
	"github.com/chanxuehong/wechat/util"
	"log"
	"net/http"
	"time"
)

// 实现 server.Agent
type CustomAgent struct {
	server.DefaultAgent // 可选, 不是必须!!! 提供了默认实现
}

// 文本消息处理函数
func (this *CustomAgent) ServeTextMsg(w http.ResponseWriter, r *http.Request,
	msg *request.Text, rawXMLMsg []byte, timestamp int64, nonce string, random [16]byte) {

	// TODO: 示例代码, 把用户发送过来的文本原样回复过去

	w.Header().Set("Content-Type", "application/xml; charset=utf-8") // 可选

	// NOTE: 时间戳也可以用传入的参数 timestamp, 即微信服务器请求 URL 中的 timestamp
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.Content, time.Now().Unix())

	// timestamp, nonce, random 也可以自己生成
	if err := server.WriteText(w, resp, timestamp, nonce, this.GetAESKey(), random[:], this.GetCorpId(), this.GetToken()); err != nil {
		// TODO: 错误处理代码
	}
}

// 自定义错误请求处理函数
func CustomInvalidRequestHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	// TODO: 这里只是简单的做下 log
	log.Println(err)
}

func init() {
	AESKey, err := util.AESKeyDecode("EncodingAESKey") // 后台里获取
	if err != nil {
		panic(err)
	}

	var agent CustomAgent
	agent.DefaultAgent.Init("CorpId", 1 /*AgentId*/, "Token", AESKey)

	agentFrontend := server.NewAgentFrontend(&agent,
		server.InvalidRequestHandlerFunc(CustomInvalidRequestHandlerFunc))

	// 注册这个 agentFrontend 到回调 URL 上
	// 比如你在公众平台后台注册的回调地址是 http://abc.xyz.com/weixin，那么可以这样注册
	http.Handle("/weixin", agentFrontend)
}

func main() {
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
```