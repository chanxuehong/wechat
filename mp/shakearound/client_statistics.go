// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com) Harry Rong(harrykobe@gmail.com)
package shakearound

import (
	"github.com/chanxuehong/wechat/mp"
)

type statistic struct {
	Ftime int64 `json:"ftime"`
	ClickPv int `json:"click_pv"`
	ClickUv int `json:"click_uv"`
	ShakePv int `json:"shake_pv"`
	ShakeUv int `json:"shake_uv"`
}

func (clt Client) GetDeviceStatistics(deviceBase *DeviceBase, beginDate, endDate int64) (statistics *[]statistic, err error) {
	var request = struct {
		DeviceIdentifier *DeviceBase `json:"device_identifier"`
		BeginDate int64 `json:"begin_date"`
		EndDate int64 `json:"end_date"`
	}{
		DeviceIdentifier: deviceBase,
		BeginDate: beginDate,
		EndDate: endDate,
	}

	var result struct {
		mp.Error
		Data []statistic `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/device?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	statistics = &result.Data
	return
}

func (clt Client) GetPageStatistics(pageId int, beginDate, endDate int64) (statistics *[]statistic, err error) {
	var request = struct {
		PageId int `json:"page_id"`
		BeginDate int64 `json:"begin_date"`
		EndDate int64 `json:"end_date"`
	}{
		PageId: pageId,
		BeginDate: beginDate,
		EndDate: endDate,
	}

	var result struct {
		mp.Error
		Data []statistic `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/page?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	statistics = &result.Data
	return
}