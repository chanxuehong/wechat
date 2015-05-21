// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com) Harry Rong(harrykobe@gmail.com)
package shakearound

import (
	"github.com/chanxuehong/wechat/mp"
	"errors"
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
	if err != nil {
		return
	}
	page = &(*pages)[0]
	return
}

func (clt Client) SearchPageByIds(pageIds *[]int) (pages *[]Page, totalCount int, err error) {
	var request = struct {
		PageIds   *[]int `json:"page_ids"`
	}{
		PageIds:   pageIds,
	}
	pages, totalCount, err = clt.SearchPage(request)
	if len(*pages) == 0{
		err = errors.New("指定ID的页面不存在")
		return
	}
	return
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

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/search?access_token="
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