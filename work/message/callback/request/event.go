package request

import (
	"github.com/chanxuehong/wechat/work/message"
)

const (
	// 普通事件类型
	EventTypeSubscribe             message.EventType = "subscribe"               // 关注事件, 包括点击关注和扫描二维码(公众号二维码和公众号带参数二维码)关注
	EventTypeUnsubscribe           message.EventType = "unsubscribe"             // 取消关注事件
	EventTypeEnterAgent            message.EventType = "enter_agent"             // 本事件在成员进入企业微信的应用时触发
	EventTypeLocation              message.EventType = "LOCATION"                // 上报地理位置事件
	EventTypeBatchJobResult        message.EventType = "batch_job_result"        // 本事件是成员在使用异步任务接口时，用于接收任务执行完毕的结果通知。
	EventTypeChangeContact         message.EventType = "change_contact"          // 新增/更新/删除/部门事件;标签成员变更事件
	EventTypeClick                 message.EventType = "click"                   // 点击菜单拉取消息的事件推送
	EventTypeView                  message.EventType = "view"                    // 点击菜单跳转链接的事件推送
	EventTypeScanCodePush          message.EventType = "scancode_push"           // 扫码推事件的事件推送
	EventTypeScanCodeWaitMsg       message.EventType = "scancode_waitmsg"        // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventTypePicSysPhoto           message.EventType = "pic_sysphoto"            // 弹出系统拍照发图的事件推送
	EventTypePicPhotoOrAlbum       message.EventType = "pic_photo_or_album"      // 弹出拍照或者相册发图的事件推送
	EventTypePicWeixin             message.EventType = "pic_weixin"              // 弹出微信相册发图器的事件推送
	EventTypeLocationSelect        message.EventType = "location_select"         // 弹出地理位置选择器的事件推送
	EventTypeOpenApprovalChange    message.EventType = "open_approval_change"    // 审批状态通知事件
	EventTypeTaskCardClick         message.EventType = "taskcard_click"          // 任务卡片事件推送
	EventTypeChangeExternalContact message.EventType = "change_external_contact" // 企业客户事件
	EventTypeChangeExternalChat    message.EventType = "change_external_chat"    // 客户群变更事件
)

// 关注事件
type SubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType message.EventType `xml:"Event" json:"Event"`                           // subscribe
	EventKey  message.CDATA     `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 此事件该值为空
}

func GetSubscribeEvent(msg *message.MixedMsg) *SubscribeEvent {
	return &SubscribeEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
	}
}

// 取消关注事件
type UnsubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType message.EventType `xml:"Event"              json:"Event"`              // unsubscribe
	EventKey  message.CDATA     `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 空值
}

func GetUnsubscribeEvent(msg *message.MixedMsg) *UnsubscribeEvent {
	return &UnsubscribeEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
	}
}

// 进入应用
type EnterAgentEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType message.EventType `xml:"Event"    json:"Event"`    // enter_agent
	EventKey  message.CDATA     `xml:"EventKey" json:"EventKey"` // 事件KEY值, 空值
}

func GetEnterAgentEvent(msg *message.MixedMsg) *EnterAgentEvent {
	return &EnterAgentEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
	}
}

// 上报地理位置事件
type LocationEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType message.EventType `xml:"Event"     json:"Event"`                     // LOCATION
	Latitude  float64           `xml:"Latitude"  json:"Latitude"`                  // 地理位置纬度
	Longitude float64           `xml:"Longitude" json:"Longitude"`                 // 地理位置经度
	Precision float64           `xml:"Precision" json:"Precision"`                 // 地理位置精度(整数? 但是微信推送过来是浮点数形式)
	AppType   message.CDATA     `xml:"AppType,omitempty" json:"AppType,omitempty"` //app类型，在企业微信固定返回wxwork，在微信不返回该字段
}

func GetLocationEvent(msg *message.MixedMsg) *LocationEvent {
	return &LocationEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		Latitude:  msg.Latitude,
		Longitude: msg.Longitude,
		Precision: msg.Precision,
		AppType:   msg.AppType,
	}
}

// 异步任务完成事件推送
type BatchJobResultEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType message.EventType `xml:"Event"     json:"Event"`                     // batch_job_result
	JobId     message.CDATA     `xml:"JobId"  json:"JobId"`                        // 异步任务id，最大长度为64字符
	JobType   message.CDATA     `xml:"JobType" json:"JobType"`                     // 操作类型，字符串，目前分别有：sync_user(增量更新成员)、 replace_user(全量覆盖成员）、invite_user(邀请成员关注）、replace_party(全量覆盖部门)
	ErrCode   int               `xml:"ErrCode,omitempty" json:"ErrCode,omitempty"` // 返回码
	ErrMsg    message.CDATA     `xml:"ErrMsg,omitempty" json:"ErrMsg,omitempty"`   // 对返回码的文本描述内容
}

func GetBatchJobResultEvent(msg *message.MixedMsg) *BatchJobResultEvent {
	return &BatchJobResultEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		JobId:     msg.JobId,
		JobType:   msg.JobType,
		ErrCode:   msg.ErrCode,
		ErrMsg:    msg.ErrMsg,
	}
}

// 新增部门事件
type CreatePartyEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType  message.EventType `xml:"Event"     json:"Event"`        // change_contact
	ChangeType message.CDATA     `xml:"ChangeType"  json:"ChangeType"` // create_party
	Id         uint64            `xml:"Id" json:"Id"`                  // 部门Id
	Name       message.CDATA     `xml:"Name" json:"Name"`              // 部门名称
	ParentId   message.CDATA     `xml:"ParentId" json:"ParentId"`      // 父部门id
	Order      int               `xml:"Order" json:"Order"`            // 部门排序

}

func GetCreatePartyEvent(msg *message.MixedMsg) *CreatePartyEvent {
	return &CreatePartyEvent{
		MsgHeader:  msg.MsgHeader,
		EventType:  msg.EventType,
		ChangeType: msg.ChangeType,
		Id:         msg.Id,
		Name:       msg.Name,
		ParentId:   msg.ParentId,
		Order:      msg.Order,
	}
}

// 更新部门事件
type UpdatePartyEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType  message.EventType `xml:"Event"     json:"Event"`        // change_contact
	ChangeType message.CDATA     `xml:"ChangeType"  json:"ChangeType"` // update_party
	Id         uint64            `xml:"Id" json:"Id"`                  // 部门Id
	Name       message.CDATA     `xml:"Name" json:"Name"`              // 部门名称
	ParentId   message.CDATA     `xml:"ParentId" json:"ParentId"`      // 父部门id

}

func GetUpdatePartyEvent(msg *message.MixedMsg) *UpdatePartyEvent {
	return &UpdatePartyEvent{
		MsgHeader:  msg.MsgHeader,
		EventType:  msg.EventType,
		ChangeType: msg.ChangeType,
		Id:         msg.Id,
		Name:       msg.Name,
		ParentId:   msg.ParentId,
	}
}

// 删除部门事件
type DeletePartyEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType  message.EventType `xml:"Event"     json:"Event"`        // change_contact
	ChangeType message.CDATA     `xml:"ChangeType"  json:"ChangeType"` // update_party
	Id         uint64            `xml:"Id" json:"Id"`                  // 部门Id

}

func GetDeletePartyEvent(msg *message.MixedMsg) *DeletePartyEvent {
	return &DeletePartyEvent{
		MsgHeader:  msg.MsgHeader,
		EventType:  msg.EventType,
		ChangeType: msg.ChangeType,
		Id:         msg.Id,
	}
}

// 标签成员变更事件
type UpdateTagEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType     message.EventType `xml:"Event"     json:"Event"`                                 // change_contact
	ChangeType    message.CDATA     `xml:"ChangeType"  json:"ChangeType"`                          // update_tag
	TagId         uint64            `xml:"TagId" json:"TagId"`                                     // 标签Id
	AddUserItems  message.CDATA     `xml:"AddUserItems,omitempty" json:"AddUserItems,omitempty"`   // 标签中新增的成员userid列表，用逗号分隔
	DelUserItems  message.CDATA     `xml:"DelUserItems,omitempty" json:"DelUserItems,omitempty"`   // 标签中删除的成员userid列表，用逗号分隔
	AddPartyItems message.CDATA     `xml:"AddPartyItems,omitempty" json:"AddPartyItems,omitempty"` // 标签中新增的部门id列表，用逗号分隔
	DelPartyItems message.CDATA     `xml:"DelPartyItems,omitempty" json:"DelPartyItems,omitempty"` // 标签中删除的部门id列表，用逗号分隔
}

func GetUpdateTagEvent(msg *message.MixedMsg) *UpdateTagEvent {
	return &UpdateTagEvent{
		MsgHeader:     msg.MsgHeader,
		EventType:     msg.EventType,
		ChangeType:    msg.ChangeType,
		TagId:         msg.TagId,
		AddUserItems:  msg.AddUserItems,
		DelUserItems:  msg.DelUserItems,
		AddPartyItems: msg.AddPartyItems,
		DelPartyItems: msg.DelPartyItems,
	}
}

// 点击菜单拉取消息的事件推送
type ClickEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType message.EventType `xml:"Event"     json:"Event"`    // click
	EventKey  message.CDATA     `xml:"EventKey"  json:"EventKey"` // 事件KEY值，与自定义菜单接口中KEY值对应
}

func GetClickEvent(msg *message.MixedMsg) *ClickEvent {
	return &ClickEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
	}
}

// 点击菜单跳转链接的事件推送
type ViewEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType message.EventType `xml:"Event"     json:"Event"`    // view
	EventKey  message.CDATA     `xml:"EventKey"  json:"EventKey"` // 事件KEY值，与自定义菜单接口中KEY值对应
}

func GetViewEvent(msg *message.MixedMsg) *ViewEvent {
	return &ViewEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
	}
}

// 扫码推事件的事件推送
type ScanCodePushEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType    message.EventType    `xml:"Event"     json:"Event"`           // scancode_push
	EventKey     message.CDATA        `xml:"EventKey"  json:"EventKey"`        // 事件KEY值，与自定义菜单接口中KEY值对应
	ScanCodeInfo message.ScanCodeInfo `xml:"ScanCodeInfo" json:"ScanCodeInfo"` // 扫描信息
}

func GetScanCodePushEvent(msg *message.MixedMsg) *ScanCodePushEvent {
	return &ScanCodePushEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		EventKey:     msg.EventKey,
		ScanCodeInfo: msg.ScanCodeInfo,
	}
}

// 弹出系统拍照发图的事件推送
type PicSysPhotoEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType    message.EventType    `xml:"Event"     json:"Event"`           // pic_sysphoto
	EventKey     message.CDATA        `xml:"EventKey"  json:"EventKey"`        // 事件KEY值，与自定义菜单接口中KEY值对应
	SendPicsInfo message.SendPicsInfo `xml:"SendPicsInfo" json:"SendPicsInfo"` // 扫描信息
}

func GetPicSysPhotoEvent(msg *message.MixedMsg) *PicSysPhotoEvent {
	return &PicSysPhotoEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		EventKey:     msg.EventKey,
		SendPicsInfo: msg.SendPicsInfo,
	}
}

// 弹出拍照或者相册发图的事件推送
type PicPhotoOrAlbumEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType    message.EventType    `xml:"Event"     json:"Event"`           // pic_photo_or_album
	EventKey     message.CDATA        `xml:"EventKey"  json:"EventKey"`        // 事件KEY值，与自定义菜单接口中KEY值对应
	SendPicsInfo message.SendPicsInfo `xml:"SendPicsInfo" json:"SendPicsInfo"` // 扫描信息
}

func GetPicPhotoOrAlbumEvent(msg *message.MixedMsg) *PicPhotoOrAlbumEvent {
	return &PicPhotoOrAlbumEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		EventKey:     msg.EventKey,
		SendPicsInfo: msg.SendPicsInfo,
	}
}

// 弹出微信相册发图器的事件推送
type PicWeixinEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType    message.EventType    `xml:"Event"     json:"Event"`           // pic_weixin
	EventKey     message.CDATA        `xml:"EventKey"  json:"EventKey"`        // 事件KEY值，与自定义菜单接口中KEY值对应
	SendPicsInfo message.SendPicsInfo `xml:"SendPicsInfo" json:"SendPicsInfo"` // 扫描信息
}

func GetPicWeixinEvent(msg *message.MixedMsg) *PicWeixinEvent {
	return &PicWeixinEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		EventKey:     msg.EventKey,
		SendPicsInfo: msg.SendPicsInfo,
	}
}

// 弹出地理位置选择器的事件推送
type LocationSelectEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType        message.EventType        `xml:"Event"     json:"Event"`                     // location_select
	EventKey         message.CDATA            `xml:"EventKey"  json:"EventKey"`                  // 事件KEY值，与自定义菜单接口中KEY值对应
	SendLocationInfo message.SendLocationInfo `xml:"SendLocationInfo" json:"SendLocationInfo"`   // 发送的位置信息
	Latitude         float64                  `xml:"Latitude"  json:"Latitude"`                  // 地理位置纬度
	Longitude        float64                  `xml:"Longitude" json:"Longitude"`                 // 地理位置经度
	Precision        float64                  `xml:"Precision" json:"Precision"`                 // 地理位置精度(整数? 但是微信推送过来是浮点数形式)
	AppType          message.CDATA            `xml:"AppType,omitempty" json:"AppType,omitempty"` //app类型，在企业微信固定返回wxwork，在微信不返回该字段
}

func GetLocationSelectEvent(msg *message.MixedMsg) *LocationSelectEvent {
	return &LocationSelectEvent{
		MsgHeader:        msg.MsgHeader,
		EventType:        msg.EventType,
		EventKey:         msg.EventKey,
		SendLocationInfo: msg.SendLocationInfo,
		AppType:          msg.AppType,
	}
}

// 任务卡片事件推送
type TaskCardClickEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType message.EventType `xml:"Event"     json:"Event"`    // taskcard_click
	EventKey  message.CDATA     `xml:"EventKey"  json:"EventKey"` // 与发送任务卡片消息时指定的按钮btn:key值相同
	TaskId    message.CDATA     `xml:"TaskId" json:"TaskId"`      // 与发送任务卡片消息时指定的task_id相同
}

func GetTaskCardClickEvent(msg *message.MixedMsg) *TaskCardClickEvent {
	return &TaskCardClickEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
		TaskId:    msg.TaskId,
	}
}

// 审批状态通知事件
type OpenApprovalChangeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType    message.EventType    `xml:"Event"     json:"Event"`           // open_approval_change
	ApprovalInfo message.ApprovalInfo `xml:"ApprovalInfo" json:"ApprovalInfo"` // 审批信息
}

func GetOpenApprovalChangeEvent(msg *message.MixedMsg) *OpenApprovalChangeEvent {
	return &OpenApprovalChangeEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		ApprovalInfo: msg.ApprovalInfo,
	}
}

type AddExternalContactEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType      message.EventType `xml:"Event"     json:"Event"`               // change_external_contact
	ChangeType     message.CDATA     `xml:"ChangeType"  json:"ChangeType"`        // add_external_contact
	UserID         message.CDATA     `xml:"UserID" json:"UserID"`                 // 企业服务人员的UserID
	ExternalUserID message.CDATA     `xml:"ExternalUserID" json:"ExternalUserID"` // 外部联系人的userid，注意不是企业成员的帐号
	State          message.CDATA     `xml:"State" json:"State"`                   // 添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
	WelcomeCode    message.CDATA     `xml:"WelcomeCode" json:"WelcomeCode"`       // 欢迎语code
}

func GetAddExternalContactEvent(msg *message.MixedMsg) *AddExternalContactEvent {
	return &AddExternalContactEvent{
		MsgHeader:      msg.MsgHeader,
		EventType:      msg.EventType,
		ChangeType:     msg.ChangeType,
		UserID:         msg.UserID,
		ExternalUserID: msg.ExternalUserID,
		State:          msg.State,
		WelcomeCode:    msg.WelcomeCode,
	}
}

type EditExternalContactEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType      message.EventType `xml:"Event"     json:"Event"`               // change_external_contact
	ChangeType     message.CDATA     `xml:"ChangeType"  json:"ChangeType"`        // edit_external_contact
	UserID         message.CDATA     `xml:"UserID" json:"UserID"`                 // 企业服务人员的UserID
	ExternalUserID message.CDATA     `xml:"ExternalUserID" json:"ExternalUserID"` // 外部联系人的userid，注意不是企业成员的帐号
	State          message.CDATA     `xml:"State" json:"State"`                   // 添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
}

func GetEditExternalContactEvent(msg *message.MixedMsg) *EditExternalContactEvent {
	return &EditExternalContactEvent{
		MsgHeader:      msg.MsgHeader,
		EventType:      msg.EventType,
		ChangeType:     msg.ChangeType,
		UserID:         msg.UserID,
		ExternalUserID: msg.ExternalUserID,
		State:          msg.State,
	}
}

type AddHalfExternalContactEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType      message.EventType `xml:"Event"     json:"Event"`               // change_external_contact
	ChangeType     message.CDATA     `xml:"ChangeType"  json:"ChangeType"`        // add_half_external_contact
	UserID         message.CDATA     `xml:"UserID" json:"UserID"`                 // 企业服务人员的UserID
	ExternalUserID message.CDATA     `xml:"ExternalUserID" json:"ExternalUserID"` // 外部联系人的userid，注意不是企业成员的帐号
	State          message.CDATA     `xml:"State" json:"State"`                   // 添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
	WelcomeCode    message.CDATA     `xml:"WelcomeCode" json:"WelcomeCode"`       // 欢迎语code
}

func GetAddHalfExternalContactEvent(msg *message.MixedMsg) *AddHalfExternalContactEvent {
	return &AddHalfExternalContactEvent{
		MsgHeader:      msg.MsgHeader,
		EventType:      msg.EventType,
		ChangeType:     msg.ChangeType,
		UserID:         msg.UserID,
		ExternalUserID: msg.ExternalUserID,
		State:          msg.State,
		WelcomeCode:    msg.WelcomeCode,
	}
}

type DelExternalContactEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType      message.EventType `xml:"Event"     json:"Event"`               // change_external_contact
	ChangeType     message.CDATA     `xml:"ChangeType"  json:"ChangeType"`        // del_external_contact
	UserID         message.CDATA     `xml:"UserID" json:"UserID"`                 // 企业服务人员的UserID
	ExternalUserID message.CDATA     `xml:"ExternalUserID" json:"ExternalUserID"` // 外部联系人的userid，注意不是企业成员的帐号
}

func GetDelExternalContactEvent(msg *message.MixedMsg) *DelExternalContactEvent {
	return &DelExternalContactEvent{
		MsgHeader:      msg.MsgHeader,
		EventType:      msg.EventType,
		ChangeType:     msg.ChangeType,
		UserID:         msg.UserID,
		ExternalUserID: msg.ExternalUserID,
	}
}

type DelExternalContactFollowUserEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType      message.EventType `xml:"Event"     json:"Event"`               // change_external_contact
	ChangeType     message.CDATA     `xml:"ChangeType"  json:"ChangeType"`        // del_follow_user
	UserID         message.CDATA     `xml:"UserID" json:"UserID"`                 // 企业服务人员的UserID
	ExternalUserID message.CDATA     `xml:"ExternalUserID" json:"ExternalUserID"` // 外部联系人的userid，注意不是企业成员的帐号
}

func GetDelExternalContactFollowUserEvent(msg *message.MixedMsg) *DelExternalContactFollowUserEvent {
	return &DelExternalContactFollowUserEvent{
		MsgHeader:      msg.MsgHeader,
		EventType:      msg.EventType,
		ChangeType:     msg.ChangeType,
		UserID:         msg.UserID,
		ExternalUserID: msg.ExternalUserID,
	}
}

type GetChangeExternalChatEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	EventType message.EventType `xml:"Event"     json:"Event"` // change_external_chat
	ChatId    message.CDATA     `xml:"ChatId" json:"ChatId"`   // 群ID
}

func GetGetChangeExternalChatEvent(msg *message.MixedMsg) *GetChangeExternalChatEvent {
	return &GetChangeExternalChatEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		ChatId:    msg.ChatId,
	}
}
