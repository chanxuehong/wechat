package template

type Template struct {
	Id      string `json:"template_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Example string `json:"example"`
}
