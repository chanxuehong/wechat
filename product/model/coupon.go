package model

type Coupon struct {
	Id         uint64       `json:"coupon_id,omitempty"`   // 优惠券
	Type       CouponType   `json:"type,omitempty"`        // 优惠券类型
	Status     CouponStatus `json:"status,omitempty"`      // 优惠券状态
	CreateTime string       `json:"create_time,omitempty"` // 优惠券创建时间
	UpdateTime string       `json:"update_time,omitempty"` // 优惠券更新时间
	Info       *CouponInfo  `json:"coupon_info,omitempty"`
	Stock      *CouponStock `json:"stock_info,omitempty"`
}

type CouponInfo struct {
	Name  string       `json:"name,omitempty"` // 优惠券名称
	Valid *CouponValid `json:"valid_info,omitempty"`
}

type CouponValid struct {
	Type      CouponValidType `json:"valid_type,omitempty"`    // 优惠券有效期类型
	DayNum    int             `json:"valid_day_num,omitempty"` // 优惠券有效天数，valid_type=2时才有意义
	StartTime string          `json:"start_time,omitempty"`    // 优惠券有效期开始时间，valid_type=1时才有意义
	EndTime   string          `json:"end_time,omitempty"`      // 优惠券有效期结束时间，valid_type=1时才有意义
}

type CouponStock struct {
	IssuedNum  int `json:"issued_num,omitempty"`  // 优惠券发放量
	ReceiveNum int `json:"receive_num,omitempty"` // 优惠券领用量
	UsedNum    int `json:"used_num,omitempty"`    // 优惠券已用量
}

// 优惠券类型
type CouponType = int

const (
	PRODUCT_CONDITION_DISCOUNT_COUPON CouponType = 1   // 商品条件折扣券
	PRODUCT_REACH_DISCOUNT_COUPON     CouponType = 2   // 商品满减券
	PRODUCT_GENERAL_DISCOUNT_COUPON   CouponType = 3   // 商品统一折扣券
	PRODUCT_PURCHASE_DISCOUNT_COUPON  CouponType = 4   // 商品直减券
	STORE_CONDITION_DISCOUNT_COUPON   CouponType = 101 // 店铺条件折扣券
	STORE_REACH_DISCOUNT_COUPON       CouponType = 102 // 店铺满减券
	STORE_GENERAL_DISCOUNT_COUPON     CouponType = 103 // 店铺统一折扣券
	STORE_PURCHASE_DISCOUNT_COUPON    CouponType = 104 // 店铺直减券
)

// 优惠券状态
type CouponStatus = int

const (
	EDITING_COUPON      CouponStatus = 1   // 未生效，编辑中
	READY_COUPON        CouponStatus = 2   // 生效
	EXPIRED_COUPON      CouponStatus = 3   // 已过期
	CANCELED_COUPON     CouponStatus = 4   // 已作废
	DELETED_COUPON      CouponStatus = 5   // 删除
	READY_TO_USE_COUPON CouponStatus = 101 // 生效中
	EXPIRED2_COUPON     CouponStatus = 102 // 已过期
	USED_COUPON         CouponStatus = 103 // 已经使用
)

// 优惠券生效类型

type CouponValidType = int

const (
	DATE_RANGE_VALID_COUPON CouponValidType = 1 // 商品指定时间区间
	DAY_NUM_VALID_COUPON    CouponValidType = 2 // 生效天数
)
