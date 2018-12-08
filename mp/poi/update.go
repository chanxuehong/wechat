package poi

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type UpdateParameters struct {
	BaseInfo struct {
		PoiId int64 `json:"poi_id"`

		// 下面7个字段，若有填写内容则为覆盖更新，若无内容则视为不修改，维持原有内容。
		// photo_list 字段为全列表覆盖，若需要增加图片，需将之前图片同样放入list 中，在其后增加新增图片。
		// 如：已有A、B、C 三张图片，又要增加D、E 两张图，则需要调用该接口，photo_list 传入A、B、C、D、E 五张图片的链接。
		Telephone    string  `json:"telephone,omitempty"`
		PhotoList    []Photo `json:"photo_list,omitempty"`
		Recommend    string  `json:"recommend,omitempty"`
		Special      string  `json:"special,omitempty"`
		Introduction string  `json:"introduction,omitempty"`
		OpenTime     string  `json:"open_time,omitempty"`
		AvgPrice     int     `json:"avg_price,omitempty"`
	} `json:"base_info"`
}

// Update 修改门店服务信息.
func Update(clt *core.Client, params *UpdateParameters) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/poi/updatepoi?access_token="

	var request = struct {
		*UpdateParameters `json:"business,omitempty"`
	}{
		UpdateParameters: params,
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
