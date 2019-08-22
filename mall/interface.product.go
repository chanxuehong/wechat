package mall

type Product struct {
	ItemCode             string     `json:"item_code"`                        // 物品ID（SPU ID），要求appid下全局唯一
	Title                string     `json:"title"`                            // 物品名称
	Desc                 string     `json:"desc,omitempty"`                   // 物品描述
	CategoryList         []string   `json:"category_list"`                    // 物品类目列表，用于搜索排序
	ImageList            []string   `json:"image_list"`                       // 物品图片链接列表，图片宽度必须大于750px，宽高比建议4:3 - 1:1之间
	AppPath              string     `json:"src_wxapp_path"`                   // 物品来源小程序路径
	SkuList              []Sku      `json:"sku_list,omitempty"`               // 物品SKU列表，单次导入不超过16个SKU，微信后台会合并多次导入的SKU
	SkuInfo              *Sku       `json:"sku_info,omitempty"`               // 物品SKU信息，微信后台会合并多次导入的SKU
	AttrList             []Attr     `json:"attr_list,omitempty"`              // 物品SPU属性
	Version              int        `json:"version,omitempty"`                // 非高并发更新数据的场景不建议填写此字段。数据版本号，需按照更新递增
	CanBeSearch          bool       `json:"can_be_search"`                    // 物品能否被搜索（默认true可以被搜索）
	BrandInfo            *Brand     `json:"brand_info,omitempty"`             // 商家信息
	PlatformCategoryList []Category `json:"platform_category_list,omitempty"` // 物品平台类目列表，填写的每个类目必须在好物圈物品类目表列出，多级类目只填最后一级（如完整类目为"运动户外-运动服饰-运动裤"，只需要填"运动裤"的类目ID与类目名）
	UpdateTime           int64      `json:"update_time,omitempty"`            // 加入购物车的时间，unix 秒级时间戳，不填默认为当前时间
}

type Attr struct {
	Name  string `json:"name"`  // 属性名称
	Value string `json:"value"` // 属性内容
}

type Poi struct {
	Lng          float64 `json:"longitude"`     // 门店的经度，WGS84标准
	Lat          float64 `json:"latitude"`      // 门店的纬度，WGS84标准
	Radius       float64 `json:"radius"`        // 门店可送达半径，单位km
	BusinessName string  `json:"business_name"` // 门店名称（仅为商户名，如：国美、麦当劳，不应包含地区、地址、分店名等信息，错误示例：北京国美），20个字符以内
	BranchName   string  `json:"branch_name"`   // 分店名称（不应包含地区信息，不应与门店名有重复，错误示例：北京王府井店），20个字符以内
	Address      string  `json:"address"`       // 门店地址（不包含省市区信息，如：新港中路123号）
}

type Sku struct {
	Id            string   `json:"sku_id"`                   // 物品sku_id，特殊情况下可以填入与item_code一致
	Price         int64    `json:"price"`                    // 物品价格，分为单位
	OriginalPrice int64    `json:"original_price,omitempty"` // 物品原价，分为单位
	Status        uint     `json:"status"`                   // 物品状态，1：在售，2：停售，3：售罄
	AttrList      []Attr   `json:"sku_attr_list,omitempty"`  // sku属性列表，参考attr_list
	PoiList       []Poi    `json:"poi_list,omitempty"`       // 物品所在门店的列表。如果物品仅支持到店提货或到家送货，该字段必填；如果物品同时支持线上物流配送，该字段应为空。
	Version       int      `json:"version,omitempty"`        // 非高并发更新数据的场景不建议填写此字段。数据版本号，需按照更新递增
	BarCodeInfo   *BarCode `json:"bar_code_info,omitempty"`  // 物品的条形码信息
}

type Category struct {
	Id   uint   `json:"category_id"`   // 平台类目ID
	Name string `json:"category_name"` // 平台类目名称
}

type Brand struct {
	Logo  string      `json:"logo,omitempty"`                // 商家logo，不填的话，默认取小程序头像
	Name  string      `json:"name,omitempty"`                // 商家名称，不填的话，默认取小程序名字
	Phone string      `json:"phone,omitempty"`               // 用于售后场景的商家联系电话，便于用户咨询和问题解决
	Page  *DetailPage `json:"contact_detail_page,omitempty"` // 联系商家页面
}

type BarCode struct {
	Type string `json:"barcode_type"` // 条形码类型, 目前支持"ean8", "ean13", 前者对应的barcode为8位纯数字字符串，后者为13位纯数字字符串
	Code string `json:"barcode"`      // 条形码数字字符串
}
