package main

import (
	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/AmazonBigSpider/public/core"
)

func main() {
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir+"/config/jp_local_config.json", AmazonBigSpider.Dir+"/config/jp_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir+"/config/jp_config.json", AmazonBigSpider.Dir+"/config/jp_log.json")
	}
	core.UrlPool()
}
