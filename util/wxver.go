package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 获取微信客户端的版本.
//  userAgent: 微信内置浏览器的 User-Agent;
//  x, y, z, w:   如果微信版本为 5.3.1.2 则有 x==5, y==3, z==1, w==2
//  err:       错误信息
func WXVersion(userAgent string) (x, y, z, w int, err error) {
	userAgent = versionRegexp.FindString(userAgent)
	if userAgent == "" {
		err = fmt.Errorf("不是有效的微信浏览器 User-Agent: %s", userAgent)
		return
	}
	userAgent = userAgent[len("MicroMessenger/"):]

	strArr := strings.Split(userAgent, ".")
	verArr := make([]int, len(strArr))

	for i, str := range strArr {
		verArr[i], err = strconv.Atoi(str)
		if err != nil {
			err = fmt.Errorf("不是有效的微信浏览器 User-Agent: %s", userAgent)
			return
		}
	}

	switch len(verArr) {
	default:
		fallthrough
	case 4:
		x = verArr[0]
		y = verArr[1]
		z = verArr[2]
		w = verArr[3]
		return
	case 3:
		x = verArr[0]
		y = verArr[1]
		z = verArr[2]
		return
	case 2:
		x = verArr[0]
		y = verArr[1]
		return
	case 1:
		x = verArr[0]
		return
	}
}

// Mozilla/5.0 (Linux; Android 4.4.4; Che1-CL10 Build/Che1-CL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.49 Mobile MQQBrowser/6.2 TBS/043128 Safari/537.36 MicroMessenger/6.5.7.1041 NetType/WIFI Language/zh_CN
var versionRegexp = regexp.MustCompile(`MicroMessenger/\d+(\.\d+)*`)
