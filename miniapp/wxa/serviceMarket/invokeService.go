package serviceMarket

import (
	"github.com/bububa/wechat/mp/core"
	"github.com/bububa/wechat/util"
)

type InvokeServiceRequest struct {
	Service     string      `json:"service"`
	Api         string      `json:"api"`
	Data        interface{} `json:"data"`
	ClientMsgId string      `json:"client_msg_id"`
}

// 调用服务平台提供的服务。
func InvokeService(clt *core.Client, req *InvokeServiceRequest) (string, error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/servicemarket?access_token="
	var result struct {
		core.Error
		Data string `json:"data"`
	}
	req.ClientMsgId = util.NonceStr()
	if err := clt.PostJSON(incompleteURL, req, &result); err != nil {
		return "", err
	}
	if result.ErrCode != core.ErrCodeOK {
		return "", &result.Error
	}
	return result.Data, nil
}
