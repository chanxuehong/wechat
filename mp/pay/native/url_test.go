// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package native

import (
	"testing"
)

func TestNativeURL(t *testing.T) {
	appid := "wxf8b4f85f3a794e77"
	appkey := "2Wozy2aksie1puXUBpWD8oZxiD1DfQuEaiC7KcRATv1Ino3mdopKaPGQQ7TtkNySuAmCaDCrw4xhPY5qKTBl7Fzm0RgR3c0WaVYIXZARsxzHV2x7iwPPzOz94dnwPWSn"
	noncestr := "adssdasssd13d"
	productid := "123456"
	var timestamp int64 = 189026618

	_url := NativeURL(appid, noncestr, timestamp, productid, appkey)
	want := "weixin://wxpay/bizpayurl?appid=wxf8b4f85f3a794e77&noncestr=adssdasssd13d&productid=123456&sign=18c6122878f0e946ae294e016eddda9468de80df&timestamp=189026618"

	if _url != want {
		t.Errorf("failed, have %#s\n, want %#s\n", _url, want)
		return
	}
}
