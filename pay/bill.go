// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

// 订单支付的详细清单.
// js api 和 native api 都需要这个, 就是那个 订单详情(package) 字符串, see Bill.Package.
type Bill struct {
	BankType     string    // 必须, 银行通道类型, 固定为 "WX"
	Body         string    // 必须, 商品描述, 128字节以内
	Attach       string    // 可选, 附加数据, 128字节以内
	PartnerId    string    // 必须, 注册时分配的财付通商户号 partnerId
	OutTradeNo   string    // 必须, 商户系统内部订单号, 32个字符内, 可包含字母, *** 确保在商户系统中唯一 ***
	TotalFee     int       // 必须, 订单总金额, 单位为分
	TransportFee int       // 可选, 物流费用, 单位为分; 如果有值, 必须保证 TransportFee + ProductFee == TotalFee
	ProductFee   int       // 可选, 商品费用, 单位为分; 如果有值, 必须保证 TransportFee + ProductFee == TotalFee
	ProductTag   string    // 可选, 商品标记, 优惠卷时可能用到
	FeeType      int       // 必须, 取值: 1(人民币); 目前暂只支持 1
	NotifyURL    string    // 必须, 在支付完成后, 接收微信通知支付结果的 URL, 需要给出绝对路径, 255个字符内
	CreateIP     string    // 必须, 订单生成的机器IP(指用户浏览器端IP, 不是商户服务器IP, 格式为IPV4), 15个字节内
	TimeStart    time.Time // 可选, 订单生成时间, 该时间取自商户服务器
	TimeExpire   time.Time // 可选, 订单失效时间, 该时间取自商户服务器
	Charset      string    // 必须, 参数字符编码, 取值范围: "GBK","UTF-8", 留空默认为 "UTF-8"
}

// 检查 Bill 设置的是否合法, 合法返回 nil, 否则返回错误信息
func (this *Bill) Check() (err error) {
	const (
		BodyLimit       = 128
		AttachLimit     = 128
		OutTradeNoLimit = 32
		NotifyURLLimit  = 255
		CreateIPLimit   = 15
	)

	switch this.BankType {
	case BANK_TYPE_WX:
	case "":
		err = errors.New("没有设置 BankType")
		return
	default:
		err = errors.New("不规范的 BankType")
		return
	}

	if n := len(this.Body); n == 0 {
		err = errors.New("没有设置 Body")
		return
	} else if n > BodyLimit {
		err = fmt.Errorf("Body 长度不能超过 %d", BodyLimit)
		return
	}

	if len(this.Attach) > AttachLimit {
		err = fmt.Errorf("Attach 长度不能超过 %d", AttachLimit)
		return
	}

	if this.PartnerId == "" {
		err = errors.New("没有设置 PartnerId")
		return
	}

	if n := len(this.OutTradeNo); n == 0 {
		err = errors.New("没有设置 OutTradeNo")
		return
	} else if n > OutTradeNoLimit {
		err = fmt.Errorf("OutTradeNo 长度不能超过 %d", OutTradeNoLimit)
		return
	}

	// TotalFee 可以为任何值, 所以没有必要检查

	if this.TransportFee != 0 && this.TransportFee+this.ProductFee != this.TotalFee {
		err = errors.New("你设置了 TransportFee, 但是 TransportFee+ProductFee != TotalFee")
		return
	}
	if this.ProductFee != 0 && this.TransportFee+this.ProductFee != this.TotalFee {
		err = errors.New("你设置了 ProductFee, 但是 TransportFee+ProductFee != TotalFee")
		return
	}

	switch this.FeeType {
	case FEE_TYPE_RMB:
	default:
		err = errors.New("不规范的 FeeType")
		return
	}

	if n := len(this.NotifyURL); n == 0 {
		err = errors.New("没有设置 NotifyURL")
		return
	} else if n > NotifyURLLimit {
		err = fmt.Errorf("NotifyURL 长度不能超过 %d", NotifyURLLimit)
		return
	}

	if n := len(this.CreateIP); n == 0 {
		err = errors.New("没有设置 CreateIP")
		return
	} else if n > CreateIPLimit {
		err = fmt.Errorf("CreateIP 长度不能超过 %d", CreateIPLimit)
		return
	} else {
		if ip := net.ParseIP(this.CreateIP); ip == nil {
			err = fmt.Errorf("CreateIP(%s) 不是一个有效的 IP", this.CreateIP)
			return
		}
	}

	switch this.Charset {
	case "", BILL_CHARSET_UTF8, BILL_CHARSET_GBK:
	default:
		err = errors.New("不规范的 Charset")
		return
	}

	return
}

// 将 Bill 打包成 订单详情(package)字符串 需要的格式.
//  @partnerKey: 财付通商户权限密钥 Key
//  NOTE: 这个函数不对 this *Bill 的字段做有效性检查, 你可以选择调用 Bill.Check()
func (this *Bill) Package(partnerKey string) (bs []byte) {
	if this.Charset == "" {
		this.Charset = BILL_CHARSET_UTF8
	}

	var (
		// 非字符串的字段都要转换成字符串
		TotalFeeStr     string = strconv.FormatInt(int64(this.TotalFee), 10)
		FeeTypeStr      string = strconv.FormatInt(int64(this.FeeType), 10)
		TimeStartStr    string
		TimeExpireStr   string
		TransportFeeStr string
		ProductFeeStr   string
	)
	var (
		// 最终结果所有字段的 value 都要经过 urlencode
		//  NOTE: 虽然有些字段, 比如 BankType, 无需 urlencode, 但是安全着想都做了 urlencode
		BankTypeURLEscapedStr   string = URLEscape(this.BankType)
		BodyURLEscapedStr       string = URLEscape(this.Body)
		PartnerIdURLEscapedStr  string = URLEscape(this.PartnerId)
		OutTradeNoURLEscapedStr string = URLEscape(this.OutTradeNo)
		NotifyURLEscapedStr     string = URLEscape(this.NotifyURL)
		CreateIPURLEscapedStr   string = URLEscape(this.CreateIP)
		CharsetURLEscapedStr    string = URLEscape(this.Charset)
		AttachURLEscapedStr     string
		ProductTagURLEscapedStr string

		// 下面这些字段的字符串都是在本 method 内生成, 肯定不需要 urlencode
		//
		// TotalFeeURLEscapedStr     string
		// FeeTypeURLEscapedStr      string
		// TimeStartURLEscapedStr    string
		// TimeExpireURLEscapedStr   string
		// TransportFeeURLEscapedStr string
		// ProductFeeURLEscapedStr   string
	)

	// 确定文档中 string1 和 string2 的长度
	// n1 == len(string1 + "&key=" + partnerKey)
	// n2 == len(string2 + "&sign=" + md5sum)

	const keysLenMin = len(`bank_type=&body=&fee_type=&input_charset=&notify_url=&out_trade_no=&partner=&spbill_create_ip=&total_fee=`)

	n1 := keysLenMin + len(this.BankType) + len(this.Body) + len(FeeTypeStr) + len(this.Charset) +
		len(this.NotifyURL) + len(this.OutTradeNo) + len(this.PartnerId) + len(this.CreateIP) + len(TotalFeeStr)
	n2 := keysLenMin + len(BankTypeURLEscapedStr) + len(BodyURLEscapedStr) + len(FeeTypeStr) +
		len(CharsetURLEscapedStr) + len(NotifyURLEscapedStr) + len(OutTradeNoURLEscapedStr) + len(PartnerIdURLEscapedStr) +
		len(CreateIPURLEscapedStr) + len(TotalFeeStr)

	if this.Attach != "" {
		AttachURLEscapedStr = URLEscape(this.Attach)

		// &attach=
		n1 += 8 + len(this.Attach)
		n2 += 8 + len(AttachURLEscapedStr)
	}
	if !this.TimeStart.IsZero() {
		TimeStartStr = FormatTime(this.TimeStart)

		// &time_start=
		n1 += 12 + len(TimeStartStr)
		n2 += 12 + len(TimeStartStr)
	}
	if !this.TimeExpire.IsZero() {
		TimeExpireStr = FormatTime(this.TimeExpire)

		// &time_expire=
		n1 += 13 + len(TimeExpireStr)
		n2 += 13 + len(TimeExpireStr)
	}
	if this.TransportFee != 0 {
		TransportFeeStr = strconv.FormatInt(int64(this.TransportFee), 10)

		// &transport_fee=
		n1 += 15 + len(TransportFeeStr)
		n2 += 15 + len(TransportFeeStr)
	}
	if this.ProductFee != 0 {
		ProductFeeStr = strconv.FormatInt(int64(this.ProductFee), 10)

		// &product_fee=
		n1 += 13 + len(ProductFeeStr)
		n2 += 13 + len(ProductFeeStr)
	}
	if this.ProductTag != "" {
		ProductTagURLEscapedStr = URLEscape(this.ProductTag)

		// &goods_tag=
		n1 += 11 + len(this.ProductTag)
		n2 += 11 + len(ProductTagURLEscapedStr)
	}

	// &key=
	n1 += 5 + len(partnerKey)
	// &sign=signature
	n2 += 6 + 32

	string1 := make([]byte, 0, n1)
	bs = make([]byte, n2)
	string2 := bs[:0]
	signature := bs[n2-32:] // md5sum

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

	hasAppend := false

	if this.Attach != "" {
		hasAppend = true

		string1 = append(string1, "attach="...)
		string1 = append(string1, this.Attach...)
		string2 = append(string2, "attach="...)
		string2 = append(string2, AttachURLEscapedStr...)
	}

	if hasAppend {
		string1 = append(string1, "&bank_type="...)
		string1 = append(string1, this.BankType...)
		string2 = append(string2, "&bank_type="...)
		string2 = append(string2, BankTypeURLEscapedStr...)
	} else {
		string1 = append(string1, "bank_type="...)
		string1 = append(string1, this.BankType...)
		string2 = append(string2, "bank_type="...)
		string2 = append(string2, BankTypeURLEscapedStr...)
	}

	string1 = append(string1, "&body="...)
	string1 = append(string1, this.Body...)
	string2 = append(string2, "&body="...)
	string2 = append(string2, BodyURLEscapedStr...)

	string1 = append(string1, "&fee_type="...)
	string1 = append(string1, FeeTypeStr...)
	string2 = append(string2, "&fee_type="...)
	string2 = append(string2, FeeTypeStr...)

	if this.ProductTag != "" {
		string1 = append(string1, "&goods_tag="...)
		string1 = append(string1, this.ProductTag...)
		string2 = append(string2, "&goods_tag="...)
		string2 = append(string2, ProductTagURLEscapedStr...)
	}

	string1 = append(string1, "&input_charset="...)
	string1 = append(string1, this.Charset...)
	string2 = append(string2, "&input_charset="...)
	string2 = append(string2, CharsetURLEscapedStr...)

	string1 = append(string1, "&notify_url="...)
	string1 = append(string1, this.NotifyURL...)
	string2 = append(string2, "&notify_url="...)
	string2 = append(string2, NotifyURLEscapedStr...)

	string1 = append(string1, "&out_trade_no="...)
	string1 = append(string1, this.OutTradeNo...)
	string2 = append(string2, "&out_trade_no="...)
	string2 = append(string2, OutTradeNoURLEscapedStr...)

	string1 = append(string1, "&partner="...)
	string1 = append(string1, this.PartnerId...)
	string2 = append(string2, "&partner="...)
	string2 = append(string2, PartnerIdURLEscapedStr...)

	if this.ProductFee != 0 {
		string1 = append(string1, "&product_fee="...)
		string1 = append(string1, ProductFeeStr...)
		string2 = append(string2, "&product_fee="...)
		string2 = append(string2, ProductFeeStr...)
	}

	string1 = append(string1, "&spbill_create_ip="...)
	string1 = append(string1, this.CreateIP...)
	string2 = append(string2, "&spbill_create_ip="...)
	string2 = append(string2, CreateIPURLEscapedStr...)

	if !this.TimeExpire.IsZero() {
		string1 = append(string1, "&time_expire="...)
		string1 = append(string1, TimeExpireStr...)
		string2 = append(string2, "&time_expire="...)
		string2 = append(string2, TimeExpireStr...)
	}

	if !this.TimeStart.IsZero() {
		string1 = append(string1, "&time_start="...)
		string1 = append(string1, TimeStartStr...)
		string2 = append(string2, "&time_start="...)
		string2 = append(string2, TimeStartStr...)
	}

	string1 = append(string1, "&total_fee="...)
	string1 = append(string1, TotalFeeStr...)
	string2 = append(string2, "&total_fee="...)
	string2 = append(string2, TotalFeeStr...)

	if this.TransportFee != 0 {
		string1 = append(string1, "&transport_fee="...)
		string1 = append(string1, TransportFeeStr...)
		string2 = append(string2, "&transport_fee="...)
		string2 = append(string2, TransportFeeStr...)
	}

	string1 = append(string1, "&key="...)
	string1 = append(string1, partnerKey...)

	string2 = append(string2, "&sign="...)
	md5sumArray := md5.Sum(string1)
	hex.Encode(signature, md5sumArray[:])
	copy(signature, bytes.ToUpper(signature))

	return
}
