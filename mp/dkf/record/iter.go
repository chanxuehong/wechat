package record

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// RecordIterator
//
//  iter, err := NewRecordIterator(clt, request)
//  if err != nil {
//      // TODO: 增加你的代码
//  }
//
//  for iter.HasNext() {
//      records, err := iter.NextPage()
//      if err != nil {
//          // TODO: 增加你的代码
//      }
//      // TODO: 增加你的代码
//  }
type RecordIterator struct {
	clt *core.Client

	nextGetRequest *GetRequest

	lastGetRecords []Record
	nextPageCalled bool
}

func (iter *RecordIterator) HasNext() bool {
	if !iter.nextPageCalled {
		return len(iter.lastGetRecords) > 0
	}
	return len(iter.lastGetRecords) >= iter.nextGetRequest.PageSize
}

func (iter *RecordIterator) NextPage() (records []Record, err error) {
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		records = iter.lastGetRecords
		return
	}

	records, err = Get(iter.clt, iter.nextGetRequest)
	if err != nil {
		return
	}

	iter.lastGetRecords = records
	iter.nextGetRequest.PageIndex++
	return
}

func NewRecordIterator(clt *core.Client, request *GetRequest) (iter *RecordIterator, err error) {
	// 逻辑上相当于第一次调用 RecordIterator.NextPage,
	// 因为第一次调用 RecordIterator.HasNext 需要数据支撑, 所以提前获取了数据
	records, err := Get(clt, request)
	if err != nil {
		return
	}

	request.PageIndex++

	iter = &RecordIterator{
		clt:            clt,
		nextGetRequest: request,
		lastGetRecords: records,
		nextPageCalled: false,
	}
	return
}
