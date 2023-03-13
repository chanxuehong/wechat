package core

import (
	"reflect"
	"strconv"

	"github.com/bububa/wechat/util"
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
	return util.StringsJoin("errcode:", strconv.FormatInt(err.ErrCode, 10), ", errmsg:", err.ErrMsg)
}
