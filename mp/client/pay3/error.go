// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"fmt"
)

type Error struct {
	ErrCode string
	ErrMsg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("errcode: %s, errmsg: %s", e.ErrCode, e.ErrMsg)
}
