// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"net/http"

	"github.com/chanxuehong/wechat/mp/message/passive/request"
)

type RequestParameters struct {
	HTTPResponseWriter http.ResponseWriter // 用于回复
	HTTPRequest        *http.Request       // r 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 r.URL.RawQuery
	Timestamp          int64               // timestamp 是请求 URL 中的时间戳
	Nonce              string              // nonce 是请求 URL 中的随机数
	RawXMLMsg          []byte              // rawXMLMsg 是"明文" xml 消息体
}

type AESRequestParameters struct {
	HTTPResponseWriter http.ResponseWriter // 用于回复
	HTTPRequest        *http.Request       // r 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 r.URL.RawQuery
	Timestamp          int64               // timestamp 是请求 URL 中的时间戳
	Nonce              string              // nonce 是请求 URL 中的随机数
	RawXMLMsg          []byte              // rawXMLMsg 是"明文" xml 消息体
	AESKey             [32]byte            // AES 加密的 key
	Random             []byte              // random 是请求 http body 中的密文消息加密时所用的 random, 16 bytes
}

// 公众号对外暴露的接口
type Agent interface {
	GetId() string    // 获取公众号的原始ID, 等于后台中的 公众号设置-->帐号详情-->原始ID
	GetToken() string // 获取公众号的Token, 和后台中的设置相等

	// fuck, AppId 貌似需要认证才会有的???
	// 如果不知道自己的 AppId 是多少, 可以先随便填入一个字符串,
	// 这样正常情况下会出现 AppId mismatch 错误, 错误中 have 后面的就是正确的 AppId
	GetAppId() string
	GetLastAESKey() [32]byte    // 获取最后一个有效的 AES 加密 Key
	GetCurrentAESKey() [32]byte // 获取当前有效的 AES 加密 Key

	// 明文模式 需要实现的方法
	// 未知类型的消息处理方法
	ServeUnknownMsg(para *RequestParameters)

	// 明文模式 需要实现的方法
	// 消息处理函数
	ServeTextMsg(msg *request.Text, para *RequestParameters)
	ServeImageMsg(msg *request.Image, para *RequestParameters)
	ServeVoiceMsg(msg *request.Voice, para *RequestParameters)
	ServeVideoMsg(msg *request.Video, para *RequestParameters)
	ServeLocationMsg(msg *request.Location, para *RequestParameters)
	ServeLinkMsg(msg *request.Link, para *RequestParameters)

	// 明文模式 需要实现的方法
	// 事件处理函数
	ServeSubscribeEvent(event *request.SubscribeEvent, para *RequestParameters)
	ServeUnsubscribeEvent(event *request.UnsubscribeEvent, para *RequestParameters)
	ServeSubscribeByScanEvent(event *request.SubscribeByScanEvent, para *RequestParameters)
	ServeScanEvent(event *request.ScanEvent, para *RequestParameters)
	ServeLocationEvent(event *request.LocationEvent, para *RequestParameters)
	ServeMenuClickEvent(event *request.MenuClickEvent, para *RequestParameters)
	ServeMenuViewEvent(event *request.MenuViewEvent, para *RequestParameters)
	ServeMenuScanCodePushEvent(event *request.MenuScanCodePushEvent, para *RequestParameters)
	ServeMenuScanCodeWaitMsgEvent(event *request.MenuScanCodeWaitMsgEvent, para *RequestParameters)
	ServeMenuPicSysPhotoEvent(event *request.MenuPicSysPhotoEvent, para *RequestParameters)
	ServeMenuPicPhotoOrAlbumEvent(event *request.MenuPicPhotoOrAlbumEvent, para *RequestParameters)
	ServeMenuPicWeixinEvent(event *request.MenuPicWeixinEvent, para *RequestParameters)
	ServeMenuLocationSelectEvent(event *request.MenuLocationSelectEvent, para *RequestParameters)
	ServeMassSendJobFinishEvent(event *request.MassSendJobFinishEvent, para *RequestParameters)
	ServeTemplateSendJobFinishEvent(event *request.TemplateSendJobFinishEvent, para *RequestParameters)
	ServeMerchantOrderEvent(event *request.MerchantOrderEvent, para *RequestParameters)

	// 兼容模式, 安全模式 需要实现的方法
	// 未知类型的消息处理方法
	ServeAESUnknownMsg(para *AESRequestParameters)

	// 兼容模式, 安全模式 需要实现的方法
	// 消息处理函数
	ServeAESTextMsg(msg *request.Text, para *AESRequestParameters)
	ServeAESImageMsg(msg *request.Image, para *AESRequestParameters)
	ServeAESVoiceMsg(msg *request.Voice, para *AESRequestParameters)
	ServeAESVideoMsg(msg *request.Video, para *AESRequestParameters)
	ServeAESLocationMsg(msg *request.Location, para *AESRequestParameters)
	ServeAESLinkMsg(msg *request.Link, para *AESRequestParameters)

	// 兼容模式, 安全模式 需要实现的方法
	// 事件处理函数
	ServeAESSubscribeEvent(event *request.SubscribeEvent, para *AESRequestParameters)
	ServeAESUnsubscribeEvent(event *request.UnsubscribeEvent, para *AESRequestParameters)
	ServeAESSubscribeByScanEvent(event *request.SubscribeByScanEvent, para *AESRequestParameters)
	ServeAESScanEvent(event *request.ScanEvent, para *AESRequestParameters)
	ServeAESLocationEvent(event *request.LocationEvent, para *AESRequestParameters)
	ServeAESMenuClickEvent(event *request.MenuClickEvent, para *AESRequestParameters)
	ServeAESMenuViewEvent(event *request.MenuViewEvent, para *AESRequestParameters)
	ServeAESMenuScanCodePushEvent(event *request.MenuScanCodePushEvent, para *AESRequestParameters)
	ServeAESMenuScanCodeWaitMsgEvent(event *request.MenuScanCodeWaitMsgEvent, para *AESRequestParameters)
	ServeAESMenuPicSysPhotoEvent(event *request.MenuPicSysPhotoEvent, para *AESRequestParameters)
	ServeAESMenuPicPhotoOrAlbumEvent(event *request.MenuPicPhotoOrAlbumEvent, para *AESRequestParameters)
	ServeAESMenuPicWeixinEvent(event *request.MenuPicWeixinEvent, para *AESRequestParameters)
	ServeAESMenuLocationSelectEvent(event *request.MenuLocationSelectEvent, para *AESRequestParameters)
	ServeAESMassSendJobFinishEvent(event *request.MassSendJobFinishEvent, para *AESRequestParameters)
	ServeAESTemplateSendJobFinishEvent(event *request.TemplateSendJobFinishEvent, para *AESRequestParameters)
	ServeAESMerchantOrderEvent(event *request.MerchantOrderEvent, para *AESRequestParameters)
}
