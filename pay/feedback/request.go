// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package feedback

import (
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

// 微信服务器推送过来的消息结构.
// 包含投诉消息, 确认消除投诉消息, 拒绝消除投诉消息
type MsgRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"`     // 公众号 id
	TimeStamp int64  `xml:"TimeStamp"` // 时间戳, unixtime

	OpenId     string `xml:"OpenId"`     // 支付该笔订单的用户 OpenId
	FeedbackId int64  `xml:"FeedBackId"` // 投诉单号
	MsgType    string `xml:"MsgType"`    // request

	TransactionId string `xml:"TransId"`  // 交易订单号
	Reason        string `xml:"Reason"`   // 用户投诉的原因
	Solution      string `xml:"Solution"` // 用户希望解决方案
	ExtInfo       string `xml:"ExtInfo"`  // 备注+电话

	Signature  string `xml:"AppSignature"` // 签名
	SignMethod string `xml:"SignMethod"`   // 签名方法, sha1

	PicInfo []struct {
		PicURL string `xml:"PicUrl"`
	} `xml:"PicInfo>item"` // 用户上传的图片凭证, 最多 5 张
}

// 从 MsgRequest 获取投诉消息
func (msg *MsgRequest) GetRequest() *Request {
	return (*Request)(msg)
}

// 从 MsgRequest 获取 Confirm 消息
func (msg *MsgRequest) GetConfirm() *Confirm {
	return &Confirm{
		XMLName: msg.XMLName,

		AppId:     msg.AppId,
		TimeStamp: msg.TimeStamp,

		OpenId:     msg.OpenId,
		FeedbackId: msg.FeedbackId,
		MsgType:    msg.MsgType,

		Reason: msg.Reason,

		Signature:  msg.Signature,
		SignMethod: msg.SignMethod,
	}
}

// 从 MsgRequest 获取 Reject 消息
func (msg *MsgRequest) GetReject() *Reject {
	return &Reject{
		XMLName: msg.XMLName,

		AppId:     msg.AppId,
		TimeStamp: msg.TimeStamp,

		OpenId:     msg.OpenId,
		FeedbackId: msg.FeedbackId,
		MsgType:    msg.MsgType,

		Reason: msg.Reason,

		Signature:  msg.Signature,
		SignMethod: msg.SignMethod,
	}
}

// 用户提交投诉消息
type Request struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"`     // 公众号 id
	TimeStamp int64  `xml:"TimeStamp"` // 时间戳, unixtime

	OpenId     string `xml:"OpenId"`     // 支付该笔订单的用户 OpenId
	FeedbackId int64  `xml:"FeedBackId"` // 投诉单号
	MsgType    string `xml:"MsgType"`    // request

	TransactionId string `xml:"TransId"`  // 交易订单号
	Reason        string `xml:"Reason"`   // 用户投诉的原因
	Solution      string `xml:"Solution"` // 用户希望解决方案
	ExtInfo       string `xml:"ExtInfo"`  // 备注+电话

	Signature  string `xml:"AppSignature"` // 签名
	SignMethod string `xml:"SignMethod"`   // 签名方法, sha1

	PicInfo []struct {
		PicURL string `xml:"PicUrl"`
	} `xml:"PicInfo>item"` // 用户上传的图片凭证, 最多 5 张
}

// 检查 req *Request 是否合法(包括签名的检查), 合法返回 nil, 否则返回错误信息.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
func (req *Request) Check(paySignKey string) (err error) {
	// 检查签名
	var hashSumLen, twoHashSumLen int
	var sumFunc hashSumFunc

	switch req.SignMethod {
	case "sha1", "SHA1":
		hashSumLen = 40
		twoHashSumLen = 80
		sumFunc = sha1Sum

	default:
		err = fmt.Errorf(`not implement for "%s" sign method`, req.SignMethod)
		return
	}

	if len(req.Signature) != hashSumLen {
		err = errors.New("签名不正确")
		return
	}

	FeedbackStr := strconv.FormatInt(req.FeedbackId, 10)
	TimeStampStr := strconv.FormatInt(req.TimeStamp, 10)

	const keysLen = len(`appid=&appkey=&extinfo=&feedbackid=&msgtype=&openid=&reason=&solution=&timestamp=&transid=`)

	n := twoHashSumLen + keysLen + len(req.AppId) + len(paySignKey) + len(req.ExtInfo) +
		len(FeedbackStr) + len(req.MsgType) + len(req.OpenId) + len(req.Reason) +
		len(req.Solution) + len(TimeStampStr) + len(req.TransactionId)

	// buf[:hashSumLen] 保存参数 req.Signature,
	// buf[hashSumLen:twoHashSumLen] 保存生成的签名
	// buf[twoHashSumLen:] 按照字典序列保存 string1
	buf := make([]byte, n)
	reqSignature := buf[:hashSumLen]
	signature := buf[hashSumLen:twoHashSumLen]
	string1 := buf[twoHashSumLen:twoHashSumLen]

	copy(reqSignature, req.Signature)

	// 字典序
	// appid
	// appkey
	// extinfo
	// feedbackid
	// msgtype
	// openid
	// reason
	// solution
	// timestamp
	// transid
	string1 = append(string1, "appid="...)
	string1 = append(string1, req.AppId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&extinfo="...)
	string1 = append(string1, req.ExtInfo...)
	string1 = append(string1, "&feedbackid="...)
	string1 = append(string1, FeedbackStr...)
	string1 = append(string1, "&msgtype="...)
	string1 = append(string1, req.MsgType...)
	string1 = append(string1, "&openid="...)
	string1 = append(string1, req.OpenId...)
	string1 = append(string1, "&reason="...)
	string1 = append(string1, req.Reason...)
	string1 = append(string1, "&solution="...)
	string1 = append(string1, req.Solution...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, TimeStampStr...)
	string1 = append(string1, "&transid="...)
	string1 = append(string1, req.TransactionId...)

	hex.Encode(signature, sumFunc(string1))

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	if subtle.ConstantTimeCompare(reqSignature, signature) != 1 {
		err = errors.New("签名不正确")
		return
	}
	return
}

// 用户确认消除投诉
type Confirm struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"`     // 公众号 id
	TimeStamp int64  `xml:"TimeStamp"` // 时间戳, unixtime

	OpenId     string `xml:"OpenId"`     // 支付该笔订单的用户 OpenId
	FeedbackId int64  `xml:"FeedBackId"` // 投诉单号
	MsgType    string `xml:"MsgType"`    // confirm

	Reason string `xml:"Reason"`

	Signature  string `xml:"AppSignature"` // 签名
	SignMethod string `xml:"SignMethod"`   // 签名方法, sha1
}

// 检查 cfm *Confirm 是否合法(包括签名的检查), 合法返回 nil, 否则返回错误信息.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
func (cfm *Confirm) Check(paySignKey string) (err error) {
	// 检查签名
	var hashSumLen, twoHashSumLen int
	var sumFunc hashSumFunc

	switch cfm.SignMethod {
	case "sha1", "SHA1":
		hashSumLen = 40
		twoHashSumLen = 80
		sumFunc = sha1Sum

	default:
		err = fmt.Errorf(`not implement for "%s" sign method`, cfm.SignMethod)
		return
	}

	if len(cfm.Signature) != hashSumLen {
		err = errors.New("签名不正确")
		return
	}

	FeedbackStr := strconv.FormatInt(cfm.FeedbackId, 10)
	TimeStampStr := strconv.FormatInt(cfm.TimeStamp, 10)

	const keysLen = len(`appid=&appkey=&feedbackid=&msgtype=&openid=&reason=&timestamp=`)

	n := twoHashSumLen + keysLen + len(cfm.AppId) + len(paySignKey) + len(FeedbackStr) +
		len(cfm.MsgType) + len(cfm.OpenId) + len(cfm.Reason) + len(TimeStampStr)

	// buf[:hashSumLen] 保存参数 cfm.Signature,
	// buf[hashSumLen:twoHashSumLen] 保存生成的签名
	// buf[twoHashSumLen:] 按照字典序列保存 string1
	buf := make([]byte, n)
	cfmSignature := buf[:hashSumLen]
	signature := buf[hashSumLen:twoHashSumLen]
	string1 := buf[twoHashSumLen:twoHashSumLen]

	copy(cfmSignature, cfm.Signature)

	// 字典序
	// appid
	// appkey
	// feedbackid
	// msgtype
	// openid
	// reason
	// timestamp
	string1 = append(string1, "appid="...)
	string1 = append(string1, cfm.AppId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&feedbackid="...)
	string1 = append(string1, FeedbackStr...)
	string1 = append(string1, "&msgtype="...)
	string1 = append(string1, cfm.MsgType...)
	string1 = append(string1, "&openid="...)
	string1 = append(string1, cfm.OpenId...)
	string1 = append(string1, "&reason="...)
	string1 = append(string1, cfm.Reason...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, TimeStampStr...)

	hex.Encode(signature, sumFunc(string1))

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	if subtle.ConstantTimeCompare(cfmSignature, signature) != 1 {
		err = errors.New("签名不正确")
		return
	}
	return
}

// 用户拒绝消除投诉
type Reject struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"`     // 公众号 id
	TimeStamp int64  `xml:"TimeStamp"` // 时间戳, unixtime

	OpenId     string `xml:"OpenId"`     // 支付该笔订单的用户 OpenId
	FeedbackId int64  `xml:"FeedBackId"` // 投诉单号
	MsgType    string `xml:"MsgType"`    // reject

	Reason string `xml:"Reason"` // 拒绝原因

	Signature  string `xml:"AppSignature"` // 签名
	SignMethod string `xml:"SignMethod"`   // 签名方法, sha1
}

// 检查 rjt *Reject 是否合法(包括签名的检查), 合法返回 nil, 否则返回错误信息.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
func (rjt *Reject) Check(paySignKey string) (err error) {
	// 检查签名
	var hashSumLen, twoHashSumLen int
	var sumFunc hashSumFunc

	switch rjt.SignMethod {
	case "sha1", "SHA1":
		hashSumLen = 40
		twoHashSumLen = 80
		sumFunc = sha1Sum

	default:
		err = fmt.Errorf(`not implement for "%s" sign method`, rjt.SignMethod)
		return
	}

	if len(rjt.Signature) != hashSumLen {
		err = errors.New("签名不正确")
		return
	}

	FeedbackStr := strconv.FormatInt(rjt.FeedbackId, 10)
	TimeStampStr := strconv.FormatInt(rjt.TimeStamp, 10)

	const keysLen = len(`appid=&appkey=&feedbackid=&msgtype=&openid=&reason=&timestamp=`)

	n := twoHashSumLen + keysLen + len(rjt.AppId) + len(paySignKey) + len(FeedbackStr) +
		len(rjt.MsgType) + len(rjt.OpenId) + len(rjt.Reason) + len(TimeStampStr)

	// buf[:hashSumLen] 保存参数 rjt.Signature,
	// buf[hashSumLen:twoHashSumLen] 保存生成的签名
	// buf[twoHashSumLen:] 按照字典序列保存 string1
	buf := make([]byte, n)
	rjtSignature := buf[:hashSumLen]
	signature := buf[hashSumLen:twoHashSumLen]
	string1 := buf[twoHashSumLen:twoHashSumLen]

	copy(rjtSignature, rjt.Signature)

	// 字典序
	// appid
	// appkey
	// feedbackid
	// msgtype
	// openid
	// reason
	// timestamp
	string1 = append(string1, "appid="...)
	string1 = append(string1, rjt.AppId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&feedbackid="...)
	string1 = append(string1, FeedbackStr...)
	string1 = append(string1, "&msgtype="...)
	string1 = append(string1, rjt.MsgType...)
	string1 = append(string1, "&openid="...)
	string1 = append(string1, rjt.OpenId...)
	string1 = append(string1, "&reason="...)
	string1 = append(string1, rjt.Reason...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, TimeStampStr...)

	hex.Encode(signature, sumFunc(string1))

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	if subtle.ConstantTimeCompare(rjtSignature, signature) != 1 {
		err = errors.New("签名不正确")
		return
	}
	return
}
