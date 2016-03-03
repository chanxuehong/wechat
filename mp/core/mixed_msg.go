package core

type (
	MsgType   string
	EventType string
)

// 微信服务器推送过来的消息(事件)的通用消息头.
type MsgHeader struct {
	ToUserName   string  `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string  `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime"   json:"CreateTime"`
	MsgType      MsgType `xml:"MsgType"      json:"MsgType"`
}

// 微信服务器推送过来的消息(事件)的合集.
type MixedMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MsgHeader

	MsgId int64 `xml:"MsgId" json:"MsgId"`
	MsgID int64 `xml:"MsgID" json:"MsgID"` // 群发消息和模板消息的消息ID, 而不是此事件消息的ID

	Content      string  `xml:"Content"      json:"Content"`
	MediaId      string  `xml:"MediaId"      json:"MediaId"`
	PicURL       string  `xml:"PicUrl"       json:"PicUrl"`
	Format       string  `xml:"Format"       json:"Format"`
	Recognition  string  `xml:"Recognition"  json:"Recognition"`
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"   json:"Location_X"`
	LocationY    float64 `xml:"Location_Y"   json:"Location_Y"`
	Scale        int     `xml:"Scale"        json:"Scale"`
	Label        string  `xml:"Label"        json:"Label"`
	Title        string  `xml:"Title"        json:"Title"`
	Description  string  `xml:"Description"  json:"Description"`
	URL          string  `xml:"Url"          json:"Url"`

	Event    EventType `xml:"Event"    json:"Event"`
	EventKey string    `xml:"EventKey" json:"EventKey"`

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`
		ScanResult string `xml:"ScanResult" json:"ScanResult"`
	} `xml:"ScanCodeInfo" json:"ScanCodeInfo"`

	SendPicsInfo struct {
		Count   int `xml:"Count" json:"Count"`
		PicList []struct {
			PicMD5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"`
	} `xml:"SendPicsInfo" json:"SendPicsInfo"`

	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X" json:"Location_X"`
		LocationY float64 `xml:"Location_Y" json:"Location_Y"`
		Scale     int     `xml:"Scale"      json:"Scale"`
		Label     string  `xml:"Label"      json:"Label"`
		PoiName   string  `xml:"Poiname"    json:"Poiname"`
	} `xml:"SendLocationInfo" json:"SendLocationInfo"`

	Ticket      string  `xml:"Ticket"      json:"Ticket"`
	Latitude    float64 `xml:"Latitude"    json:"Latitude"`
	Longitude   float64 `xml:"Longitude"   json:"Longitude"`
	Precision   float64 `xml:"Precision"   json:"Precision"`
	Status      string  `xml:"Status"      json:"Status"`
	TotalCount  int     `xml:"TotalCount"  json:"TotalCount"`
	FilterCount int     `xml:"FilterCount" json:"FilterCount"`
	SentCount   int     `xml:"SentCount"   json:"SentCount"`
	ErrorCount  int     `xml:"ErrorCount"  json:"ErrorCount"`

	// merchant
	OrderId     string `xml:"OrderId"     json:"OrderId"`
	OrderStatus int    `xml:"OrderStatus" json:"OrderStatus"`
	ProductId   string `xml:"ProductId"   json:"ProductId"`
	SKUInfo     string `xml:"SkuInfo"     json:"SkuInfo"`

	// card
	CardId          string `xml:"CardId"          json:"CardId"`
	IsGiveByFriend  int    `xml:"IsGiveByFriend"  json:"IsGiveByFriend"`
	FriendUserName  string `xml:"FriendUserName"  json:"FriendUserName"`
	UserCardCode    string `xml:"UserCardCode"    json:"UserCardCode"`
	OldUserCardCode string `xml:"OldUserCardCode" json:"OldUserCardCode"`
	ConsumeSource   string `xml:"ConsumeSource"   json:"ConsumeSource"`
	OuterId         int64  `xml:"OuterId"         json:"OuterId"`

	// poi
	UniqId string `xml:"UniqId" json:"UniqId"`
	PoiId  int64  `xml:"PoiId"  json:"PoiId"`
	Result string `xml:"Result" json:"Result"`
	Msg    string `xml:"Msg"    json:"Msg"`

	// dkf
	KfAccount     string `xml:"KfAccount"     json:"KfAccount"`
	FromKfAccount string `xml:"FromKfAccount" json:"FromKfAccount"`
	ToKfAccount   string `xml:"ToKfAccount"   json:"ToKfAccount"`

	// shakearound
	ChosenBeacon struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"ChosenBeacon"                         json:"ChosenBeacon"`
	AroundBeacons []struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"AroundBeacons>AroundBeacon,omitempty" json:"AroundBeacons,omitempty"`

	// bizwifi
	ConnectTime int64  `xml:"ConnectTime" json:"ConnectTime"`
	ExpireTime  int64  `xml:"ExpireTime"  json:"ExpireTime"`
	VendorId    string `xml:"VendorId"    json:"VendorId"`
	PlaceId     int64  `xml:"PlaceId"     json:"PlaceId"`
	DeviceNo    string `xml:"DeviceNo"    json:"DeviceNo"`
}
