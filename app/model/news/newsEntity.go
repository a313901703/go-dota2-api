package news

//Entity  新闻实例
type Entity struct {
	Title   string `json:"title"`   //标题
	Thumb   string `json:"thumb"`   //标题
	Content string `json:"content"` //内容
	Date    string `json:"date"`    //时间
}
