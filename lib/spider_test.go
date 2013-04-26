package lib

import (
	"testing"
)

var spider Spider

func TestFetch(t *testing.T) {
	if spider.Fetch("http://www.atido.com.cn") == "atido.com" {
		t.Log("http://www.atido.com")
		t.Fail()
	}
}
