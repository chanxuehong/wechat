package union

// ShopInfo 商品所属小商店数据
type ShopInfo struct {
	// Name 小商店名称
	Name string `json:"name,omitempty"`
	// AppID 小商店AppID
	AppID string `json:"appId,omitempty"`
	// Username 小商店原始id
	Username string `json:"username,omitempty"`
	// HeadImgUrl 	小商店店铺头像
	HeadImgUrl string `json:"headImgUrl,omitempty"`
	// ShippingMethods 配送方式，结构见shippingMethods 的结构
	ShippingMethods *ShippingMethods `json:"shippingMethods,omitempty"`
	// AddressList 发货地，只有当配送方式包含「同城配送、上门自提」才出该项
	AddressList []Address `json:"addressList,omitempty"`
	// SameCityTemplate 配送范围，只有当配送方式包含「同城配送」才出该项
	SameCityTemplate *SameCityTemplate `json:"sameCityTemplate,omitempty"`
	// FreightTemplate 运费模板，只有当配送方式包含「快递」才出此项
	FreightTemplate *FreightTemplate `json:"freightTemplate,omitempty"`
}

// SameCityTemplate 配送范围，只有当配送方式包含「同城配送」才出该项
type SameCityTemplate struct {
	// DeliverScopeType 配送范围的定义方式，0：按照距离定义配送范围，1：按照区域定义配送范围
	DeliverScopeType int `json:"deliverScopeType,omitempty"`
	// Scope 配送范围
	Scope string `json:"scope,omitempty"`
	// Region 全城配送时的配送范围，结构同addressInfo
	Region *AddressInfo `json:"region,omitempty"`
}

// FreightTemplate 运费模板，只有当配送方式包含「快递」才出此项
type FreightTemplate struct {
	// NotSendArea 不发货地区
	NotSendArea struct {
		// AddressInfoList 不发货地区地址列表
		AddressInfoList []AddressInfo `json:"addressInfoList,omitempty"`
	} `json:"notSendArea"`
}
