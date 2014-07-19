// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package wechat

import (
	"testing"
)

func TestWXVersion(t *testing.T) {
	userAgent := `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/5.0`

	ver, err := WXVersion(userAgent)
	if err != nil {
		t.Error(err)
		return
	}
	if ver != 5.0 { // 5.0 可以用 ==
		t.Error("获取了错误的版本号")
		return
	}

	userAgent = `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/5.0x`
	_, err = WXVersion(userAgent)
	if err == nil {
		t.Errorf("从 %#s 获取版本号应该出错, 但是目前却没有错误!", userAgent)
		return
	}

	userAgent = `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/`
	_, err = WXVersion(userAgent)
	if err == nil {
		t.Errorf("从 %#s 获取版本号应该出错, 但是目前却没有错误!", userAgent)
		return
	}

	userAgent = `Mozilla5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit534.46(KHTML,like Geocko)Mobile9B206 MicroMessenger5.0`
	_, err = WXVersion(userAgent)
	if err == nil {
		t.Errorf("从 %#s 获取版本号应该出错, 但是目前却没有错误!", userAgent)
		return
	}
}
