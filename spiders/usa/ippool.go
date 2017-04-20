package main

import (
	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/AmazonBigSpider/public/core"
	"fmt"
)

func main() {
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir+"/config/usa_local_config.json", AmazonBigSpider.Dir+"/config/usa_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir+"/config/usa_config.json", AmazonBigSpider.Dir+"/config/usa_log.json")
	}
	//Todo
	go func() {
		host := ":12345"
		fmt.Println("开始监控IP")
		ac := &core.AmazonController{Message: "usa spider running", SpiderType: "IP process is running"}
		err := core.ServePort(host, ac)
		if err != nil {
			panic(err.Error())
		}
	}()

	core.IPPool()
}
