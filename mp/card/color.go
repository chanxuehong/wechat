package card

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type Color struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// 获取卡券最新的颜色列表.
func GetColors(clt *core.Client) (colors []Color, err error) {
	var result struct {
		core.Error
		Colors []Color `json:"colors"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/getcolors?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	colors = result.Colors
	return
}
