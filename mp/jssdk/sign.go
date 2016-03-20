<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package jssdk

import (
	"crypto/sha1"
	"encoding/hex"
)

// 微信 js-sdk wx.config 的参数签名.
=======
package jssdk

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"sort"
)

// JS-SDK wx.config 的参数签名.
>>>>>>> github/v2
func WXConfigSign(jsapiTicket, nonceStr, timestamp, url string) (signature string) {
	n := len("jsapi_ticket=") + len(jsapiTicket) +
		len("&noncestr=") + len(nonceStr) +
		len("&timestamp=") + len(timestamp) +
		len("&url=") + len(url)
<<<<<<< HEAD

=======
>>>>>>> github/v2
	buf := make([]byte, 0, n)

	buf = append(buf, "jsapi_ticket="...)
	buf = append(buf, jsapiTicket...)
	buf = append(buf, "&noncestr="...)
	buf = append(buf, nonceStr...)
	buf = append(buf, "&timestamp="...)
	buf = append(buf, timestamp...)
	buf = append(buf, "&url="...)
	buf = append(buf, url...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}
<<<<<<< HEAD
=======

// stringWriter is the interface that wraps the WriteString method.
type stringWriter interface {
	WriteString(s string) (n int, err error)
}

// JS-SDK 卡券 API 参数签名.
func CardSign(strs []string) (signature string) {
	sort.Strings(strs)

	h := sha1.New()
	if sw, ok := h.(stringWriter); ok {
		for _, str := range strs {
			sw.WriteString(str)
		}
	} else {
		bufw := bufio.NewWriterSize(h, 256)
		for _, str := range strs {
			bufw.WriteString(str)
		}
		bufw.Flush()
	}
	return hex.EncodeToString(h.Sum(nil))
}
>>>>>>> github/v2
