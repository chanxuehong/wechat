// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package util

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

// 微信公众号 明文模式/URL认证 签名
func Sign(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	n := len(token) + len(timestamp) + len(nonce)
	buf := make([]byte, 0, n)

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

// 微信公众号/企业号 密文模式消息签名
func MsgSign(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strs.Sort()

	n := len(token) + len(timestamp) + len(nonce) + len(encryptedMsg)
	buf := make([]byte, 0, n)

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)
	buf = append(buf, strs[3]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

// 微信的JS签名

func JsSign(noncestr, jsapi_ticket, timestamp, url string) (signature string) {
	params := make(map[string]string)
	params["noncestr"] = noncestr
	params["jsapi_ticket"] = jsapi_ticket
	params["timestamp"] = timestamp
	params["url"] = url
	result := util.JsCommonSign(params)
	return result
}

// 微信位置签名，仅在需要兼容6.02版本之前时使用
func JsAddrSign(noncestr, appId, timestamp, url, accesstoken string) (signature string) {
	params := make(map[string]string)
	params["noncestr"] = noncestr
	params["appId"] = appId
	params["timestamp"] = timestamp
	params["url"] = url
	params["accesstoken"] = accesstoken
	result := util.JsCommonSign(params)
	return result
}

// 通用的签名方法
func JsCommonSign(params map[string]string) (signature string) {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var uri_keys []string
	for _, k := range keys {
		uri_keys = append(uri_keys, k+"="+params[k])
	}
	buf := strings.Join(uri_keys, "&")
	hashsum := sha1.Sum([]byte(buf))
	return hex.EncodeToString(hashsum[:])
}
