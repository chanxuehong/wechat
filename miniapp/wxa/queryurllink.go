package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

type URLLinkInfo struct {
	// AppID 小程序 appid
	AppID string `json:"appid,omitempty"`
	// Path 小程序页面路径
	Path string `json:"path,omitempty"`
	// Query 小程序页面query
	Query string `json:"query,omitempty"`
	// CreateTime 创建时间，为 Unix 时间戳
	CreateTime int64 `json:"create_time,omitempty"`
	// ExpireTime 到期失效时间，为 Unix 时间戳，0 表示永久生效
	ExpireTime int64 `json:"expire_time,omitempty"`
	// EnvVersion 要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"
	EnvVersion string `json:"env_version,omitempty"`
	// CloudBase 云开发静态网站自定义 H5 配置参数，可配置中转的云开发 H5 页面。不填默认用官方 H5 页面
	CloudBase *CloudBase `json:"cloud_base,omitempty"`
}

// URLLinkQuota
type URLLinkQuota struct {
	// LongTimeUsed 长期有效 url_link 已生成次数
	LongTimeUsed int64 `json:"long_time_used,omitempty"`
	// LongTimeLimit 长期有效 url_link 生成次数上限
	LongTimeLimit int64 `json:"long_time_limit,omitempty"`
}

type URLLinkResult struct {
	// Info url_link 配置
	Info *URLLinkInfo `json:"url_link_info,omitempty"`
	// Quota quota 配置
	Quota *URLLinkQuota `json:"url_link_quota,omitempty"`
	// VisitOpenID 访问 Link 的用户openid，为空表示未被访问过
	VisitOpenID string `json:"visit_openid,omitempty"`
}

// QueryURLLink 查询 URL Link
func QueryURLLink(clt *core.Client, urllink string) (ret *URLLinkResult, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/query_urllink?access_token="
	var result struct {
		core.Error
		URLLinkResult
	}
	req := map[string]string{"url_link": urllink}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = &result.URLLinkResult
	return
}
