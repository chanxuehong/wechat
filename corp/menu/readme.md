### 创建菜单的示例

```Go
	// 典型的创建菜单的过程
	
	var subButtons = make([]Button, 2)
	subButtons[0].InitToViewButton("搜索", "http://www.soso.com/")
	subButtons[1].InitToClickButton("赞一下我们", "V1001_GOOD")

	var mn Menu
	mn.Buttons = make([]Button, 3)
	mn.Buttons[0].InitToClickButton("今日歌曲", "V1001_TODAY_MUSIC")
	mn.Buttons[1].InitToViewButton("视频", "http://v.qq.com/")
	mn.Buttons[2].InitToSubMenuButton("菜单", subButtons)

	// =========================================================================
	
	var subButtons0 = make([]Button, 2)
	subButtons0[0].InitToScanCodeWaitMsgButton("扫码带提示", "rselfmenu_0_0")
	subButtons0[1].InitToScanCodePushButton("扫码推事件", "rselfmenu_0_1")

	var subButtons1 = make([]Button, 3)
	subButtons1[0].InitToPicSysPhotoButton("系统拍照发图", "rselfmenu_1_0")
	subButtons1[1].InitToPicPhotoOrAlbumButton("拍照或者相册发图", "rselfmenu_1_1")
	subButtons1[2].InitToPicWeixinButton("微信相册发图", "rselfmenu_1_2")

	var mn1 Menu
	mn1.Buttons = make([]Button, 3)
	mn1.Buttons[0].InitToSubMenuButton("扫码", subButtons0)
	mn1.Buttons[1].InitToSubMenuButton("发图", subButtons1)
	mn1.Buttons[2].InitToLocationSelectButton("发送位置", "rselfmenu_2_0")
```
