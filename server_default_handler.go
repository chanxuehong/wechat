package wechat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/chanxuehong/wechat/message"
	"io"
	"net/http"
)

func invalidRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
	io.WriteString(w, "")
	fmt.Printf("invalid request, %s\n", err.Error())
}

func unknownRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	io.WriteString(w, "")
	b, err := json.Marshal(rqstMsg)
	if err != nil {
		fmt.Printf("unknown request: %+v\n", rqstMsg)
		return
	}
	fmt.Printf("unknown request: %s\n", b)
}

// common request handler ======================================================

func textRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		fmt.Sprintf("您刚才发过来的内容是:\n%s\n", rqstMsg.Content),
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

func imageRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewImageResponse(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		rqstMsg.MediaId,
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

func voiceRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewVoiceResponse(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		rqstMsg.MediaId,
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

func voiceRecognitionRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
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

func videoRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewVideoResponse(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		rqstMsg.MediaId,
		"测试",
		"返回你刚才发送的视频",
	)

	b, err := xml.Marshal(respMsg)
	if err != nil {
		//TODO: 增加相应的处理
		io.WriteString(w, "")
		return
	}
	io.Copy(w, bytes.NewReader(b))
}

func locationRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		fmt.Sprintf(
			"您的地理位置是:\n北纬:%f\n东经:%f\n位置信息:%s\n",
			rqstMsg.Location_X, rqstMsg.Location_Y, rqstMsg.Label,
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

func linkRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
		rqstMsg.FromUserName,
		rqstMsg.ToUserName,
		fmt.Sprintf(
			"您刚才发过来的链接是:\n消息类型:%s\n消息描述:%s\n链接地址:%s\n",
			rqstMsg.Title, rqstMsg.Description, rqstMsg.Url,
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

// event request handler =======================================================

func subscribeEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
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
func subscribeEventByScanRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
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
func unsubscribeEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	io.WriteString(w, "")
}

// 已经订阅用户扫描二维码
func scanEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
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
func locationEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
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
func clickEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
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
func viewEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
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
func masssendjobfinishEventRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.Request) {
	respMsg := message.NewTextResponse(
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
