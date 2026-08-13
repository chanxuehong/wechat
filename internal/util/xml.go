package util

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
	"strings"
)

// DecodeXMLToMap decodes xml reading from io.Reader and returns the first-level sub-node key-value set,
// if the first-level sub-node contains child nodes, skip it.
func DecodeXMLToMap(r io.Reader) (m map[string]string, err error) {
	m = make(map[string]string)
	var (
		decoder = xml.NewDecoder(r)
		depth   = 0
		token   xml.Token
		key     string
		value   strings.Builder
	)
	for {
		token, err = decoder.Token()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}

		switch v := token.(type) {
		case xml.StartElement:
			depth++
			switch depth {
			case 2:
				key = v.Name.Local
				value.Reset()
			case 3:
				if err = decoder.Skip(); err != nil {
					return
				}
				depth--
				key = "" // key == "" indicates that the node with depth==2 has children
			}
		case xml.CharData:
			if depth == 2 && key != "" {
				value.Write(v)
			}
		case xml.EndElement:
			if depth == 2 && key != "" {
				m[key] = value.String()
			}
			depth--
		}
	}
}

// EncodeXMLFromMap encodes map[string]string to io.Writer with xml format.
//
//	NOTE: This function requires the rootname argument and the keys of m (type map[string]string) argument
//	are legitimate xml name string that does not contain the required escape character!
func EncodeXMLFromMap(w io.Writer, m map[string]string, rootname string) (err error) {
	switch v := w.(type) {
	case *bytes.Buffer:
		bufw := v
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return nil
	case *strings.Builder:
		bufw := v
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return nil

	case *bufio.Writer:
		bufw := v
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return bufw.Flush()

	default:
		bufw := bufio.NewWriterSize(w, 256)
		if err = bufw.WriteByte('<'); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}

		for k, v := range m {
			if err = bufw.WriteByte('<'); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}

			if err = xml.EscapeText(bufw, []byte(v)); err != nil {
				return
			}

			if _, err = bufw.WriteString("</"); err != nil {
				return
			}
			if _, err = bufw.WriteString(k); err != nil {
				return
			}
			if err = bufw.WriteByte('>'); err != nil {
				return
			}
		}

		if _, err = bufw.WriteString("</"); err != nil {
			return
		}
		if _, err = bufw.WriteString(rootname); err != nil {
			return
		}
		if err = bufw.WriteByte('>'); err != nil {
			return
		}
		return bufw.Flush()
	}
}
