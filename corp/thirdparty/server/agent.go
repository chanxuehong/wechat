// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"net/http"
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

// 套件对外暴露的接口
type Agent interface {
	GetSuiteId() string         // 套件Id
	GetToken() string           // 套件的Token
	GetLastAESKey() [32]byte    // 获取最后一个有效的 AES 加密 Key
	GetCurrentAESKey() [32]byte // 获取当前有效的 AES 加密 Key

	// 未知类型的消息处理方法
	ServeUnknownMsg(para *RequestParameters)

	// 消息处理函数
	ServeSuiteTicketMsg(msg *SuiteTicket, para *RequestParameters)
	ServeChangeAuthMsg(msg *ChangeAuth, para *RequestParameters)
	ServeCancelAuthMsg(msg *CancelAuth, para *RequestParameters)
}
