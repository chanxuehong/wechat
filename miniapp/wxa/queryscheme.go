package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

type SchemeInfo struct {
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
}

// QueryScheme 查询 scheme 码
func QueryScheme(clt *core.Client, scheme string) (openid string, schemeInfo *SchemeInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/queryscheme?access_token="
	var result struct {
		core.Error
		// VisitOpenID 访问 scheme 的用户openid，为空表示未被访问过
		VisitOpenID string `json:"visit_openid"`
		// SchemeInfo scheme 配置
		SchemeInfo *SchemeInfo `json:"scheme_info,omitempty"`
	}
	req := map[string]string{"scheme": scheme}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	openid = result.VisitOpenID
	schemeInfo = result.SchemeInfo
	return
}
