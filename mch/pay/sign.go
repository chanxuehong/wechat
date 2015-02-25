// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"sort"
)

// 微信支付签名.
//  parameters: 待签名的参数集合
//  apiKey:     API密钥
//  h:          hash.Hash, 如果 h == nil 则默认用 md5.New()
func Sign(parameters map[string]string, apiKey string, h hash.Hash) string {
	ks := make([]string, 0, len(parameters))
	for k := range parameters {
		if k == "sign" {
			continue
		}
		ks = append(ks, k)
	}
	sort.Strings(ks)

	if h == nil {
		h = md5.New()
	}
	signature := make([]byte, h.Size()*2)

	for _, k := range ks {
		v := parameters[k]
		if v == "" {
			continue
		}
		h.Write([]byte(k))
		h.Write([]byte{'='})
		h.Write([]byte(v))
		h.Write([]byte{'&'})
	}
	h.Write([]byte("key="))
	h.Write([]byte(apiKey))

	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}
