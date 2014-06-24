package category

// 获取指定子分类的所有SKU 成功时返回结果的数据结构
type SKU struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Values []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"value_list"`
}
