package device

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 配置设备与门店的关联关系
func BindLocation(clt *core.Client, deviceIdentifier *DeviceIdentifier, poiId int64) (err error) {
	request := struct {
		DeviceIdentifier *DeviceIdentifier `json:"device_identifier,omitempty"`
		PoiId            int64             `json:"poi_id"`
	}{
		DeviceIdentifier: deviceIdentifier,
		PoiId:            poiId,
	}

	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/bindlocation?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
