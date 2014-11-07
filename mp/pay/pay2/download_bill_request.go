// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/chanxuehong/wechat/mp/pay"
)

// 对账单下载接口 请求参数
type DownloadBillRequest map[string]string

func (req DownloadBillRequest) SetSPId(str string) {
	req["spid"] = str
}
func (req DownloadBillRequest) SetTransactionTime(t time.Time) {
	req["trans_time"] = t.In(pay.BeijingLocation).Format("2006-01-02")
}
func (req DownloadBillRequest) SetTimestamp(t time.Time) {
	req["stamp"] = strconv.FormatInt(t.Unix(), 10)
}
func (req DownloadBillRequest) SetCftSignMethod(n int) {
	req["cft_signtype"] = strconv.FormatInt(int64(n), 10)
}
func (req DownloadBillRequest) SetMchType(n int) {
	req["mchtype"] = strconv.FormatInt(int64(n), 10)
}

// 设置签名字段.
//  Key: 商户密钥
//
//  NOTE: 要求在 DownloadBillRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req DownloadBillRequest) SetSignature(Key string) (err error) {
	Hash := md5.New()
	hashsum := make([]byte, 32)
	var str string

	if str = req["spid"]; str != "" {
		Hash.Write([]byte("spid="))
		Hash.Write([]byte(str))
		Hash.Write([]byte{'&'})
	}
	if str = req["trans_time"]; str != "" {
		Hash.Write([]byte("trans_time="))
		Hash.Write([]byte(str))
		Hash.Write([]byte{'&'})
	}
	if str = req["stamp"]; str != "" {
		Hash.Write([]byte("stamp="))
		Hash.Write([]byte(str))
		Hash.Write([]byte{'&'})
	}
	if str = req["cft_signtype"]; str != "" {
		Hash.Write([]byte("cft_signtype="))
		Hash.Write([]byte(str))
		Hash.Write([]byte{'&'})
	}
	if str = req["mchtype"]; str != "" {
		Hash.Write([]byte("mchtype="))
		Hash.Write([]byte(str))
		Hash.Write([]byte{'&'})
	}
	Hash.Write([]byte("key="))
	Hash.Write([]byte(Key))

	hex.Encode(hashsum, Hash.Sum(nil))
	req["sign"] = string(hashsum)
	return
}
