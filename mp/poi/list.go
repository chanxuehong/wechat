package poi

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp/core"
)

type ListResult struct {
	TotalCount int   `json:"total_count"`   // 门店总数量
	ItemCount  int   `json:"item_count"`    // 本次调用获取的门店数量
	List       []Poi `json:"business_list"` // 本次调用获取的门店列表
}

// List 查询门店列表.
//  begin: 开始位置，0 即为从第一条开始查询
//  limit: 返回数据条数，最大允许50，默认为20
func List(clt *core.Client, begin, limit int) (rslt *ListResult, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/poi/getpoilist?access_token="

	if begin < 0 {
		err = fmt.Errorf("invalid begin: %d", begin)
		return
	}
	if limit <= 0 {
		err = fmt.Errorf("invalid limit: %d", limit)
		return
	}

	var request = struct {
		Begin int `json:"begin"`
		Limit int `json:"limit"`
	}{
		Begin: begin,
		Limit: limit,
	}
	var result struct {
		core.Error
		ListResult
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	result.ListResult.ItemCount = len(result.ListResult.List)
	rslt = &result.ListResult
	return
}

// =====================================================================================================================

// PoiIterator
//
//  iter, err := NewPoiIterator(clt, 0, 10)
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
type PoiIterator struct {
	clt *core.Client

	nextOffset int
	count      int

	lastListResult *ListResult
	nextPageCalled bool
}

func (iter *PoiIterator) TotalCount() int {
	return iter.lastListResult.TotalCount
}

func (iter *PoiIterator) HasNext() bool {
	if !iter.nextPageCalled {
		return iter.lastListResult.ItemCount > 0 || iter.nextOffset < iter.lastListResult.TotalCount
	}
	return iter.nextOffset < iter.lastListResult.TotalCount
}

func (iter *PoiIterator) NextPage() (list []Poi, err error) {
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		list = iter.lastListResult.List
		return
	}

	rslt, err := List(iter.clt, iter.nextOffset, iter.count)
	if err != nil {
		return
	}

	iter.lastListResult = rslt
	iter.nextOffset += rslt.ItemCount

	list = rslt.List
	return
}

func NewPoiIterator(clt *core.Client, begin, limit int) (iter *PoiIterator, err error) {
	// 逻辑上相当于第一次调用 PoiIterator.NextPage,
	// 因为第一次调用 PoiIterator.HasNext 需要数据支撑, 所以提前获取了数据
	rslt, err := List(clt, begin, limit)
	if err != nil {
		return
	}

	iter = &PoiIterator{
		clt: clt,

		nextOffset: begin + rslt.ItemCount,
		count:      limit,

		lastListResult: rslt,
		nextPageCalled: false,
	}
	return
}
