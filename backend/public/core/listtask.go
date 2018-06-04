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
package core

import (
	"github.com/hunterhug/parrot/util"
	"math/rand"
)

var (
	listnum int
	endchan chan string
)

func listtask(taskname string) {
	second := rand.Intn(5)
	AmazonListLog.Debugf("%s:%d second", taskname, second)
	util.Sleep(second)
	if MyConfig.Proxycategory {
		err := GetUrls()
		if err != nil {
			AmazonListLog.Error(taskname + "-error:" + err.Error())
		}
	} else {
		err := GetNoneProxyUrls(taskname)
		if err != nil {
			AmazonListLog.Error(taskname + "-error:" + err.Error())
		}
	}
	endchan <- "done!"
}

func ListTask() {
	OpenMysql()
	err := CreateAsinTables()
	if err != nil {
		AmazonListLog.Errorf("createtables:%s,error:%s", Today, err.Error())
	}
	err = CreateAsinRankTables()
	if err != nil {
		AmazonListLog.Errorf("createtables:%s,error:%s", "Asin"+Today, err.Error())
	}
	listnum = MyConfig.Listtasknum
	endchan = make(chan string, listnum)
	for i := 0; i < listnum; i++ {
		go listtask("ltask" + util.IS(i))
	}
	go Clean()
	for i := 0; i < listnum; i++ {
		<-endchan
	}
	AmazonListLog.Log("List All done")
}
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