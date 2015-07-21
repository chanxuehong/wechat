// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"errors"

	"github.com/chanxuehong/wechat/mp"
)

type TestWhiteListSetParameters struct {
	OpenIdList   []string `json:"openid,omitempty"`   // 测试的openid 列表
	UserNameList []string `json:"username,omitempty"` // 测试的微信号列表
}

// 设置测试用户白名单.
//  由于卡券有审核要求, 为方便公众号调试, 可以设置一些测试帐号, 这些帐号可领取未通过审核的卡券, 体验整个流程.
//  注: 同时支持"openid", "username"两种字段设置白名单, 总数上限为10 个.
func (clt *Client) TestWhiteListSet(para *TestWhiteListSetParameters) (err error) {
	if para == nil {
		return errors.New("nil TestWhiteListSetParameters")
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/testwhitelist/set?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
