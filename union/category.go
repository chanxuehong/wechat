package union

// Category 类目ID
type Category struct {
	// CatID 类目ID
	CatID string `json:"catId,omitempty"`
	// Name 类目名称
	Name string `json:"name,omitempty"`
	// FCatID fCatId
	FCatID string `json:"fCatId,omitempty"`
	// CatType
	CatType int `json:"catType,omitempty"`
	// Bizuin
	Bizuin string `json:"bizuin,omitempty"`
	// BrandCat
	BrandCat int `json:"brandCat,omitempty"`
	// CatInfo
	CatInfo *CategoryInfo `json:"catInfo,omitempty"`
}

type CategoryInfo struct {
	IsRequired  int `json:"isRequired,omitempty"`
	IsCustomize int `json:"isCustomize,omitempty"`
	Level       int `json:"level,omitempty"`
}
