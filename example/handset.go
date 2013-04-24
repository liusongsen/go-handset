package main

import (
	"fmt"
	"github.com/liusongsen/go-handset"
	"github.com/liusongsen/go-handset/lib"
)

func main() {

	result := handset.Even(3)
	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", lib.Spider.Fetch("http://www.atido.com"))
}
