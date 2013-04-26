package handset

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

//模拟php iconv 编码转换
func iconv(fromEncoding string, toEncoding string, str string) (content string) {

	if fromEncoding == "GBK" && toEncoding == "UTF-8" {
		content = str + "OK"
	} else {
		content = str + "NOT OK"
	}
	return
}

//生成showji.com URl
func createShowjiUrl(mobileNo int) *url.URL {

	cc := new(url.URL)

	cc.Scheme = "http"
	cc.Host = "opendata.baidu.com/api.php"
	q := cc.Query()
	q.Set("query", strconv.Itoa(mobileNo)+".")
	q.Set("co", "")
	q.Set("resource_id", "6004")
	q.Set("t", string([]byte(fmt.Sprintf("%v", time.Now().UnixNano()))[0:13]))
	q.Set("ie", "utf8")
	q.Set("oe", "gbk")
	q.Set("cb", "")
	q.Set("format", "json")
	q.Set("tn", "baidu")
	cc.RawQuery = q.Encode()
	return cc
}

//生成ip138.com URL
func createIp138Url(mobileNo int) *url.URL {

	cc := new(url.URL)

	cc.Scheme = "http"
	cc.Host = "www.ip138.com:8080/search.asp"
	q := cc.Query()
	q.Set("mobile", strconv.Itoa(mobileNo))
	q.Set("action", "mobile")
	cc.RawQuery = q.Encode()
	return cc
}

//生成imobile.com URL
func createImobileUrl(mobileNo int) *url.URL {

	cc := new(url.URL)

	cc.Scheme = "http"
	cc.Host = "search.imobile.com.cn/index.php"
	q := cc.Query()
	q.Set("keyword", strconv.Itoa(mobileNo))
	q.Set("a", "search")
	q.Set("scope", "all")
	cc.RawQuery = q.Encode()
	return cc
}
