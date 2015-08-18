// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chanxuehong/util"
	"github.com/chanxuehong/wechat/mch"
)

// 统一下单.
func UnifiedOrder(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/pay/unifiedorder", req)
}

// 查询订单.
func OrderQuery(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/pay/orderquery", req)
}

// 关闭订单.
func CloseOrder(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/pay/closeorder", req)
}

// 申请退款.
//  NOTE: 请求需要双向证书.
func Refund(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/secapi/pay/refund", req)
}

// 查询退款.
func RefundQuery(pxy *mch.Proxy, req map[string]string) (resp map[string]string, err error) {
	return pxy.PostXML("https://api.mch.weixin.qq.com/pay/refundquery", req)
}

// 下载对账单.
func DownloadBill(req map[string]string, httpClient *http.Client) (data []byte, err error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	bodyBuf := bytes.NewBuffer(make([]byte, 0, 1024))
	if err = util.FormatMapToXML(bodyBuf, req); err != nil {
		return
	}

	httpResp, err := httpClient.Post("https://api.mch.weixin.qq.com/pay/downloadbill", "text/xml; charset=utf-8", bodyBuf)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	var result mch.Error
	if err = xml.Unmarshal(respBody, &result); err == nil {
		err = &result
		return
	}

	data = respBody
	err = nil
	return
}
