package util

import (
	"reflect"
	"strings"
	"testing"
)

func TestDecodeXMLToMap(t *testing.T) {
	var xmlArr = []string{
		`<xml>
			<a>a</a>
			<b>b</b>
		</xml>`,
		`<xml>
			<a>a</a>
			<b>
				<ba>ba</ba>
			</b>
			<c>c</c>
		</xml>`,
		`<xml>
			<a>a</a>
			<b>
				bchara
				<ba>ba</ba>
			</b>
			<c>c</c>
		</xml>`,
		`<xml>
			<a>a</a>
			<b>
				bchara
				<ba>ba</ba>
				bchara
				<bb>bb</bb>
				bchara
			</b>
			<c>c</c>
		</xml>`,
		`<xml>
			chara
			<a>a</a>
			<b>
				<ba>ba</ba>
				bchara
			</b>
			<c>c</c>
		</xml>`,
	}

	var mapArr = []map[string]string{
		{
			"a": "a",
			"b": "b",
		},
		{
			"a": "a",
			"c": "c",
		},
		{
			"a": "a",
			"c": "c",
		},
		{
			"a": "a",
			"c": "c",
		},
		{
			"a": "a",
			"c": "c",
		},
	}

	for i, src := range xmlArr {
		m, err := DecodeXMLToMap(strings.NewReader(src))
		if err != nil {
			t.Errorf("DecodeXMLToMap(%s) failed: %s\n", src, err.Error())
			continue
		}
		if !reflect.DeepEqual(m, mapArr[i]) {
			t.Errorf("DecodeXMLToMap(%s) failed:\nhave %+v\nwant %+v\n", src, m, mapArr[i])
			continue
		}
	}
}
