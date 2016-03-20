package request

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/chanxuehong/wechat/mp/core"
)

func TestUnsubscribeEvent(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457752129</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[unsubscribe]]></Event>
<EventKey><![CDATA[]]></EventKey>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetUnsubscribeEvent(mixedMsg)

	var wantObject = &UnsubscribeEvent{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457752129,
			MsgType:      "event",
		},
		EventType: EventTypeUnsubscribe,
		EventKey:  "",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}

func TestSubscribeEvent(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457752147</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[subscribe]]></Event>
<EventKey><![CDATA[]]></EventKey>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetSubscribeEvent(mixedMsg)

	var wantObject = &SubscribeEvent{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457752147,
			MsgType:      "event",
		},
		EventType: EventTypeSubscribe,
		EventKey:  "",
		Ticket:    "",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}

	msg = []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457752430</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[subscribe]]></Event>
<EventKey><![CDATA[qrscene_1000000]]></EventKey>
<Ticket><![CDATA[gQHr8ToAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xLzMwT19GN2ZsSkkxZFBGTkpkMjEzAAIEIYnjVgMECAcAAA==]]></Ticket>
</xml>`)

	mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	haveObject = GetSubscribeEvent(mixedMsg)

	wantObject = &SubscribeEvent{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457752430,
			MsgType:      "event",
		},
		EventType: EventTypeSubscribe,
		EventKey:  "qrscene_1000000",
		Ticket:    "gQHr8ToAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xLzMwT19GN2ZsSkkxZFBGTkpkMjEzAAIEIYnjVgMECAcAAA==",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}

func TestScanEvent(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457752395</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[SCAN]]></Event>
<EventKey><![CDATA[1000000]]></EventKey>
<Ticket><![CDATA[gQHr8ToAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xLzMwT19GN2ZsSkkxZFBGTkpkMjEzAAIEIYnjVgMECAcAAA==]]></Ticket>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetScanEvent(mixedMsg)

	var wantObject = &ScanEvent{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457752395,
			MsgType:      "event",
		},
		EventType: EventTypeScan,
		EventKey:  "1000000",
		Ticket:    "gQHr8ToAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xLzMwT19GN2ZsSkkxZFBGTkpkMjEzAAIEIYnjVgMECAcAAA==",
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}

func TestLocationEvent(t *testing.T) {
	msg := []byte(`<xml><ToUserName><![CDATA[gh_21ee2dc92d7d]]></ToUserName>
<FromUserName><![CDATA[os-IKuHd9pJ6xsn4mS7GyL4HxqI4]]></FromUserName>
<CreateTime>1457752531</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[LOCATION]]></Event>
<Latitude>30.443105</Latitude>
<Longitude>116.654735</Longitude>
<Precision>2703.000000</Precision>
</xml>`)

	var mixedMsg = &core.MixedMsg{}
	if err := xml.Unmarshal(msg, mixedMsg); err != nil {
		t.Errorf("unmarshal failed: %s\n", err.Error())
		return
	}
	var haveObject = GetLocationEvent(mixedMsg)

	var wantObject = &LocationEvent{
		MsgHeader: core.MsgHeader{
			ToUserName:   "gh_21ee2dc92d7d",
			FromUserName: "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
			CreateTime:   1457752531,
			MsgType:      "event",
		},
		EventType: EventTypeLocation,
		Latitude:  30.443105,
		Longitude: 116.654735,
		Precision: 2703.000000,
	}
	if !reflect.DeepEqual(haveObject, wantObject) {
		t.Errorf("compare failed,\nhave:\n%+v\nwant:\n%+v\n", haveObject, wantObject)
		return
	}
}
