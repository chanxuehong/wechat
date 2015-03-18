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
)

// 统一下单.
func (clt *Client) UnifiedOrder(req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/pay/unifiedorder", req)
}

// 订单查询.
func (clt *Client) OrderQuery(req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/pay/orderquery", req)
}

// 关闭订单.
func (clt *Client) CloseOrder(req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/pay/closeorder", req)
}

// 申请退款.
//  NOTE: 请求需要双向证书.
func (clt *Client) Refund(req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/secapi/pay/refund", req)
}

// 退款查询.
func (clt *Client) RefundQuery(req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/pay/refundquery", req)
}

// 测速上报.
func (clt *Client) Report(req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/payitil/report", req)
}

// 转换短链接.
func (clt *Client) ShortURL(req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/tools/shorturl", req)
}

// 下载对账单.
func (clt *Client) DownloadBill(req map[string]string) (data []byte, err error) {
	bodyBuf := textBufferPool.Get().(*bytes.Buffer)
	bodyBuf.Reset()
	defer textBufferPool.Put(bodyBuf)

	if err = util.FormatMapToXML(bodyBuf, req); err != nil {
		return
	}

	url := "https://api.mch.weixin.qq.com/pay/downloadbill"
	httpResp, err := clt.httpClient.Post(url, "text/xml; charset=utf-8", bodyBuf)
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

	var result Error
	if err = xml.Unmarshal(respBody, &result); err == nil {
		err = &result
		return
	}

	data = respBody
	err = nil
	return
}
