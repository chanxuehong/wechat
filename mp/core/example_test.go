package core_test

import (
	"net/http"

	"github.com/chanxuehong/wechat/mp/core"
)

func ExampleServer_ServeHTTP() {
	// 在回调 URL 的 Handler 里处理消息(事件)
	http.HandleFunc("/wechat_callback", func(w http.ResponseWriter, r *http.Request) {
		mux := core.NewServeMux() // 创建 core.Handler, 也可以用自己实现的 core.Handler

		// 注册消息(事件)处理 Handler, 都不是必须的!
		{
			mux.UseFunc(func(ctx *core.Context) { // 注册中间件, 处理所有的消息(事件)
				// TODO: 中间件处理逻辑
			})
			mux.UseFuncForMsg(func(ctx *core.Context) { // 注册中间件, 处理所有的消息
				// TODO: 中间件处理逻辑
			})
			mux.UseFuncForEvent(func(ctx *core.Context) { // 注册中间件, 处理所有的事件
				// TODO: 中间件处理逻辑
			})

			mux.DefaultMsgHandleFunc(func(ctx *core.Context) { // 设置默认消息处理 Handler
				// TODO: 消息处理逻辑
			})
			mux.DefaultEventHandleFunc(func(ctx *core.Context) { // 设置默认事件处理 Handler
				// TODO: 事件处理逻辑
			})

			mux.MsgHandleFunc("{MsgType}", func(ctx *core.Context) { // 设置具体类型的消息处理 Handler
				// TODO: 消息处理逻辑
			})
			mux.EventHandleFunc("{EventType}", func(ctx *core.Context) { // 设置具体类型的事件处理 Handler
				// TODO: 事件处理逻辑
			})
		}

		// 创建 Server, 设置正确的参数
		srv := core.NewServer("{oriId}", "{appId}", "{token}", "{base64AESKey}", mux, nil)

		// 处理回调请求
		srv.ServeHTTP(w, r, nil)
	})
}
