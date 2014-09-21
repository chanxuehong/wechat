// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

const (
	MenuButtonCountLimit    = 3 // 一级菜单最多包含 3 个按钮
	SubMenuButtonCountLimit = 5 // 二级菜单最多包含 5 个按钮
)

const (
	MenuButtonLenLimit    = 16 // 菜单标题不超过16个字节
	SubMenuButtonLenLimit = 40 // 子菜单标题不超过40个字节
)

const (
	ButtonKeyLenLimit = 128 // 菜单KEY值不能超过128字节
	ButtonURLLenLimit = 256 // 网页链接不能超过256字节
)

const (
	BUTTON_TYPE_CLICK              = "click"              // 点击推事件
	BUTTON_TYPE_VIEW               = "view"               // 跳转URL
	BUTTON_TYPE_SCANCODE_WAITMSG   = "scancode_waitmsg"   // 扫码带提示
	BUTTON_TYPE_SCANCODE_PUSH      = "scancode_push"      // 扫码推事件
	BUTTON_TYPE_PIC_SYSPHOTO       = "pic_sysphoto"       // 系统拍照发图
	BUTTON_TYPE_PIC_PHOTO_OR_ALBUM = "pic_photo_or_album" // 拍照或者相册发图
	BUTTON_TYPE_PIC_WEIXIN         = "pic_weixin"         // 微信相册发图
	BUTTON_TYPE_LOCATION_SELECT    = "location_select"    // 发送位置
)
