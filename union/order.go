package union

import "github.com/chanxuehong/wechat/util"

// Order 订单
type Order struct {
	// OrderID 订单ID
	OrderID string `json:"orderId,omitempty"`
	// PayTime 支付时间戳，单位为s
	PayTime int64 `json:"payTime,omitempty"`
	// ConfirmReceiveTime 确认收货时间戳，单位为s，没有时为0
	ConfirmReceiveTime int64 `json:"confirmReceiveTime,omitempty"`
	// ShopName 店铺名称
	ShopName string `json:"shopName,omitempty"`
	// ShopAppID 店铺 Appid
	ShopAppID string `json:"shopAppId,omitempty"`
	// ProductList 	商品列表
	ProductList []OrderProduct `json:"productList,omitempty"`
	// CustomizeInfo 自定义信息
	CustomizeInfo string `json:"customizeInfo,omitempty"`
	// CustomUserID 	自定义用户参数
	CustomUserID string `json:"customUserId,omitempty"`
	// UserNickName 	用户昵称
	UserNickName string `json:"userNickName,omitempty"`
}

type OrderProduct struct {
	// ProductID 商品SPU ID
	ProductID string `json:"productId,omitempty"`
	// SkuID sku ID
	SkuID string `json:"skuId,omitempty"`
	// Title 商品名称
	Title string `json:"title,omitempty"`
	// ThumbImg 商品缩略图 url
	ThumbImg string `json:"thumbImg,omitempty"`
	// Price 	商品成交总价，前带单位 ¥
	Price util.MoneyFloat `json:"price,omitempty"`
	// ProductCnt 成交数量
	ProductCnt int `json:"productCnt,omitempty"`
	// Ratio 	分佣比例，单位为万分之一
	Ratio int64 `json:"ratio,omitempty"`
	// CommissionStatus 分佣状态
	CommissionStatus CommissionStatus `json:"commissionStatus,omitempty"`
	// CommissionStatusUpdateTime 分佣状态更新时间戳，单位为s
	CommissionStatusUpdateTime util.Int64 `json:"commissionStatusUpdateTime,omitempty"`
	// ProfitShardingSucTime 结算时间，当分佣状态为已结算才有值，单位为s
	ProfitShardingSucTime util.Int64 `json:"profitShardingSucTime,omitempty"`
	// Commission 分佣金额，前带单位 ¥
	Commission util.MoneyFloat `json:"commission,omitempty"`
	// EstimatedCommission 预估分佣金额，单位为分
	EstimatedCommission int64 `json:"estimatedCommission,omitempty"`
	// CategoryStr 类目名称，多个用英文逗号分隔
	CategoryStr string `json:"categoryStr,omitempty"`
	// CustomizeInfo 自定义信息
	CustomizeInfo string `json:"customizeInfo,omitempty"`
	// PromotionInfo 	推广信息
	PromotionInfo *PromotionInfo `json:"promotionInfo,omitempty"`
}

// PromotionInfo 	推广信息
type PromotionInfo struct {
	// PromotionSourcePid 推广位 id
	PromotionSourcePid string `json:"promotionSourcePid,omitempty"`
	// PromotionSourceName 推广位名称
	PromotionSourceName string `json:"promotionSourceName,omitempty"`
}
