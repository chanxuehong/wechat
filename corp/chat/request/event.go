// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

import (
	"github.com/zihuxinyu/wechat/corp/chat"
)

const (
// 微信服务器推送过来的事件类型
	EventTypeSubscribe = "subscribe"       // 订阅
	EventTypeUnsubscribe = "unsubscribe"     // 取消订阅
	EventTypeEnterAgent = "enter_agent"     // 进入企业号消息服务
	EventTypeCreateChat = "create_chat"     //创建会话
	EventTypeQuitChat = "quit_chat"       //创建会话
	EventTypeUpdateChat = "update_chat"     //创建会话
)

// 关注事件
type SubscribeEvent struct {
	XMLName struct {}   `xml:"xml" json:"-"`
	chat.ItemHeader
	Event   string      `xml:"Event" json:"Event"` // 事件类型, subscribe(订阅)
}

func GetSubscribeEvent(msg *chat.MixedMessage) *SubscribeEvent {
	return &SubscribeEvent{
		ItemHeader   :msg.CurrentItem.ItemHeader,
		Event       :msg.CurrentItem.Event,
	}
}

// 取消关注
type UnsubscribeEvent struct {
	XMLName struct {}   `xml:"xml" json:"-"`
	chat.ItemHeader
	Event   string      `xml:"Event" json:"Event"` // 事件类型, unsubscribe(取消订阅)
}

func GetUnsubscribeEvent(msg *chat.MixedMessage) *UnsubscribeEvent {
	return &UnsubscribeEvent{
		ItemHeader   :msg.CurrentItem.ItemHeader,
		Event       :msg.CurrentItem.Event,
	}
}



// 进入企业号
type EnterAgentEvent struct {
	XMLName struct {}   `xml:"xml" json:"-"`
	chat.ItemHeader
	Event   string      `xml:"Event" json:"Event"` // 事件类型, enter_agent（进入企业号消息服务）
}

func GetEnterAgentEvent(msg *chat.MixedMessage) *EnterAgentEvent {
	return &EnterAgentEvent{
		ItemHeader   :msg.CurrentItem.ItemHeader,
		Event       :msg.CurrentItem.Event,
	}
}

// 创建会话
type CreateChatEvent struct {
	XMLName struct {}   `xml:"xml" json:"-"`
	chat.ItemHeader
	chat.ChatInfo
	Event   string      `xml:"Event" json:"Event"` // 事件类型, 此时固定为：create_chat
}

func GetCreateChatEvent(msg *chat.MixedMessage) *CreateChatEvent {
	return &CreateChatEvent{
		ItemHeader   :msg.CurrentItem.ItemHeader,
		ChatInfo    :msg.CurrentItem.ChatInfo,
		Event       :msg.CurrentItem.Event,
	}
}



// 退出会话
type QuitChatEvent struct {
	XMLName struct {}   `xml:"xml" json:"-"`
	chat.ItemHeader
	Event   string      `xml:"Event" json:"Event"` // 事件类型, 此时固定为：quit_chat
	ChatId  string      `xml:"ChatId"  json:"ChatId"`

}

func GetQuitChatEvent(msg *chat.MixedMessage) *QuitChatEvent {
	return &QuitChatEvent{
		ItemHeader   :msg.CurrentItem.ItemHeader,
		Event       :msg.CurrentItem.Event,
		ChatId      :msg.CurrentItem.ChatId,
	}
}


//修改会话
type UpdateChatEvent struct {
	XMLName     struct {} `xml:"xml" json:"-"`
	chat.ItemHeader
	Event       string   `xml:"Event" json:"Event"`              // 事件类型, 此时固定为：update_chat
	Name        string   `xml:"Name"  json:"Name"`               //会话名称
	Owner       string   `xml:"Owner"  json:"Owner"`             //会话所有者（管理员）
	AddUserList string   `xml:"AddUserList"  json:"AddUserList"` // 会话新增成员列表，成员间以竖线分隔，如：zhangsan|lisi
	DelUserList string   `xml:"DelUserList"  json:"DelUserList"` // 会话删除成员列表，成员间以竖线分隔，如：zhangsan|lisi
	ChatId      string   `xml:"ChatId"  json:"ChatId"`           // 会话id
}

func GetUpdateChatEvent(msg *chat.MixedMessage) *UpdateChatEvent {
	return &UpdateChatEvent{
		ItemHeader  :msg.CurrentItem.ItemHeader,
		Event      :msg.CurrentItem.Event,
		Name       :msg.CurrentItem.Name,
		Owner      :msg.CurrentItem.Owner,
		AddUserList:msg.CurrentItem.AddUserList,
		DelUserList:msg.CurrentItem.DelUserList,
		ChatId     :msg.CurrentItem.ChatId,
	}
}
