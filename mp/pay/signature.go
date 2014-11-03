// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"sort"
)

// 对 parameters 里的参数做 MD5 签名.
// 签名方法:
// 1. 对参数 parameters 按照 key 的 ASCII 码从小到大排序（字典序）后，使用 URL 键值对的
// 格式（即 key1=value1&key2=value2...）拼接成字符串 string1，
// 注意：值为空的参数不参与签名；
// 2. 在 string1 最后拼接上 key=Key(商户支付密钥) 得到 stringSignTemp 字符串， 并对
// stringSignTemp 进行 md5 运算，再将得到的字符串所有字符转换为大写，得到 sign 值
// signValue。
//
//  parameters:  待签名的参数
//  Key:         支付签名的 Key
func MD5Signature(parameters map[string]string, Key string) (signature string) {
	keys := make([]string, 0, len(parameters))
	for key, value := range parameters {
		if value == "" { // 值为空不参加签名
			continue
		}
		if key == "sign" {
			continue
		}

		keys = append(keys, key)
	}
	sort.Strings(keys)

	Hash := md5.New()
	hashsum := make([]byte, md5.Size*2)

	for _, key := range keys {
		value := parameters[key]

		Hash.Write([]byte(key))
		Hash.Write([]byte{'='})
		Hash.Write([]byte(value))
		Hash.Write([]byte{'&'})
	}
	Hash.Write([]byte("key="))
	Hash.Write([]byte(Key))

	hex.Encode(hashsum, Hash.Sum(nil))
	signature = string(bytes.ToUpper(hashsum))
	return
}
