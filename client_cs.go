package wechat

import (
	"errors"
	"github.com/chanxuehong/wechat/cs"
)

// 获取客服聊天记录
func (c *Client) CSRecordGet(request *cs.RecordGetRequest) ([]cs.Record, error) {
	if request == nil {
		return nil, errors.New("request == nil")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}
	_url := clientCSRecordGetURL(token)

	var result struct {
		RecordList []cs.Record `json:"recordlist"`
		Error
	}
	if err = c.postJSON(_url, request, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return result.RecordList, nil
}

// 该结构实现了 github.com/chanxuehong/wechat/cs.RecordIterator 接口
type csRecordIterator struct {
	recordGetRequest  *cs.RecordGetRequest // 对于 NextPage() 表示当前的 request
	recordGetResponse []cs.Record          // 对于 HasNext() 表示上一次的 response

	wechatClient   *Client // 关联的微信 Client
	nextPageCalled bool    // NextPage() 是否调用过
}

func (iter *csRecordIterator) HasNext() bool {
	// 第一批数据不需要通过 NextPage() 来获取, 因为在创建这个对象的时候就获取了;
	// 后续的数据都要通过 NextPage() 来获取.
	if !iter.nextPageCalled {
		return len(iter.recordGetResponse) > 0
	}
	// 如果当前读取的数据等于 PageSize, 则有可能还有数据; 否则肯定是没有数据了.
	return len(iter.recordGetResponse) == iter.recordGetRequest.PageSize
}
func (iter *csRecordIterator) NextPage() ([]cs.Record, error) {
	// 第一次调用 NextPage(), 因为在创建这个对象的时候已经获取了数据, 所以直接返回.
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		iter.recordGetRequest.PageIndex++ // 为下一页准备数据
		return iter.recordGetResponse, nil
	}

	// 不是第一次调用的都要从服务器拉取数据
	resp, err := iter.wechatClient.CSRecordGet(iter.recordGetRequest)
	if err != nil {
		return nil, err
	}

	iter.recordGetResponse = resp     // 覆盖老数据
	iter.recordGetRequest.PageIndex++ // 为下一页准备数据
	return resp, nil
}

// 聊天记录遍历器
func (c *Client) CSRecordIterator(queryRequest *cs.RecordGetRequest) (cs.RecordIterator, error) {
	// CSRecordGet 会做参数检查, 这里就不用了
	resp, err := c.CSRecordGet(queryRequest)
	if err != nil {
		return nil, err
	}
	var iter csRecordIterator
	iter.recordGetRequest = queryRequest
	iter.recordGetResponse = resp
	iter.wechatClient = c
	return &iter, nil
}
