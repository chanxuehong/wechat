package core

import (
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

	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf0 [256]byte

	buf1 := buf0[:]
	h := fn()
	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		buf1 = buf1[:0]
		buf1 = append(buf1, k...)
		buf1 = append(buf1, '=')
		buf1 = append(buf1, v...)
		buf1 = append(buf1, '&')
		h.Write(buf1)
	}
	buf1 = buf1[:0]
	buf1 = append(buf1, "key="...)
	buf1 = append(buf1, apiKey...)
	h.Write(buf1)

	var signature []byte
	if size := hex.EncodedLen(h.Size()); size > len(buf0) {
		signature = make([]byte, size)
	} else {
		signature = buf0[:size]
	}
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

	var buf0 [256]byte
	buf1 := buf0[:]

	// appId
	// nonceStr
	// package
	// signType
	// timeStamp
	if appId != "" {
		buf1 = buf1[:0]
		buf1 = append(buf1, "appId="...)
		buf1 = append(buf1, appId...)
		h.Write(buf1)
	}
	if nonceStr != "" {
		buf1 = buf1[:0]
		buf1 = append(buf1, "&nonceStr="...)
		buf1 = append(buf1, nonceStr...)
		h.Write(buf1)
	}
	if packageStr != "" {
		buf1 = buf1[:0]
		buf1 = append(buf1, "&package="...)
		buf1 = append(buf1, packageStr...)
		h.Write(buf1)
	}
	if signType != "" {
		buf1 = buf1[:0]
		buf1 = append(buf1, "&signType="...)
		buf1 = append(buf1, signType...)
		h.Write(buf1)
	}
	if timeStamp != "" {
		buf1 = buf1[:0]
		buf1 = append(buf1, "&timeStamp="...)
		buf1 = append(buf1, timeStamp...)
		h.Write(buf1)
	}
	buf1 = buf1[:0]
	buf1 = append(buf1, "&key="...)
	buf1 = append(buf1, apiKey...)
	h.Write(buf1)

	var signature []byte
	if size := hex.EncodedLen(h.Size()); size > len(buf0) {
		signature = make([]byte, size)
	} else {
		signature = buf0[:size]
	}
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}

// EditAddressSign 收货地址共享接口签名
func EditAddressSign(appId, url, timestamp, nonceStr, accessToken string) string {
	var buf0 [256]byte

	buf := buf0[:]
	h := sha1.New()

	// accesstoken
	// appid
	// noncestr
	// timestamp
	// url
	buf = buf[:0]
	buf = append(buf, "accesstoken="...)
	buf = append(buf, accessToken...)
	h.Write(buf)

	buf = buf[:0]
	buf = append(buf, "&appid="...)
	buf = append(buf, appId...)
	h.Write(buf)

	buf = buf[:0]
	buf = append(buf, "&noncestr="...)
	buf = append(buf, nonceStr...)
	h.Write(buf)

	buf = buf[:0]
	buf = append(buf, "&timestamp="...)
	buf = append(buf, timestamp...)
	h.Write(buf)

	buf = buf[:0]
	buf = append(buf, "&url="...)
	buf = append(buf, url...)
	h.Write(buf)

	signature := buf0[:sha1.Size*2]
	hex.Encode(signature, h.Sum(nil))
	return string(signature)
}
