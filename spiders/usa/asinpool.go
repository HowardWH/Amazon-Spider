package main

import (
	"github.com/hunterhug/AmazonBigSpider/public/core"
	"github.com/hunterhug/AmazonBigSpider"
)

func main() {
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir+"/config/usa_local_config.json", AmazonBigSpider.Dir+"/config/usa_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir+"/config/usa_config.json", AmazonBigSpider.Dir+"/config/usa_log.json")
	}
	core.AsinPool()
}
