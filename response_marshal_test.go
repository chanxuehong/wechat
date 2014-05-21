package wechat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"regexp"
	"testing"
)

// 对于XML, 返回的不是 <![CDATA[toUser]]> 格式, 而是经过了 Escape 后的结果, 不做测试.

var responseMarshalTests = []struct {
	Value      interface{}
	ExpectXML  []byte
	ExpectJSON []byte
}{
	{ // 回复文本消息
		TextResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "text",
			},
			responseText: responseText{
				Content: "你好",
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>text</MsgType>
		<Content>你好</Content>
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"text",
			"text":
			{
				"content":"你好"
			}
		}`),
	},
	{ // 回复图片消息
		ImageResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "image",
			},
			responseImage: responseImage{
				MediaId: "media_id",
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>image</MsgType>
		<Image>
			<MediaId>media_id</MediaId>
		</Image>
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"image",
			"image":
			{
				"media_id":"media_id"
			}
		}`),
	},
	{ // 回复语音消息
		VoiceResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "voice",
			},
			responseVoice: responseVoice{
				MediaId: "media_id",
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>voice</MsgType>
		<Voice>
			<MediaId>media_id</MediaId>
		</Voice>
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"voice",
			"voice":
			{
				"media_id":"media_id"
			}
		}`),
	},
	{ // 回复视频消息
		VideoResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "video",
			},
			responseVideo: responseVideo{
				MediaId:     "media_id",
				Title:       "title",
				Description: "description",
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>video</MsgType>
		<Video>
			<MediaId>media_id</MediaId>
			<Title>title</Title>
			<Description>description</Description>
		</Video> 
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"video",
			"video":
			{
				"media_id":"media_id",
				"title":"title",
				"description":"description"
			}
		}`),
	},
	{ // 回复视频消息, 没有 title
		VideoResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "video",
			},
			responseVideo: responseVideo{
				MediaId:     "media_id",
				Description: "description",
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>video</MsgType>
		<Video>
			<MediaId>media_id</MediaId>
			<Description>description</Description>
		</Video> 
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"video",
			"video":
			{
				"media_id":"media_id",
				"description":"description"
			}
		}`),
	},
	{ // 回复视频消息, 没有 description
		VideoResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "video",
			},
			responseVideo: responseVideo{
				MediaId: "media_id",
				Title:   "title",
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>video</MsgType>
		<Video>
			<MediaId>media_id</MediaId>
			<Title>title</Title>
		</Video> 
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"video",
			"video":
			{
				"media_id":"media_id",
				"title":"title"
			}
		}`),
	},
	{ // 发送音乐消息
		MusicResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "music",
			},
			responseMusic: responseMusic{
				Title:        "TITLE",
				Description:  "DESCRIPTION",
				MusicUrl:     "MUSIC_Url",
				HQMusicUrl:   "HQ_MUSIC_Url",
				ThumbMediaId: "media_id",
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>music</MsgType>
		<Music>
			<Title>TITLE</Title>
			<Description>DESCRIPTION</Description>
			<MusicUrl>MUSIC_Url</MusicUrl>
			<HQMusicUrl>HQ_MUSIC_Url</HQMusicUrl>
			<ThumbMediaId>media_id</ThumbMediaId>
		</Music>
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"music",
			"music":
			{
				"title":"TITLE",
				"description":"DESCRIPTION",
				"musicurl":"MUSIC_Url",
				"hqmusicurl":"HQ_MUSIC_Url",
				"thumb_media_id":"media_id" 
			}
		}`),
	},
	{ // 发送音乐消息, 没有 title 和 DESCRIPTION
		MusicResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "music",
			},
			responseMusic: responseMusic{
				MusicUrl:     "MUSIC_Url",
				HQMusicUrl:   "HQ_MUSIC_Url",
				ThumbMediaId: "media_id",
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>music</MsgType>
		<Music>
			<MusicUrl>MUSIC_Url</MusicUrl>
			<HQMusicUrl>HQ_MUSIC_Url</HQMusicUrl>
			<ThumbMediaId>media_id</ThumbMediaId>
		</Music>
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"music",
			"music":
			{
				"musicurl":"MUSIC_Url",
				"hqmusicurl":"HQ_MUSIC_Url",
				"thumb_media_id":"media_id" 
			}
		}`),
	},
	{ // 回复图文消息, 文章数量 == 0
		NewsResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "news",
			},
			responseNews: responseNews{
				ArticleCount: 0,
				Articles:     make([]*Article, 0),
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>news</MsgType>
		<ArticleCount>0</ArticleCount>
		<Articles>
		</Articles>
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"news",
			"news":{
				"articles":[]
			}
		}`),
	},
	{ // 回复图文消息
		NewsResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "news",
			},
			responseNews: responseNews{
				ArticleCount: 1,
				Articles: []*Article{
					&Article{
						Title:       "title1",
						Description: "description1",
						PicUrl:      "picurl",
						Url:         "url",
					},
				},
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>news</MsgType>
		<ArticleCount>1</ArticleCount>
		<Articles>
			<item>
				<Title>title1</Title> 
				<Description>description1</Description>
				<PicUrl>picurl</PicUrl>
				<Url>url</Url>
			</item>
		</Articles>
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"news",
			"news":{
				"articles":[
					{
						"title":"title1",
						"description":"description1",
						"picurl":"picurl",
						"url":"url"
					}
				]
			}
		}`),
	},
	{ // 回复图文消息
		NewsResponseMsg{
			responseCommon: responseCommon{
				ToUserName:   "toUser",
				FromUserName: "fromUser",
				CreateTime:   12345678,
				MsgType:      "news",
			},
			responseNews: responseNews{
				ArticleCount: 2,
				Articles: []*Article{
					&Article{
						Title:       "title1",
						Description: "description1",
						PicUrl:      "picurl",
						Url:         "url",
					},
					&Article{
						Title:       "title",
						Description: "description",
						PicUrl:      "picurl",
						Url:         "url",
					},
				},
			},
		},
		[]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>news</MsgType>
		<ArticleCount>2</ArticleCount>
		<Articles>
			<item>
				<Title>title1</Title> 
				<Description>description1</Description>
				<PicUrl>picurl</PicUrl>
				<Url>url</Url>
			</item>
			<item>
				<Title>title</Title>
				<Description>description</Description>
				<PicUrl>picurl</PicUrl>
				<Url>url</Url>
			</item>
		</Articles>
		</xml>`),
		[]byte(`{
			"touser":"toUser",
			"msgtype":"news",
			"news":{
				"articles":[
					{
						"title":"title1",
						"description":"description1",
						"picurl":"picurl",
						"url":"url"
					},
					{
						"title":"title",
						"description":"description",
						"picurl":"picurl",
						"url":"url"
					}
				]
			}
		}`),
	},
}

func TestResponseXMLandJSONMarshal(t *testing.T) {
	// 匹配一个文本文件每行开头和结尾的空白(包括换行符)
	var trim_rnt = regexp.MustCompile(`(?m)(^[\x20\t]*)|([\x20\t]*(\n|\r\n|\r))`)

	for _, test := range responseMarshalTests {
		// xml
		b, err := xml.Marshal(test.Value)
		if err != nil {
			t.Errorf("xml.Marshal(%#v):\nError: %s\n", test.Value, err)
			continue
		}
		want := trim_rnt.ReplaceAllLiteral(test.ExpectXML, []byte(""))
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#v):\nhave %#s\nwant %#s\n", test.Value, b, want)
			continue
		}

		// json
		b, err = json.Marshal(test.Value)
		if err != nil {
			t.Errorf("json.Marshal(%#v):\nError: %s\n", test.Value, err)
			continue
		}
		want = trim_rnt.ReplaceAllLiteral(test.ExpectJSON, []byte(""))
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#v):\nhave %#s\nwant %#s\n", test.Value, b, want)
			continue
		}
	}
}
