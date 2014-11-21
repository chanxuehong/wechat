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

func timeoutRetryWait() {
	time.Sleep(300 * time.Millisecond)
}

type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", e.ErrCode, e.ErrMsg)
}
