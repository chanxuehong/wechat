package express

import (
	"github.com/bububa/wechat/component/core"
)

// OrderBatchGet 批量获取运单数据
func OrderBatchGet(clt *core.Client, req []OrderGetRequest) (list []Order, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/express/business/order/batchget?access_token="
	var result struct {
		core.Error
		// List 批量获取运单数据
		List []Order `json:"order_list,omitempty"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.List
	return
}
