package model

type Attribute struct {
	Key   string `json:"attr_key,omitempty"`   // 属性键key（属性自定义用）
	Value string `json:"attr_value,omitempty"` // 属性值（属性自定义用）
}
