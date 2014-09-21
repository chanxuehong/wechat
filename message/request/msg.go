// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

// 文本消息
type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	Content string `xml:"Content" json:"Content"` // 文本消息内容
}

func (req *Request) Text() (text *Text) {
	text = &Text{
		CommonHead: req.CommonHead,
		MsgId:      req.MsgId,
		Content:    req.Content,
	}
	return
}

// 图片消息
type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	PicURL  string `xml:"PicUrl"  json:"PicUrl"`  // 图片链接
}

func (req *Request) Image() (image *Image) {
	image = &Image{
		CommonHead: req.CommonHead,
		MsgId:      req.MsgId,
		MediaId:    req.MediaId,
		PicURL:     req.PicURL,
	}
	return
}

// 语音消息
type Voice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 语音消息媒体id，可以调用多媒体文件下载接口拉取该媒体
	Format  string `xml:"Format"  json:"Format"`  // 语音格式，如amr，speex等

	// 语音识别结果，UTF8编码，
	// NOTE: 需要开通语音识别功能，否则该字段为空，即使开通了语音识别该字段还是有可能为空
	Recognition string `xml:"Recognition" json:"Recognition"`
}

func (req *Request) Voice() (voice *Voice) {
	voice = &Voice{
		CommonHead:  req.CommonHead,
		MsgId:       req.MsgId,
		MediaId:     req.MediaId,
		Format:      req.Format,
		Recognition: req.Recognition,
	}
	return
}

// 视频消息
type Video struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId        int64  `xml:"MsgId"        json:"MsgId"`        // 消息id, 64位整型
	MediaId      string `xml:"MediaId"      json:"MediaId"`      // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
}

func (req *Request) Video() (video *Video) {
	video = &Video{
		CommonHead:   req.CommonHead,
		MsgId:        req.MsgId,
		MediaId:      req.MediaId,
		ThumbMediaId: req.ThumbMediaId,
	}
	return
}

// 地理位置消息
type Location struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId     int64   `xml:"MsgId"      json:"MsgId"`      // 消息id, 64位整型
	LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
	LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
	Scale     int     `xml:"Scale"      json:"Scale"`      // 地图缩放大小
	Label     string  `xml:"Label"      json:"Label"`      // 地理位置信息
}

func (req *Request) Location() (location *Location) {
	location = &Location{
		CommonHead: req.CommonHead,
		MsgId:      req.MsgId,
		LocationX:  req.LocationX,
		LocationY:  req.LocationY,
		Scale:      req.Scale,
		Label:      req.Label,
	}
	return
}

// 链接消息
type Link struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	MsgId       int64  `xml:"MsgId"       json:"MsgId"`       // 消息id, 64位整型
	Title       string `xml:"Title"       json:"Title"`       // 消息标题
	Description string `xml:"Description" json:"Description"` // 消息描述
	URL         string `xml:"Url"         json:"Url"`         // 消息链接
}

func (req *Request) Link() (link *Link) {
	link = &Link{
		CommonHead:  req.CommonHead,
		MsgId:       req.MsgId,
		Title:       req.Title,
		Description: req.Description,
		URL:         req.URL,
	}
	return
}
