package lib

import (
	"testing"
)

func TestFetch(t *testing.T) {
	if Fetch("http://www.atido.com.cn") == "atido.com" {
		t.Log("http://www.atido.com")
		t.Fail()
	}
}
