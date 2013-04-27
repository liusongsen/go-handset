package lib

import (
	"testing"
)

func TestFetch(t *testing.T) {

	spider := Spider{"127.0.0.1", 23456, "http://www.baidu.com"}
	if spider.Fetch("http://www.atido.com.cn") == "atido.com" {
		t.Log("http://www.atido.com")
		t.Fail()
	}
}

func TestSubmit(t *testing.T) {

	spider := Spider{"127.0.0.1", 23456, "http://www.google.com"}
	if spider.Submit("http://www.atido.com.cn") == "http://www.atido.com.cn" {
		t.Log("http://www.atido.com")
		t.Fail()
	}
}
