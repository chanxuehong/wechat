package tag

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type Tag struct {
	Id        int    `json:"id"`    // tag id
	Name      string `json:"name"`  // tag name
	UserCount int    `json:"count"` // Tag内用户数量
}

// 获取用户列表返回的数据结构
type GetResult struct {
	Count int `json:"count"` // 拉取的OPENID个数, 最大值为10000

	Data struct {
		OpenIdList []string `json:"openid,omitempty"`
	} `json:"data"` // 列表数据, OPENID的列表

	// 拉取列表的最后一个用户的OPENID, 如果 next_openid == "" 则表示没有了用户数据
	NextOpenId string `json:"next_openid"`
}

func Create(clt *core.Client, name string) (tag *Tag, err error) {
	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/create?access_token="
	var request struct {
		Tag struct {
			Name string `json:"name"`
		} `json:"tag"`
	}
	request.Tag.Name = name
	var result struct {
		core.Error
		Tag Tag `json:"tag"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	result.Tag.UserCount = 0
	tag = &result.Tag
	return
}

// List 查询所有Tag.
func List(clt *core.Client) (tags []Tag, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/get?access_token="

	var result struct {
		core.Error
		Tags []Tag `json:"tags"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	tags = result.Tags
	return
}

// Update 修改Tag名.
func Update(clt *core.Client, tagId int, name string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/update?access_token="

	var request struct {
		Tag struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tag"`
	}
	request.Tag.Id = tagId
	request.Tag.Name = name

	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// TagGet 根据TagId获取用户列表.
//  NOTE: 每次最多能获取 10000 个用户, 可以多次指定 nextOpenId 来获取以满足需求, 如果 nextOpenId == "" 则表示从头获取
func TagGet(clt *core.Client, tagId int, nextOpenId string) (rslt *GetResult, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token="

	var request = struct {
		Id     int    `json:"tagid"`
		OpenId string `json:"next_openid"`
	}{
		Id:     tagId,
		OpenId: nextOpenId,
	}
	var result struct {
		core.Error
		GetResult
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	rslt = &result.GetResult
	return
}

// Delete 删除Tag.
func Delete(clt *core.Client, tagId int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token="

	var request struct {
		Tag struct {
			Id int `json:"id"`
		} `json:"tag"`
	}
	request.Tag.Id = tagId

	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// BatchTag 批量打标签.
func BatchTag(clt *core.Client, openIdList []string, tagId int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token="

	if len(openIdList) <= 0 {
		return
	}

	var request = struct {
		OpenIdList []string `json:"openid_list,omitempty"`
		TagId      int      `json:"tagid"`
	}{
		OpenIdList: openIdList,
		TagId:      tagId,
	}
	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// BatchUntag 批量取消标签.
func BatchUntag(clt *core.Client, openIdList []string, tagId int) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token="

	if len(openIdList) <= 0 {
		return
	}

	var request = struct {
		OpenIdList []string `json:"openid_list,omitempty"`
		TagId      int      `json:"tagid"`
	}{
		OpenIdList: openIdList,
		TagId:      tagId,
	}
	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
