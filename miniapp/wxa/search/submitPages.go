package search

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type SubmitPagesRequest struct {
	Pages []Page `json:"pages"` // 小程序页面信息列表
}

type Page struct {
	Path  string `json:"path"`  // 页面路径
	Query string `json:"query"` // 页面参数
}

// 小程序开发者可以通过本接口提交小程序页面url及参数信息，让微信可以更及时的收录到小程序的页面信息，开发者提交的页面信息将可能被用于小程序搜索结果展示。
func SubmitPages(clt *core.Client, req *SubmitPagesRequest) error {
	const incompleteURL = "api.weixin.qq.com/wxa/search/wxaapi_submitpages?access_token="
	var result struct {
		core.Error
	}
	if err := clt.PostJSON(incompleteURL, req, &result); err != nil {
		return err
	}
	if result.ErrCode != core.ErrCodeOK {
		return &result.Error
	}
	return nil
}
