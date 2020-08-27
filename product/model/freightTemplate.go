package model

// 计费类型
type ValuationType = int

const (
	PIECE_VALUATION  ValuationType = 1 // 按件
	WEIGHT_VALUATION ValuationType = 2 // 按重量
)

type FreightTemplate struct {
	Id            uint64        `json:"template_id,omitempty"` // 模板ID
	Name          string        `json:"name,omitempty"`        // 模板名称
	ValuationType ValuationType `json:"valuation_type,omitempty"`
}
