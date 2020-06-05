package externalcontact

import (
	"github.com/chanxuehong/wechat/work/core"
)

// EditCorpTag 编辑企业客户标签
// id: 标签或标签组的id列表
// name: 新的标签或标签组名称，最长为30个字符
// order: 标签/标签组的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
func EditCorpTag(clt *core.Client, id string, name string, order uint64) (err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag?access_token="

	var request = struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Order uint64 `json:"order,omitempty"`
	}{
		Id:    id,
		Name:  name,
		Order: order,
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
