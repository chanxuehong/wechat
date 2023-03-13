package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

// 授权给第三方的小程序，其业务域名只可以为第三方的服务器，当小程序通过第三方发布代码上线后，小程序原先自己配置的业务域名将被删除，只保留第三方平台的域名，所以第三方平台在代替小程序发布代码之前，需要调用接口为小程序添加业务域名。提示：1、需要先将域名登记到第三方平台的小程序业务域名中，才可以调用接口进行配置。 2、为授权的小程序配置域名时支持配置子域名，例如第三方登记的业务域名如为qq.com，则可以直接将qq.com及其子域名（如xxx.qq.com）也配置到授权的小程序中。
func SetWebViewDomain(clt *core.Client, req *ModifyDomainRequest) (domain []string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/setwebviewdomain?access_token="
	var result struct {
		core.Error
		WebViewDomain []string `json:"webviewdomain,omitempty"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.WebViewDomain, nil
}
