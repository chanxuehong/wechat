// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"net/http"

	"github.com/chanxuehong/wechat/corp/message/passive/request"
)

// 企业号单个应用对外暴露的接口
type Agent interface {
	GetCorpId() string   // 获取应用所属的企业号Id
	GetAgentId() int64   // 获取应用的Id
	GetToken() string    // 获取应用的Token
	GetAESKey() [32]byte // 获取 32bytes 的 AES 加密 Key

	// 未知类型的消息处理方法
	//  rawXMLMsg   是解密后的"明文" xml 消息体
	//  timestamp   是请求 URL 中的时间戳
	//  nonce       是请求 URL 中的随机数
	//  random      是请求 http body 中的密文消息加密时所用的 random, 16 bytes
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 r.URL.RawQuery
	ServeUnknownMsg(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)

	// 消息处理函数
	//  msg 是成功解析的消息结构体
	//  rawXMLMsg   是解密后的"明文" xml 消息体
	//  timestamp   是请求 URL 中的时间戳
	//  nonce       是请求 URL 中的随机数
	//  random      是请求 http body 中的密文消息加密时所用的 random, 16 bytes
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 r.URL.RawQuery
	ServeTextMsg(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeImageMsg(w http.ResponseWriter, r *http.Request, msg *request.Image, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeVoiceMsg(w http.ResponseWriter, r *http.Request, msg *request.Voice, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeVideoMsg(w http.ResponseWriter, r *http.Request, msg *request.Video, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeLocationMsg(w http.ResponseWriter, r *http.Request, msg *request.Location, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)

	// 事件处理函数
	//  event 是成功解析的消息结构体
	//  rawXMLMsg   是解密后的"明文" xml 消息体
	//  timestamp   是请求 URL 中的时间戳
	//  nonce       是请求 URL 中的随机数
	//  random      是请求 http body 中的密文消息加密时所用的 random, 16 bytes
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 r.URL.RawQuery
	ServeSubscribeEvent(w http.ResponseWriter, r *http.Request, event *request.SubscribeEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeUnsubscribeEvent(w http.ResponseWriter, r *http.Request, event *request.UnsubscribeEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeLocationEvent(w http.ResponseWriter, r *http.Request, event *request.LocationEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeMenuClickEvent(w http.ResponseWriter, r *http.Request, event *request.MenuClickEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeMenuViewEvent(w http.ResponseWriter, r *http.Request, event *request.MenuViewEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeMenuScanCodePushEvent(w http.ResponseWriter, r *http.Request, event *request.MenuScanCodePushEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeMenuScanCodeWaitMsgEvent(w http.ResponseWriter, r *http.Request, event *request.MenuScanCodeWaitMsgEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeMenuPicSysPhotoEvent(w http.ResponseWriter, r *http.Request, event *request.MenuPicSysPhotoEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeMenuPicPhotoOrAlbumEvent(w http.ResponseWriter, r *http.Request, event *request.MenuPicPhotoOrAlbumEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeMenuPicWeixinEvent(w http.ResponseWriter, r *http.Request, event *request.MenuPicWeixinEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
	ServeMenuLocationSelectEvent(w http.ResponseWriter, r *http.Request, event *request.MenuLocationSelectEvent, rawXMLMsg []byte, timestamp int64, nonce string, random []byte)
}
