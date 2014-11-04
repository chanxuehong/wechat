// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"bytes"
	"crypto/md5"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
)

// 检查 parameters 的签名是否正确, 正确返回 nil, 否则返回错误信息
//  parameters:  待签名的参数
//  signKey:     支付签名的 signKey
//
//  NOTE: 调用之前一般要确保有 sign 字段, 特别是有 return_code 时要判断是否为 RET_CODE_SUCCESS
func CheckMD5Signature(parameters map[string]string, signKey string) (err error) {
	if parameters == nil {
		return errors.New("parameters == nil")
	}

	signature1 := parameters["sign"]
	if signature1 == "" {
		return errors.New("sign is empty")
	}
	if len(signature1) != 32 {
		err = fmt.Errorf("不正确的签名: %q, 长度不对, have: %d, want: %d",
			signature1, len(signature1), 32)
		return
	}

	signature2 := MD5Sign(parameters, signKey)

	if subtle.ConstantTimeCompare([]byte(signature2), []byte(signature1)) != 1 {
		err = fmt.Errorf("签名不匹配, \r\nlocal: %q, \r\ninput: %q", signature2, signature1)
		return
	}
	return
}

// 设置 parameters 签名, 一般最后调用, 正确返回 nil, 否则返回错误信息
//  parameters:  待签名的参数
//  signKey:     支付签名的 signKey
func SetMD5Signature(parameters map[string]string, signKey string) (err error) {
	if parameters == nil {
		return errors.New("parameters == nil")
	}

	parameters["sign"] = MD5Sign(parameters, signKey)
	return
}

// 对 parameters 里的参数做 MD5 签名.
// 签名方法:
// a.对所有传入参数按照字段名的 ASCII 码从小到大排序（字典序）后，使用 URL 键值
// 对的格式（即 key1=value1&key2=value2…）拼接成字符串 string1，注意：值为空的参数不
// 参与签名；
// b. 在 string1 最 后 拼 接 上 key=signKey 得 到 stringSignTemp 字符串，并对
// stringSignTemp 进行 md5 运算，再将得到的字符串所有字符转换为大写，得到 sign 值
// signValue。
//
//  parameters:  待签名的参数
//  signKey:     支付签名的 Key
func MD5Sign(parameters map[string]string, signKey string) string {
	keys := make([]string, 0, len(parameters))
	for key, value := range parameters {
		if value == "" {
			continue
		}
		if key == "sign" {
			continue
		}

		keys = append(keys, key)
	}
	sort.Strings(keys)

	Hash := md5.New()
	hashsum := make([]byte, 32)

	for _, key := range keys {
		value := parameters[key]

		Hash.Write([]byte(key))
		Hash.Write([]byte{'='})
		Hash.Write([]byte(value))
		Hash.Write([]byte{'&'})
	}
	Hash.Write([]byte("key="))
	Hash.Write([]byte(signKey))

	hex.Encode(hashsum, Hash.Sum(nil))
	return string(bytes.ToUpper(hashsum))
}
