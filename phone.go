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

//手机
type phoner interface {
	parse() sim
}

//imobile.com
type imobile struct {
	card int
	url  string
}

//正则解析匹配
func (h imobile) parse() sim {

	//fetch url
	//parse content 
	return sim{"广东", "广州", "中国联通3G卡", "imobile"}
}

//ip138.com
type ip138 struct {
	card int
	url  string
}

//正则解析匹配
func (h ip138) parse() sim {

	//fetch url
	//parse content 
	return sim{"广东", "广州", "中国联通3G卡", "ip138"}
}

//showji.com
type shouji struct {
	card int
	url  string
}

//正则解析匹配
func (h shouji) parse() sim {

	spider := lib.Spider{"127.0.0.1", 80, "http://www.baidu.com"}
	if content, ok := spider.Fetch(h.url); ok == nil {
		fmt.Printf("%v\n", content)
	} else {
		fmt.Printf("%v\n", "sdfsdf")
	}

	//fetch url
	return sim{"广东", "广州", "中国联通3G卡", "shouji"}
}
