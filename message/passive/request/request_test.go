// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

import (
	"encoding/xml"
	"testing"
)

func TestRequestUnmarshalAndZero(t *testing.T) {
	var req Request
	var msgBytes []byte

	// 测试文本消息===============================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName> 
		<CreateTime>1348831860</CreateTime>
		<MsgType><![CDATA[text]]></MsgType>
		<Content><![CDATA[this is a test]]></Content>
		<MsgId>1234567890123456</MsgId>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   1348831860,
				MsgType:      MSG_TYPE_TEXT,
			},

			MsgId:   1234567890123456,
			Content: "this is a test",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 测试图片消息===============================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1348831860</CreateTime>
		<MsgType><![CDATA[image]]></MsgType>
		<PicUrl><![CDATA[this is a url]]></PicUrl>
		<MediaId><![CDATA[media_id]]></MediaId>
		<MsgId>1234567890123456</MsgId>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   1348831860,
				MsgType:      MSG_TYPE_IMAGE,
			},

			MsgId:   1234567890123456,
			MediaId: "media_id",
			PicURL:  "this is a url",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 测试语音识别结果消息=========================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1357290913</CreateTime>
		<MsgType><![CDATA[voice]]></MsgType>
		<MediaId><![CDATA[media_id]]></MediaId>
		<Format><![CDATA[Format]]></Format>
		<Recognition><![CDATA[腾讯微信团队]]></Recognition>
		<MsgId>1234567890123456</MsgId>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   1357290913,
				MsgType:      MSG_TYPE_VOICE,
			},

			MsgId:       1234567890123456,
			MediaId:     "media_id",
			Format:      "Format",
			Recognition: "腾讯微信团队",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 测试视频消息===============================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1357290913</CreateTime>
		<MsgType><![CDATA[video]]></MsgType>
		<MediaId><![CDATA[media_id]]></MediaId>
		<ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
		<MsgId>1234567890123456</MsgId>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   1357290913,
				MsgType:      MSG_TYPE_VIDEO,
			},

			MsgId:        1234567890123456,
			MediaId:      "media_id",
			ThumbMediaId: "thumb_media_id",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 测试地理位置消息============================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1351776360</CreateTime>
		<MsgType><![CDATA[location]]></MsgType>
		<Location_X>23.134525</Location_X>
		<Location_Y>113.358805</Location_Y>
		<Scale>20</Scale>
		<Label><![CDATA[位置信息]]></Label>
		<MsgId>1234567890123456</MsgId>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   1351776360,
				MsgType:      MSG_TYPE_LOCATION,
			},

			MsgId:     1234567890123456,
			LocationX: 23.134525,  // 最后一位是 5 才能精确表示
			LocationY: 113.358805, // 最后一位是 5 才能精确表示
			Scale:     20,
			Label:     "位置信息",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 测试链接消息===============================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1351776360</CreateTime>
		<MsgType><![CDATA[link]]></MsgType>
		<Title><![CDATA[公众平台官网链接]]></Title>
		<Description><![CDATA[公众平台官网链接]]></Description>
		<Url><![CDATA[url]]></Url>
		<MsgId>1234567890123456</MsgId>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   1351776360,
				MsgType:      MSG_TYPE_LINK,
			},

			MsgId:       1234567890123456,
			Title:       "公众平台官网链接",
			Description: "公众平台官网链接",
			URL:         "url",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 测试关注事件消息===============================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[subscribe]]></Event>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "FromUser",
				CreateTime:   123456789,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event: EVENT_TYPE_SUBSCRIBE,
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 测试取消关注事件============================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[unsubscribe]]></Event>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "FromUser",
				CreateTime:   123456789,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event: EVENT_TYPE_UNSUBSCRIBE,
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 扫描带参数二维码事件, 用户未关注时，进行关注后的事件推送============================

	msgBytes = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[subscribe]]></Event>
		<EventKey><![CDATA[qrscene_123123]]></EventKey>
		<Ticket><![CDATA[TICKET]]></Ticket>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "FromUser",
				CreateTime:   123456789,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_SUBSCRIBE,
			EventKey: "qrscene_123123",
			Ticket:   "TICKET",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 扫描带参数二维码事件, 用户已关注时的事件推送=====================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[SCAN]]></Event>
		<EventKey><![CDATA[SCENE_VALUE]]></EventKey>
		<Ticket><![CDATA[TICKET]]></Ticket>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "FromUser",
				CreateTime:   123456789,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_SCAN,
			EventKey: "SCENE_VALUE",
			Ticket:   "TICKET",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 上报地理位置事件============================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[LOCATION]]></Event>
		<Latitude>23.137465</Latitude>
		<Longitude>113.352425</Longitude>
		<Precision>119.385045</Precision>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   123456789,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:     EVENT_TYPE_LOCATION,
			Latitude:  23.137465,  // 最后一位是 5 才能精确表示
			Longitude: 113.352425, // 最后一位是 5 才能精确表示
			Precision: 119.385045, // 最后一位是 5 才能精确表示
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 点击菜单拉取消息时的事件推送===================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[CLICK]]></Event>
		<EventKey><![CDATA[EVENTKEY]]></EventKey>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "FromUser",
				CreateTime:   123456789,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_CLICK,
			EventKey: "EVENTKEY",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 点击菜单跳转链接时的事件推送===================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[FromUser]]></FromUserName>
		<CreateTime>123456789</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[VIEW]]></Event>
		<EventKey><![CDATA[www.qq.com]]></EventKey>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "toUser",
				FromUserName: "FromUser",
				CreateTime:   123456789,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_VIEW,
			EventKey: "www.qq.com",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// scancode_push：扫码推事件的事件推送=========================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
		<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
		<CreateTime>1408090502</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[scancode_push]]></Event>
		<EventKey><![CDATA[6]]></EventKey>
		<ScanCodeInfo>
			<ScanType><![CDATA[qrcode]]></ScanType>
			<ScanResult><![CDATA[1]]></ScanResult>
		</ScanCodeInfo>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "gh_e136c6e50636",
				FromUserName: "oMgHVjngRipVsoxg6TuX3vz6glDg",
				CreateTime:   1408090502,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_SCANCODE_PUSH,
			EventKey: "6",
		}
		expectReq.ScanCodeInfo.ScanType = "qrcode"
		expectReq.ScanCodeInfo.ScanResult = "1"

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框的事件推送==================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
		<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
		<CreateTime>1408090606</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[scancode_waitmsg]]></Event>
		<EventKey><![CDATA[6]]></EventKey>
		<ScanCodeInfo>
			<ScanType><![CDATA[qrcode]]></ScanType>
			<ScanResult><![CDATA[2]]></ScanResult>
		</ScanCodeInfo>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "gh_e136c6e50636",
				FromUserName: "oMgHVjngRipVsoxg6TuX3vz6glDg",
				CreateTime:   1408090606,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_SCANCODE_WAITMSG,
			EventKey: "6",
		}
		expectReq.ScanCodeInfo.ScanType = "qrcode"
		expectReq.ScanCodeInfo.ScanResult = "2"

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// pic_sysphoto：弹出系统拍照发图的事件推送=====================================

	msgBytes = []byte(`<xml>
	<ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
	<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
	<CreateTime>1408090651</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[pic_sysphoto]]></Event>
	<EventKey><![CDATA[6]]></EventKey>
	<SendPicsInfo>
		<Count>1</Count>
		<PicList>
			<item>
				<PicMd5Sum><![CDATA[1b5f7c23b5bf75682a53e7b6d163e185]]></PicMd5Sum>
			</item>
		</PicList>
	</SendPicsInfo>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "gh_e136c6e50636",
				FromUserName: "oMgHVjngRipVsoxg6TuX3vz6glDg",
				CreateTime:   1408090651,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_PIC_SYSPHOTO,
			EventKey: "6",
		}
		expectReq.SendPicsInfo.Count = 1
		expectReq.SendPicsInfo.PicList = make([]struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		}, 1)
		expectReq.SendPicsInfo.PicList[0].PicMd5Sum = "1b5f7c23b5bf75682a53e7b6d163e185"

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// pic_photo_or_album：弹出拍照或者相册发图的事件推送============================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
		<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
		<CreateTime>1408090816</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[pic_photo_or_album]]></Event>
		<EventKey><![CDATA[6]]></EventKey>
		<SendPicsInfo>
			<Count>1</Count>
			<PicList>
				<item>
					<PicMd5Sum><![CDATA[5a75aaca956d97be686719218f275c6b]]></PicMd5Sum>
				</item>
			</PicList>
		</SendPicsInfo>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "gh_e136c6e50636",
				FromUserName: "oMgHVjngRipVsoxg6TuX3vz6glDg",
				CreateTime:   1408090816,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_PIC_PHOTO_OR_ALBUM,
			EventKey: "6",
		}
		expectReq.SendPicsInfo.Count = 1
		expectReq.SendPicsInfo.PicList = make([]struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		}, 1)
		expectReq.SendPicsInfo.PicList[0].PicMd5Sum = "5a75aaca956d97be686719218f275c6b"

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// pic_weixin：弹出微信相册发图器的事件推送=====================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
		<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
		<CreateTime>1408090816</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[pic_weixin]]></Event>
		<EventKey><![CDATA[6]]></EventKey>
		<SendPicsInfo>
			<Count>1</Count>
			<PicList>
				<item>
					<PicMd5Sum><![CDATA[5a75aaca956d97be686719218f275c6b]]></PicMd5Sum>
				</item>
			</PicList>
		</SendPicsInfo>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "gh_e136c6e50636",
				FromUserName: "oMgHVjngRipVsoxg6TuX3vz6glDg",
				CreateTime:   1408090816,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_PIC_WEIXIN,
			EventKey: "6",
		}
		expectReq.SendPicsInfo.Count = 1
		expectReq.SendPicsInfo.PicList = make([]struct {
			PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
		}, 1)
		expectReq.SendPicsInfo.PicList[0].PicMd5Sum = "5a75aaca956d97be686719218f275c6b"

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// location_select：弹出地理位置选择器的事件推送================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
		<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
		<CreateTime>1408091189</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[location_select]]></Event>
		<EventKey><![CDATA[6]]></EventKey>
		<SendLocationInfo>
			<Location_X><![CDATA[23]]></Location_X>
			<Location_Y><![CDATA[113]]></Location_Y>
			<Scale><![CDATA[15]]></Scale>
			<Label><![CDATA[ 广州市海珠区客村艺苑路 106号]]></Label>
			<Poiname><![CDATA[]]></Poiname>
		</SendLocationInfo>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "gh_e136c6e50636",
				FromUserName: "oMgHVjngRipVsoxg6TuX3vz6glDg",
				CreateTime:   1408091189,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:    EVENT_TYPE_LOCATION_SELECT,
			EventKey: "6",
		}
		expectReq.SendLocationInfo.LocationX = 23
		expectReq.SendLocationInfo.LocationY = 113
		expectReq.SendLocationInfo.Scale = 15
		expectReq.SendLocationInfo.Label = " 广州市海珠区客村艺苑路 106号"
		expectReq.SendLocationInfo.Poiname = ""

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 高级群发消息，事件推送群发结果=================================================

	msgBytes = []byte(`<xml>
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
		<ErrorCount>5</ErrorCount>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "gh_3e8adccde292",
				FromUserName: "oR5Gjjl_eiZoUpGozMo7dbBJ362A",
				CreateTime:   1394524295,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:       EVENT_TYPE_MASSSENDJOBFINISH,
			MsgID:       1988,
			Status:      "sendsuccess",
			TotalCount:  100,
			FilterCount: 80,
			SentCount:   75,
			ErrorCount:  5,
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 模板消息发送事件推送结果======================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[gh_7f083739789a]]></ToUserName>
		<FromUserName><![CDATA[oia2TjuEGTNoeX76QEjQNrcURxG8]]></FromUserName>
		<CreateTime>1395658920</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[TEMPLATESENDJOBFINISH]]></Event>
		<MsgID>200163836</MsgID>
		<Status><![CDATA[success]]></Status>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "gh_7f083739789a",
				FromUserName: "oia2TjuEGTNoeX76QEjQNrcURxG8",
				CreateTime:   1395658920,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:  EVENT_TYPE_TEMPLATESENDJOBFINISH,
			MsgID:  200163836,
			Status: "success",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}

	// 微信小店, 订单付款通知=======================================================

	msgBytes = []byte(`<xml>
		<ToUserName><![CDATA[weixin_media1]]></ToUserName>
		<FromUserName><![CDATA[oDF3iYyVlek46AyTBbMRVV8VZVlI]]></FromUserName>
		<CreateTime>1398144192</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[merchant_order]]></Event>
		<OrderId><![CDATA[test_order_id]]></OrderId>
		<OrderStatus>2</OrderStatus>
		<ProductId><![CDATA[test_product_id]]></ProductId>
		<SkuInfo><![CDATA[10001:1000012;10002:100021]]></SkuInfo>
	</xml>`)

	req.Zero()
	if err := xml.Unmarshal(msgBytes, &req); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", msgBytes, err)
	} else {
		expectReq := Request{
			CommonHead: CommonHead{
				ToUserName:   "weixin_media1",
				FromUserName: "oDF3iYyVlek46AyTBbMRVV8VZVlI",
				CreateTime:   1398144192,
				MsgType:      MSG_TYPE_EVENT,
			},

			Event:       EVENT_TYPE_MERCHANTORDER,
			OrderId:     "test_order_id",
			OrderStatus: 2,
			ProductId:   "test_product_id",
			SkuInfo:     "10001:1000012;10002:100021",
		}

		if !compareRequest(&req, &expectReq) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		}
	}
}

// 比较两个 Request 是否相等, 相等返回 true, 否则返回 false
func compareRequest(req1, req2 *Request) bool {
	if req1.XMLName != req2.XMLName {
		return false
	}
	if req1.CommonHead != req2.CommonHead {
		return false
	}
	if req1.MsgId != req2.MsgId {
		return false
	}
	if req1.MsgID != req2.MsgID {
		return false
	}
	if req1.Content != req2.Content {
		return false
	}
	if req1.MediaId != req2.MediaId {
		return false
	}
	if req1.PicURL != req2.PicURL {
		return false
	}
	if req1.Format != req2.Format {
		return false
	}
	if req1.Recognition != req2.Recognition {
		return false
	}
	if req1.ThumbMediaId != req2.ThumbMediaId {
		return false
	}
	if req1.LocationX != req2.LocationX {
		return false
	}
	if req1.LocationY != req2.LocationY {
		return false
	}
	if req1.Scale != req2.Scale {
		return false
	}
	if req1.Label != req2.Label {
		return false
	}
	if req1.Title != req2.Title {
		return false
	}
	if req1.Description != req2.Description {
		return false
	}
	if req1.URL != req2.URL {
		return false
	}
	if req1.Event != req2.Event {
		return false
	}
	if req1.EventKey != req2.EventKey {
		return false
	}
	if req1.ScanCodeInfo != req2.ScanCodeInfo {
		return false
	}

	if false == func() bool {
		if req1.SendPicsInfo.Count != req2.SendPicsInfo.Count {
			return false
		}
		if len(req1.SendPicsInfo.PicList) != len(req2.SendPicsInfo.PicList) {
			return false
		}
		for i := 0; i < len(req1.SendPicsInfo.PicList); i++ {
			if req1.SendPicsInfo.PicList[i] != req2.SendPicsInfo.PicList[i] {
				return false
			}
		}
		return true
	}() {
		return false
	}

	if req1.SendLocationInfo != req2.SendLocationInfo {
		return false
	}
	if req1.Ticket != req2.Ticket {
		return false
	}
	if req1.Latitude != req2.Latitude {
		return false
	}
	if req1.Longitude != req2.Longitude {
		return false
	}
	if req1.Precision != req2.Precision {
		return false
	}
	if req1.Status != req2.Status {
		return false
	}
	if req1.TotalCount != req2.TotalCount {
		return false
	}
	if req1.FilterCount != req2.FilterCount {
		return false
	}
	if req1.SentCount != req2.SentCount {
		return false
	}
	if req1.ErrorCount != req2.ErrorCount {
		return false
	}
	if req1.OrderId != req2.OrderId {
		return false
	}
	if req1.OrderStatus != req2.OrderStatus {
		return false
	}
	if req1.ProductId != req2.ProductId {
		return false
	}
	if req1.SkuInfo != req2.SkuInfo {
		return false
	}
	return true
}
