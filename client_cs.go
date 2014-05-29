package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/chanxuehong/wechat/cs"
	"io/ioutil"
	"net/http"
)

// 获取客服聊天记录
func (c *Client) CSGetRecord(request *cs.RecordRequest) (*cs.RecordResponse, error) {
	if request == nil {
		return nil, errors.New("request == nil")
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	_url := csGetRecordUrlPrefix + token
	resp, err := http.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		cs.RecordResponse
		Error
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.RecordResponse, nil
}

// 该结构实现了 cs.RecordIterator 接口
type csRecordIterator struct {
	recordRequest  *cs.RecordRequest
	recordResponse *cs.RecordResponse

	wechatClient   *Client // 关联的微信 Client
	nextPageCalled bool    // NextPage() 是否调用过
}

func (iter *csRecordIterator) HasNext() bool {
	// 第一批数据不需要通过 NextPage() 来获取, 因为在创建这个对象的时候就获取了;
	// 后续的数据都要通过 NextPage() 来获取.
	if !iter.nextPageCalled {
		return len(iter.recordResponse.RecordList) != 0
	}
	// 如果当前读取的数据等于 PageSize, 则有可能还有数据; 否则肯定是没有数据了.
	return len(iter.recordResponse.RecordList) == iter.recordRequest.PageSize
}
func (iter *csRecordIterator) NextPage() ([]cs.RecordItem, error) {
	// 第一次调用 NextPage(), 因为在创建这个对象的时候已经获取了数据, 所以直接返回.
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		iter.recordRequest.PageIndex++ // 为下一页准备数据
		return iter.recordResponse.RecordList, nil
	}

	// 不是第一次调用的都要从服务器拉取数据
	resp, err := iter.wechatClient.CSGetRecord(iter.recordRequest)
	if err != nil {
		return nil, err
	}

	iter.recordResponse = resp     // 覆盖老数据
	iter.recordRequest.PageIndex++ // 为下一页准备数据
	return resp.RecordList, nil
}

// 聊天记录遍历器
func (c *Client) CSRecordIterator(queryRequest *cs.RecordRequest) (cs.RecordIterator, error) {
	resp, err := c.CSGetRecord(queryRequest)
	if err != nil {
		return nil, err
	}
	var iter csRecordIterator
	iter.recordRequest = queryRequest
	iter.recordResponse = resp
	iter.wechatClient = c
	return &iter, nil
}
