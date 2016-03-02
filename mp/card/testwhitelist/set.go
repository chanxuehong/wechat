package testwhitelist

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type SetParameters struct {
	OpenIdList   []string `json:"openid,omitempty"`   // 测试的openid列表
	UserNameList []string `json:"username,omitempty"` // 测试的微信号列表
}

// 设置测试白名单
func Set(clt *core.Client, para *SetParameters) (err error) {
	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/card/testwhitelist/set?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
