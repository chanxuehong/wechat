package menu

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 获取自定义菜单配置接口.
func GetMenuInfo(clt *core.Client) (info MenuInfo, isMenuOpen bool, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token="

	var result struct {
		core.Error
		IsMenuOpen int      `json:"is_menu_open"`
		MenuInfo   MenuInfo `json:"selfmenu_info"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = result.MenuInfo
	if result.IsMenuOpen != 0 {
		isMenuOpen = true
	}
	return
}

type MenuInfo struct {
	Buttons []ButtonEx `json:"button,omitempty"`
}

type ButtonEx struct {
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	Key     string `json:"key,omitempty"`
	URL     string `json:"url,omitempty"`
	MediaId string `json:"media_id,omitempty"`

	Value    string `json:"value,omitempty"`
	NewsInfo struct {
		Articles []Article `json:"list,omitempty"`
	} `json:"news_info"`

	SubButton struct {
		Buttons []ButtonEx `json:"list,omitempty"`
	} `json:"sub_button"`
}

type Article struct {
	Title      string `json:"title,omitempty"`       // 图文消息的标题
	Author     string `json:"author,omitempty"`      // 作者
	Digest     string `json:"digest,omitempty"`      // 摘要
	ShowCover  int    `json:"show_cover"`            // 是否显示封面, 0为不显示, 1为显示
	CoverURL   string `json:"cover_url,omitempty"`   // 封面图片的URL
	ContentURL string `json:"content_url,omitempty"` // 正文的URL
	SourceURL  string `json:"source_url,omitempty"`  // 原文的URL, 若置空则无查看原文入口
}
