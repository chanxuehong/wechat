package externalcontact

import (
	"fmt"
	"net/url"

	"github.com/chanxuehong/wechat/work/core"
)

// List 获取客户列表.
// userid: 企业成员的userid
func List(clt *core.Client, userId string) (list []string, err error) {
	incompleteURL := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?userid=%s&access_token=", url.QueryEscape(userId))

	var result struct {
		core.Error
		ExternalUserIds []string `json:"external_userid"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.ExternalUserIds
	return
}
