package union

import (
	"encoding/json"
	"strconv"

	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/util"
)

// ProductListRequest 查询全量商品 API Request
type ProductListRequest struct {
	// From 偏移，从0开始
	From int `json:"from,omitempty"`
	// Limit 每页条数
	Limit int `json:"limit,omitempty"`
	// Query 搜索关键词，具体含义与queryType相关
	Query string `json:"query,omitempty"`
	// QueryType 搜索类型
	QueryType int `json:"queryType,omitempty"`
	// MaxPrice 商品最高价格，单位分
	MaxPrice int64 `json:"maxPrice,omitempty"`
	// MinPrice 商品最低价格，单位分
	MinPrice int64 `json:"minPrice,omitempty"`
	// MinCommissionValue 佣金金额下限，单位分
	MinCommissionValue int64 `json:"minCommissionValue,omitempty"`
	// MinCommissionRatio 佣金比例下限，单位万分之一
	MinCommissionRatio int64 `json:"minCommissionRatio,omitempty"`
	// SortType
	SortType int `json:"sortType,omitempty"`
	// CategoryID 单个类目ID，值来自获取类目列表接口
	CategoryID uint64 `json:"categoryId,omitempty"`
	// ShopAppIDs 小商店AppID列表
	ShopAppIDs []string `json:"shopAppIds,omitempty"`
	// HasCoupon 是否有联盟券,1为含券商品，0为全部商品
	HasCoupon int `json:"hasCoupon,omitempty"`
	// Category 多个类目ID，多个用英文逗号分隔。若与 categoryid 同时存在，取categoryid
	Category string `json:"category,omitempty"`
	// NoCategory 黑名单类目ID，不拉出黑名单类目商品，多个用英文逗号分隔
	NoCategory string `json:"noCategory,omitempty"`
	// ProductID 商品SPUID，多个用英文逗号分隔
	ProductID string `json:"productId,omitempty"`
	// ShippingMethod 配送方式，JSON对象字符串，具体结构见shippingMethods 的结构
	ShippingMethods *ShippingMethods `json:"shippingMethods,omitempty"`
	// AddressList 发货地址列表，JSON数组字符串，单个结构见address
	AddressList []Address `json:"addressList,omitempty"`
}

// ShippingMethods 配送方式，JSON对象字符串，具体结构见shippingMethods 的结构
type ShippingMethods struct {
	// Express 是否支持快递，1：是，0：否
	Express int `json:"express,omitempty"`
	// SameCity 是否支持同城配送，1：是，0：否
	SameCity int `json:"sameCity,omitempty"`
	// Pickup 是否支持上门自提，1：是，0：否
	Pickup int `json:"pickup,omitempty"`
}

// Address 发货地址
type Address struct {
	// AddressInfo 地址信息
	AddressInfo *AddressInfo `json:"addressInfo,omitempty"`
	// AddressType 地址类型，结构同 shippingMethods
	AddressType *ShippingMethods `json:"addressType,omitempty"`
}

// AddressInfo 地址信息
type AddressInfo struct {
	// ProvinceName 国标收货地址第一级地址
	ProvinceName string `json:"provinceName,omitempty"`
	// CityName 国标收货地址第二级地址
	CityName string `json:"cityName,omitempty"`
	// CountryName 国标收货地址第三级地址
	CountryName string `json:"countryName,omitempty"`
}

type ProductListResult struct {
	// ProductList 商品列表数据
	ProductList []Product `json:"productList,omitempty"`
	// Total 商品总数
	Total int64 `json:"total,omitempty"`
}

func ProductList(clt *core.Client, req *ProductListRequest) (total int64, list []Product, err error) {
	values := util.GetUrlValues()
	values.Set("from", strconv.Itoa(req.From))
	values.Set("limit", strconv.Itoa(req.Limit))
	if req.Query != "" {
		values.Set("query", req.Query)
	}
	if req.QueryType > 0 {
		values.Set("queryType", strconv.Itoa(req.QueryType))
	}
	if req.MaxPrice > 0 {
		values.Set("maxPrice", strconv.FormatInt(req.MaxPrice, 10))
	}
	if req.MinPrice > 0 {
		values.Set("minPrice", strconv.FormatInt(req.MinPrice, 10))
	}
	if req.MinCommissionValue > 0 {
		values.Set("minCommissionValue", strconv.FormatInt(req.MinCommissionValue, 10))
	}
	if req.MinCommissionRatio > 0 {
		values.Set("minCommissionRatio", strconv.FormatInt(req.MinCommissionRatio, 10))
	}
	if req.SortType > 0 {
		values.Set("sortType", strconv.Itoa(req.SortType))
	}
	if req.CategoryID > 0 {
		values.Set("categoryId", strconv.FormatUint(req.CategoryID, 10))
	}
	if len(req.ShopAppIDs) > 0 {
		bs, _ := json.Marshal(req.ShopAppIDs)
		values.Set("shopAppIds", string(bs))
	}
	if req.HasCoupon == 1 {
		values.Set("hasCoupon", "1")
	}
	if req.Category != "" {
		values.Set("category", req.Category)
	}
	if req.NoCategory != "" {
		values.Set("noCategory", req.NoCategory)
	}
	if req.ProductID != "" {
		values.Set("productId", req.ProductID)
	}
	if req.ShippingMethods != nil {
		bs, _ := json.Marshal(req.ShippingMethods)
		values.Set("shippingMethods", string(bs))
	}
	if len(req.AddressList) > 0 {
		bs, _ := json.Marshal(req.AddressList)
		values.Set("addressList", string(bs))
	}
	query := values.Encode()
	util.PutUrlValues(values)
	incompleteURL := util.StringsJoin("https://api.weixin.qq.com/union/promoter/product/list?", query, "&access_token=")

	var result struct {
		core.Error
		ProductListResult
	}

	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	total = result.Total
	list = result.ProductList
	return
}
