// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package poi

import (
	"errors"

	"github.com/chanxuehong/wechat/mp"
)

type PoiUpdateParameters struct {
	BaseInfo struct {
		PoiId string `json:"poi_id,omitempty"`

		Telephone    string   `json:"telephone,omitempty"`    // 必须, 门店的电话（纯数字，区号、分机号均由“-”隔开）
		PhotoList    []string `json:"photo_list,omitempty"`   // 必须, 图片列表，url 形式，可以有多张图片，尺寸为640*340px。必须为上一接口生成的url
		Recommend    string   `json:"recommend,omitempty"`    // 可选, 推荐品，餐厅可为推荐菜；酒店为推荐套房；景点为推荐游玩景点等，针对自己行业的推荐内容
		Special      string   `json:"special,omitempty"`      // 必须, 特色服务，如免费wifi，免费停车，送货上门等商户能提供的特色功能或服务
		Introduction string   `json:"introduction,omitempty"` // 可选, 商户简介，主要介绍商户信息等
		OpenTime     string   `json:"open_time,omitempty"`    // 必须, 营业时间，24 小时制表示，用“-”连接，如8:00-20:00
		AvgPrice     int      `json:"avg_price,omitempty"`    // 可选, 人均价格，大于0 的整数
	} `json:"base_info"`
}

// 修改门店服务信息.
//  商户可以通过该接口，修改门店的服务信息，包括：图片列表、营业时间、推荐、特色服务、简
//  介、人均价格、电话7 个字段。目前基础字段包括（名称、坐标、地址等不可修改）
func (clt *Client) PoiUpdate(para *PoiUpdateParameters) (err error) {
	if para == nil {
		return errors.New("nil PoiUpdateParameters")
	}

	var request = struct {
		*PoiUpdateParameters `json:"business,omitempty"`
	}{
		PoiUpdateParameters: para,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/poi/updatepoi?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
