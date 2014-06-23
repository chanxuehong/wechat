package shelf

// with shelf_data
type Shelf struct {
	Id   int64  `json:"shelf_id,omitempty"`
	Name string `json:"shelf_name"`

	// 货架招牌图片URL(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，
	// 建议尺寸为640*120，仅控件1-4有banner，控件5没有banner)
	Banner string `json:"shelf_banner,omitempty"`

	Info struct {
		ModuleInfos []*Module `json:"module_infos,omitempty"`
	} `json:"shelf_data"`
}

// with shelf_info
type ShelfX struct {
	Id   int64  `json:"shelf_id,omitempty"`
	Name string `json:"shelf_name"`

	// 货架招牌图片URL(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，
	// 建议尺寸为640*120，仅控件1-4有banner，控件5没有banner)
	Banner string `json:"shelf_banner,omitempty"`

	Info struct {
		ModuleInfos []*Module `json:"module_infos,omitempty"`
	} `json:"shelf_info"`
}

// 货架控件
type Module struct {
	EId int `json:"eid"` // 控件id, 标识控件 1,2,3,4,5

	GroupInfo  *GroupInfo  `json:"group_info,omitempty"`  // 分组信息, 控件 1,3   的属性
	GroupInfos *GroupInfos `json:"group_infos,omitempty"` // 分组信息, 控件 2,4,5 的属性
}

func NewModule1(groupId int64, count int) *Module {
	return &Module{
		EId: 1,
		GroupInfo: &GroupInfo{
			GroupId: groupId,
			Filter: &GroupInfoFilter{
				Count: count,
			},
		},
	}
}
func NewModule2(groupIds []int64) *Module {
	groups := make([]Group, len(groupIds))
	for i := 0; i < len(groupIds); i++ {
		groups[i].GroupId = groupIds[i]
	}

	return &Module{
		EId: 2,
		GroupInfos: &GroupInfos{
			Groups: groups,
		},
	}
}
func NewModule3(groupId int64, image string) *Module {
	return &Module{
		EId: 3,
		GroupInfo: &GroupInfo{
			GroupId: groupId,
			Image:   image,
		},
	}
}
func NewModule4(groups []Group) *Module {
	return &Module{
		EId: 4,
		GroupInfos: &GroupInfos{
			Groups: groups,
		},
	}
}
func NewModule5(groupIds []int64, imageBackground string) *Module {
	groups := make([]Group, len(groupIds))
	for i := 0; i < len(groupIds); i++ {
		groups[i].GroupId = groupIds[i]
	}

	return &Module{
		EId: 5,
		GroupInfos: &GroupInfos{
			Groups:          groups,
			ImageBackground: imageBackground,
		},
	}
}

// 控件 1,3 包含这个结构
type GroupInfo struct {
	GroupId int64 `json:"group_id"` // 分组ID, 控件 1,3 的属性
	// 分组照片(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，建议分辨率600*208),
	// 控件 3 的属性
	Image  string           `json:"img,omitempty"`
	Filter *GroupInfoFilter `json:"filter,omitempty"` // 控件 1 的属性
}

// 控件1 的 Filter
type GroupInfoFilter struct {
	Count int `json:"count"` // 该控件展示商品个数, 控件 1 的属性
}

// 控件 2,4,5 包含这个结构
type GroupInfos struct {
	// 分组列表, 控件 2,4,5 的属性
	Groups []Group `json:"groups"`

	// 分组照片(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，建议分辨率640*1008),
	// 控件 5 的属性
	ImageBackground string `json:"img_background,omitempty"`
}

type Group struct {
	GroupId int64 `json:"group_id"` // 分组ID, 控件 2,4,5 的属性

	// 分组照片(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，
	// 3个分组建议分辨率分别为: 350*350, 244*172, 244*172),
	// 控件 4 的属性
	Image string `json:"img,omitempty"`
}
