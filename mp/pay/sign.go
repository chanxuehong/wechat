// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

// 财付通签名, 微信支付 2.x 涉及到财付通的交易 和 微信支付 3.x 都是使用这种签名.
//
// 签名方法为:
// a.对所有传入参数按照字段名的 ASCII 码从小到大排序（字典序）后，使用 URL 键值
// 对的格式（即 key1=value1&key2=value2…）拼接成字符串 string1，注意：值为空的参数不
// 参与签名；
// b. 在 string1 最 后 拼 接 上 key=signKey 得 到 stringSignTemp 字符串，并对
// stringSignTemp 进行 md5 运算，再将得到的字符串所有字符转换为大写，得到 sign 值
// signValue。
//
//  parameters:  待签名的参数
//  signKey:     支付签名的 Key
func TenpayMD5Sign(parameters map[string]string, signKey string) string {
	kvs := make(KVSlice, 0, len(parameters))
	for k, v := range parameters {
		if v == "" { // 值为空不参加签名
			continue
		}
		if k == "sign" {
			continue
		}

		kvs = append(kvs, KV{k, v})
	}
	sort.Sort(kvs)

	Hash := md5.New()
	hashsum := make([]byte, 32)

	for _, kv := range kvs {
		Hash.Write([]byte(kv.Key))
		Hash.Write([]byte{'='})
		Hash.Write([]byte(kv.Value))
		Hash.Write([]byte{'&'})
	}
	Hash.Write([]byte("key="))
	Hash.Write([]byte(signKey))

	hex.Encode(hashsum, Hash.Sum(nil))
	return string(bytes.ToUpper(hashsum))
}

// 微信签名, 微信支付 2.x 的签名方法.
//
// 签名方法为:
// a.对所有待签名参数按照字段名的 ASCII 码从小到大排序（字典序）后，使用 URL 键
// 值对的格式（即key1=value1&key2=value2…）拼接成字符串string1。这里需要注意的是所
// 有参数名均为小写字符，例如 appId 在排序后字符串则为 appid；
// b.对 string1 作签名算法，字段名和字段值都采用原始值（此时package的value就对应
// 了使用 2.6中描述的方式生成的 package），不进行 URL 转义。具体签名算法为 paySign =
// SHA1(string)。
//
//  parameters:       待签名的参数
//  signKey:          支付签名的 Key
//  noSignParaNames:  指定 parameters 里面不参与签名的字段
func WXSHA1SignWithoutNames(parameters map[string]string, signKey string, noSignParaNames []string) string {
	return wxSHA1Sign(parameters, signKey, noSignParaNames, false)
}

// 微信签名, 微信支付 2.x 的签名方法.
//
// 签名方法为:
// a.对所有待签名参数按照字段名的 ASCII 码从小到大排序（字典序）后，使用 URL 键
// 值对的格式（即key1=value1&key2=value2…）拼接成字符串string1。这里需要注意的是所
// 有参数名均为小写字符，例如 appId 在排序后字符串则为 appid；
// b.对 string1 作签名算法，字段名和字段值都采用原始值（此时package的value就对应
// 了使用 2.6中描述的方式生成的 package），不进行 URL 转义。具体签名算法为 paySign =
// SHA1(string)。
//
//  parameters:    待签名的参数
//  signKey:       支付签名的 Key
//  signParaNames: 指定 parameters 里面参与签名的字段
func WXSHA1SignWithNames(parameters map[string]string, signKey string, signParaNames []string) string {
	return wxSHA1Sign(parameters, signKey, signParaNames, true)
}

// 微信签名, 微信支付 2.x 的签名方法.
//
// 签名方法为:
// a.对所有待签名参数按照字段名的 ASCII 码从小到大排序（字典序）后，使用 URL 键
// 值对的格式（即key1=value1&key2=value2…）拼接成字符串string1。这里需要注意的是所
// 有参数名均为小写字符，例如 appId 在排序后字符串则为 appid；
// b.对 string1 作签名算法，字段名和字段值都采用原始值（此时package的value就对应
// 了使用 2.6中描述的方式生成的 package），不进行 URL 转义。具体签名算法为 paySign =
// SHA1(string)。
//
//  parameters:      待签名的参数
//  signKey:         支付签名的 Key
//  paraNameArray:   一般是 parameters 里面某些参数数组
//  behaviourSwitch: true:  指定 parameters 出现在 paraNameArray 里的参数才能参加签名
//                   false: 指定 parameters 出现在 paraNameArray 里的参数不能参加签名
func wxSHA1Sign(parameters map[string]string, signKey string,
	paraNameArray []string, behaviourSwitch bool) string {

	kvs := make(KVSlice, 0, len(parameters)+1)
	for k, v := range parameters {
		if behaviourSwitch {
			if !isIn(paraNameArray, k) {
				continue
			}
		} else {
			if isIn(paraNameArray, k) {
				continue
			}
		}

		lowerKey := strings.ToLower(k)
		if lowerKey == "appkey" {
			continue
		}
		kvs = append(kvs, KV{lowerKey, v})
	}
	kvs = append(kvs, KV{"appkey", signKey})
	sort.Sort(kvs)

	Hash := sha1.New()
	hasWrite := false

	for _, kv := range kvs {
		if hasWrite {
			Hash.Write([]byte{'&'})
			Hash.Write([]byte(kv.Key))
			Hash.Write([]byte{'='})
			Hash.Write([]byte(kv.Value))
		} else {
			hasWrite = true //
			Hash.Write([]byte(kv.Key))
			Hash.Write([]byte{'='})
			Hash.Write([]byte(kv.Value))
		}
	}

	return hex.EncodeToString(Hash.Sum(nil))
}

// 判断 str 是否在 strs 里面, 如果在返回 true, 否则返回 false
func isIn(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
