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
	"strconv"
)

type Request struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	OpenId       string   `xml:"OpenId"`       // 点击链接准备购买商品的用户 OpenId
	AppId        string   `xml:"AppId"`        // 公众帐号的 appid
	IsSubscribe  int      `xml:"IsSubscribe"`  // 标记用户是否订阅该公众帐号，1为关注，0为未关注
	ProductId    string   `xml:"ProductId"`    // 第三方的商品	ID号
	TimeStamp    int64    `xml:"TimeStamp"`    // 时间戳
	NonceStr     string   `xml:"NonceStr"`     // 随机串
	AppSignature string   `xml:"AppSignature"` // 参数的加密签名
	SignMethod   string   `xml:"SignMethod"`   // 签名方式，目前只支持“SHA1”，该字段不参与签名
}

func (req *Request) CheckSignature(paySignKey string) bool {
	const hashsumLen = 40 // sha1

	if len(req.AppSignature) != hashsumLen {
		return false
	}

	IsSubscribeStr := strconv.FormatInt(int64(req.IsSubscribe), 10)
	TimeStampStr := strconv.FormatInt(req.TimeStamp, 10)

	// hashsum, len(`appid=&appkey=&issubscribe=&noncestr=&openid=&productid=&timestamp=`) == 67
	n := hashsumLen + 67 + len(req.AppId) + len(paySignKey) + len(IsSubscribeStr) +
		len(req.NonceStr) + len(req.OpenId) + len(req.ProductId) + len(TimeStampStr)

	buf := make([]byte, hashsumLen, n)

	// appid
	// appkey
	// issubscribe
	// noncestr
	// openid
	// productid
	// timestamp
	buf = append(buf, "appid="...)
	buf = append(buf, req.AppId...)
	buf = append(buf, "&appkey="...)
	buf = append(buf, paySignKey...)
	buf = append(buf, "&issubscribe="...)
	buf = append(buf, IsSubscribeStr...)
	buf = append(buf, "&noncestr="...)
	buf = append(buf, req.NonceStr...)
	buf = append(buf, "&openid="...)
	buf = append(buf, req.OpenId...)
	buf = append(buf, "&productid="...)
	buf = append(buf, req.ProductId...)
	buf = append(buf, "&timestamp="...)
	buf = append(buf, TimeStampStr...)

	hashsumArray := sha1.Sum(buf[hashsumLen:])
	hashsumHexBytes := buf[:hashsumLen]
	hex.Encode(hashsumHexBytes, hashsumArray[:])

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	if rslt := subtle.ConstantTimeCompare(hashsumHexBytes, []byte(req.AppSignature)); rslt == 1 {
		return true
	}
	return false
}

type Response struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	AppId        string   `xml:"AppId"`
	Package      string   `xml:"Package"`
	TimeStamp    int64    `xml:"TimeStamp"`
	NonceStr     string   `xml:"NonceStr"`
	RetCode      int      `xml:"RetCode"`
	RetErrMsg    string   `xml:"RetErrMsg"`
	AppSignature string   `xml:"AppSignature"`
	SignMethod   string   `xml:"SignMethod"`
}

func (resp *Response) SetAppSignature(paySignKey string) {
	RetCodeStr := strconv.FormatInt(int64(resp.RetCode), 10)
	TimeStampStr := strconv.FormatInt(resp.TimeStamp, 10)

	// len(`appid=&appkey=&noncestr=&package=&retcode=&reterrmsg=&timestamp=`) == 64
	n := 64 + len(resp.AppId) + len(paySignKey) + len(resp.NonceStr) +
		len(resp.Package) + len(RetCodeStr) + len(resp.RetErrMsg) + len(TimeStampStr)

	buf := make([]byte, 0, n)

	// appid
	// appkey
	// noncestr
	// package
	// retcode
	// reterrmsg
	// timestamp
	buf = append(buf, "appid="...)
	buf = append(buf, resp.AppId...)
	buf = append(buf, "&appkey="...)
	buf = append(buf, paySignKey...)
	buf = append(buf, "&noncestr="...)
	buf = append(buf, resp.NonceStr...)
	buf = append(buf, "&package="...)
	buf = append(buf, resp.Package...)
	buf = append(buf, "&retcode="...)
	buf = append(buf, RetCodeStr...)
	buf = append(buf, "&reterrmsg="...)
	buf = append(buf, resp.RetErrMsg...)
	buf = append(buf, "&timestamp="...)
	buf = append(buf, TimeStampStr...)

	hashsumArray := sha1.Sum(buf)
	resp.AppSignature = hex.EncodeToString(hashsumArray[:])
}

func (resp *Response) MarshalToXML() []byte {
	bs, _ := xml.Marshal(resp)
	return bs
}
