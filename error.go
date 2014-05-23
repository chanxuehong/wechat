package wechat

import (
	"fmt"
)

// 微信服务器返回的错误基本都是这个格式
type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误编号: %d, 错误信息: %s", e.ErrCode, e.ErrMsg)
}
