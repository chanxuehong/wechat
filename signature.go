package wechat

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"sort"
)

// 校验消息是否是从微信服务器发送过来的
func CheckSignature(signature, timestamp, nonce, token string) bool {
	if len(signature) != 40 {
		return false
	}

	strArr := sort.StringSlice{token, timestamp, nonce}
	strArr.Sort()

	// buf := []byte(strArr[0] + strArr[1] + strArr[2])
	// 因为目前 golang 还不支持 copy-on-write, 为了减少内存复制采用了这个的方法
	str0Len := len(strArr[0])
	str0str1Len := str0Len + len(strArr[1])
	str0str1str2Len := str0str1Len + len(strArr[2])
	buf := make([]byte, str0str1str2Len)
	copy(buf, strArr[0])
	copy(buf[str0Len:], strArr[1])
	copy(buf[str0str1Len:], strArr[2])

	hashSumArray := sha1.Sum(buf) // require go1.2+
	hashSumHexBytes := make([]byte, 40)
	hex.Encode(hashSumHexBytes, hashSumArray[:])

	// 现在 len(signature) == 40, 不会 panic
	if rslt := subtle.ConstantTimeCompare(hashSumHexBytes, []byte(signature)); rslt == 1 {
		return true
	}
	return false
}
