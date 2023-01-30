package union

// CouponInfo 联盟优惠券数据
type CouponInfo struct {
	// HasCoupon 是否有联盟券，1为含券商品，0为全部商品
	HasCoupon int `json:"hasCoupon,omitempty"`
	// CouponID 券id
	CouponID string `json:"couponId,omitempty"`
	// CouponDetail 券详情
	CouponDetail *CouponDetail `json:"couponDetail,omitempty"`
}

// CouponDetail 券详情
type CouponDetail struct {
	// RestNum 券库存
	RestNum int64 `json:"restNum,omitempty"`
	// Type 券类型
	Type int `json:"type,omitempty"`
	// DiscountInfo 	券面额
	DiscountInfo *DiscountInfo `json:"discountInfo,omitempty"`
	// ValidInfo 有效期
	ValidInfo *CouponValidInfo `json:"validInfo,omitempty"`
	// ReceiveInfo 领券时间
	ReceiveInfo *CouponReceiveInfo `json:"receiveInfo,omitempty"`
}

// DiscountInfo 	券面额
type DiscountInfo struct {
	DiscountCondition *DiscountCondition `json:"discountCondition,omitempty"`
	// DiscountNum 	折扣数，如 5.1 折 为 5.1 * 1000
	DiscountNum int `json:"discountNum,omitempty"`
	// DiscountFee 直减金额，单位为分
	DiscountFee int64 `json:"discountFee,omitempty"`
}

type DiscountCondition struct {
	// ProductIDs 指定商品 id
	ProductIDs string `json:"productIds,omitempty"`
	// ProductCnt 商品数
	ProductCnt int `json:"productCnt,omitempty"`
	// ProductPrice 商品金额
	ProductPrice int64 `json:"productPrice,omitempty"`
}

// CouponValidInfo 有效期
type CouponValidInfo struct {
	// ValidType 有效期类型，1 为商品指定时间区间，2 为生效天数
	ValidType int `json:"validType,omitempty"`
	// ValidDayNum 生效天数
	ValidDayNum int `json:"validDayNum,omitempty"`
	// StartTime 有效开始时间
	StartTime string `json:"startTime,omitempty"`
	// EndTime 有效结束时间
	EndTime string `json:"endTime,omitempty"`
}

// CouponReceiveInfo 领券时间
type CouponReceiveInfo struct {
	// StartTime 领取开始时间戳
	StartTime string `json:"startTime,omitempty"`
	// EndTime 	领取结束时间戳
	EndTime string `json:"endTime,omitempty"`
	// LimitNumOnePerson 每人限领张数
	LimitNumOnePerson int `json:"limitNumOnePerson,omitempty"`
}
