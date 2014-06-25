package wechat

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"sort"
)

// 校验消息是否是从微信服务器发送过来的
func CheckSignature(signature, timestamp, nonce, token string) bool {
	const hashsumLen = 40 // sha1
	if len(signature) != hashsumLen {
		return false
	}

	strArr := sort.StringSlice{token, timestamp, nonce}
	strArr.Sort()

	// buf := []byte(strArr[0] + strArr[1] + strArr[2])
	// 因为目前 golang 还不支持 string 和 []byte 之间的 copy-on-write,
	// 为了减少内存复制采用了这个的方法
	hashsumStr0Len := hashsumLen + len(strArr[0])
	hashsumStr0Str1Len := hashsumStr0Len + len(strArr[1])
	hashsumStr0Str1Str2Len := hashsumStr0Str1Len + len(strArr[2])
	buf := make([]byte, hashsumStr0Str1Str2Len)
	hashsumHexBytes := buf[:hashsumLen]

	copy(buf[hashsumLen:], strArr[0])
	copy(buf[hashsumStr0Len:], strArr[1])
	copy(buf[hashsumStr0Str1Len:], strArr[2])
	hashsumArray := sha1.Sum(buf[hashsumLen:]) // require go1.2+

	hex.Encode(hashsumHexBytes, hashsumArray[:])

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	// 现在 len(signature) == hashsumLen, 不会 panic
	if rslt := subtle.ConstantTimeCompare(hashsumHexBytes, []byte(signature)); rslt == 1 {
		return true
	}
	return false
}

// 校验消息是否是从微信服务器发送过来的
//  NOTE: 确保 len(buf) > 40, 否则会 panic
func CheckSignatureEx(signature, timestamp, nonce, token string, buf []byte) bool {
	const hashsumLen = 40 // sha1
	if len(signature) != hashsumLen {
		return false
	}

	strArr := sort.StringSlice{token, timestamp, nonce}
	strArr.Sort()

	buf = buf[:hashsumLen]
	buf = append(buf, strArr[0]...)
	buf = append(buf, strArr[1]...)
	buf = append(buf, strArr[2]...)

	hashsumArray := sha1.Sum(buf[hashsumLen:]) // require go1.2+
	hashsumHexBytes := buf[:hashsumLen]
	hex.Encode(hashsumHexBytes, hashsumArray[:])

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	// 现在 len(signature) == hashsumLen, 不会 panic
	if rslt := subtle.ConstantTimeCompare(hashsumHexBytes, []byte(signature)); rslt == 1 {
		return true
	}
	return false
}
