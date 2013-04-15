package main

import (
	"fmt"
	curl "github.com/andelf/go-curl"
)

const (
	SHOUJI = 1
	IP138  = 2
	MOBILE = 3
)

type Handset struct {
	card     int
	province string
	city     string
	post     int
	area     string
	ctype    string
	source   int
	url      string
	size     int
	content  string
	err      error
}

//根据类型构造RUL
func (h *Handset) buildUrl(t int) {
	switch {
	case t == SHOUJI:
		h.url = "http://www.atido.com/tb.php?a=shouji"
		h.source = SHOUJI
	case t == IP138:
		h.url = "http://www.atido.com/tb.php?a=ip138"
		h.source = IP138
	case t == MOBILE:
		h.url = "http://www.atido.com/tb.php?a=mobile"
		h.source = MOBILE
	}
}

//模拟蜘蛛抓取网页内容
func (h *Handset) fetch() {
	if h.url != "" {
		easy := curl.EasyInit()
		defer easy.Cleanup()

		easy.Setopt(curl.OPT_URL, h.url)

		// make a callback function
		fooTest := func(buf []byte, userdata interface{}) bool {
			h.size = len(buf)
			h.content = string(buf)
			return true
		}

		easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

		h.err = easy.Perform()
	}
}

//解析手机号码信息
func (h *Handset) Parse() {
	h.buildUrl(MOBILE)
	fmt.Printf("%d======%s\n", h.card, h.url)
	h.fetch()
	if h.err == nil {
		h.province = "广东"
		h.city = "广州"
		h.post = 510000
		h.ctype = "中国联通3G"
		h.area = "天河"
		fmt.Printf("%v\n", h)
	} else {
		fmt.Printf("EROOR: %v\n", h.err)
	}
}

func main() {

	handset := Handset{card: 1320201}
	handset.Parse()
}
