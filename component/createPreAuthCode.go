package component

import (
	"github.com/bububa/wechat/component/core"
)

// 该API用于获取预授权码。预授权码用于公众号或小程序授权时的第三方平台方安全验证.
func CreatePreAuthCode(clt *core.Client, appId string) (preAuthCode string, expiresIn int64, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token="
	var result struct {
		core.Error
		PreAuthCode string `json:"pre_auth_code"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	req := map[string]string{"component_appid": appId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.PreAuthCode, result.ExpiresIn, nil
}
