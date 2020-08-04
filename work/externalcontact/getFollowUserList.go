package externalcontact

import (
	"github.com/chanxuehong/wechat/work/core"
)

// GetFollowUserList 获取客户列表.
func GetFollowUserList(clt *core.Client) (users []string, err error) {
	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list?access_token="

	var result struct {
		core.Error
		Users []string `json:"follow_user"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	users = result.Users
	return
}
