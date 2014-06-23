package express

const (
	// 快递ID列表
	EXPRESS_ID_PINGYOU = 10000027 // 平邮
	EXPRESS_ID_KUAIDI  = 10000028 // 快递
	EXPRESS_ID_EMS     = 10000029 // EMS
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

const (
	// 支付方式(0-买家承担运费, 1-卖家承担运费)
	ASSUMER_BUYER  = 0
	ASSUMER_SELLER = 1
)

const (
	// 计费单位(0-按件计费, 1-按重量计费, 2-按体积计费，目前只支持按件计费，默认为0)
	VALUATION_BY_ITEM   = 0
	VALUATION_BY_WEIGHT = 1
	VALUATION_BY_VOLUME = 2
)
