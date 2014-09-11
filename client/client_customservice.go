// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"github.com/chanxuehong/wechat/customservice"
)

// 获取客服聊天记录
func (c *Client) CustomServiceRecordGet(request *customservice.RecordGetRequest) (recordList []customservice.Record, err error) {
	if request == nil {
		err = errors.New("request == nil")
		return
	}

	var result struct {
		RecordList []customservice.Record `json:"recordlist"`
		Error
	}
	// 预分配一定的容量
	if size := request.PageSize; size >= 64 {
		result.RecordList = make([]customservice.Record, 0, 64)
	} else {
		result.RecordList = make([]customservice.Record, 0, size)
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := customServiceRecordGetURL(token)

	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		recordList = result.RecordList
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 该结构实现了 github.com/chanxuehong/wechat/customservice.RecordIterator 接口
type customServiceRecordIterator struct {
	recordGetRequest  *customservice.RecordGetRequest // 上一次查询的 request
	recordGetResponse []customservice.Record          // 上一次查询的 response

	wechatClient   *Client // 关联的微信 Client
	nextPageCalled bool    // NextPage() 是否调用过
}

func (iter *customServiceRecordIterator) HasNext() bool {
	// 第一批数据不需要通过 NextPage() 来获取, 在创建这个对象的时候就获取了;
	if !iter.nextPageCalled {
		return len(iter.recordGetResponse) > 0
	}
	// 如果上一次读取的数据等于 PageSize, 则*有可能*还有数据; 否则肯定是没有数据了.
	return len(iter.recordGetResponse) == iter.recordGetRequest.PageSize
}

func (iter *customServiceRecordIterator) NextPage() (records []customservice.Record, err error) {
	// 第一次调用 NextPage(), 因为在创建这个对象的时候已经获取了数据, 所以直接返回.
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		records = iter.recordGetResponse
		return
	}

	// 不是第一次调用的都要从服务器拉取数据
	iter.recordGetRequest.PageIndex++
	records, err = iter.wechatClient.CustomServiceRecordGet(iter.recordGetRequest)
	if err != nil {
		return
	}

	iter.recordGetResponse = records
	return
}

// 聊天记录遍历器
func (c *Client) CustomServiceRecordIterator(request *customservice.RecordGetRequest) (iter customservice.RecordIterator, err error) {
	resp, err := c.CustomServiceRecordGet(request)
	if err != nil {
		return
	}

	iter = &customServiceRecordIterator{
		recordGetRequest:  request,
		recordGetResponse: resp,
		wechatClient:      c,
	}
	return
}

// 获取客服基本信息
func (c *Client) CustomServiceKfList() (kfList []customservice.KfInfo, err error) {
	var result struct {
		KfList []customservice.KfInfo `json:"kf_list"`
		Error
	}
	// 预分配一定的容量
	result.KfList = make([]customservice.KfInfo, 0, 16)

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := customServiceKfListURL(token)

	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		kfList = result.KfList
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}

// 获取在线客服接待信息
func (c *Client) CustomServiceOnlineKfList() (kfList []customservice.OnlineKfInfo, err error) {
	var result struct {
		KfList []customservice.OnlineKfInfo `json:"kf_online_list"`
		Error
	}
	// 预分配一定的容量
	result.KfList = make([]customservice.OnlineKfInfo, 0, 16)

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := customServiceOnlineKfListURL(token)

	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		kfList = result.KfList
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = &result.Error
		return
	}
}
