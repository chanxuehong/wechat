// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package sns

import (
	"testing"
)

func TestHeadImageSize(t *testing.T) {
	urls := []string{
		"http://www.xxx.com/abcd/0",
		"http://www.xxx.com/abcd/46",
		"http://www.xxx.com/abcd/64",
		"http://www.xxx.com/abcd/96",
		"http://www.xxx.com/abcd/132",
	}
	expect := []int{
		640,
		46,
		64,
		96,
		132,
	}

	for i := 0; i < len(urls); i++ {
		size, err := HeadImageSize(urls[i])
		if err != nil {
			t.Error(err)
			continue
		}
		if size != expect[i] {
			t.Errorf("HeadImageSize(%#s):\nhave %d\nwant %d\n", urls[i], size, expect[i])
			continue
		}
	}
}
