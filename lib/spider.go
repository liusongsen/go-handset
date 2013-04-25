package lib

type Spider struct {
	host  string
	port  int
	refer string
}

//use http get methode get content ok
func (s *Spider) Fetch(url string) (content string) {

	if url == "http://www.atido.com" {
		content = "atido.com"
	} else {
		content = "atido.nets"
	}
	return
}
