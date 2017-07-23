/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:569929309

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:569929309

	2017.7 by hunterhug
*/
package core

import (
	"fmt"
	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/util"
	"testing"
)

// https://www.amazon.com/Best-Sellers-Automotive-Performance-ABS-Brake-Parts/zgbs/automotive/15710931/ref=zg_bs_pg_1?_encoding=UTF8&pg=1&ajax=1
// https://www.amazon.com/Best-Sellers-Automotive-Performance-ABS-Brake-Parts/zgbs/automotive/15710931/ref=zg_bs_pg_2?_encoding=UTF8&pg=2&ajax=1
// https://www.amazon.com/Best-Sellers-Automotive-Performance-ABS-Brake-Parts/zgbs/automotive/15710931/ref=zg_bs_pg_3?_encoding=UTF8&pg=3&ajax=1
// https://www.amazon.com/Best-Sellers-Automotive-Performance-ABS-Brake-Parts/zgbs/automotive/15710931/ref=zg_bs_pg_4?_encoding=UTF8&pg=4&ajax=1
// https://www.amazon.com/Best-Sellers-Automotive-Performance-ABS-Brake-Parts/zgbs/automotive/15710931/ref=zg_bs_pg_5?_encoding=UTF8&pg=5&ajax=1
// https://www.amazon.com/dp/B001IHBLPC
func TestAsinDownload(t *testing.T) {
	util.MakeDir(AmazonBigSpider.Dir + "/test/asin/")
	ip := "104.128.124.122:808"
	// debug info will no appear |nothing
	spider.SetLogLevel("info")
	url := "https://www.amazon.com/dp/B016L36UZI"
	prefix := "asin"
	testtimes := 1000
	for {
		testtimes--
		if testtimes == 0 {
			break
		}
		robotime := 0
		maxtime := 1000
		times := 0
		for {
			if times > maxtime {
				break
			}
			temp := url
			content, err := Download(ip, temp)
			if err != nil {
				fmt.Printf("%#v", err.Error())
			} else {
				err = spider.TooSortSizes(content, 10)
				// robot continue
				if err != nil {
					robotime++
					times++
					break
				} else {
					// and then out
					fmt.Printf("The %d try Asin page :%d times | robbot max times:%d\n", testtimes, times, robotime)
				}
				util.SaveToFile(AmazonBigSpider.Dir+"/test/asin/"+prefix+util.IS(testtimes)+".html", content)
				break
			}
		}
	}
}

func TestListDownload(t *testing.T) {
	util.MakeDir(AmazonBigSpider.Dir + "/test/list/")
	ip := "104.128.124.122:808"
	// debug info will no appear |nothing
	spider.SetLogLevel("info")
	url := "https://www.amazon.co.jp/gp/bestsellers/dvd/ref=zg_bs_nav_0"
	content, err := Download(ip, url)
	if err != nil {
		fmt.Printf("%#v", err.Error())
	}
	util.SaveToFile(AmazonBigSpider.Dir+"/test/list/xxx2.html", content)
}
/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:569929309

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:569929309

	2017.7 by hunterhug
*/