package core

type (
	MsgType   string
	EventType string
)

// 微信服务器推送过来的消息(事件)的通用消息头.
type MsgHeader struct {
	AppId        string  `xml:"AppId,omitempty" json:"AppId,omitempty"`
	InfoType     string  `xml:"InfoType,omitempty" json:"InfoType,omitempty"`
	ToUserName   string  `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string  `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime"   json:"CreateTime"`
	MsgType      MsgType `xml:"MsgType"      json:"MsgType"`
}

// 微信服务器推送过来的消息(事件)的合集.
type MixedMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MsgHeader
	EventType EventType `xml:"Event,omitempty" json:"Event,omitempty"`

	MsgId        int64   `xml:"MsgId,omitempty"        json:"MsgId,omitempty"`        // request
	Content      string  `xml:"Content,omitempty"      json:"Content,omitempty"`      // request
	MediaId      string  `xml:"MediaId,omitempty"      json:"MediaId,omitempty"`      // request
	PicURL       string  `xml:"PicUrl,omitempty"       json:"PicUrl,omitempty"`       // request
	Format       string  `xml:"Format,omitempty"       json:"Format,omitempty"`       // request
	Recognition  string  `xml:"Recognition,omitempty"  json:"Recognition,omitempty"`  // request
	ThumbMediaId string  `xml:"ThumbMediaId,omitempty" json:"ThumbMediaId,omitempty"` // request
	LocationX    float64 `xml:"Location_X,omitempty"   json:"Location_X,omitempty"`   // request
	LocationY    float64 `xml:"Location_Y,omitempty"   json:"Location_Y,omitempty"`   // request
	Scale        int     `xml:"Scale,omitempty"        json:"Scale,omitempty"`        // request
	Label        string  `xml:"Label,omitempty"        json:"Label,omitempty"`        // request
	Title        string  `xml:"Title,omitempty"        json:"Title,omitempty"`        // request
	Description  string  `xml:"Description,omitempty"  json:"Description,omitempty"`  // request
	URL          string  `xml:"Url,omitempty"          json:"Url,omitempty"`          // request
	EventKey     string  `xml:"EventKey,omitempty"     json:"EventKey,omitempty"`     // request, menu
	Ticket       string  `xml:"Ticket"       json:"Ticket"`                           // request
	Latitude     float64 `xml:"Latitude"     json:"Latitude"`                         // request
	Longitude    float64 `xml:"Longitude"    json:"Longitude"`                        // request
	Precision    float64 `xml:"Precision"    json:"Precision"`                        // request

	ComponentVerifyTicket        string `xml:"ComponentVerifyTicket,omitempty"      json:"ComponentVerifyTicket,omitempty"`
	AuthorizerAppid              string `xml:"AuthorizerAppid,omitempty" json:"AuthorizerAppid,omitempty"`                           // 公众号或小程序
	AuthorizationCode            string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty"`                       // 授权码，可用于换取公众号的接口调用凭据
	AuthorizationCodeExpiredTime string `xml:"AuthorizationCodeExpiredTime,omitempty" json:"AuthorizationCodeExpiredTime,omitempty"` // 授权码过期时间
	PreAuthCode                  string `xml:"PreAuthCode,omitempty" json:"PreAuthCode,omitempty"`                                   // 预授权码

	// menu
	MenuId       int64 `xml:"MenuId,omitempty" json:"MenuId,omitempty"`
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

	MsgID       int64  `xml:"MsgID,omitempty"  json:"MsgID,omitempty"`            // template, mass
	Status      string `xml:"Status,omitempty" json:"Status,omitempty"`           // template, mass
	TotalCount  int    `xml:"TotalCount,omitempty"  json:"TotalCount,omitempty"`  // template, mass
	FilterCount int    `xml:"FilterCount,omitempty" json:"FilterCount,omitempty"` // template, mass
	SentCount   int    `xml:"SentCount,omitempty"   json:"SentCount,omitempty"`   // template, mass
	ErrorCount  int    `xml:"ErrorCount,omitempty"  json:"ErrorCount,omitempty"`  // template, mass

	ExpiredTime int64  `xml:"ExpiredTime,omitempty" json:"ExpiredTime,omitempty"`
	FailTime    int64  `xml:"FailTime,omitempty"    json:"FailTime,omitempty"`
	FailReason  string `xml:"FailReason,omitempty"  json:"FailReason,omitempty"`

	KfAccount     string `xml:"KfAccount,omitempty"     json:"KfAccount,omitempty"`
	FromKfAccount string `xml:"FromKfAccount,omitempty" json:"FromKfAccount,omitempty"`
	ToKfAccount   string `xml:"ToKfAccount,omitempty"   json:"ToKfAccount,omitempty"`

	// poi
	UniqId string `xml:"UniqId,omitempty" json:"UniqId,omitempty"`
	PoiId  int64  `xml:"PoiId,omitempty"  json:"PoiId,omitempty"`
	Result string `xml:"Result,omitempty" json:"Result,omitempty"`
	Msg    string `xml:"Msg,omitempty"    json:"Msg,omitempty"`

	// card
	CardId              string `xml:"CardId,omitempty"              json:"CardId,omitempty"`
	RefuseReason        string `xml:"RefuseReason,omitempty"        json:"RefuseReason,omitempty"`
	IsGiveByFriend      int    `xml:"IsGiveByFriend,omitempty"      json:"IsGiveByFriend,omitempty"`
	FriendUserName      string `xml:"FriendUserName,omitempty"      json:"FriendUserName,omitempty"`
	UserCardCode        string `xml:"UserCardCode,omitempty"        json:"UserCardCode,omitempty"`
	OldUserCardCode     string `xml:"OldUserCardCode,omitempty"     json:"OldUserCardCode,omitempty"`
	ConsumeSource       string `xml:"ConsumeSource,omitempty"       json:"ConsumeSource,omitempty"`
	OuterId             int64  `xml:"OuterId,omitempty"             json:"OuterId,omitempty"`
	LocationName        string `xml:"LocationName,omitempty"        json:"LocationName,omitempty"`
	StaffOpenId         string `xml:"StaffOpenId,omitempty"         json:"StaffOpenId,omitempty"`
	VerifyCode          string `xml:"VerifyCode,omitempty"          json:"VerifyCode,omitempty"`
	RemarkAmount        string `xml:"RemarkAmount,omitempty"        json:"RemarkAmount,omitempty"`
	OuterStr            string `xml:"OuterStr,omitempty"            json:"OuterStr,omitempty"`
	Detail              string `xml:"Detail,omitempty"              json:"Detail,omitempty"`
	IsReturnBack        int    `xml:"IsReturnBack,omitempty"        json:"IsReturnBack,omitempty"`
	IsChatRoom          int    `xml:"IsChatRoom,omitempty"          json:"IsChatRoom,omitempty"`
	IsRestoreMemberCard int    `xml:"IsRestoreMemberCard" json:"IsRestoreMemberCard,omitempty"`
	IsRecommendByFriend int    `xml:"IsRecommendByFriend,omitempty" json:"IsRecommendByFriend,omitempty"`
	PageId              string `xml:"PageId,omitempty"              json:"PageId,omitempty"`
	OrderId             string `xml:"OrderId,omitempty"             json:"OrderId,omitempty"`

	// bizwifi
	ConnectTime int64  `xml:"ConnectTime,omitempty" json:"ConnectTime,omitempty"`
	ExpireTime  int64  `xml:"ExpireTime,omitempty"  json:"ExpireTime,omitempty"`
	VendorId    string `xml:"VendorId,omitempty"    json:"VendorId,omitempty"`
	PlaceId     int64  `xml:"PlaceId,omitempty"     json:"PlaceId,omitempty"`
	DeviceNo    string `xml:"DeviceNo,omitempty"    json:"DeviceNo,omitempty"`

	// file
	FileKey      string `xml:"FileKey,omitempty"      json:"FileKey,omitempty"`
	FileMd5      string `xml:"FileMd5,omitempty"      json:"FileMd5,omitempty"`
	FileTotalLen string `xml:"FileTotalLen,omitempty" json:"FileTotalLen,omitempty"`

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
	ArticleUrlResult *ArticleUrlResult `xml:"ArticleUrlResult,omitempty" json:"ArticleUrlResult,omitempty"`

	// add express path
	DeliveryID string                 `xml:"DeliveryID,omitempty" json:"DeliveryID,omitempty"` // 快递公司ID
	WaybillID  string                 `xml:"WayBillId,omitempty" json:"WayBillId,omitempty"`   // 运单ID
	Version    int                    `xml:"Version,omitempty" json:"Version,omitempty"`       // 轨迹版本号（整型）
	Count      int                    `xml:"Count,omitempty" json:"Count,omitempty"`           // 轨迹节点数（整型）
	Actions    []AddExpressPathAction `xml:"Actions,omitempty" json:"Actions,omitempty"`       // 轨迹列表

	// 小程序订阅消息
	SubscribeMsgPopupEvent *SubscribeMsgPopupEvent `xml:"SubscribeMsgPopupEvent,omitempty" json:"SubscribeMsgPopupEvent,omitempty"`
}

type ArticleUrlResult struct {
	Count      int `xml:"Count" json:"Count"`
	ResultList []struct {
		ArticleIdx uint   `xml:"ArticleIdx" json:"ArticleIdx"`
		ArticleUrl string `xml:"ArticleUrl" json:"ArticleUrl"`
	} `xml:"ResultList>item,omitempty" json:"ResultList>item,omitempty"`
}

// AddExpressPathAction 轨迹列表
type AddExpressPathAction struct {
	// ActionTime 轨迹节点 Unix 时间戳
	ActionTime int64 `xml:"ActionTime" json:"ActionTime"`
	// ActionType 轨迹节点类型
	ActionType int `xml:"ActionType" json:"ActionType"`
	// ActionMsg 轨迹节点详情
	ActionMsg string `xml:"ActionMsg" json:"ActionMsg"`
}

// 小程序订阅消息事件
type SubscribeMsgPopupEvent struct {
	List []SubscribeMsgPopupEventItem `xml:"List" json:"List"`
}

type SubscribeMsgPopupEventItem struct {
	TemplateID            string `xml:"TemplateId" json:"TemplateId"`
	SubscribeStatusString string `xml:"SubscribeStatusString" json:"SubscribeStatusString"`
	PopupScene            string `xml:"PopupScene" json:"PopupScene"`
}
