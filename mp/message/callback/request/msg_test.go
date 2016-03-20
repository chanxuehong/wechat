package request

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/chanxuehong/wechat/mp/core"
)

func TestTextMessage(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457742803</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[hhhhhjjjjjjjjj]]></Content>
<MsgId>6260957665266699303</MsgId>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetText(mixedMsg)

	var wantObject = &Text{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457742803,
			MsgType:      MsgTypeText,
		},
		MsgId:   6260957665266699303,
		Content: "hhhhhjjjjjjjjj",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}

func TestImageMessage(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457742821</CreateTime>
<MsgType><![CDATA[image]]></MsgType>
<PicUrl><![CDATA[http://mmbiz.qpic.cn/mmbiz/eUIjD3Aun1G3tmxfVLgNsNI5UVrf5fFZXC9bkNmibssibA4aluCxmqp6ibyyxT4iaXmQaMR72hncyp65Yt2ZRIsH9g/0]]></PicUrl>
<MsgId>6260957742576110634</MsgId>
<MediaId><![CDATA[XbQZ2aq1i7G6ZPZzP8DoySg25X5iDcjefx_cbO3kFzww5I9Y_NgNVSsDqte1ZNOR]]></MediaId>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetImage(mixedMsg)

	var wantObject = &Image{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457742821,
			MsgType:      MsgTypeImage,
		},
		MsgId:   6260957742576110634,
		MediaId: "XbQZ2aq1i7G6ZPZzP8DoySg25X5iDcjefx_cbO3kFzww5I9Y_NgNVSsDqte1ZNOR",
		PicURL:  "http://mmbiz.qpic.cn/mmbiz/eUIjD3Aun1G3tmxfVLgNsNI5UVrf5fFZXC9bkNmibssibA4aluCxmqp6ibyyxT4iaXmQaMR72hncyp65Yt2ZRIsH9g/0",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}

func TestVoiceMessage(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457742830</CreateTime>
<MsgType><![CDATA[voice]]></MsgType>
<MediaId><![CDATA[NUjTR19l2n7kxNmmol1K3unuUrk-b2LRMri_SkEoFADjnVYeV8tFjx8EQB7trHJk]]></MediaId>
<Format><![CDATA[amr]]></Format>
<MsgId>6260957780828487680</MsgId>
<Recognition><![CDATA[傻逼！]]></Recognition>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetVoice(mixedMsg)

	var wantObject = &Voice{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457742830,
			MsgType:      MsgTypeVoice,
		},
		MsgId:       6260957780828487680,
		MediaId:     "NUjTR19l2n7kxNmmol1K3unuUrk-b2LRMri_SkEoFADjnVYeV8tFjx8EQB7trHJk",
		Format:      "amr",
		Recognition: "傻逼！",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}

func TestVideoMessage(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>1357290913</CreateTime>
<MsgType><![CDATA[video]]></MsgType>
<MediaId><![CDATA[media_id]]></MediaId>
<ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
<MsgId>1234567890123456</MsgId>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetVideo(mixedMsg)

	var wantObject = &Video{
		MsgHeader: core.MsgHeader{
			ToUserName:   "toUser",
			FromUserName: "fromUser",
			CreateTime:   1357290913,
			MsgType:      MsgTypeVideo,
		},
		MsgId:        1234567890123456,
		MediaId:      "media_id",
		ThumbMediaId: "thumb_media_id",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}

func TestShortVideoMessage(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457742856</CreateTime>
<MsgType><![CDATA[shortvideo]]></MsgType>
<MediaId><![CDATA[hcik2O_g6Pa6d2RjvZze-BgLjocAiBfsNqXd_JBkw570RUb0qh8ayO12TaC5TPhc]]></MediaId>
<ThumbMediaId><![CDATA[5Sykx9I2TiVMmU0EAYTiSVZ7rdU0q7zHw8NEBL_COo_hedMcd9RFuz0-zv_Hqplk]]></ThumbMediaId>
<MsgId>6260957892899965999</MsgId>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetShortVideo(mixedMsg)

	var wantObject = &ShortVideo{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457742856,
			MsgType:      MsgTypeShortVideo,
		},
		MsgId:        6260957892899965999,
		MediaId:      "hcik2O_g6Pa6d2RjvZze-BgLjocAiBfsNqXd_JBkw570RUb0qh8ayO12TaC5TPhc",
		ThumbMediaId: "5Sykx9I2TiVMmU0EAYTiSVZ7rdU0q7zHw8NEBL_COo_hedMcd9RFuz0-zv_Hqplk",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}

func TestLocationMessage(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457742894</CreateTime>
<MsgType><![CDATA[location]]></MsgType>
<Location_X>30.195360</Location_X>
<Location_Y>109.32305</Location_Y>
<Scale>10</Scale>
<Label><![CDATA[湖北省恩施土家族苗族自治州恩施市]]></Label>
<MsgId>6260958056108723257</MsgId>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetLocation(mixedMsg)

	var wantObject = &Location{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457742894,
			MsgType:      MsgTypeLocation,
		},
		MsgId:     6260958056108723257,
		LocationX: 30.195360,
		LocationY: 109.32305,
		Scale:     10,
		Label:     "湖北省恩施土家族苗族自治州恩施市",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}

func TestLinkMessage(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457742914</CreateTime>
<MsgType><![CDATA[link]]></MsgType>
<Title><![CDATA[当日本的“变态”遇上德国的严谨，结果房子疯了！]]></Title>
<Description><![CDATA[http://mp.weixin.qq.com/s?__biz=MjM5OTI1NzMwMQ==&mid=401506474&idx=2&sn=2a58169e0c6cf9a1d0b364d3415ef9f2&scene=2&srcid=1213NSFj35aDx3AQxygUNmJK#rd]]></Description>
<Url><![CDATA[http://mp.weixin.qq.com/s?__biz=MjM5OTI1NzMwMQ==&mid=401506474&idx=2&sn=2a58169e0c6cf9a1d0b364d3415ef9f2&scene=2&srcid=1213NSFj35aDx3AQxygUNmJK&from=timeline&isappinstalled=0#rd]]></Url>
<MsgId>6260958142008069179</MsgId>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetLink(mixedMsg)

	var wantObject = &Link{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457742914,
			MsgType:      MsgTypeLink,
		},
		MsgId:       6260958142008069179,
		Title:       "当日本的“变态”遇上德国的严谨，结果房子疯了！",
		Description: "http://mp.weixin.qq.com/s?__biz=MjM5OTI1NzMwMQ==&mid=401506474&idx=2&sn=2a58169e0c6cf9a1d0b364d3415ef9f2&scene=2&srcid=1213NSFj35aDx3AQxygUNmJK#rd",
		URL:         "http://mp.weixin.qq.com/s?__biz=MjM5OTI1NzMwMQ==&mid=401506474&idx=2&sn=2a58169e0c6cf9a1d0b364d3415ef9f2&scene=2&srcid=1213NSFj35aDx3AQxygUNmJK&from=timeline&isappinstalled=0#rd",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}
