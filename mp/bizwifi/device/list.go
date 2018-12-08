package device

import (
	"github.com/chanxuehong/wechat/internal/util"
	"github.com/chanxuehong/wechat/mp/core"
)

type SearchQuery struct {
	ShopId    *int64 `json:"shop_id,omitempty"`   // 根据门店id查询
	PageIndex int    `json:"pageindex,omitempty"` // 分页下标，默认从1开始
	PageSize  int    `json:"pagesize,omitempty"`  // 每页的个数，默认10个，最大20个
}

func NewSearchQuery1(PageIndex, PageSize int) *SearchQuery {
	return &SearchQuery{
		PageIndex: PageIndex,
		PageSize:  PageSize,
	}
}

func NewSearchQuery2(ShopId int64, PageIndex, PageSize int) *SearchQuery {
	return &SearchQuery{
		ShopId:    util.Int64(ShopId),
		PageIndex: PageIndex,
		PageSize:  PageSize,
	}
}

type ListResult struct {
	PageIndex int `json:"pageindex"` // 分页下标
	PageCount int `json:"pagecount"` // 分页页数

	TotalCount int `json:"totalcount"` // 总数
	ItemCount  int `json:"itemcount"`  // 当前页列表大小

	Records []Device `json:"records"` // 当前页列表数组
}

type Device struct {
	ShopId int64  `json:"shop_id"` // 门店ID
	SSID   string `json:"ssid"`    // 连网设备ssid
	BSSID  string `json:"bssid"`   // 无线MAC地址
}

// 查询设备.
func List(clt *core.Client, query *SearchQuery) (rslt *ListResult, err error) {
	var result struct {
		core.Error
		ListResult `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/bizwifi/device/list?access_token="
	if err = clt.PostJSON(incompleteURL, query, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}

	result.ListResult.ItemCount = len(result.ListResult.Records)
	rslt = &result.ListResult
	return
}

// DeviceIterator
//
//  iter, err := NewDeviceIterator(*core.Client, *SearchQuery)
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
type DeviceIterator struct {
	clt *core.Client

	nextQuery *SearchQuery

	lastListResult *ListResult // 最近一次获取的数据
	nextPageCalled bool        // NextPage() 是否调用过
}

func (iter *DeviceIterator) TotalCount() int {
	return iter.lastListResult.TotalCount
}

func (iter *DeviceIterator) HasNext() bool {
	if !iter.nextPageCalled { // 第一次调用需要特殊对待
		return iter.lastListResult.ItemCount > 0 ||
			iter.nextQuery.PageIndex <= iter.lastListResult.PageCount
	}

	return iter.nextQuery.PageIndex <= iter.lastListResult.PageCount
}

func (iter *DeviceIterator) NextPage() (records []Device, err error) {
	if !iter.nextPageCalled { // 第一次调用需要特殊对待
		iter.nextPageCalled = true

		records = iter.lastListResult.Records
		return
	}

	rslt, err := List(iter.clt, iter.nextQuery)
	if err != nil {
		return
	}

	iter.nextQuery.PageIndex++
	iter.lastListResult = rslt

	records = rslt.Records
	return
}

func NewDeviceIterator(clt *core.Client, query *SearchQuery) (iter *DeviceIterator, err error) {
	// 逻辑上相当于第一次调用 DeviceIterator.NextPage, 因为第一次调用 DeviceIterator.HasNext 需要数据支撑, 所以提前获取了数据

	rslt, err := List(clt, query)
	if err != nil {
		return
	}

	query.PageIndex++

	iter = &DeviceIterator{
		clt: clt,

		nextQuery: query,

		lastListResult: rslt,
		nextPageCalled: false,
	}
	return
}
