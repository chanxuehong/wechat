package order

const (
	// 订单状态(2-待发货, 3-已发货, 5-已完成, 8-维权中)
	ORDER_STATUS_DELIVERING       = 2
	ORDER_STATUS_DELIVERED        = 3
	ORDER_STATUS_DONE             = 5
	ORDER_STATUS_RIGHTS_DEFENDING = 8
)

const (
	// 物流公司ID
	DELIVERY_COMPANY_ID_EMS        = "Fsearch_code"  // 邮政EMS
	DELIVERY_COMPANY_ID_SHENTONG   = "002shentong"   // 申通快递
	DELIVERY_COMPANY_ID_ZHONGTONG  = "066zhongtong"  // 中通速递
	DELIVERY_COMPANY_ID_YUANTONG   = "056yuantong"   // 圆通速递
	DELIVERY_COMPANY_ID_TIANTIAN   = "042tiantian"   // 天天快递
	DELIVERY_COMPANY_ID_SHUNFENG   = "003shunfeng"   // 顺丰速运
	DELIVERY_COMPANY_ID_YUNDA      = "059Yunda"      // 韵达快运
	DELIVERY_COMPANY_ID_ZHAIJISONG = "064zhaijisong" // 宅急送
	DELIVERY_COMPANY_ID_HUITONG    = "020huitong"    // 汇通快运
	DELIVERY_COMPANY_ID_YIXUN      = "zj001yixun"    // 易迅快递
)
