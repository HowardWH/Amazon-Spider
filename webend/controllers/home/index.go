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
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/rabbit/lib"
	"github.com/hunterhug/rabbit/models/blog"
)

type MainController struct {
	baseController
}

// global web config
var config *blog.Config

// Index: web info struct
type webInfo map[string]map[string]interface{}

// init in every request
func (this *MainController) Prepare() {
	this.baseController.Prepare()

	config = new(blog.Config)
	config.Id = 1
	config.Read()

	// global set
	this.Data["config"] = config

	// paper
	this.Data["category"] = GetNav(0, 0)

	// production
	this.Data["photo"] = GetNav(0, 1)
}

func (this *MainController) DetectIndex() {

	// change lang|home by cookie
	h := this.GetString("h", "")
	if h != "" {
		switch h {
		case "first", "default", "hunterhug":
			this.Ctx.SetCookie("X-Home", "home/"+h, false, "/", false, false, true)
		default:
			break
		}

	}

	lang := this.GetString("lang", "")
	if lang != "" {
		this.Ctx.SetCookie("lang", lang)
		switch lang {
		case "en":
			this.Lang = "en-US"
		case "cn":
			this.Lang = "zh-CN"
		default:
			this.Lang = "zh-CN"
		}
		this.Data["Lang"] = this.Lang
	}
}

// Home
func (this *MainController) Index() {

	// before index do some detect
	this.DetectIndex()

	// get roll
	roll := new(blog.Roll)
	rolls := []orm.Params{}
	roll.Query().Filter("Status", 1).OrderBy("-Sort", "Createtime").Values(&rolls)
	this.Data["roll"] = rolls

	// index blocks
	index := webInfo{}
	err := json.Unmarshal([]byte(lib.TripAll(config.Webinfo)), &index)
	if err != nil {
		beego.Trace(err.Error())
	}
	for i, item := range index {
		// one block meta and it's content
		e, block, blockc := GetBlocks(item["name"].(string), int(item["limit"].(float64)))
		if e != nil {
			beego.Error(e.Error())
		}
		this.Data["block"+i] = block
		this.Data["block"+i+"c"] = blockc
	}

	this.TplName = this.GetTemplate() + "/index.html"
}

// Get Nav
func GetNav(beautyid int, blogtype int) []orm.Params {
	category := new(blog.Category)
	categorys := []orm.Params{}
	// query：beautyid 1 level
	category.Query().Filter("Status", 1).Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", blogtype).OrderBy("-Sort", "Createtime").Values(&categorys, "Id", "Title", "Alias")
	for _, cate := range categorys {
		// query: 2 level
		son := []orm.Params{}
		category.Query().Filter("Pid", cate["Id"]).OrderBy("-Sort", "Createtime").Values(&son, "Id", "Title", "Alias")
		cate["Son"] = son
	}
	return categorys

}

func GetBlocks(alias string, count int) (error, []orm.Params, orm.Params) {
	err, category := GetCategory(alias)
	if err != nil {
		err, album := GetAlbum(alias)
		if err != nil {
			return errors.New("can't find block" + alias), []orm.Params{}, album
		} else {
			id := album["Id"].(int64)
			return nil, GetPhoto(id, count), album
		}
	} else {
		id := category["Id"].(int64)
		return nil, GetPaper(id, count), category
	}
}

//获取开启的文章，按置顶
func GetPaper(id int64, count int) []orm.Params {
	paper := new(blog.Paper)
	papers := []orm.Params{}
	paper.Query().Filter("Cid", id).Filter("Type", 0).Filter("Status", 1).OrderBy("-Istop", "-Createtime").Limit(count, 0).Values(&papers)
	return papers
}

//获取开启的图片，按轮转，置顶
func GetPhoto(id int64, count int) []orm.Params {
	paper := new(blog.Paper)
	papers := []orm.Params{}
	paper.Query().Filter("Cid", id).Filter("Type", 1).Filter("Status", 1).OrderBy("-Isroll", "-Istop", "-Createtime").Limit(count, 0).Values(&papers)
	return papers
}

func GetCategoryOrAlbum(alias string, id interface{}) (error, orm.Params) {
	category := new(blog.Category)
	categorys := []orm.Params{}
	query := category.Query().Filter("Siteid", 0).Filter("Alias", alias).Limit(1)
	if id != nil {
		query = query.Filter("Type", id)
	}
	query.Values(&categorys)
	if len(categorys) == 0 {
		return errors.New("can't get category:" + alias), orm.Params{}
	} else {
		return nil, categorys[0]
	}

}

//获取文章目录
func GetCategory(alias string) (error, orm.Params) {
	return GetCategoryOrAlbum(alias, 0)
}

//获取相册目录
func GetAlbum(alias string) (error, orm.Params) {
	return GetCategoryOrAlbum(alias, 1)
}
