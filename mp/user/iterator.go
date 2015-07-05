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
//  iter, err := Client.UserIterator("BeginOpenId")
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
	clt Client // 关联的微信 Client

	lastUserListData  *UserListResult // 最近一次获取的用户数据
	nextPageHasCalled bool            // NextPage() 是否调用过
}

func (iter *UserIterator) TotalCount() int {
	return iter.lastUserListData.TotalCount
}

func (iter *UserIterator) HasNext() bool {
	if !iter.nextPageHasCalled { // 第一次调用需要特殊对待
		return iter.lastUserListData.GotCount > 0 || iter.lastUserListData.NextOpenId != ""
	}

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
	return iter.lastUserListData.NextOpenId != "" && iter.lastUserListData.GotCount == UserPageSizeLimit
}

func (iter *UserIterator) NextPage() (openidList []string, err error) {
	if !iter.nextPageHasCalled { // 第一次调用需要特殊对待
		iter.nextPageHasCalled = true

		openidList = iter.lastUserListData.Data.OpenIdList
		return
	}

	data, err := iter.clt.UserList(iter.lastUserListData.NextOpenId)
	if err != nil {
		return
	}

	iter.lastUserListData = data

	openidList = data.Data.OpenIdList
	return
}

// 获取用户遍历器, BeginOpenId 表示开始遍历用户, 如果 BeginOpenId == "" 则表示从头遍历.
func (clt Client) UserIterator(BeginOpenId string) (iter *UserIterator, err error) {
	// 逻辑上相当于第一次调用 UserIterator.NextPage, 因为第一次调用 UserIterator.HasNext 需要数据支撑, 所以提前获取了数据

	data, err := clt.UserList(BeginOpenId)
	if err != nil {
		return
	}

	iter = &UserIterator{
		clt:               clt,
		lastUserListData:  data,
		nextPageHasCalled: false,
	}
	return
}
