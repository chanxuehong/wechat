// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

import (
	"testing"
)

func TestSubscribeByScanEventSceneId(t *testing.T) {
	event := SubscribeByScanEvent{
		EventKey: "qrscene_1000",
	}
	sceneid, err := event.SceneId()
	if err != nil {
		t.Error(err)
	} else if sceneid != 1000 {
		t.Errorf("SubscribeByScanEvent.SceneId():\nhave %d\nwant 1000\n", sceneid)
	}
}
