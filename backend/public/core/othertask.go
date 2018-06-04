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
	"fmt"
	"github.com/hunterhug/parrot/util"
	"os"
)

func Montior() {
	for {
		urltotal, e1 := RedisClient.Llen(MyConfig.Urlpool)
		urldone, e2 := RedisClient.Hlen(MyConfig.Urlhashpool)
		asintotal, e3 := RedisClient.Llen(MyConfig.Asinpool)
		asindone, e4 := RedisClient.Hlen(MyConfig.Asinhashpool)
		ipremain, e5 := RedisClient.Llen(MyConfig.Proxypool)
		sql := "INSERT INTO `smart_monitor`(id,redistype,doing,done,createtime)VALUES(?,?,?,?,?)on duplicate key update doing=?,done=?,updatetime=?"
		fmt.Printf("e1:%v,e2:%v,e3:%v,e4:%v,e5:%v\n", e1, e2, e3, e4, e5)
		if e1 == nil && e2 == nil {
			_, err := BasicDb.Insert(sql, Today, "urlpool", urltotal, urldone, util.TodayString(6), urltotal, urldone, util.TodayString(6))
			if err == nil {
				fmt.Printf("dbur-%v,%v,%v,%v,%v\n", urltotal, urldone, asintotal, asindone, ipremain)
			} else {
				fmt.Println("urlpoolinsert:" + err.Error())
			}
		}
		if e3 == nil && e4 == nil {
			_, err := BasicDb.Insert(sql, Today, "asinpool", asintotal, asindone, util.TodayString(6), asintotal, asindone, util.TodayString(6))
			if err == nil {
				fmt.Printf("dbas-%v,%v,%v,%v,%v\n", urltotal, urldone, asintotal, asindone, ipremain)
			} else {
				fmt.Println("asinpoolinsert:" + err.Error())
			}
		}
		if e5 == nil {
			//_, err := BasicDb.Insert(sql, Today, "ippool", ipremain, 0, util.TodayString(6), ipremain, 0, util.TodayString(6))
			//if err == nil {
				fmt.Printf("dbip-%v,%v,%v,%v,%v\n", urltotal, urldone, asintotal, asindone, ipremain)
			//} else {
			//	fmt.Println("ippoolinsert:" + err.Error())
			//}
		}
		fmt.Println("----------------")
		util.Sleep(30)

	}

}

func Clean() {
	today, _ := util.SI(Today)
	for {

		Newday := util.TodayString(3)
		newday, _ := util.SI(Newday)
		if newday > today {
			fmt.Println("today out!,now is " + Newday)
			os.Exit(1)
		}
		util.Sleep(1800)
	}
}

func smart2016() string {
	urltotal, _ := RedisClient.Llen(MyConfig.Urlpool)
	urldone, _ := RedisClient.Hlen(MyConfig.Urlhashpool)
	asintotal, _ := RedisClient.Llen(MyConfig.Asinpool)
	asindone, _ := RedisClient.Hlen(MyConfig.Asinhashpool)
	ipremain, _ := RedisClient.Llen(MyConfig.Proxypool)
	return fmt.Sprintf(`
	<table border="1" style="text-align:center;font-size:2.0em">
	<tr>
	<th>URLPOOL</th><th>URLDONE</th><th>ASINPOOL</th><th>ASINDONE</th><th>IPPOOL</th>
	</tr>
	<tr>
	<td>%v</td>
	<td>%v</td>
	<td>%v</td>
	<td>%v</td>
	<td>%v</td>
	</tr>
	</table>
	`, urltotal, urldone, asintotal, asindone, ipremain)

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