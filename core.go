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

	if t, ok := p.(phoner); ok {

		sim := p.parse()
		sims[fmt.Sprintf("%T", t)] = sim
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

	pIp138 := ip138{mobileno, createIp138Url(mobileno).String()}
	pImobile := imobile{mobileno, createImobileUrl(mobileno).String()}
	pShouji := shouji{mobileno, createShowjiUrl(mobileno).String()}

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
