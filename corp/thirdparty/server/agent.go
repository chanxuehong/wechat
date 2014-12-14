// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"net/http"
)

type InputParameters struct {
	w         http.ResponseWriter // 提供这个参数暂时不需要写任何数据, 预留的
	r         *http.Request       // r 的 Body 已经读取过了, 不要再读取了, 但是可以获取其他信息, 比如 r.URL.RawQuery
	rawXMLMsg []byte              // rawXMLMsg 是解密后的"明文" xml 消息体
	timestamp int64               // timestamp 是请求 URL 中的时间戳
	nonce     string              // nonce 是请求 URL 中的随机数
	AESKey    [32]byte            // AES 加密的 key
	random    []byte              // random 是请求 http body 中的密文消息加密时所用的 random, 16 bytes
}

// 套件对外暴露的接口
type Agent interface {
	GetSuiteId() string         // 套件Id
	GetToken() string           // 套件的Token
	GetLastAESKey() [32]byte    // 获取最后一个有效的 AES 加密 Key
	GetCurrentAESKey() [32]byte // 获取当前有效的 AES 加密 Key

	// 未知类型的消息处理方法
	ServeUnknownMsg(para *InputParameters)

	// 消息处理函数
	ServeSuiteTicketMsg(para *InputParameters, msg *SuiteTicket)
	ServeChangeAuthMsg(para *InputParameters, msg *ChangeAuth)
	ServeCancelAuthMsg(para *InputParameters, msg *CancelAuth)
}
