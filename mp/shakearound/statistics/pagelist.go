package statistics

import (
	"github.com/mingjunyang/wechat.v2/mp/core"
)

const PageListPageSize = 50

type PageListResult struct {
	PageIndex int   `json:"page_index"`
	Date      int64 `json:"date"`

	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`

	Data struct {
		PageStatisticsList []PageStatistics `json:"pages"`
	} `json:"data"`
}

// 批量查询设备统计数据接口
func PageList(clt *core.Client, date int64, pageIndex int) (rslt *PageListResult, err error) {
	request := struct {
		Date      int64 `json:"date"`
		PageIndex int   `json:"page_index"`
	}{
		Date:      date,
		PageIndex: pageIndex,
	}

	var result struct {
		core.Error
		PageListResult
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/statistics/pagelist?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}

	pages := result.PageListResult.Data.PageStatisticsList
	for i := 0; i < len(pages); i++ {
		pages[i].Ftime = result.PageListResult.Date
	}
	result.PageListResult.ItemCount = len(pages)
	rslt = &result.PageListResult
	return
}

// PageStatisticsIterator
//
//  iter, err := NewPageStatisticsIterator(clt, date, pageIndex)
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
type PageStatisticsIterator struct {
	clt *core.Client

	date          int64
	nextPageIndex int

	lastPageListResult *PageListResult // 最近一次获取的数据
	nextPageCalled     bool            // NextPage() 是否调用过
}

func (iter *PageStatisticsIterator) TotalCount() int {
	return iter.lastPageListResult.TotalCount
}

func (iter *PageStatisticsIterator) HasNext() bool {
	if !iter.nextPageCalled { // 第一次调用需要特殊对待
		return iter.lastPageListResult.ItemCount > 0
	}

	return iter.lastPageListResult.ItemCount >= PageListPageSize
}

func (iter *PageStatisticsIterator) NextPage() (statisticsList []PageStatistics, err error) {
	if !iter.nextPageCalled { // 第一次调用需要特殊对待
		iter.nextPageCalled = true

		statisticsList = iter.lastPageListResult.Data.PageStatisticsList
		return
	}

	rslt, err := PageList(iter.clt, iter.date, iter.nextPageIndex)
	if err != nil {
		return
	}

	iter.nextPageIndex++
	iter.lastPageListResult = rslt

	statisticsList = rslt.Data.PageStatisticsList
	return
}

func NewPageStatisticsIterator(clt *core.Client, date int64, pageIndex int) (iter *PageStatisticsIterator, err error) {
	// 逻辑上相当于第一次调用 PageStatisticsIterator.NextPage, 因为第一次调用 PageStatisticsIterator.HasNext 需要数据支撑, 所以提前获取了数据

	rslt, err := PageList(clt, date, pageIndex)
	if err != nil {
		return
	}

	iter = &PageStatisticsIterator{
		clt: clt,

		date:          date,
		nextPageIndex: pageIndex + 1,

		lastPageListResult: rslt,
		nextPageCalled:     false,
	}
	return
}
