package externalcontact

import (
	"net/url"

	"github.com/bububa/wechat/util"
	"github.com/bububa/wechat/work/core"
)

// Get 获取客户详情.
// external_userid: 外部联系人的userid，注意不是企业成员的帐号
func Get(clt *core.Client, externalUserId string) (contact *ExternalContact, followUsers []FollowUser, err error) {
	incompleteURL := util.StringsJoin("https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?external_userid=", url.QueryEscape(externalUserId), "&access_token=")

	var result struct {
		core.Error
		ExternalContact *ExternalContact `json:"external_contact"`
		FollowUsers     []FollowUser     `json:"follow_user"` // 添加了此外部联系人的企业成员
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	contact = result.ExternalContact
	followUsers = result.FollowUsers
	return
}
