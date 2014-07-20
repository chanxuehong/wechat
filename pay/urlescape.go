// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

// 转义 s string 到 URL 编码格式.
//  NOTE: 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_.~ 不转义,
//  和 net/url.QueryEscape 不同的地方在于空格转义成 %20 而不是 +
func URLEscape(s string) string {
	hexCount := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscapeMap[c] {
			hexCount++
		}
	}

	if hexCount == 0 {
		return s
	}

	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case shouldEscapeMap[c]:
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&0x0f]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}

	return string(t)
}

// 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_.~ 不转义
var shouldEscapeMap [256]bool

func init() {
	for i := 0; i < 256; i++ {
		if '0' <= i && i <= '9' || 'A' <= i && i <= 'Z' || 'a' <= i && i <= 'z' {
			shouldEscapeMap[i] = false
			continue
		}

		switch i {
		case '-', '_', '.', '~':
			shouldEscapeMap[i] = false

		default:
			shouldEscapeMap[i] = true
		}
	}
}
