package message

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
)

// 校验消息是否是从微信服务器发送过来的
func CheckSignature(signature, timestamp, nonce, token string) bool {
	strArr := sort.StringSlice{token, timestamp, nonce}
	strArr.Sort()

	// buf := []byte(strArr[0] + strArr[1] + strArr[2])
	// 因为目前 golang 还不支持 copy-on-write, 为了减少内存复制采用了变通的方法
	str0Len := len(strArr[0])
	str0str1Len := str0Len + len(strArr[1])
	str0str1str2Len := str0str1Len + len(strArr[2])
	buf := make([]byte, str0str1str2Len)
	copy(buf, strArr[0])
	copy(buf[str0Len:], strArr[1])
	copy(buf[str0str1Len:], strArr[2])

	sha1SumBytes := sha1.Sum(buf) // 要求 go1.2+
	sha1SumStr := hex.EncodeToString(sha1SumBytes[:])

	if sha1SumStr == signature {
		return true
	} else {
		return false
	}
}
