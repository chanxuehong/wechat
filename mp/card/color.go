// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"github.com/chanxuehong/wechat/mp"
)

type Color struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// 获得卡券的最新颜色列表，用于卡券创建.
func (clt Client) GetColors() (colors []Color, err error) {
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
