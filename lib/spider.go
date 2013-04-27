/*
	lib package
	封装spider模拟curl函数
*/
package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//爬虫类
type Spider struct {
	HOST  string
	PORT  int
	REFER string
}

//use http get methode get content
func (s *Spider) Fetch(url string) ([]byte, error) {

	fmt.Sprintf("%v\n", url)
	//create client object
	client := &http.Client{}
	//create resposne object
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		//set header
		//make resposne
		resp, err := client.Do(req)
		if err == nil {

			defer resp.Body.Close()

			content, err := ioutil.ReadAll(resp.Body)
			return content, err
		}
	}

	return []byte(""), err
}

//use http post method  sumit  data 
func (s *Spider) Submit(url string, values url.Values) ([]byte, error) {

	client := http.Client{}

	resp, err := client.PostForm(url, values)
	defer resp.Body.Close()

	if err == nil {
		content, err := ioutil.ReadAll(resp.Body)
		return content, err
	}
	return []byte(""), err
}
