package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

type ModifyDomainRequest struct {
	Action          string   `json:"action"`                    // add添加, delete删除, set覆盖, get获取。当参数是get时不需要填四个域名字段
	RequestDomain   []string `json:"requestdomain,omitempty"`   // request合法域名，当action参数是get时不需要此字段
	WsRequestDomain []string `json:"wsrequestdomain,omitempty"` // socket合法域名，当action参数是get时不需要此字段
	UploadDomain    []string `json:"uploaddomain,omitempty"`    // uploadFile合法域名，当action参数是get时不需要此字段
	DownloadDomain  []string `json:"downloaddomain,omitempty"`  // downloadFile合法域名，当action参数是get时不需要此字段
	WebViewDomain   []string `json:"webviewdomain,omitempty"`   // request合法域名，当action参数是get时不需要此字段
}

type ModifyDomainResponse struct {
	RequestDomain   []string `json:"requestdomain,omitempty"`   // request合法域名，当action参数是get时不需要此字段
	WsRequestDomain []string `json:"wsrequestdomain,omitempty"` // socket合法域名，当action参数是get时不需要此字段
	UploadDomain    []string `json:"uploaddomain,omitempty"`    // uploadFile合法域名，当action参数是get时不需要此字段
	DownloadDomain  []string `json:"downloaddomain,omitempty"`  // downloadFile合法域名，当action参数是get时不需要此字段
}

// 授权给第三方的小程序，其服务器域名只可以为第三方的服务器，当小程序通过第三方发布代码上线后，小程序原先自己配置的服务器域名将被删除，只保留第三方平台的域名，所以第三方平台在代替小程序发布代码之前，需要调用接口为小程序添加第三方自身的域名。提示：需要先将域名登记到第三方平台的小程序服务器域名中，才可以调用接口进行配置。
func ModifyDomain(clt *core.Client, req *ModifyDomainRequest) (resp *ModifyDomainResponse, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/modify_domain?access_token="
	var result struct {
		core.Error
		RequestDomain   []string `json:"requestdomain,omitempty"`
		WsRequestDomain []string `json:"wsrequestdomain,omitempty"`
		UploadDomain    []string `json:"uploaddomain,omitempty"`
		DownloadDomain  []string `json:"downloaddomain,omitempty"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return &ModifyDomainResponse{
		RequestDomain:   result.RequestDomain,
		WsRequestDomain: result.WsRequestDomain,
		UploadDomain:    result.UploadDomain,
		DownloadDomain:  result.DownloadDomain,
	}, nil
}
