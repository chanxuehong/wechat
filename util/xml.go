// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package util

import (
	"encoding/xml"
	"io"
)

// 解析xml, 返回第一级元素键值对。如果第一级元素有子节点，则此节点的值是子节点的xml数据。（用于微信支付）
func ParseXMLToMap(xmlReader io.Reader) (Map map[string]string, err error) {
	// 目前微信支付的 xml 都只有一级元素, 所以用此简单的实现,
	// 如果以后更改了, 再重新实现!

	d := xml.NewDecoder(xmlReader)
	Map = make(map[string]string)

	var key, value string // 保存当前扫描的节点 key, value
	depth := 0            // 当前节点的深度
	for {
		var tk xml.Token
		tk, err = d.Token()
		if err != nil {
			if err != io.EOF {
				Map = nil
				return
			}
			err = nil
			return
		}

		switch v := tk.(type) {
		case xml.StartElement:
			depth++
			if depth == 2 {
				key = v.Name.Local
			}

		case xml.CharData:
			if depth == 2 {
				value = string(v.Copy())
			}

		case xml.EndElement:
			if depth == 2 {
				Map[key] = value
			}
			depth--
		}
	}
}

// 格式化 Map 为 xml 格式（用于微信支付）
func FormatMapToXML(xmlWriter io.Writer, Map map[string]string) (err error) {
	_, err = io.WriteString(xmlWriter, "<xml>\n")
	if err != nil {
		return
	}

	for key, value := range Map {
		_, err = io.WriteString(xmlWriter, "<"+key+">")
		if err != nil {
			return
		}
		if err = xml.EscapeText(xmlWriter, []byte(value)); err != nil {
			return
		}
		_, err = io.WriteString(xmlWriter, "</"+key+">\n")
		if err != nil {
			return
		}
	}

	_, err = io.WriteString(xmlWriter, "</xml>")
	if err != nil {
		return
	}

	return
}
