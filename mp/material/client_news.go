// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package material

import (
	"errors"
	"fmt"

	"github.com/chanxuehong/wechat/mp"
)

const (
	NewsArticleCountLimit = 10 // 图文消息里文章的个数限制
)

type Article struct {
	ThumbMediaId     string `json:"thumb_media_id"`               // 必须; 图文消息的封面图片素材id（必须是永久mediaID）
	Title            string `json:"title"`                        // 必须; 标题
	Author           string `json:"author,omitempty"`             // 必须; 作者
	Digest           string `json:"digest,omitempty"`             // 必须; 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
	Content          string `json:"content"`                      // 必须; 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS
	ContentSourceURL string `json:"content_source_url,omitempty"` // 必须; 图文消息的原文地址，即点击“阅读原文”后的URL
	ShowCoverPic     int    `json:"show_cover_pic"`               // 必须; 是否显示封面，0为false，即不显示，1为true，即显示
}

func (article *Article) SetShowCoverPic(b bool) {
	if b {
		article.ShowCoverPic = 1
	} else {
		article.ShowCoverPic = 0
	}
}

type News []Article

// 新增永久图文素材.
func (clt *Client) CreateNews(news News) (mediaId string, err error) {
	if news == nil {
		err = errors.New("nil news")
		return
	}
	if len(news) == 0 {
		err = errors.New("图文消息是空的")
		return
	}
	if len(news) > NewsArticleCountLimit {
		err = fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", NewsArticleCountLimit, len(news))
		return
	}

	var request = struct {
		Articles []Article `json:"articles,omitempty"`
	}{
		Articles: news,
	}

	var result struct {
		mp.Error
		MediaId string `json:"media_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	mediaId = result.MediaId
	return
}

// 获取图文消息素材
func (clt *Client) GetNews(mediaId string) (news News, err error) {
	request := struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}

	var result struct {
		mp.Error
		NewsItem News `json:"news_item"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	news = result.NewsItem
	return
}
