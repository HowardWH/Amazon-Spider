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
package log

//打印的level等级

const (
	NOSET  = 0
	NO     = 110
	PANIC  = 120
	ERROR  = 130
	WARN   = 140
	NOTICE = 150
	LOG    = 160
	DEBUG  = 170
	ALL    = 180
)

var LogLevelMap map[string]int = map[string]int{
	"NO":     NO,
	"DEBUG":  DEBUG,
	"WARN":   WARN,
	"NOTICE": NOTICE,
	"LOG":    LOG,
	"ERROR":  ERROR,
	"PANIC":  PANIC,
	"ALL":    ALL,
}
var logLevelStringMap map[int]string = map[int]string{
	NO:     "NO",
	DEBUG:  "DEBUG",
	WARN:   "WARN",
	NOTICE: "NOTICE",
	LOG:    "LOG",
	ERROR:  "ERROR",
	PANIC:  "PANIC",
	ALL:    "ALL",
}
