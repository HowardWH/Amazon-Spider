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
	"strings"
)

func CollectAsin(files []string) {
	// you can open in many place
	OpenMysql()
	for _, file := range files {
		AmazonIpLog.Debugf("deal file:%s", file)
		text, err := util.ReadfromFile(file)
		if err != nil {
			AmazonIpLog.Errorf("Read %s-error:%s", file, err.Error())
			continue
		}
		fileinfo, _ := util.GetFilenameInfo(file)
		createtime := util.GetSecond2DateTimes(fileinfo.ModTime().Unix())
		insertlist, err := ParseList(text)
		if err != nil {
			AmazonIpLog.Errorf("Parse %s-error:%s", file, err.Error())
			continue
		}
		category := strings.Split(fileinfo.Name(), ",")
		if len(category) != 2 {
			AmazonIpLog.Errorf("Filename %s:error", file)
			continue
		}
		category1 := strings.Split(category[1], ".")[0]
		err = InsertAsinMysql(insertlist, createtime, category1)
		if err != nil {
			AmazonIpLog.Errorf("InsertMysql %s-error:%s", file, err.Error())
			continue
		}
		err = util.Rename(file, file+"sql")

		//err = os.Remove(file)
		if err != nil {
			AmazonIpLog.Errorf("Rename %s-error:%s", file, err.Error())
		} else {
			AmazonIpLog.Debugf("Rename %s", file+"sql")
		}
	}
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