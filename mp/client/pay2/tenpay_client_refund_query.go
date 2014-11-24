// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"errors"
	"fmt"

	"github.com/chanxuehong/wechat/mp/pay/pay2"
)

func (c *TenpayClient) NormalRefundQuery(req pay2.NormalRefundQueryRequest) (resp pay2.NormalRefundQueryResponse, err error) {
	if req == nil {
		err = errors.New("req == nil")
		return
	}

	resp = make(map[string]string)
	url_ := "https://gw.tenpay.com/gateway/normalrefundquery.xml"

	if err = c.postXML(url_, req, resp); err != nil {
		return
	}

	if havePartnerId := resp.PartnerId(); havePartnerId != c.partnerId {
		err = fmt.Errorf("PartnerId mismatch, \r\nhave: %q, \r\nwant: %q", havePartnerId, c.partnerId)
		return
	}
	if err = resp.CheckSignature(c.partnerKey); err != nil {
		return
	}

	return
}
