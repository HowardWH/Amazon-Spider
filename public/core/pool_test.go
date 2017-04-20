package core

import (
	"fmt"
	"github.com/hunterhug/AmazonBigSpider"
	"testing"
)

func TestGetIP(t *testing.T) {
	if AmazonBigSpider.Local {
		InitConfig(AmazonBigSpider.Dir+"/config/usa_local_config.json", AmazonBigSpider.Dir+"/config/usa_log.json")
	} else {
		InitConfig(AmazonBigSpider.Dir+"/config/usa_config.json", AmazonBigSpider.Dir+"/config/usa_log.json")
	}
	fmt.Printf("%#v", getips())
}
