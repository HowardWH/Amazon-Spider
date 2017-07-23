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
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/AmazonBigSpider/public/core"
	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/util"
	"strings"
)

var urlchan chan string
var proxy bool = false
var num int = 20
var wait int = 0

func main() {
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir+"/config/"+"de_local_config.json", AmazonBigSpider.Dir+"/config/"+"de_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir+"/config/"+"de_config.json", AmazonBigSpider.Dir+"/config/"+"de_log.json")
	}
	//6级别
	//26-28-14-4-10-0,https://www.amazon.co.jp/gp/bestsellers/books/3525971,ヴェルディ
	util.MakeDir(core.MyConfig.Datadir + "/url/0")
	util.MakeDir(core.MyConfig.Datadir + "/url/1")
	util.MakeDir(core.MyConfig.Datadir + "/url/2")
	util.MakeDir(core.MyConfig.Datadir + "/url/3")
	util.MakeDir(core.MyConfig.Datadir + "/url/4")
	listlevel1 := index() //1
	level0(listlevel1)    //2
	//Good(1) //3
	//Good(2)     //4
	//Good(3) //5
	//Good(4) //6
}

// so ! what !
func Good(level int) {
	num = 20
	wait = 0
	urlchan = make(chan string, 1)
	parentdir := core.MyConfig.Datadir + "/url/" + (util.IS(level - 1))
	dir := core.MyConfig.Datadir + "/url/" + (util.IS(level))
	files, _ := util.ListDir(parentdir, ".md")
	//core.LocalLogger.Debugf("%#v",files)
	process, e := util.DevideStringList(files, num)
	if e != nil {
		panic(e.Error())
	}
	for index, v := range process {
		go func(v []string) {
			ip := "*" + util.IS(index)
			ipbegintimes := ""
			if proxy {
				ip = core.GetIP()

				// before use, send to hash pool
				ipbegintimes = util.GetSecord2DateTimes(util.GetSecordTimes())
				core.RedisClient.Hset(core.MyConfig.Proxyhashpool, ip, ipbegintimes)
			}
			for _, file := range v {

				tempbytecont, _ := util.ReadfromFile(file)
				tempurls := strings.Split(string(tempbytecont), "\n")
				//core.LocalLogger.Debugf("%#v",tempurls)
				for _, url := range tempurls {
					tempurl := strings.Split(url, ",")
					num := tempurl[0]
					a := dir + "/" + num + ".html"
					b := dir + "/" + num + ".md"
					c := dir + "/" + num + ".mdxx"
					fileok := util.FileExist(a)
					filemdok := util.FileExist(b)
					filemdokxx := util.FileExist(c)
					reallyurl := tempurl[1]
					if strings.Contains(reallyurl, "books") {
						fmt.Printf("%s,%s是书籍，重复！！\n", num, reallyurl)
						continue
					}
					if filemdokxx {
						fmt.Printf("%s,%s没有下级\n", num, reallyurl)
						continue
					}
					if filemdok && fileok {
						fmt.Printf("%s,%s已经抓过和处理过\n", num, reallyurl)
						continue
					}
					urlcont := []byte("")
					var e error = nil
					if fileok {
						fmt.Printf("%s,%s已经抓过\n", num, reallyurl)
						urlcont, e = util.ReadfromFile(a)
						if e != nil {
							fmt.Printf("%s,%s打开失败:%s\n", num, reallyurl, e.Error())
						}
					} else {
						for {
							urlcont, e = core.Download(ip, reallyurl)
							util.Sleep(wait)
							if e != nil {
								fmt.Printf("%s,%s抓取失败:%s\n", num, reallyurl, e.Error())
								continue
							}
							spider, ok := core.Spiders.Get(ip)
							if robot(urlcont) {
								fmt.Printf("%s,%s抓取机器人！！！\n", num, reallyurl)
								if ok {
									spider.Errortimes = spider.Errortimes + 1
								}
							} else {
								break
							}
							// if proxy ip err more than config, change ip
							if proxy && ok && spider.Errortimes > core.MyConfig.Proxymaxtrytimes {
								// die sent
								ipendtimes := util.GetSecord2DateTimes(util.GetSecordTimes())
								insidetemp := ipbegintimes + "|" + ipendtimes + "|" + util.IS(spider.Fetchtimes-spider.Errortimes) + "|" + util.IS(spider.Errortimes)
								core.RedisClient.Hset(core.MyConfig.Proxyhashpool, ip, insidetemp)
								// you know it
								core.Spiders.Delete(ip)
								// get new proxy again
								ip = core.GetIP()
								ipbegintimes = util.GetSecord2DateTimes(util.GetSecordTimes())
								core.RedisClient.Hset(core.MyConfig.Proxyhashpool, ip, ipbegintimes)
							}
						}
						util.SaveToFile(a, urlcont)
					}
					if filemdok == false {
						md := parseurl(num, urlcont, level+2)
						if len(md) == 0 {
							util.SaveToFile(c, []byte(""))
						} else {
							util.SaveToFile(b, []byte(strings.Join(md, "\n")))
						}
					}
				}

			}
			urlchan <- "done"
		}(v)
	}
	for i := 0; i < num; i++ {
		<-urlchan
	}
	fmt.Println("Done!")
}

func robot(b []byte) bool {
	doc, e := query.QueryBytes(b)
	if e == nil {
		if strings.Contains(doc.Find("title").Text(), "Robot Check") {
			return true
		}
	} else {
		return true
	}
	return false

}
func level0(listlevel []string) {
	for _, v := range listlevel {
		bytescontents := []byte("")
		var err error = nil
		temp := strings.Split(v, ",")
		filename := temp[0]
		url := temp[1]
		file := core.MyConfig.Datadir + "/url/0/" + filename + ".html"
		filemd := core.MyConfig.Datadir + "/url/0/" + filename + ".md"
		fileok := util.FileExist(file)
		filemdok := util.FileExist(filemd)
		filemdokxx := util.FileExist(filemd + "xx")
		if filemdokxx {
			fmt.Printf("%s,%s没有下级\n", filename, url)
			continue
		}
		if fileok && filemdok {
			fmt.Printf("%s,%s已经抓过和处理过\n", filename, url)
			continue
		}
		if fileok == false {
			for {
				bytescontents, err = core.NonProxyDownload("*level1", url)

				if err != nil {
					fmt.Printf("%s,%s抓取失败:%s\n", filename, url, err.Error())
					continue
				} else {

					if robot(bytescontents) {
						fmt.Printf("%s,%s机器人\n", filename, url)
						continue
					}
					fmt.Printf("%s,%s抓取成功\n", filename, url)
					util.SaveToFile(file, bytescontents)
					break
				}
			}

		} else {
			bytescontents, err = util.ReadfromFile(file)
			if err != nil {
				fmt.Printf("%s,%s文件打开失败:%s\n", filename, url, err.Error())
				continue
			}
			fmt.Printf("%s.%s已经抓过\n", filename, url)
		}
		if filemdok == false {
			md := parseurl(filename, bytescontents, 1)
			if len(md) == 0 {
				util.SaveToFile(filemd+"xx", []byte(""))
			} else {
				util.SaveToFile(filemd, []byte(strings.Join(md, "\n")))
			}
		}

	}
}
func index() []string {
	index := core.MyConfig.Datadir + "/url/index.html"
	indexmd := core.MyConfig.Datadir + "/url/index.md"
	indexok := util.FileExist(index)
	indexmdok := util.FileExist(indexmd)
	url := core.URL + "/gp/bestsellers"
	ip := "*task1"
	if indexmdok && indexok {
		fmt.Println("首页抓取和处理成功")
		temp, err := util.ReadfromFile(indexmd)
		if err != nil {
			fmt.Println(err.Error())
		}
		return strings.Split(string(temp), "\n")
	} else {
		bytescontents := []byte("")
		var err error = nil
		if indexok {
			bytescontents, err = util.ReadfromFile(index)
		} else {
			for {
				bytescontents, err = core.NonProxyDownload(ip, url)
				if err != nil {
					continue
				}
				if robot(bytescontents) {
					continue
				}
				break
			}
		}
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if indexok == false {
				util.SaveToFile(index, bytescontents)
			}
			//zg_browseRoot
			returnlist := []string{}
			doc, _ := query.QueryBytes(bytescontents)
			root := doc.Find("#zg_browseRoot")
			i := 1
			root.Find("li").Each(func(num int, node *goquery.Selection) {
				a := node.Find("a")
				title := a.Text()
				href, e := a.Attr("href")
				if e || href != "" {
					returnlist = append(returnlist, util.IS(i)+","+strings.Split(href, "/ref")[0]+","+strings.Replace(title, ",", "", -1))
					i++
				}
			})
			//
			//for _, v := range returnlist {
			//	fmt.Printf("提取%v\n", v)
			//}
			util.SaveToFile(indexmd, []byte(strings.Join(returnlist, "\n")))
			return returnlist
		}
	}
	return []string{}
}

func parseurl(pfilename string, bytescontents []byte, level int) []string {
	doc, _ := query.QueryBytes(bytescontents)
	returnlist := []string{}
	mark := "#zg_browseRoot"
	for i := 0; i < level; i++ {
		mark = mark + " ul"
	}
	root := doc.Find(mark)
	i := 1
	root.Find("li").Each(func(num int, node *goquery.Selection) {
		a := node.Find("a")
		title := a.Text()
		href, e := a.Attr("href")
		if e || href != "" {
			returnlist = append(returnlist, pfilename+"-"+util.IS(i)+","+strings.Split(href, "/ref")[0]+","+strings.Replace(title, ",", "", -1))
			i++
		}
	})
	if len(returnlist) == 0 {
		fmt.Println("链接最底部！")
	} else {
		for _, v := range returnlist {
			fmt.Println(v)
		}
	}
	return returnlist
}
