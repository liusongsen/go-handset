package main

import (
	"fmt"
	"github.com/liusongsen/go-handset"
	"github.com/liusongsen/go-handset/lib"
	"time"
)

func main() {

	fmt.Printf("=========%v:%d==========\n", "Start", time.Now().Unix())
	handset.Parse(1320201)
	fmt.Printf("=========%v:%d==========\n", "End", time.Now().Unix())

	spider := new(lib.Spider)
	content, _ := spider.Fetch("http://www.baidu.com")
	fmt.Printf("%v\n", string(content))

}
