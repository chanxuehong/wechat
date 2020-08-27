package model

type Category struct {
	Id    uint64 `json:"cat_id,omitempty"`   // 类目ID
	Pid   uint64 `json:"f_cat_id,omitempty"` // 类目父ID
	Name  string `json:"name,omitempty"`     // 类目名称
	Level int    `json:"level,omitempty"`
}
