// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"github.com/chanxuehong/wechat/message/passive/request"
	"net/http"
)

// 公众号对外暴露的接口
type Agent interface {
	// 获取公众号的原始ID, 等于后台中的 公众号设置-->帐号详情-->原始ID
	GetId() string

	// 获取公众号的 Token, 和后台中的设置相等
	GetToken() string

	// 未知类型的消息处理方法
	//  rawXMLMsg 是 xml 消息体
	//  timestamp 是请求 URL 中的时间戳
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 user-agent
	ServeUnknownMsg(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte, timestamp int64)

	// 消息处理函数
	//  rawXMLMsg 是 xml 消息体
	//  timestamp 是请求 URL 中的时间戳
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 user-agent
	ServeTextMsg(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64)
	ServeImageMsg(w http.ResponseWriter, r *http.Request, msg *request.Image, rawXMLMsg []byte, timestamp int64)
	ServeVoiceMsg(w http.ResponseWriter, r *http.Request, msg *request.Voice, rawXMLMsg []byte, timestamp int64)
	ServeVideoMsg(w http.ResponseWriter, r *http.Request, msg *request.Video, rawXMLMsg []byte, timestamp int64)
	ServeLocationMsg(w http.ResponseWriter, r *http.Request, msg *request.Location, rawXMLMsg []byte, timestamp int64)
	ServeLinkMsg(w http.ResponseWriter, r *http.Request, msg *request.Link, rawXMLMsg []byte, timestamp int64)

	// 事件处理函数
	//  rawXMLMsg 是 xml 消息体
	//  timestamp 是请求 URL 中的时间戳
	//  r *http.Request 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 user-agent
	ServeSubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeEvent, rawXMLMsg []byte, timestamp int64)
	ServeUnsubscribeEvent(w http.ResponseWriter, r *http.Request, msg *request.UnsubscribeEvent, rawXMLMsg []byte, timestamp int64)
	ServeSubscribeByScanEvent(w http.ResponseWriter, r *http.Request, msg *request.SubscribeByScanEvent, rawXMLMsg []byte, timestamp int64)
	ServeScanEvent(w http.ResponseWriter, r *http.Request, msg *request.ScanEvent, rawXMLMsg []byte, timestamp int64)
	ServeLocationEvent(w http.ResponseWriter, r *http.Request, msg *request.LocationEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuClickEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuClickEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuViewEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuViewEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuScanCodePushEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodePushEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuScanCodeWaitMsgEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodeWaitMsgEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuPicSysPhotoEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicSysPhotoEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuPicPhotoOrAlbumEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicPhotoOrAlbumEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuPicWeixinEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuPicWeixinEvent, rawXMLMsg []byte, timestamp int64)
	ServeMenuLocationSelectEvent(w http.ResponseWriter, r *http.Request, msg *request.MenuLocationSelectEvent, rawXMLMsg []byte, timestamp int64)
	ServeMassSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.MassSendJobFinishEvent, rawXMLMsg []byte, timestamp int64)
	ServeTemplateSendJobFinishEvent(w http.ResponseWriter, r *http.Request, msg *request.TemplateSendJobFinishEvent, rawXMLMsg []byte, timestamp int64)
	ServeMerchantOrderEvent(w http.ResponseWriter, r *http.Request, msg *request.MerchantOrderEvent, rawXMLMsg []byte, timestamp int64)
}
