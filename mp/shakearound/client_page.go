// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)
package shakearound

// 测试通过
import (
	"github.com/chanxuehong/wechat/mp"
)

type ShakearoundPage struct {
	PageId      int64  `json:"page_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PageURL     string `json:"page_url"`
	Comment     string `json:"comment"`
	IconURL     string `json:"icon_url,omitempty"`
}

func (clt Client) PageAdd(page ShakearoundPage) (pageId int64, err error) {
	var result struct {
		mp.Error
		Data struct {
			PageId int64 `json:"page_id"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/add?access_token="
	if err = clt.PostJSON(incompleteURL, &page, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	pageId = result.Data.PageId

	return
}

func (clt Client) PageUpdate(page ShakearoundPage) (pageId int64, err error) {
	var result struct {
		mp.Error
		Data struct {
			PageId int64 `json:"page_id"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/update?access_token="
	if err = clt.PostJSON(incompleteURL, &page, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	pageId = result.Data.PageId
	return
}

// 需要查询指定页面时：
// {
//     "page_ids":[12345, 23456, 34567]
// }
// 需要分页查询或者指定范围内的页面时：
// {
//     "begin": 0,
//     "count": 3
// }
func (clt Client) PageSearch(pageIds []int64, begin, count int64) (totalCount int64, pages []ShakearoundPage, err error) {
	var request = struct {
		PageIds []int64 `json:"page_ids,omtiempty"`
		Begin   int64   `json:"begin,omitempty"` // 页面列表的起始索引值
		Count   int64   `json:"count,omitempty"` // 待查询的页面个数
	}{
		PageIds: pageIds,
		Begin:   begin,
		Count:   count,
	}
	var result struct {
		mp.Error
		Data struct {
			Pages      []ShakearoundPage `json:"pages"`
			TotalCount int64             `json:"total_count"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/search?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	totalCount = result.Data.TotalCount
	pages = result.Data.Pages
	return
}

func (clt Client) PageDelete(pageIds []int64) (err error) {
	var request = struct {
		PageIds []int64 `json:"page_ids,omitempty"`
	}{
		PageIds: pageIds,
	}
	var result struct {
		mp.Error
		Data struct {
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}
