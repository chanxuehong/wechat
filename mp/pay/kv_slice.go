// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

type KV struct {
	Key   string
	Value string
}

type KVSlice []KV

// implement sort.Interface
func (p KVSlice) Len() int           { return len(p) }
func (p KVSlice) Less(i, j int) bool { return p[i].Key < p[j].Key }
func (p KVSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
