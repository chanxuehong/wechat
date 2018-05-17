package util

import (
	"testing"
)

func TestWXVersion(t *testing.T) {
	userAgent := `Mozilla/5.0 (Linux; Android 4.4.4; Che1-CL10 Build/Che1-CL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.49 Mobile MQQBrowser/6.2 TBS/043128 Safari/537.36 MicroMessenger/6.5.7.1041 NetType/WIFI Language/zh_CN`
	x, y, z, w, err := WXVersion(userAgent)
	if err != nil {
		t.Error(err)
		return
	}
	if x != 6 || y != 5 || z != 7 || w != 1041 {
		t.Error("获取了错误的版本号")
		return
	}

	userAgent = `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/5.3`
	x, y, z, w, err = WXVersion(userAgent)
	if err != nil {
		t.Error(err)
		return
	}
	if x != 5 || y != 3 || z != 0 {
		t.Error("获取了错误的版本号")
		return
	}

	userAgent = `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/5.3.1`
	x, y, z, w, err = WXVersion(userAgent)
	if err != nil {
		t.Error(err)
		return
	}
	if x != 5 || y != 3 || z != 1 {
		t.Error("获取了错误的版本号")
		return
	}

	userAgent = `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/5.3.1.5`
	x, y, z, w, err = WXVersion(userAgent)
	if err != nil {
		t.Error(err)
		return
	}
	if x != 5 || y != 3 || z != 1 {
		t.Error("获取了错误的版本号")
		return
	}

	userAgent = `Mozilla5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit534.46(KHTML,like Geocko)Mobile9B206 MicroMessenger5.0`
	_, _, _, _, err = WXVersion(userAgent)
	if err == nil {
		t.Errorf("从 %#q 获取版本号应该出错, 但是目前却没有错误!", userAgent)
		return
	}

	userAgent = `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/`
	_, _, _, _, err = WXVersion(userAgent)
	if err == nil {
		t.Errorf("从 %#q 获取版本号应该出错, 但是目前却没有错误!", userAgent)
		return
	}

	userAgent = `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/5x`
	x, y, z, w, err = WXVersion(userAgent)
	if err != nil {
		t.Error(err)
		return
	}
	if x != 5 || y != 0 || z != 0 || w != 0 {
		t.Error("获取了错误的版本号")
		return
	}

	userAgent = `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/5.3x`
	x, y, z, w, err = WXVersion(userAgent)
	if err != nil {
		t.Error(err)
		return
	}
	if x != 5 || y != 3 || z != 0 || w != 0 {
		t.Error("获取了错误的版本号")
		return
	}

	userAgent = `Mozilla/5.0(iphone;CPU iphone OS 5_1_1 like Mac OS X) AppleWebKit/534.46(KHTML,like Geocko)Mobile/9B206 MicroMessenger/5.3.1x`
	x, y, z, w, err = WXVersion(userAgent)
	if err != nil {
		t.Error(err)
		return
	}
	if x != 5 || y != 3 || z != 1 || w != 0 {
		t.Error("获取了错误的版本号")
		return
	}

	userAgent = `Mozilla/5.0 (Linux; Android 8.0.0; MI 6 Build/OPR1.170623.027; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/62.0.3202.84 Mobile Safari/537.36 MicroMessenger/6.6.2.1240(0x26060240) NetType/WIFI Language/zh_CN`
	x, y, z, w, err = WXVersion(userAgent)
	if err != nil {
		t.Error(err)
		return
	}
	if x != 6 || y != 6 || z != 2 || w != 1240 {
		t.Error("获取了错误的版本号")
		return
	}
}
