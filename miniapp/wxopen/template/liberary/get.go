package liberary

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 获取模板库某个模板标题下关键词库
func Get(clt *core.Client, templateId string) (tmp Template, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/library/get?access_token="
	var result struct {
		core.Error
		Id       string    `json:"id"`
		Title    string    `json:"title"`
		Keywords []Keyword `json:"keyword_list"`
	}
	req := map[string]string{"id": templateId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return Template{
		Id:       result.Id,
		Title:    result.Title,
		Keywords: result.Keywords,
	}, nil
}
