package lib

type Spider struct {
	HOST  string
	PORT  int
	REFER string
}

//use http get methode get content
func (s Spider) Fetch(url string) (content string) {

	if url == "http://www.atido.com" {
		content = "atido.com"
	} else {
		content = "atido.nets"
	}
	return
}

//模拟表单POST数据
func (s Spider) Submit(url string) (content string) {

	content = url + "post"
	return
}
