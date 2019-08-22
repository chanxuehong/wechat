package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type CommitRequest struct {
	TemplateId  uint64 `json:"template_id"`  // 代码库中的代码模版ID
	ExtJson     string `json:"ext_json"`     // 第三方自定义的配置
	UserVersion string `json:"user_version"` // 代码版本号，开发者可自定义（长度不要超过64个字符）
	UserDesc    string `json:"user_desc"`    // 代码描述，开发者可自定义
}

// 第三方平台在开发者工具上开发完成后，可点击上传，代码将上传到开放平台草稿箱中，第三方平台可选择将代码添加到模板中，获得代码模版ID后，可调用以下接口进行代码管理。
func Commit(clt *core.Client, req *CommitRequest) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/commit?access_token="
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
