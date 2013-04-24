package lib

//use http get methode get content
func Fetch(url string) (content string) {

	if url == "http://www.atido.com" {
		content = "atido.com"
	} else {
		content = "atido.nets"
	}
	return
}
