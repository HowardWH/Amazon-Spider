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
	"fmt"
	"github.com/hunterhug/rabbit/models/admin"
)

type UserController struct {
	CommonController
}

func (this *UserController) Index() {
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
		users, count := admin.Getuserlist(page, page_size, sort)

		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		this.ServeJSON()
		return
	} else {
		this.Layout = this.GetTemplate() + "/public/layout.html"
		this.TplName = this.GetTemplate() + "/rbac/user.html"
	}

}

func (this *UserController) AddUser() {
	u := admin.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := admin.AddUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, fmt.Sprintf("%v", err))
		return
	}

}

func (this *UserController) UpdateUser() {
	u := admin.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := admin.UpdateUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, fmt.Sprintf("%v", err))
		return
	}

}

//UpdateUserPasswd
func (this *UserController) UpdateUserPasswd() {
	u := admin.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := admin.UpdateUserPasswd(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, fmt.Sprintf("%v", err))
		return
	}

}

func (this *UserController) DelUser() {
	/*	 Id, _ := this.GetInt64("Id")
		 status, err := admin.DelUserById(Id)
		 if err == nil && status > 0 {
		 	this.Rsp(true, "Success")
		 	return
		 } else {
		 	this.Rsp(false, err.Error())
		 	return
		 }*/

	this.Rsp(false, "本系统用户一添加则不允许删除")
	return
}
