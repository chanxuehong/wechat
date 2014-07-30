### 创建菜单的示例

```Go
	var buttons = make([]Button, 3)
	buttons[0].InitToViewButton("搜索", "http://www.soso.com/")
	buttons[1].InitToViewButton("视频", "http://v.qq.com/")
	buttons[2].InitToClickButton("赞一下我们", "V1001_GOOD")

	var mn Menu
	mn.Buttons = make([]Button, 3)
	mn.Buttons[0].InitToClickButton("今日歌曲", "V1001_TODAY_MUSIC")
	mn.Buttons[1].InitToClickButton("歌手简介", "V1001_TODAY_SINGER")
	mn.Buttons[2].InitToSubMenuButton("菜单", buttons)
```

### 已测试通过