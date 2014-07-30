// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package shelf

// with shelf_data
type Shelf struct {
	Id   int64  `json:"shelf_id,omitempty"`
	Name string `json:"shelf_name"`

	// 货架招牌图片URL(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，
	// 建议尺寸为640*120，仅控件1-4有banner，控件5没有banner)
	Banner string `json:"shelf_banner,omitempty"`

	Info struct {
		ModuleInfos []Module `json:"module_infos,omitempty"`
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
		ModuleInfos []Module `json:"module_infos,omitempty"`
	} `json:"shelf_info"`
}

// 货架控件
type Module struct {
	EId int `json:"eid"` // 控件id, 标识控件 1,2,3,4,5

	GroupInfo  *groupInfo  `json:"group_info,omitempty"`  // 分组信息, 控件 1,3   的属性
	GroupInfos *groupInfos `json:"group_infos,omitempty"` // 分组信息, 控件 2,4,5 的属性
}

// 初始化 md 指向的 Module 为 控件1
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule1(groupId int64, count int) {
	md.EId = 1
	md.GroupInfo = &groupInfo{
		GroupId: groupId,
		Filter: &groupInfoFilter{
			Count: count,
		},
	}
}

// 初始化 md 指向的 Module 为 控件2
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule2(groupIds []int64) {
	groups := make([]Group, len(groupIds))
	for i := 0; i < len(groupIds); i++ {
		groups[i].GroupId = groupIds[i]
	}

	md.EId = 2
	md.GroupInfos = &groupInfos{
		Groups: groups,
	}
}

// 初始化 md 指向的 Module 为 控件3
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule3(groupId int64, image string) {
	md.EId = 3
	md.GroupInfo = &groupInfo{
		GroupId: groupId,
		Image:   image,
	}
}

// 初始化 md 指向的 Module 为 控件4
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule4(groups []Group) {
	md.EId = 4
	md.GroupInfos = &groupInfos{
		Groups: groups,
	}
}

// 初始化 md 指向的 Module 为 控件5
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule5(groupIds []int64, imageBackground string) {
	groups := make([]Group, len(groupIds))
	for i := 0; i < len(groupIds); i++ {
		groups[i].GroupId = groupIds[i]
	}

	md.EId = 5
	md.GroupInfos = &groupInfos{
		Groups:          groups,
		ImageBackground: imageBackground,
	}
}

// 控件 1,3 包含这个结构
type groupInfo struct {
	GroupId int64 `json:"group_id"` // 分组ID, 控件 1,3 的属性
	// 分组照片(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，建议分辨率600*208),
	// 控件 3 的属性
	Image  string           `json:"img,omitempty"`
	Filter *groupInfoFilter `json:"filter,omitempty"` // 控件 1 的属性
}

type groupInfoFilter struct {
	Count int `json:"count"` // 该控件展示商品个数, 控件 1 的属性
}

// 控件 2,4,5 包含这个结构
type groupInfos struct {
	// 分组列表, 控件 2,4,5 的属性
	Groups []Group `json:"groups,omitempty"`

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
