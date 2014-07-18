// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"sort"
)

// 校验消息是否是从微信服务器发送过来的.
//  使用 buf 能提高一点性能 和 减少一些对 GC 的压力, buf 的长度要足够大, 建议不小于 256;
//  因为是内部用, 所以多个参数也无所谓, 也许以后 go 的 runtime 对于 string 和 []byte 之间
//  的转换支持  copy-on-write, 会改变实现, 如下:
//
//  func checkSignature(signature, timestamp, nonce, token string) bool {
//      const hashSumLen = 40 // sha1
//
//      if len(signature) != hashSumLen {
//          return false
//      }
//
//      strArray := sort.StringSlice{token, timestamp, nonce}
//      strArray.Sort()
//
//      h := sha1.New()
//      h.Write([]byte(strArray[0]))
//      h.Write([]byte(strArray[1]))
//      h.Write([]byte(strArray[2]))
//
//      hashsumBytes := h.Sum(nil)
//      hashsumHexBytes := make([]byte, hashSumLen)
//
//      hex.Encode(hashsumHexBytes, hashsumBytes)
//
//      // 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
//      if rslt := subtle.ConstantTimeCompare(hashsumHexBytes, []byte(signature)); rslt == 1 {
//          return true
//      }
//      return false
//  }
//
func checkSignature(signature, timestamp, nonce, token string, buf []byte) bool {
	const hashSumLen = 40 // sha1
	const twoHashSumLen = hashSumLen * 2

	if len(signature) != hashSumLen {
		return false
	}

	strArray := sort.StringSlice{token, timestamp, nonce}
	strArray.Sort()

	// buf[:hashSumLen] 保存参数 signature, buf[hashSumLen:twoHashSumLen] 保存生成的签名
	// buf[twoHashSumLen:] 按照字典序列保存 timestamp, nonce, token
	n := twoHashSumLen + len(timestamp) + len(nonce) + len(token)
	if len(buf) < n {
		buf = make([]byte, twoHashSumLen, n)
	} else {
		buf = buf[:twoHashSumLen]
	}
	buf = append(buf, strArray[0]...)
	buf = append(buf, strArray[1]...)
	buf = append(buf, strArray[2]...)

	hashsumArray := sha1.Sum(buf[twoHashSumLen:]) // require go1.2+

	hashsumHexBytes := buf[hashSumLen:twoHashSumLen]
	hex.Encode(hashsumHexBytes, hashsumArray[:])

	copy(buf, signature)

	// 采用 subtle.ConstantTimeCompare 是防止 计时攻击!
	if rslt := subtle.ConstantTimeCompare(hashsumHexBytes, buf[:hashSumLen]); rslt == 1 {
		return true
	}
	return false
}
