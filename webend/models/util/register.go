package util

import (
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/rabbit/models/admin"
	"github.com/hunterhug/rabbit/models/blog"
)

func RegisterDBAdminModel() {
	orm.RegisterModel(new(admin.Group), new(admin.Node), new(admin.Role), new(admin.User))
}

func RegisterDBBlogModel() {
	orm.RegisterModel(new(blog.Category), new(blog.Config), new(blog.Paper), new(blog.Roll))
}

func RegisterDBModel() {
	RegisterDBAdminModel()
	RegisterDBBlogModel()
}
