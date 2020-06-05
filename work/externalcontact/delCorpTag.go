package externalcontact

import (
	"github.com/chanxuehong/wechat/work/core"
)

// DelCorpTag 删除企业客户标签
// tag_id: 标签的id列表
// group_id: 标签组的id列表
func DelCorpTag(clt *core.Client, tagIds []string, groupIds []string) (err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag?access_token="

	var request = struct {
		TagIds   []string `json:"tag_id,omitempty"`
		GroupIds []string `json:"group_id,omitempty"`
	}{
		TagIds:   tagIds,
		GroupIds: groupIds,
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
