package core

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"sort"
)

// Sign 微信支付签名.
//  params: 待签名的参数集合
//  apiKey: api密钥
//  fn:     func() hash.Hash, 如果为 nil 则默认用 md5.New
func Sign(params map[string]string, apiKey string, fn func() hash.Hash) string {
	if fn == nil {
		fn = md5.New
	}
	h := fn()
	bufw := bufio.NewWriterSize(h, 128)

	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		bufw.WriteString(k)
		bufw.WriteByte('=')
		bufw.WriteString(v)
		bufw.WriteByte('&')
	}
	bufw.WriteString("key=")
	bufw.WriteString(apiKey)

	bufw.Flush()
	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}

// jssdk 支付签名, signType 只支持 "MD5", "SHA1", 传入其他的值会 panic.
func JsapiSign(appId, timeStamp, nonceStr, packageStr, signType string, apiKey string) string {
	var h hash.Hash
	switch signType {
	case "MD5":
		h = md5.New()
	case "SHA1":
		h = sha1.New()
	default:
		panic("unsupported signType")
	}
	bufw := bufio.NewWriterSize(h, 128)

	// appId
	// nonceStr
	// package
	// signType
	// timeStamp
	bufw.WriteString("appId=")
	bufw.WriteString(appId)
	bufw.WriteString("&nonceStr=")
	bufw.WriteString(nonceStr)
	bufw.WriteString("&package=")
	bufw.WriteString(packageStr)
	bufw.WriteString("&signType=")
	bufw.WriteString(signType)
	bufw.WriteString("&timeStamp=")
	bufw.WriteString(timeStamp)
	bufw.WriteString("&key=")
	bufw.WriteString(apiKey)

	bufw.Flush()
	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}

// EditAddressSign 收货地址共享接口签名
func EditAddressSign(appId, url, timestamp, nonceStr, accessToken string) string {
	h := sha1.New()
	bufw := bufio.NewWriterSize(h, 128)

	// accesstoken
	// appid
	// noncestr
	// timestamp
	// url
	bufw.WriteString("accesstoken=")
	bufw.WriteString(accessToken)
	bufw.WriteString("&appid=")
	bufw.WriteString(appId)
	bufw.WriteString("&noncestr=")
	bufw.WriteString(nonceStr)
	bufw.WriteString("&timestamp=")
	bufw.WriteString(timestamp)
	bufw.WriteString("&url=")
	bufw.WriteString(url)

	bufw.Flush()
	return hex.EncodeToString(h.Sum(nil))
}
