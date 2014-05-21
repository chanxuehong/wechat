// 请求消息

package wechat

import (
	"encoding/xml"
)

// 包括了所有的请求消息类型
type RequestMsg struct {
	XMLName xml.Name `xml:"xml" json:"-"`

	// head
	ToUserName   string `xml:"ToUserName"   json:"ToUserName,omitempty"`   // 开发者微信号
	FromUserName string `xml:"FromUserName" json:"FromUserName,omitempty"` // 发送方帐号(一个OpenID)
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime,omitempty"`   // 消息创建时间(整型), unixtime
	MsgType      string `xml:"MsgType"      json:"MsgType,omitempty"`      // text, image, voice, video, location, link, event

	// body

	// fuck weixin, MsgId != MsgID
	MsgId int64 `xml:"MsgId" json:"MsgId,omitempty"` // 消息id, 64位整型
	MsgID int64 `xml:"MsgID" json:"MsgID,omitempty"` // 消息id, 64位整型

	// common message
	Content      string  `xml:"Content"      json:"Content,omitempty"`      // text, 文本消息内容
	MediaId      string  `xml:"MediaId"      json:"MediaId,omitempty"`      // image, voice, video 消息媒体id, 可以调用多媒体文件下载接口拉取数据
	PicUrl       string  `xml:"PicUrl"       json:"PicUrl,omitempty"`       // image, 图片链接
	Format       string  `xml:"Format"       json:"Format,omitempty"`       // voice, 语音格式, 如amr, speex等
	Recognition  string  `xml:"Recognition"  json:"Recognition,omitempty"`  // voice, 语音识别结果, UTF8编码
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId,omitempty"` // video, 视频消息缩略图的媒体id, 可以调用多媒体文件下载接口拉取数据
	Location_X   float64 `xml:"Location_X"   json:"Location_X,omitempty"`   // location, 地理位置纬度
	Location_Y   float64 `xml:"Location_Y"   json:"Location_Y,omitempty"`   // location, 地理位置经度
	Scale        int     `xml:"Scale"        json:"Scale,omitempty"`        // location, 地图缩放大小
	Label        string  `xml:"Label"        json:"Label,omitempty"`        // location, 地理位置信息
	Title        string  `xml:"Title"        json:"Title,omitempty"`        // link, 消息标题
	Description  string  `xml:"Description"  json:"Description,omitempty"`  // link, 消息描述
	Url          string  `xml:"Url"          json:"Url,omitempty"`          // link, 消息链接

	// event message
	Event     string  `xml:"Event"     json:"Event,omitempty"`     // subscribe, unsubscribe, SCAN, LOCATION, CLICK, VIEW, MASSSENDJOBFINISH
	EventKey  string  `xml:"EventKey"  json:"EventKey,omitempty"`  // 不同的 Event 不同的功能
	Ticket    string  `xml:"Ticket"    json:"Ticket,omitempty"`    // 二维码的ticket, 可用来换取二维码图片
	Latitude  float64 `xml:"Latitude"  json:"Latitude,omitempty"`  // 地理位置纬度
	Longitude float64 `xml:"Longitude" json:"Longitude,omitempty"` // 地理位置经度
	Precision float64 `xml:"Precision" json:"Precision,omitempty"` // 地理位置精度

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
	Status     string `xml:"Status"     json:"Status,omitempty"`
	TotalCount int    `xml:"TotalCount" json:"TotalCount,omitempty"` // group_id 下粉丝数, 或者 openid_list 中的粉丝数
	// 过滤(过滤是指特定地区, 性别的过滤, 用户设置拒收的过滤; 用户接收已超4条的过滤）后,
	// 准备发送的粉丝数, 原则上, FilterCount = SentCount + ErrorCount
	FilterCount int `xml:"FilterCount" json:"FilterCount,omitempty"`
	SentCount   int `xml:"SentCount"   json:"SentCount,omitempty"`  // 发送成功的粉丝数
	ErrorCount  int `xml:"ErrorCount"  json:"ErrorCount,omitempty"` // 发送失败的粉丝数
}

// 因为 RequestMsg 结构体比较大, 每次都申请比较不划算, 并且这个结构体一般都是过度,
// 不会常驻内存, 所以建议用对象池技术; 用对象池最好都要每次都 清零, 以防旧数据干扰.
func (msg *RequestMsg) Zero() {
	*msg = _zeroRequestMsg
}
