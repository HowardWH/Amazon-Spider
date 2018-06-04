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
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hunterhug/rabbit/conf"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = conf.AdminTemplate + "/" + "error.html"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = conf.AdminTemplate + "/" + "error.html"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = conf.AdminTemplate + "/" + "error.html"
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "server error 500"
	c.TplName = conf.AdminTemplate + "/" + "error.html"
}
