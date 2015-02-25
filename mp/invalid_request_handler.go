// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import (
	"net/http"
)

// 无效请求(非法或者错误)的处理接口.
type InvalidRequestHandler interface {
	// err 是错误信息
	ServeInvalidRequest(w http.ResponseWriter, r *http.Request, err error)
}

type InvalidRequestHandlerFunc func(http.ResponseWriter, *http.Request, error)

func (fn InvalidRequestHandlerFunc) ServeInvalidRequest(w http.ResponseWriter, r *http.Request, err error) {
	fn(w, r, err)
}

var DefaultInvalidRequestHandler = InvalidRequestHandlerFunc(func(http.ResponseWriter, *http.Request, error) {})
