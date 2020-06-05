package externalcontact

import (
	"github.com/chanxuehong/wechat/work/core"
)

// MarkTag 编辑客户企业标签
// userid: 添加外部联系人的userid
// external_userid: 外部联系人userid
// add_tag: 要标记的标签列表
// remove_tag: 要移除的标签列表
func MarkTag(clt *core.Client, userId string, externalUserId string, addTags []string, removeTags []string) (err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag?access_token="

	var request = struct {
		UserId         string   `json:"userid,omitempty"`
		ExternalUserId string   `json:"external_userid,omitempty"`
		AddTags        []string `json:"add_tag,omitempty"`
		RemoveTags     []string `json:"remove_tag,omitempty"`
	}{
		UserId:         userId,
		ExternalUserId: externalUserId,
		AddTags:        addTags,
		RemoveTags:     removeTags,
	}

	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}
