package handset

import (
	"fmt"
	"github.com/liusongsen/go-handset/lib"
)

//sim信息元数据
type sim struct {
	province string "省份"
	city     string "城市"
	ctype    string "卡类型"
	source   string "来源"
}

//手机号码
type phone struct {
	card int
}

//匹配imobile.com.cn
func (h phone) parseImobile() sim {

	//fetch url
	url := createImobileUrl(h.card).String()
	fmt.Printf("%v\n", url)
	s := new(lib.Spider)
	s.HOST = "127.0.0.1"
	s.PORT = 80
	s.REFER = "http://www.baidu.com"

	content, _ := s.Fetch(url)

	fmt.Printf("%v\n", content)
	//parse content 
	return sim{"广东", "广州", "中国联通3G卡", "imobile"}
}

//匹配ip138.com
func (h phone) parseIp138() sim {

	//fetch url
	//parse content 
	return sim{"广东", "广州", "中国联通3G卡", "ip138"}
}

//匹配showji.com
func (h phone) parseShouji() sim {

	//fetch url
	return sim{"广东", "广州", "中国联通3G卡", "shouji"}
}
