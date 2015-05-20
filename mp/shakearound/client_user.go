package shakearound

import (
	"github.com/chanxuehong/wechat/mp"
)

type BeaconInfo struct {
	Distance float64 `json:"distance"`
	Uuid string `json:"uuid"`
	Major int `json:"major"`
	Minor int `json:"minor"`
}

type Shakeinfo struct {
	PageId int `json:"page_id"`
	BeaconInfo BeaconInfo `json:"beacon_info"`
	Openid string `json:"openid"`
	PoiId int `json:"poi_id"`
}

func (clt Client) Getshakeinfo(ticket string, needPoi ...bool) (shakeinfoerr *Shakeinfo, err error) {
	var need_poi int = 0
	if len(needPoi) > 0{
		if needPoi[0] == true {
			need_poi = 1
		}
	}
	var request = struct {
		Ticket string `json:"ticket"`
		Need_poi int `json:"need_poi"`
	}{
		Ticket: ticket,
		Need_poi: need_poi,
	}

	var result struct {
		mp.Error
		Data Shakeinfo `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/user/getshakeinfo?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	shakeinfoerr = &result.Data
	return
}