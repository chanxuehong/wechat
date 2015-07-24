// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com), magicshui(shuiyuzhe@gmail.com), Harry Rong(harrykobe@gmail.com)

package statistics

import (
	"github.com/chanxuehong/wechat/mp"
)

type DeviceListResult struct {
	PageIndex int   `json:"page_index"`
	Date      int64 `json:"date"`

	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`

	Data struct {
		DeviceStatisticsList []DeviceStatistics `json:"devices"`
	} `json:"data"`
}

// 批量查询设备统计数据接口
func DeviceList(clt *mp.Client, date int64, pageIndex int) (rslt *DeviceListResult, err error) {
	request := struct {
		Date      int64 `json:"date"`
		PageIndex int   `json:"page_index"`
	}{
		Date:      date,
		PageIndex: pageIndex,
	}

	var result struct {
		mp.Error
		DeviceListResult
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/devicelist?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}

	for i, l := 0, len(rslt.Data.DeviceStatisticsList); i < l; i++ {
		rslt.Data.DeviceStatisticsList[i].Ftime = rslt.Date
	}
	rslt.ItemCount = len(rslt.Data.DeviceStatisticsList)
	rslt = &result.DeviceListResult
	return
}
