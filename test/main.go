package main

import (
	"fmt"
	"github.com/hunterhug/GoSpider/spider"
)

func main() {
	a := spider.NewAPI()
	a.SetUrl("http://aws.lenggirl.com:12345/help?user=smart&password=smart2016")
	for {
		for i := 0; i < 500; i++ {
			func() {
				_, e := a.Get()
				if e != nil {
					fmt.Println(e.Error())
				}
				fmt.Println(i)
			}()
		}
		fmt.Println("xx")
	}
}
