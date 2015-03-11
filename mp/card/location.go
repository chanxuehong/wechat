// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"errors"
	"fmt"

	"github.com/chanxuehong/wechat/mp"
)

type LocationAddParameters struct {
	BusinessName string  `json:"business_name"`         // 必须; 门店名称
	BranchName   string  `json:"branch_name,omitempty"` // 可选; 分店名
	Province     string  `json:"province"`              // 必须; 门店所在的省
	City         string  `json:"city"`                  // 必须; 门店所在的市
	District     string  `json:"district"`              // 必须; 门店所在的区
	Address      string  `json:"address"`               // 必须; 门店所在的详细街道地址
	Telephone    string  `json:"telephone"`             // 必须; 门店的电话
	Category     string  `json:"category"`              // 必须; 门店的类型（酒店、餐饮、购物...）
	Longitude    float64 `json:"longitude"`             // 必须; 门店所在地理位置的经度（建议使用腾讯地图定位经纬度）
	Latitude     float64 `json:"latitude"`              // 必须; 门店所在地理位置的纬度（建议使用腾讯地图定位经纬度）
}

// 批量导入门店信息.
//  1.支持商户调用该接口批量导入/新建门店信息，获取门店ID。
//    通过该接口导入的门店信息将进入门店审核流程，审核期间可正常使用。若导入的
//    门店信息未通过审核，则会被剔除出门店列表。
//  2.LocationList 和 LocationIdList 长度相等, 如果 LocationList 某个门店导入失败,
//    那么 LocationIdList 对应的位置就是等于 -1
func (clt *Client) LocationBatchAdd(LocationList []LocationAddParameters) (LocationIdList []int64, err error) {
	if len(LocationList) <= 0 {
		return
	}

	var request = struct {
		LocationList []LocationAddParameters `json:"location_list,omitempty"`
	}{
		LocationList: LocationList,
	}

	var result struct {
		mp.Error
		LocationIdList []int64 `json:"location_id_list"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/location/batchadd?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	LocationIdList = result.LocationIdList
	return
}

type Location struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// 拉取门店列表, 获取在公众平台上申请创建及API导入的门店列表，用于创建卡券.
//  offset: 偏移量，0 开始
//  count:  拉取数量
//  注：“offset”，“count”为0 时默认拉取全部门店。
func (clt *Client) LocationBatchGet(offset, count int) (LocationList []Location, err error) {
	if offset < 0 {
		err = fmt.Errorf("invalid offset: %d", offset)
		return
	}
	if count < 0 {
		err = fmt.Errorf("invalid count: %d", count)
		return
	}

	var request = struct {
		Offset int `json:"offset"`
		Count  int `json:"count"`
	}{
		Offset: offset,
		Count:  count,
	}

	var result struct {
		mp.Error
		LocationList []Location `json:"location_list"`
		Count        int        `json:"count"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/location/batchget?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	if result.Count != len(result.LocationList) {
		err = errors.New("the count and length of location_list does not match")
		return
	}
	LocationList = result.LocationList
	return
}
