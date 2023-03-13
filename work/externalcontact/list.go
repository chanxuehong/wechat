package externalcontact

import (
	"net/url"

	"github.com/bububa/wechat/util"
	"github.com/bububa/wechat/work/core"
)

// List 获取客户列表.
// userid: 企业成员的userid
func List(clt *core.Client, userId string) (list []string, err error) {
	incompleteURL := util.StringsJoin("https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?userid=", url.QueryEscape(userId), "&access_token=")

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
