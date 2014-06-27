package request

import (
	"fmt"
	"strconv"
	"strings"
)

type CommonHead struct {
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`   // 开发者微信号
	FromUserName string `xml:"FromUserName" json:"FromUserName"` // 发送方帐号(一个OpenID)
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`   // 消息创建时间(整型), unixtime
	MsgType      string `xml:"MsgType"      json:"MsgType"`      // text, image, voice, video, location, link, event
}

// 包括了所有从微信服务器推送过来的消息类型
type Request struct {
	XMLName struct{} `xml:"xml"`
	CommonHead

	// fuck weixin, MsgId != MsgID
	MsgId int64 `xml:"MsgId"` // 消息id, 64位整型
	MsgID int64 `xml:"MsgID"` // 消息id, 64位整型; 高级群发接口的 事件推送群发结果 貌似用的是这个!

	// common message
	Content      string  `xml:"Content"`      // text, 文本消息内容
	MediaId      string  `xml:"MediaId"`      // image, voice, video 消息媒体id, 可以调用多媒体文件下载接口拉取数据
	PicURL       string  `xml:"PicUrl"`       // image, 图片链接
	Format       string  `xml:"Format"`       // voice, 语音格式, 如amr, speex等
	Recognition  string  `xml:"Recognition"`  // voice, 语音识别结果, UTF8编码
	ThumbMediaId string  `xml:"ThumbMediaId"` // video, 视频消息缩略图的媒体id, 可以调用多媒体文件下载接口拉取数据
	Location_X   float64 `xml:"Location_X"`   // location, 地理位置纬度
	Location_Y   float64 `xml:"Location_Y"`   // location, 地理位置经度
	Scale        int     `xml:"Scale"`        // location, 地图缩放大小
	Label        string  `xml:"Label"`        // location, 地理位置信息
	Title        string  `xml:"Title"`        // link, 消息标题
	Description  string  `xml:"Description"`  // link, 消息描述
	URL          string  `xml:"Url"`          // link, 消息链接

	// event message
	Event     string  `xml:"Event"`     // subscribe, unsubscribe, SCAN, LOCATION, CLICK, VIEW, MASSSENDJOBFINISH
	EventKey  string  `xml:"EventKey"`  // 不同的 Event 不同的功能
	Ticket    string  `xml:"Ticket"`    // 二维码的ticket, 可用来换取二维码图片
	Latitude  float64 `xml:"Latitude"`  // 地理位置纬度
	Longitude float64 `xml:"Longitude"` // 地理位置经度
	Precision float64 `xml:"Precision"` // 地理位置精度

	// 群发的结构, 为 "send success" 或 "send fail" 或 "err(num)".
	// 但 send success 时, 也有可能因用户拒收公众号的消息, 系统错误等原因造成少量用户接收失败.
	// err(num) 是审核失败的具体原因, 可能的情况如下:
	// err(10001), //涉嫌广告
	// err(20001), //涉嫌政治
	// err(20004), //涉嫌社会
	// err(20002), //涉嫌色情
	// err(20006), //涉嫌违法犯罪
	// err(20008), //涉嫌欺诈
	// err(20013), //涉嫌版权
	// err(22000), //涉嫌互推(互相宣传)
	// err(21000), //涉嫌其他
	Status string `xml:"Status"`

	TotalCount int `xml:"TotalCount"` // group_id 下粉丝数, 或者 openid_list 中的粉丝数

	// 过滤(过滤是指特定地区, 性别的过滤, 用户设置拒收的过滤; 用户接收已超4条的过滤）后,
	// 准备发送的粉丝数, 原则上, FilterCount = SentCount + ErrorCount
	FilterCount int `xml:"FilterCount"`
	SentCount   int `xml:"SentCount"`  // 发送成功的粉丝数
	ErrorCount  int `xml:"ErrorCount"` // 发送失败的粉丝数

	// 微信小店, 订单付款通知
	OrderId     string `xml:"OrderId"`
	OrderStatus int    `xml:"OrderStatus"`
	ProductId   string `xml:"ProductId"`
	SkuInfo     string `xml:"SkuInfo"`
}

var _zeroRequest Request

// 因为 Request 结构体比较大, 每次都申请比较不划算, 并且这个结构体一般都是过度,
// 不会常驻内存, 所以建议用对象池技术; 用对象池最好都要每次都 清零, 以防旧数据干扰.
func (msg *Request) Zero() *Request {
	*msg = _zeroRequest
	return msg
}

// 文本消息
type Text struct {
	CommonHead

	MsgId   int64  `json:"MsgId"`   // 消息id, 64位整型
	Content string `json:"Content"` // 文本消息内容
}

func (msg *Request) Text() *Text {
	var r Text
	r.CommonHead = msg.CommonHead
	r.MsgId = msg.MsgId
	r.Content = msg.Content

	return &r
}

// 图片消息
type Image struct {
	CommonHead

	MsgId   int64  `json:"MsgId"`   // 消息id, 64位整型
	MediaId string `json:"MediaId"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	PicURL  string `json:"PicUrl"`  // 图片链接
}

func (msg *Request) Image() *Image {
	var r Image
	r.CommonHead = msg.CommonHead
	r.MsgId = msg.MsgId
	r.MediaId = msg.MediaId
	r.PicURL = msg.PicURL

	return &r
}

// 语音消息
type Voice struct {
	CommonHead

	MsgId   int64  `json:"MsgId"`   // 消息id, 64位整型
	MediaId string `json:"MediaId"` // 语音消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Format  string `json:"Format"`  // 语音格式，如amr，speex等
}

func (msg *Request) Voice() *Voice {
	var r Voice
	r.CommonHead = msg.CommonHead
	r.MsgId = msg.MsgId
	r.MediaId = msg.MediaId
	r.Format = msg.Format

	return &r
}

// 接收语音识别结果
type VoiceRecognition struct {
	CommonHead

	MsgId       int64  `json:"MsgId"`       // 消息id, 64位整型
	MediaId     string `json:"MediaId"`     // 语音消息媒体id，可以调用多媒体文件下载接口拉取该媒体
	Format      string `json:"Format"`      // 语音格式：amr
	Recognition string `json:"Recognition"` // 语音识别结果，UTF8编码
}

func (msg *Request) VoiceRecognition() *VoiceRecognition {
	var r VoiceRecognition
	r.CommonHead = msg.CommonHead
	r.MsgId = msg.MsgId
	r.MediaId = msg.MediaId
	r.Format = msg.Format
	r.Recognition = msg.Recognition

	return &r
}

// 视频消息
type Video struct {
	CommonHead

	MsgId        int64  `json:"MsgId"`        // 消息id, 64位整型
	MediaId      string `json:"MediaId"`      // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaId string `json:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
}

func (msg *Request) Video() *Video {
	var r Video
	r.CommonHead = msg.CommonHead
	r.MsgId = msg.MsgId
	r.MediaId = msg.MediaId
	r.ThumbMediaId = msg.ThumbMediaId

	return &r
}

// 地理位置消息
type Location struct {
	CommonHead

	MsgId      int64   `json:"MsgId"`      // 消息id, 64位整型
	Location_X float64 `json:"Location_X"` // 地理位置纬度
	Location_Y float64 `json:"Location_Y"` // 地理位置经度
	Scale      int     `json:"Scale"`      // 地图缩放大小
	Label      string  `json:"Label"`      // 地理位置信息
}

func (msg *Request) Location() *Location {
	var r Location
	r.CommonHead = msg.CommonHead
	r.MsgId = msg.MsgId
	r.Location_X = msg.Location_X
	r.Location_Y = msg.Location_Y
	r.Scale = msg.Scale
	r.Label = msg.Label

	return &r
}

// 链接消息
type Link struct {
	CommonHead

	MsgId       int64  `json:"MsgId"`       // 消息id, 64位整型
	Title       string `json:"Title"`       // 消息标题
	Description string `json:"Description"` // 消息描述
	URL         string `json:"Url"`         // 消息链接
}

func (msg *Request) Link() *Link {
	var r Link
	r.CommonHead = msg.CommonHead
	r.MsgId = msg.MsgId
	r.Title = msg.Title
	r.Description = msg.Description
	r.URL = msg.URL

	return &r
}

// 关注事件
type SubscribeEvent struct {
	CommonHead

	Event string `json:"Event"` // 事件类型，subscribe(订阅)
}

func (msg *Request) SubscribeEvent() *SubscribeEvent {
	var r SubscribeEvent
	r.CommonHead = msg.CommonHead
	r.Event = msg.Event

	return &r
}

// 取消关注
type UnsubscribeEvent struct {
	CommonHead

	Event string `json:"Event"` // 事件类型，unsubscribe(取消订阅)
}

func (msg *Request) UnsubscribeEvent() *UnsubscribeEvent {
	var r UnsubscribeEvent
	r.CommonHead = msg.CommonHead
	r.Event = msg.Event

	return &r
}

// 用户未关注时，扫描带参数二维码进行关注后的事件推送
type SubscribeByScanEvent struct {
	CommonHead

	Event    string `json:"Event"`    // 事件类型，subscribe
	EventKey string `json:"EventKey"` // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket   string `json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

// 获取二维码参数
func (event *SubscribeByScanEvent) SceneId() (id uint32, err error) {
	const prefix = "qrscene_"

	if !strings.HasPrefix(event.EventKey, prefix) {
		err = fmt.Errorf("EventKey(%s) 应该以 %s 为前缀", event.EventKey, prefix)
		return
	}

	id64, err := strconv.ParseUint(event.EventKey[len(prefix):], 10, 32)
	if err != nil {
		return
	}
	id = uint32(id64)
	return
}

func (msg *Request) SubscribeByScanEvent() *SubscribeByScanEvent {
	var r SubscribeByScanEvent
	r.CommonHead = msg.CommonHead
	r.Event = msg.Event
	r.EventKey = msg.EventKey
	r.Ticket = msg.Ticket

	return &r
}

// 用户已关注时，扫描带参数二维码的事件推送
type ScanEvent struct {
	CommonHead

	Event    string `json:"Event"`    // 事件类型，SCAN
	EventKey string `json:"EventKey"` // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket   string `json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

// 获取二维码参数
func (event *ScanEvent) SceneId() (id uint32, err error) {
	id64, err := strconv.ParseUint(event.EventKey, 10, 32)
	if err != nil {
		return
	}
	id = uint32(id64)
	return
}

func (msg *Request) ScanEvent() *ScanEvent {
	var r ScanEvent
	r.CommonHead = msg.CommonHead
	r.Event = msg.Event
	r.EventKey = msg.EventKey
	r.Ticket = msg.Ticket

	return &r
}

// 上报地理位置事件
type LocationEvent struct {
	CommonHead

	Event     string  `json:"Event"`     // 事件类型，LOCATION
	Latitude  float64 `json:"Latitude"`  // 地理位置纬度
	Longitude float64 `json:"Longitude"` // 地理位置经度
	Precision float64 `json:"Precision"` // 地理位置精度
}

func (msg *Request) LocationEvent() *LocationEvent {
	var r LocationEvent
	r.CommonHead = msg.CommonHead
	r.Event = msg.Event
	r.Latitude = msg.Latitude
	r.Longitude = msg.Longitude
	r.Precision = msg.Precision

	return &r
}

// 点击菜单拉取消息时的事件推送
type MenuClickEvent struct {
	CommonHead

	Event    string `json:"Event"`    // 事件类型，CLICK
	EventKey string `json:"EventKey"` // 事件KEY值，与自定义菜单接口中KEY值对应
}

func (msg *Request) MenuClickEvent() *MenuClickEvent {
	var r MenuClickEvent
	r.CommonHead = msg.CommonHead
	r.Event = msg.Event
	r.EventKey = msg.EventKey

	return &r
}

// 点击菜单跳转链接时的事件推送
type MenuViewEvent struct {
	CommonHead

	Event    string `json:"Event"`    // 事件类型，VIEW
	EventKey string `json:"EventKey"` // 事件KEY值，设置的跳转URL
}

func (msg *Request) MenuViewEvent() *MenuViewEvent {
	var r MenuViewEvent
	r.CommonHead = msg.CommonHead
	r.Event = msg.Event
	r.EventKey = msg.EventKey

	return &r
}

// 高级群发消息, 事件推送群发结果
type MassSendJobFinishEvent struct {
	CommonHead

	MsgId int64 `json:"MsgId"` // 群发的消息ID, 64位整型

	Event string `json:"Event"` // 事件信息，此处为MASSSENDJOBFINISH

	// 群发的结构, 为 "send success" 或 "send fail" 或 "err(num)".
	// 但 send success 时, 也有可能因用户拒收公众号的消息, 系统错误等原因造成少量用户接收失败.
	// err(num) 是审核失败的具体原因, 可能的情况如下:
	// err(10001), //涉嫌广告
	// err(20001), //涉嫌政治
	// err(20004), //涉嫌社会
	// err(20002), //涉嫌色情
	// err(20006), //涉嫌违法犯罪
	// err(20008), //涉嫌欺诈
	// err(20013), //涉嫌版权
	// err(22000), //涉嫌互推(互相宣传)
	// err(21000), //涉嫌其他
	Status string `json:"Status"`

	TotalCount int `json:"TotalCount"` // group_id 下粉丝数, 或者 openid_list 中的粉丝数

	// 过滤(过滤是指特定地区, 性别的过滤, 用户设置拒收的过滤; 用户接收已超4条的过滤）后,
	// 准备发送的粉丝数, 原则上, FilterCount = SentCount + ErrorCount
	FilterCount int `json:"FilterCount"`
	SentCount   int `json:"SentCount"`  // 发送成功的粉丝数
	ErrorCount  int `json:"ErrorCount"` // 发送失败的粉丝数
}

func (msg *Request) MassSendJobFinishEvent() *MassSendJobFinishEvent {
	var r MassSendJobFinishEvent
	r.CommonHead = msg.CommonHead
	r.MsgId = msg.MsgID // NOTE
	r.Event = msg.Event
	r.Status = msg.Status
	r.TotalCount = msg.TotalCount
	r.FilterCount = msg.FilterCount
	r.SentCount = msg.SentCount
	r.ErrorCount = msg.ErrorCount

	return &r
}

// 微信小店, 订单付款通知
type MerchantOrderEvent struct {
	CommonHead

	Event       string `json:"Event"` // 事件类型，merchant_order
	OrderId     string `json:"OrderId"`
	OrderStatus int    `json:"OrderStatus"` // 订单状态(2-待发货, 3-已发货, 5-已完成, 8-维权中)
	ProductId   string `json:"ProductId"`
	SkuInfo     string `json:"SkuInfo"`
}

func (msg *Request) MerchantOrderEvent() *MerchantOrderEvent {
	var r MerchantOrderEvent
	r.CommonHead = msg.CommonHead
	r.Event = msg.Event
	r.OrderId = msg.OrderId
	r.OrderStatus = msg.OrderStatus
	r.ProductId = msg.ProductId
	r.SkuInfo = msg.SkuInfo

	return &r
}
