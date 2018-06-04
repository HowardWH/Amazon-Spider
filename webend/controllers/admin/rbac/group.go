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
package rbac

import (
	"github.com/hunterhug/rabbit/models/admin"
)

type GroupController struct {
	CommonController
}

func (this *GroupController) Index() {
	if this.IsAjax() {
		page, _ := this.GetInt64("page")
		page_size, _ := this.GetInt64("rows")
		sort := this.GetString("sort")
		order := this.GetString("order")
		if len(order) > 0 {
			if order == "desc" {
				sort = "-" + sort
			}
		} else {
			sort = "Id"
		}
		nodes, count := admin.GetGrouplist(page, page_size, sort)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &nodes}
		this.ServeJSON()
		return
	} else {
		this.Layout = this.GetTemplate() + "/public/layout.html"
		this.TplName = this.GetTemplate() + "/rbac/group.html"
	}

}
func (this *GroupController) AddGroup() {
	g := admin.Group{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := admin.AddGroup(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *GroupController) UpdateGroup() {
	g := admin.Group{}
	if err := this.ParseForm(&g); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := admin.UpdateGroup(&g)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *GroupController) DelGroup() {
	Id, _ := this.GetInt64("Id")
	status, err := admin.DelGroupById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
