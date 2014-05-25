package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/message/mass"
	"io/ioutil"
	"net/http"
)

// 根据分组群发 ==================================================================

func (c *Client) massSendGroupMsg(jsonData []byte) (*mass.MassResponse, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(massSendMessageByGroupUrlFormat, token)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result mass.MassResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &Error{
			ErrCode: result.ErrCode,
			ErrMsg:  result.ErrMsg,
		}
	}
	return &result, nil
}

func (c *Client) MassSendGroupNews(msg *mass.GroupNews) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendGroupMsg(msgBody)
}

func (c *Client) MassSendGroupText(msg *mass.GroupText) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendGroupMsg(msgBody)
}

func (c *Client) MassSendGroupVoice(msg *mass.GroupVoice) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendGroupMsg(msgBody)
}

func (c *Client) MassSendGroupImage(msg *mass.GroupImage) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendGroupMsg(msgBody)
}

func (c *Client) MassSendGroupVideo(msg *mass.GroupVideo) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendGroupMsg(msgBody)
}

// 根据 OpenId 列表群发 ==========================================================

func (c *Client) massSendOpenIdMsg(jsonData []byte) (*mass.MassResponse, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(massSendMessageByOpenIdUrlFormat, token)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result mass.MassResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &Error{
			ErrCode: result.ErrCode,
			ErrMsg:  result.ErrMsg,
		}
	}
	return &result, nil
}

func (c *Client) MassSendOpenIdNews(msg *mass.OpenIdNews) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendOpenIdMsg(msgBody)
}

func (c *Client) MassSendOpenIdText(msg *mass.OpenIdText) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendOpenIdMsg(msgBody)
}

func (c *Client) MassSendOpenIdVoice(msg *mass.OpenIdVoice) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendOpenIdMsg(msgBody)
}

func (c *Client) MassSendOpenIdImage(msg *mass.OpenIdImage) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendOpenIdMsg(msgBody)
}

func (c *Client) MassSendOpenIdVideo(msg *mass.OpenIdVideo) (*mass.MassResponse, error) {
	if msg == nil {
		return nil, errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return c.massSendOpenIdMsg(msgBody)
}

// 删除群发======================================================================
func (c *Client) MassDelete(msg *mass.DeleteMassRequest) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	token, err := c.Token()
	if err != nil {
		return err
	}

	url := fmt.Sprintf(massDeleteUrlFormat, token)
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(msgBody))
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
