package lib

import (
	"net/url"
	"testing"
)

func TestFetch(t *testing.T) {

	spider := Spider{"127.0.0.1", 23456, "http://www.baidu.com"}

	body, err := spider.Fetch("http://www.atido.com")

	if err != nil {
		t.Log(string(body))
		t.Fail()
	}
}

func TestSubmit(t *testing.T) {

	spider := Spider{"127.0.0.1", 23456, "http://www.baidu.com"}

	v := url.Values{}
	v.Set("username", "pickerliu")
	v.Add("password", "sdfsdf")
	v.Add("imgCode", "3345")

	body, err := spider.Submit("http://u.atido.com/index.php?m=ucenter&c=user&a=userLogin", v)

	if err != nil {
		t.Log(string(body))
		t.Fail()
	}
}
