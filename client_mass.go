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
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
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

func (c *Client) MassSendGroupNews(msg *mass.GroupNews) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendGroupMsg(msg)
}

func (c *Client) MassSendGroupText(msg *mass.GroupText) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendGroupMsg(msg)
}

func (c *Client) MassSendGroupVoice(msg *mass.GroupVoice) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendGroupMsg(msg)
}

func (c *Client) MassSendGroupImage(msg *mass.GroupImage) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendGroupMsg(msg)
}

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
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
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

func (c *Client) MassSendOpenIdNews(msg *mass.OpenIdNews) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

func (c *Client) MassSendOpenIdText(msg *mass.OpenIdText) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

func (c *Client) MassSendOpenIdVoice(msg *mass.OpenIdVoice) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

func (c *Client) MassSendOpenIdImage(msg *mass.OpenIdImage) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

func (c *Client) MassSendOpenIdVideo(msg *mass.OpenIdVideo) (msgid int, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return c.massSendOpenIdMsg(msg)
}

// 删除群发======================================================================
func (c *Client) MassDelete(msg *mass.DeleteMassRequest) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	token, err := c.Token()
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	url := massDeleteUrlPrefix + token
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
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
