// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

import (
	"testing"
)

var _test_request Request

func TestRequestZero(t *testing.T) {
	_test_request.ToUserName = "touser"

	if _test_request == _zero_request {
		t.Error("_test_request must not be zero")
		return
	}

	_test_request.Zero()
	if _test_request.ToUserName != "" {
		t.Error("_test_request must be zero")
		return
	}
}
func BenchmarkRequestZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_test_request.Zero()
	}
}
func BenchmarkRequest_Zero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_test_request._Zero()
	}
}
func (msg *Request) _Zero() *Request {
	msg.CommonHead.ToUserName = ""
	msg.CommonHead.FromUserName = ""
	msg.CommonHead.CreateTime = 0
	msg.CommonHead.MsgType = ""

	msg.MsgId = 0
	msg.MsgID = 0

	// common message
	msg.Content = ""
	msg.MediaId = ""
	msg.PicURL = ""
	msg.Format = ""
	msg.Recognition = ""
	msg.ThumbMediaId = ""
	msg.Location_X = 0
	msg.Location_Y = 0
	msg.Scale = 0
	msg.Label = ""
	msg.Title = ""
	msg.Description = ""
	msg.URL = ""

	// event message
	msg.Event = ""
	msg.EventKey = ""
	msg.Ticket = ""
	msg.Latitude = 0
	msg.Longitude = 0
	msg.Precision = 0

	msg.Status = ""

	msg.TotalCount = 0

	msg.FilterCount = 0
	msg.SentCount = 0
	msg.ErrorCount = 0

	msg.OrderId = ""
	msg.OrderStatus = 0
	msg.ProductId = ""
	msg.SkuInfo = ""

	return msg
}
