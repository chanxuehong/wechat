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
	Comment string   `json:"comment,omitempty"`
	PageIds []string `json:"page_ids,omitempty"`
	Status  int      `json:"status,omitempty"`
	PoiId   int64    `json:"poi_id,omitempty"`
}

func (clt *Client) DeviceApplyId(quantity int64, applyReason, comment string, poiId int64) (applyId int64,
	deviceIdentifiers []ShakeDeviceIdentifier,
	auditStatus int,
	auditComment string, err error) {
	var request = struct {
		Quantity    int64  `json:"quantity"`
		ApplyReason string `json:"apply_reason"`
		Comment     string `json:"comment,omtiempty"`
		PoiId       int64  `json:"poi_id,omtiempty"`
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
		Data interface{} `json:"data"`
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
		Data interface{} `json:"data"`
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

// 测试没有通过
func (clt *Client) DeviceSearch(deviceIndentifier []ShakeDeviceIdentifier, applyId, begin, count int64) (devices []ShakearoundDevice, totalCount int64, err error) {
	var request = struct {
		DeviceIdentifier []ShakeDeviceIdentifier `json:"device_identifiers,omitempty"`
		ApplyId          int64                   `json:"apply_id,omitempty"`
		Begin            int64                   `json:"begin,omitempty"`
		Count            int64                   `json:"count,omitempty"`
	}{
		DeviceIdentifier: deviceIndentifier,
		ApplyId:          applyId,
		Begin:            begin,
		Count:            count,
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
