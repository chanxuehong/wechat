// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/chanxuehong/wechat/mp/customservice"
)

// 获取客服聊天记录
func (c *Client) CustomServiceRecordGet(request *customservice.RecordGetRequest) (recordList []customservice.Record, err error) {
	if request == nil {
		err = errors.New("request == nil")
		return
	}

	var result struct {
		Error
		RecordList []customservice.Record `json:"recordlist"`
	}
	// 预分配一定的容量
	if size := request.PageSize; size >= 64 {
		result.RecordList = make([]customservice.Record, 0, 64)
	} else {
		result.RecordList = make([]customservice.Record, 0, size)
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceRecordGetURL(token)

	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		recordList = result.RecordList
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
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
	lastRecordGetRequest *customservice.RecordGetRequest // 上一次查询的 request
	lastRecordGetResult  []customservice.Record          // 上一次查询的 result

	wechatClient   *Client // 关联的微信 Client
	nextPageCalled bool    // NextPage() 是否调用过
}

func (iter *customServiceRecordIterator) HasNext() bool {
	if !iter.nextPageCalled { // 还没有调用 NextPage(), 从创建的时候获取的数据来判断
		return len(iter.lastRecordGetResult) > 0
	}

	// 如果上一次读取的数据等于 PageSize, 则***有可能***还有数据; 否则肯定是没有数据了.
	return len(iter.lastRecordGetResult) == iter.lastRecordGetRequest.PageSize
}

func (iter *customServiceRecordIterator) NextPage() (records []customservice.Record, err error) {
	if !iter.nextPageCalled { // 还没有调用 NextPage(), 从创建的时候获取的数据中获取
		iter.nextPageCalled = true
		records = iter.lastRecordGetResult
		return
	}

	// 不是第一次调用的都要从服务器拉取数据
	iter.lastRecordGetRequest.PageIndex++
	records, err = iter.wechatClient.CustomServiceRecordGet(iter.lastRecordGetRequest)
	if err != nil {
		return
	}

	iter.lastRecordGetResult = records
	return
}

// 聊天记录遍历器
func (c *Client) CustomServiceRecordIterator(request *customservice.RecordGetRequest) (iter customservice.RecordIterator, err error) {
	records, err := c.CustomServiceRecordGet(request)
	if err != nil {
		return
	}

	iter = &customServiceRecordIterator{
		lastRecordGetRequest: request,
		lastRecordGetResult:  records,
		wechatClient:         c,
	}
	return
}

// 获取客服基本信息
func (c *Client) CustomServiceKFList() (kfList []customservice.KFInfo, err error) {
	var result struct {
		Error
		KFList []customservice.KFInfo `json:"kf_list"`
	}
	result.KFList = make([]customservice.KFInfo, 0, 16) // 预分配一定的容量

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceKFListURL(token)

	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		kfList = result.KFList
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return
	}
}

// 获取在线客服接待信息
func (c *Client) CustomServiceOnlineKFList() (kfList []customservice.OnlineKFInfo, err error) {
	var result struct {
		Error
		KFList []customservice.OnlineKFInfo `json:"kf_online_list"`
	}
	result.KFList = make([]customservice.OnlineKFInfo, 0, 16) // 预分配一定的容量

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceOnlineKFListURL(token)

	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		kfList = result.KFList
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return
	}
}

// 添加客服账号
//  isPwdPlain: 标识 password 是否为明文格式
func (c *Client) CustomServiceKFAccountAdd(account, nickname, password string, isPwdPlain bool) (err error) {
	if isPwdPlain {
		md5Sum := md5.Sum([]byte(password))
		password = hex.EncodeToString(md5Sum[:])
	}

	request := struct {
		Account  string `json:"kf_account"`
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}{
		Account:  account,
		Nickname: nickname,
		Password: password,
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceKFAccountAddURL(token)

	if err = c.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return
	}
}

// 设置客服信息
//  isPwdPlain: 标识 password 是否为明文格式
func (c *Client) CustomServiceKFAccountSet(account, nickname, password string, isPwdPlain bool) (err error) {
	if isPwdPlain {
		md5Sum := md5.Sum([]byte(password))
		password = hex.EncodeToString(md5Sum[:])
	}

	request := struct {
		Account  string `json:"kf_account"`
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}{
		Account:  account,
		Nickname: nickname,
		Password: password,
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceKFAccountSetURL(token)

	if err = c.postJSON(url_, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return
	}
}

// 删除客服账号
func (c *Client) CustomServiceKFAccountDelete(account string) (err error) {
	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := customServiceKFAccountDeleteURL(token, account)

	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result
		return
	}
}
