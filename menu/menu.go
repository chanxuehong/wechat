// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

// 菜单, 按钮个数不能超过 MenuButtonCountLimit
type Menu []*Button

// 菜单的按钮
type Button struct {
	Name    string `json:"name"`                 // 菜单标题，不超过16个字节，子菜单不超过40个字节
	Type    string `json:"type,omitempty"`       // 菜单的响应动作类型，目前有click、view两种类型
	Key     string `json:"key,omitempty"`        // click类型必须; 菜单KEY值, 用于消息接口推送, 不超过128字节
	URL     string `json:"url,omitempty"`        // view类型必须; 网页链接, 用户点击菜单可打开链接, 不超过256字节
	SubMenu Menu   `json:"sub_button,omitempty"` // 二级菜单, 按钮个数不能超过 SubMenuButtonCountLimit
}

func NewClickButton(name, key string) *Button {
	return &Button{
		Name: name,
		Type: BUTTON_TYPE_CLICK,
		Key:  key,
	}
}

func NewViewButton(name, url string) *Button {
	return &Button{
		Name: name,
		Type: BUTTON_TYPE_VIEW,
		URL:  url,
	}
}

func NewSubMenuButton(name string, submenu Menu) *Button {
	if submenu == nil {
		submenu = make([]*Button, 0, SubMenuButtonCountLimit)
	}
	return &Button{
		Name:    name,
		SubMenu: submenu,
	}
}
