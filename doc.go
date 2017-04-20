// a go spider!
package AmazonBigSpider

import (
	"github.com/hunterhug/GoSpider/util"
	"fmt"
)

var Dir = util.CurDir()
var Local = true

func init() {
	// 在根目录建一个远程.txt使用远程配置
	if util.FileExist(Dir + "/远程.txt") {
		Local = false
		fmt.Println("远程方式！！！")
	}
}
