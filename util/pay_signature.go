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
	"errors"
	"fmt"
	"hash"
	"sort"
	"strings"
)

// 微信支付签名, 微信支付接口 2.x 系列基本都是用这种方式签名, 参考微信支付接口文档 v2.7 的 2.7 章节
//  parameters:        待签名的参数集合
//  appKey:            签名的 key
//  signatureKeyName:  parameters 里面 "签名" 的 key name, 不参与签名
//  signMethodKeyName: parameters 里面 "签名方法" 的 key name, 不参与签名
func Pay2Signature(parameters map[string]string, appKey string,
	signatureKeyName, signMethodKeyName string) (signature string, err error) {

	if parameters == nil {
		err = errors.New("parameters == nil")
		return
	}

	signMethod, ok := parameters[signMethodKeyName]
	if !ok {
		err = errors.New("sign method is empty")
		return
	}

	var Hash hash.Hash
	var hashsum []byte

	switch signMethod {
	case "sha1", "SHA1":
		Hash = sha1.New()
		hashsum = make([]byte, sha1.Size*2)

	case "md5", "MD5":
		Hash = md5.New()
		hashsum = make([]byte, md5.Size*2)

	default:
		err = fmt.Errorf(`unknown sign method: %q`, signMethod)
		return
	}

	lowerSignatureKeyName := strings.ToLower(signatureKeyName)
	lowerSignMethodKeyName := strings.ToLower(signMethodKeyName)

	lowerParameters := make(map[string]string, len(parameters)+1)
	for key, value := range parameters {
		lowerParameters[strings.ToLower(key)] = value
	}
	lowerParameters["appkey"] = appKey

	keys := make([]string, 0, len(lowerParameters))
	for key := range lowerParameters {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	hasWrite := false
	for _, key := range keys {
		// 不参与签名
		if key == lowerSignatureKeyName || key == lowerSignMethodKeyName {
			continue
		}

		value := lowerParameters[key]

		if !hasWrite {
			hasWrite = true

			Hash.Write([]byte(key))
			Hash.Write([]byte{'='})
			Hash.Write([]byte(value))
		} else {
			Hash.Write([]byte{'&'})
			Hash.Write([]byte(key))
			Hash.Write([]byte{'='})
			Hash.Write([]byte(value))
		}
	}

	hex.Encode(hashsum, Hash.Sum(nil))

	signature = string(hashsum)
	return
}

// 微信支付签名, 微信支付接口 3.x 系列基本都是用这种方式签名, 参考微信支付接口文档 v3.3.6 的 3.2 章节
//  parameters:        待签名的参数集合
//  appKey:            签名的 key
//  signatureKeyName:  parameters 里面 "签名" 的 key name, 不参与签名, 如果留空则默认为 "sign"
func Pay3Signature(parameters map[string]string, appKey string,
	signatureKeyName string) (signature string, err error) {

	if parameters == nil {
		err = errors.New("parameters == nil")
		return
	}
	if signatureKeyName == "" {
		signatureKeyName = "sign"
	}

	keys := make([]string, 0, len(parameters))
	for key := range parameters {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	Hash := md5.New()
	md5sum := make([]byte, md5.Size*2)
	for _, key := range keys {
		// 不参与签名
		if key == signatureKeyName {
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

	hex.Encode(md5sum, Hash.Sum(nil))
	copy(md5sum, bytes.ToUpper(md5sum))

	signature = string(md5sum)
	return
}
