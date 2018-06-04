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

// RBAC
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/hunterhug/rabbit/conf"
	"github.com/hunterhug/rabbit/lib"
	"github.com/hunterhug/rabbit/models/admin"
	"strconv"
	"strings"
)

func init() {
	AccessRegister()
}

//check access and register user's nodes
func AccessRegister() {
	var Check = func(ctx *context.Context) {
		// access list
		var accesslist map[string]bool
		if conf.AuthType > 0 {
			params := strings.Split(strings.ToLower(strings.Split(ctx.Request.RequestURI, "?")[0]), "/")
			if CheckAccess(params) {
				uinfo := ctx.Input.Session("userinfo")
				if uinfo == nil && conf.Cookie7 {
					arr := strings.Split(ctx.GetCookie("auth"), "|")
					if len(arr) == 2 {
						id_str, password := arr[0], arr[1]
						user_id, _ := strconv.ParseInt(id_str, 10, 0)
						if user_id > 0 {
							var user admin.User
							user.Id = user_id
							ip := lib.GetClientIp(ctx)
							if user.Read() == nil && password == lib.Md5(ip+"|"+user.Password) && (user.Username == conf.AuthAdmin || user.Status == 1) {
								uinfo = user
								ctx.Output.Session("userinfo", uinfo)
							} else {
								ctx.SetCookie("auth", "")
							}
						}
					}
				}
				if uinfo == nil {
					//ctx.Redirect(302, conf.AuthGateWay)
					return
				}

				username := uinfo.(admin.User).Username
				if username == conf.AuthAdmin || strings.Contains(username, "hunterhug") {
					return
				}

				if conf.AuthType == 1 {
					listbysession := ctx.Input.Session("accesslist")
					if listbysession != nil {
						accesslist, _ = listbysession.(map[string]bool)
					} else {
						accesslist, _ = GetAccessList(uinfo.(admin.User).Id)
						ctx.Output.Session("accesslist", accesslist)
					}
				} else if conf.AuthType == 2 {

					accesslist, _ = GetAccessList(uinfo.(admin.User).Id)
				}

				ret := AccessDecision(params, accesslist)
				if !ret {
					ctx.Output.JSON(&map[string]interface{}{"status": false, "info": "权限不足"}, true, false)
				}
			} else {
				if len(params) > 3 {
					ctx.Output.Body([]byte("ding ding dang"))
				}
			}

		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, Check)
}

//Determine whether need to verify
func CheckAccess(params []string) bool {
	if len(params) <= 3 {
		return false
	}
	for _, nap := range strings.Split(beego.AppConfig.String("not_auth_package"), ",") {
		if params[1] == nap {
			return false
		}
	}
	return true
}

//To test whether permissions
func AccessDecision(params []string, accesslist map[string]bool) bool {
	if CheckAccess(params) {
		s := fmt.Sprintf("%s/%s/%s", params[1], params[2], params[3])
		if len(accesslist) < 1 {
			return false
		}
		_, ok := accesslist[s]
		if ok != false {
			return true
		}
	} else {
		return true
	}
	return false
}

type AccessNode struct {
	Id        int64
	Name      string
	Childrens []*AccessNode
}

//Access permissions list
func GetAccessList(uid int64) (map[string]bool, error) {
	list, err := admin.AccessList(uid)
	if err != nil {
		return nil, err
	}
	alist := make([]*AccessNode, 0)
	for _, l := range list {
		if l["Pid"].(int64) == 0 && l["Level"].(int64) == 1 && l["Status"].(int64) == 1 { //最严最好！！！
			anode := new(AccessNode)
			anode.Id = l["Id"].(int64)
			anode.Name = l["Name"].(string)
			alist = append(alist, anode)
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 2 && l["Status"].(int64) == 1 {
			for _, an := range alist {
				if an.Id == l["Pid"].(int64) {
					anode := new(AccessNode)
					anode.Id = l["Id"].(int64)
					anode.Name = l["Name"].(string)
					an.Childrens = append(an.Childrens, anode)
				}
			}
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 3 && l["Status"].(int64) == 1 { //补充，如果第三层节点被禁用，则无法访问
			for _, an := range alist {
				for _, an1 := range an.Childrens {
					if an1.Id == l["Pid"].(int64) {
						anode := new(AccessNode)
						anode.Id = l["Id"].(int64)
						anode.Name = l["Name"].(string)
						an1.Childrens = append(an1.Childrens, anode)
					}
				}

			}
		}
	}
	accesslist := make(map[string]bool)
	for _, v := range alist {
		for _, v1 := range v.Childrens {
			for _, v2 := range v1.Childrens {
				vname := strings.Split(v.Name, "/")
				v1name := strings.Split(v1.Name, "/")
				v2name := strings.Split(v2.Name, "/")
				str := fmt.Sprintf("%s/%s/%s", strings.ToLower(vname[0]), strings.ToLower(v1name[0]), strings.ToLower(v2name[0]))
				accesslist[str] = true
			}
		}
	}
	return accesslist, nil
}
