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
	EventType EventType `xml:"Event" json:"Event"`

	MsgId        int64   `xml:"MsgId"        json:"MsgId"`        // request
	Content      string  `xml:"Content"      json:"Content"`      // request
	MediaId      string  `xml:"MediaId"      json:"MediaId"`      // request
	PicURL       string  `xml:"PicUrl"       json:"PicUrl"`       // request
	Format       string  `xml:"Format"       json:"Format"`       // request
	Recognition  string  `xml:"Recognition"  json:"Recognition"`  // request
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId"` // request
	LocationX    float64 `xml:"Location_X"   json:"Location_X"`   // request
	LocationY    float64 `xml:"Location_Y"   json:"Location_Y"`   // request
	Scale        int     `xml:"Scale"        json:"Scale"`        // request
	Label        string  `xml:"Label"        json:"Label"`        // request
	Title        string  `xml:"Title"        json:"Title"`        // request
	Description  string  `xml:"Description"  json:"Description"`  // request
	URL          string  `xml:"Url"          json:"Url"`          // request
	EventKey     string  `xml:"EventKey"     json:"EventKey"`     // request, menu
	Ticket       string  `xml:"Ticket"       json:"Ticket"`       // request
	Latitude     float64 `xml:"Latitude"     json:"Latitude"`     // request
	Longitude    float64 `xml:"Longitude"    json:"Longitude"`    // request
	Precision    float64 `xml:"Precision"    json:"Precision"`    // request

	// menu
	MenuId       int64 `xml:"MenuId" json:"MenuId"`
	ScanCodeInfo *struct {
		ScanType   string `xml:"ScanType"   json:"ScanType"`
		ScanResult string `xml:"ScanResult" json:"ScanResult"`
	} `xml:"ScanCodeInfo,omitempty" json:"ScanCodeInfo,omitempty"`
	SendPicsInfo *struct {
		Count   int `xml:"Count" json:"Count"`
		PicList []struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		} `xml:"PicList>item,omitempty" json:"PicList,omitempty"`
	} `xml:"SendPicsInfo,omitempty" json:"SendPicsInfo,omitempty"`
	SendLocationInfo *struct {
		LocationX float64 `xml:"Location_X" json:"Location_X"`
		LocationY float64 `xml:"Location_Y" json:"Location_Y"`
		Scale     int     `xml:"Scale"      json:"Scale"`
		Label     string  `xml:"Label"      json:"Label"`
		PoiName   string  `xml:"Poiname"    json:"Poiname"`
	} `xml:"SendLocationInfo,omitempty" json:"SendLocationInfo,omitempty"`

	MsgID    int64  `xml:"MsgID"  json:"MsgID"`  // template, mass
	Status   string `xml:"Status" json:"Status"` // template, mass
	*mass           // mass
	*account        // account
	*dkf            // dkf
	*poi            // poi
	*card           // card
	*bizwifi        // bizwifi
	*file           // MsgType is file

	// shakearound
	ChosenBeacon *struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"ChosenBeacon,omitempty" json:"ChosenBeacon,omitempty"`
	AroundBeacons []struct {
		UUID     string  `xml:"Uuid"     json:"Uuid"`
		Major    int     `xml:"Major"    json:"Major"`
		Minor    int     `xml:"Minor"    json:"Minor"`
		Distance float64 `xml:"Distance" json:"Distance"`
	} `xml:"AroundBeacons>AroundBeacon,omitempty" json:"AroundBeacons,omitempty"`
}

type mass struct {
	//MsgID       int64  `xml:"MsgID"       json:"MsgID"`
	//Status      string `xml:"Status"      json:"Status"`
	TotalCount  int `xml:"TotalCount"  json:"TotalCount"`
	FilterCount int `xml:"FilterCount" json:"FilterCount"`
	SentCount   int `xml:"SentCount"   json:"SentCount"`
	ErrorCount  int `xml:"ErrorCount"  json:"ErrorCount"`
}

type account struct {
	ExpiredTime int64  `xml:"ExpiredTime" json:"ExpiredTime"`
	FailTime    int64  `xml:"FailTime"    json:"FailTime"`
	FailReason  string `xml:"FailReason"  json:"FailReason"`
}

type dkf struct {
	KfAccount     string `xml:"KfAccount"     json:"KfAccount"`
	FromKfAccount string `xml:"FromKfAccount" json:"FromKfAccount"`
	ToKfAccount   string `xml:"ToKfAccount"   json:"ToKfAccount"`
}

type poi struct {
	UniqId string `xml:"UniqId" json:"UniqId"`
	PoiId  int64  `xml:"PoiId"  json:"PoiId"`
	Result string `xml:"Result" json:"Result"`
	Msg    string `xml:"Msg"    json:"Msg"`
}

type card struct {
	CardId              string `xml:"CardId"              json:"CardId"`
	RefuseReason        string `xml:"RefuseReason"        json:"RefuseReason"`
	IsGiveByFriend      int    `xml:"IsGiveByFriend"      json:"IsGiveByFriend"`
	FriendUserName      string `xml:"FriendUserName"      json:"FriendUserName"`
	UserCardCode        string `xml:"UserCardCode"        json:"UserCardCode"`
	OldUserCardCode     string `xml:"OldUserCardCode"     json:"OldUserCardCode"`
	ConsumeSource       string `xml:"ConsumeSource"       json:"ConsumeSource"`
	OuterId             int64  `xml:"OuterId"             json:"OuterId"`
	LocationName        string `xml:"LocationName"        json:"LocationName"`
	StaffOpenId         string `xml:"StaffOpenId"         json:"StaffOpenId"`
	VerifyCode          string `xml:"VerifyCode"          json:"VerifyCode"`
	RemarkAmount        string `xml:"RemarkAmount"        json:"RemarkAmount"`
	OuterStr            string `xml:"OuterStr"            json:"OuterStr"`
	Detail              string `xml:"Detail"              json:"Detail"`
	IsReturnBack        int    `xml:"IsReturnBack"        json:"IsReturnBack"`
	IsChatRoom          int    `xml:"IsChatRoom"          json:"IsChatRoom"`
	IsRestoreMemberCard int    `xml:"IsRestoreMemberCard" json:"IsRestoreMemberCard"`
	IsRecommendByFriend int    `xml:"IsRecommendByFriend" json:"IsRecommendByFriend"`
	PageId              string `xml:"PageId"              json:"PageId"`
	OrderId             string `xml:"OrderId"             json:"OrderId"`
}

type bizwifi struct {
	ConnectTime int64  `xml:"ConnectTime" json:"ConnectTime"`
	ExpireTime  int64  `xml:"ExpireTime"  json:"ExpireTime"`
	VendorId    string `xml:"VendorId"    json:"VendorId"`
	PlaceId     int64  `xml:"PlaceId"     json:"PlaceId"`
	DeviceNo    string `xml:"DeviceNo"    json:"DeviceNo"`
}

type file struct {
	FileKey      string `xml:"FileKey"      json:"FileKey"`
	FileMd5      string `xml:"FileMd5"      json:"FileMd5"`
	FileTotalLen string `xml:"FileTotalLen" json:"FileTotalLen"`
}
