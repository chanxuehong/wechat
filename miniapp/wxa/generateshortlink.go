package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

type GenerateShortLinkRequest struct {
	// PageURL 通过 Short Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，可携带 query，最大1024个字符
	PageURL string `json:"page_url,omitempty"`
	// PageTitle 页面标题，不能包含违法信息，超过20字符会用... 截断代替
	PageTitle string `json:"page_title,omitempty"`
	// IsPermanent 默认值false。生成的 Short Link 类型，短期有效：false，永久有效：true
	IsPermanent bool `json:"is_permanent,omitempty"`
}

// GenerateShortLink 获取 Short Link
func GenerateShortLink(clt *core.Client, req *GenerateShortLinkRequest) (link string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/genwxashortlink?access_token="
	var result struct {
		core.Error
		Link string `json:"link"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	link = result.Link
	return
}
