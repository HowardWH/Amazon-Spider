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

import (
	"runtime"
	"strings"
)

var defaultManager *LoggerManager

func init() {
	var defaultJson = `
		{
		    "Appenders":{
		        "stdout":{
		            "Type":"console"
		        }
		    },
		    "Root":{
		        "Level":"DEBUG",
		        "Appenders":["stdout"]
		    }
		}
		`
	var err error
	defaultManager, err = NewLoggerManagerWithJsconf(defaultJson)
	if err != nil {
		panic(err)
	}
}

func Init(jsconf string) (err error) {
	cfg, err := LoadConf(jsconf)
	if err != nil {
		return err
	}
	return defaultManager.UpdateConf(cfg)
}
func InitConf(cfg *Config) (err error) {
	return defaultManager.UpdateConf(cfg)
}

func CurLoggerMananger() (cfg *LoggerManager) {
	return defaultManager
}

func Get(name string) *Logger {
	return defaultManager.Logger(name)
}

func UseRoot(name string) error {
	return defaultManager.UseRoot(name)
}
func SetRootAppender(appenders ...Appender) {
	defaultManager.SetRootAppender(appenders...)
}

func SetRootSeparationAppender(fileName string) {
	SetRootAppender(NewLevelSeparationDailyAppender("root", fileName))
}
func SetRootFileAppender(fileName string) {
	SetRootAppender(NewFileAppender("root", fileName))
}

func SetRootLevel(l int)         { defaultManager.SetRootLevel(l) }
func SetRootOnlyLevel(ls ...int) { defaultManager.SetRootOnlyLevel(ls...) }

func Debug(args ...interface{})  { defaultLogger().Debug(args...) }
func Log(args ...interface{})    { defaultLogger().Log(args...) }
func Notice(args ...interface{}) { defaultLogger().Notice(args...) }
func Warn(args ...interface{})   { defaultLogger().Warn(args...) }
func Error(args ...interface{})  { defaultLogger().Error(args...) }

func Debugf(format string, args ...interface{})  { defaultLogger().Debugf(format, args...) }
func Logf(format string, args ...interface{})    { defaultLogger().Logf(format, args...) }
func Noticef(format string, args ...interface{}) { defaultLogger().Noticef(format, args...) }
func Warnf(format string, args ...interface{})   { defaultLogger().Warnf(format, args...) }
func Errorf(format string, args ...interface{})  { defaultLogger().Errorf(format, args...) }

func IsAll() bool    { return defaultLogger().IsAll() }
func IsInfo() bool   { return defaultLogger().IsInfo() }
func IsDebug() bool  { return defaultLogger().IsDebug() }
func IsNotice() bool { return defaultLogger().IsNotice() }
func IsWarn() bool   { return defaultLogger().IsWarn() }
func IsError() bool  { return defaultLogger().IsError() }

func defaultLogger() (logger *Logger) {
	name := pathInGoPath(2)
	logger = defaultManager.Logger(name)
	logger.SetCallpath(DefaultLoggerCallpath + 1)
	return
}
func pathInGoPath(level int) (inGoPath string) {
	_, name, _, _ := runtime.Caller(level + 1)
	if arr := strings.Split(name, "src/"); len(arr) > 1 {
		inGoPath = arr[1]
	} else {
		inGoPath = strings.Trim(name, "/")
	}
	return
}
