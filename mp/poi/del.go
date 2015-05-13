// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package poi

import (
	"github.com/chanxuehong/wechat/mp"
)

// 删除门店.
func (clt Client) PoiDelete(poiId string) (err error) {
	var request = struct {
		PoiId string `json:"poi_id"`
	}{
		PoiId: poiId,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
