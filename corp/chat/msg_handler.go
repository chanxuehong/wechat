// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package chat

import (
	"net/http"
	"net/url"
)

// 微信服务器推送过来的消息(事件)处理接口
type MessageHandler interface {
	ServeMessage(http.ResponseWriter, *Request)
}

type MessageHandlerFunc func(http.ResponseWriter, *Request)

func (fn MessageHandlerFunc) ServeMessage(w http.ResponseWriter, r *Request) {
	fn(w, r)
}

// 消息(事件)请求信息
type Request struct {
	HttpRequest  *http.Request // 可以为 nil, 因为某些 http 框架没有提供此参数


	QueryValues  url.Values    // 回调请求 URL 中的查询参数集合
	MsgSignature string        // 回调请求 URL 中的消息体签名: msg_signature
	Timestamp    int64         // 回调请求 URL 中的时间戳: timestamp
	Nonce        string        // 回调请求 URL 中的随机数: nonce

	RawMsgXML    []byte        // 消息的"明文"XML 文本
	MixedMsg     *MixedMessage // RawMsgXML 解析后的消息

	AESKey       [32]byte      // 当前消息 AES 加密的 key
	Random       []byte        // 当前消息加密时所用的 random, 16 bytes

                               // 下面字段是企业号应用的基本信息
	CorpId       string        // 请求消息所属企业号的 ID
	CorpToken    string        // 请求消息所属企业号应用的 Token

}

// 微信服务器推送过来的消息(事件)通用的消息头
type MessageHeader struct {
	AgentType  string `xml:"AgentType"     json:"AgentType"` //固定为chat
	ToUserName string `xml:"ToUserName"    json:"ToUserName"`
	ItemCount  int64  `xml:"ItemCount"     json:"ItemCount"`
	PackageId  int64  `xml:"PackageId"     json:"PackageId"`

}

// 微信服务器推送过来的消息(事件)的合集.
type MixedMessage struct {
	XMLName     struct {} `xml:"xml"       json:"-"`
	MessageHeader
	Item        []Item    `xml:"Item"      json:"-"`
	CurrentItem Item //当前处理的Item
}


// 微信服务器推送过来的Item(事件)通用的消息头
type ItemHeader struct {
	FromUserName string   `xml:"FromUserName"  json:"FromUserName"` //成员UserID
	CreateTime   string   `xml:"CreateTime"    json:"CreateTime  "`   // 消息创建时间（整型）
	MsgType      string   `xml:"MsgType"       json:"MsgType "`          //消息类型

}

// 微信服务器推送过来的消息(事件)的合集.
type Item struct {

	ItemHeader
	Event       string   `xml:"Event"        json:"Event"`
	MsgId       int64    `xml:"MsgId"        json:"MsgId"`
	Name        string   `xml:"Name"         json:"Name"`
	Owner       string   `xml:"Owner"        json:"Owner"`
	AddUserList string   `xml:"AddUserList"  json:"AddUserList"`
	DelUserList string   `xml:"DelUserList"  json:"DelUserList"`
	ChatId      string   `xml:"ChatId"       json:"ChatId"`
	ChatInfo    ChatInfo `xml:"ChatInfo"     json:"ChatInfo"`
	Content     string   `xml:"Content"      json:"Content"`
	Receiver    Receiver `xml:"Receiver"     json:"Receiver"`
	PicUrl      string   `xml:"PicUrl"       json:"PicUrl"`
	MediaId     string   `xml:"MediaId"      json:"MediaId"`


}


type  Receiver     struct {
	Type string   `xml:"Type"  json:"Type"` //接收人类型：single|group，分别表示：群聊|单聊
	Id   string   `xml:"Id"    json:"Id"`     //接收人的值，为userid|chatid，分别表示：成员id|会话id
}


type ChatInfo struct {
	ChatId   string `xml:"ChatId"   json:"chatid"` // 会话id
	Name     string `xml:"Name"     json:"name"`   // 会话标题
	Owner    string `xml:"Owner"    json:"owner"`
	UserList string `xml:"UserList" json:"userlist"`
}
