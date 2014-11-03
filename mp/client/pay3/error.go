// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"fmt"
)

type Error struct {
	XMLName struct{} `xml:"xml"                   json:"-"`
	RetCode string   `xml:"return_code"           json:"return_code"`
	RetMsg  string   `xml:"return_msg,omitempty"  json:"return_msg,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("return_code: %s, return_msg: %s", e.RetCode, e.RetMsg)
}
