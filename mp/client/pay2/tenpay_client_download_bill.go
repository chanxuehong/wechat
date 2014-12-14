// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chanxuehong/wechat/mp/pay"
	"github.com/chanxuehong/wechat/mp/pay/pay2"
)

// <html>
//   <body>03020003:该日期对帐单还没有生成</body>
// </html>
type downloadBillFailResponse struct {
	XMLName struct{} `xml:"html"`
	Body    string   `xml:"body"`
}

func (c *TenpayClient) DownloadBill(req pay2.DownloadBillRequest) (data []byte, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	buf := textBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	defer textBufferPool.Put(buf)               // important
	buf.Reset()                                 // important

	if err = pay.FormatMapToXML(buf, req); err != nil {
		return
	}

	url_ := "http://mch.tenpay.com/cgi-bin/mchdown_real_new.cgi"

	resp, err := c.httpClient.Post(url_, "text/xml; charset=utf-8", buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", resp.Status)
		return
	}

	httpBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result downloadBillFailResponse
	if err = xml.Unmarshal(httpBody, &result); err != nil {
		data = httpBody
		err = nil
		return
	} else {
		err = errors.New(result.Body)
		return
	}
}
