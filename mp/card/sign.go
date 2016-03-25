package card

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"sort"
)

// 卡券通用签名方法.
//  将 strs 里面的字符串字典排序, 然后拼接成一个字符串后做 sha1 签名.
func Sign(strs []string) (signature string) {
	sort.Strings(strs)

	h := sha1.New()

	bufw := bufio.NewWriterSize(h, 128) // sha1.BlockSize 的整数倍
	for _, str := range strs {
		bufw.WriteString(str)
	}
	bufw.Flush()

	return hex.EncodeToString(h.Sum(nil))
}
