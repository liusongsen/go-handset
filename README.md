go-handset
==========

解析手机号码

install

go get github.com/liusongsen/go-handset


use:

import (
	"fmt"
	"handset"
)

handset.Parse(1320201)


[ `replay` | done: 5.668ms ]

	=========Start:1366986942==========
	http://www.ip138.com:8080/search.asp?action=mobile&mobile=1320201
	http://opendata.baidu.com/api.php?t=1366986942736&query=1320201.&cb=&resource_id=6004&co=&oe=gbk&format=json&ie=utf8&tn=baidu
	http://search.imobile.com.cn/index.php?a=search&scope=all&keyword=1320201
	=========End:1366986942==========


example:

	go build handset.go

	./handset

test:

	go test 

[ `go test` | done: 1.21986s ]

	http://www.ip138.com:8080/search.asp?mobile=1320201&action=mobile
	http://opendata.baidu.com/api.php?t=1366986995908&co=&oe=gbk&format=json&resource_id=6004&ie=utf8&cb=&query=1320201.&tn=baidu
	http://search.imobile.com.cn/index.php?a=search&scope=all&keyword=1320201
	PASS
	ok  	_/home/liusongsen/go-work/go-handset	0.008s

