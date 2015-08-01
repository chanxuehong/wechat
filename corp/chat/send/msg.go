// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

// 发送消息.
package send


//创建会话用结构体
type ChatInfo struct {
	Chatid   string `json:"chatid"`     // 必须;会话id。字符串类型，最长32个字符。只允许字符0-9及字母a-zA-Z,
                                        //                   如果值内容为64bit无符号整型：要求值范围在[1, 2^63)之间，
                                        //                   [2^63, 2^64)为系统分配会话id区间
	Name     string `json:"name"`       // 必须;会话标题
	Owner    string `json:"owner"`      //必须;管理员userid，必须是该会话userlist的成员之一
	Userlist []string `json:"userlist"` //必须;会话成员列表，成员用userid来标识。会话成员必须在3人或以上，1000人以下
}

//更新会话用结构体
type UpdateInfo struct {
	ChatId      string       `json:"chatid"`        //必须;会话id
	OpUser      string       `json:"op_user"`       //必须;操作人userid
	Name        string       `json:"name"`          //非必须;会话标题
	Owner       string       `json:"owner"`         //非必须;管理员userid，必须是该会话userlist的成员之一
	AddUserList []string     `json:"add_user_list"` //非必须;会话新增成员列表，成员用userid来标识
	DelUserList []string     `json:"del_user_list"` //非必须;会话退出成员列表，成员用userid来标识
}

//清除消息未读状态结构体
type ClearNotify struct {
	OpUser string       `json:"op_user"` //必须; 会话所有者的userid
	Chat   struct {
		       Type string `json:"type"` //必须; 会话类型：single|group，分别表示：群聊|单聊
		       Id   string   `json:"id"` //必须; 会话值，为userid|chatid，分别表示：成员id|会话id
	       }  `json:"chat"`
}

// 成员新消息免打扰参数，数组，最大支持10000个成员
type UserMute struct {
		Userid string `json:"userid"` //成员UserID
		Status int64 `json:"status"`  //免打扰状态，0关闭，1打开,默认为0。
	                                  // 当打开时所有消息不提醒；当关闭时，以成员对会话的设置为准。
}


const (
	MsgTypeText = "text"
	MsgTypeImage = "image"
//	MsgTypeVoice  = "voice"
//	MsgTypeVideo  = "video"
	MsgTypeFile = "file"


	ReceiverTypeSingle = "single"
	ReceiverTypeGroup = "group"
)

type MessageHeader struct {
	Receiver struct {
		         Type string `json:"type"` //必须; 会话类型：single|group，分别表示：群聊|单聊
		         Id   string   `json:"id"` //必须; 会话值，为userid|chatid，分别表示：成员id|会话id
	         } `json:"receiver"`     //必须; 接收人
	MsgType  string `json:"msgtype"` // 必须; 消息类型
	Sender   string  `json:"sender"` // 必须; 发送人
}

type Text struct {
	MessageHeader

	Text struct {
		     Content string `json:"content"`
	     } `json:"text"`
}

type Image struct {
	MessageHeader

	Image struct {
		      MediaId string `json:"media_id"` // 图片媒体文件id, 可以调用上传媒体文件接口获取
	      } `json:"image"`
}

type Voice struct {
	MessageHeader

	Voice struct {
		      MediaId string `json:"media_id"` // 语音文件id, 可以调用上传媒体文件接口获取
	      } `json:"voice"`
}

type Video struct {
	MessageHeader

	Video struct {
		      MediaId     string `json:"media_id"`              // 视频媒体文件id, 可以调用上传媒体文件接口获取
		      Title       string `json:"title,omitempty"`       // 视频消息的标题
		      Description string `json:"description,omitempty"` // 视频消息的描述
	      } `json:"video"`
}

type File struct {
	MessageHeader

	File struct {
		     MediaId string `json:"media_id"` // 媒体文件id, 可以调用上传媒体文件接口获取
	     } `json:"file"`
}
