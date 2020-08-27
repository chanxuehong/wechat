package model

type Brand struct {
	Cid1 uint64     `json:"first_cat_id,omitempty"`  // 第一级类目
	Cid2 uint64     `json:"second_cat_id,omitempty"` // 第二级类目
	Cid3 uint64     `json:"third_cat_id,omitempty"`  // 第三级类目
	Info *BrandInfo `json:"brand_info,omitempty"`
}

type BrandInfo struct {
	Id   uint64 `json:"brand_id,omitempty"`   // 品牌ID
	Name string `json:"brand_name,omitempty"` // 品牌名称
}
