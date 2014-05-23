package wechat

import (
	"github.com/chanxuehong/wechat/message/mass"
)

// 根据分组群发

func (c *Client) MassSendGroupNews(msg *mass.GroupNews) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendGroupText(msg *mass.GroupText) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendGroupVoice(msg *mass.GroupVoice) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendGroupImage(msg *mass.GroupImage) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendGroupVideo(msg *mass.GroupVideo) (*mass.MassResponse, error) {
	return nil, nil
}

// 根据 OpenId 列表群发

func (c *Client) MassSendOpenIdNews(msg *mass.OpenIdNews) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendOpenIdText(msg *mass.OpenIdText) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendOpenIdVoice(msg *mass.OpenIdVoice) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendOpenIdImage(msg *mass.OpenIdImage) (*mass.MassResponse, error) {
	return nil, nil
}

func (c *Client) MassSendOpenIdVideo(msg *mass.OpenIdVideo) (*mass.MassResponse, error) {
	return nil, nil
}

// 删除群发
func (c *Client) MassDelete(msg *mass.DeleteMassRequest) error {
	return nil
}
