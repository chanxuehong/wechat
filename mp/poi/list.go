<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

=======
>>>>>>> github/v2
package poi

import (
	"fmt"

<<<<<<< HEAD
	"github.com/chanxuehong/wechat/mp"
)

type PoiListResult struct {
	TotalCount int   `json:"total_count"`   // 门店总数量
	ItemCount  int   `json:"item_count"`    // 本次调用获取的门店数量
	PoiList    []Poi `json:"business_list"` // 本次调用获取的门店列表
}

// 查询门店列表.
//  begin: 开始位置, 0 即为从第一条开始查询
//  limit: 返回数据条数, 最大允许50, 默认为20
func (clt *Client) PoiList(begin, limit int) (rslt *PoiListResult, err error) {
=======
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

>>>>>>> github/v2
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
<<<<<<< HEAD

	var result struct {
		mp.Error
		PoiListResult
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/poi/getpoilist?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}

	result.PoiListResult.ItemCount = len(result.PoiListResult.PoiList)
	rslt = &result.PoiListResult
	return
}

// PoiIterator
//
//  iter, err := Client.PoiIterator(0, 10)
=======
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
>>>>>>> github/v2
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
<<<<<<< HEAD
	clt *Client // 关联的微信 Client

	nextOffset int // 下一次获取数据时的 offset
	count      int // 步长

	lastPoiListResult *PoiListResult // 最近一次获取的数据
	nextPageHasCalled bool           // NextPage() 是否调用过
}

func (iter *PoiIterator) TotalCount() int {
	return iter.lastPoiListResult.TotalCount
}

func (iter *PoiIterator) HasNext() bool {
	if !iter.nextPageHasCalled { // 第一次调用需要特殊对待
		return iter.lastPoiListResult.ItemCount > 0 ||
			iter.nextOffset < iter.lastPoiListResult.TotalCount
	}

	return iter.nextOffset < iter.lastPoiListResult.TotalCount
}

func (iter *PoiIterator) NextPage() (poiList []Poi, err error) {
	if !iter.nextPageHasCalled { // 第一次调用需要特殊对待
		iter.nextPageHasCalled = true

		poiList = iter.lastPoiListResult.PoiList
		return
	}

	rslt, err := iter.clt.PoiList(iter.nextOffset, iter.count)
=======
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
>>>>>>> github/v2
	if err != nil {
		return
	}

<<<<<<< HEAD
	iter.nextOffset += rslt.ItemCount
	iter.lastPoiListResult = rslt

	poiList = rslt.PoiList
	return
}

func (clt *Client) PoiIterator(begin, limit int) (iter *PoiIterator, err error) {
	// 逻辑上相当于第一次调用 PoiIterator.NextPage, 因为第一次调用 PoiIterator.HasNext 需要数据支撑, 所以提前获取了数据

	rslt, err := clt.PoiList(begin, limit)
=======
	iter.lastListResult = rslt
	iter.nextOffset += rslt.ItemCount

	list = rslt.List
	return
}

func NewPoiIterator(clt *core.Client, begin, limit int) (iter *PoiIterator, err error) {
	// 逻辑上相当于第一次调用 PoiIterator.NextPage,
	// 因为第一次调用 PoiIterator.HasNext 需要数据支撑, 所以提前获取了数据
	rslt, err := List(clt, begin, limit)
>>>>>>> github/v2
	if err != nil {
		return
	}

	iter = &PoiIterator{
		clt: clt,

		nextOffset: begin + rslt.ItemCount,
		count:      limit,

<<<<<<< HEAD
		lastPoiListResult: rslt,
		nextPageHasCalled: false,
=======
		lastListResult: rslt,
		nextPageCalled: false,
>>>>>>> github/v2
	}
	return
}
