// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package js

// -.0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~ 不转义
func urlEscape(s string) string {
	hexCount := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
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
		case shouldEscape(c):
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

func shouldEscape(c byte) bool {
	if 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || 'A' <= c && c <= 'Z' {
		return false
	}

	switch c {
	case '-', '_', '.', '~':
		return false

	default:
		return true
	}
}
