// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"github.com/chanxuehong/wechat/message/passive/request"
	"net/http"
)

// 消息处理接口
type MsgHandler interface {
	// 生成签名
	Signature(timestamp, nonce string) (signature string)

	// 未知类型的消息处理方法
	//  rawXMLMsg 是 xml 消息体
	//  timestamp 是请求中的时间戳
	UnknownMsgHandler(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte, timestamp int64)

	// 消息处理函数
	//  rawXMLMsg 是 xml 消息体
	//  timestamp 是请求中的时间戳
	TextMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64)
	ImageMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Image, rawXMLMsg []byte, timestamp int64)
	VoiceMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Voice, rawXMLMsg []byte, timestamp int64)
	VideoMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Video, rawXMLMsg []byte, timestamp int64)
	LocationMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Location, rawXMLMsg []byte, timestamp int64)
	LinkMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Link, rawXMLMsg []byte, timestamp int64)

	// 事件处理函数
	//  rawXMLMsg 是 xml 消息体
	//  timestamp 是请求中的时间戳
	SubscribeEventHandler(w http.ResponseWriter, r *http.Request, msg *request.SubscribeEvent, rawXMLMsg []byte, timestamp int64)
	UnsubscribeEventHandler(w http.ResponseWriter, r *http.Request, msg *request.UnsubscribeEvent, rawXMLMsg []byte, timestamp int64)
	SubscribeByScanEventHandler(w http.ResponseWriter, r *http.Request, msg *request.SubscribeByScanEvent, rawXMLMsg []byte, timestamp int64)
	ScanEventHandler(w http.ResponseWriter, r *http.Request, msg *request.ScanEvent, rawXMLMsg []byte, timestamp int64)
	LocationEventHandler(w http.ResponseWriter, r *http.Request, msg *request.LocationEvent, rawXMLMsg []byte, timestamp int64)
	MenuClickEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuClickEvent, rawXMLMsg []byte, timestamp int64)
	MenuViewEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuViewEvent, rawXMLMsg []byte, timestamp int64)
	MenuScanCodePushEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodePushEvent, rawXMLMsg []byte, timestamp int64)
	MenuScanCodeWaitMsgEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodeWaitMsgEvent, rawXMLMsg []byte, timestamp int64)
	MenuPicSysPhotoEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuPicSysPhotoEvent, rawXMLMsg []byte, timestamp int64)
	MenuPicPhotoOrAlbumEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuPicPhotoOrAlbumEvent, rawXMLMsg []byte, timestamp int64)
	MenuPicWeixinEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuPicWeixinEvent, rawXMLMsg []byte, timestamp int64)
	MenuLocationSelectEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuLocationSelectEvent, rawXMLMsg []byte, timestamp int64)
	MassSendJobFinishEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MassSendJobFinishEvent, rawXMLMsg []byte, timestamp int64)
	TemplateSendJobFinishEventHandler(w http.ResponseWriter, r *http.Request, msg *request.TemplateSendJobFinishEvent, rawXMLMsg []byte, timestamp int64)
	MerchantOrderEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MerchantOrderEvent, rawXMLMsg []byte, timestamp int64)
}
