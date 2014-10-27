// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package util

import (
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
