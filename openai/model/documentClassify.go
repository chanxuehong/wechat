package model

type DocumentClassify struct {
	L1 string `json:"level1_cls,omitempty"`
	L2 string `json:"level2_cls,omitempty"`
	L3 string `json:"level3_cls,omitempty"`
}
