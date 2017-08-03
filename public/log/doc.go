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
// 日志记录类
package log

//用法：
//
// import "sunteng/commons/log"
// var jsconf = `
// {
// 	"Appenders":{
// 		"test_appender":{
// 			"Type":"file",
// 			"Target":"/tmp/test.log"
// 		},
// 		"a_appender":{
// 			"Type":"console"
// 		}
// 	},
// 	"Loggers":{
// 		"sunteng/commons/log/a":{
//          "@Appenders":"日志输出到test_appender和a_appender",
// 			"Appenders":["test_appender","a_appender"],
//          "@Level":"记录debug和debug等级以上的数据",
// 			"Level":"debug"
// 		},
// 		"sunteng/commons/log/b":{
//          "@Appenders":"日志输出到最近上级的appender,即Root的Appenders",
//          "@Level":"只记录debug和error等级的数据",
// 			"Level":["debug","error"]
// 		}
// 	},
// 	"Root":{
// 		"Level":"log",
// 		"Appenders":["test_appender"]
// 	}
// }
// `
// log.Init(jsconf)
// logger := log.Get("sunteng/commons/log/a")
// logger.Log("hello logger")
// logger := log.Get("sunteng/commons/log/a/b")
// logger.Log("hello logger")
