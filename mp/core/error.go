package core

import (
	"fmt"
	"reflect"
)

const (
	ErrCodeOK                 = 0
	ErrCodeInvalidCredential  = 40001 // access_token 过期错误码
	ErrCodeAccessTokenExpired = 42001 // access_token 过期错误码(maybe!!!)
)

var (
	errorType      = reflect.TypeOf(Error{})
	errorZeroValue = reflect.Zero(errorType)
)

const (
	errorErrCodeIndex = 0
	errorErrMsgIndex  = 1
)

type Error struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", err.ErrCode, err.ErrMsg)
}
