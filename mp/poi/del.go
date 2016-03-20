<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package poi

import (
	"github.com/chanxuehong/wechat/mp"
)

// 删除门店.
func (clt *Client) PoiDelete(poiId int64) (err error) {
	var request = struct {
		PoiId int64 `json:"poi_id,string"`
	}{
		PoiId: poiId,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
=======
package poi

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// Delete 删除门店.
func Delete(clt *core.Client, poiId int64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token="

	var request = struct {
		PoiId int64 `json:"poi_id"`
	}{
		PoiId: poiId,
	}
	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result
		return
	}
	return
}
