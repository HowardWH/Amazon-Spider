/*
	Copyright 2017 by rabbit author: gdccmcm14@live.com.
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License
*/

// Main Web Entrance
package main

import (
	"flag"
	"mime"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/hunterhug/rabbit/conf"
	"github.com/hunterhug/rabbit/controllers"
	"github.com/hunterhug/rabbit/lib"
	"github.com/hunterhug/rabbit/models"
	"github.com/hunterhug/rabbit/routers"
)

func init() {
	// init flag
	flags := conf.FlagConfig{}

	// user that hide
	flags.User = flag.String("user", "", "user")

	// db init or rebuild
	flags.DbInit = flag.Bool("db", false, "init db")
	flags.DbInitForce = flag.Bool("f", false, "force init db first drop db then rebuild it")

	// rbac config rebuild
	flags.Rbac = flag.Bool("rbac", false, "rebuild rbac database tables")

	// front-end  view
	home := flag.String("home", "", "home template")

	// config file position
	config := flag.String("config", "", "config file position if empty use default")

	flag.Parse()

	// init config
	if *config != "" {
		beego.Trace("use diy config")
		err := beego.LoadAppConfig("ini", *config)
		if err != nil {
			beego.Trace(err.Error())
		} else {
			beego.Trace("Use config:" + *config)
		}
	}

	if *home != "" {
		beego.Trace("Home template is " + *home)
		beego.AppConfig.Set(beego.BConfig.RunMode+"::"+"home_template", *home)
	}

	conf.InitConfig()

	// init lang
	// just add some ini in conf such locale_zh-CN.ini and edit app.conf
	langTypes := strings.Split(beego.AppConfig.String("lang_types"), "|")

	for _, lang := range langTypes {
		beego.Trace("Load language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Load language error:", err)
			return
		}
	}

	// add func map
	beego.Trace("add i18n function map")
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Trace("add stringsToJson function  map")
	beego.AddFuncMap("stringsToJson", lib.StringsToJson)
	mime.AddExtensionType(".css", "text/css") // some not important

	// init model
	beego.Trace("model run")
	models.RunBaseDb(flags)
	// init amazon
	if conf.Amazon {
		//mkdir file
		//chmod 777 file
		//
		//mkdir file/back
		//mkdir file/data
		models.RunAmazonDb()
	}

	// init router
	beego.Trace("router run")
	routers.Run()
	if conf.Amazon {
		routers.RunAmazon()
	}

	beego.Trace("start open error template")
	beego.ErrorController(&controllers.ErrorController{})
}

// Start!
func main() {
	beego.Trace("Start Listen ...")
	beego.Run()
}
