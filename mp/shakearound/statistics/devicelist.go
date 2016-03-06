package statistics

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const DeviceListPageSize = 50

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
func DeviceList(clt *core.Client, date int64, pageIndex int) (rslt *DeviceListResult, err error) {
	request := struct {
		Date      int64 `json:"date"`
		PageIndex int   `json:"page_index"`
	}{
		Date:      date,
		PageIndex: pageIndex,
	}

	var result struct {
		core.Error
		DeviceListResult
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/devicelist?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}

	devices := result.DeviceListResult.Data.DeviceStatisticsList
	for i := 0; i < len(devices); i++ {
		devices[i].Ftime = result.DeviceListResult.Date
	}
	result.DeviceListResult.ItemCount = len(devices)
	rslt = &result.DeviceListResult
	return
}

// DeviceStatisticsIterator
//
//  iter, err := NewDeviceStatisticsIterator(clt, date, pageIndex)
//  if err != nil {
//      // TODO: 增加你的代码
//  }
//
//  for iter.HasNext() {
//      items, err := iter.NextPage()
//      if err != nil {
//          // TODO: 增加你的代码
//      }
//      // TODO: 增加你的代码
//  }
type DeviceStatisticsIterator struct {
	clt *core.Client

	date          int64
	nextPageIndex int

	lastDeviceListResult *DeviceListResult // 最近一次获取的数据
	nextPageCalled       bool              // NextPage() 是否调用过
}

func (iter *DeviceStatisticsIterator) TotalCount() int {
	return iter.lastDeviceListResult.TotalCount
}

func (iter *DeviceStatisticsIterator) HasNext() bool {
	if !iter.nextPageCalled { // 第一次调用需要特殊对待
		return iter.lastDeviceListResult.ItemCount > 0
	}

	return iter.lastDeviceListResult.ItemCount >= DeviceListPageSize
}

func (iter *DeviceStatisticsIterator) NextPage() (statisticsList []DeviceStatistics, err error) {
	if !iter.nextPageCalled { // 第一次调用需要特殊对待
		iter.nextPageCalled = true

		statisticsList = iter.lastDeviceListResult.Data.DeviceStatisticsList
		return
	}

	rslt, err := DeviceList(iter.clt, iter.date, iter.nextPageIndex)
	if err != nil {
		return
	}

	iter.nextPageIndex++
	iter.lastDeviceListResult = rslt

	statisticsList = rslt.Data.DeviceStatisticsList
	return
}

func NewDeviceStatisticsIterator(clt *core.Client, date int64, pageIndex int) (iter *DeviceStatisticsIterator, err error) {
	// 逻辑上相当于第一次调用 DeviceStatisticsIterator.NextPage, 因为第一次调用 DeviceStatisticsIterator.HasNext 需要数据支撑, 所以提前获取了数据

	rslt, err := DeviceList(clt, date, pageIndex)
	if err != nil {
		return
	}

	iter = &DeviceStatisticsIterator{
		clt: clt,

		date:          date,
		nextPageIndex: pageIndex + 1,

		lastDeviceListResult: rslt,
		nextPageCalled:       false,
	}
	return
}
