// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

// 关注事件
//  如果id为0，则表示是整个企业号的关注/取消关注事件
type SubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event string `xml:"Event" json:"Event"` // 事件类型，subscribe(订阅)
}

func (req *Request) SubscribeEvent() (event *SubscribeEvent) {
	event = &SubscribeEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
	}
	return
}

// 取消关注事件
//  如果id为0，则表示是整个企业号的关注/取消关注事件
type UnsubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event string `xml:"Event" json:"Event"` // 事件类型，unsubscribe(取消订阅)
}

func (req *Request) UnsubscribeEvent() (event *UnsubscribeEvent) {
	event = &UnsubscribeEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
	}
	return
}

// 上报地理位置事件
type LocationEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event     string  `xml:"Event"     json:"Event"`     // 事件类型，LOCATION
	Latitude  float64 `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64 `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64 `xml:"Precision" json:"Precision"` // 地理位置精度
}

func (req *Request) LocationEvent() (event *LocationEvent) {
	event = &LocationEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		Precision:  req.Precision,
	}
	return
}
