package main

import (
	"testing"
	"github.com/hunterhug/AmazonBigSpider/public/core"
	"github.com/hunterhug/parrot/util"
	"fmt"
	"github.com/hunterhug/AmazonBigSpider"
)

func TestGood2(t *testing.T) {
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir+"/config/"+"usa_local_config.json", AmazonBigSpider.Dir+"/config/"+"usa_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir+"/config/"+"usa_config.json", AmazonBigSpider.Dir+"/config/"+"usa_log.json")
	}
	level := 1
	parentdir := core.MyConfig.Datadir + "/url/" + (util.IS(level - 1))
	fmt.Println(parentdir)
	files, e := util.ListDir(parentdir, ".md")
	fmt.Printf("%#v,%#v", files, e)

}

func TestGood(t *testing.T) {
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir+"/config/"+"usa_local_config.json", AmazonBigSpider.Dir+"/config/"+"usa_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir+"/config/"+"usa_config.json", AmazonBigSpider.Dir+"/config/"+"usa_log.json")
	}
	num = 1
	wait = 1
	Good(4)
}
