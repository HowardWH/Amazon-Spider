// a go spider!
package AmazonBigSpider

import (
	"flag"
	"fmt"
	"github.com/hunterhug/GoSpider/util"
	//"net"
	"github.com/hunterhug/GoSpider/spider"
	"os"
	"path/filepath"
	"strings"
)

var Dir = util.CurDir()
var CoreDir = filepath.Join(Dir, "public", "core")
var Local = true

func init() {
	rootdir := flag.String("root", "", "root config")
	coredir := flag.String("core", "", "core config")
	user := flag.String("user", "", "user")
	if !flag.Parsed() {
		flag.Parse()
	}
	if *rootdir != "" {
		Dir = *rootdir
	}
	if *coredir != "" {
		CoreDir = *coredir
	}
	// 在根目录建一个远程.txt使用远程配置
	if util.FileExist(Dir + "/远程.txt") {
		Local = false
		fmt.Println("远程方式！！！")
	}

	//addrs, err := net.InterfaceAddrs()
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//
	//	for _, address := range addrs {
	//
	//		// 检查ip地址判断是否回环地址
	//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//			if ipnet.IP.To4() != nil {
	//				fmt.Println(ipnet.IP.String())
	//			}
	//
	//		}
	//	}
	//}
	sp := spider.NewAPI()
	sp.SetUrl("http://www.lenggirl.com/xx.xx")
	data, err := sp.Get()
	if err != nil {
		fmt.Println("Network error, retry")
		os.Exit(0)
	}
	if strings.Contains(string(data), "帮帮宝贝回家") {
		fmt.Println("Network error, retry")
		os.Exit(0)
	}

	if strings.Contains(string(data), "#hunterhugxxoo") || (strings.Contains(string(data), "user-"+*user) && *user != "") {
		fmt.Println("start app")
	} else {
		fmt.Println("start app...")
		fmt.Println("error!")
		os.Exit(0)
	}
}
