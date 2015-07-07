// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package util

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// 把整数 n 格式化成 4 字节的网络字节序
func encodeNetworkByteOrder(orderBytes []byte, n int) {
	orderBytes[0] = byte(n >> 24)
	orderBytes[1] = byte(n >> 16)
	orderBytes[2] = byte(n >> 8)
	orderBytes[3] = byte(n)
}

// 从 4 字节的网络字节序里解析出整数
func decodeNetworkByteOrder(orderBytes []byte) (n int) {
	n = int(orderBytes[0])<<24 |
		int(orderBytes[1])<<16 |
		int(orderBytes[2])<<8 |
		int(orderBytes[3])
	return
}

// ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + AppId]
func AESEncryptMsg(random, rawXMLMsg []byte, AppId string, AESKey [32]byte) (ciphertext []byte) {
	const (
		BLOCK_SIZE = 32 // PKCS#7
		BLOCK_MASK = BLOCK_SIZE - 1
	)

	rawXMLMsgEnd := 20 + len(rawXMLMsg)
	contentLen := rawXMLMsgEnd + len(AppId)
	amountToPad := BLOCK_SIZE - contentLen&BLOCK_MASK
	plaintextLen := contentLen + amountToPad

	plaintext := make([]byte, plaintextLen)

	// 拼接
	copy(plaintext[:16], random)
	encodeNetworkByteOrder(plaintext[16:20], len(rawXMLMsg)) // len(rawXMLMsg) 正常情况下不会越界
	copy(plaintext[20:], rawXMLMsg)
	copy(plaintext[rawXMLMsgEnd:], AppId)

	// PKCS#7 补位
	for i := contentLen; i < plaintextLen; i++ {
		plaintext[i] = byte(amountToPad)
	}

	// 加密
	block, err := aes.NewCipher(AESKey[:])
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, AESKey[:16])
	mode.CryptBlocks(plaintext, plaintext)

	ciphertext = plaintext
	return
}

// ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + AppId]
func AESDecryptMsg(ciphertext []byte, AppId string, AESKey [32]byte) (random, rawXMLMsg []byte, err error) {
	const (
		BLOCK_SIZE = 32 // PKCS#7
		BLOCK_MASK = BLOCK_SIZE - 1
	)

	if len(ciphertext) < BLOCK_SIZE {
		err = fmt.Errorf("the length of encryptedMsg too short: %d", len(ciphertext))
		return
	}
	if len(ciphertext)&BLOCK_MASK != 0 {
		err = fmt.Errorf("encryptedMsg is not a multiple of the block size, the length is %d", len(ciphertext))
		return
	}

	plaintext := make([]byte, len(ciphertext)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	block, err := aes.NewCipher(AESKey[:])
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, AESKey[:16])
	mode.CryptBlocks(plaintext, ciphertext)

	// PKCS#7 去除补位
	amountToPad := int(plaintext[len(plaintext)-1])
	if amountToPad < 1 || amountToPad > BLOCK_SIZE {
		err = fmt.Errorf("the amount to pad is invalid: %d", amountToPad)
		return
	}
	plaintext = plaintext[:len(plaintext)-amountToPad]

	// 反拼装
	// len(plain) == 16+4+len(rawXMLMsg)+len(AppId)
	// len(AppId) > 0
	if len(plaintext) <= 20 {
		err = fmt.Errorf("plain msg too short, the length is %d", len(plaintext))
		return
	}
	msgLen := decodeNetworkByteOrder(plaintext[16:20])
	if msgLen < 0 {
		err = fmt.Errorf("invalid msg length: %d", msgLen)
		return
	}
	msgEnd := 20 + msgLen
	if len(plaintext) <= msgEnd {
		err = fmt.Errorf("msg length too large: %d", msgLen)
		return
	}

	AppIdHave := string(plaintext[msgEnd:])
	if AppIdHave != AppId {
		err = fmt.Errorf("AppId mismatch, have: %s, want: %s", AppIdHave, AppId)
		return
	}

	random = plaintext[:16:20]
	rawXMLMsg = plaintext[20:msgEnd]
	return
}
