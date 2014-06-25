package server

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"sort"
)

// 校验消息是否是从微信服务器发送过来的.
func CheckSignature(signature, timestamp, nonce, token string) bool {
	const hashsumLen = 40 // sha1

	if len(signature) != hashsumLen {
		return false
	}

	buf := make([]byte, hashsumLen, hashsumLen+len(timestamp)+len(nonce)+len(token))

	strArray := sort.StringSlice{token, timestamp, nonce}
	strArray.Sort()

	buf = append(buf, strArray[0]...)
	buf = append(buf, strArray[1]...)
	buf = append(buf, strArray[2]...)

	hashsumArray := sha1.Sum(buf[hashsumLen:]) // require go1.2+

	hashsumHexBytes := buf[:hashsumLen]
	hex.Encode(hashsumHexBytes, hashsumArray[:])

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	if rslt := subtle.ConstantTimeCompare(hashsumHexBytes, []byte(signature)); rslt == 1 {
		return true
	}
	return false
}

// 校验消息是否是从微信服务器发送过来的.
// 使用 buf 能提高一点性能 和 减少一些对 GC 的压力, buf 的长度最好 >=128
func CheckSignatureEx(signature, timestamp, nonce, token string, buf []byte) bool {
	const hashsumLen = 40 // sha1

	if len(signature) != hashsumLen {
		return false
	}

	bufLen := hashsumLen + len(timestamp) + len(nonce) + len(token)
	if len(buf) < bufLen {
		buf = make([]byte, hashsumLen, bufLen)
	} else {
		buf = buf[:hashsumLen]
	}

	strArray := sort.StringSlice{token, timestamp, nonce}
	strArray.Sort()

	buf = append(buf, strArray[0]...)
	buf = append(buf, strArray[1]...)
	buf = append(buf, strArray[2]...)

	hashsumArray := sha1.Sum(buf[hashsumLen:]) // require go1.2+

	hashsumHexBytes := buf[:hashsumLen]
	hex.Encode(hashsumHexBytes, hashsumArray[:])

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	if rslt := subtle.ConstantTimeCompare(hashsumHexBytes, []byte(signature)); rslt == 1 {
		return true
	}
	return false
}
