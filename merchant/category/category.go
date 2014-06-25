package category

// 获取指定分类的所有子分类 成功时返回结果的数据结构
type Category struct {
	Id   string `json:"id"`   // 分类id
	Name string `json:"name"` // 分类name
}
