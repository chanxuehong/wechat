// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// copy from github.com/chanxuehong/wechat/mp/client/error.go

package pay2

import (
	"errors"
	"fmt"
	"time"

	"github.com/chanxuehong/wechat"
)

const (
	errCodeOK                = 0
	errCodeInvalidCredential = 40001 // access_token 过期（无效）返回这个错误
	errCodeTimeout           = 42001 // access_token 过期（无效）返回这个错误（maybe!!!）
)

// 查看 TokenService.Token 是否有更新, 如果更新了返回新的 token, 否则返回错误.
func getNewToken(tokenService wechat.TokenService, currentToken string) (token string, err error) {
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)

		token, err = tokenService.Token()
		if err != nil {
			return
		}
		if token != currentToken {
			return
		}
	}

	err = errors.New("get new access token failed")
	return
}

type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", e.ErrCode, e.ErrMsg)
}
