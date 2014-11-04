// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package feedback

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
)

// 微信服务器推送过来的消息结构, xml 格式.
// 包含投诉消息, 确认消除投诉消息, 拒绝消除投诉消息
type MixedRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"     json:"AppId"`     // 公众号 id
	TimeStamp string `xml:"TimeStamp" json:"TimeStamp"` // 时间戳, unixtime

	OpenId     string `xml:"OpenId"     json:"OpenId"`     // 支付该笔订单的用户 OpenId
	FeedbackId string `xml:"FeedBackId" json:"FeedBackId"` // 投诉单号
	MsgType    string `xml:"MsgType"    json:"MsgType"`    // request, confirm, reject

	TransactionId string `xml:"TransId"  json:"TransId"`  // 交易订单号
	Reason        string `xml:"Reason"   json:"Reason"`   // 用户投诉的原因
	Solution      string `xml:"Solution" json:"Solution"` // 用户希望解决方案
	ExtInfo       string `xml:"ExtInfo"  json:"ExtInfo"`  // 备注+电话

	Signature  string `xml:"AppSignature" json:"AppSignature"` // 签名
	SignMethod string `xml:"SignMethod"   json:"SignMethod"`   // 签名方法, sha1

	PicInfo []struct {
		PicURL string `xml:"PicUrl" json:"PicUrl"`
	} `xml:"PicInfo>item,omitempty" json:"PicInfo,omitempty"` // 用户上传的图片凭证, 最多 5 张
}

// 检查 req *MixedRequest 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 即 paySignKey, 公众号支付请求中用于加密的密钥 Key
func (req *MixedRequest) CheckSignature(appKey string) (err error) {
	var Hash hash.Hash
	var hashsum []byte

	switch req.SignMethod {
	case "sha1", "SHA1":
		if len(req.Signature) != 40 {
			err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
				req.Signature, len(req.Signature), 40)
			return
		}

		Hash = sha1.New()
		hashsum = make([]byte, 40)

	default:
		err = fmt.Errorf(`unknown sign method: %q`, req.SignMethod)
		return
	}

	// 字典序
	// appid
	// appkey
	// openid
	// timestamp
	Hash.Write([]byte("appid="))
	Hash.Write([]byte(req.AppId))
	Hash.Write([]byte("&appkey="))
	Hash.Write([]byte(appKey))
	Hash.Write([]byte("&openid="))
	Hash.Write([]byte(req.OpenId))
	Hash.Write([]byte("&timestamp="))
	Hash.Write([]byte(req.TimeStamp))

	hex.Encode(hashsum, Hash.Sum(nil))

	if subtle.ConstantTimeCompare(hashsum, []byte(req.Signature)) != 1 {
		err = fmt.Errorf("签名不匹配, \r\nlocal: %q, \r\ninput: %q", hashsum, req.Signature)
		return
	}
	return
}

// 用户提交投诉消息, MsgType == request
type Complaint MixedRequest

func (this *Complaint) GetTimeStamp() (n int64, err error) {
	return strconv.ParseInt(this.TimeStamp, 10, 64)
}
func (this *Complaint) GetFeedbackId() (n int64, err error) {
	return strconv.ParseInt(this.FeedbackId, 10, 64)
}

// 从 MixedRequest 获取投诉消息
func (req *MixedRequest) GetComplaint() *Complaint {
	var ret Complaint = Complaint(*req)
	return &ret
}

// 用户确认消除投诉, MsgType == confirm
type Confirmation struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId     string `xml:"AppId"     json:"AppId"`     // 公众号 id
	TimeStamp string `xml:"TimeStamp" json:"TimeStamp"` // 时间戳, unixtime

	OpenId     string `xml:"OpenId"     json:"OpenId"`     // 支付该笔订单的用户 OpenId
	FeedbackId string `xml:"FeedBackId" json:"FeedBackId"` // 投诉单号
	MsgType    string `xml:"MsgType"    json:"MsgType"`    // confirm

	Reason string `xml:"Reason" json:"Reason"`

	Signature  string `xml:"AppSignature" json:"AppSignature"` // 签名
	SignMethod string `xml:"SignMethod"   json:"SignMethod"`   // 签名方法, sha1
}

func (this *Confirmation) GetTimeStamp() (n int64, err error) {
	return strconv.ParseInt(this.TimeStamp, 10, 64)
}
func (this *Confirmation) GetFeedbackId() (n int64, err error) {
	return strconv.ParseInt(this.FeedbackId, 10, 64)
}

// 从 MixedRequest 获取 Confirm 消息
func (req *MixedRequest) GetConfirmation() *Confirmation {
	return &Confirmation{
		AppId:     req.AppId,
		TimeStamp: req.TimeStamp,

		OpenId:     req.OpenId,
		FeedbackId: req.FeedbackId,
		MsgType:    req.MsgType,

		Reason: req.Reason,

		Signature:  req.Signature,
		SignMethod: req.SignMethod,
	}
}

// 用户拒绝消除投诉, MsgType == reject
type Rejection Confirmation

func (this *Rejection) GetTimeStamp() (n int64, err error) {
	return strconv.ParseInt(this.TimeStamp, 10, 64)
}
func (this *Rejection) GetFeedbackId() (n int64, err error) {
	return strconv.ParseInt(this.FeedbackId, 10, 64)
}

// 从 MixedRequest 获取 Reject 消息
func (req *MixedRequest) GetRejection() *Rejection {
	return (*Rejection)(req.GetConfirmation())
}
