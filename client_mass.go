package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/chanxuehong/wechat/message/mass"
	"io/ioutil"
	"net/http"
)

// 根据分组群发 ==================================================================

// 根据分组群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) massSendGroupMsg(msg interface{}) (msgid int, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	url := massSendMessageByGroupUrlPrefix + token
	resp, err := http.Post(url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result struct {
		Error
		MsgId int `json:"msg_id"`
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}

// 根据分组群发图文消息.
func (c *Client) MassSendGroupNews(msg *mass.GroupNews) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendGroupMsg(msg)
}

// 根据分组群发文本消息.
func (c *Client) MassSendGroupText(msg *mass.GroupText) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendGroupMsg(msg)
}

// 根据分组群发语音消息.
func (c *Client) MassSendGroupVoice(msg *mass.GroupVoice) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendGroupMsg(msg)
}

// 根据分组群发图片消息.
func (c *Client) MassSendGroupImage(msg *mass.GroupImage) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendGroupMsg(msg)
}

// 根据分组群发视频消息.
func (c *Client) MassSendGroupVideo(msg *mass.GroupVideo) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendGroupMsg(msg)
}

// 根据 OpenId 列表群发 ==========================================================

// 根据 OpenId列表 群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) massSendOpenIdMsg(msg interface{}) (msgid int, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	url := massSendMessageByOpenIdUrlPrefix + token
	resp, err := http.Post(url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result struct {
		Error
		MsgId int `json:"msg_id"`
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}

// 根据用户列表群发图文消息.
func (c *Client) MassSendOpenIdNews(msg *mass.OpenIdNews) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

// 根据用户列表群发文本消息.
func (c *Client) MassSendOpenIdText(msg *mass.OpenIdText) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

// 根据用户列表群发语音消息.
func (c *Client) MassSendOpenIdVoice(msg *mass.OpenIdVoice) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

// 根据用户列表群发图片消息.
func (c *Client) MassSendOpenIdImage(msg *mass.OpenIdImage) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

// 根据用户列表群发视频消息.
func (c *Client) MassSendOpenIdVideo(msg *mass.OpenIdVideo) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

// 删除群发 =====================================================================
//  NOTE: 只有已经发送成功的消息才能删除删除消息只是将消息的图文详情页失效，已经收到的用户，
//  还是能在其本地看到消息卡片。 另外，删除群发消息只能删除图文消息和视频消息，
//  其他类型的消息一经发送，无法删除。
func (c *Client) MassDelete(msgid int) error {
	token, err := c.Token()
	if err != nil {
		return err
	}

	var deleteRequest struct {
		MsgId int `json:"msgid"`
	}
	deleteRequest.MsgId = msgid

	jsonData, err := json.Marshal(deleteRequest)
	if err != nil {
		return err
	}

	url := massDeleteUrlPrefix + token
	resp, err := http.Post(url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result Error
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.ErrCode != 0 {
		return &result
	}
	return nil
}
