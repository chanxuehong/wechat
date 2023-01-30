package union

// Product 商品
type Product struct {
	// ProductID 商品SPU ID
	ProductID string `json:"productId,omitempty"`
	// AppID 商品所在小商店的AppID
	AppID string `json:"appId,omitempty"`
	// Product 商品具体信息
	Product *ProductInfo `json:"product,omitempty"`
	// LeagueExInfo 联盟佣金相关数据
	LeagueExInfo *LeagueExInfo `json:"leagueExInfo,omitempty"`
	// ShopInfo 商品所属小商店数据
	ShopInfo *ShopInfo `json:"shopInfo,omitempty"`
	// CouponInfo 联盟优惠券数据
	CouponInfo *CouponInfo `json:"couponInfo,omitempty"`
}

// LeagueExInfo 联盟佣金相关数据
type LeagueExInfo struct {
	// HasCommission 是否有佣金，1/0
	HasCommission int `json:"hasCommission,omitempty"`
	// CommissionRatio 佣金比例，万分之一
	CommissionRatio int64 `json:"commissionRatio,omitempty"`
	// CommissionValue 	佣金金额，单位分
	CommissionValue int64 `json:"commissionValue,omitempty"`
}

// ProductInfo 商品具体信息
type ProductInfo struct {
	// Title 商品标题
	Title string `json:"title,omitempty"`
	// SubTitle 商品子标题
	SubTitle string `json:"subTitle,omitempty"`
	// HeadImg 商品主图
	HeadImg []string `json:"headImg,omitempty"`
	// Category 商品类目
	Category []string `json:"category,omitempty"`
	// ShopName 商店名称
	ShopName string `json:"shopName,omitempty"`
	// Brand 品牌名称
	Brand string `json:"brand,omitempty"`
	// BrandID 品牌ID
	BrandID string `json:"brandId,omitempty"`
	// Model 	型号
	Model string `json:"model,omitempty"`
	// Detail 商品详细数据
	Detail *ProductDetail `json:"detail,omitempty"`
	// MinPrice 商品最低价格，单位分
	MinPrice int64 `json:"minPrice,omitempty"`
	// Discount 商品优惠金额，单位分
	Discount int64 `json:"discount,omitempty"`
	// DiscountPrice 	商品券后价
	DiscountPrice int `json:"discountPrice,omitempty"`
	// TotalStockNum 	总库存
	TotalStock int64 `json:"totalStock,omitempty"`
	// TotalSoldNum 累计销量
	TotalSoldNum int64 `json:"totalSoldNum,omitempty"`
	// TotalOrderNum 	累计订单量
	TotalOrderNum int64 `json:"totalOrderNum,omitempty"`
	// Skus 商品SKU
	Skus []Sku `json:"skus,omitempty"`
	// PluginResult 是否引用小商店组件（未引用组件的商品不可推广），0：否，1：是
	PluginResult int `json:"pluginResult,omitempty"`
}

// ProductDetail 商品详细数据
type ProductDetail struct {
	// DetailImg 商品详情图片
	DetailImg []string `json:"detailImg,omitempty"`
	// Param 商品参数
	Param []interface{} `json:"param,omitempty"`
}

// ShareProduct 商品
type ShareProduct struct {
	// ProductID 商品SPU ID
	ProductID string `json:"productId,omitempty"`
	// AppID 商品所在小商店的AppID
	AppID string `json:"appId,omitempty"`
	// CustomizeInfo 自定义参数，最多包含80个字符
	CustomizeInfo string `json:"customizeInfo,omitempty"`
	// ProductInfo 商品具体信息
	ProductInfo *ProductInfo `json:"productInfo,omitempty"`
	// ShareInfo 推广相关信息
	ShareInfo *ShareInfo `json:"shareInfo,omitempty"`
}

// ShareInfo 推广相关信息
type ShareInfo struct {
	// AppID 推广商品的小程序AppID
	AppID string `json:"appId,omitempty"`
	// Path 推广商品的小程序Path
	Path string `json:"path,omitempty"`
	// CouponPath 推广商品的带券小程序Path
	CouponPath string `json:"couponPath,omitempty"`
	// PromotionURL 推广商品短链
	PromotionURL string `json:"promotionUrl,omitempty"`
	// CouponPromotionURL 推广商品带券短链
	CouponPromotionURL string `json:"couponPromotionUrl,omitempty"`
	// PromotionWording 推广商品文案
	PromotionWording string `json:"promotionWording,omitempty"`
	// CouponPromotionWording 推广商品带券文案
	CouponPromotionWording string `json:"couponPromotionWording,omitempty"`
	// PromotionTag 推广商品tag
	PromotionTag string `json:"promotionTag,omitempty"`
	// CouponPromotionTag 推广商品带券tag
	CouponPromotionTag string `json:"couponPromotionTag,omitempty"`
}
