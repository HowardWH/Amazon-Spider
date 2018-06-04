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
	spider "github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/marmot/expert"
	"github.com/hunterhug/parrot/util"
	"regexp"
	"testing"
)

func TestListDownload1(t *testing.T) {
	util.MakeDir(Dir + "/test/list/")
	ip := "*104.128.124.122:808"

	// debug info will no appear |nothing
	spider.SetLogLevel("info")

	url := "https://www.amazon.com/dp/B01C8RN7VG"
	content, err := Download(ip, url)
	if err != nil {
		fmt.Printf("%#v", err.Error())
	}
	util.SaveToFile(Dir+"/test/list/xxx2.html", content)
}

func TestParseRank1(t *testing.T) {
	bytecontent, _ := util.ReadfromFile(Dir + "/test/list/xxx2.html")
	fmt.Printf("%#v\n", Urlmap)
	doc, _ := expert.QueryBytes(bytecontent)
	test := doc.Find("body").Text()
	fmt.Printf("%#v\n", test)
	r, _ := regexp.Compile(`#([,\d]{1,10})[\s]{0,1}[A-Za-z0-9]{0,6} in ([^#;)(\n]{2,30})[\s\n]{0,1}[(]{0,1}`)
	god := r.FindAllStringSubmatch(test, -1)
	fmt.Printf("%#v\n", god)
}

func TestParsedd1(t *testing.T) {
	bytecontent, _ := util.ReadfromFile(Dir + "/test/list/xxx2.html")
	t.Logf("%#v", ParseDetail("/dp/dd", bytecontent))
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