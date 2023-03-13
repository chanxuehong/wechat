package express

import (
	"github.com/chanxuehong/wechat/component/core"
)

type PathGetRequest struct {
	// OpenID 用户openid
	OpenID string `json:"openid,omitempty"`
	// DeliveryID 快递公司ID，参见getAllDelivery
	DeliveryID string `json:"delivery_id,omitempty"`
	// WaybillID 运单ID
	WaybillID string `json:"waybill_id,omitempty"`
}

// PathGetResult
type PathGetResult struct {
	// OpenID 用户openid
	OpenID string `json:"openid,omitempty"`
	// DeliveryID 快递公司ID，参见getAllDelivery
	DeliveryID string `json:"delivery_id,omitempty"`
	// WaybillID 运单ID
	WaybillID string `json:"waybill_id,omitempty"`
	// PathItemNum 轨迹节点数量
	PathItemNum int `json:"path_item_num,omitempty"`
	// PathItemList 轨迹节点列表
	PathItemList []PathItem `json:"path_item_list,omitempty"`
}

// PathItemList 轨迹节点列表
type PathItem struct {
	// ActionTime 轨迹节点 Unix 时间戳
	ActionTime int64 `json:"action_time,omitempty"`
	// ActionType 轨迹节点类型
	ActionType int `json:"action_type,omitempty"`
	// ActionMsg 轨迹节点详情
	ActionMsg string `json:"action_msg,omitempty"`
}

// PathGet 查询运单轨迹
func PathGet(clt *core.Client, req *PathGetRequest) (res *PathGetResult, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/express/business/path/get?access_token="
	var result struct {
		core.Error
		PathGetResult
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	res = &result.PathGetResult
	return
}
