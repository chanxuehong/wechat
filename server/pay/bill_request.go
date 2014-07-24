// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/wechat/pay/native"
	"net/http"
)

// native api 请求订单详情的 Handler
type BillRequestHandler struct {
	paySignKey            string
	invalidRequestHandler InvalidRequestHandlerFunc
	billRequestHandler    BillRequestHandlerFunc
}

// NOTE: 所有参数必须有效
func NewBillRequestHandler(
	paySignKey string,
	invalidRequestHandler InvalidRequestHandlerFunc,
	billRequestHandler BillRequestHandlerFunc,

) (handler *BillRequestHandler) {

	if paySignKey == "" {
		panic(`paySignKey == ""`)
	}
	if invalidRequestHandler == nil {
		panic("invalidRequestHandler == nil")
	}
	if billRequestHandler == nil {
		panic("billRequestHandler == nil")
	}

	handler = &BillRequestHandler{
		paySignKey:            paySignKey,
		invalidRequestHandler: invalidRequestHandler,
		billRequestHandler:    billRequestHandler,
	}

	return
}

// BillRequestHandler 实现 http.Handler 接口
func (handler *BillRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		handler.invalidRequestHandler(w, r, err)
		return
	}

	var billReq native.BillRequest
	if err := xml.NewDecoder(r.Body).Decode(&billReq); err != nil {
		handler.invalidRequestHandler(w, r, err)
		return
	}

	if err := billReq.Check(handler.paySignKey); err != nil {
		handler.invalidRequestHandler(w, r, err)
		return
	}

	handler.billRequestHandler(w, r, &billReq)
}
