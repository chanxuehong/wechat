// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package send

import (
	"errors"
	"net/http"

	"github.com/chanxuehong/wechat/corp"
)

type Client struct {
	corp.CorpClient
}

// 创建一个新的 Client.
//  如果 HttpClient == nil 则默认用 http.DefaultClient
func NewClient(AccessTokenServer corp.AccessTokenServer, HttpClient *http.Client) *Client {
	if AccessTokenServer == nil {
		panic("AccessTokenServer == nil")
	}
	if HttpClient == nil {
		HttpClient = http.DefaultClient
	}

	return &Client{
		CorpClient: corp.CorpClient{
			AccessTokenServer: AccessTokenServer,
			HttpClient:        HttpClient,
		},
	}
}

// 发送消息返回的数据结构
type Result struct {
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

func (clt *Client) SendText(msg *Text) (r *Result, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendImage(msg *Image) (r *Result, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendVoice(msg *Voice) (r *Result, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendVideo(msg *Video) (r *Result, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendFile(msg *File) (r *Result, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendNews(msg *News) (r *Result, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendMPNews(msg *MPNews) (r *Result, err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return clt.send(msg)
}

func (clt *Client) send(msg interface{}) (r *Result, err error) {
	var result struct {
		corp.Error
		Result
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="
	if err = clt.PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	r = &result.Result
	return
}
