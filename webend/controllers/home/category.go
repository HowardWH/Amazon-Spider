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
package home

import (
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/rabbit/lib"
	"github.com/hunterhug/rabbit/models/blog"
)

func (this *MainController) Category() {
	id := this.Ctx.Input.Param(":id")
	types := 0
	err, category := GetCategory(id)
	if err != nil {
		err, category = GetAlbum(id)
		if err != nil {
			this.Rsp(false, "not this category")
		} else {
			types = 1
		}
	}

	//本大爷
	this.Data["thiscategory"] = category

	//附录爸爸
	cid := category["Pid"]
	father := new(blog.Category)
	father.Id = cid.(int64)
	err1 := father.Read()
	if err1 != nil {

	} else {
		this.Data["father"] = father
	}

	//附录儿子
	son := []orm.Params{}
	new(blog.Category).Query().Filter("Pid", category["Id"].(int64)).Filter("Status", 1).OrderBy("-Sort", "Createtime").Values(&son, "Title")
	this.Data["son"] = son

	//文章儿子
	var limit int64 = 8
	papers := []orm.Params{}
	query1 := new(blog.Paper).Query().Filter("Cid", category["Id"].(int64)).Filter("Status", 1).OrderBy("-Istop", "Createtime")
	page, _ := this.GetInt64("page", 1)
	if page <= 0 {
		page = 1
	}
	//总数
	count, _ := query1.Count()

	temp := new(lib.Pager)
	temp.Page = page
	temp.Pagesize = limit
	temp.Totalnum = count
	temp.Urlpath = "/" + category["Alias"].(string)
	//beego.Trace("Dddd"+temp.ToString())
	this.Data["nums"] = temp.ToString()
	query1.Limit(limit, (page-1)*limit).Values(&papers)

	this.Data["papers"] = papers

	////图片轮转
	//roll := new(blog.Roll)
	//rolls := []orm.Params{}
	//roll.Query().Filter("Status", 1).OrderBy("-Sort", "Createtime").Values(&rolls)
	//this.Data["roll"] = rolls

	if types == 0 {
		this.TplName = this.GetTemplate() + "/category.html"
	} else {
		this.TplName = this.GetTemplate() + "/album.html"
	}
}
