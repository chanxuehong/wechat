package externalcontact

type TagGroup struct {
	GroupId    string `json:"group_id,omitempty"`    // 标签组id
	GroupName  string `json:"group_name,omitempty"`  // 标签组名称
	CreateTime int64  `json:"create_time,omitempty"` // 标签组创建时间
	Order      uint64 `json:"order,omitempty"`       // 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Deleted    bool   `json:"deleted,omitempty"`     // 标签组是否已经被删除，只在指定tag_id进行查询时返回
	Tags       []Tag  `json:"tag,omitempty"`         // 标签组内的标签列表
}

type Tag struct {
	Id         string `json:"id,omitempty"`          // 标签id
	Name       string `json:"name,omitempty"`        // 标签名称
	CreateTime int64  `json:"create_time,omitempty"` // 标签创建时间
	Order      uint64 `json:"order,omitempty"`       // 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Deleted    bool   `json:"deleted,omitempty"`     // 标签是否已经被删除，只在指定tag_id进行查询时返回
}
