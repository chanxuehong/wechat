package wxa

import (
	"github.com/chanxuehong/wechat/component/core"
)

// 获取模板库某个模板标题下关键词库
func GetLibraryTemplateKeywords(clt *core.Client, id string) (keywords []TemplateKeyword, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/library/get?access_token="
	var result struct {
		core.Error
		Id       uint64            `json:"id"`
		Title    string            `json:"title"`
		Keywords []TemplateKeyword `json:"keyword_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.Keywords, nil
}
