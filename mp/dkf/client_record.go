// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package dkf

import (
	"errors"

	"github.com/chanxuehong/wechat/mp"
)

// 一条聊天记录
type Record struct {
	Worker string `json:"worker"` // 客服账号
	OpenId string `json:"openid"` // 用户的标识，对当前公众号唯一

	// 操作ID（会话状态）:
	// 1000	 创建未接入会话
	// 1001	 接入会话
	// 1002	 主动发起会话
	// 1004	 关闭会话
	// 1005	 抢接会话
	// 2001	 公众号收到消息
	// 2002	 客服发送消息
	// 2003	 客服收到消息
	OperCode  int    `json:"opercode"`
	Timestamp int64  `json:"time"` // 操作时间，UNIX时间戳
	Text      string `json:"text"` // 聊天记录
}

const (
	RecordPageSizeLimit = 1000 // 客户聊天记录每页最多拉取1000条
)

// 获取客服聊天记录 请求消息结构
type GetRecordRequest struct {
	StartTime int64  `json:"starttime"` // 查询开始时间，UNIX时间戳
	EndTime   int64  `json:"endtime"`   // 查询结束时间，UNIX时间戳，每次查询不能跨日查询
	OpenId    string `json:"openid"`    // 普通用户的标识，对当前公众号唯一
	PageSize  int    `json:"pagesize"`  // 每页大小，每页最多拉取1000条
	PageIndex int    `json:"pageindex"` // 查询第几页，从1开始
}

// 获取客服聊天记录
func (clt Client) GetRecord(request *GetRecordRequest) (recordList []Record, err error) {
	if request == nil {
		err = errors.New("nil request")
		return
	}

	var result struct {
		mp.Error
		RecordList []Record `json:"recordlist"`
	}
	// 预分配一定的容量
	if size := request.PageSize; size >= 64 {
		result.RecordList = make([]Record, 0, 64)
	} else {
		result.RecordList = make([]Record, 0, size)
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/customservice/getrecord?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	recordList = result.RecordList
	return
}

// 聊天记录遍历器.
//
//  iter, err := Client.RecordIterator(request)
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
	lastGetRecordRequest *GetRecordRequest // 上一次查询的 request
	lastGetRecordResult  []Record          // 上一次查询的 result

	wechatClient   Client // 关联的微信 Client
	nextPageCalled bool   // NextPage() 是否调用过
}

func (iter *RecordIterator) HasNext() bool {
	if !iter.nextPageCalled { // 还没有调用 NextPage(), 从创建的时候获取的数据来判断
		return len(iter.lastGetRecordResult) > 0
	}

	// 如果上一次读取的数据等于 PageSize, 则"可能"还有数据; 否则肯定是没有数据了.
	return len(iter.lastGetRecordResult) == iter.lastGetRecordRequest.PageSize
}

func (iter *RecordIterator) NextPage() (records []Record, err error) {
	if !iter.nextPageCalled { // 还没有调用 NextPage(), 从创建的时候获取的数据中获取
		records = iter.lastGetRecordResult
		iter.nextPageCalled = true
		return
	}

	// 不是第一次调用的都要从服务器拉取数据
	iter.lastGetRecordRequest.PageIndex++
	records, err = iter.wechatClient.GetRecord(iter.lastGetRecordRequest)
	if err != nil {
		iter.lastGetRecordRequest.PageIndex-- //
		return
	}

	iter.lastGetRecordResult = records
	return
}

// 获取聊天记录遍历器.
func (clt Client) RecordIterator(request *GetRecordRequest) (iter *RecordIterator, err error) {
	records, err := clt.GetRecord(request)
	if err != nil {
		return
	}

	iter = &RecordIterator{
		lastGetRecordRequest: request,
		lastGetRecordResult:  records,
		wechatClient:         clt,
		nextPageCalled:       false,
	}
	return
}
