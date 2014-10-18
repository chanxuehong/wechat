// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package shelf

type Shelf struct {
	Id   int64  `json:"shelf_id,omitempty"`
	Name string `json:"shelf_name"`

	// 货架招牌图片URL(图片需调用图片上传接口获得图片URL填写至此，否则添加货架失败，
	// 建议尺寸为640*120，仅控件1-4有banner，控件5没有banner)
	Banner string `json:"shelf_banner,omitempty"`

	Info struct {
		ModuleInfos []Module `json:"module_infos,omitempty"`
	} `json:"shelf_info"`
}
