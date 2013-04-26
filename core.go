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
func collector(p phoner) {

	var url string
	var sim sim

	url = p.makeUrl()
	sim = p.parse()
	sims[p.getKey()] = sim
	fmt.Printf("%s======%v\n", url, sim)
	c <- 1

}

//解析手机号码
func Parse(mobileno int) {

	defer func() {
		fmt.Printf("%v\n", sims)
	}()

	c = make(chan int)
	sims = make(map[string]sim)

	pIp138 := ip138{mobileno, ""}
	pImobile := imobile{mobileno, ""}
	pShouji := shouji{mobileno, ""}

	go collector(pIp138)
	go collector(pImobile)
	go collector(pShouji)

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
