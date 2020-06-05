package message

import (
	"encoding/xml"
)

type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

type EncryptedMsg struct {
	XMLName    struct{} `xml:"xml" json:"-"`
	ToUserName CDATA    `xml:"ToUserName" json:"ToUserName"`
	AgentID    CDATA    `xml:"AgentID" json:"AgentID"`
	Encrypt    CDATA    `xml:"Encrypt" json:"Encrypt"`
}

type MsgRequest struct {
	MsgSignature string `form:"msg_signature" json:"msg_signature" binding:"required"`
	Timestamp    string `form:"timestamp" json:"timestamp" binding:"required"`
	Nonce        string `form:"nonce" json:"nonce" binding:"required"`
	EchoStr      string `form:"echostr" json:"echostr"`
}

type (
	MsgType   string
	EventType string
)

type MsgHeader struct {
	ToUserName   CDATA   `xml:"ToUserName" json:"ToUserName"`
	FromUserName CDATA   `xml:"FromUserName" json:"FromUserName"`
	AgentID      CDATA   `xml:"AgentID" json:"AgentID"`
	MsgType      MsgType `xml:"MsgType" json:"MsgType"`
	CreateTime   int64   `xml:"CreateTime" json:"CreateTime"`
}

// 微信服务器推送过来的消息(事件)的合集.
type MixedMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MsgHeader
	EventType        EventType        `xml:"Event" json:"Event"`
	MsgId            int64            `xml:"MsgId"        json:"MsgId"`                              // request
	Content          CDATA            `xml:"Content"      json:"Content"`                            // request
	MediaId          CDATA            `xml:"MediaId"      json:"MediaId"`                            // request
	PicURL           CDATA            `xml:"PicUrl"       json:"PicUrl"`                             // request
	Format           CDATA            `xml:"Format"       json:"Format"`                             // request
	ThumbMediaId     CDATA            `xml:"ThumbMediaId" json:"ThumbMediaId"`                       // request
	LocationX        float64          `xml:"Location_X"   json:"Location_X"`                         // request
	LocationY        float64          `xml:"Location_Y"   json:"Location_Y"`                         // request
	Scale            int              `xml:"Scale"        json:"Scale"`                              // request
	Label            CDATA            `xml:"Label"        json:"Label"`                              // request
	Title            CDATA            `xml:"Title"        json:"Title"`                              // request
	Description      CDATA            `xml:"Description"  json:"Description"`                        // request
	URL              CDATA            `xml:"Url"          json:"Url"`                                // request
	EventKey         CDATA            `xml:"EventKey"     json:"EventKey"`                           // request, event
	Latitude         float64          `xml:"Latitude"     json:"Latitude"`                           // request
	Longitude        float64          `xml:"Longitude"    json:"Longitude"`                          // request
	Precision        float64          `xml:"Precision"    json:"Precision"`                          // 地理位置精度(整数? 但是微信推送过来是浮点数形式)
	AppType          CDATA            `xml:"AppType" json:"AppType"`                                 //app类型，在企业微信固定返回wxwork，在微信不返回该字段
	JobId            CDATA            `xml:"JobId"  json:"JobId"`                                    // 异步任务id，最大长度为64字符
	JobType          CDATA            `xml:"JobType" json:"JobType"`                                 // 操作类型，字符串，目前分别有：sync_user(增量更新成员)、 replace_user(全量覆盖成员）、invite_user(邀请成员关注）、replace_party(全量覆盖部门)
	ErrCode          int              `xml:"ErrCode,omitempty" json:"ErrCode,omitempty"`             // 返回码
	ErrMsg           CDATA            `xml:"ErrMsg,omitempty" json:"ErrMsg,omitempty"`               // 对返回码的文本描述内容
	ChangeType       CDATA            `xml:"ChangeType"  json:"ChangeType"`                          // create_party
	Id               uint64           `xml:"Id" json:"Id"`                                           // 部门Id
	Name             CDATA            `xml:"Name" json:"Name"`                                       // 部门名称
	ParentId         CDATA            `xml:"ParentId" json:"ParentId"`                               // 父部门id
	Order            int              `xml:"Order" json:"Order"`                                     // 部门排序
	TagId            uint64           `xml:"TagId" json:"TagId"`                                     // 标签Id
	AddUserItems     CDATA            `xml:"AddUserItems,omitempty" json:"AddUserItems,omitempty"`   // 标签中新增的成员userid列表，用逗号分隔
	DelUserItems     CDATA            `xml:"DelUserItems,omitempty" json:"DelUserItems,omitempty"`   // 标签中删除的成员userid列表，用逗号分隔
	AddPartyItems    CDATA            `xml:"AddPartyItems,omitempty" json:"AddPartyItems,omitempty"` // 标签中新增的部门id列表，用逗号分隔
	DelPartyItems    CDATA            `xml:"DelPartyItems,omitempty" json:"DelPartyItems,omitempty"` // 标签中删除的部门id列表，用逗号分隔
	ScanCodeInfo     ScanCodeInfo     `xml:"ScanCodeInfo" json:"ScanCodeInfo"`                       // 扫描信息
	SendPicsInfo     SendPicsInfo     `xml:"SendPicsInfo" json:"SendPicsInfo"`                       // 扫描信息
	SendLocationInfo SendLocationInfo `xml:"SendLocationInfo" json:"SendLocationInfo"`               // 发送的位置信息
	TaskId           CDATA            `xml:"TaskId" json:"TaskId"`                                   // 与发送任务卡片消息时指定的task_id相同
	ApprovalInfo     ApprovalInfo     `xml:"ApprovalInfo" json:"ApprovalInfo"`                       // 审批信息
}

type ScanCodeInfo struct {
	ScanType   CDATA `xml:"ScanType" json:"ScanType"`     // 扫描类型，一般是qrcode
	ScanResult CDATA `xml:"ScanResult" json:"ScanResult"` // 扫描结果，即二维码对应的字符串信息
}

type SendPicsInfo struct {
	Count   int           `xml:"Count" json:"Count"`          // 发送的图片数量
	PicList []PicListItem `xml:"PicList>item" json:"PicList"` // 图片列表
}

type PicListItem struct {
	PicMd5Sum CDATA `xml:"PicMd5Sum" json:"PicMd5Sum"` // 图片的MD5值，开发者若需要，可用于验证接收到图片
}

type SendLocationInfo struct {
	LocationX float64 `xml:"Location_X"   json:"Location_X"` // request
	LocationY float64 `xml:"Location_Y"   json:"Location_Y"`
	Scale     int     `xml:"Scale"        json:"Scale"` // request
	Label     CDATA   `xml:"Label"        json:"Label"` // request
	Poiname   CDATA   `xml:"Poiname"        json:"Poiname"`
}

type ApprovalInfo struct {
	ThirdNo        CDATA          `xml:"ThirdNo" json:"ThirdNo"`                          // 审批单编号，由开发者在发起申请时自定义
	OpenSpName     CDATA          `xml:"OpenSpName" json:"OpenSpName"`                    // 审批模板名称
	OpenTemplateId CDATA          `xml:"OpenTemplateId" json:"OpenTemplateId"`            // 审批模板id
	OpenSpStatus   int            `xml:"OpenSpStatus" json:"OpenSpStatus"`                // 申请单当前审批状态：1-审批中；2-已通过；3-已驳回；4-已取消
	ApplyTime      int64          `xml:"ApplyTime" json:"ApplyTime"`                      // 提交申请时间
	ApplyUserName  CDATA          `xml:"ApplyUserName" json:"ApplyUserName"`              // 提交者姓名
	ApplyUserId    CDATA          `xml:"ApplyUserId" json:"ApplyUserId"`                  // 提交者userid
	ApplyUserParty CDATA          `xml:"ApplyUserParty" json:"ApplyUserParty"`            // 提交者所在部门
	ApplyUserImage CDATA          `xml:"ApplyUserImage" json:"ApplyUserImage"`            // 提交者头像
	ApprovalNodes  []ApprovalNode `xml:"ApprovalNodes>ApprovalNode" json:"ApprovalNodes"` // 审批流程信息
	NotifyNodes    []NotifyNode   `xml:"NotifyNodes>NotifyNode" json:"NotifyNodes"`       // 抄送信息，可能有多个抄送人
	ApproverStep   int            `xml:"approverstep" json:"approverstep"`                // 当前审批节点：0-第一个审批节点；1-第二个审批节点…以此类推
}

type ApprovalNode struct {
	NodeStatus int                `xml:"NodeStatus" json:"NodeStatus"` // 节点审批操作状态：1-审批中；2-已同意；3-已驳回；4-已转审
	NodeAttr   int                `xml:"NodeAttr" json:"NodeAttr"`     // 审批节点属性：1-或签；2-会签
	NodeType   int                `xml:"NodeType" json:"NodeType"`     // 审批节点类型：1-固定成员；2-标签；3-上级
	Items      []ApprovalNodeItem `xml:"Items>Item" json:"Items"`      //
}

type ApprovalNodeItem struct {
	ItemName   CDATA `xml:"ItemName" json:"ItemName"`     // 分支审批人姓名
	ItemUserId CDATA `xml:"ItemUserId" json:"ItemUserId"` // 分支审批人userid
	ItemImage  CDATA `xml:"ItemImage" json:"ItemImage"`   // 分支审批人头像
	ItemStatus int   `xml:"ItemStatus" json:"ItemStatus"` // 分支审批审批操作状态：1-审批中；2-已同意；3-已驳回；4-已转审
	ItemSpeech CDATA `xml:"ItemSpeech" json:"ItemSpeech"` // 分支审批人审批意见
	ItemOpTime int64 `xml:"ItemOpTime" json:"ItemOpTime"` // 分支审批人操作时间
}

type NotifyNode struct {
	ItemName   CDATA `xml:"ItemName" json:"ItemName"`     // 抄送人姓名
	ItemUserId CDATA `xml:"ItemUserId" json:"ItemUserId"` // 抄送人userid
	ItemImage  CDATA `xml:"ItemImage" json:"ItemImage"`   // 抄送人头像
}
