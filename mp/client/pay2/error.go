// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// copy from github.com/chanxuehong/wechat/mp/client/error.go

package pay2

import (
	"fmt"
	"time"
)

const (
	errCodeOK                = 0
	errCodeInvalidCredential = 40001 // access_token 过期（无效）返回这个错误
	errCodeTimeout           = 42001 // access_token 过期（无效）返回这个错误（maybe!!!）
)

// 当 TokenService 定时更新 access_token 的时候, 之前从缓存中获取的 access_token 就会失效,
// 但是安全考虑不能自己去服务器获取新的 access_token, 而是等待 TokenService 完成 access_token 的
// 更新动作, 再到缓存中读取新的 access_token!
//
// 这就是等待函数, 这里设置为等待 200ms, 一般稍微大于
// TokenService 从微信服务器获取 access_token 的时间
// +
// TokenService 把这个 access_token 存入缓存的时间
func timeoutRetryWait() {
	time.Sleep(200 * time.Millisecond)
}

type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", e.ErrCode, e.ErrMsg)
}
