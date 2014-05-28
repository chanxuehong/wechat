package user

// 分组
type Group struct {
	Id        int    `json:"id"`    // 分组id，由微信分配
	Name      string `json:"name"`  // 分组名字，UTF8编码
	UserCount int    `json:"count"` // 分组内用户数量
}

type UserInfo struct {
	OpenId   string `json:"openid"`   // 用户的标识，对当前公众号唯一
	Nickname string `json:"nickname"` // 用户的昵称
	Sex      int    `json:"sex"`      // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language string `json:"language"` // 用户的语言，简体中文为zh_CN
	City     string `json:"city"`     // 用户所在城市
	Province string `json:"province"` // 用户所在省份
	Country  string `json:"country"`  // 用户所在国家
	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，
	// 0代表640*640正方形头像），用户没有头像时该项为空
	HeadImgUrl string `json:"headimgurl"`
	// 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	SubscribeTime int64 `json:"subscribe_time"`
}

/*
关注者列表的遍历器.

	iter, err := Client.UserIterator("beginOpenId")
	if err != nil {
		...
	}

	for iter.HasNext() {
		openids, err := iter.Next()
		if err != nil {
			...
		}
		...
	}
*/
type UserIterator interface {
	HasNext() bool
	Next() ([]string, error)
	Total() int
}
