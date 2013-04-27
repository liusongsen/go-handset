/*
	手机号码采集器
*/
package handset

import (
	"fmt"
)

var c chan int
var i int
var sims map[string]sim

//采集手机号码段信息
func collector(p phone, t string) {

	switch t {
	case "ip138":
		sims[t] = p.parseIp138()
	case "imobile":
		sims[t] = p.parseImobile()
	case "shouji":
		sims[t] = p.parseShouji()
	}

	c <- 1

}

//解析手机号码
func Parse(mobileno int) {

	defer func() {
		fmt.Printf("%v\n", sims)
	}()

	c = make(chan int)
	sims = make(map[string]sim)

	p := phone{mobileno}

	go collector(p, "ip138")
	go collector(p, "imobile")
	go collector(p, "shouji")

L:
	for {
		select {
		case <-c:
			i++
			if i > 0 {
				break L
			}
		}
	}

}
