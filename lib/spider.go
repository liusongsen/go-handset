package lib

//spider object 
type Spider struct {
	host  string
	port  int
	refer string
}

//use http get methode get content
func (s *Spider) Fetch(url string) (content string) {

	content = url + "content:sdfsdfsfdsdfsdf method:GET"
	return
}

//use http post method submit data 
func (s *Spider) Submit(url string, data interface{}) (content string) {

	content = url + "content:sdfsdfsdf method:POST"
	return
}
