package menu

// 菜单的按钮.
type MenuButton struct {
	Name string `json:"name"` // 菜单标题，不超过16个字节，子菜单不超过40个字节
	// NOTE: (MenuButton.Type, MenuButton.Key, MenuButton.URL) 不能和 MenuButton.SubButton 不能同时设置
	Type string `json:"type,omitempty"` // 菜单的响应动作类型，目前有click、view两种类型
	Key  string `json:"key,omitempty"`  // click类型必须; 菜单KEY值, 用于消息接口推送, 不超过128字节
	URL  string `json:"url,omitempty"`  // view类型必须; 网页链接, 用户点击菜单可打开链接, 不超过256字节
	// 二级菜单数组, 个数不能超过 SubMenuButtonCountLimit
	SubButton []*MenuButton `json:"sub_button,omitempty"`
}

// 如果总的子按钮数超过限制, 则截除多余的.
func (mbtn *MenuButton) AppendButton(btn ...*MenuButton) {
	if len(btn) <= 0 {
		return
	}

	switch n := SubMenuButtonCountLimit - len(mbtn.SubButton); {
	case n > 0:
		if len(btn) > n {
			btn = btn[:n]
		}
		mbtn.SubButton = append(mbtn.SubButton, btn...)
	case n == 0:
	default: // n < 0
		mbtn.SubButton = mbtn.SubButton[:SubMenuButtonCountLimit]
	}
}

type Menu struct {
	// Button 的个数不能超过 MenuButtonCountLimit
	Button []*MenuButton `json:"button"`
}

// 如果总的按钮数超过限制, 则截除多余的.
func (m *Menu) AppendButton(btn ...*MenuButton) {
	if len(btn) <= 0 {
		return
	}

	switch n := MenuButtonCountLimit - len(m.Button); {
	case n > 0:
		if len(btn) > n {
			btn = btn[:n]
		}
		m.Button = append(m.Button, btn...)
	case n == 0:
	default: // n < 0
		m.Button = m.Button[:MenuButtonCountLimit]
	}
}
