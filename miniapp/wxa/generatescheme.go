package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type GenerateSchemeRequest struct {
	JumpWxa JumpWxa `json:"jump_wxa"` // 跳转到的目标小程序信息
	// IsExpire 生成的scheme码类型，到期失效：true，永久有效：false。
	IsExpire bool `json:"is_expire,omitempty"`
	// ExpireTime 到期失效的scheme码的失效时间，为Unix时间戳。生成的到期失效scheme码在该时间前有效。最长有效期为1年。生成到期失效的scheme时必填。
	ExpireTime int64 `json:"expire_time,omitempty"`
	// ExpireType 默认值0，到期失效的 scheme 码失效类型，失效时间：0，失效间隔天数：1
	ExpireType int `json:"expire_type,omitempty"`
	// ExpireInterval 到期失效的 scheme 码的失效间隔天数。生成的到期失效 scheme 码在该间隔时间到达前有效。最长间隔天数为30天。is_expire 为 true 且 expire_type 为 1 时必填
	ExpireInterval int `json:"expire_interval,omitempty"`
}

type JumpWxa struct {
	// Path 通过scheme码进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带query。path为空时会跳转小程序主页。
	Path string `json:"path,omitempty"`
	// Query 通过scheme码进入小程序时的query，最大128个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~
	Query string `json:"query,omitempty"`
	// EnvVersion 默认值"release"。要打开的小程序版本。正式版为"release"，体验版为"trial"，开发版为"develop"，仅在微信外打开时生效。
	EnvVersion string `json:"env_version,omitempty"`
}

// 获取小程序scheme码，适用于短信、邮件、外部网页等拉起小程序的业务场景。通过该接口，可以选择生成到期失效和永久有效的小程序码，目前仅针对国内非个人主体的小程序开放
func GenerateScheme(clt *core.Client, req *GenerateSchemeRequest) (openlink string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/generatescheme?access_token="
	var result struct {
		core.Error
		Openlink string `json:"openlink"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.Openlink, nil
}
