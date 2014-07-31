// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

import (
	"testing"
)

func TestRequestZero(t *testing.T) {
	var req Request
	req.ToUserName = "touser"

	if req == zeroRequest {
		t.Error("req must not be zero")
		return
	}

	req.Zero()
	if req.ToUserName != "" {
		t.Error("req must be zero")
		return
	}
}

func BenchmarkRequestZero(b *testing.B) {
	var req Request
	for i := 0; i < b.N; i++ {
		req.Zero()
	}
}
func BenchmarkRequestZeroX(b *testing.B) {
	var req Request
	for i := 0; i < b.N; i++ {
		req.ZeroX()
	}
}

func (msg *Request) ZeroX() *Request {
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
	msg.LocationX = 0
	msg.LocationY = 0
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
