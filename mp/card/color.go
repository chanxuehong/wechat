<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package card

import (
	"github.com/chanxuehong/wechat/mp"
=======
package card

import (
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
)

type Color struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// 获取卡券最新的颜色列表.
<<<<<<< HEAD
func GetColors(clt *mp.Client) (colors []Color, err error) {
	var result struct {
		mp.Error
=======
func GetColors(clt *core.Client) (colors []Color, err error) {
	var result struct {
		core.Error
>>>>>>> github/v2
		Colors []Color `json:"colors"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/getcolors?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

<<<<<<< HEAD
	if result.ErrCode != mp.ErrCodeOK {
=======
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result.Error
		return
	}
	colors = result.Colors
	return
}
