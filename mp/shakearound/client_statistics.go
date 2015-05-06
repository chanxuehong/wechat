// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)
package shakearound

import (
	"github.com/chanxuehong/wechat/mp"
)

type ShakeBeaconInfo struct {
	Distince float64 `json:"distince"` // Beacon信号与手机的距离，单位为米
	Major    int64   `json:"major"`
	Minor    int64   `json:"minor"`
	UUID     string  `json:"uuid"`
}

type ShakeBeaconStatistics struct {
	ClickPv int64 `json:"click_pv"`
	ClickUv int64 `json:"click_uv"`
	Ftime   int64 `json:"ftime"`
	ShakePv int64 `json:"shake_pv"`
	ShakeUv int64 `json:"shake_uv"`
}

func (clt *Client) StatisticsDevice(device ShakeDeviceIdentifier, beginDate, endDate int64) (data []ShakeBeaconStatistics, err error) {
	var request = struct {
		DeviceIdentifier ShakeDeviceIdentifier `json:"device_identifier"`
		BeginDate        int64                 `json:"begin_date"`
		EndDate          int64                 `json:"end_date"`
	}{
		DeviceIdentifier: device,
		BeginDate:        beginDate,
		EndDate:          endDate,
	}
	var result struct {
		mp.Error
		Data []ShakeBeaconStatistics `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/device?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	data = result.Data

	return
}

func (clt *Client) StatisticsPage(pageId, beginDate, endDate int64) (data []ShakeBeaconStatistics, err error) {
	var request = struct {
		PageId    int64 `json:"page_id"`
		BeginDate int64 `json:"begin_date"`
		EndDate   int64 `json:"end_date"`
	}{
		PageId:    pageId,
		BeginDate: beginDate,
		EndDate:   endDate,
	}
	var result struct {
		mp.Error
		Data []ShakeBeaconStatistics `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/page?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	data = result.Data

	return
}
