package handset

//模拟php iconv 编码转换
func iconv(fromEncoding string, toEncoding string, str string) (content string) {

	if fromEncoding == "GBK" && toEncoding == "UTF-8" {
		content = str + "OK"
	} else {
		content = str + "NOT OK"
	}
	return
}
