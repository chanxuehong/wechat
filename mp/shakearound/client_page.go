package shakearound

import (
	"github.com/chanxuehong/wechat/mp"
)

type Page struct {
	PageId int `json:"page_id,omitempty"`
	Title string `json:"title"`
	Description string `json:"description"`
	IconUrl string `json:"icon_url"`
	PageUrl string `json:"page_url"`
	Comment string `json:"comment"`
}


func (clt Client) AddPage(page *Page) (err error) {
	var result struct {
		mp.Error
		Data struct {
				 PageId int `json:"page_id"`
			 } `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/add?access_token="
	if err = clt.PostJSON(incompleteURL, page, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	page.PageId = result.Data.PageId
	return
}

func (clt Client) UpdatePage(page *Page) (err error) {
	var result struct {
		mp.Error
		Data struct {
				 PageId int `json:"page_id"`
			 } `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/update?access_token="
	if err = clt.PostJSON(incompleteURL, page, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

func (clt Client) SearchPageById(pageId int) (page *Page, totalCount int, err error) {
	pages, totalCount, err := clt.SearchPageByIds(&([]int{pageId}))
	page = &(*pages)[0]
	return
}

func (clt Client) SearchPageByIds(pageIds *[]int) (pages *[]Page, totalCount int, err error) {
	var request = struct {
		PageIds   *[]int `json:"page_ids"`
	}{
		PageIds:   pageIds,
	}
	return clt.SearchPage(request)
}

func (clt Client) SearchPageByCount(begin, count int) (pages *[]Page, totalCount int, err error) {
	var request = struct {
		Begin   int `json:"begin"`
		Count   int `json:"count"`
	}{
		Begin:  begin,
		Count:	count,
	}
	return clt.SearchPage(request)
}

func (clt Client) SearchPage(v interface{}) (pages *[]Page, totalCount int, err error) {
	var result struct {
		mp.Error
		Data struct {
			 Pages []Page `json:"pages"`
			 TotalCount int `json:"total_count"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/search?access_token="
	if err = clt.PostJSON(incompleteURL, v, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	pages = &result.Data.Pages
	totalCount = result.Data.TotalCount
	return
}

func (clt Client) DeletePage(pageIds *[]int) (err error) {
	var request = struct {
		PageIds *[]int `json:"page_ids"`
	}{
		PageIds: pageIds,
	}
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}