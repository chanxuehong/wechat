package base

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 获取微信服务器IP地址.
//  如果公众号基于安全等考虑，需要获知微信服务器的IP地址列表，以便进行相关限制，可以通过该接口获得微信服务器IP地址列表。
func GetCallbackIP(clt *core.Client) (ipList []string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token="

	var result struct {
		core.Error
		List []string `json:"ip_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	ipList = result.List
	return
}
