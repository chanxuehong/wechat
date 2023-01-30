package union

import (
	"strconv"

	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/util"
)

// CommissionStatus 分佣状态
type CommissionStatus string

const (
	// SETTLEMENT_PENDING 待结算
	SETTLEMENT_PENDING CommissionStatus = "SETTLEMENT_PENDING"
	// SETTLEMENT_SUCCESS 已结算
	SETTLEMENT_SUCCESS CommissionStatus = "SETTLEMENT_SUCCESS"
	// SETTLEMENT_CANCELED 取消结算
	SETTLEMENT_CANCELED CommissionStatus = "SETTLEMENT_CANCELED"
)

// OrderSearchRequest 根据订单支付时间、订单分佣状态拉取订单详情
type OrderSearchRequest struct {
	// Page 页码，起始为 1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小，最大 200
	PageSize int `json:"pageSize,omitempty"`
	// StartTimestamp 起始时间戳，单位为秒
	StartTimestamp int64 `json:"startTimestamp,omitempty"`
	// EndTimestamp 结束时间戳，单位为秒
	EndTimestamp int64 `json:"endTimestamp,omitempty"`
	// CommissionStatus 分佣状态
	CommissionStatus string `json:"commissionStatus,omitempty"`
	// SortByCommissionUpdateTime 是否按照分佣状态更新时间排序和筛选订单，1：是，0：否
	SortByCommissionUpdateTime int `json:"sortByCommissionUpdateTime,omitempty"`
	// StartCommissionUpdateTime 分佣状态更新时间起始时间戳，单位为秒
	StartCommissionUpdateTime int64 `json:"startCommissionUpdateTime,omitempty"`
	// EndCommissionUpdateTime 分佣状态更新时间结束时间戳，单位为秒
	EndCommissionUpdateTime int64 `json:"endCommissionUpdateTime,omitempty"`
}

// OrderSearchResult
type OrderSearchResult struct {
	// PageSize 分页大小
	PageSize int `json:"pageSize,omitempty"`
	// TotalNum 订单总数
	TotalNum int `json:"totalNum,omitempty"`
	// OrderList 订单列表
	OrderList []Order `json:"orderList,omitempty"`
}

// OrderSearch 获取订单列表
func OrderSearch(clt *core.Client, req *OrderSearchRequest) (total int, orders []Order, err error) {
	values := util.GetUrlValues()
	values.Set("page", strconv.Itoa(req.Page))
	values.Set("pageSize", strconv.Itoa(req.PageSize))
	if req.StartTimestamp > 0 {
		values.Set("startTimestamp", strconv.FormatInt(req.StartTimestamp, 10))
	}
	if req.EndTimestamp > 0 {
		values.Set("endTimestamp", strconv.FormatInt(req.EndTimestamp, 10))
	}
	if req.CommissionStatus != "" {
		values.Set("commissionStatus", req.CommissionStatus)
	}
	if req.SortByCommissionUpdateTime == 1 {
		values.Set("sortByCommissionUpdateTime", "1")
	}
	if req.StartCommissionUpdateTime > 0 {
		values.Set("startCommissionUpdateTime", strconv.FormatInt(req.StartCommissionUpdateTime, 10))
	}
	if req.EndCommissionUpdateTime > 0 {
		values.Set("endCommissionUpdateTime", strconv.FormatInt(req.EndCommissionUpdateTime, 10))
	}
	query := values.Encode()
	util.PutUrlValues(values)
	incompleteURL := util.StringsJoin("https://api.weixin.qq.com/union/promoter/order/search?", query, "&access_token=")

	var result struct {
		core.Error
		OrderSearchResult
	}

	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	total = result.TotalNum
	orders = result.OrderList
	return
}
