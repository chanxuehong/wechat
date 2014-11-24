// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"errors"
	"fmt"

	"github.com/chanxuehong/wechat/mp/user"
)

// 创建分组
func (c *Client) UserGroupCreate(name string) (_group *user.Group, err error) {
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
		Error

		Group struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"group"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := userGroupCreateURL(token)

	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		_group = &user.Group{
			Id:   result.Group.Id,
			Name: result.Group.Name,
		}
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

// 查询所有分组
func (c *Client) UserGroupGet() (groups []user.Group, err error) {
	var result = struct {
		Error
		Groups []user.Group `json:"groups"`
	}{
		Groups: make([]user.Group, 0, 16),
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := userGroupGetURL(token)

	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		groups = result.Groups
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

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := userGroupRenameURL(token)

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

// 查询用户所在分组
func (c *Client) UserInWhichGroup(openid string) (groupid int64, err error) {
	if len(openid) == 0 {
		err = errors.New(`openid == ""`)
		return
	}

	var request = struct {
		OpenId string `json:"openid"`
	}{
		OpenId: openid,
	}

	var result struct {
		Error
		GroupId int64 `json:"groupid"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := userInWhichGroupURL(token)

	if err = c.postJSON(url_, request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		groupid = result.GroupId
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

// 移动用户分组
func (c *Client) UserMoveToGroup(openid string, toGroupId int64) (err error) {
	if len(openid) == 0 {
		err = errors.New(`openid == ""`)
		return
	}

	var request = struct {
		OpenId    string `json:"openid"`
		ToGroupId int64  `json:"to_groupid"`
	}{
		OpenId:    openid,
		ToGroupId: toGroupId,
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := userMoveToGroupURL(token)

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

// 开发者可以通过该接口对指定用户设置备注名
//  NOTE: 该接口暂时开放给微信认证的服务号
func (c *Client) UserUpdateRemark(openId, remark string) (err error) {
	if len(openId) == 0 {
		err = errors.New(`openId == ""`)
		return
	}

	var request = struct {
		OpenId string `json:"openid"`
		Remark string `json:"remark"`
	}{
		OpenId: openId,
		Remark: remark,
	}

	var result Error

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := userUpdateRemarkURL(token)

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

// 获取用户基本信息, 如果用户没有订阅公众号, 返回 user.ErrNotSubscribe 错误.
//  lang 可能的取值是 zh_CN, zh_TW, en; 如果留空 "" 则默认为 zh_CN.
func (c *Client) UserInfo(openid string, lang string) (userinfo *user.UserInfo, err error) {
	if openid == "" {
		err = errors.New(`openid == ""`)
		return
	}

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
		Error
		Subscribe int `json:"subscribe"` // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
		user.UserInfo
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := userInfoURL(token, openid, lang)

	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		if result.Subscribe == 0 {
			err = user.ErrNotSubscribe
			return
		}
		userinfo = &result.UserInfo
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

// 获取关注者列表, 每次最多能获取 10000 个用户, 如果 beginOpenId == "" 则表示从头获取
func (c *Client) UserList(beginOpenId string) (data *user.UserListResult, err error) {
	var result struct {
		Error
		user.UserListResult
	}
	result.UserListResult.Data.OpenId = make([]string, 0, 256)

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := userGetURL(token, beginOpenId)

	if err = c.getJSON(url_, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		data = &result.UserListResult
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

// 该结构实现了 user.UserIterator 接口
type userIterator struct {
	lastUserListData *user.UserListResult // 最近一次获取的用户数据

	wechatClient   *Client // 关联的微信 Client
	nextPageCalled bool    // NextPage() 是否调用过
}

func (iter *userIterator) Total() int {
	return iter.lastUserListData.TotalCount
}

func (iter *userIterator) HasNext() bool {
	if !iter.nextPageCalled { // 还没有调用 NextPage(), 从创建的时候获取的数据来判断
		return iter.lastUserListData.GotCount > 0
	}

	// 已经调用过 NextPage(), 根据 next_openid 字段是否为空来判断
	//
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
	return len(iter.lastUserListData.NextOpenId) != 0 &&
		iter.lastUserListData.GotCount == user.UserPageSizeLimit
}

func (iter *userIterator) NextPage() (openids []string, err error) {
	if !iter.nextPageCalled { // 还没有调用 NextPage(), 从创建的时候获取的数据中获取
		iter.nextPageCalled = true
		openids = iter.lastUserListData.Data.OpenId
		return
	}

	// 不是第一次调用的都要从服务器拉取数据
	data, err := iter.wechatClient.UserList(iter.lastUserListData.NextOpenId)
	if err != nil {
		return
	}

	iter.lastUserListData = data // 覆盖老数据
	openids = data.Data.OpenId
	return
}

// 关注用户遍历器, 如果 beginOpenId == "" 则表示从头遍历
func (c *Client) UserIterator(beginOpenId string) (iter user.UserIterator, err error) {
	data, err := c.UserList(beginOpenId)
	if err != nil {
		return
	}

	iter = &userIterator{
		lastUserListData: data,
		wechatClient:     c,
	}
	return
}
