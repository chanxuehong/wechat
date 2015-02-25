// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package account

import (
	"github.com/chanxuehong/wechat/mp"
)

// 将一条长链接转成短链接.
//  主要使用场景：
//  开发者用于生成二维码的原链接（商品、支付二维码等）太长导致扫码速度和成功率下降，
//  将原长链接通过此接口转成短链接再生成二维码将大大提升扫码速度和成功率。
func (clt *Client) ShortURL(LongURL string) (ShortURL string, err error) {
	var request = struct {
		Action  string `json:"action"`
		LongURL string `json:"long_url"`
	}{
		Action:  "long2short",
		LongURL: LongURL,
	}

	var result struct {
		mp.Error
		ShortURL string `json:"short_url"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/shorturl?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	ShortURL = result.ShortURL
	return
}
