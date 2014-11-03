// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/chanxuehong/wechat/util"
	"strconv"
	"time"
)

// 订单详情, 微信根据这个信息生成订单.
// js api 和 native api 都需要这个, 就是那个 订单详情(package) 字符串, see PayPackage.Package
type PayPackage struct {
	BankType     string `json:"bank_type"`               // 必须, 银行通道类型, 固定为 "WX"
	Body         string `json:"body"`                    // 必须, 商品描述, 128字节以内
	Attach       string `json:"attach,omitempty"`        // 可选, 附加数据, 128字节以内
	PartnerId    string `json:"partner"`                 // 必须, 注册时分配的财付通商户号 partnerId
	OutTradeNo   string `json:"out_trade_no"`            // 必须, 商户系统内部订单号, 32个字符内, 可包含字母, *** 确保在商户系统中唯一 ***
	TotalFee     string `json:"total_fee"`               // 必须, 订单总金额, 单位为分
	TransportFee string `json:"transport_fee,omitempty"` // 可选, 物流费用, 单位为分; 如果有值, 必须保证 TransportFee + ProductFee == TotalFee
	ProductFee   string `json:"product_fee,omitempty"`   // 可选, 商品费用, 单位为分; 如果有值, 必须保证 TransportFee + ProductFee == TotalFee
	ProductTag   string `json:"goods_tag,omitempty"`     // 可选, 商品标记, 优惠卷时可能用到
	FeeType      string `json:"fee_type"`                // 必须, 取值: 1(人民币); 目前暂只支持 1
	NotifyURL    string `json:"notify_url"`              // 必须, 在支付完成后, 接收微信通知支付结果的 URL, 需要给出绝对路径, 255个字符内
	BillCreateIP string `json:"spbill_create_ip"`        // 必须, 订单生成的机器IP(指用户浏览器端IP, 不是商户服务器IP, 格式为IPV4), 15个字节内
	TimeStart    string `json:"time_start,omitempty"`    // 可选, 订单生成时间, 该时间取自商户服务器
	TimeExpire   string `json:"time_expire,omitempty"`   // 可选, 订单失效时间, 该时间取自商户服务器
	Charset      string `json:"input_charset"`           // 必须, 参数字符编码, 取值范围: "GBK", "UTF-8"
}

// setter
func (this *PayPackage) SetTotalFee(n int) {
	this.TotalFee = strconv.FormatInt(int64(n), 10)
}
func (this *PayPackage) SetTransportFee(n int) {
	this.TransportFee = strconv.FormatInt(int64(n), 10)
}
func (this *PayPackage) SetProductFee(n int) {
	this.ProductFee = strconv.FormatInt(int64(n), 10)
}
func (this *PayPackage) SetFeeType(n int) {
	this.FeeType = strconv.FormatInt(int64(n), 10)
}
func (this *PayPackage) SetTimeStart(t time.Time) {
	this.TimeStart = util.FormatTime(t)
}
func (this *PayPackage) SetTimeExpire(t time.Time) {
	this.TimeExpire = util.FormatTime(t)
}

// 将 PayPackage 打包成 订单详情(package)字符串 需要的格式.
//  partnerKey: 财付通商户权限密钥 Key
//
//  NOTE: 这个函数不对 this *PayPackage 的字段做有效性检查, 请确保设置正确
func (this *PayPackage) Package(partnerKey string) (package_ []byte) {
	ks := make([]string, 0, 16)  // 包含不为空值的字段的 key, 字典序, 目前只有 15 个字段
	vs1 := make([]string, 0, 16) // 包含不为空值的字段的 value,
	vs2 := make([]string, 0, 16) // 包含不为空值的字段的 value, 经过了 URLEscape

	// 字典序
	// attach
	// bank_type
	// body
	// fee_type
	// goods_tag
	// input_charset
	// notify_url
	// out_trade_no
	// partner
	// product_fee
	// spbill_create_ip
	// time_expire
	// time_start
	// total_fee
	// transport_fee
	if len(this.Attach) > 0 {
		ks = append(ks, "&attach=")
		vs1 = append(vs1, this.Attach)
		vs2 = append(vs2, util.URLEscape(this.Attach))
	}
	if len(this.BankType) > 0 {
		ks = append(ks, "&bank_type=")
		vs1 = append(vs1, this.BankType)
		vs2 = append(vs2, util.URLEscape(this.BankType))
	}
	if len(this.Body) > 0 {
		ks = append(ks, "&body=")
		vs1 = append(vs1, this.Body)
		vs2 = append(vs2, util.URLEscape(this.Body))
	}
	if len(this.FeeType) > 0 {
		ks = append(ks, "&fee_type=")
		vs1 = append(vs1, this.FeeType)
		vs2 = append(vs2, util.URLEscape(this.FeeType))
	}
	if len(this.ProductTag) > 0 {
		ks = append(ks, "&goods_tag=")
		vs1 = append(vs1, this.ProductTag)
		vs2 = append(vs2, util.URLEscape(this.ProductTag))
	}
	if len(this.Charset) > 0 {
		ks = append(ks, "&input_charset=")
		vs1 = append(vs1, this.Charset)
		vs2 = append(vs2, util.URLEscape(this.Charset))
	}
	if len(this.NotifyURL) > 0 {
		ks = append(ks, "&notify_url=")
		vs1 = append(vs1, this.NotifyURL)
		vs2 = append(vs2, util.URLEscape(this.NotifyURL))
	}
	if len(this.OutTradeNo) > 0 {
		ks = append(ks, "&out_trade_no=")
		vs1 = append(vs1, this.OutTradeNo)
		vs2 = append(vs2, util.URLEscape(this.OutTradeNo))
	}
	if len(this.PartnerId) > 0 {
		ks = append(ks, "&partner=")
		vs1 = append(vs1, this.PartnerId)
		vs2 = append(vs2, util.URLEscape(this.PartnerId))
	}
	if len(this.ProductFee) > 0 {
		ks = append(ks, "&product_fee=")
		vs1 = append(vs1, this.ProductFee)
		vs2 = append(vs2, util.URLEscape(this.ProductFee))
	}
	if len(this.BillCreateIP) > 0 {
		ks = append(ks, "&spbill_create_ip=")
		vs1 = append(vs1, this.BillCreateIP)
		vs2 = append(vs2, util.URLEscape(this.BillCreateIP))
	}
	if len(this.TimeExpire) > 0 {
		ks = append(ks, "&time_expire=")
		vs1 = append(vs1, this.TimeExpire)
		vs2 = append(vs2, util.URLEscape(this.TimeExpire))
	}
	if len(this.TimeStart) > 0 {
		ks = append(ks, "&time_start=")
		vs1 = append(vs1, this.TimeStart)
		vs2 = append(vs2, util.URLEscape(this.TimeStart))
	}
	if len(this.TotalFee) > 0 {
		ks = append(ks, "&total_fee=")
		vs1 = append(vs1, this.TotalFee)
		vs2 = append(vs2, util.URLEscape(this.TotalFee))
	}
	if len(this.TransportFee) > 0 {
		ks = append(ks, "&transport_fee=")
		vs1 = append(vs1, this.TransportFee)
		vs2 = append(vs2, util.URLEscape(this.TransportFee))
	}

	if len(ks) > 0 {
		ks[0] = ks[0][1:] // len(ks[0]) > 0, 去掉 ks[0] 的首字母 &
	}

	ksTotalLen := 0
	for _, str := range ks {
		ksTotalLen += len(str)
	}
	vs1TotalLen := 0
	for _, str := range vs1 {
		vs1TotalLen += len(str)
	}
	vs2TotalLen := 0
	for _, str := range vs2 {
		vs2TotalLen += len(str)
	}

	var n1, n2 int
	if len(ks) > 0 {
		n1 = ksTotalLen + vs1TotalLen + 5 + len(partnerKey) // &key=partnerKey
		n2 = ksTotalLen + vs2TotalLen + 6 + md5.Size*2      // &sign=signature, md5
	} else {
		n1 = 4 + len(partnerKey) // key=partnerKey
		n2 = 5 + md5.Size*2      // sign=signature, md5
	}

	buf := make([]byte, n1+n2)
	string1 := buf[0:0:n1]
	string2 := buf[n1:n1]
	package_ = buf[n1:]
	signature := package_[len(package_)-md5.Size*2:]

	for i := 0; i < len(ks); i++ {
		string1 = append(string1, ks[i]...)
		string1 = append(string1, vs1[i]...)

		string2 = append(string2, ks[i]...)
		string2 = append(string2, vs2[i]...)
	}

	if len(ks) > 0 {
		string1 = append(string1, "&key="...)
		string2 = append(string2, "&sign="...)

	} else {
		string1 = append(string1, "key="...)
		string2 = append(string2, "sign="...)
	}
	string1 = append(string1, partnerKey...)
	md5sum := md5.Sum(string1)
	hex.Encode(signature, md5sum[:])
	copy(signature, bytes.ToUpper(signature))

	return
}
