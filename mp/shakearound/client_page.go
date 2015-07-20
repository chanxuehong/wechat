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
	PageId int `json:"page_id,omitempty"`		//页面的页面id
	Title string `json:"title"`					//在摇一摇页面展示的主标题
	Description string `json:"description"`		//在摇一摇页面展示的副标题
	IconUrl string `json:"icon_url"`			//在摇一摇页面展示的图片，图片需先上传至微信侧服务器
	PageUrl string `json:"page_url"`			//跳转的页面
	Comment string `json:"comment"`				//页面的备注信息
}

//	新增页面
//	page: 	新增的页面，无需设置PageId
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

//	编辑页面信息
//	page: 	需要编辑的页面，需要设置PageId
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

//	根据ID查询页面
//	pageId:		页面的ID
func (clt Client) SearchPageById(pageId int) (page *Page, totalCount int, err error) {
	pages, totalCount, err := clt.SearchPageByIds(&([]int{pageId}))
	if err != nil {
		return
	}
	if len(*pages) == 0{
		err = errors.New("指定的页面不存在")
		return
	}
	page = &(*pages)[0]
	return
}

//	根据ID查询页面列表
//	pageIds:		页面的ID列表
func (clt Client) SearchPageByIds(pageIds *[]int) (pages *[]Page, totalCount int, err error) {
	var request = struct {
		PageIds   *[]int `json:"page_ids"`
	}{
		PageIds:   pageIds,
	}
	pages, totalCount, err = clt.searchPage(request)
	return
}

//	根据分页查询或者指定范围内查询页面列表
//	begin:	页面列表的起始索引值
//	count: 	待查询的页面个数
func (clt Client) SearchPageByCount(begin, count int) (pages *[]Page, totalCount int, err error) {
	var request = struct {
		Begin   int `json:"begin"`
		Count   int `json:"count"`
	}{
		Begin:  begin,
		Count:	count,
	}
	return clt.searchPage(request)
}


func (clt Client) searchPage(v interface{}) (pages *[]Page, totalCount int, err error) {
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

//	删除页面
//	pageId:		需要删除的指定ID的页面
func (clt Client) DeletePage(pageId int) (err error) {
	return clt.DeletePages(&([]int{pageId}))
}

//	删除页面
//	pageIds:	需要删除的指定ID列表的页面
func (clt Client) DeletePages(pageIds *[]int) (err error) {
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