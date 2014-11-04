// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/chanxuehong/wechat/mp/pay"
	"sort"
	"strconv"
	"time"
)

// 订单详情的参数
type PayPackageParameters map[string]string

func (para PayPackageParameters) SetBankType(BankType string) {
	para["bank_type"] = BankType
}
func (para PayPackageParameters) SetBody(Body string) {
	para["body"] = Body
}
func (para PayPackageParameters) SetAttach(Attach string) {
	para["attach"] = Attach
}
func (para PayPackageParameters) SetPartnerId(PartnerId string) {
	para["partner"] = PartnerId
}
func (para PayPackageParameters) SetOutTradeNo(OutTradeNo string) {
	para["out_trade_no"] = OutTradeNo
}
func (para PayPackageParameters) SetGoodsTag(GoodsTag string) {
	para["goods_tag"] = GoodsTag
}
func (para PayPackageParameters) SetNotifyURL(NotifyURL string) {
	para["notify_url"] = NotifyURL
}
func (para PayPackageParameters) SetBillCreateIP(BillCreateIP string) {
	para["spbill_create_ip"] = BillCreateIP
}
func (para PayPackageParameters) SetCharset(Charset string) {
	para["input_charset"] = Charset
}
func (para PayPackageParameters) SetTotalFee(n int) {
	para["total_fee"] = strconv.FormatInt(int64(n), 10)
}
func (para PayPackageParameters) SetTransportFee(n int) {
	para["transport_fee"] = strconv.FormatInt(int64(n), 10)
}
func (para PayPackageParameters) SetProductFee(n int) {
	para["product_fee"] = strconv.FormatInt(int64(n), 10)
}
func (para PayPackageParameters) SetFeeType(n int) {
	para["fee_type"] = strconv.FormatInt(int64(n), 10)
}
func (para PayPackageParameters) SetTimeStart(t time.Time) {
	para["time_start"] = pay.FormatTime(t)
}
func (para PayPackageParameters) SetTimeExpire(t time.Time) {
	para["time_expire"] = pay.FormatTime(t)
}

// 生成订单详情字符串(package)
//  PayPackageParameters: 订单详情参数
//  partnerKey:           财付通商户权限密钥 Key
func MakePayPackage(para PayPackageParameters, partnerKey string) []byte {
	KVS := make(pay.KVSlice, 0, len(para))
	for K, V := range para {
		if K == "sign" {
			continue
		}
		KVS = append(KVS, pay.KV{K, V})
	}
	sort.Sort(KVS)

	string1KVS := make(pay.KVSlice, 0, len(KVS))
	string2KVS := make(pay.KVSlice, 0, len(KVS))
	string1Count := 0
	string2Count := 0

	for _, KV := range KVS {
		escapedKey := pay.URLEscape(KV.Key) // 安全起见也做 escape
		escapedValue := pay.URLEscape(KV.Value)
		string2KVS = append(string2KVS, pay.KV{escapedKey, escapedValue})
		string2Count += len(escapedKey) + len(escapedValue) + 2 // key=value&

		if KV.Value == "" { // 空值不参加签名
			continue
		}
		string1KVS = append(string1KVS, KV)
		string1Count += len(KV.Key) + len(KV.Value) + 2 // key=value&
	}
	string1Count += len("key=") + len(partnerKey)
	string2Count += len("sign=") + 32 // md5sum

	var buf []byte
	if string1Count >= string2Count {
		buf = make([]byte, string1Count)
	} else {
		buf = make([]byte, string2Count)
	}
	string1 := buf[:0]
	string2 := buf[:0]

	for _, KV := range string1KVS {
		string1 = append(string1, KV.Key...)
		string1 = append(string1, '=')
		string1 = append(string1, KV.Value...)
		string1 = append(string1, '&')
	}
	string1 = append(string1, "key="...)
	string1 = append(string1, partnerKey...)

	md5sum := md5.Sum(string1)
	signature := make([]byte, 32)
	hex.Encode(signature, md5sum[:])
	signature = bytes.ToUpper(signature)

	for _, KV := range string2KVS {
		string2 = append(string2, KV.Key...)
		string2 = append(string2, '=')
		string2 = append(string2, KV.Value...)
		string2 = append(string2, '&')
	}
	string2 = append(string2, "sign="...)
	string2 = append(string2, signature...)

	return string2
}
