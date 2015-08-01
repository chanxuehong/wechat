// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package send

import (
	"net/http"
	"errors"

	"github.com/chanxuehong/wechat/corp"

)


type Client corp.Client

func NewClient(srv corp.AccessTokenServer, clt *http.Client) *Client {
	return (*Client)(corp.NewClient(srv, clt))
}



// 创建会话
func (clt *Client) Create(chatInfo *ChatInfo) (err error) {
	var result corp.Error
	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/chat/create?access_token="
	if err = ((*corp.Client)(clt)).PostJSON(incompleteURL, &chatInfo, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result
		return
	}

	return
}

func (clt *Client) GetChatInfo(chatId string) (chatInfo ChatInfo, err error) {
	var result struct {
		corp.Error
		ChatInfo ChatInfo `json:"chat_info"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/chat/get?chatid=" +
	chatId + "&access_token="

	if err = ((*corp.Client)(clt)).GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	chatInfo=result.ChatInfo
	return

}


// 退出会话
func (clt *Client) Quit(chatId, opUser string) (err error) {
	var request = struct {
		ChatId string `json:"chatid"`
		OpUser string `json:"op_user"`
	}{
		ChatId  :    chatId,
		OpUser  :   opUser,
	}

	var result struct {
		corp.Error
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/chat/quit?access_token="
	if err = ((*corp.Client)(clt)).PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}



// 更新会话
func (clt *Client) Update(updateChatInfo *UpdateInfo) (err error) {
	var result struct {
		corp.Error
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/chat/update?access_token="
	if err = ((*corp.Client)(clt)).PostJSON(incompleteURL, &updateChatInfo, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}


//清除消息未读状态
func (clt *Client) ClearNotify(request *ClearNotify) (err error) {
	var result struct {
		corp.Error
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/chat/clearnotify?access_token="
	if err = ((*corp.Client)(clt)).PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}


//设置成员新消息免打扰
// 成员新消息免打扰参数，数组，最大支持10000个成员
func (clt *Client) SetMute(userMuteList []*UserMute) (invaliduser []string, err error) {
	var request = struct {
		UserMuteList []*UserMute `json:"user_mute_list"`
	}{
		UserMuteList:userMuteList,
	}


	var result struct {
		corp.Error
		Invaliduser []string `json:"invaliduser"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/chat/setmute?access_token="
	if err = ((*corp.Client)(clt)).PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	invaliduser=result.Invaliduser
	return
}




//发送消息
func (clt *Client) send(msg interface{}) (err error) {
	var result struct {
		corp.Error

	}
	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/chat/send?access_token="
	if err = ((*corp.Client)(clt)).PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}

	if result.ErrCode != corp.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}


func (clt *Client) SendText(msg *Text) (err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendImage(msg *Image) (err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendVoice(msg *Voice) (err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendVideo(msg *Video) (err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendFile(msg *File) (err error) {
	if msg == nil {
		err = errors.New("nil msg")
		return
	}
	return clt.send(msg)
}

