// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package custom

import (
	"errors"
	"net/http"

	"github.com/chanxuehong/wechat/mp"
)

type Client struct {
	*mp.Client
}

// 兼容保留, 建議實際項目全局維護一個 *mp.Client
func NewClient(AccessTokenServer mp.AccessTokenServer, httpClient *http.Client) Client {
	return Client{
		Client: mp.NewClient(AccessTokenServer, httpClient),
	}
}

// 发送客服消息, 文本.
func (clt Client) SendText(msg *Text) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 图片.
func (clt Client) SendImage(msg *Image) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 语音.
func (clt Client) SendVoice(msg *Voice) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 视频.
func (clt Client) SendVideo(msg *Video) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 音乐.
func (clt Client) SendMusic(msg *Music) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	return clt.send(msg)
}

// 发送客服消息, 图文.
func (clt Client) SendNews(msg *News) (err error) {
	if msg == nil {
		return errors.New("msg == nil")
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return clt.send(msg)
}

func (clt Client) send(msg interface{}) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="
	if err = clt.PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
