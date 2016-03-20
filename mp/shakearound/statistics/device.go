<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com), magicshui(shuiyuzhe@gmail.com), Harry Rong(harrykobe@gmail.com)

package statistics

import (
	"github.com/chanxuehong/wechat/mp"
=======
package statistics

import (
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
	"github.com/chanxuehong/wechat/mp/shakearound/device"
)

// 以设备为维度的数据统计接口
<<<<<<< HEAD
func Device(clt *mp.Client, deviceIdentifier *device.DeviceIdentifier, beginDate, endDate int64) (data []StatisticsBase, err error) {
=======
func Device(clt *core.Client, deviceIdentifier *device.DeviceIdentifier, beginDate, endDate int64) (data []StatisticsBase, err error) {
>>>>>>> github/v2
	request := struct {
		DeviceIdentifier *device.DeviceIdentifier `json:"device_identifier,omitempty"`
		BeginDate        int64                    `json:"begin_date"`
		EndDate          int64                    `json:"end_date"`
	}{
		DeviceIdentifier: deviceIdentifier,
		BeginDate:        beginDate,
		EndDate:          endDate,
	}

	var result struct {
<<<<<<< HEAD
		mp.Error
=======
		core.Error
>>>>>>> github/v2
		Data []StatisticsBase `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/device?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

<<<<<<< HEAD
	if result.ErrCode != mp.ErrCodeOK {
=======
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result.Error
		return
	}
	data = result.Data
	return
}
