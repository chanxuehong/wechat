package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type GenerateSchemeRequest struct {
	JumpWxa    JumpWxa `json:"jump_wxa"`              // 跳转到的目标小程序信息
	IsExpire   bool    `json:"is_expire,omitempty"`   // 生成的scheme码类型，到期失效：true，永久有效：false。
	ExpireTime int64   `json:"expire_time,omitempty"` // 到期失效的scheme码的失效时间，为Unix时间戳。生成的到期失效scheme码在该时间前有效。最长有效期为1年。生成到期失效的scheme时必填。
}

type JumpWxa struct {
	Path  string `json:"path"`  // 通过scheme码进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带query。path为空时会跳转小程序主页。
	Query string `json:"query"` // 通过scheme码进入小程序时的query，最大128个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~
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
