package material

import (
	"fmt"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

// 删除永久素材.
func Delete(clt *core.Client, mediaId string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token="

	var request = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
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

// 公众号永久素材总数信息
type MaterialCountInfo struct {
	VoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"`
}

// 获取素材总数数据.
func GetMaterialCount(clt *core.Client) (info *MaterialCountInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token="

	var result struct {
		core.Error
		MaterialCountInfo
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.MaterialCountInfo
	return
}

type BatchGetResult struct {
	TotalCount int            `json:"total_count"` // 该类型的素材的总数
	ItemCount  int            `json:"item_count"`  // 本次调用获取的素材的数量
	Items      []MaterialInfo `json:"item"`        // 本次调用获取的素材列表
}

type MaterialInfo struct {
	MediaId    string `json:"media_id"`    // 素材id
	Name       string `json:"name"`        // 文件名称
	UpdateTime int64  `json:"update_time"` // 最后更新时间
	URL        string `json:"url"`         // 当获取的列表是图片素材列表时, 该字段是图片的URL
}

// 获取素材列表.
//  materialType: 素材的类型, 图片(image), 视频(video), 语音 (voice)
//  offset:       从全部素材的该偏移位置开始返回, 0表示从第一个素材
//  count:        返回素材的数量, 取值在1到20之间
func BatchGet(clt *core.Client, materialType string, offset, count int) (rslt *BatchGetResult, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token="

	switch materialType {
	case MaterialTypeImage, MaterialTypeVideo, MaterialTypeVoice, MaterialTypeNews:
	default:
		err = fmt.Errorf("Incorrect materialType: %s", materialType)
		return
	}

	if offset < 0 {
		err = fmt.Errorf("Incorrect offset: %d", offset)
		return
	}
	if count <= 0 {
		err = fmt.Errorf("Incorrect count: %d", count)
		return
	}

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
		core.Error
		BatchGetResult
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	rslt = &result.BatchGetResult
	return
}

// =====================================================================================================================

// MaterialIterator
//
//  iter, err := NewMaterialIterator(clt, MaterialTypeImage, 0, 10)
//  if err != nil {
//      // TODO: 增加你的代码
//  }
//
//  for iter.HasNext() {
//      items, err := iter.NextPage()
//      if err != nil {
//          // TODO: 增加你的代码
//      }
//      // TODO: 增加你的代码
//  }
type MaterialIterator struct {
	clt *core.Client

	materialType string
	nextOffset   int
	count        int

	lastBatchGetResult *BatchGetResult
	nextPageCalled     bool
}

func (iter *MaterialIterator) TotalCount() int {
	return iter.lastBatchGetResult.TotalCount
}

func (iter *MaterialIterator) HasNext() bool {
	if !iter.nextPageCalled {
		return iter.lastBatchGetResult.ItemCount > 0 || iter.nextOffset < iter.lastBatchGetResult.TotalCount
	}
	return iter.nextOffset < iter.lastBatchGetResult.TotalCount
}

func (iter *MaterialIterator) NextPage() (items []MaterialInfo, err error) {
	if !iter.nextPageCalled {
		iter.nextPageCalled = true
		items = iter.lastBatchGetResult.Items
		return
	}

	rslt, err := BatchGet(iter.clt, iter.materialType, iter.nextOffset, iter.count)
	if err != nil {
		return
	}

	iter.lastBatchGetResult = rslt
	iter.nextOffset += rslt.ItemCount

	items = rslt.Items
	return
}

func NewMaterialIterator(clt *core.Client, materialType string, offset, count int) (iter *MaterialIterator, err error) {
	// 逻辑上相当于第一次调用 MaterialIterator.NextPage,
	// 因为第一次调用 MaterialIterator.HasNext 需要数据支撑, 所以提前获取了数据
	rslt, err := BatchGet(clt, materialType, offset, count)
	if err != nil {
		return
	}

	iter = &MaterialIterator{
		clt: clt,

		materialType: materialType,
		nextOffset:   offset + rslt.ItemCount,
		count:        count,

		lastBatchGetResult: rslt,
		nextPageCalled:     false,
	}
	return
}
