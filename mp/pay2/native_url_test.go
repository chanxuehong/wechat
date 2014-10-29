// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"testing"
)

func TestNativeURL(t *testing.T) {
	appid := "wxf8b4f85f3a794e77"
	appkey := "2Wozy2aksie1puXUBpWD8oZxiD1DfQuEaiC7KcRATv1Ino3mdopKaPGQQ7TtkNySuAmCaDCrw4xhPY5qKTBl7Fzm0RgR3c0WaVYIXZARsxzHV2x7iwPPzOz94dnwPWSn"
	noncestr := "adssdasssd13d"
	productid := "123456"
	timestamp := "189026618"

	have := NativeURL(appid, appkey, noncestr, timestamp, productid)
	want := "weixin://wxpay/bizpayurl?sign=18c6122878f0e946ae294e016eddda9468de80df&appid=wxf8b4f85f3a794e77&productid=123456&timestamp=189026618&noncestr=adssdasssd13d"

	if have != want {
		t.Errorf("failed, \r\nhave %#s, \r\nwant %#s", have, want)
		return
	}
}
