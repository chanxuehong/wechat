package search

import (
	"github.com/bububa/wechat/mp/core"
)

type PageItem struct {
	Title string `json:"title"`                 // 小程序页面标题
	Desc  string `json:"description,omitempty"` // 小程序页面摘要
	Image string `json:"image,omitempty"`       // 小程序页面代表图
	Path  string `json:"path"`                  // 小程序页面路径
}

type SiteSearchResponse struct {
	Items        []PageItem `json:"items,omitempty"`
	HasNextPage  int        `json:"has_next_page,omitempty"`
	NextPageInfo string     `json:"next_page_info,omitempty"`
	Total        int        `json:"hit_count,omitempty"`
}

// 小程序内部搜索API提供针对页面的查询能力，小程序开发者输入搜索词后，将返回自身小程序和搜索词相关的页面。因此，利用该接口，开发者可以查看指定内容的页面被微信平台的收录情况；同时，该接口也可供开发者在小程序内应用，给小程序用户提供搜索能力。
func SiteSearch(clt *core.Client, keyword string, nextPageInfo string) (*SiteSearchResponse, error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/sitesearch?access_token="
	var result struct {
		core.Error
		Items        []PageItem `json:"items"`
		HasNextPage  int        `json:"has_next_page"`
		NextPageInfo string     `json:"next_page_info"`
		Total        int        `json:"hit_count"`
	}
	req := map[string]string{"keyword": keyword, "next_page_info": nextPageInfo}
	if err := clt.PostJSON(incompleteURL, req, &result); err != nil {
		return nil, err
	}
	if result.ErrCode != core.ErrCodeOK {
		return nil, &result.Error
	}
	return &SiteSearchResponse{
		Total:        result.Total,
		Items:        result.Items,
		HasNextPage:  result.HasNextPage,
		NextPageInfo: result.NextPageInfo,
	}, nil
}
