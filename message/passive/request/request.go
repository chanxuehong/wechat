// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

type CommonHead struct {
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`   // 开发者微信号
	FromUserName string `xml:"FromUserName" json:"FromUserName"` // 发送方帐号(一个OpenID)
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`   // 消息创建时间(整型), unixtime
	MsgType      string `xml:"MsgType"      json:"MsgType"`      // 消息类型
}

// 微信服务器推送到开发者 URL 的所有已知的消息类型的组合体
type Request struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	// fuck weixin, MsgId != MsgID
	MsgId int64 `xml:"MsgId,omitempty" json:"MsgId,omitempty"`
	MsgID int64 `xml:"MsgID,omitempty" json:"MsgID,omitempty"`

	Content      string  `xml:"Content,omitempty"      json:"Content,omitempty"`
	MediaId      string  `xml:"MediaId,omitempty"      json:"MediaId,omitempty"`
	PicURL       string  `xml:"PicUrl,omitempty"       json:"PicUrl,omitempty"`
	Format       string  `xml:"Format,omitempty"       json:"Format,omitempty"`
	Recognition  string  `xml:"Recognition,omitempty"  json:"Recognition,omitempty"`
	ThumbMediaId string  `xml:"ThumbMediaId,omitempty" json:"ThumbMediaId,omitempty"`
	LocationX    float64 `xml:"Location_X,omitempty"   json:"Location_X,omitempty"`
	LocationY    float64 `xml:"Location_Y,omitempty"   json:"Location_Y,omitempty"`
	Scale        int     `xml:"Scale,omitempty"        json:"Scale,omitempty"`
	Label        string  `xml:"Label,omitempty"        json:"Label,omitempty"`
	Title        string  `xml:"Title,omitempty"        json:"Title,omitempty"`
	Description  string  `xml:"Description,omitempty"  json:"Description,omitempty"`
	URL          string  `xml:"Url,omitempty"          json:"Url,omitempty"`

	Event        string `xml:"Event,omitempty"       json:"Event,omitempty"`
	EventKey     string `xml:"EventKey,omitempty"    json:"EventKey,omitempty"`
	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`
		ScanResult string `xml:"ScanResult" json:"ScanResult"`
	} `xml:"ScanCodeInfo" json:"ScanCodeInfo"`
	SendPicsInfo struct {
		Count   int `xml:"Count"   json:"Count"`
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"`
	} `xml:"SendPicsInfo" json:"SendPicsInfo"`
	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X" json:"Location_X"`
		LocationY float64 `xml:"Location_Y" json:"Location_Y"`
		Scale     int     `xml:"Scale"      json:"Scale"`
		Label     string  `xml:"Label"      json:"Label"`
		Poiname   string  `xml:"Poiname"    json:"Poiname"`
	} `xml:"SendLocationInfo" json:"SendLocationInfo"`
	Ticket      string  `xml:"Ticket,omitempty"      json:"Ticket,omitempty"`
	Latitude    float64 `xml:"Latitude,omitempty"    json:"Latitude,omitempty"`
	Longitude   float64 `xml:"Longitude,omitempty"   json:"Longitude,omitempty"`
	Precision   float64 `xml:"Precision,omitempty"   json:"Precision,omitempty"`
	Status      string  `xml:"Status,omitempty"      json:"Status,omitempty"`
	TotalCount  int     `xml:"TotalCount,omitempty"  json:"TotalCount,omitempty"`
	FilterCount int     `xml:"FilterCount,omitempty" json:"FilterCount,omitempty"`
	SentCount   int     `xml:"SentCount,omitempty"   json:"SentCount,omitempty"`
	ErrorCount  int     `xml:"ErrorCount,omitempty"  json:"ErrorCount,omitempty"`
	OrderId     string  `xml:"OrderId,omitempty"     json:"OrderId,omitempty"`
	OrderStatus int     `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty"`
	ProductId   string  `xml:"ProductId,omitempty"   json:"ProductId,omitempty"`
	SkuInfo     string  `xml:"SkuInfo,omitempty"     json:"SkuInfo,omitempty"`
}

var zeroRequest Request

func (req *Request) Zero() *Request {
	*req = zeroRequest
	return req
}
