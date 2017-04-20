package main

import (
	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/AmazonBigSpider/public/core"
)

func main() {
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir+"/config/de_local_config.json", AmazonBigSpider.Dir+"/config/de_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir+"/config/de_config.json", AmazonBigSpider.Dir+"/config/de_log.json")
	}
	core.LocalListParseTask()

}
