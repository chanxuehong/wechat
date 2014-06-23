package product

// 获取指定分类的所有属性 成功时返回结果的数据结构
type Property struct {
	CategoryId int64  `json:"-"`
	Id         string `json:"id"`   // 属性id
	Name       string `json:"name"` // 属性name
	Values     []struct {
		Id   string `json:"id"`   // 属性值id
		Name string `json:"name"` // 属性值name
	} `json:"property_value"`
}
