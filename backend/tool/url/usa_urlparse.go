/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
*/
package main

// insert url into mysql
import (
	"fmt"
	"strings"

	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/AmazonBigSpider/public/core"
	"github.com/hunterhug/parrot/util"
)

func main() {
	if AmazonBigSpider.Local {
		core.InitConfig(AmazonBigSpider.Dir+"/config/"+"usa_local_config.json", AmazonBigSpider.Dir+"/config/"+"usa_log.json")
	} else {
		core.InitConfig(AmazonBigSpider.Dir+"/config/"+"usa_config.json", AmazonBigSpider.Dir+"/config/"+"usa_log.json")
	}
	core.OpenMysql()
	dir := core.MyConfig.Datadir + "/url"
	files, e := util.WalkDir(dir, "md")
	filesxx, exx := util.WalkDir(dir, "mdxx")
	if exx != nil {
		fmt.Println(exx.Error())
		panic("dudu")
	}
	if e != nil {
		fmt.Println(e.Error())
		panic("dudu")
	} else {
		ismallbool := map[string]bool{}
		for _, v := range filesxx {
			xxtemp := strings.Split(v, "\\")
			xxlen := len(xxtemp)
			ismallbool[xxtemp[xxlen-1]] = true

		}
		for _, file := range files {
			fmt.Printf("处理%s\n", file)
			dudu, dudue := util.ReadfromFile(file)
			if dudue != nil {
				fmt.Printf("打开%s失败\n")
			} else {
				filecont := string(dudu)
				filelist := strings.Split(filecont, "\n")
				for _, onefile := range filelist {
					mysqllist := strings.Split(onefile, ",")
					if len(mysqllist) != 3 {
						continue
					}
					temp := strings.Split(mysqllist[0], "-")
					pid := "0"
					ismall := 0
					level := len(temp)
					if level > 1 {
						pid = strings.Join(temp[0:len(temp)-1], "-")
					}
					if level == 6 {
						ismall = 1
					} else {
						xx := mysqllist[0] + ".mdxx"
						if _, ok := ismallbool[xx]; ok {
							ismall = 1
						}
					}
					bigpid := temp[0]
					bigname, ok := core.Urlnumdudumap[bigpid]
					if ok == false {
						continue
					}
					url := mysqllist[1]
					name := mysqllist[2]
					//fmt.Printf("%d,%s,%s,%s,%s\n",level,bigpid,bigname,url,name)
					// url must set as unique
					//Todo robot!!!!!and url repeat
					sql := "INSERT IGNORE INTO `smart_category`(`id`,`url`,`name`,`level`,`pid`,`createtime`,`bigpname`,`bigpid`,`ismall`) VALUES(?,?,?,?,?,?,?,?,?);"
					_, mysqle := core.BasicDb.Insert(sql, mysqllist[0], url, name, level, pid, util.TodayString(6), bigname, bigpid, ismall)
					if mysqle != nil {
						fmt.Printf("插入错误:%s\n", mysqle.Error())
					} else {
						fmt.Printf("插入成功:%s\n", onefile)
					}
				}
			}
		}
	}
	//fmt.Printf("%#v",core.Urlnumdudumap)
}
