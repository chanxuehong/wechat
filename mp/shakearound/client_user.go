// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)
package shakearound

import (
	"github.com/chanxuehong/wechat/mp"
)

// 新增页面
func (clt *Client) UserGetShakeInfo(ticket string, needPoi int) (pageId int64, beaconInfo ShakeBeaconInfo, err error) {
	var request = struct {
		Ticket  string `json:"ticket"`
		NeedPoi int    `json:"need_poi,omtiempty"`
	}{
		Ticket:  ticket,
		NeedPoi: needPoi,
	}
	var result struct {
		mp.Error
		Data struct {
			PageId     int64           `json:"page_id"`
			BeaconInfo ShakeBeaconInfo `json:"beacon_info"`
			OpenId     string          `json:"openid"`
			PoiId      int64           `json:"poi_id,omtiempty"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/user/getshakeinfo?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	pageId = result.Data.PageId
	beaconInfo = result.Data.BeaconInfo

	return
}
