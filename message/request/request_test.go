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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := Text{
				CommonHead: CommonHead{
					ToUserName:   "toUser",
					FromUserName: "fromUser",
					CreateTime:   1348831860,
					MsgType:      MSG_TYPE_TEXT,
				},

				MsgId:   1234567890123456,
				Content: "this is a test",
			}

			text := req.Text()
			if *text != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, text, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := Image{
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

			image := req.Image()
			if *image != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, image, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := Voice{
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

			voiceRecognition := req.Voice()
			if *voiceRecognition != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, voiceRecognition, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := Video{
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

			video := req.Video()
			if *video != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, video, expect)
			}
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

			MsgId:      1234567890123456,
			Location_X: 23.134525,  // 最后一位是 5 才能精确表示
			Location_Y: 113.358805, // 最后一位是 5 才能精确表示
			Scale:      20,
			Label:      "位置信息",
		}
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := Location{
				CommonHead: CommonHead{
					ToUserName:   "toUser",
					FromUserName: "fromUser",
					CreateTime:   1351776360,
					MsgType:      MSG_TYPE_LOCATION,
				},

				MsgId:      1234567890123456,
				Location_X: 23.134525,  // 最后一位是 5 才能精确表示
				Location_Y: 113.358805, // 最后一位是 5 才能精确表示
				Scale:      20,
				Label:      "位置信息",
			}

			location := req.Location()
			if *location != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, location, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := Link{
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

			link := req.Link()
			if *link != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, link, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := SubscribeEvent{
				CommonHead: CommonHead{
					ToUserName:   "toUser",
					FromUserName: "FromUser",
					CreateTime:   123456789,
					MsgType:      MSG_TYPE_EVENT,
				},

				Event: EVENT_TYPE_SUBSCRIBE,
			}

			SubscribeEvent := req.SubscribeEvent()
			if *SubscribeEvent != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, SubscribeEvent, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := UnsubscribeEvent{
				CommonHead: CommonHead{
					ToUserName:   "toUser",
					FromUserName: "FromUser",
					CreateTime:   123456789,
					MsgType:      MSG_TYPE_EVENT,
				},

				Event: EVENT_TYPE_UNSUBSCRIBE,
			}

			UnsubscribeEvent := req.UnsubscribeEvent()
			if *UnsubscribeEvent != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, UnsubscribeEvent, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := SubscribeByScanEvent{
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

			SubscribeByScanEvent := req.SubscribeByScanEvent()
			if *SubscribeByScanEvent != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, SubscribeByScanEvent, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := ScanEvent{
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

			ScanEvent := req.ScanEvent()
			if *ScanEvent != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, ScanEvent, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := LocationEvent{
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

			LocationEvent := req.LocationEvent()
			if *LocationEvent != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, LocationEvent, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := MenuClickEvent{
				CommonHead: CommonHead{
					ToUserName:   "toUser",
					FromUserName: "FromUser",
					CreateTime:   123456789,
					MsgType:      MSG_TYPE_EVENT,
				},

				Event:    EVENT_TYPE_CLICK,
				EventKey: "EVENTKEY",
			}

			MenuClickEvent := req.MenuClickEvent()
			if *MenuClickEvent != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, MenuClickEvent, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := MenuViewEvent{
				CommonHead: CommonHead{
					ToUserName:   "toUser",
					FromUserName: "FromUser",
					CreateTime:   123456789,
					MsgType:      MSG_TYPE_EVENT,
				},

				Event:    EVENT_TYPE_VIEW,
				EventKey: "www.qq.com",
			}

			MenuViewEvent := req.MenuViewEvent()
			if *MenuViewEvent != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, MenuViewEvent, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := MassSendJobFinishEvent{
				CommonHead: CommonHead{
					ToUserName:   "gh_3e8adccde292",
					FromUserName: "oR5Gjjl_eiZoUpGozMo7dbBJ362A",
					CreateTime:   1394524295,
					MsgType:      MSG_TYPE_EVENT,
				},

				Event:       EVENT_TYPE_MASSSENDJOBFINISH,
				MsgId:       1988,
				Status:      "sendsuccess",
				TotalCount:  100,
				FilterCount: 80,
				SentCount:   75,
				ErrorCount:  5,
			}

			MassSendJobFinishEvent := req.MassSendJobFinishEvent()
			if *MassSendJobFinishEvent != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, MassSendJobFinishEvent, expect)
			}
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
		if req != expectReq {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, req, expectReq)
		} else {
			expect := MerchantOrderEvent{
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

			MerchantOrderEvent := req.MerchantOrderEvent()
			if *MerchantOrderEvent != expect {
				t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", msgBytes, MerchantOrderEvent, expect)
			}
		}
	}
}
