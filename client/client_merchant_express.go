// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

package client

import (
	"errors"
	"github.com/chanxuehong/wechat/merchant/express"
)

// 增加邮费模板
func (c *Client) MerchantExpressAdd(template *express.DeliveryTemplate) (templateId int64, err error) {
	if template == nil {
		err = errors.New("template == nil")
		return
	}

	template.Id = 0

	token, err := c.Token()
	if err != nil {
		return
	}
	_url := merchantExpressAddURL(token)

	var request = struct {
		DeliveryTemplate *express.DeliveryTemplate `json:"delivery_template"`
	}{
		DeliveryTemplate: template,
	}

	var result struct {
		Error
		TemplateId int64 `json:"template_id"`
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	templateId = result.TemplateId
	return
}

// 删除邮费模板
func (c *Client) MerchantExpressDelete(templateId int64) error {
	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := merchantExpressDeleteURL(token)

	var request = struct {
		TemplateId int64 `json:"template_id"`
	}{
		TemplateId: templateId,
	}

	var result Error
	if err = c.postJSON(_url, request, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}

	return nil
}

// 修改邮费模板
//  NOTE: 需要指定 template.Id 字段
func (c *Client) MerchantExpressUpdate(template *express.DeliveryTemplate) error {
	if template == nil {
		return errors.New("template == nil")
	}

	token, err := c.Token()
	if err != nil {
		return err
	}
	_url := merchantExpressUpdateURL(token)

	var request = struct {
		TemplateId       int64                     `json:"template_id"`
		DeliveryTemplate *express.DeliveryTemplate `json:"delivery_template"`
	}{
		TemplateId:       template.Id,
		DeliveryTemplate: template,
	}

	template.Id = 0 // 请求的时候不携带这个参数

	var result Error
	if err = c.postJSON(_url, request, &result); err != nil {
		return err
	}

	if result.ErrCode != 0 {
		return &result
	}

	return nil
}

// 获取指定ID的邮费模板
func (c *Client) MerchantExpressGetById(templateId int64) (*express.DeliveryTemplate, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := merchantExpressGetByIdURL(token)

	var request = struct {
		TemplateId int64 `json:"template_id"`
	}{
		TemplateId: templateId,
	}

	var result struct {
		Error
		TemplateInfo express.DeliveryTemplate `json:"template_info"`
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return &result.TemplateInfo, nil
}

// 获取所有邮费模板
func (c *Client) MerchantExpressGetAll() ([]express.DeliveryTemplate, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := merchantExpressGetAllURL(token)

	var result struct {
		Error
		TemplatesInfo []express.DeliveryTemplate `json:"templates_info"`
	}
	result.TemplatesInfo = make([]express.DeliveryTemplate, 0, 64)
	if err = c.getJSON(_url, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	return result.TemplatesInfo, nil
}
