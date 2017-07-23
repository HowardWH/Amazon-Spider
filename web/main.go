// 应用主函数包
package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/hunterhug/AmazonBigSpiderWeb/controllers"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	"github.com/hunterhug/AmazonBigSpiderWeb/models"
	"github.com/hunterhug/AmazonBigSpiderWeb/routers"
	"mime"
	"strings"
)

// 国际化语言数组
var langTypes []string

// 加载、初始化国际化
func init() {
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	for _, lang := range langTypes {
		beego.Trace("加载语言: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("加载语言文件失败:", err)
			return
		}
	}

	// 添加映射
	beego.Trace("添加i18n函数映射")
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Trace("添加json格式化函数映射")
	beego.AddFuncMap("stringsToJson", StringsToJson)
	mime.AddExtensionType(".css", "text/css")

	// 模型初始化
	beego.Trace("模型初始化")
	models.Run()

	beego.Trace("路由开始")
	routers.Run()

	beego.Trace("错误模板开启")
	beego.ErrorController(&controllers.ErrorController{})
}

func main() {
	beego.Trace("监听开始")
	beego.Run()
}
