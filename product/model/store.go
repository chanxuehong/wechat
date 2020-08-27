package model

type Store struct {
	Name string `json:"store_name,omitempty"` // 商店名称
	Logo string `json:"store_logo,omitempty"` // 商店头像
}
