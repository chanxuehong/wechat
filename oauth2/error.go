package oauth2

import (
	"strconv"

	"github.com/chanxuehong/wechat/util"
)

const (
	ErrCodeOK = 0
)

type Error struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (err *Error) Error() string {
	return util.StringsJoin("errcode:", strconv.FormatInt(err.ErrCode, 10), ", errmsg:", err.ErrMsg)
}
