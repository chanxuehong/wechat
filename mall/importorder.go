package mall

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
)

type ImportOrderRequest struct {
	OrderList []Order `json:"order_list"`
}

// 导入订单
// 开发者可以在用户支付完成后，同步小程序/H5/APP订单数据至好物圈（H5/APP订单需保证调用支付接口的H5/APP与导入数据的小程序绑定在同一个微信开放平台帐号下）。
// 历史订单导入：在首次接入时，开发者需导入最近三个月的”历史订单“数据，导入数据有助于物品在搜索中获得更好的曝光，帮助新接入的商家在搜索中实现冷启动。
func AddOrder(clt *core.Client, req *ImportOrderRequest, isHistory uint) (err error) {
	incompleteURL := fmt.Sprintf("https://api.weixin.qq.com/mall/importorder?action=add-order&is_history=%d&access_token=", isHistory)
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

// 更新订单信息
// 开发者对订单信息进行更新，如订单状态改变等
func UpdateOrder(clt *core.Client, req *ImportOrderRequest, isHistory uint) (err error) {
	incompleteURL := fmt.Sprintf("https://api.weixin.qq.com/mall/importorder?action=update-order&is_history=%d&access_token=", isHistory)
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
