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
package admin

import (
	"fmt"
	"math/rand"

	"github.com/astaxie/beego"
	"github.com/hunterhug/rabbit/lib"
	"github.com/hunterhug/rabbit/models/blog"
	"github.com/hunterhug/rabbit/conf"
)

func InitData() {
	InsertUser()
	InsertGroup()
	InsertRole()
	InsertNodes()
	InsertConfig()
	InsertCategory()
	InsertRoll()
	InsertPaper()
}

//插入网站配置
func InsertConfig() {
	fmt.Println("insert config start")
	c := new(blog.Config)
	c.Photo = "/file/image/64/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png"
	c.Id = 1
	c.Title = "亚马逊大数据智能选款平台Plus"
	c.Webinfo = `
	{
		"1":{"name":"About","limit":6},
        "2":{"name":"News","limit":6},
        "3":{"name":"Lifes","limit":6},
        "4":{"name":"Production","limit":6},
        "5":{"name":"Flower","limit":6},
        "6":{"name":"TeaCup","limit":6}
	}
	`
	c.Phone = "0750-12345678"
	c.Content = `
<div align="center">
	<p>
		<span style="font-size:32px;">大数据智能选款平台，高效，快捷</span><br />
<img src="/file/image/53/68756e7465726875671e5573ac53bb5813b6b51d47d2db806b.gif" alt="" /><br />
<span style="font-size:32px;">吴邪为你打电话好运不断运气不断</span>
	</p>
</div>
	`
	c.Slogan = "选款难，如何百万商品千里挑一？亚马逊跨境电商利器，助你一臂之力。"
	c.Address = `<meta description="rabbit" >
<!-- some other script put in here -->`
	c.Code3 = `Power by hunterhug at 2017 此处页脚版权`
	c.Code2 = `Stats Code  此处放统计代码`
	c.Code1 = "Comment Code 此处放跟帖代码"
	err := c.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert config end")
}

func InsertCategory() {
	fmt.Println("insert category start")
	cs := map[int64]string{1: "About", 2: "News", 3: "Lifes", 4: "Production", 5: "Flower", 6: "TeaCup", 7: "Books", 8: "Musics"}
	for k, v := range cs {
		c := new(blog.Category)
		c.Id = k
		c.Title = v + "-T"
		c.Alias = v
		c.Createtime = lib.GetTime()
		c.Status = 1
		c.Image = "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png"
		c.Content = v
		if k == 4 {
			c.Type = 1
		}
		if k > 4 {
			c.Pid = 4
			c.Type = 1
		}
		err := c.Insert()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("insert category end")
}

func InsertPaper() {
	aaa := []string{
		"/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png",
		"/file/image/12/68756e74657268756795ad72d42c7ef1b56c04c66297db1c27.jpeg",
		"/file/image/64/68756e74657268756759fc5ee18fa5210bb76003976900fae9.jpeg",
	}
	k := 140
	for k > 0 {
		k = k - 1
		paper := new(blog.Paper)
		paper.Title = "Test test test test data"
		paper.Status = 1
		paper.Photo = aaa[rand.Intn(3)]
		paper.Descontent = "淭一厘晢臹隒丌蒛霒冘，庳乜砐淈琲葙丌漍厹刌。仈"
		paper.Content = `
So her / Ce are Xi Zan, but what Xiao this Wu may not put up the Yun Zhang ping. These Jie Tong Kang Schrodinger
Jiu Da Chi Yu a Feng Bing Wu Dao tea, Ke Bei Fu Tu cou not drag type go to squall process gas Bureau ze. A Qu% Xiu
Yan Que Yin find not weighted, Bi E Qu Pei cases Guo are not Rou cut. He Ge howl Xi Cang Hu Bo Kuang Nan Xin loop Yi
Fu Ru Chu Le Xie row full of a surname. He got Shan creating Zhi w Gun and Dun Cu Qiong Jian Hu You Yi Wei Ding Hai Li.
After a Xuan Bei Tu yao type seal Zeng ang, Guo Wu You Liu Wen Torr + Quan market dissatisfied. A Xian how prepared and
not Zhu Che Han he Xia, Chu Zhong Chen Yi Yan not falsification of Ji E. A Gong Wu Fei Guo zhe Ju Er through, please.
Guan melancholy that he was Jiang urn donburi. A host Gui Yuncheng Qin TA type. Zhi Ji, Ji type He Xun Zan Mei not Jiao
Dao le. Kun a sincere Sun what taste as a surname Ju Qi, sow not You Si Xiao type Hou tonnes through the. A Yan Zhao
`
		paper.Author = "hunterhug"
		paper.Createtime = lib.GetTime()
		paper.Cid = int64(rand.Intn(8) + 1)
		if paper.Cid >= 4 {
			paper.Type = 1
		}
		paper.Istop = int64(rand.Intn(2))
		paper.Insert()
	}
}

func InsertRoll() {
	rolls := map[string]string{
		"选款智能平台1": "/file/image/35/roll.jpg",
		"选款智能平台2": "/file/image/37/68756e7465726875673308fd68c821f8fb4180732625ef10ba.png",
		"选款智能平台3": "/file/image/35/roll.jpg",
		"选款智能平台4": "/file/image/37/68756e7465726875673308fd68c821f8fb4180732625ef10ba.png",
	}
	for k, v := range rolls {
		t := new(blog.Roll)
		t.Photo = v
		t.Status = 1
		t.Title = k
		t.Url = "/public"
		t.Createtime = lib.GetTime()
		t.Insert()
	}
}

// 用户数据
func InsertUser() {
	fmt.Println("insert user ...")
	u := new(User)
	u.Username = beego.AppConfig.String("rbac_admin_user")
	u.Nickname = "Admin"
	u.Password = lib.Pwdhash(beego.AppConfig.String("rbac_admin_user"))
	u.Email = "459527502@qq.com"
	u.Remark = "最高权限的王"
	// 2 stand for close, but it has very high authority
	u.Status = 2
	u.Createtime = lib.GetTime()
	err := u.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	u1 := new(User)
	u1.Username = "test"
	u1.Nickname = "Test"
	u1.Password = lib.Pwdhash("test")
	u1.Email = "459527502@qq.com"
	u1.Remark = "Just a Test User"
	u1.Status = 1
	u1.Createtime = lib.GetTime()
	err1 := u1.Insert()
	if err1 != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert user end")
}

// 模组数据
func InsertGroup() {
	fmt.Println("insert group ...")
	g := new(Group)
	g.Name = "兔子后台"
	g.Title = "后台管理"
	g.Sort = 1
	g.Id = 1
	g.Status = 1
	e := g.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	g1 := new(Group)
	g1.Name = "兔子后台"
	g1.Title = "文章管理"
	g1.Sort = 2
	g1.Id = 2
	g1.Status = 1
	e = g1.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	g2 := new(Group)
	g2.Name = "兔子后台"
	g2.Title = "图片管理"
	g2.Sort = 3
	g2.Id = 3
	g2.Status = 1
	e = g2.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println("insert group end")

	if conf.Amazon {
		gb := new(Group)
		gb.Name = "美国亚马逊"
		gb.Title = "美国亚马逊"
		gb.Sort = 4
		gb.Id = 4
		gb.Status = 1
		e = gb.Insert()
		if e != nil {
			fmt.Println(e.Error())
		}

		jgb := new(Group)
		jgb.Name = "日本亚马逊"
		jgb.Title = "日本亚马逊"
		jgb.Sort = 5
		jgb.Id = 5
		jgb.Status = 1
		e = jgb.Insert()
		if e != nil {
			fmt.Println(e.Error())
		}

		de := new(Group)
		de.Name = "德国亚马逊"
		de.Title = "德国亚马逊"
		de.Sort = 6
		de.Id = 6
		de.Status = 1
		e = de.Insert()
		if e != nil {
			fmt.Println(e.Error())
		}

		uk := new(Group)
		uk.Name = "英国亚马逊"
		uk.Title = "英国亚马逊"
		uk.Sort = 7
		uk.Id = 7
		uk.Status = 1
		e = uk.Insert()
		if e != nil {
			fmt.Println(e.Error())
		}

		yw := new(Group)
		yw.Name = "业务数据"
		yw.Title = "业务数据"
		yw.Sort = 8
		yw.Id = 8
		yw.Status = 1
		e = yw.Insert()
		if e != nil {
			fmt.Println(e.Error())
		}

	}
}

// 角色数据
func InsertRole() {
	fmt.Println("insert role ...")
	r := new(Role)
	r.Name = "管理员"
	r.Remark = "权限最高的一群人"
	r.Status = 1
	r.Title = "管理员角色"
	r.Insert()
	fmt.Println("insert role end")
}

// 节点数据
func InsertNodes() {
	fmt.Println("insert node ...")
	g := new(Group)
	g.Id = 1
	g1 := new(Group)
	g1.Id = 2
	g2 := new(Group)
	g2.Id = 3
	nodes := []Node{
		/*

			RBAC管理中心

		*/
		{Id: 1, Name: "rbac", Title: "权限中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		{Id: 2, Name: "node/index", Title: "节点管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 3, Name: "Index", Title: "节点首页", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 4, Name: "AddAndEdit", Title: "增编节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 5, Name: "DelNode", Title: "删除节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},

		{Id: 6, Name: "user/index", Title: "用户管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 7, Name: "Index", Title: "用户首页", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 8, Name: "AddUser", Title: "增加用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 9, Name: "UpdateUser", Title: "更新用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 10, Name: "DelUser", Title: "删除用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},

		{Id: 11, Name: "group/index", Title: "分组管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 12, Name: "Index", Title: "分组首页", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 13, Name: "AddGroup", Title: "增加分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 14, Name: "UpdateGroup", Title: "更新分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 15, Name: "DelGroup", Title: "删除分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},

		{Id: 16, Name: "role/index", Title: "角色管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 17, Name: "index", Title: "角色首页", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 18, Name: "AddAndEdit", Title: "增编角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 19, Name: "DelRole", Title: "删除角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 20, Name: "GetList", Title: "列出角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 21, Name: "AccessToNode", Title: "显示权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 22, Name: "AddAccess", Title: "增加权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 23, Name: "RoleToUserList", Title: "列出角色下用户", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 24, Name: "AddRoleToUser", Title: "授予用户角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},

		/*

			配置中心

		*/
		{Id: 25, Name: "config", Title: "配置中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		//-------
		//网站配置
		{Id: 26, Name: "option/index", Title: "网站配置", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 27, Name: "Index", Title: "网站配置首页", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		{Id: 28, Name: "UpdateOption", Title: "更新网站配置", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		//网站配置
		//个人信息
		{Id: 29, Name: "user/index", Title: "个人信息", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 30, Name: "Index", Title: "个人信息首页", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		{Id: 31, Name: "UpdateUser", Title: "更新个人信息", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		//个人信息

		/*

			文章中心

		*/
		{Id: 32, Name: "blog", Title: "文章中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g1},
		//------
		//文章目录
		{Id: 33, Name: "category/index", Title: "目录列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 34, Name: "Index", Title: "目录列表首页", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 35, Name: "AddCategory", Title: "增加目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 36, Name: "UpdateCategory", Title: "修改目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		//文章目录
		//文章
		{Id: 37, Name: "paper/index", Title: "文章列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 38, Name: "Index", Title: "文章列表首页", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 39, Name: "AddPaper", Title: "增加文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 40, Name: "UpdatePaper", Title: "修改文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 41, Name: "DeletePaper", Title: "回收文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 42, Name: "RealDelPaper", Title: "删除文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		//文章

		/*

			图片管理

		*/
		{Id: 43, Name: "picture", Title: "图片中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g2},
		//---------
		//相册
		{Id: 44, Name: "album/index", Title: "相册管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 45, Name: "Index", Title: "相册首页", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 46, Name: "AddAlbum", Title: "增加相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 47, Name: "DeleteAlbum", Title: "删除相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 48, Name: "UpdateAlbum", Title: "修改相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		//相册
		//图片
		{Id: 49, Name: "photo/index", Title: "图片管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 50, Name: "Index", Title: "图片首页", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 51, Name: "AddPhoto", Title: "增加图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 52, Name: "DeletePhoto", Title: "回收图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 53, Name: "UpdatePhoto", Title: "修改图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 54, Name: "RealDelPhoto", Title: "删除图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		//图片

		//补充的
		{Id: 55, Name: "DeleteCategory", Title: "删除目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},

		{Id: 56, Name: "paper/rubbish", Title: "文章回收站", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 57, Name: "Rubbish", Title: "文章回收站", Remark: "", Level: 3, Pid: 56, Status: 1, Group: g1},

		{Id: 58, Name: "photo/rubbish", Title: "图片回收站", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 59, Name: "Rubbish", Title: "图片回收站", Remark: "", Level: 3, Pid: 58, Status: 1, Group: g2},

		//首页图片轮转
		{Id: 60, Name: "roll/index", Title: "首页轮转", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 61, Name: "Index", Title: "轮转列表", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 62, Name: "AddRoll", Title: "增加轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 63, Name: "UpdateRoll", Title: "更新轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 64, Name: "DeleteRoll", Title: "删除轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
	}

	if conf.Amazon {
		gb := new(Group)
		gb.Id = 4

		jp := new(Group)
		jp.Id = 5

		de := new(Group)
		de.Id = 6

		uk := new(Group)
		uk.Id = 7

		yw := new(Group)
		yw.Id = 8
		amazonnode := []Node{
			{Id: 89, Name: "csv", Title: "业务数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: yw},
			{Id: 90, Name: "report/index", Title: "报告数据", Remark: "", Level: 2, Pid: 89, Status: 1, Group: yw},
			{Id: 91, Name: "Index", Title: "报告数据列表", Remark: "", Level: 3, Pid: 90, Status: 1, Group: yw},
			{Id: 92, Name: "Query", Title: "报告数据查询", Remark: "", Level: 3, Pid: 90, Status: 1, Group: yw},
			{Id: 93, Name: "Export", Title: "报告数据导出", Remark: "", Level: 3, Pid: 90, Status: 1, Group: yw},
			{Id: 94, Name: "Import", Title: "报告数据导入", Remark: "", Level: 3, Pid: 90, Status: 1, Group: yw},
			{Id: 95, Name: "Delete", Title: "报告数据删除", Remark: "", Level: 3, Pid: 90, Status: 1, Group: yw},


			{Id: 65, Name: "auas", Title: "基础数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: gb},
			{Id: 66, Name: "base/index", Title: "小类数据", Remark: "", Level: 2, Pid: 65, Status: 1, Group: gb},
			{Id: 67, Name: "Index", Title: "美国站小类数据列表", Remark: "", Level: 3, Pid: 66, Status: 1, Group: gb},
			{Id: 671, Name: "Query", Title: "美国站小类数据查询", Remark: "", Level: 3, Pid: 66, Status: 1, Group: gb},
			{Id: 672, Name: "Export", Title: "美国站小类数据导出", Remark: "", Level: 3, Pid: 66, Status: 1, Group: gb},

			{Id: 68, Name: "big/index", Title: "大类数据", Remark: "", Level: 2, Pid: 65, Status: 1, Group: gb},
			{Id: 69, Name: "Index", Title: "美国站大类数据列表", Remark: "", Level: 3, Pid: 68, Status: 1, Group: gb},
			{Id: 70, Name: "Query", Title: "美国站大类数据查询", Remark: "", Level: 3, Pid: 68, Status: 1, Group: gb},
			{Id: 71, Name: "Export", Title: "美国站大类数据导出", Remark: "", Level: 3, Pid: 68, Status: 1, Group: gb},
			{Id: 711, Name: "Asin", Title: "美国站Asin历史记录", Remark: "", Level: 3, Pid: 68, Status: 1, Group: gb},

			{Id: 72, Name: "asin/index", Title: "Asin数据", Remark: "", Level: 2, Pid: 65, Status: 1, Group: gb},
			{Id: 73, Name: "Index", Title: "美国站Asin数据列表", Remark: "", Level: 3, Pid: 72, Status: 1, Group: gb},
			{Id: 74, Name: "Query", Title: "美国站Asin数据查询", Remark: "", Level: 3, Pid: 72, Status: 1, Group: gb},
			{Id: 75, Name: "Export", Title: "美国站Asin数据导出", Remark: "", Level: 3, Pid: 72, Status: 1, Group: gb},

			{Id: 76, Name: "url/index", Title: "类目数据", Remark: "", Level: 2, Pid: 65, Status: 1, Group: gb},
			{Id: 77, Name: "Index", Title: "美国站类目数据列表", Remark: "", Level: 3, Pid: 76, Status: 1, Group: gb},
			{Id: 78, Name: "Query", Title: "美国站类目数据查询", Remark: "", Level: 3, Pid: 76, Status: 1, Group: gb},
			{Id: 79, Name: "Update", Title: "美国站类目数据更新", Remark: "", Level: 3, Pid: 76, Status: 1, Group: gb},

			{Id: 791, Name: "monitor/index", Title: "采集监控", Remark: "", Level: 2, Pid: 65, Status: 1, Group: gb},
			{Id: 792, Name: "Index", Title: "监控列表", Remark: "", Level: 3, Pid: 791, Status: 1, Group: gb},

			//---------------------------------------------------
			//日本数据
			{Id: 102, Name: "ajp", Title: "基础数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: jp},
			{Id: 103, Name: "base/index", Title: "小类数据", Remark: "", Level: 2, Pid: 102, Status: 1, Group: jp},
			{Id: 104, Name: "Index", Title: "日本站小类数据列表", Remark: "", Level: 3, Pid: 103, Status: 1, Group: jp},
			{Id: 105, Name: "Query", Title: "日本站小类数据查询", Remark: "", Level: 3, Pid: 103, Status: 1, Group: jp},
			{Id: 106, Name: "Export", Title: "日本站小类数据导出", Remark: "", Level: 3, Pid: 103, Status: 1, Group: jp},

			{Id: 107, Name: "big/index", Title: "大类数据", Remark: "", Level: 2, Pid: 102, Status: 1, Group: jp},
			{Id: 108, Name: "Index", Title: "日本站大类数据列表", Remark: "", Level: 3, Pid: 107, Status: 1, Group: jp},
			{Id: 109, Name: "Query", Title: "日本站大类数据查询", Remark: "", Level: 3, Pid: 107, Status: 1, Group: jp},
			{Id: 110, Name: "Export", Title: "日本站大类数据导出", Remark: "", Level: 3, Pid: 107, Status: 1, Group: jp},
			{Id: 712, Name: "Asin", Title: "日本站Asin历史记录", Remark: "", Level: 3, Pid: 107, Status: 1, Group: jp},

			{Id: 111, Name: "asin/index", Title: "Asin数据", Remark: "", Level: 2, Pid: 102, Status: 1, Group: jp},
			{Id: 112, Name: "Index", Title: "日本站Asin数据列表", Remark: "", Level: 3, Pid: 111, Status: 1, Group: jp},
			{Id: 113, Name: "Query", Title: "日本站Asin数据查询", Remark: "", Level: 3, Pid: 111, Status: 1, Group: jp},
			{Id: 114, Name: "Export", Title: "日本站Asin数据导出", Remark: "", Level: 3, Pid: 111, Status: 1, Group: jp},

			{Id: 115, Name: "url/index", Title: "类目数据", Remark: "", Level: 2, Pid: 102, Status: 1, Group: jp},
			{Id: 116, Name: "Index", Title: "日本站类目数据列表", Remark: "", Level: 3, Pid: 115, Status: 1, Group: jp},
			{Id: 117, Name: "Query", Title: "日本站类目数据查询", Remark: "", Level: 3, Pid: 115, Status: 1, Group: jp},
			{Id: 118, Name: "Update", Title: "日本站类目数据更新", Remark: "", Level: 3, Pid: 115, Status: 1, Group: jp},

			{Id: 1181, Name: "monitor/index", Title: "采集监控", Remark: "", Level: 2, Pid: 102, Status: 1, Group: jp},
			{Id: 1182, Name: "Index", Title: "监控列表", Remark: "", Level: 3, Pid: 1181, Status: 1, Group: jp},

			//---------------------------------------------------
			//亚马逊德国数据
			{Id: 119, Name: "de", Title: "基础数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: de},
			{Id: 120, Name: "base/index", Title: "小类数据", Remark: "", Level: 2, Pid: 119, Status: 1, Group: de},
			{Id: 121, Name: "Index", Title: "德国站小类数据列表", Remark: "", Level: 3, Pid: 120, Status: 1, Group: de},
			{Id: 122, Name: "Query", Title: "德国站小类数据查询", Remark: "", Level: 3, Pid: 120, Status: 1, Group: de},
			{Id: 123, Name: "Export", Title: "德国站小类数据导出", Remark: "", Level: 3, Pid: 120, Status: 1, Group: de},

			{Id: 124, Name: "big/index", Title: "大类数据", Remark: "", Level: 2, Pid: 119, Status: 1, Group: de},
			{Id: 125, Name: "Index", Title: "德国站大类数据列表", Remark: "", Level: 3, Pid: 124, Status: 1, Group: de},
			{Id: 126, Name: "Query", Title: "德国站大类数据查询", Remark: "", Level: 3, Pid: 124, Status: 1, Group: de},
			{Id: 127, Name: "Export", Title: "德国站大类数据导出", Remark: "", Level: 3, Pid: 124, Status: 1, Group: de},
			{Id: 128, Name: "Asin", Title: "德国站Asin历史记录", Remark: "", Level: 3, Pid: 124, Status: 1, Group: de},

			{Id: 129, Name: "asin/index", Title: "Asin数据", Remark: "", Level: 2, Pid: 119, Status: 1, Group: de},
			{Id: 130, Name: "Index", Title: "德国站Asin数据列表", Remark: "", Level: 3, Pid: 129, Status: 1, Group: de},
			{Id: 131, Name: "Query", Title: "德国站Asin数据查询", Remark: "", Level: 3, Pid: 129, Status: 1, Group: de},
			{Id: 132, Name: "Export", Title: "德国站Asin数据导出", Remark: "", Level: 3, Pid: 129, Status: 1, Group: de},

			{Id: 133, Name: "url/index", Title: "类目数据", Remark: "", Level: 2, Pid: 119, Status: 1, Group: de},
			{Id: 1331, Name: "Index", Title: "德国站类目数据列表", Remark: "", Level: 3, Pid: 133, Status: 1, Group: de},
			{Id: 1332, Name: "Query", Title: "德国站类目数据查询", Remark: "", Level: 3, Pid: 133, Status: 1, Group: de},
			{Id: 134, Name: "Update", Title: "德国站类目数据更新", Remark: "", Level: 3, Pid: 133, Status: 1, Group: de},

			{Id: 1341, Name: "monitor/index", Title: "采集监控", Remark: "", Level: 2, Pid: 119, Status: 1, Group: de},
			{Id: 1342, Name: "Index", Title: "监控列表", Remark: "", Level: 3, Pid: 1341, Status: 1, Group: de},

			//---------------------------------------------------
			//亚马逊英国数据
			{Id: 135, Name: "uk", Title: "基础数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: uk},
			{Id: 136, Name: "base/index", Title: "小类数据", Remark: "", Level: 2, Pid: 135, Status: 1, Group: uk},
			{Id: 137, Name: "Index", Title: "英国站小类数据列表", Remark: "", Level: 3, Pid: 136, Status: 1, Group: uk},
			{Id: 138, Name: "Query", Title: "英国站小类数据查询", Remark: "", Level: 3, Pid: 136, Status: 1, Group: uk},
			{Id: 139, Name: "Export", Title: "英国站小类数据导出", Remark: "", Level: 3, Pid: 136, Status: 1, Group: uk},

			{Id: 140, Name: "big/index", Title: "大类数据", Remark: "", Level: 2, Pid: 135, Status: 1, Group: uk},
			{Id: 141, Name: "Index", Title: "英国站大类数据列表", Remark: "", Level: 3, Pid: 140, Status: 1, Group: uk},
			{Id: 142, Name: "Query", Title: "英国站大类数据查询", Remark: "", Level: 3, Pid: 140, Status: 1, Group: uk},
			{Id: 143, Name: "Export", Title: "英国站大类数据导出", Remark: "", Level: 3, Pid: 140, Status: 1, Group: uk},
			{Id: 144, Name: "Asin", Title: "英国站Asin历史记录", Remark: "", Level: 3, Pid: 140, Status: 1, Group: uk},

			{Id: 145, Name: "asin/index", Title: "Asin数据", Remark: "", Level: 2, Pid: 135, Status: 1, Group: uk},
			{Id: 146, Name: "Index", Title: "英国站Asin数据列表", Remark: "", Level: 3, Pid: 145, Status: 1, Group: uk},
			{Id: 147, Name: "Query", Title: "英国站Asin数据查询", Remark: "", Level: 3, Pid: 145, Status: 1, Group: uk},
			{Id: 148, Name: "Export", Title: "英国站Asin数据导出", Remark: "", Level: 3, Pid: 145, Status: 1, Group: uk},

			{Id: 149, Name: "url/index", Title: "类目数据", Remark: "", Level: 2, Pid: 135, Status: 1, Group: uk},
			{Id: 150, Name: "Index", Title: "英国站类目数据列表", Remark: "", Level: 3, Pid: 149, Status: 1, Group: uk},
			{Id: 151, Name: "Query", Title: "英国站类目数据查询", Remark: "", Level: 3, Pid: 149, Status: 1, Group: uk},
			{Id: 152, Name: "Update", Title: "英国站类目数据更新", Remark: "", Level: 3, Pid: 149, Status: 1, Group: uk},

			{Id: 1521, Name: "monitor/index", Title: "采集监控", Remark: "", Level: 2, Pid: 135, Status: 1, Group: uk},
			{Id: 1522, Name: "Index", Title: "监控列表", Remark: "", Level: 3, Pid: 1521, Status: 1, Group: uk},
		}
		nodes = append(nodes, amazonnode...)
	}
	for _, v := range nodes {
		n := new(Node)
		n.Id = v.Id // 这句是无效的,后来 bug 被 beego 官方改好了
		n.Name = v.Name
		n.Title = v.Title
		n.Remark = v.Remark
		n.Level = v.Level
		n.Pid = v.Pid
		n.Status = v.Status
		n.Group = v.Group
		e := n.Insert()
		if e != nil {
			fmt.Printf("%#v:%#v\n", n, e.Error())
		}
	}
	fmt.Println("insert node end")
}
