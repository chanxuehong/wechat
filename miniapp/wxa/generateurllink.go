package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

type GenerateURLLinkRequest struct {
	// Path 通过 URL Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query 。path 为空时会跳转小程序主页
	Path string `json:"path,omitempty"`
	// Query 通过 URL Link 进入小程序时的query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~%
	Query string `json:"query,omitempty"`
	// IsExpire 生成的scheme码类型，到期失效：true，永久有效：false。
	IsExpire bool `json:"is_expire,omitempty"`
	// ExpireTime 到期失效的scheme码的失效时间，为Unix时间戳。生成的到期失效scheme码在该时间前有效。最长有效期为1年。生成到期失效的scheme时必填。
	ExpireTime int64 `json:"expire_time,omitempty"`
	// ExpireType 默认值0，到期失效的 scheme 码失效类型，失效时间：0，失效间隔天数：1
	ExpireType int `json:"expire_type,omitempty"`
	// ExpireInterval 到期失效的 scheme 码的失效间隔天数。生成的到期失效 scheme 码在该间隔时间到达前有效。最长间隔天数为30天。is_expire 为 true 且 expire_type 为 1 时必填
	ExpireInterval int `json:"expire_interval,omitempty"`
	// EnvVersion 默认值"release"。要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"，仅在微信外打开时生效。
	EnvVersion string `json:"env_version,omitempty"`
	// CloudBase 云开发静态网站自定义 H5 配置参数，可配置中转的云开发 H5 页面。不填默认用官方 H5 页面
	CloudBase *CloudBase `json:"cloud_base,omitempty"`
}

// CloudBase 云开发静态网站自定义 H5 配置参数，可配置中转的云开发 H5 页面。不填默认用官方 H5 页面
type CloudBase struct {
	// Env 云开发环境
	Env string `json:"env,omitempty"`
	// Domain 静态网站自定义域名，不填则使用默认域名
	Domain string `json:"domain,omitempty"`
	// Path 云开发静态网站 H5 页面路径，不可携带 query
	Path string `json:"path,omitempty"`
	// Query 云开发静态网站 H5 页面 query 参数，最大 1024 个字符，只支持数字，大小写英文以及部分特殊字符：`!#$&'()*+,/:;=?@-._~%``
	Query string `json:"query,omitempty"`
	// ResourceAppID 第三方批量代云开发时必填，表示创建该 env 的 appid （小程序/第三方平台）
	ResourceAppID string `json:"resource_appid,omitempty"`
}

// GeneratURLLink 获取 URL Link
func GenerateURLLink(clt *core.Client, req *GenerateURLLinkRequest) (urllink string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/generate_urllink?access_token="
	var result struct {
		core.Error
		URLLink string `json:"url_link"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	urllink = result.URLLink
	return
}
