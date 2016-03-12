package device

import (
	"github.com/chanxuehong/wechat/internal/util"
	"github.com/chanxuehong/wechat/mp/core"
)

// 设备标识
type DeviceIdentifier struct {
	// 设备编号，若填了UUID、major、minor，则可不填设备编号，若二者都填，则以设备编号为优先
	DeviceId *int64 `json:"device_id,omitempty"`

	// UUID、major、minor，三个信息需填写完整，若填了设备编号，则可不填此信息。
	UUID  string `json:"uuid,omitempty"`
	Major *int   `json:"major,omitempty"`
	Minor *int   `json:"minor,omitempty"`
}

func NewDeviceIdentifier1(deviceId int64) *DeviceIdentifier {
	return &DeviceIdentifier{
		DeviceId: util.Int64(deviceId),
	}
}

func NewDeviceIdentifier2(uuid string, major, minor int) *DeviceIdentifier {
	return &DeviceIdentifier{
		UUID:  uuid,
		Major: util.Int(major),
		Minor: util.Int(minor),
	}
}

func NewDeviceIdentifier3(deviceId int64, uuid string, major, minor int) *DeviceIdentifier {
	return &DeviceIdentifier{
		DeviceId: util.Int64(deviceId),
		UUID:     uuid,
		Major:    util.Int(major),
		Minor:    util.Int(minor),
	}
}

// 编辑设备信息
func Update(clt *core.Client, deviceIdentifier *DeviceIdentifier, comment string) (err error) {
	request := struct {
		DeviceIdentifier *DeviceIdentifier `json:"device_identifier,omitempty"`
		Comment          string            `json:"comment"`
	}{
		DeviceIdentifier: deviceIdentifier,
		Comment:          comment,
	}

	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
