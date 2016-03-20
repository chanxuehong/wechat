<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package card

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"sort"
)

=======
package card

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"sort"
)

// stringWriter is the interface that wraps the WriteString method.
type stringWriter interface {
	WriteString(s string) (n int, err error)
}

>>>>>>> github/v2
// 卡券通用签名方法.
//  将 strs 里面的字符串字典排序, 然后拼接成一个字符串后做 sha1 签名.
func Sign(strs []string) (signature string) {
	sort.Strings(strs)

	h := sha1.New()
<<<<<<< HEAD
	for _, str := range strs {
		io.WriteString(h, str)
=======
	if sw, ok := h.(stringWriter); ok {
		for _, str := range strs {
			sw.WriteString(str)
		}
	} else {
		bufw := bufio.NewWriterSize(h, 256)
		for _, str := range strs {
			bufw.WriteString(str)
		}
		bufw.Flush()
>>>>>>> github/v2
	}
	return hex.EncodeToString(h.Sum(nil))
}
