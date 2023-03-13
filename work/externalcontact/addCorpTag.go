package externalcontact

import (
	"github.com/bububa/wechat/work/core"
)

// AddCorpTag 添加企业客户标签
func AddCorpTag(clt *core.Client, tagGroup *TagGroup) (rslt *TagGroup, err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token="

	var result struct {
		core.Error
		Group *TagGroup `json:"tag_group"`
	}
	if err = clt.PostJSON(incompleteURL, tagGroup, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	rslt = result.Group
	return
}
