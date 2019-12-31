package bot

// 消息类型，此时固定为news
type News struct {
	Articles []Article `json:"articles"` // 图文消息，一个图文消息支持1到8条图文
}

type Article struct {
	Title  string `json:"title"`                 // 标题，不超过128个字节，超过会自动截断
	Desc   string `json:"description,omitempty"` // 描述，不超过512个字节，超过会自动截断
	Url    string `json:"url"`                   // 点击后跳转的链接
	PicUrl string `json:"picurl,omitempty"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150
}

func NewNews(articles []Article) *News {
	return &Message{
		Type: NEWS,
		News: &News{
			Articles: articles,
		},
	}
}
