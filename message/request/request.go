// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

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
	XMLName struct{} `xml:"xml" json:"-"`
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

var _zero_request Request

// 因为 Request 结构体比较大, 每次都申请比较不划算, 并且这个结构体一般都是过渡,
// 不会常驻内存, 所以建议用对象池技术;
// 用对象池最好都要每次都 清零, 以防旧数据干扰.
func (req *Request) Zero() *Request {
	*req = _zero_request
	return req
}

// 文本消息
type Text struct {
	CommonHead

	MsgId   int64  `json:"MsgId"`   // 消息id, 64位整型
	Content string `json:"Content"` // 文本消息内容
}

func (req *Request) Text() (text *Text) {
	text = &Text{
		CommonHead: req.CommonHead,
		MsgId:      req.MsgId,
		Content:    req.Content,
	}
	return
}

// 图片消息
type Image struct {
	CommonHead

	MsgId   int64  `json:"MsgId"`   // 消息id, 64位整型
	MediaId string `json:"MediaId"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	PicURL  string `json:"PicUrl"`  // 图片链接
}

func (req *Request) Image() (image *Image) {
	image = &Image{
		CommonHead: req.CommonHead,
		MsgId:      req.MsgId,
		MediaId:    req.MediaId,
		PicURL:     req.PicURL,
	}
	return
}

// 语音消息
type Voice struct {
	CommonHead

	MsgId   int64  `json:"MsgId"`   // 消息id, 64位整型
	MediaId string `json:"MediaId"` // 语音消息媒体id，可以调用多媒体文件下载接口拉取该媒体
	Format  string `json:"Format"`  // 语音格式，如amr，speex等

	// 语音识别结果，UTF8编码，
	// NOTE: 需要开通语音识别功能，否则该字段为空，即使开通了语音识别该字段还是有可能为空
	Recognition string `json:"Recognition"`
}

func (req *Request) Voice() (voice *Voice) {
	voice = &Voice{
		CommonHead:  req.CommonHead,
		MsgId:       req.MsgId,
		MediaId:     req.MediaId,
		Format:      req.Format,
		Recognition: req.Recognition,
	}
	return
}

// 视频消息
type Video struct {
	CommonHead

	MsgId        int64  `json:"MsgId"`        // 消息id, 64位整型
	MediaId      string `json:"MediaId"`      // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaId string `json:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
}

func (req *Request) Video() (video *Video) {
	video = &Video{
		CommonHead:   req.CommonHead,
		MsgId:        req.MsgId,
		MediaId:      req.MediaId,
		ThumbMediaId: req.ThumbMediaId,
	}
	return
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

func (req *Request) Location() (location *Location) {
	location = &Location{
		CommonHead: req.CommonHead,
		MsgId:      req.MsgId,
		Location_X: req.Location_X,
		Location_Y: req.Location_Y,
		Scale:      req.Scale,
		Label:      req.Label,
	}
	return
}

// 链接消息
type Link struct {
	CommonHead

	MsgId       int64  `json:"MsgId"`       // 消息id, 64位整型
	Title       string `json:"Title"`       // 消息标题
	Description string `json:"Description"` // 消息描述
	URL         string `json:"Url"`         // 消息链接
}

func (req *Request) Link() (link *Link) {
	link = &Link{
		CommonHead:  req.CommonHead,
		MsgId:       req.MsgId,
		Title:       req.Title,
		Description: req.Description,
		URL:         req.URL,
	}
	return
}

// 关注事件
type SubscribeEvent struct {
	CommonHead

	Event string `json:"Event"` // 事件类型，subscribe(订阅)
}

func (req *Request) SubscribeEvent() (event *SubscribeEvent) {
	event = &SubscribeEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
	}
	return
}

// 取消关注
type UnsubscribeEvent struct {
	CommonHead

	Event string `json:"Event"` // 事件类型，unsubscribe(取消订阅)
}

func (req *Request) UnsubscribeEvent() (event *UnsubscribeEvent) {
	event = &UnsubscribeEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
	}
	return
}

// 用户未关注时，扫描带参数二维码进行关注后的事件推送
type SubscribeByScanEvent struct {
	CommonHead

	Event    string `json:"Event"`    // 事件类型，subscribe
	EventKey string `json:"EventKey"` // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket   string `json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

func (req *Request) SubscribeByScanEvent() (event *SubscribeByScanEvent) {
	event = &SubscribeByScanEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		EventKey:   req.EventKey,
		Ticket:     req.Ticket,
	}
	return
}

// 获取二维码参数
func (event *SubscribeByScanEvent) SceneId() (id uint32, err error) {
	const prefix = "qrscene_"

	if !strings.HasPrefix(event.EventKey, prefix) {
		err = fmt.Errorf("EventKey(%s) 应该以 %s 为前缀", event.EventKey, prefix)
		return
	}

	idUint64, err := strconv.ParseUint(event.EventKey[len(prefix):], 10, 32)
	if err != nil {
		return
	}
	id = uint32(idUint64)
	return
}

// 用户已关注时，扫描带参数二维码的事件推送
type ScanEvent struct {
	CommonHead

	Event    string `json:"Event"`    // 事件类型，SCAN
	EventKey string `json:"EventKey"` // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket   string `json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

func (req *Request) ScanEvent() (event *ScanEvent) {
	event = &ScanEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		EventKey:   req.EventKey,
		Ticket:     req.Ticket,
	}
	return
}

// 获取二维码参数
func (event *ScanEvent) SceneId() (id uint32, err error) {
	idUint64, err := strconv.ParseUint(event.EventKey, 10, 32)
	if err != nil {
		return
	}
	id = uint32(idUint64)
	return
}

// 上报地理位置事件
type LocationEvent struct {
	CommonHead

	Event     string  `json:"Event"`     // 事件类型，LOCATION
	Latitude  float64 `json:"Latitude"`  // 地理位置纬度
	Longitude float64 `json:"Longitude"` // 地理位置经度
	Precision float64 `json:"Precision"` // 地理位置精度
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

// 点击菜单拉取消息时的事件推送
type MenuClickEvent struct {
	CommonHead

	Event    string `json:"Event"`    // 事件类型，CLICK
	EventKey string `json:"EventKey"` // 事件KEY值，与自定义菜单接口中KEY值对应
}

func (req *Request) MenuClickEvent() (event *MenuClickEvent) {
	event = &MenuClickEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		EventKey:   req.EventKey,
	}
	return
}

// 点击菜单跳转链接时的事件推送
type MenuViewEvent struct {
	CommonHead

	Event    string `json:"Event"`    // 事件类型，VIEW
	EventKey string `json:"EventKey"` // 事件KEY值，设置的跳转URL
}

func (req *Request) MenuViewEvent() (event *MenuViewEvent) {
	event = &MenuViewEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		EventKey:   req.EventKey,
	}
	return
}

// 高级群发消息, 事件推送群发结果
type MassSendJobFinishEvent struct {
	CommonHead

	Event string `json:"Event"` // 事件信息，此处为MASSSENDJOBFINISH

	MsgId int64 `json:"MsgId"` // 群发的消息ID, 64位整型

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

func (req *Request) MassSendJobFinishEvent() (event *MassSendJobFinishEvent) {
	event = &MassSendJobFinishEvent{
		CommonHead:  req.CommonHead,
		Event:       req.Event,
		MsgId:       req.MsgID, // NOTE
		Status:      req.Status,
		TotalCount:  req.TotalCount,
		FilterCount: req.FilterCount,
		SentCount:   req.SentCount,
		ErrorCount:  req.ErrorCount,
	}
	return
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

func (req *Request) MerchantOrderEvent() (event *MerchantOrderEvent) {
	event = &MerchantOrderEvent{
		CommonHead:  req.CommonHead,
		Event:       req.Event,
		OrderId:     req.OrderId,
		OrderStatus: req.OrderStatus,
		ProductId:   req.ProductId,
		SkuInfo:     req.SkuInfo,
	}
	return
}
