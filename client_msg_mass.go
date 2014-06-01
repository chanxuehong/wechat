package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/chanxuehong/wechat/message/mass"
	"io/ioutil"
)

// 根据分组群发 ==================================================================

// 根据分组群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgMassSendByGroup(msg interface{}) (msgid int, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	_url := clientMessageMassSendByGroupURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
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
func (c *Client) MsgMassSendNewsByGroup(msg *mass.GroupNews) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发文本消息.
func (c *Client) MsgMassSendTextByGroup(msg *mass.GroupText) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发语音消息.
func (c *Client) MsgMassSendVoiceByGroup(msg *mass.GroupVoice) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发图片消息.
func (c *Client) MsgMassSendImageByGroup(msg *mass.GroupImage) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据分组群发视频消息.
func (c *Client) MsgMassSendVideoByGroup(msg *mass.GroupVideo) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByGroup(msg)
}

// 根据 OpenId 列表群发 ==========================================================

// 根据 OpenId列表 群发消息, 之所以不暴露这个接口是因为怕接收到不合法的参数.
func (c *Client) msgMassSendByOpenId(msg interface{}) (msgid int, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	_url := clientMessageMassSendByOpenIdURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
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
func (c *Client) MsgMassSendNewsByOpenId(msg *mass.OpenIdNews) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发文本消息.
func (c *Client) MsgMassSendTextByOpenId(msg *mass.OpenIdText) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发语音消息.
func (c *Client) MsgMassSendVoiceByOpenId(msg *mass.OpenIdVoice) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发图片消息.
func (c *Client) MsgMassSendImageByOpenId(msg *mass.OpenIdImage) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 根据用户列表群发视频消息.
func (c *Client) MsgMassSendVideoByOpenId(msg *mass.OpenIdVideo) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.msgMassSendByOpenId(msg)
}

// 删除群发 =====================================================================
//  NOTE: 只有已经发送成功的消息才能删除删除消息只是将消息的图文详情页失效，已经收到的用户，
//  还是能在其本地看到消息卡片。 另外，删除群发消息只能删除图文消息和视频消息，
//  其他类型的消息一经发送，无法删除。
func (c *Client) MsgMassDelete(msgid int) error {
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

	_url := clientMessageMassDeleteURL(token)
	resp, err := c.httpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
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
