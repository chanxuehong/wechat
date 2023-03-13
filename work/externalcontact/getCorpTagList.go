package externalcontact

import (
	"github.com/bububa/wechat/work/core"
)

// GetCorpTagList 获取企业标签库
// tag_id: 要查询的标签id，如果不填则获取该企业的所有客户标签，目前暂不支持标签组id
func GetCorpTagList(clt *core.Client, tagIds []string) (rslt []TagGroup, err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list?access_token="

	request := make(map[string][]string, 1)
	if len(tagIds) > 0 {
		request["tag_id"] = tagIds
	}
	var result struct {
		core.Error
		Groups []TagGroup `json:"tag_group"`
	}
	if err = clt.PostJSON(incompleteURL, request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	rslt = result.Groups
	return
}
