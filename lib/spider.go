package lib

type Spider struct {
	HOST  string
	PORT  int
	REFER string
}

//use http get methode get content
func (s *Spider) Fetch(url string) (content string) {

	if url == "http://www.atido.com" {
		content = "atido.com"
	} else {
		content = "atido.nets"
	}
	return
}
