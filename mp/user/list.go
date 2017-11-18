package user

import (
	"net/url"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

// 获取用户列表返回的数据结构
type ListResult struct {
	TotalCount int `json:"total"` // 关注该公众账号的总用户数
	ItemCount  int `json:"count"` // 拉取的OPENID个数, 最大值为10000

	Data struct {
		OpenIdList []string `json:"openid,omitempty"`
	} `json:"data"` // 列表数据, OPENID的列表

	// 拉取列表的最后一个用户的OPENID, 如果 next_openid == "" 则表示没有了用户数据
	NextOpenId string `json:"next_openid"`
}

// List 获取用户列表.
//  NOTE: 每次最多能获取 10000 个用户, 可以多次指定 nextOpenId 来获取以满足需求, 如果 nextOpenId == "" 则表示从头获取
func List(clt *core.Client, nextOpenId string) (rslt *ListResult, err error) {
	var incompleteURL string
	if nextOpenId == "" {
		incompleteURL = "https://api.weixin.qq.com/cgi-bin/user/get?access_token="
	} else {
		incompleteURL = "https://api.weixin.qq.com/cgi-bin/user/get?next_openid=" + url.QueryEscape(nextOpenId) + "&access_token="
	}

	var result struct {
		core.Error
		ListResult
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	rslt = &result.ListResult
	return
}

// =====================================================================================================================

// UserIterator
//
//  iter, err := NewUserIterator(clt, "NextOpenId")
//  if err != nil {
//      // TODO: 增加你的代码
//  }
//
//  for iter.HasNext() {
//      openids, err := iter.NextPage()
//      if err != nil {
//          // TODO: 增加你的代码
//      }
//      // TODO: 增加你的代码
//  }
type UserIterator struct {
	clt *core.Client

	lastListResult *ListResult
	nextPageCalled bool
}

func (iter *UserIterator) TotalCount() int {
	return iter.lastListResult.TotalCount
}

func (iter *UserIterator) HasNext() bool {
	if !iter.nextPageCalled {
		return iter.lastListResult.ItemCount > 0 || iter.lastListResult.NextOpenId != ""
	}
	return iter.lastListResult.NextOpenId != ""
}

func (iter *UserIterator) NextPage() (openIdList []string, err error) {
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		openIdList = iter.lastListResult.Data.OpenIdList
		return
	}

	rslt, err := List(iter.clt, iter.lastListResult.NextOpenId)
	if err != nil {
		return
	}

	iter.lastListResult = rslt

	openIdList = rslt.Data.OpenIdList
	return
}

// NewUserIterator 获取用户遍历器, 从 nextOpenId 开始遍历, 如果 nextOpenId == "" 则表示从头遍历.
func NewUserIterator(clt *core.Client, nextOpenId string) (iter *UserIterator, err error) {
	// 逻辑上相当于第一次调用 UserIterator.NextPage,
	// 因为第一次调用 UserIterator.HasNext 需要数据支撑, 所以提前获取了数据
	rslt, err := List(clt, nextOpenId)
	if err != nil {
		return
	}

	iter = &UserIterator{
		clt:            clt,
		lastListResult: rslt,
		nextPageCalled: false,
	}
	return
}
