// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/chanxuehong/wechat/mp/pay"
	"io/ioutil"
	"net/http"
)

func (c *Client) DownloadBill(req map[string]string) (data []byte, err error) {
	buf := textBufferPool.Get().(*bytes.Buffer) // io.ReadWriter
	buf.Reset()                                 // important
	defer textBufferPool.Put(buf)               // important

	if err = pay.FormatMapToXML(buf, req); err != nil {
		return
	}

	url_ := "https://api.mch.weixin.qq.com/pay/downloadbill"

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

	var result Error
	if err = xml.Unmarshal(httpBody, &result); err != nil {
		data = httpBody
		err = nil
		return
	} else {
		err = &result
		return
	}
}
