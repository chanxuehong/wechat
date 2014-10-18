// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

// 菜单
//
//  {
//      "button": [
//          {
//              "type": "click",
//              "name": "今日歌曲",
//              "key": "V1001_TODAY_MUSIC"
//          },
//          {
//              "name": "菜单",
//              "sub_button": [
//                  {
//                      "type": "view",
//                      "name": "搜索",
//                      "url": "http://www.soso.com/"
//                  },
//                  {
//                      "type": "view",
//                      "name": "视频",
//                      "url": "http://v.qq.com/"
//                  },
//                  {
//                      "type": "click",
//                      "name": "赞一下我们",
//                      "key": "V1001_GOOD"
//                  }
//              ]
//          }
//      ]
//  }
//
//  {
//      "button": [
//          {
//              "name": "扫码",
//              "sub_button": [
//                  {
//                      "type": "scancode_waitmsg",
//                      "name": "扫码带提示",
//                      "key": "rselfmenu_0_0",
//                      "sub_button": [ ]
//                  },
//                  {
//                      "type": "scancode_push",
//                      "name": "扫码推事件",
//                      "key": "rselfmenu_0_1",
//                      "sub_button": [ ]
//                  }
//              ]
//          },
//          {
//              "name": "发图",
//              "sub_button": [
//                  {
//                      "type": "pic_sysphoto",
//                      "name": "系统拍照发图",
//                      "key": "rselfmenu_1_0",
//                      "sub_button": [ ]
//                  },
//                  {
//                      "type": "pic_photo_or_album",
//                      "name": "拍照或者相册发图",
//                      "key": "rselfmenu_1_1",
//                      "sub_button": [ ]
//                  },
//                  {
//                      "type": "pic_weixin",
//                      "name": "微信相册发图",
//                      "key": "rselfmenu_1_2",
//                      "sub_button": [ ]
//                  }
//              ]
//          },
//          {
//              "name": "发送位置",
//              "type": "location_select",
//              "key": "rselfmenu_2_0"
//          }
//      ]
//  }

type Menu struct {
	Buttons []Button `json:"button,omitempty"` // 一级菜单数组，个数应为1~3个
}

// 菜单的按钮
type Button struct {
	Type       string   `json:"type,omitempty"`       // 菜单的响应动作类型
	Name       string   `json:"name"`                 // 菜单标题，不超过16个字节，子菜单不超过40个字节
	Key        string   `json:"key,omitempty"`        // 菜单KEY值，用于消息接口推送，不超过128字节
	URL        string   `json:"url,omitempty"`        // 网页链接，用户点击菜单可打开链接，不超过256字节
	SubButtons []Button `json:"sub_button,omitempty"` // 二级菜单数组，个数应为1~5个
}

// 初始化 btn 指向的 Button 为 子菜单 类型按钮
func (btn *Button) InitToSubMenuButton(name string, subButtons []Button) {
	btn.Name = name
	btn.SubButtons = subButtons

	// 容错性考虑, 清除其他字段
	btn.Type = ""
	btn.Key = ""
	btn.URL = ""
}

// 初始化 btn 指向的 Button 为 click 类型按钮
func (btn *Button) InitToClickButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_CLICK
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 view 类型按钮
func (btn *Button) InitToViewButton(name, url string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_VIEW
	btn.URL = url

	// 容错性考虑, 清除其他字段
	btn.Key = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 扫码推事件 类型按钮
func (btn *Button) InitToScanCodePushButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_SCANCODE_PUSH
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 扫码推事件且弹出“消息接收中”提示框 类型按钮
func (btn *Button) InitToScanCodeWaitMsgButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_SCANCODE_WAITMSG
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 弹出系统拍照发图 类型按钮
func (btn *Button) InitToPicSysPhotoButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_PIC_SYSPHOTO
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 弹出拍照或者相册发图 类型按钮
func (btn *Button) InitToPicPhotoOrAlbumButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_PIC_PHOTO_OR_ALBUM
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 弹出微信相册发图器 类型按钮
func (btn *Button) InitToPicWeixinButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_PIC_WEIXIN
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}

// 初始化 btn 指向的 Button 为 弹出地理位置选择器 类型按钮
func (btn *Button) InitToLocationSelectButton(name, key string) {
	btn.Name = name
	btn.Type = BUTTON_TYPE_LOCATION_SELECT
	btn.Key = key

	// 容错性考虑, 清除其他字段
	btn.URL = ""
	btn.SubButtons = nil
}
