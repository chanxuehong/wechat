<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com), magicshui(shuiyuzhe@gmail.com), Harry Rong(harrykobe@gmail.com)

package device

import (
	"github.com/chanxuehong/wechat/mp"
)

// 配置设备与门店的关联关系
func BindLocation(clt *mp.Client, deviceIdentifier *DeviceIdentifier, poiId int64) (err error) {
=======
package device

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 配置设备与门店的关联关系
func BindLocation(clt *core.Client, deviceIdentifier *DeviceIdentifier, poiId int64) (err error) {
>>>>>>> github/v2
	request := struct {
		DeviceIdentifier *DeviceIdentifier `json:"device_identifier,omitempty"`
		PoiId            int64             `json:"poi_id"`
	}{
		DeviceIdentifier: deviceIdentifier,
		PoiId:            poiId,
	}

<<<<<<< HEAD
	var result mp.Error
=======
	var result core.Error
>>>>>>> github/v2

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/bindlocation?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

<<<<<<< HEAD
	if result.ErrCode != mp.ErrCodeOK {
=======
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result
		return
	}
	return
}
