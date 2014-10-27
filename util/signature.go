// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package util

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"sort"
)

// 微信公众号 明文模式签名
func Signature(token, timestamp, nonce string) (signature string) {
	strArray := sort.StringSlice{token, timestamp, nonce}
	strArray.Sort()

	n := len(token) + len(timestamp) + len(nonce)
	buf := make([]byte, 0, n)

	buf = append(buf, strArray[0]...)
	buf = append(buf, strArray[1]...)
	buf = append(buf, strArray[2]...)

	hashSumArray := sha1.Sum(buf)
	return hex.EncodeToString(hashSumArray[:])
}

// 微信公众号/企业号 密文模式签名
func MsgSignature(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strArray := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strArray.Sort()

	n := len(token) + len(timestamp) + len(nonce) + len(encryptedMsg)
	buf := make([]byte, 0, n)

	buf = append(buf, strArray[0]...)
	buf = append(buf, strArray[1]...)
	buf = append(buf, strArray[2]...)
	buf = append(buf, strArray[3]...)

	hashSumArray := sha1.Sum(buf)
	return hex.EncodeToString(hashSumArray[:])
}

// 微信支付签名
//  NOTE: 一般是 parameters 除了 sign 参数之外，其他参数都设置好才调用此函数
func PaySignature(parameters map[string]string, appKey string) (signature string) {
	keys := make([]string, 0, len(parameters))
	for key := range parameters {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	Hash := md5.New()
	for _, key := range keys {
		if key == "sign" {
			continue
		}

		value := parameters[key]

		if len(value) > 0 {
			Hash.Write([]byte(key))
			Hash.Write([]byte{'='})
			Hash.Write([]byte(value))
			Hash.Write([]byte{'&'})
		}
	}
	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	md5sum := make([]byte, md5.Size*2)
	hex.Encode(md5sum, Hash.Sum(nil))
	copy(md5sum, bytes.ToUpper(md5sum))

	return string(md5sum)
}
