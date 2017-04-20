package main

import (
	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/AmazonBigSpider/public/core"
)

func main() {
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir+"/config/uk_local_config.json", AmazonBigSpider.Dir+"/config/uk_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir+"/config/uk_config.json", AmazonBigSpider.Dir+"/config/uk_log.json")
	}
	core.InitDB()
}
