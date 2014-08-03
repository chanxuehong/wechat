// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package shelf

// 货架控件
type Module struct {
	EId int `json:"eid"` // 控件id, 标识控件 1,2,3,4,5

	GroupInfo  *groupInfo `json:"group_info,omitempty"` // 分组信息, 控件 1,3   的属性
	_GroupInfo groupInfo  `json:"-"`

	GroupInfos  *groupInfos `json:"group_infos,omitempty"` // 分组信息, 控件 2,4,5 的属性
	_GroupInfos groupInfos  `json:"-"`
}

// 初始化 md 指向的 Module 为 控件1
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule1(groupId int64, count int) {
	md.EId = 1
	md._GroupInfo.GroupId = groupId
	md._GroupInfo._Filter.Count = count

	md._GroupInfo.Filter = &md._GroupInfo._Filter
	md.GroupInfo = &md._GroupInfo

	// 容错
	md._GroupInfo.Image = ""
	md.GroupInfos = nil
}

// 初始化 md 指向的 Module 为 控件2
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule2(groupIds []int64) {
	groups := make([]Group, len(groupIds))
	for i := 0; i < len(groupIds); i++ {
		groups[i].GroupId = groupIds[i]
	}

	md.EId = 2
	md._GroupInfos.Groups = groups

	md.GroupInfos = &md._GroupInfos

	// 容错
	md._GroupInfos.ImageBackground = ""
	md.GroupInfo = nil
}

// 初始化 md 指向的 Module 为 控件3
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule3(groupId int64, image string) {
	md.EId = 3
	md._GroupInfo.GroupId = groupId
	md._GroupInfo.Image = image

	md.GroupInfo = &md._GroupInfo

	// 容错
	md._GroupInfo.Filter = nil
	md.GroupInfos = nil
}

// 初始化 md 指向的 Module 为 控件4
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule4(groups []Group) {
	md.EId = 4
	md._GroupInfos.Groups = groups

	md.GroupInfos = &md._GroupInfos

	// 容错
	md._GroupInfos.ImageBackground = ""
	md.GroupInfo = nil
}

// 初始化 md 指向的 Module 为 控件5
//  NOTE: 要求 md 指向的 Module 是 zero value, 即是刚创建的全0值, 否则有不可预料的错误!
func (md *Module) InitToModule5(groupIds []int64, imageBackground string) {
	groups := make([]Group, len(groupIds))
	for i := 0; i < len(groupIds); i++ {
		groups[i].GroupId = groupIds[i]
	}

	md.EId = 5
	md._GroupInfos.Groups = groups
	md._GroupInfos.ImageBackground = imageBackground

	md.GroupInfos = &md._GroupInfos

	// 容错
	md.GroupInfo = nil
}

type Group struct {
	GroupId int64 `json:"group_id"` // 分组ID, 控件 2,4,5 的属性

	// 分组照片(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，
	// 3个分组建议分辨率分别为: 350*350, 244*172, 244*172),
	// 控件 4 的属性
	Image string `json:"img,omitempty"`
}

// 控件 1,3 包含这个结构
type groupInfo struct {
	GroupId int64 `json:"group_id"` // 分组ID, 控件 1,3 的属性
	// 分组照片(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，建议分辨率600*208),
	// 控件 3 的属性
	Image   string           `json:"img,omitempty"`
	Filter  *groupInfoFilter `json:"filter,omitempty"` // 控件 1 的属性
	_Filter groupInfoFilter  `json:"-"`
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
