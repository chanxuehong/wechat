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
	MsgType      string `xml:"MsgType"      json:"MsgType"`      // 消息类型
}

// 包括了微信服务器推送到开发者 URL 的所有的消息类型
type Request struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	// fuck weixin, MsgId != MsgID
	MsgId int64 `xml:"MsgId,omitempty" json:"MsgId,omitempty"`
	MsgID int64 `xml:"MsgID,omitempty" json:"MsgID,omitempty"`

	Content      string  `xml:"Content,omitempty" json:"Content,omitempty"`
	MediaId      string  `xml:"MediaId,omitempty" json:"MediaId,omitempty"`
	PicURL       string  `xml:"PicUrl,omitempty" json:"PicUrl,omitempty"`
	Format       string  `xml:"Format,omitempty" json:"Format,omitempty"`
	Recognition  string  `xml:"Recognition,omitempty" json:"Recognition,omitempty"`
	ThumbMediaId string  `xml:"ThumbMediaId,omitempty" json:"ThumbMediaId,omitempty"`
	LocationX    float64 `xml:"Location_X,omitempty" json:"Location_X,omitempty"`
	LocationY    float64 `xml:"Location_Y,omitempty" json:"Location_Y,omitempty"`
	Scale        int     `xml:"Scale,omitempty" json:"Scale,omitempty"`
	Label        string  `xml:"Label,omitempty" json:"Label,omitempty"`
	Title        string  `xml:"Title,omitempty" json:"Title,omitempty"`
	Description  string  `xml:"Description,omitempty" json:"Description,omitempty"`
	URL          string  `xml:"Url,omitempty" json:"Url,omitempty"`

	Event       string  `xml:"Event,omitempty" json:"Event,omitempty"`
	EventKey    string  `xml:"EventKey,omitempty" json:"EventKey,omitempty"`
	Ticket      string  `xml:"Ticket,omitempty" json:"Ticket,omitempty"`
	Latitude    float64 `xml:"Latitude,omitempty" json:"Latitude,omitempty"`
	Longitude   float64 `xml:"Longitude,omitempty" json:"Longitude,omitempty"`
	Precision   float64 `xml:"Precision,omitempty" json:"Precision,omitempty"`
	Status      string  `xml:"Status,omitempty" json:"Status,omitempty"`
	TotalCount  int     `xml:"TotalCount,omitempty" json:"TotalCount,omitempty"`
	FilterCount int     `xml:"FilterCount,omitempty" json:"FilterCount,omitempty"`
	SentCount   int     `xml:"SentCount,omitempty" json:"SentCount,omitempty"`
	ErrorCount  int     `xml:"ErrorCount,omitempty" json:"ErrorCount,omitempty"`
	OrderId     string  `xml:"OrderId,omitempty" json:"OrderId,omitempty"`
	OrderStatus int     `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty"`
	ProductId   string  `xml:"ProductId,omitempty" json:"ProductId,omitempty"`
	SkuInfo     string  `xml:"SkuInfo,omitempty" json:"SkuInfo,omitempty"`
}

var zeroRequest Request

func (req *Request) Zero() *Request {
	*req = zeroRequest
	return req
}

// 文本消息
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406600283</CreateTime>
//      <MsgType><![CDATA[text]]></MsgType>
//      <Content><![CDATA[测试]]></Content>
//      <MsgId>6041302214229962560</MsgId>
//  </xml>
type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	Content string `xml:"Content" json:"Content"` // 文本消息内容
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
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406599389</CreateTime>
//      <MsgType><![CDATA[image]]></MsgType>
//      <PicUrl><![CDATA[http://mmbiz.qpic.cn/mmbiz/eUIjD3Aun1HwKC5Kn7KRAVnkibd8gYmI0ky6ywuWJjheZibWA6Zefj1tN5aJ1Shfv86yGxO0v8mF1VYmeUZdhJYw/0]]></PicUrl>
//      <MsgId>6041298374529199806</MsgId>
//      <MediaId><![CDATA[q4tXXIynvOOv15foN-8Z28VaAfy3zv33yy7jskNmax-8bxIf7XZB18pLKiKQ_U-G]]></MediaId>
//  </xml>
type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	PicURL  string `xml:"PicUrl"  json:"PicUrl"`  // 图片链接
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
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406599448</CreateTime>
//      <MsgType><![CDATA[voice]]></MsgType>
//      <MediaId><![CDATA[2suTDW9Bni2SkGQALX4y-8Jb1zWupzcUi69FKFzavnRkVP50aTb_y0uO5OgcEULP]]></MediaId>
//      <Format><![CDATA[amr]]></Format>
//      <MsgId>6041298627731652608</MsgId>
//      <Recognition><![CDATA[你好你好]]></Recognition>
//  </xml>
type Voice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 语音消息媒体id，可以调用多媒体文件下载接口拉取该媒体
	Format  string `xml:"Format"  json:"Format"`  // 语音格式，如amr，speex等

	// 语音识别结果，UTF8编码，
	// NOTE: 需要开通语音识别功能，否则该字段为空，即使开通了语音识别该字段还是有可能为空
	Recognition string `xml:"Recognition" json:"Recognition"`
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
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406600636</CreateTime>
//      <MsgType><![CDATA[video]]></MsgType>
//      <MediaId><![CDATA[c5r39YYfpHMJG5RLyr6TnGe7Iz0tymBpCvTNqe9pYEJcpzi15OXa2P_Qm5rBrcBJ]]></MediaId>
//      <ThumbMediaId><![CDATA[ZMmoLN8qzZd1u7vrfGUOolkdgyt8U3mIzY73XkYGgL-UrGyGpbRWjsm8J3TvWOEx]]></ThumbMediaId>
//      <MsgId>6041303730353418073</MsgId>
//  </xml>
type Video struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId        int64  `xml:"MsgId"        json:"MsgId"`        // 消息id, 64位整型
	MediaId      string `xml:"MediaId"      json:"MediaId"`      // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
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
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406600885</CreateTime>
//      <MsgType><![CDATA[location]]></MsgType>
//      <Location_X>23.444099</Location_X>
//      <Location_Y>113.632614</Location_Y>
//      <Scale>16</Scale>
//      <Label><![CDATA[测试位置]]></Label>
//      <MsgId>6041304799800274795</MsgId>
//  </xml>
type Location struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId     int64   `xml:"MsgId"      json:"MsgId"`      // 消息id, 64位整型
	LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
	LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
	Scale     int     `xml:"Scale"      json:"Scale"`      // 地图缩放大小
	Label     string  `xml:"Label"      json:"Label"`      // 地理位置信息
}

func (req *Request) Location() (location *Location) {
	location = &Location{
		CommonHead: req.CommonHead,
		MsgId:      req.MsgId,
		LocationX:  req.LocationX,
		LocationY:  req.LocationY,
		Scale:      req.Scale,
		Label:      req.Label,
	}
	return
}

// 链接消息
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406599749</CreateTime>
//      <MsgType><![CDATA[link]]></MsgType>
//      <Title><![CDATA[电线电缆基本知识！]]></Title>
//      <Description><![CDATA[http://mp.weixin.qq.com/s?__biz=MzA3NTAzMzIwMQ==&mid=201078176&idx=1&sn=fc46eb543aa2819c01ccabad546f44bf&scene=2#rd]]></Description>
//      <Url><![CDATA[http://mp.weixin.qq.com/s?__biz=MzA3NTAzMzIwMQ==&mid=201078176&idx=1&sn=fc46eb543aa2819c01ccabad546f44bf&scene=2&from=timeline&isappinstalled=0#rd]]></Url>
//      <MsgId>6041299920717426420</MsgId>
//  </xml>
type Link struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId       int64  `xml:"MsgId"       json:"MsgId"`       // 消息id, 64位整型
	Title       string `xml:"Title"       json:"Title"`       // 消息标题
	Description string `xml:"Description" json:"Description"` // 消息描述
	URL         string `xml:"Url"         json:"Url"`         // 消息链接
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

// 关注事件(普通关注)
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[ovx6euNq-hN2do74jeVSqZB82DiE]]></FromUserName>
//      <CreateTime>1406601711</CreateTime>
//      <MsgType><![CDATA[event]]></MsgType>
//      <Event><![CDATA[subscribe]]></Event>
//      <EventKey><![CDATA[]]></EventKey>
//  </xml>
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

// 取消关注
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406602046</CreateTime>
//      <MsgType><![CDATA[event]]></MsgType>
//      <Event><![CDATA[unsubscribe]]></Event>
//      <EventKey><![CDATA[]]></EventKey>
//  </xml>
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

// 用户未关注时，扫描带参数二维码进行关注后的事件推送
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406602445</CreateTime>
//      <MsgType><![CDATA[event]]></MsgType>
//      <Event><![CDATA[subscribe]]></Event>
//      <EventKey><![CDATA[qrscene_100000]]></EventKey>
//      <Ticket><![CDATA[gQEq8ToAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL1hrUHRHSC1sWUkwWmV0SlJKRzEzAAIEhAzXUwMEZAAAAA==]]></Ticket>
//  </xml>
type SubscribeByScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，subscribe
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket   string `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
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
		err = fmt.Errorf("EventKey 应该以 %s 为前缀, 但是现在是 %s", prefix, event.EventKey)
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
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406602958</CreateTime>
//      <MsgType><![CDATA[event]]></MsgType>
//      <Event><![CDATA[SCAN]]></Event>
//      <EventKey><![CDATA[100000]]></EventKey>
//      <Ticket><![CDATA[gQGT8DoAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL0ZrUDltS3psYUkwUllackpORzEzAAIEaw7XUwME6AMAAA==]]></Ticket>
//  </xml>
type ScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，SCAN
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket   string `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
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
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406603221</CreateTime>
//      <MsgType><![CDATA[event]]></MsgType>
//      <Event><![CDATA[LOCATION]]></Event>
//      <Latitude>23.446735</Latitude>
//      <Longitude>113.627739</Longitude>
//      <Precision>120.000000</Precision>
//  </xml>
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

// 点击菜单拉取消息时的事件推送
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406603449</CreateTime>
//      <MsgType><![CDATA[event]]></MsgType>
//      <Event><![CDATA[CLICK]]></Event>
//      <EventKey><![CDATA[key]]></EventKey>
//  </xml>
type MenuClickEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，CLICK
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，与自定义菜单接口中KEY值对应
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
//
//  <xml>
//      <ToUserName><![CDATA[gh_xxxxxxxxxxxx]]></ToUserName>
//      <FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
//      <CreateTime>1406603565</CreateTime>
//      <MsgType><![CDATA[event]]></MsgType>
//      <Event><![CDATA[VIEW]]></Event>
//      <EventKey><![CDATA[http://www.qq.com]]></EventKey>
//  </xml>
type MenuViewEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，VIEW
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，设置的跳转URL
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
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event string `xml:"Event" json:"Event"` // 事件信息，此处为 MASSSENDJOBFINISH

	MsgId int64 `xml:"MsgId" json:"MsgId"` // 群发的消息ID, 64位整型

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
	Status string `xml:"Status" json:"Status"`

	TotalCount int `xml:"TotalCount" json:"TotalCount"` // group_id 下粉丝数, 或者 openid_list 中的粉丝数

	// 过滤(过滤是指特定地区, 性别的过滤, 用户设置拒收的过滤; 用户接收已超4条的过滤）后,
	// 准备发送的粉丝数, 原则上, FilterCount = SentCount + ErrorCount
	FilterCount int `xml:"FilterCount" json:"FilterCount"`
	SentCount   int `xml:"SentCount"   json:"SentCount"`  // 发送成功的粉丝数
	ErrorCount  int `xml:"ErrorCount"  json:"ErrorCount"` // 发送失败的粉丝数
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

// 在模版消息发送任务完成后，微信服务器会将是否送达成功作为通知，发送到开发者中心中填写的服务器配置地址中。
type TemplateSendJobFinishEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event string `xml:"Event" json:"Event"` // 事件信息，此处为 TEMPLATESENDJOBFINISH

	MsgId int64 `xml:"MsgId" json:"MsgId"` // 模板消息ID

	// 送达成功时:                                     success
	// 送达由于用户拒收（用户设置拒绝接收公众号消息）而失败时:  failed:user block
	// 送达由于其他原因失败时:                            failed: system failed
	Status string `xml:"Status" json:"Status"`
}

func (req *Request) TemplateSendJobFinishEvent() (event *TemplateSendJobFinishEvent) {
	event = &TemplateSendJobFinishEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		MsgId:      req.MsgID, // NOTE
		Status:     req.Status,
	}
	return
}

// 微信小店, 订单付款通知
type MerchantOrderEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event       string `xml:"Event"       json:"Event"`       // 事件类型, merchant_order
	OrderId     string `xml:"OrderId"     json:"OrderId"`     // 订单 id
	OrderStatus int    `xml:"OrderStatus" json:"OrderStatus"` // 订单状态(2-待发货, 3-已发货, 5-已完成, 8-维权中)
	ProductId   string `xml:"ProductId"   json:"ProductId"`   // 商品 id
	SkuInfo     string `xml:"SkuInfo"     json:"SkuInfo"`     // sku 信息
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
