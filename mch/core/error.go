package core

import (
	"encoding/xml"
	"errors"

	"github.com/bububa/wechat/util"
)

var (
	ErrNotFoundReturnCode = errors.New("not found return_code parameter")
	ErrNotFoundResultCode = errors.New("not found result_code parameter")
	ErrNotFoundSign       = errors.New("not found sign parameter")
)

var _ error = (*Error)(nil)

// 协议错误, return_code 不为 SUCCESS 时有返回.
type Error struct {
	XMLName    struct{} `xml:"xml"                  json:"-"`
	ReturnCode string   `xml:"return_code"          json:"return_code"`
	ReturnMsg  string   `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
}

func (e *Error) Error() string {
	bs, err := xml.Marshal(e)
	if err != nil {
		return util.StringsJoin("return_code:", e.ReturnCode, ", return_msg:", e.ReturnMsg)
	}
	return string(bs)
}

var _ error = (*BizError)(nil)

// 业务错误, result_code 不为 SUCCESS 时有返回.
type BizError struct {
	XMLName     struct{} `xml:"xml"                    json:"-"`
	ResultCode  string   `xml:"result_code"            json:"result_code"`
	ErrCode     string   `xml:"err_code,omitempty"     json:"err_code,omitempty"`
	ErrCodeDesc string   `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
}

func (e *BizError) Error() string {
	bs, err := xml.Marshal(e)
	if err != nil {
		return util.StringsJoin("result_code:", e.ResultCode, ", err_code:", e.ErrCode, ", err_code_des:", e.ErrCodeDesc)
	}
	return string(bs)
}
