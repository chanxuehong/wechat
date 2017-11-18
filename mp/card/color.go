// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://gopkg.in/chanxuehong/wechat.v1 for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package card

import (
	"gopkg.in/chanxuehong/wechat.v1/mp"
)

type Color struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// 获取卡券最新的颜色列表.
func GetColors(clt *mp.Client) (colors []Color, err error) {
	var result struct {
		mp.Error
		Colors []Color `json:"colors"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/getcolors?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	colors = result.Colors
	return
}
