package menu

// 菜单的按钮.
// NOTE: (MenuButton.Type, MenuButton.Key) 不能和 MenuButton.SubButton 不能同时设置
type MenuButton struct {
	Name      string        `json:"name"`
	Type      string        `json:"type,omitempty"`
	Key       string        `json:"key,omitempty"`
	SubButton []*MenuButton `json:"sub_button,omitempty"`
}

type Menu struct {
	Button []*MenuButton `json:"button"`
}

// 获取自定义菜单的回复
type MenuResponse struct {
	Menu Menu `json:"menu"`
}
