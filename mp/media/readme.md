### 上传图片示例

```Go
package main

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/media"
)

var TokenServer = mp.NewDefaultTokenServer("appid", "appsecret", nil)

func main() {
	clt := media.NewClient(TokenServer, nil)
	info, err := clt.UploadImage("d:\\img.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}
```
