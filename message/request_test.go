package message

import (
	"encoding/xml"
	"testing"
)

// 测试 RequestMsg 的 xml.Unmarshal() 和 RequestMsg.Zero()

var requestUnmarshalTests = []struct {
	XML              []byte
	ExpectRequestMsg RequestMsg
}{
	{ // 文本消息
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName> 
		<CreateTime>1348831860</CreateTime>
		<MsgType><![CDATA[text]]></MsgType>
		<Content><![CDATA[this is a test]]></Content>
		<MsgId>1234567890123456</MsgId>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "fromUser",
			CreateTime:   1348831860,
			MsgType:      RQST_MSG_TYPE_TEXT,
			Content:      "this is a test",
			MsgId:        1234567890123456,
		},
	},
	{ // 图片消息
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1348831860</CreateTime>
		<MsgType><![CDATA[image]]></MsgType>
		<PicUrl><![CDATA[this is a url]]></PicUrl>
		<MediaId><![CDATA[media_id]]></MediaId>
		<MsgId>1234567890123456</MsgId>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "fromUser",
			CreateTime:   1348831860,
			MsgType:      RQST_MSG_TYPE_IMAGE,
			PicUrl:       "this is a url",
			MediaId:      "media_id",
			MsgId:        1234567890123456,
		},
	},
	{ // 语音消息
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1357290913</CreateTime>
		<MsgType><![CDATA[voice]]></MsgType>
		<MediaId><![CDATA[media_id]]></MediaId>
		<Format><![CDATA[Format]]></Format>
		<MsgId>1234567890123456</MsgId>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "fromUser",
			CreateTime:   1357290913,
			MsgType:      RQST_MSG_TYPE_VOICE,
			MediaId:      "media_id",
			Format:       "Format",
			MsgId:        1234567890123456,
		},
	},
	{ // 接收语音识别结果
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1357290913</CreateTime>
		<MsgType><![CDATA[voice]]></MsgType>
		<MediaId><![CDATA[media_id]]></MediaId>
		<Format><![CDATA[Format]]></Format>
		<Recognition><![CDATA[腾讯微信团队]]></Recognition>
		<MsgId>1234567890123456</MsgId>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "fromUser",
			CreateTime:   1357290913,
			MsgType:      RQST_MSG_TYPE_VOICE,
			MediaId:      "media_id",
			Format:       "Format",
			Recognition:  "腾讯微信团队",
			MsgId:        1234567890123456,
		},
	},
	{ // 视频消息
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1357290913</CreateTime>
		<MsgType><![CDATA[video]]></MsgType>
		<MediaId><![CDATA[media_id]]></MediaId>
		<ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
		<MsgId>1234567890123456</MsgId>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "fromUser",
			CreateTime:   1357290913,
			MsgType:      RQST_MSG_TYPE_VIDEO,
			MediaId:      "media_id",
			ThumbMediaId: "thumb_media_id",
			MsgId:        1234567890123456,
		},
	},
	{ // 地理位置消息
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1351776360</CreateTime>
		<MsgType><![CDATA[location]]></MsgType>
		<Location_X>23.555555</Location_X>
		<Location_Y>113.555555</Location_Y>
		<Scale>20</Scale>
		<Label><![CDATA[位置信息]]></Label>
		<MsgId>1234567890123456</MsgId>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "fromUser",
			CreateTime:   1351776360,
			MsgType:      RQST_MSG_TYPE_LOCATION,
			Location_X:   23.555555,
			Location_Y:   113.555555,
			Scale:        20,
			Label:        "位置信息",
			MsgId:        1234567890123456,
		},
	},
	{ // 链接消息
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1351776360</CreateTime>
		<MsgType><![CDATA[link]]></MsgType>
		<Title><![CDATA[公众平台官网链接]]></Title>
		<Description><![CDATA[公众平台官网链接]]></Description>
		<Url><![CDATA[url]]></Url>
		<MsgId>1234567890123456</MsgId>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "fromUser",
			CreateTime:   1351776360,
			MsgType:      RQST_MSG_TYPE_LINK,
			Title:        "公众平台官网链接",
			Description:  "公众平台官网链接",
			Url:          "url",
			MsgId:        1234567890123456,
		},
	},
	{ // 关注事件
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[subscribe]]></Event>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "FromUser",
			CreateTime:   123456789,
			MsgType:      RQST_MSG_TYPE_EVENT,
			Event:        RQST_EVENT_TYPE_SUBSCRIBE,
		},
	},
	{ // 取消关注事件
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[unsubscribe]]></Event>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "FromUser",
			CreateTime:   123456789,
			MsgType:      RQST_MSG_TYPE_EVENT,
			Event:        RQST_EVENT_TYPE_UNSUBSCRIBE,
		},
	},
	{ // 用户未关注时，进行关注后的事件推送
		[]byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[subscribe]]></Event>
		<EventKey><![CDATA[qrscene_123123]]></EventKey>
		<Ticket><![CDATA[TICKET]]></Ticket>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "FromUser",
			CreateTime:   123456789,
			MsgType:      RQST_MSG_TYPE_EVENT,
			Event:        RQST_EVENT_TYPE_SUBSCRIBE,
			EventKey:     "qrscene_123123",
			Ticket:       "TICKET",
		},
	},
	{ // 用户已关注时的事件推送
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[SCAN]]></Event>
		<EventKey><![CDATA[SCENE_VALUE]]></EventKey>
		<Ticket><![CDATA[TICKET]]></Ticket>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "FromUser",
			CreateTime:   123456789,
			MsgType:      RQST_MSG_TYPE_EVENT,
			Event:        RQST_EVENT_TYPE_SCAN,
			EventKey:     "SCENE_VALUE",
			Ticket:       "TICKET",
		},
	},
	{ // 上报地理位置事件
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[LOCATION]]></Event>
		<Latitude>23.555555</Latitude>
		<Longitude>113.555555</Longitude>
		<Precision>119.555555</Precision>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "fromUser",
			CreateTime:   123456789,
			MsgType:      RQST_MSG_TYPE_EVENT,
			Event:        RQST_EVENT_TYPE_LOCATION,
			Latitude:     23.555555,
			Longitude:    113.555555,
			Precision:    119.555555,
		},
	},
	{ // 点击菜单拉取消息时的事件推送
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[CLICK]]></Event>
		<EventKey><![CDATA[EVENTKEY]]></EventKey>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "FromUser",
			CreateTime:   123456789,
			MsgType:      RQST_MSG_TYPE_EVENT,
			Event:        RQST_EVENT_TYPE_CLICK,
			EventKey:     "EVENTKEY",
		},
	},
	{ // 点击菜单跳转链接时的事件推送
		[]byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[VIEW]]></Event>
		<EventKey><![CDATA[www.qq.com]]></EventKey>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "toUser",
			FromUserName: "FromUser",
			CreateTime:   123456789,
			MsgType:      RQST_MSG_TYPE_EVENT,
			Event:        RQST_EVENT_TYPE_VIEW,
			EventKey:     "www.qq.com",
		},
	},
	{ // 事件推送群发结果
		[]byte(`<xml>
		<ToUserName><![CDATA[gh_3e8adccde292]]></ToUserName>
		<FromUserName><![CDATA[oR5Gjjl_eiZoUpGozMo7dbBJ362A]]></FromUserName>
		<CreateTime>1394524295</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[MASSSENDJOBFINISH]]></Event>
		<MsgID>1988</MsgID>
		<Status><![CDATA[sendsuccess]]></Status>
		<TotalCount>100</TotalCount>
		<FilterCount>80</FilterCount>
		<SentCount>75</SentCount>
		<ErrorCount>75</ErrorCount>
		</xml>`),
		RequestMsg{
			XMLName:      xml.Name{Space: "", Local: "xml"},
			ToUserName:   "gh_3e8adccde292",
			FromUserName: "oR5Gjjl_eiZoUpGozMo7dbBJ362A",
			CreateTime:   1394524295,
			MsgType:      RQST_MSG_TYPE_EVENT,
			Event:        RQST_EVENT_TYPE_MASSSENDJOBFINISH,
			MsgID:        1988,
			Status:       "sendsuccess",
			TotalCount:   100,
			FilterCount:  80,
			SentCount:    75,
			ErrorCount:   75,
		},
	},
}

// 只测试 unmarshal 了, marshal 比 unmarshal 简单, 没有层级结构,
// 同时测试 RequestMsg.Zero 是否能正确工作.
func TestRequestMsgUnmarshalAndZero(t *testing.T) {
	var msg RequestMsg

	for _, test := range requestUnmarshalTests {
		msg.Zero() // 去掉这个肯定失败

		if err := xml.Unmarshal(test.XML, &msg); err != nil {
			t.Errorf("unmarshal(%#q):\nError: %s\n", test.XML, err)
			continue
		}
		if got, want := msg, test.ExpectRequestMsg; got != want {
			t.Errorf("unmarshal(%#q):\nhave %#v\nwant %#v\n", test.XML, got, want)
			continue
		}
	}
}
