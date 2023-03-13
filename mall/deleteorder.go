package mall

import (
	"github.com/bububa/wechat/mp/core"
)

type DeleteOrderRequest struct {
	OrderId string `json:"order_id"`
	OpenId  string `json:"user_open_id"`
}

// 删除订单
// 用户可以对订单进行删除
func DeleteOrder(clt *core.Client, req *DeleteOrderRequest) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/mall/deleteorder?access_token="
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
