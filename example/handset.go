package main

import (
	"fmt"
	"github.com/liusongsen/go-handset"
)

func main() {

	fmt.Printf("%v\n", "Start Collect")
	handset.Parse(1320201)
	fmt.Printf("End Collect")
}
