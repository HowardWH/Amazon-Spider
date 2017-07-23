package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = beego.AppConfig.String("admin_template") + "/" + "error.html"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = beego.AppConfig.String("admin_template") + "/" + "error.html"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = beego.AppConfig.String("admin_template") + "/" + "error.html"
}
