package jssdk

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

// JS-SDK wx.config 的参数签名.
func WXConfigSign(jsapiTicket, nonceStr, timestamp, url string) (signature string) {
	if i := strings.IndexByte(url, '#'); i >= 0 {
		url = url[:i]
	}

	n := len("jsapi_ticket=") + len(jsapiTicket) +
		len("&noncestr=") + len(nonceStr) +
		len("&timestamp=") + len(timestamp) +
		len("&url=") + len(url)
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

// JS-SDK 卡券 API 参数签名.
func CardSign(strs []string) (signature string) {
	sort.Strings(strs)

	h := sha1.New()

	bufw := bufio.NewWriterSize(h, 128) // sha1.BlockSize 的整数倍
	for _, str := range strs {
		bufw.WriteString(str)
	}
	bufw.Flush()

	return hex.EncodeToString(h.Sum(nil))
}
