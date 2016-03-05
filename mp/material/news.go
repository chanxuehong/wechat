package material

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type Article struct {
	ThumbMediaId     string `json:"thumb_media_id"`               // 图文消息的封面图片素材id(必须是永久mediaID)
	Title            string `json:"title"`                        // 标题
	Author           string `json:"author,omitempty"`             // 作者
	Digest           string `json:"digest,omitempty"`             // 图文消息的摘要, 仅有单图文消息才有摘要, 多图文此处为空
	Content          string `json:"content"`                      // 图文消息的具体内容, 支持HTML标签, 必须少于2万字符, 小于1M, 且此处会去除JS
	ContentSourceURL string `json:"content_source_url,omitempty"` // 图文消息的原文地址, 即点击"阅读原文"后的URL
	ShowCoverPic     int    `json:"show_cover_pic"`               // 是否显示封面, 0为false, 即不显示, 1为true, 即显示
	URL              string `json:"url,omitempty"`                // !!!创建时不需要此参数!!! 图文页的URL, 文章创建成功以后, 会由微信自动生成
}

type News struct {
	Articles []Article `json:"articles,omitempty"`
}

// 新增永久图文素材.
func AddNews(clt *core.Client, news *News) (mediaId string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token="

	var result struct {
		core.Error
		MediaId string `json:"media_id"`
	}
	if err = clt.PostJSON(incompleteURL, news, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	mediaId = result.MediaId
	return
}

// 修改永久图文素材.
func UpdateNews(clt *core.Client, mediaId string, index int, article *Article) (err error) {
	var request = struct {
		MediaId string   `json:"media_id"`
		Index   int      `json:"index"`
		Article *Article `json:"articles,omitempty"`
	}{
		MediaId: mediaId,
		Index:   index,
		Article: article,
	}

	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/update_news?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 获取永久图文素材.
func GetNews(clt *core.Client, mediaId string) (news *News, err error) {
	var request = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}

	var result struct {
		core.Error
		Articles []Article `json:"news_item"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	news = &News{
		Articles: result.Articles,
	}
	return
}

type NewsInfo struct {
	MediaId string `json:"media_id"` // 素材id
	Content struct {
		Articles []Article `json:"news_item,omitempty"`
	} `json:"content"`
	UpdateTime int64 `json:"update_time"` // 最后更新时间
}

type BatchGetNewsResult struct {
	TotalCount int        `json:"total_count"` // 该类型的素材的总数
	ItemCount  int        `json:"item_count"`  // 本次调用获取的素材的数量
	Items      []NewsInfo `json:"item"`        // 本次调用获取的素材列表
}

// 获取图文素材列表.
//
//  offset:       从全部素材的该偏移位置开始返回, 0表示从第一个素材 返回
//  count:        返回素材的数量, 取值在1到20之间
func BatchGetNews(clt *core.Client, offset, count int) (rslt *BatchGetNewsResult, err error) {
	var request = struct {
		MaterialType string `json:"type"`
		Offset       int    `json:"offset"`
		Count        int    `json:"count"`
	}{
		MaterialType: MaterialTypeNews,
		Offset:       offset,
		Count:        count,
	}

	var result struct {
		core.Error
		BatchGetNewsResult
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}

	rslt = &result.BatchGetNewsResult
	return
}

// NewsIterator
//
//  iter, err := Client.NewsIterator(0, 10)
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
type NewsIterator struct {
	clt *core.Client // 关联的微信 Client

	nextOffset int // 下一次获取数据时的 offset
	count      int // 步长

	lastBatchGetNewsResult *BatchGetNewsResult // 最近一次获取的数据
	nextPageHasCalled      bool                // NextPage() 是否调用过
}

func (iter *NewsIterator) TotalCount() int {
	return iter.lastBatchGetNewsResult.TotalCount
}

func (iter *NewsIterator) HasNext() bool {
	if !iter.nextPageHasCalled { // 第一次调用需要特殊对待
		return iter.lastBatchGetNewsResult.ItemCount > 0 ||
			iter.nextOffset < iter.lastBatchGetNewsResult.TotalCount
	}

	return iter.nextOffset < iter.lastBatchGetNewsResult.TotalCount
}

func (iter *NewsIterator) NextPage() (items []NewsInfo, err error) {
	if !iter.nextPageHasCalled { // 第一次调用需要特殊对待
		iter.nextPageHasCalled = true

		items = iter.lastBatchGetNewsResult.Items
		return
	}

	rslt, err := BatchGetNews(iter.clt, iter.nextOffset, iter.count)
	if err != nil {
		return
	}

	iter.nextOffset += rslt.ItemCount
	iter.lastBatchGetNewsResult = rslt

	items = rslt.Items
	return
}

func CreateNewsIterator(clt *core.Client, offset, count int) (iter *NewsIterator, err error) {
	// 逻辑上相当于第一次调用 NewsIterator.NextPage, 因为第一次调用 NewsIterator.HasNext 需要数据支撑, 所以提前获取了数据

	rslt, err := BatchGetNews(clt, offset, count)
	if err != nil {
		return
	}

	iter = &NewsIterator{
		clt: clt,

		nextOffset: offset + rslt.ItemCount,
		count:      count,

		lastBatchGetNewsResult: rslt,
		nextPageHasCalled:      false,
	}
	return
}
