// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package native

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strconv"
)

// 获取订单详情 package 的请求消息
type BillRequest struct {
	XMLName   struct{} `xml:"xml" json:"-"`
	AppId     string   `xml:"AppId"`     // 必须, 公众帐号的 appid
	NonceStr  string   `xml:"NonceStr"`  // 必须, 随机串
	TimeStamp int64    `xml:"TimeStamp"` // 必须, 时间戳

	ProductId string `xml:"ProductId"` // 必须, 第三方的商品	id 号

	OpenId      string `xml:"OpenId"`      // 必须, 点击链接准备购买商品的用户 OpenId
	IsSubscribe int    `xml:"IsSubscribe"` // 必须, 标记用户是否订阅该公众帐号, 1为关注, 0为未关注

	Signature  string `xml:"AppSignature"` // 必须, 参数的加密签名
	SignMethod string `xml:"SignMethod"`   // 必须, 签名方式, 目前只支持"SHA1", 该字段不参与签名
}

// 校验 req *BillRequest 的签名是否有效, isValid == true && err == nil 时表示有效.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
func (req *BillRequest) CheckSignature(paySignKey string) (isValid bool, err error) {
	switch req.SignMethod {
	case "sha1", "SHA1":
		const hashSumLen = 40 // sha1
		const twoHashSumLen = hashSumLen * 2

		if len(req.Signature) != hashSumLen {
			return
		}

		TimeStampStr := strconv.FormatInt(req.TimeStamp, 10)
		IsSubscribeStr := strconv.FormatInt(int64(req.IsSubscribe), 10)

		const keysLen = len(`appid=&appkey=&issubscribe=&noncestr=&openid=&productid=&timestamp=`)

		n := twoHashSumLen + keysLen + len(req.AppId) + len(paySignKey) + len(IsSubscribeStr) +
			len(req.NonceStr) + len(req.OpenId) + len(req.ProductId) + len(TimeStampStr)

		// buf[:hashSumLen] 保存参数 req.Signature,
		// buf[hashSumLen:twoHashSumLen] 保存生成的签名
		// buf[twoHashSumLen:] 按照字典序列保存 string1
		buf := make([]byte, n)
		copy(buf, req.Signature)
		signature := buf[hashSumLen:twoHashSumLen]
		string1 := buf[twoHashSumLen:twoHashSumLen]

		// appid
		// appkey
		// issubscribe
		// noncestr
		// openid
		// productid
		// timestamp
		string1 = append(string1, "appid="...)
		string1 = append(string1, req.AppId...)
		string1 = append(string1, "&appkey="...)
		string1 = append(string1, paySignKey...)
		string1 = append(string1, "&issubscribe="...)
		string1 = append(string1, IsSubscribeStr...)
		string1 = append(string1, "&noncestr="...)
		string1 = append(string1, req.NonceStr...)
		string1 = append(string1, "&openid="...)
		string1 = append(string1, req.OpenId...)
		string1 = append(string1, "&productid="...)
		string1 = append(string1, req.ProductId...)
		string1 = append(string1, "&timestamp="...)
		string1 = append(string1, TimeStampStr...)

		hashsumArray := sha1.Sum(string1)
		hex.Encode(signature, hashsumArray[:])

		// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
		if rslt := subtle.ConstantTimeCompare(signature, buf[:hashSumLen]); rslt == 1 {
			isValid = true
			return
		}
		return

	default:
		err = fmt.Errorf("not implement for %s sign method", req.SignMethod)
		return
	}
}

type BillResponse struct {
	AppId     string `xml:"AppId"`     // 必须, 公众帐号的 appid
	NonceStr  string `xml:"NonceStr"`  // 必须, 随机串
	TimeStamp int64  `xml:"TimeStamp"` // 必须, 时间戳

	Package string `xml:"Package"` // 必须,订单详情组合成的字符串, 4096个字符以内, see ../Bill.Package

	// 可以自己定义错误信息
	ErrCode int    `xml:"RetCode"`   // 可选, 0  表示正确
	ErrMsg  string `xml:"RetErrMsg"` // 可选, 错误信息, 要求 utf8 编码格式

	SignMethod string `xml:"SignMethod"` // 必须, 签名方式, 目前只支持 sha1
}

// 将 BillResponse 格式化成文档规定的格式, XML 格式.
//  @paySignKey: 公众号支付请求中用于加密的密钥 Key, 对应于支付场景中的 appKey
func (resp *BillResponse) MarshalToXML(paySignKey string) (xmlBytes []byte, err error) {
	var dst = struct {
		XMLName struct{} `xml:"xml" json:"-"`
		*BillResponse
		Signature string `xml:"AppSignature"`
	}{
		BillResponse: resp,
	}

	TimeStampStr := strconv.FormatInt(resp.TimeStamp, 10)
	ErrCodeStr := strconv.FormatInt(int64(resp.ErrCode), 10)

	const keysLen = len(`appid=&appkey=&noncestr=&package=&retcode=&reterrmsg=&timestamp=`)
	n := keysLen + len(resp.AppId) + len(paySignKey) + len(resp.NonceStr) +
		len(resp.Package) + len(ErrCodeStr) + len(resp.ErrMsg) + len(TimeStampStr)

	string1 := make([]byte, 0, n)

	// 字典序
	// appid
	// appkey
	// noncestr
	// package
	// retcode
	// reterrmsg
	// timestamp
	string1 = append(string1, "appid="...)
	string1 = append(string1, resp.AppId...)
	string1 = append(string1, "&appkey="...)
	string1 = append(string1, paySignKey...)
	string1 = append(string1, "&noncestr="...)
	string1 = append(string1, resp.NonceStr...)
	string1 = append(string1, "&package="...)
	string1 = append(string1, resp.Package...)
	string1 = append(string1, "&retcode="...)
	string1 = append(string1, ErrCodeStr...)
	string1 = append(string1, "&reterrmsg="...)
	string1 = append(string1, resp.ErrMsg...)
	string1 = append(string1, "&timestamp="...)
	string1 = append(string1, TimeStampStr...)

	switch resp.SignMethod {
	case RESPONSE_SIGN_METHOD_SHA1:
		hashsumArray := sha1.Sum(string1)
		dst.Signature = hex.EncodeToString(hashsumArray[:])
	default:
		err = fmt.Errorf("not implement for %s sign method", resp.SignMethod)
		return
	}

	xmlBytes, err = xml.Marshal(dst)
	return
}
