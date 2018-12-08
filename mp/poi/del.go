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
		err = &result
		return
	}
	return
}
