package main

import (
	"fmt"
	"github.com/liusongsen/go-handset"
	"time"
)

func main() {

	fmt.Printf("=========%v:%d==========\n", "Start", time.Now().Unix())
	handset.Parse(1320201)
	fmt.Printf("=========%v:%d==========\n", "End", time.Now().Unix())

}
