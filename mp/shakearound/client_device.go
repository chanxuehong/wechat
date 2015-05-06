// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)
package shakearound

import (
	"github.com/chanxuehong/wechat/mp"
)

type ShakeDeviceIdentifier struct {
	DeviceId int64  `json:"device_id"`
	UUID     string `json:"uuid"`
	Major    int64  `json:"major"`
	Minor    int64  `json:"minor"`
}
type ShakearoundDevice struct {
	ShakeDeviceIdentifier
	// omtiempty
	Comment string   `json:"comment,omtiempty"`
	PageIds []string `json:"page_ids,omtiempty"`
	Status  int      `json:"status,omtiempty"`
	PoiId   int64    `json:"poi_id,omtiempty"`
}

// 接口说明 申请配置设备所需的UUID、Major、Minor。申请成功后返回批次ID，可用返回的批次ID用“查询设备列表”接口拉取本次申请的设备ID。
// 单次新增设备超过500个，需走人工审核流程，大概需要三个工作日；单次新增设备不超过500个的，当日可返回申请的设备ID。
// 一个公众账号最多可申请99999个设备ID，如需申请的设备ID数超过最大限额，请邮件至zhoubian@tencent.com，
// 邮件格式如下： 标题：申请提升设备ID额度 内容：1、公众账号名称及appid（wx开头的字符串，在mp平台可查看）
// 2、用途 3、预估需要多少设备ID

// 新增页面
func (clt *Client) DeviceApplyId(quantity int64, applyReason, comment string, poiId int64) (applyId int64,
	deviceIdentifiers []ShakeDeviceIdentifier,
	auditStatus int,
	auditComment string, err error) {
	var request = struct {
		Quantity    int64  `json:"quantity"`
		ApplyReason string `json:"apply_reason"`
		Comment     string `json:"comment"`
		PoiId       int64  `json:"poi_id"`
	}{
		Quantity:    quantity,
		ApplyReason: applyReason,
		Comment:     comment,
		PoiId:       poiId,
	}
	var result struct {
		mp.Error
		Data struct {
			ApplyId           int64                   `json:"apply_id"`
			AuditStatus       int                     `json:"audit_status,omtiempty"`
			AuditComment      string                  `json:"audit_comment,omtiempty"`
			DeviceIdentifiers []ShakeDeviceIdentifier `json:"device_identifiers,omtiempty"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/applyid?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	applyId = result.Data.ApplyId
	deviceIdentifiers = result.Data.DeviceIdentifiers
	auditStatus = result.Data.AuditStatus
	auditComment = result.Data.AuditComment

	return
}

// 新增页面
func (clt *Client) DeviceUpdate(device ShakeDeviceIdentifier, comment string) (err error) {
	var request = struct {
		DeviceIdentifier ShakeDeviceIdentifier `json:"device_identifier"`
		Comment          string                `json:"comment"`
	}{
		DeviceIdentifier: device,
		Comment:          comment,
	}
	var result struct {
		mp.Error
		Data struct {
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}

	return
}

// 新增页面
func (clt *Client) DeviceBindLocation(device ShakeDeviceIdentifier, poiId int64) (err error) {
	var request = struct {
		DeviceIdentifier ShakeDeviceIdentifier `json:"device_identifier"`
		PoiId            int64                 `json:"poi_id"`
	}{
		DeviceIdentifier: device,
		PoiId:            poiId,
	}
	var result struct {
		mp.Error
		Data struct {
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/bindlocation?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}

	return
}

// 新增页面
func (clt *Client) DeviceSearch(device ShakeDeviceIdentifier, applyId, begin, end int64) (devices []ShakearoundDevice, totalCount int64, err error) {
	var request = struct {
		DeviceIdentifier ShakeDeviceIdentifier `json:"device_identifier"`
		ApplyId          int64                 `json:"apply_id"`
		Begin            int64                 `json:"begin"`
		End              int64                 `json:"end"`
	}{
		DeviceIdentifier: device,
		ApplyId:          applyId,
		Begin:            begin,
		End:              end,
	}
	var result struct {
		mp.Error
		Data struct {
			Devices    []ShakearoundDevice `json:"devices"`
			TotalCount int64               `json:"total_count"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/search?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	devices = result.Data.Devices
	totalCount = result.Data.TotalCount

	return
}

// 新增页面
func (clt *Client) DeviceBindPage(device ShakeDeviceIdentifier, pageIds []int64, bind, append_ int64) (err error) {
	var request = struct {
		DeviceIdentifier ShakeDeviceIdentifier `json:"device_identifier"`
		PageIds          []int64               `json:"page_ids"`
		Bind             int64                 `json:"bind"`
		Append           int64                 `json:"append"`
	}{
		DeviceIdentifier: device,
		PageIds:          pageIds,
		Bind:             bind,
		Append:           append_,
	}
	var result struct {
		mp.Error
		Data struct {
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/bindpage?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}

	return
}
