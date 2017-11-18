// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://gopkg.in/chanxuehong/wechat.v1 for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package testwhitelist

import (
	"gopkg.in/chanxuehong/wechat.v1/mp"
)

type SetParameters struct {
	OpenIdList   []string `json:"openid,omitempty"`   // 测试的openid列表
	UserNameList []string `json:"username,omitempty"` // 测试的微信号列表
}

// 设置测试白名单
func Set(clt *mp.Client, para *SetParameters) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/testwhitelist/set?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
