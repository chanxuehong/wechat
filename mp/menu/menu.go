package menu

const (
	// 下面6个类型(包括view类型)的按钮是在公众平台官网发布的菜单按钮类型
	ButtonTypeText  = "text"
	ButtonTypeImage = "img"
	ButtonTypePhoto = "photo"
	ButtonTypeVideo = "video"
	ButtonTypeVoice = "voice"

	// 上面5个类型的按钮不能通过API设置

	ButtonTypeView        = "view"        // 跳转URL
	ButtonTypeClick       = "click"       // 点击推事件
	ButtonTypeMiniProgram = "miniprogram" // 小程序

	// 下面的按钮类型仅支持微信 iPhone5.4.1 以上版本, 和 Android5.4 以上版本的微信用户,
	// 旧版本微信用户点击后将没有回应, 开发者也不能正常接收到事件推送.
	ButtonTypeScanCodePush    = "scancode_push"      // 扫码推事件
	ButtonTypeScanCodeWaitMsg = "scancode_waitmsg"   // 扫码带提示
	ButtonTypePicSysPhoto     = "pic_sysphoto"       // 系统拍照发图
	ButtonTypePicPhotoOrAlbum = "pic_photo_or_album" // 拍照或者相册发图
	ButtonTypePicWeixin       = "pic_weixin"         // 微信相册发图
	ButtonTypeLocationSelect  = "location_select"    // 发送位置

	// 下面的按钮类型专门给第三方平台旗下未微信认证(具体而言, 是资质认证未通过)的订阅号准备的事件类型,
	// 它们是没有事件推送的, 能力相对受限, 其他类型的公众号不必使用.
	ButtonTypeMediaId     = "media_id"     // 下发消息
	ButtonTypeViewLimited = "view_limited" // 跳转图文消息URL
)

type Menu struct {
	Buttons   []Button   `json:"button,omitempty"`
	MatchRule *MatchRule `json:"matchrule,omitempty"`
	MenuId    int64      `json:"menuid,omitempty"` // 有个性化菜单时查询接口返回值包含这个字段
}

type MatchRule struct {
	GroupId            string `json:"group_id,omitempty"`
	Sex                string `json:"sex,omitempty"`
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	ClientPlatformType string `json:"client_platform_type,omitempty"`
	Language           string `json:"language,omitempty"`
	TagId              string `json:"tag_id,omitempty"`
}

type Button struct {
	Type       string   `json:"type,omitempty"`       // 非必须; 菜单的响应动作类型
	Name       string   `json:"name,omitempty"`       // 必须;  菜单标题
	Key        string   `json:"key,omitempty"`        // 非必须; 菜单KEY值, 用于消息接口推送
	URL        string   `json:"url,omitempty"`        // 非必须; 网页链接, 用户点击菜单可打开链接
	MediaId    string   `json:"media_id,omitempty"`   // 非必须; 调用新增永久素材接口返回的合法media_id
	AppId      string   `json:"appid,omitempty"`      // 非必须; 跳转到小程序的appid
	PagePath   string   `json:"pagepath,omitempty"`   // 非必须; 跳转到小程序的path
	SubButtons []Button `json:"sub_button,omitempty"` // 非必须; 二级菜单数组
}

// 设置 btn 指向的 Button 为 子菜单 类型按钮.
func (btn *Button) SetAsSubMenuButton(name string, subButtons []Button) {
	btn.Name = name
	btn.SubButtons = subButtons

	btn.Type = ""
	btn.Key = ""
	btn.URL = ""
	btn.MediaId = ""
}

// 设置 btn 指向的 Button 为 click 类型按钮.
func (btn *Button) SetAsClickButton(name, key string) {
	btn.Type = ButtonTypeClick
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaId = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 view 类型按钮.
func (btn *Button) SetAsViewButton(name, url string) {
	btn.Type = ButtonTypeView
	btn.Name = name
	btn.URL = url

	btn.Key = ""
	btn.MediaId = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 扫码推事件 类型按钮.
func (btn *Button) SetAsScanCodePushButton(name, key string) {
	btn.Type = ButtonTypeScanCodePush
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaId = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 扫码推事件且弹出"消息接收中"提示框 类型按钮.
func (btn *Button) SetAsScanCodeWaitMsgButton(name, key string) {
	btn.Type = ButtonTypeScanCodeWaitMsg
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaId = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 弹出系统拍照发图 类型按钮.
func (btn *Button) SetAsPicSysPhotoButton(name, key string) {
	btn.Type = ButtonTypePicSysPhoto
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaId = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 弹出拍照或者相册发图 类型按钮.
func (btn *Button) SetAsPicPhotoOrAlbumButton(name, key string) {
	btn.Type = ButtonTypePicPhotoOrAlbum
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaId = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 弹出微信相册发图器 类型按钮.
func (btn *Button) SetAsPicWeixinButton(name, key string) {
	btn.Type = ButtonTypePicWeixin
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaId = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 弹出地理位置选择器 类型按钮.
func (btn *Button) SetAsLocationSelectButton(name, key string) {
	btn.Type = ButtonTypeLocationSelect
	btn.Name = name
	btn.Key = key

	btn.URL = ""
	btn.MediaId = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 下发消息(除文本消息) 类型按钮.
func (btn *Button) SetAsMediaIdButton(name, mediaId string) {
	btn.Type = ButtonTypeMediaId
	btn.Name = name
	btn.MediaId = mediaId

	btn.Key = ""
	btn.URL = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 跳转图文消息URL 类型按钮.
func (btn *Button) SetAsViewLimitedButton(name, mediaId string) {
	btn.Type = ButtonTypeViewLimited
	btn.Name = name
	btn.MediaId = mediaId

	btn.Key = ""
	btn.URL = ""
	btn.SubButtons = nil
}

// 设置 btn 指向的 Button 为 打开小程序.
func (btn *Button) SetAsMiniProgramButton(name, appId, pagePath, url string) {
	btn.Type = ButtonTypeMiniProgram
	btn.Name = name
	btn.URL = url
	btn.AppId = appId
	btn.PagePath = pagePath
}
