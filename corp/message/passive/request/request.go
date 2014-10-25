// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

type CommonHead struct {
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`   // 企业号CorpID
	FromUserName string `xml:"FromUserName" json:"FromUserName"` // 员工UserID
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`   // 消息创建时间（整型）, unixtime
	MsgType      string `xml:"MsgType"      json:"MsgType"`      // 消息类型
	AgentId      int64  `xml:"AgentID"      json:"AgentID"`      // 企业应用的id，可在应用的设置页面获取；如果id为0，则表示是整个企业号的关注/取消关注事件
}

// 微信服务器推送到开发者 URL 的所有已知的消息类型的组合体
type Request struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId int64 `xml:"MsgId" json:"MsgId"`

	Content      string  `xml:"Content"      json:"Content"`
	MediaId      string  `xml:"MediaId"      json:"MediaId"`
	PicURL       string  `xml:"PicUrl"       json:"PicUrl"`
	Format       string  `xml:"Format"       json:"Format"`
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"   json:"Location_X"`
	LocationY    float64 `xml:"Location_Y"   json:"Location_Y"`
	Scale        int     `xml:"Scale"        json:"Scale"`
	Label        string  `xml:"Label"        json:"Label"`

	Event        string `xml:"Event"       json:"Event"`
	EventKey     string `xml:"EventKey"    json:"EventKey"`
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
	Latitude  float64 `xml:"Latitude"    json:"Latitude"`
	Longitude float64 `xml:"Longitude"   json:"Longitude"`
	Precision float64 `xml:"Precision"   json:"Precision"`
}

var zeroRequest Request

func (req *Request) Zero() *Request {
	*req = zeroRequest
	return req
}
