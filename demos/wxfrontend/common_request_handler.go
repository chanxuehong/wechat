package wxfrontend

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/chanxuehong/wechat/message"
	"io"
	"net/http"
)

func textRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
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

func imageRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewImageResponseMsg(
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

func voiceRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewVoiceResponseMsg(
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

func voiceRecognitionRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
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

func videoRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewVideoResponseMsg(
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

func locationRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
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

func linkRequestHandler(w http.ResponseWriter, r *http.Request, rqstMsg *message.RequestMsg) {
	respMsg := message.NewTextResponseMsg(
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
