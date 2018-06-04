package home

import (
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/rabbit/lib"
	"github.com/hunterhug/rabbit/models/blog"
)

// not so good must rewrite!
func (this *MainController) Paper() {
	id := this.Ctx.Input.Param(":cid")
	paperid := this.Ctx.Input.Param(":id") //文章id

	patemp := new(blog.Paper)

	pid, errp := strconv.Atoi(paperid)
	if errp != nil {
		this.Rsp(false, "你要干嘛？")
	}
	patemp.Id = int64(pid)

	n, errp1 := patemp.Query().Count()
	if n == 0 || errp1 != nil {
		this.Rsp(false, "不存在文章。。。")
	}
	patemp.Read()
	patemp.View = patemp.View + 1
	patemp.Update()

	this.Data["paper"] = patemp

	types := 0
	err, category := GetCategory(id)
	if err != nil {
		err, category = GetAlbum(id)
		if err != nil {
			this.Rsp(false, "没有这个分类啊，哥哥")
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
	temp.Urlpath = "/" + category["Alias"].(string) + "/" + strconv.Itoa(int(patemp.Id))
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
		this.TplName = this.GetTemplate() + "/paper.html"
	} else {
		this.TplName = this.GetTemplate() + "/paper.html"
	}
}
