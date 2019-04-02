package card

type AdvancedInfo struct {
	UseCondition    *UseCondition   `json:"use_condition,omitempty"`    //使用门槛（条件）字段，若不填写使用条件则在券面拼写 ：无最低消费限制，全场通用，不限品类；并在使用说明显示： 可与其他优惠共享
	Abstract        *Abstract       `json:"abstract,omitempty"`         //封面摘要结构体名称
	TextImageList   []TextImageList `json:"text_image_list,omitempty"`  //图文列表，显示在详情内页 ，优惠券券开发者须至少传入 一组图文列表
	TimeLimit       []TimeLimit     `json:"time_limit,omitempty"`       //使用时段限制，包含以下字段
	BusinessService []string        `json:"business_service,omitempty"` //商家服务类型： BIZ_SERVICE_DELIVER 外卖服务； BIZ_SERVICE_FREE_PARK 停车位； BIZ_SERVICE_WITH_PET 可带宠物； BIZ_SERVICE_FREE_WIFI 免费wifi， 可多选
}

type UseCondition struct {
	AcceptCategory          string `json:"accept_category,omitempty"` //指定可用的商品类目，仅用于代金券类型 ，填入后将在券面拼写适用于xxx
	RejectCategory          string `json:"reject_category,omitempty"` //指定不可用的商品类目，仅用于代金券类型 ，填入后将在券面拼写不适用于xxxx
	LeastCost               int    `json:"least_cost,omitempty"`      //满减门槛字段，可用于兑换券和代金券 ，填入后将在全面拼写消费满xx元可用。
	CanUseWithOtherDiscount bool   `json:"can_use_with_other_discount,omitempty"`
}

type Abstract struct {
	Abstract    string   `json:"abstract,omitempty"`      //封面摘要简介。
	IconURLList []string `json:"icon_url_list,omitempty"` //封面图片列表，仅支持填入一 个封面图片链接， 上传图片接口 上传获取图片获得链接，填写 非CDN链接会报错，并在此填入。 建议图片尺寸像素850*350
}

type TextImageList struct {
	ImageURL string `json:"image_url,omitempty"` //图片链接，必须调用 上传图片接口 上传图片获得链接，并在此填入， 否则报错
	Text     string `json:"text,omitempty"`      //图文描述
}

type TimeLimit struct {
	Type        string `json:"type,omitempty"`         //限制类型枚举值：支持填入 MONDAY 周一 TUESDAY 周二 WEDNESDAY 周三 THURSDAY 周四 FRIDAY 周五 SATURDAY 周六 SUNDAY 周日 此处只控制显示， 不控制实际使用逻辑，不填默认不显示
	BeginHour   int    `json:"begin_hour,omitempty"`   //当前type类型下的起始时间（小时） ，如当前结构体内填写了MONDAY， 此处填写了10，则此处表示周一 10:00可用
	EndHour     int    `json:"end_hour,omitempty"`     //当前type类型下的起始时间（分钟） ，如当前结构体内填写了MONDAY， begin_hour填写10，此处填写了59， 则此处表示周一 10:59可用
	BeginMinute int    `json:"begin_minute,omitempty"` //当前type类型下的结束时间（小时） ，如当前结构体内填写了MONDAY， 此处填写了20， 则此处表示周一 10:00-20:00可用
	EndMinute   int    `json:"end_minute,omitempty"`   //当前type类型下的结束时间（分钟） ，如当前结构体内填写了MONDAY， begin_hour填写10，此处填写了59， 则此处表示周一 10:59-00:59可用
}
