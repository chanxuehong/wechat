// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

package menu

const (
	// 目前自定义菜单最多包括3个一级菜单，每个一级菜单最多包含5个二级菜单。
	MenuButtonCountLimit    = 3
	SubMenuButtonCountLimit = 5
)

const (
	MENUBUTTON_TYPE_CLICK = "click"
	MENUBUTTON_TYPE_VIEW  = "view"
)
