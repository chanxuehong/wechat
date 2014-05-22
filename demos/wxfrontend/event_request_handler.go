package wxfrontend

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/chanxuehong/wechat/message"
	"io"
	"net/http"
)

func subscribeEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		"欢迎关 产先生",
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

// 通过扫描二维码订阅
func subscribeEventByScanRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		"欢迎关 产先生, 您是通过扫描二维码订阅的",
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

// 取消订阅
func unsubscribeEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	io.WriteString(w, "")
}

// 已经订阅用户扫描二维码
func scanEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		"谢谢扫描, 暂时还没有实现更多动能, 敬请期待...",
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

// 上报地理位置事件
func locationEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		fmt.Sprintf(
			"您的地理位置是:\n北纬:%f\n东经:%f\n精度:%f\n",
			rqstMsg.Latitude, rqstMsg.Longitude, rqstMsg.Precision,
		),
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

// 点击菜单拉取消息时的事件推送
func clickEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		"暂时还没有实现这个动能, 敬请期待...",
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

// 点击菜单跳转链接时的事件推送
func viewEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		"暂时还没有实现这个动能, 敬请期待...",
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

// 群发, 事件推送群发结果
func masssendjobfinishEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		"暂时还没有实现这个动能, 敬请期待...",
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}
