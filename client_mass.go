package wechat

import (
	"github.com/chanxuehong/wechat/message/mass"
)

// 根据分组群发

func (c *Client) MassSendGroupNews(msg *mass.GroupNewsMsg) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendGroupText(msg *mass.GroupTextMsg) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendGroupVoice(msg *mass.GroupVoiceMsg) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendGroupImage(msg *mass.GroupImageMsg) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendGroupVideo(msg *mass.GroupVideoMsg) (*mass.MassResponse, error) {
	return nil, nil
}

// 根据 OpenId 列表群发

func (c *Client) MassSendOpenIdNews(msg *mass.OpenIdNewsMsg) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendOpenIdText(msg *mass.OpenIdTextMsg) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendOpenIdVoice(msg *mass.OpenIdVoiceMsg) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendOpenIdImage(msg *mass.OpenIdImageMsg) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendOpenIdVideo(msg *mass.OpenIdVideoMsg) (*mass.MassResponse, error) {
	return nil, nil
}

// 删除群发
func (c *Client) MassDelete(msg *mass.DeleteMassRequest) error {
	return nil
}
