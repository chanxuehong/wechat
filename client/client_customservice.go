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
func (c *Client) CustomServiceRecordGet(request *customservice.RecordGetRequest) ([]customservice.Record, error) {
	if request == nil {
		return nil, errors.New("request == nil")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := customServiceRecordGetURL(token)

	var result = struct {
		RecordList []customservice.Record `json:"recordlist"`
		Error
	}{
		RecordList: make([]customservice.Record, 0, customservice.RecordPageSizeLimit),
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return result.RecordList, nil
}

// 该结构实现了 github.com/chanxuehong/wechat/customservice.RecordIterator 接口
type csRecordIterator struct {
	recordGetRequest  *customservice.RecordGetRequest // 上一次查询的 request
	recordGetResponse []customservice.Record          // 上一次查询的 response

	wechatClient   *Client // 关联的微信 Client
	nextPageCalled bool    // NextPage() 是否调用过
}

func (iter *csRecordIterator) HasNext() bool {
	// 第一批数据不需要通过 NextPage() 来获取, 在创建这个对象的时候就获取了;
	if !iter.nextPageCalled {
		return len(iter.recordGetResponse) > 0
	}
	// 如果上一次读取的数据等于 PageSize, 则*有可能*还有数据; 否则肯定是没有数据了.
	return len(iter.recordGetResponse) == iter.recordGetRequest.PageSize
}
func (iter *csRecordIterator) NextPage() ([]customservice.Record, error) {
	// 第一次调用 NextPage(), 因为在创建这个对象的时候已经获取了数据, 所以直接返回.
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		return iter.recordGetResponse, nil
	}

	// 不是第一次调用的都要从服务器拉取数据
	iter.recordGetRequest.PageIndex++
	resp, err := iter.wechatClient.CustomServiceRecordGet(iter.recordGetRequest)
	if err != nil {
		return nil, err
	}

	iter.recordGetResponse = resp
	return resp, nil
}

// 聊天记录遍历器
func (c *Client) CustomServiceRecordIterator(queryRequest *customservice.RecordGetRequest) (customservice.RecordIterator, error) {
	// CSRecordGet 会做参数检查, 这里就不用了
	resp, err := c.CustomServiceRecordGet(queryRequest)
	if err != nil {
		return nil, err
	}
	var iter csRecordIterator
	iter.recordGetRequest = queryRequest
	iter.recordGetResponse = resp
	iter.wechatClient = c
	return &iter, nil
}
