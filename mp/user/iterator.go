// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package user

const (
	UserPageSizeLimit = 10000 // 每次拉取的 OPENID 个数最大值为 10000
)

// 用户遍历器
//
//  iter, err := Client.UserIterator("beginOpenId")
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
	lastUserListData *UserListResult // 最近一次获取的用户数据

	wechatClient   *Client // 关联的微信 Client
	nextPageCalled bool    // NextPage() 是否调用过
}

func (iter *UserIterator) Total() int {
	return iter.lastUserListData.TotalCount
}

func (iter *UserIterator) HasNext() bool {
	if !iter.nextPageCalled { // 还没有调用 NextPage(), 从创建的时候获取的数据来判断
		return iter.lastUserListData.GotCount > 0
	}

	// 已经调用过 NextPage(), 根据上一次 next_openid 字段是否为空来判断
	//
	// 跟文档的描述貌似有点不一样, 即使后续没有用户, 貌似 next_openid 还是不为空!
	// 所以增加了一个判断 iter.userGetResponse.GetCount == UserPageCountLimit
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
	return iter.lastUserListData.NextOpenId != "" &&
		iter.lastUserListData.GotCount == UserPageSizeLimit
}

func (iter *UserIterator) NextPage() (openids []string, err error) {
	if !iter.nextPageCalled { // 还没有调用 NextPage(), 从创建的时候获取的数据中获取
		openids = iter.lastUserListData.Data.OpenId
		iter.nextPageCalled = true
		return
	}

	// 不是第一次调用的都要从服务器拉取数据
	data, err := iter.wechatClient.UserList(iter.lastUserListData.NextOpenId)
	if err != nil {
		return
	}

	openids = data.Data.OpenId
	iter.lastUserListData = data //
	return
}

// 获取用户遍历器, beginOpenId 表示开始遍历用户, 如果 beginOpenId == "" 则表示从头遍历.
func (clt *Client) UserIterator(beginOpenId string) (iter *UserIterator, err error) {
	data, err := clt.UserList(beginOpenId)
	if err != nil {
		return
	}

	iter = &UserIterator{
		lastUserListData: data,
		wechatClient:     clt,
		nextPageCalled:   false,
	}
	return
}
