// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"crypto/md5"
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/mp/pay"
)

// 检查 parameters 的签名是否正确, 正确返回 nil, 否则返回错误信息
//  parameters:  待签名的参数
//  Key:         支付签名的 Key
//
//  NOTE: 调用之前一般要确保有 sign 字段, 特别是有 return_code 时要判断是否为 RET_CODE_SUCCESS
func CheckSignature(parameters map[string]string, Key string) (err error) {
	if parameters == nil {
		return errors.New("parameters == nil")
	}

	signature1 := parameters["sign"]
	if signature1 == "" {
		return errors.New("sign is empty")
	}
	if len(signature1) != md5.Size*2 {
		err = fmt.Errorf("不正确的签名: %q, 长度不对, have: %d, want: %d",
			signature1, len(signature1), md5.Size*2)
		return
	}

	signature2 := pay.MD5Signature(parameters, Key)

	if subtle.ConstantTimeCompare([]byte(signature2), []byte(signature1)) != 1 {
		err = fmt.Errorf("签名不匹配, \r\nlocal: %q, \r\ninput: %q", signature2, signature1)
		return
	}
	return
}

// 设置 parameters 签名, 一般最后调用, 正确返回 nil, 否则返回错误信息
//  parameters:  待签名的参数
//  Key:         支付签名的 Key
func SetSignature(parameters map[string]string, Key string) (err error) {
	if parameters == nil {
		return errors.New("parameters == nil")
	}

	parameters["sign"] = pay.MD5Signature(parameters, Key)
	return
}
