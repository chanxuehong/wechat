// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package js

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// 订单详情.
// jsapi getBrandWCPayRequest 接口的 package 参数的原型, see OrderInfo.Package()
type OrderInfo struct {
	BankType     string    // 必须, 银行通道类型, 固定为 "WX"
	Body         string    // 必须, 商品描述, 128字节以内
	Attach       string    // 可选, 附加数据, 原样返回, 128字节以内
	PartnerId    string    // 必须, 注册时分配的财付通商户号 partnerId
	OutTradeNo   string    // 必须, 商户系统内部订单号, 32个字符内, 可包含字母, *确保在商户系统中唯一*
	TotalFee     int       // 必须, 订单总金额, 单位为分
	FeeType      int       // 必须, 取值：1（人民币）, 目前暂只支持 1
	NotifyURL    string    // 必须, 在支付完成后, 接收微信通知支付结果的URL, 需要给出绝对路径, 255个字符内
	CreateIP     string    // 必须, 订单生成的机器IP（指用户浏览器端IP, 不是商户服务器IP, 格式为IPV4）, 15个字节内
	TimeStart    time.Time // 可选, 订单生成时间, 格式为 yyyyMMDDHHmmss, GMT+8, 该时间取自商户服务器
	TimeExpire   time.Time // 可选, 订单失效时间, 格式为 yyyyMMDDHHmmss, GMT+8, 该时间取自商户服务器
	TransportFee int       // 可选, 物流费用, 单位为分, 如果有值, 必须保证 TransportFee + ProductFee == TotalFee
	ProductFee   int       // 可选, 商品费用, 单位为分, 如果有值, 必须保证 TransportFee + ProductFee == TotalFee
	ProductTag   string    // 可选, 商品标记, 优惠卷时可能用到
	Charset      string    // 必须, 参数字符编码, 取值范围: "GBK","UTF-8"
}

// 检查 info *OrderInfo 是否合法, 合法返回 nil, 否则返回错误信息
func (info *OrderInfo) Check() (err error) {
	const (
		BodyLimit       = 128
		AttachLimit     = 128
		OutTradeNoLimit = 32
		NotifyURLLimit  = 255
		CreateIPLimit   = 15
	)

	if info.BankType != ORDER_INFO_BANK_TYPE_WX {
		err = errors.New("请正确设置 BankType")
		return
	}
	if n := len(info.Body); n == 0 {
		err = errors.New("请设置 Body")
		return
	} else if n > BodyLimit {
		err = fmt.Errorf("Body 长度不能超过 %d", BodyLimit)
		return
	}
	if len(info.Attach) > AttachLimit {
		err = fmt.Errorf("Attach 长度不能超过 %d", AttachLimit)
		return
	}
	if info.PartnerId == "" {
		err = errors.New("请设置 PartnerId")
		return
	}
	if n := len(info.OutTradeNo); n == 0 {
		err = errors.New("请设置 OutTradeNo")
		return
	} else if n > OutTradeNoLimit {
		err = fmt.Errorf("OutTradeNo 长度不能超过 %d", OutTradeNoLimit)
		return
	}
	if info.FeeType != ORDER_INFO_FEE_TYPE_RMB {
		err = errors.New("请正确设置 FeeType")
		return
	}
	if n := len(info.NotifyURL); n == 0 {
		err = errors.New("请设置 NotifyURL")
		return
	} else if n > NotifyURLLimit {
		err = fmt.Errorf("NotifyURL 长度不能超过 %d", NotifyURLLimit)
		return
	}
	if n := len(info.CreateIP); n == 0 {
		err = errors.New("请设置 CreateIP")
		return
	} else if n > CreateIPLimit {
		err = fmt.Errorf("CreateIP 长度不能超过 %d", CreateIPLimit)
		return
	}
	if info.TimeExpire.Sub(info.TimeStart) < 0 {
		err = errors.New("失效时间不能小于生成时间")
		return
	}
	if info.TransportFee != 0 && info.TransportFee+info.ProductFee != info.TotalFee {
		err = errors.New("你设置了 TransportFee, 但是 TransportFee+ProductFee != TotalFee")
		return
	}
	if info.ProductFee != 0 && info.TransportFee+info.ProductFee != info.TotalFee {
		err = errors.New("你设置了 ProductFee, 但是 TransportFee+ProductFee != TotalFee")
		return
	}
	if info.Charset != ORDER_INFO_CHARSET_UTF8 &&
		info.Charset != ORDER_INFO_CHARSET_GBK {

		err = errors.New("请正确设置 Charset")
		return
	}
	return
}

// 格式化时间到 yyyyMMDDHHmmss
func formatTime(t time.Time) string {
	const layout = "20060102150405"
	return t.Format(layout)
}

// 将 info *OrderInfo 打包成 jsapi getBrandWCPayRequest 接口的 package 参数需要的格式.
//  NOTE: 这个函数不对 info *OrderInfo 的字段做有效性检查, 你可以选择调用 OrderInfo.Check()
//  @partnerKey: 财付通商户权限密钥 Key
func (info *OrderInfo) Package(partnerKey string) (bs []byte) {
	var (
		TotalFeeStr     string = strconv.FormatInt(int64(info.TotalFee), 10)
		FeeTypeStr      string = strconv.FormatInt(int64(info.FeeType), 10)
		TimeStartStr    string
		TimeExpireStr   string
		TransportFeeStr string
		ProductFeeStr   string
	)
	var (
		BankTypeURLEscapedStr   string = URLEscape(info.BankType)
		BodyURLEscapedStr       string = URLEscape(info.Body)
		AttachURLEscapedStr     string
		PartnerIdURLEscapedStr  string = URLEscape(info.PartnerId)
		OutTradeNoURLEscapedStr string = URLEscape(info.OutTradeNo)
		NotifyURLEscapedStr     string = URLEscape(info.NotifyURL)
		CreateIPURLEscapedStr   string = URLEscape(info.CreateIP)
		ProductTagURLEscapedStr string
		CharsetURLEscapedStr    string = URLEscape(info.Charset)
		// TotalFeeURLEscapedStr     string
		// FeeTypeURLEscapedStr      string
		// TimeStartURLEscapedStr    string
		// TimeExpireURLEscapedStr   string
		// TransportFeeURLEscapedStr string
		// ProductFeeURLEscapedStr   string
	)

	// md5sum == 32, len(`bank_type=&body=&partner=&out_trade_no=&total_fee=&fee_type=&notify_url=&spbill_create_ip=&input_charset=`) == 105
	n1 := 32 + 105 + len(info.BankType) + len(info.Body) + len(info.PartnerId) + len(info.OutTradeNo) +
		len(TotalFeeStr) + len(FeeTypeStr) + len(info.NotifyURL) + len(info.CreateIP) + len(info.Charset)
	// len(`bank_type=&body=&partner=&out_trade_no=&total_fee=&fee_type=&notify_url=&spbill_create_ip=&input_charset=`) == 105
	n2 := 105 + len(BankTypeURLEscapedStr) + len(BodyURLEscapedStr) + len(PartnerIdURLEscapedStr) +
		len(OutTradeNoURLEscapedStr) + len(TotalFeeStr) + len(FeeTypeStr) + len(NotifyURLEscapedStr) +
		len(CreateIPURLEscapedStr) + len(CharsetURLEscapedStr)

	if info.Attach != "" {
		AttachURLEscapedStr = URLEscape(info.Attach)

		// &attach=
		n1 += 8 + len(info.Attach)
		n2 += 8 + len(AttachURLEscapedStr)
	}
	if !info.TimeStart.IsZero() {
		TimeStartStr = formatTime(info.TimeStart)

		// &time_start=
		n1 += 12 + len(TimeStartStr)
		n2 += 12 + len(TimeStartStr)
	}
	if !info.TimeExpire.IsZero() {
		TimeExpireStr = formatTime(info.TimeExpire)

		// &time_expire=
		n1 += 13 + len(TimeExpireStr)
		n2 += 13 + len(TimeExpireStr)
	}
	if info.TransportFee != 0 {
		TransportFeeStr = strconv.FormatInt(int64(info.TransportFee), 10)

		// &transport_fee=
		n1 += 15 + len(TransportFeeStr)
		n2 += 15 + len(TransportFeeStr)
	}
	if info.ProductFee != 0 {
		ProductFeeStr = strconv.FormatInt(int64(info.ProductFee), 10)

		// &product_fee=
		n1 += 13 + len(ProductFeeStr)
		n2 += 13 + len(ProductFeeStr)
	}
	if info.ProductTag != "" {
		ProductTagURLEscapedStr = URLEscape(info.ProductTag)

		// &goods_tag=
		n1 += 11 + len(info.ProductTag)
		n2 += 11 + len(ProductTagURLEscapedStr)
	}

	// &key=
	n1 += 5 + len(partnerKey)
	// &sign=md5sum
	n2 += 6 + 32

	buf := make([]byte, 32, n1)
	bs = make([]byte, 0, n2)

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

	if info.Attach != "" {
		hasAppend = true

		buf = append(buf, "attach="...)
		buf = append(buf, info.Attach...)
		bs = append(bs, "attach="...)
		bs = append(bs, AttachURLEscapedStr...)
	}

	if hasAppend {
		buf = append(buf, "&bank_type="...)
		buf = append(buf, info.BankType...)
		bs = append(bs, "&bank_type="...)
		bs = append(bs, BankTypeURLEscapedStr...)
	} else {
		buf = append(buf, "bank_type="...)
		buf = append(buf, info.BankType...)
		bs = append(bs, "bank_type="...)
		bs = append(bs, BankTypeURLEscapedStr...)
	}

	buf = append(buf, "&body="...)
	buf = append(buf, info.Body...)
	bs = append(bs, "&body="...)
	bs = append(bs, BodyURLEscapedStr...)

	buf = append(buf, "&fee_type="...)
	buf = append(buf, FeeTypeStr...)
	bs = append(bs, "&fee_type="...)
	bs = append(bs, FeeTypeStr...)

	if info.ProductTag != "" {
		buf = append(buf, "&goods_tag="...)
		buf = append(buf, info.ProductTag...)
		bs = append(bs, "&goods_tag="...)
		bs = append(bs, ProductTagURLEscapedStr...)
	}

	buf = append(buf, "&input_charset="...)
	buf = append(buf, info.Charset...)
	bs = append(bs, "&input_charset="...)
	bs = append(bs, CharsetURLEscapedStr...)

	buf = append(buf, "&notify_url="...)
	buf = append(buf, info.NotifyURL...)
	bs = append(bs, "&notify_url="...)
	bs = append(bs, NotifyURLEscapedStr...)

	buf = append(buf, "&out_trade_no="...)
	buf = append(buf, info.OutTradeNo...)
	bs = append(bs, "&out_trade_no="...)
	bs = append(bs, OutTradeNoURLEscapedStr...)

	buf = append(buf, "&partner="...)
	buf = append(buf, info.PartnerId...)
	bs = append(bs, "&partner="...)
	bs = append(bs, PartnerIdURLEscapedStr...)

	if info.ProductFee != 0 {
		buf = append(buf, "&product_fee="...)
		buf = append(buf, ProductFeeStr...)
		bs = append(bs, "&product_fee="...)
		bs = append(bs, ProductFeeStr...)
	}

	buf = append(buf, "&spbill_create_ip="...)
	buf = append(buf, info.CreateIP...)
	bs = append(bs, "&spbill_create_ip="...)
	bs = append(bs, CreateIPURLEscapedStr...)

	if !info.TimeExpire.IsZero() {
		buf = append(buf, "&time_expire="...)
		buf = append(buf, TimeExpireStr...)
		bs = append(bs, "&time_expire="...)
		bs = append(bs, TimeExpireStr...)
	}

	if !info.TimeStart.IsZero() {
		buf = append(buf, "&time_start="...)
		buf = append(buf, TimeStartStr...)
		bs = append(bs, "&time_start="...)
		bs = append(bs, TimeStartStr...)
	}

	buf = append(buf, "&total_fee="...)
	buf = append(buf, TotalFeeStr...)
	bs = append(bs, "&total_fee="...)
	bs = append(bs, TotalFeeStr...)

	if info.TransportFee != 0 {
		buf = append(buf, "&transport_fee="...)
		buf = append(buf, TransportFeeStr...)
		bs = append(bs, "&transport_fee="...)
		bs = append(bs, TransportFeeStr...)
	}

	buf = append(buf, "&key="...)
	buf = append(buf, partnerKey...)

	md5sumArray := md5.Sum(buf[32:])
	md5sumHexBytes := buf[:32]
	hex.Encode(md5sumHexBytes, md5sumArray[:])
	md5sumHexBytes = bytes.ToUpper(md5sumHexBytes)

	bs = append(bs, "&sign="...)
	bs = append(bs, md5sumHexBytes...)

	return
}
