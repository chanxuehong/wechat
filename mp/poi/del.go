package poi

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 删除门店.
func PoiDelete(clt *core.Client, poiId int64) (err error) {
	var request = struct {
		PoiId int64 `json:"poi_id,string"`
	}{
		PoiId: poiId,
	}

	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
