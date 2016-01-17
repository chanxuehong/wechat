package core_test

import (
	"net/http"

	"github.com/chanxuehong/wechat/mp/core"
)

func ExampleServer_ServeHTTP() {
	// 在回调 URL 的 Handler 里处理消息(事件)
	http.HandleFunc("/wecaht_callback", func(w http.ResponseWriter, r *http.Request) {
		// 创建默认的 core.Handler, 也可以用自己实现的 core.Handler
		mux := core.NewServeMux()
		// 注册中间件
		mux.UseFunc(func(ctx *core.Context) {
			// TODO: 中间件处理逻辑
		})
		// 设置默认消息处理 Handler
		mux.DefaultMsgHandleFunc(func(ctx *core.Context) {
			// TODO: 消息处理逻辑
		})
		// 设置默认事件处理 Handler
		mux.DefaultEventHandleFunc(func(ctx *core.Context) {
			// TODO: 事件处理逻辑
		})

		// 创建 Server, 设置正确的参数
		srv := core.NewServer("", "", "{token}", "", mux, nil)

		// 处理回调请求
		srv.ServeHTTP(w, r, nil)
	})
}
