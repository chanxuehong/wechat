// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
)

// 解析xml, 返回第一级元素键值对。
// 如果第一级元素有子节点，则此节点的值是子节点的xml数据。
// 用于微信支付
func ParseXMLToMap(xmlReader io.Reader, Map map[string]string) (err error) {
	if xmlReader == nil {
		err = errors.New("xmlReader == nil")
		return
	}
	if Map == nil {
		err = errors.New("Map == nil")
		return
	}

	// 目前微信支付的 xml 都只有一级元素, 所以用此简单的实现,
	// 如果以后更改了, 再重新实现!

	d := xml.NewDecoder(xmlReader)

	var key string          // 保存当前第一级节点的 key
	var buffer bytes.Buffer // 保存当前第一级节点的 value
	depth := 0              // 当前节点的深度

	for {
		var tk xml.Token
		tk, err = d.Token()
		if err != nil {
			if err == io.EOF {
				err = nil
				return
			}
			return
		}

		switch v := tk.(type) {
		case xml.StartElement:
			depth++
			if depth == 2 {
				key = v.Name.Local
				buffer.Reset()
			} else {
				key = "" // 不是第一级节点的都不保存
			}

		case xml.CharData:
			if depth == 2 && key != "" {
				buffer.Write(v)
			}

		case xml.EndElement:
			if depth == 2 && key != "" {
				Map[key] = buffer.String()
			}
			depth--
		}
	}
}

// 格式化 map[string]string 为 xml 格式
// 用于微信支付
func FormatMapToXML(xmlWriter io.Writer, Map map[string]string) (err error) {
	if xmlWriter == nil {
		return errors.New("xmlWriter == nil")
	}

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
