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
	"github.com/hunterhug/AmazonBigSpider"
	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/util"
	"testing"
)

func TestRobot(t *testing.T) {
	bytecontent, _ := util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/404.html")
	t.Log(IsRobot(bytecontent))
	bytecontent, _ = util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/categorynotexist.html")
	t.Log(IsRobot(bytecontent))
	bytecontent, _ = util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/listnull.html")
	t.Log(IsRobot(bytecontent))
	bytecontent, _ = util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/listnormal.html")
	t.Log(IsRobot(bytecontent))
	bytecontent, _ = util.ReadfromFile(AmazonBigSpider.Dir + "/test/robot/robot.html")
	t.Log(IsRobot(bytecontent))
}

func Test404(t *testing.T) {
	bytecontent, _ := util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/404.html")
	t.Log(Is404(bytecontent))
	bytecontent, _ = util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/categorynotexist.html")
	t.Log(Is404(bytecontent))
	bytecontent, _ = util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/listnull.html")
	t.Log(Is404(bytecontent))
	bytecontent, _ = util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/listnormal.html")
	t.Log(Is404(bytecontent))
	bytecontent, _ = util.ReadfromFile(AmazonBigSpider.Dir + "/test/robot/robot.html")
	t.Log(Is404(bytecontent))
}

func TestParselist(t *testing.T) {
	bytecontent, _ := util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/1,18-2-5-1-10.html")
	results, err := ParseList(bytecontent)
	for _, result := range results {
		t.Logf("%v:%v", result, err)
	}
}

func TestParseRank(t *testing.T) {
	bytecontent, _ := util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/xxx2.html")
	doc, _ := query.QueryBytes(bytecontent)
	test := doc.Find("body").Text()
	fmt.Printf("%#v\n", test)
	t.Logf("%#v", ParseRank(test))
}

func TestParsedd(t *testing.T) {
	bytecontent, _ := util.ReadfromFile(AmazonBigSpider.Dir + "/test/list/xxx2.html")
	t.Logf("%#v", ParseDetail("/dp/dd", bytecontent))
}

func TestManyRank(t *testing.T) {
	files, _ := util.ListDir(DataDir+"/asin/20161114", "html")
	for _, file := range files {
		fmt.Printf("%s\n", file)
		bytecontent, _ := util.ReadfromFile(file)
		doc, _ := query.QueryBytes(bytecontent)
		test := doc.Find("body").Text()
		//fmt.Printf("%#v\n", test)
		fmt.Printf("%#v\n", ParseRank(test))
	}
}

func TestParserankk(t *testing.T) {
	fmt.Printf("%#v", ParseRank("#1 in Computers & Accessories > Computer Accessories > Computer Cable Adapters > Serial Adapters "))
}

func TestBig(t *testing.T) {
	fmt.Println(BigReallyName("artscr_afts"))
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