// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package material

import (
	"github.com/chanxuehong/wechat/mp"
)

// 删除永久素材.
func (clt Client) DeleteMaterial(mediaId string) (err error) {
	var request = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 公众号永久素材总数信息
type MaterialCountInfo struct {
	VoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"`
}

// 获取素材总数.
func (clt Client) GetMaterialCount() (info *MaterialCountInfo, err error) {
	var result struct {
		mp.Error
		MaterialCountInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.MaterialCountInfo
	return
}

type MaterialInfo struct {
	MediaId    string `json:"media_id"`    // 素材id
	Name       string `json:"name"`        // 文件名称
	UpdateTime int64  `json:"update_time"` // 最后更新时间
	URL        string `json:"url"`         // 当获取的列表是图片素材列表时, 该字段是图片的URL
}

// 获取素材列表.
//
//  materialType: 素材的类型, 图片(image), 视频(video), 语音 (voice)
//  offset:       从全部素材的该偏移位置开始返回, 0表示从第一个素材 返回
//  count:        返回素材的数量, 取值在1到20之间
//
//  TotalCount:   该类型的素材的总数
//  ItemCount:    本次调用获取的素材的数量
//  Items:        本次调用获取的素材
func (clt Client) BatchGetMaterial(materialType string, offset, count int) (TotalCount, ItemCount int, Items []MaterialInfo, err error) {
	var request = struct {
		MaterialType string `json:"type"`
		Offset       int    `json:"offset"`
		Count        int    `json:"count"`
	}{
		MaterialType: materialType,
		Offset:       offset,
		Count:        count,
	}

	var result struct {
		mp.Error
		TotalCount int            `json:"total_count"`
		ItemCount  int            `json:"item_count"`
		Items      []MaterialInfo `json:"item"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	TotalCount = result.TotalCount
	ItemCount = result.ItemCount
	Items = result.Items
	return
}
