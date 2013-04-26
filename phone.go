package handset

import (
	"strconv"
)

const (
	IMOBILE = "http://www.imobile.com.cn/"
	IP138   = "http://www.ip138.com/"
	SHOJI   = "http://www.showji.com/"
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
	makeUrl() string
	parse() sim
	getKey() string
}

//imobile.com
type imobile struct {
	card int
	url  string
}

//构造请求URL
func (h imobile) makeUrl() string {

	h.url = createImobileUrl(h.card).String()
	return h.url
}

//正则解析匹配
func (h imobile) parse() sim {

	//fetch url
	//parse content 
	return sim{"广东", "广州", "中国联通3G卡", "imobile"}
}

//获取key
func (h imobile) getKey() string {

	return strconv.Itoa(h.card) + "_IMOBILE"
}

//ip138.com
type ip138 struct {
	card int
	url  string
}

//构造请求URL
func (h ip138) makeUrl() string {

	h.url = createIp138Url(h.card).String()
	return h.url
}

//正则解析匹配
func (h ip138) parse() sim {

	//fetch url
	//parse content 
	return sim{"广东", "广州", "中国联通3G卡", "ip138"}
}

//获取key
func (h ip138) getKey() string {

	return strconv.Itoa(h.card) + "_IP138"
}

//showji.com
type shouji struct {
	card int
	url  string
}

//构造请求URL
func (h shouji) makeUrl() string {

	h.url = createShowjiUrl(h.card).String()
	return h.url
}

//正则解析匹配
func (h shouji) parse() sim {

	//fetch url
	//parse content 
	return sim{"广东", "广州", "中国联通3G卡", "shouji"}
}

//获取key
func (h shouji) getKey() string {

	return strconv.Itoa(h.card) + "_SHOUJI"
}
