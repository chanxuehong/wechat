// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/user"
)

// 创建分组
func (c *Client) UserGroupCreate(name string) (_group user.Group, err error) {
	if len(name) == 0 {
		err = errors.New(`name == ""`)
		return
	}

	var request struct {
		Group struct {
			Name string `json:"name"`
		} `json:"group"`
	}
	request.Group.Name = name

	var result struct {
		Group struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"group"`
		Error
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := userGroupCreateURL(token)
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		_group = user.Group{
			Id:   result.Group.Id,
			Name: result.Group.Name,
		}
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result.Error
		return
	}
}

// 查询所有分组
func (c *Client) UserGroupGet() (groups []user.Group, err error) {
	var result = struct {
		Groups []user.Group `json:"groups"`
		Error
	}{
		Groups: make([]user.Group, 0, 16),
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := userGroupGetURL(token)
	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		groups = result.Groups
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result.Error
		return
	}
}

// 修改分组名
func (c *Client) UserGroupRename(groupid int64, name string) (err error) {
	if len(name) == 0 {
		return errors.New(`name == ""`)
	}

	var request struct {
		Group struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"group"`
	}
	request.Group.Id = groupid
	request.Group.Name = name

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := userGroupRenameURL(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result
		return
	}
}

// 查询用户所在分组
func (c *Client) UserInWhichGroup(openid string) (groupid int64, err error) {
	var request = struct {
		OpenId string `json:"openid"`
	}{
		OpenId: openid,
	}

	var result struct {
		GroupId int64 `json:"groupid"`
		Error
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := userInWhichGroupURL(token)
	if err = c.postJSON(_url, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		groupid = result.GroupId
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result.Error
		return
	}
}

// 移动用户分组
func (c *Client) UserMoveToGroup(openid string, toGroupId int64) (err error) {
	var request = struct {
		OpenId    string `json:"openid"`
		ToGroupId int64  `json:"to_groupid"`
	}{
		OpenId:    openid,
		ToGroupId: toGroupId,
	}

	var result Error

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := userMoveToGroupURL(token)
	if err = c.postJSON(_url, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result
		return
	}
}

// 获取用户基本信息.
//  lang 可能的取值是 zh_CN, zh_TW, en; 如果留空 "" 则默认为 zh_CN.
func (c *Client) UserInfo(openid string, lang string) (userinfo *user.UserInfo, err error) {
	switch lang {
	case "":
		lang = user.Language_zh_CN
	case user.Language_zh_CN, user.Language_zh_TW, user.Language_en:
	default:
		err = fmt.Errorf("lang 必须是 \"\",%s,%s,%s 其中之一",
			user.Language_zh_CN, user.Language_zh_TW, user.Language_en)
		return
	}

	var result struct {
		Subscribe int `json:"subscribe"` // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
		user.UserInfo
		Error
	}

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := userInfoURL(token, openid, lang)
	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		if result.Subscribe == 0 {
			err = fmt.Errorf("该用户 %s 没有订阅这个公众号", openid)
			return
		}
		userinfo = &result.UserInfo
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result.Error
		return
	}
}

// 获取关注者返回的数据结构
type userGetResponse struct {
	TotalCount int `json:"total"` // 关注该公众账号的总用户数
	GetCount   int `json:"count"` // 拉取的OPENID个数，最大值为10000
	Data       struct {
		OpenId []string `json:"openid"`
	} `json:"data"` // 列表数据，OPENID的列表
	// 拉取列表的后一个用户的OPENID, 如果 next_openid == "" 则表示没有了用户数据
	NextOpenId string `json:"next_openid"`
}

// 获取关注者列表, 如果 beginOpenId == "" 则表示从头遍历
func (c *Client) userGet(beginOpenId string) (resp *userGetResponse, err error) {
	var result struct {
		userGetResponse
		Error
	}
	result.userGetResponse.Data.OpenId = make([]string, 0, 256)

	hasRetry := false
RETRY:
	token, err := c.Token()
	if err != nil {
		return
	}
	_url := userGetURL(token, beginOpenId)
	if err = c.getJSON(_url, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		resp = &result.userGetResponse
		return

	case errCodeTimeout:
		if !hasRetry {
			hasRetry = true
			timeoutRetryWait()
			goto RETRY
		}
		fallthrough

	default:
		err = result.Error
		return
	}
}

// 该结构实现了 user.UserIterator 接口
type userGetIterator struct {
	userGetResponse *userGetResponse // 最近一次获取的用户数据

	wechatClient   *Client // 关联的微信 Client
	nextPageCalled bool    // NextPage() 是否调用过
}

func (iter *userGetIterator) Total() int {
	return iter.userGetResponse.TotalCount
}

func (iter *userGetIterator) HasNext() bool {
	// 第一批数据不需要通过 NextPage() 来获取, 因为在创建这个对象的时候就获取了;
	if !iter.nextPageCalled {
		return iter.userGetResponse.GetCount > 0
	}

	// 后面的判断都要根据上一次获取的数据来判断

	// 跟文档的描述貌似有点不一样, 即使后续没有用户, 貌似 next_openid 还是不为空!
	// 所以增加了一个判断 iter.userGetResponse.GetCount == user.UserPageCountLimit
	//
	// 200	OK
	// Connection: keep-alive
	// Date: Sat, 28 Jun 2014 07:00:10 GMT
	// Server: nginx/1.4.4
	// Content-Type: application/json; encoding=utf-8
	// Content-Length: 117
	// {
	//     "total": 1,
	//     "count": 1,
	//     "data": {
	//         "openid": [
	//             "os-IKuHd9pJ6xsn4mS7GyL4HxqI4"
	//         ]
	//     },
	//     "next_openid": "os-IKuHd9pJ6xsn4mS7GyL4HxqI4"
	// }
	return iter.userGetResponse.GetCount == user.UserPageCountLimit &&
		len(iter.userGetResponse.NextOpenId) != 0
}
func (iter *userGetIterator) NextPage() (openids []string, err error) {
	// 第一次调用 NextPage(), 因为在创建这个对象的时候已经获取了数据, 所以直接返回.
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		openids = iter.userGetResponse.Data.OpenId
		return
	}

	// 不是第一次调用的都要从服务器拉取数据
	resp, err := iter.wechatClient.userGet(iter.userGetResponse.NextOpenId)
	if err != nil {
		return
	}

	iter.userGetResponse = resp // 覆盖老数据
	openids = resp.Data.OpenId
	return
}

// 关注用户遍历器, 如果 beginOpenId == "" 则表示从头遍历
func (c *Client) UserIterator(beginOpenId string) (iter user.UserIterator, err error) {
	resp, err := c.userGet(beginOpenId)
	if err != nil {
		return
	}

	iter = &userGetIterator{
		userGetResponse: resp,
		wechatClient:    c,
	}
	return
}
