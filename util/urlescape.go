// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package util

// 转义 s string 到 URL 编码格式.
//  NOTE: 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_.~ 不转义,
//        和 net/url.QueryEscape 不同的地方在于 空格 转义成 "%20" 而不是 "+"
func URLEscape(src string) string {
	hexCount := 0
	for i := 0; i < len(src); i++ {
		if shouldEscapeTable[src[i]] {
			hexCount++
		}
	}

	if hexCount == 0 {
		return src
	}

	dst := make([]byte, len(src)+2*hexCount)
	for i, j := 0, 0; i < len(src); i++ {
		switch c := src[i]; {
		case shouldEscapeTable[c]:
			dst[j] = '%'
			dst[j+1] = "0123456789ABCDEF"[c>>4]
			dst[j+2] = "0123456789ABCDEF"[c&0x0f]
			j += 3
		default:
			dst[j] = src[i]
			j++
		}
	}

	return string(dst)
}

// 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_.~ 不转义
var shouldEscapeTable [256]bool

func init() {
	for i := 0; i < 256; i++ {
		c := byte(i)

		if '0' <= c && c <= '9' || 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' {
			//shouldEscapeTable[i] = false
			continue
		}

		switch c {
		case '-', '_', '.', '~':
			//shouldEscapeTable[i] = false
		default:
			shouldEscapeTable[i] = true
		}
	}
}
