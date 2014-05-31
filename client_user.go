package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/user"
	"io/ioutil"
)

// 创建分组
func (c *Client) UserGroupCreate(name string) (*user.Group, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	var request struct {
		Group struct {
			Name string `json:"name"`
		} `json:"group"`
	}

	request.Group.Name = name

	jsonData, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	_url := userGroupCreateUrlPrefix + token
	resp, err := commonHttpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Group struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"group"`
		Error
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, &result.Error
	}

	var group user.Group
	group.Id = result.Group.Id
	group.Name = result.Group.Name
	return &group, nil
}

// 查询所有分组
func (c *Client) UserGroupGet() ([]user.Group, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	_url := userGroupGetUrlPrefix + token
	resp, err := commonHttpClient.Get(_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Groups []user.Group `json:"groups"`
		Error
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return result.Groups, nil
}

// 修改分组名
func (c *Client) UserGroupRename(groupid int, name string) (err error) {
	token, err := c.Token()
	if err != nil {
		return
	}

	var request struct {
		Group struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"group"`
	}

	request.Group.Id = groupid
	request.Group.Name = name

	jsonData, err := json.Marshal(&request)
	if err != nil {
		return
	}

	_url := userGroupRenameUrlPrefix + token
	resp, err := commonHttpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result Error
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		return &result
	}

	return
}

// 查询用户所在分组
func (c *Client) UserInWhichGroup(openid string) (groupid int, err error) {
	token, err := c.Token()
	if err != nil {
		return
	}

	var request = struct {
		OpenId string `json:"openid"`
	}{OpenId: openid}

	jsonData, err := json.Marshal(&request)
	if err != nil {
		return
	}

	_url := userInWhichGroupUrlPrefix + token
	resp, err := commonHttpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result struct {
		GroupId int `json:"groupid"`
		Error
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	groupid = result.GroupId
	return
}

// 移动用户分组
func (c *Client) UserMoveToGroup(openid string, toGroupId int) (err error) {
	token, err := c.Token()
	if err != nil {
		return
	}

	var request = struct {
		OpenId    string `json:"openid"`
		ToGroupId int    `json:"to_groupid"`
	}{
		OpenId:    openid,
		ToGroupId: toGroupId,
	}

	jsonData, err := json.Marshal(&request)
	if err != nil {
		return
	}

	_url := userMoveToGroupUrlPrefix + token
	resp, err := commonHttpClient.Post(_url, postJSONContentType, bytes.NewReader(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result Error
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		return &result
	}

	return
}

// 获取用户基本信息.
// lang 可能的取值是 zh_CN, zh_TW, en; 如果留空 "" 则默认为 zh_CN.
func (c *Client) UserInfo(openid string, lang string) (*user.UserInfo, error) {
	switch lang {
	case "", user.Language_zh_CN, user.Language_zh_TW, user.Language_en:
	default:
		return nil, errors.New(`lang 必须是 "", zh_CN, zh_TW, en 之一`)
	}

	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	_url := fmt.Sprintf(userInfoUrlFormat, token, openid, lang)
	resp, err := commonHttpClient.Get(_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		// 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
		Subscribe int `json:"subscribe"`
		user.UserInfo
		Error
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	if result.Subscribe == 0 {
		return nil, fmt.Errorf("该用户 %s 没有订阅这个公众号", openid)
	}
	return &result.UserInfo, nil
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
func (c *Client) userGet(beginOpenId string) (*userGetResponse, error) {
	token, err := c.Token()
	if err != nil {
		return nil, err
	}

	var _url string
	if beginOpenId == "" {
		_url = userGetUrlPrefix + token
	} else {
		_url = userGetUrlPrefix + token + "&next_openid=" + beginOpenId
	}

	resp, err := commonHttpClient.Get(_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		userGetResponse
		Error
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, &result.Error
	}
	return &result.userGetResponse, nil
}

// 该结构实现了 user.UserIterator 接口
type userGetIterator struct {
	userGetResponse *userGetResponse // NextPage() 返回的数据

	wechatClient   *Client // 关联的微信 Client
	nextPageCalled bool    // NextPage() 是否调用过
}

func (iter *userGetIterator) Total() int {
	return iter.userGetResponse.TotalCount
}
func (iter *userGetIterator) HasNext() bool {
	// 第一批数据不需要通过 NextPage() 来获取, 因为在创建这个对象的时候就获取了;
	// 后续的数据都要通过 NextPage() 来获取, 所以要通过上一次的 NextOpenId 来判断了.
	if !iter.nextPageCalled {
		return iter.userGetResponse.GetCount != 0
	}
	return iter.userGetResponse.NextOpenId != ""
}
func (iter *userGetIterator) NextPage() ([]string, error) {
	// 第一次调用 NextPage(), 因为在创建这个对象的时候已经获取了数据, 所以直接返回.
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		return iter.userGetResponse.Data.OpenId, nil
	}

	// 不是第一次调用的都要从服务器拉取数据
	resp, err := iter.wechatClient.userGet(iter.userGetResponse.NextOpenId)
	if err != nil {
		return nil, err
	}

	iter.userGetResponse = resp // 覆盖老数据
	return resp.Data.OpenId, nil
}

// 关注用户遍历器, 如果 beginOpenId == "" 则表示从头遍历
func (c *Client) UserIterator(beginOpenId string) (user.UserIterator, error) {
	resp, err := c.userGet(beginOpenId)
	if err != nil {
		return nil, err
	}
	var iter userGetIterator
	iter.userGetResponse = resp
	iter.wechatClient = c
	return &iter, nil
}
