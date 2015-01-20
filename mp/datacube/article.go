// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)

package datacube

type ArticleBaseData struct {
	IntPageReadUser  int `json:"int_page_read_user"`  // 图文页（点击群发图文卡片进入的页面）的阅读人数
	IntPageReadCount int `json:"int_page_read_count"` // 图文页的阅读次数
	OriPageReadUser  int `json:"ori_page_read_user"`  // 原文页（点击图文页“阅读原文”进入的页面）的阅读人数，无原文页时此处数据为0
	OriPageReadCount int `json:"ori_page_read_count"` // 原文页的阅读次数
	ShareUser        int `json:"share_user"`          // 分享的人数
	ShareCount       int `json:"share_count"`         // 分享的次数
	AddToFavUser     int `json:"add_to_fav_user"`     // 收藏的人数
	AddToFavCount    int `json:"add_to_fav_count"`    // 收藏的次数
}

type ArticleSummaryData struct {
	RefDate    string `json:"ref_date"`    // 数据的日期，需在begin_date和end_date之间
	UserSource int    `json:"user_source"` //
	Msgid      string `json:"msgid"`       // 这里的msgid实际上是由msgid（图文消息id）和index（消息次序索引）组成， 例如12003_3， 其中12003是msgid，即一次群发的id消息的； 3为index，假设该次群发的图文消息共5个文章（因为可能为多图文）， 3表示5个中的第3个
	Title      string `json:"title"`       // 图文消息的标题
	ArticleBaseData
}

type ArticleTotalData struct {
	RefDate    string `json:"ref_date"`
	Msgid      string `json:"msgid"`
	Title      string `json:"title"`
	UserSource int    `json:"user_source"`
	Details    []struct {
		StatDate   string `json:"stat_date"`
		TargetUser int    `json:"target_user"` // 送达人数，一般约等于总粉丝数（需排除黑名单或其他异常情况下无法收到消息的粉丝）
		ArticleBaseData
	} `json:"details"`
}

type UserReadData struct {
	RefDate    string `json:"ref_date"`
	UserSource int    `json:"user_source"`
	ArticleBaseData
}

type UserReadHourData struct {
	RefHour int `json:"ref_hour"` // 数据的小时，包括从000到2300，分别代表的是[000,100)到[2300,2400)，即每日的第1小时和最后1小时
	UserReadData
	TotalOnlineTime int64 `json:"total_online_time"`
}

type UserShareData struct {
	RefDate    string `json:"ref_date"`
	UserSource int    `json:"user_source"`
	ShareScene int    `json:"share_scene"` // 分享的场景：1代表好友转发 2代表朋友圈 3代表腾讯微博 255代表其他
	ShareCount int    `json:"shareCount"`
	ShareUser  int    `json:"share_user"`
}

type UserShareHourData struct {
	UserShareData
	RefDate string `json:"ref_date"`
}
