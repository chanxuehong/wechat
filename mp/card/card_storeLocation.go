// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com)

package card

// 门店
type StoreLocation struct {
	// 门店名称
	BusinessName string `json:"business_name"`
	// 门店所在省
	Province string `json:"province"`
	// 城市
	City string `json:"city"`
	// 区
	District string `json:"district"`
	// 街道详细地址
	Address string `json:"address"`
	// 门店电话
	Telephone string `json:"telephone"`
	// 门店的类型（酒店、餐饮、购物...）
	Category string `json:"category"`
	// 门店所在地理位置的经度（建议使用腾讯地图定位经纬度）
	Longitude string `json:"longitude"`
	// 门店所在地理位置的纬度
	Latitude string `json:"latitude"`
}

// 门店地址集合
type LocationList struct {
	List []StoreLocation `json:"location_list"`
}

type LocationResponse struct {
	LocationId int     `json:"location_id"`
	Name       string  `json:"name"`
	Phone      string  `json:"phone"`
	Address    string  `json:"address"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
}
