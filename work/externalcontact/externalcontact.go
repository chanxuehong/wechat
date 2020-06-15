package externalcontact

type ExternalContact struct {
	ExternalUserId  string           `json:"external_userid,omitempty"`  // 外部联系人的userid
	Name            string           `json:"name,omitempty"`             // 外部联系人的名称*
	Avatar          string           `json:"avatar,omitempty"`           // 外部联系人头像，第三方不可获取
	Type            uint             `json:"type,omitempty"`             // 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
	Gender          uint             `json:"gender,omitempty"`           // 外部联系人性别 0-未知 1-男性 2-女性
	UnionId         string           `json:"unionid,omitempty"`          // 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段。
	Position        string           `json:"position,omitempty"`         // 外部联系人的职位，如果外部企业或用户选择隐藏职位，则不返回，仅当联系人类型是企业微信用户时有此字段
	CorpName        string           `json:"corp_name,omitempty"`        // 外部联系人所在企业的简称，仅当联系人类型是企业微信用户时有此字段
	CorpFullName    string           `json:"corp_full_name,omitempty"`   // 外部联系人所在企业的主体名称，仅当联系人类型是企业微信用户时有此字段
	ExternalProfile *ExternalProfile `json:"external_profile,omitempty"` // 外部联系人的自定义展示信息，可以有多个字段和多种类型，包括文本，网页和小程序，仅当联系人类型是企业微信用户时有此字段，字段详情见对外属性；
}

type ExternalProfile struct {
	CorpName string         `json:"external_corp_name,omitempty"` // 企业对外简称，需从已认证的企业简称中选填。可在“我的企业”页中查看企业简称认证状态。
	Attrs    []ExternalAttr `json:"external_attr,omitempty"`      // 属性列表，目前支持文本、网页、小程序三种类型
}

type ExternalAttr struct {
	Type        uint                     `json:"type,omitempty"`        // 属性类型: 0-文本 1-网页 2-小程序
	Name        string                   `json:"name,omitempty"`        // 属性名称： 需要先确保在管理端有创建该属性，否则会忽略
	Text        *ExternalAttrText        `json:"text,omitempty"`        // 文本类型的属性
	Web         *ExternalAttrWeb         `json:"web,omitempty"`         // 网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空
	MiniProgram *ExternalAttrMiniProgram `json:"miniprogram,omitempty"` // 小程序类型的属性，appid和title字段要么同时为空表示清除改属性，要么同时不为空
}

type ExternalAttrText struct {
	Value string `json:"value,omitempty"` // 文本属性内容,长度限制12个UTF8字符
}

type ExternalAttrWeb struct {
	Url   string `json:"url,omitempty"`   // 网页的url,必须包含http或者https头
	Title string `json:"title,omitempty"` // 网页的展示标题,长度限制12个UTF8字符
}

type ExternalAttrMiniProgram struct {
	AppId string `json:"appid,omitempty"`    // 小程序appid，必须是有在本企业安装授权的小程序，否则会被忽略
	Title string `json:"title,omitempty"`    // 小程序的展示标题,长度限制12个UTF8字符
	Path  string `json:"pagepath,omitempty"` // 小程序的页面路径
}

type FollowUser struct {
	UserId         string          `json:"userid,omitempty"`           // 添加了此外部联系人的企业成员userid
	Remark         string          `json:"remark,omitempty"`           // 该成员对此外部联系人的备注
	Description    string          `json:"description,omitempty"`      // 该成员对此外部联系人的描述
	CreateTime     int64           `json:"createtime,omitempty"`       // 该成员添加此外部联系人的时间
	Tags           []FollowUserTag `json:"tags,omitempty"`             // 该成员添加此外部联系人所打标签
	RemarkCorpName string          `json:"remark_corp_name,omitempty"` // 该成员对此客户备注的企业名称
	RemarkMobiles  []string        `json:"remark_mobiles,omitempty"`   // 该成员对此客户备注的手机号码，第三方不可获取
	State          string          `json:"state,omitempty"`            // 该成员添加此客户的渠道，由用户通过创建「联系我」方式指定
}

type FollowUserTag struct {
	GroupName string `json:"group_name,omitepmty"` // 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	Name      string `json:"tag_name,omitempty"`   // 该成员添加此外部联系人所打标签名称
	Type      uint   `json:"type,omitempty"`       // 该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义
}
