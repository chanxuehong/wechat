// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package merchant

import (
	"errors"

	"github.com/chanxuehong/wechat/mp/merchant/express"
)

// 增加邮费模板
//  NOTE: 无需指定 Id 字段
func (c *Client) MerchantExpressAddDeliveryTemplate(template *express.DeliveryTemplate) (templateId int64, err error) {
	if template == nil {
		err = errors.New("template == nil")
		return
	}

	template.Id = 0

	var request = struct {
		DeliveryTemplate *express.DeliveryTemplate `json:"delivery_template"`
	}{
		DeliveryTemplate: template,
	}

	var result struct {
		Error
		TemplateId int64 `json:"template_id"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantExpressAddURL(token)
	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		templateId = result.TemplateId
		return

	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 删除邮费模板
func (c *Client) MerchantExpressDeleteDeliveryTemplate(templateId int64) (err error) {
	var request = struct {
		TemplateId int64 `json:"template_id"`
	}{
		TemplateId: templateId,
	}

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantExpressDeleteURL(token)
	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result
		return
	}
}

// 修改邮费模板
//  NOTE: 需要指定 template.Id 字段
func (c *Client) MerchantExpressUpdateDeliveryTemplate(template *express.DeliveryTemplate) (err error) {
	if template == nil {
		return errors.New("template == nil")
	}

	var request = struct {
		TemplateId       int64                     `json:"template_id"`
		DeliveryTemplate *express.DeliveryTemplate `json:"delivery_template"`
	}{
		TemplateId:       template.Id,
		DeliveryTemplate: template,
	}

	template.Id = 0 // 请求的时候不携带这个参数

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantExpressUpdateURL(token)
	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result
		return
	}
}

// 获取指定ID的邮费模板
func (c *Client) MerchantExpressGetDeliveryTemplateById(templateId int64) (dt *express.DeliveryTemplate, err error) {
	var request = struct {
		TemplateId int64 `json:"template_id"`
	}{
		TemplateId: templateId,
	}

	var result struct {
		Error
		TemplateInfo express.DeliveryTemplate `json:"template_info"`
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantExpressGetByIdURL(token)
	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		dt = &result.TemplateInfo
		return

	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 获取所有邮费模板
func (c *Client) MerchantExpressGetAllDeliveryTemplate() (dts []express.DeliveryTemplate, err error) {
	var result struct {
		Error
		TemplatesInfo []express.DeliveryTemplate `json:"templates_info"`
	}
	result.TemplatesInfo = make([]express.DeliveryTemplate, 0, 16)

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	url_ := merchantExpressGetAllURL(token)
	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		dts = result.TemplatesInfo
		return

	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}
