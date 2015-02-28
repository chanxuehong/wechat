// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"net/http"
)

// 处理单个APP的消息.
type MessageServerFrontend struct {
	messageServer         MessageServer
	invalidRequestHandler InvalidRequestHandler
}

func NewMessageServerFrontend(server MessageServer, handler InvalidRequestHandler) *MessageServerFrontend {
	if server == nil {
		panic("pay: nil MessageServer")
	}
	if handler == nil {
		handler = DefaultInvalidRequestHandler
	}

	return &MessageServerFrontend{
		messageServer:         server,
		invalidRequestHandler: handler,
	}
}

// 实现 http.Handler.
func (frontend *MessageServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	messageServer := frontend.messageServer
	invalidRequestHandler := frontend.invalidRequestHandler

	ServeHTTP(w, r, nil, messageServer, invalidRequestHandler)
}
