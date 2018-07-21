package core

import (
	"encoding/xml"
	"testing"
)

type xmlTestData struct {
	Data           []byte
	WantToUserName string
	WantAppID      string
	WantEncrypt    string
}

func TestXmlUnmarshal(t *testing.T) {
	testTable := []xmlTestData{
		xmlTestData{
			[]byte(`<xml>
				<ToUserName><![CDATA[gh_b1eb3f8bd6c6]]></ToUserName>
				<Encrypt><![CDATA[DlCGq+lWQuyjNNK+vDaO0zUltpdUW3u4V00WCzsdNzmZGEhrU7TPxG52viOKCWYPwTMbCzgbCtakZHyNxr5hjoZJ7ORAUYoIAGQy/LDWtAnYgDO+ppKLp0rDq+67Dv3yt+vatMQTh99NII6x9SEGpY3O2h8RpG99+NYevQiOLVKqiQYzan21sX/jE4Y3wZaeudsb4QVjqzRAPaCJ5nS3T31uIR9fjSRgHTDRDOzjQ1cHchge+t6faUhniN5VQVTE+wIYtmnejc55BmHYPfBnTkYah9+cTYnI3diUPJRRiyVocJyHlb+XOZN22dsx9yzKHBAyagaoDIV8Yyb/PahcUbsqGv5wziOgLJQIa6z93/VY7d2Kq2C2oBS+Qb+FI9jLhgc3RvCi+Yno2X3cWoqbsRwoovYdyg6jme/H7nMZn77PSxOGRt/dYiWx2NuBAF7fNFigmbRiive3DyOumNCMvA==]]></Encrypt>
			</xml>`),
			"gh_b1eb3f8bd6c6",
			"",
			`DlCGq+lWQuyjNNK+vDaO0zUltpdUW3u4V00WCzsdNzmZGEhrU7TPxG52viOKCWYPwTMbCzgbCtakZHyNxr5hjoZJ7ORAUYoIAGQy/LDWtAnYgDO+ppKLp0rDq+67Dv3yt+vatMQTh99NII6x9SEGpY3O2h8RpG99+NYevQiOLVKqiQYzan21sX/jE4Y3wZaeudsb4QVjqzRAPaCJ5nS3T31uIR9fjSRgHTDRDOzjQ1cHchge+t6faUhniN5VQVTE+wIYtmnejc55BmHYPfBnTkYah9+cTYnI3diUPJRRiyVocJyHlb+XOZN22dsx9yzKHBAyagaoDIV8Yyb/PahcUbsqGv5wziOgLJQIa6z93/VY7d2Kq2C2oBS+Qb+FI9jLhgc3RvCi+Yno2X3cWoqbsRwoovYdyg6jme/H7nMZn77PSxOGRt/dYiWx2NuBAF7fNFigmbRiive3DyOumNCMvA==`,
		},
		xmlTestData{
			[]byte(`<xml>
				<AppId><![CDATA[wxc792941638093b10]]></AppId>
				<Encrypt><![CDATA[r2sCuDQ4vOp4r4y3TjE/HK+LiethCVOiwMdQAasoLjxqiPuGdJG+hDBhXj06uL5MQyoCMJ34sRl1ilstnmuRrThy43J+M7egX0SYe01S721F3uOjCv2sDEzA04PEu+3GFydJ01VPj4uYzldHrH4F3YLCJBrxHWR/0InA2bnERozZ34IZkZjdDbmAjdlPXgM+6iWpy7R70yPlN+cZFGhwpzSyiINwlOSKQZ8eoXSufrn2RXobq9V5x6WfwJuZfq47m6yaE6Mjvm9g6YWJ0POsw8wfQ0pFL2XwkYj85O8yOfay5UxXR1cTKz3H5ezfxQYSFiIv+qtGWRyZfGBV+qLdCsbrknbo8NsOv4/ECr4DJvrPMbv0vXsua396yS1QAmQoE9iNpB938iXZ7Uk+okU0/vrDA0OgdYQeM35lD3plA7E5zkd1cSkoGyrMOYR2i2b6fM2G0s4Pj3Aw3MqxQw6SVg==]]></Encrypt>
			</xml>`),
			"",
			"wxc792941638093b10",
			`r2sCuDQ4vOp4r4y3TjE/HK+LiethCVOiwMdQAasoLjxqiPuGdJG+hDBhXj06uL5MQyoCMJ34sRl1ilstnmuRrThy43J+M7egX0SYe01S721F3uOjCv2sDEzA04PEu+3GFydJ01VPj4uYzldHrH4F3YLCJBrxHWR/0InA2bnERozZ34IZkZjdDbmAjdlPXgM+6iWpy7R70yPlN+cZFGhwpzSyiINwlOSKQZ8eoXSufrn2RXobq9V5x6WfwJuZfq47m6yaE6Mjvm9g6YWJ0POsw8wfQ0pFL2XwkYj85O8yOfay5UxXR1cTKz3H5ezfxQYSFiIv+qtGWRyZfGBV+qLdCsbrknbo8NsOv4/ECr4DJvrPMbv0vXsua396yS1QAmQoE9iNpB938iXZ7Uk+okU0/vrDA0OgdYQeM35lD3plA7E5zkd1cSkoGyrMOYR2i2b6fM2G0s4Pj3Aw3MqxQw6SVg==`,
		},
	}

	for _, d := range testTable {
		x := cipherRequestHttpBody{}
		if err := xmlUnmarshal(d.Data, &x); err != nil {
			t.Error(err)
			return
		}
		if x.ToUserName != d.WantToUserName {
			t.Errorf("ToUserName mismatch,\nhave: %s\nwant: %s\n", x.ToUserName, d.WantToUserName)
			return
		}

		if x.AppID != d.WantAppID {
			t.Errorf("ToUserName mismatch,\nhave: %s\nwant: %s\n", x.AppID, d.WantAppID)
			return
		}

		if string(x.Base64EncryptedMsg) != d.WantEncrypt {
			t.Errorf("Encrypt mismatch,\nhave: %s\nwant: %s\n", x.Base64EncryptedMsg, d.WantEncrypt)
			return
		}
	}

}

func BenchmarkXmlUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	data := []byte(`<xml>
    <ToUserName><![CDATA[gh_b1eb3f8bd6c6]]></ToUserName>
    <Encrypt><![CDATA[DlCGq+lWQuyjNNK+vDaO0zUltpdUW3u4V00WCzsdNzmZGEhrU7TPxG52viOKCWYPwTMbCzgbCtakZHyNxr5hjoZJ7ORAUYoIAGQy/LDWtAnYgDO+ppKLp0rDq+67Dv3yt+vatMQTh99NII6x9SEGpY3O2h8RpG99+NYevQiOLVKqiQYzan21sX/jE4Y3wZaeudsb4QVjqzRAPaCJ5nS3T31uIR9fjSRgHTDRDOzjQ1cHchge+t6faUhniN5VQVTE+wIYtmnejc55BmHYPfBnTkYah9+cTYnI3diUPJRRiyVocJyHlb+XOZN22dsx9yzKHBAyagaoDIV8Yyb/PahcUbsqGv5wziOgLJQIa6z93/VY7d2Kq2C2oBS+Qb+FI9jLhgc3RvCi+Yno2X3cWoqbsRwoovYdyg6jme/H7nMZn77PSxOGRt/dYiWx2NuBAF7fNFigmbRiive3DyOumNCMvA==]]></Encrypt>
</xml>`)
	var x cipherRequestHttpBody
	for i := 0; i < b.N; i++ {
		xmlUnmarshal(data, &x)
	}
}

func BenchmarkStdXmlUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	data := []byte(`<xml>
    <ToUserName><![CDATA[gh_b1eb3f8bd6c6]]></ToUserName>
    <Encrypt><![CDATA[DlCGq+lWQuyjNNK+vDaO0zUltpdUW3u4V00WCzsdNzmZGEhrU7TPxG52viOKCWYPwTMbCzgbCtakZHyNxr5hjoZJ7ORAUYoIAGQy/LDWtAnYgDO+ppKLp0rDq+67Dv3yt+vatMQTh99NII6x9SEGpY3O2h8RpG99+NYevQiOLVKqiQYzan21sX/jE4Y3wZaeudsb4QVjqzRAPaCJ5nS3T31uIR9fjSRgHTDRDOzjQ1cHchge+t6faUhniN5VQVTE+wIYtmnejc55BmHYPfBnTkYah9+cTYnI3diUPJRRiyVocJyHlb+XOZN22dsx9yzKHBAyagaoDIV8Yyb/PahcUbsqGv5wziOgLJQIa6z93/VY7d2Kq2C2oBS+Qb+FI9jLhgc3RvCi+Yno2X3cWoqbsRwoovYdyg6jme/H7nMZn77PSxOGRt/dYiWx2NuBAF7fNFigmbRiive3DyOumNCMvA==]]></Encrypt>
</xml>`)
	var x cipherRequestHttpBody
	for i := 0; i < b.N; i++ {
		xml.Unmarshal(data, &x)
	}
}
