package handset

type Imobile struct {
	card    int
	url     string
	content string
}

//构造请求URL
func (h *Imobile) MakeUrl() {
	//pass
}

//抓去网页内容
func (h *Imobile) Fetch() {
	//pass
}

//正则解析匹配
func (h *Imobile) Parse() {
	//pass
}
