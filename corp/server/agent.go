// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"net/http"

	"github.com/chanxuehong/wechat/corp/message/passive/request"
)

type RequestParameters struct {
	HTTPResponseWriter http.ResponseWriter // 用于回复
	HTTPRequest        *http.Request       // r 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 r.URL.RawQuery
	Timestamp          int64               // timestamp 是请求 URL 中的时间戳
	Nonce              string              // nonce 是请求 URL 中的随机数
	RawXMLMsg          []byte              // rawXMLMsg 是"明文" xml 消息体
	AESKey             [32]byte            // AES 加密的 key
	Random             []byte              // random 是请求 http body 中的密文消息加密时所用的 random, 16 bytes
}

// 企业号单个应用对外暴露的接口
type Agent interface {
	GetCorpId() string   // 获取应用所属的企业号Id
	GetAgentId() int64   // 获取应用的Id
	GetToken() string    // 获取应用的Token
	GetAESKey() [32]byte // 获取 32bytes 的 AES 加密 Key

	// 未知类型的消息处理方法
	ServeUnknownMsg(para *RequestParameters)

	// 消息处理函数
	ServeTextMsg(msg *request.Text, para *RequestParameters)
	ServeImageMsg(msg *request.Image, para *RequestParameters)
	ServeVoiceMsg(msg *request.Voice, para *RequestParameters)
	ServeVideoMsg(msg *request.Video, para *RequestParameters)
	ServeLocationMsg(msg *request.Location, para *RequestParameters)

	// 事件处理函数
	ServeSubscribeEvent(event *request.SubscribeEvent, para *RequestParameters)
	ServeUnsubscribeEvent(event *request.UnsubscribeEvent, para *RequestParameters)
	ServeLocationEvent(event *request.LocationEvent, para *RequestParameters)
	ServeMenuClickEvent(event *request.MenuClickEvent, para *RequestParameters)
	ServeMenuViewEvent(event *request.MenuViewEvent, para *RequestParameters)
	ServeMenuScanCodePushEvent(event *request.MenuScanCodePushEvent, para *RequestParameters)
	ServeMenuScanCodeWaitMsgEvent(event *request.MenuScanCodeWaitMsgEvent, para *RequestParameters)
	ServeMenuPicSysPhotoEvent(event *request.MenuPicSysPhotoEvent, para *RequestParameters)
	ServeMenuPicPhotoOrAlbumEvent(event *request.MenuPicPhotoOrAlbumEvent, para *RequestParameters)
	ServeMenuPicWeixinEvent(event *request.MenuPicWeixinEvent, para *RequestParameters)
	ServeMenuLocationSelectEvent(event *request.MenuLocationSelectEvent, para *RequestParameters)
}
