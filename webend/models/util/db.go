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
package util

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hunterhug/rabbit/conf"
	"github.com/hunterhug/rabbit/models/admin"
	"time"
)

func init() {
	orm.DefaultTimeLoc = time.UTC
}

func Createtb() {
	beego.Trace("data init start")
	admin.InitData()
	beego.Trace("data init end")
}

func Syncdb(force bool) {
	beego.Trace("db, sync db start")

	Createdb(force)
	Connect()
	CreateConfig()
	Createtb()

	beego.Trace("sync db end, please reopen app again")
}

func UpdateRbac() {
	TruncateRbacTable([]string{beego.AppConfig.String("rbac_group_table"), beego.AppConfig.String("rbac_node_table")})
	Connect()
	admin.InsertGroup()
	admin.InsertNodes()
}

func CreateConfig() {
	name := "default" // database alias name
	force := true     // drop table force
	verbose := true   // print log
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error("database config set to force error:" + err.Error())
	}
}

//创建数据库
func Createdb(force bool) {
	beego.Trace("create database start")
	var dns, createdbsql, dropdbsql string

	switch conf.DbType {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", conf.DbUser, conf.DbPass, conf.DbHost, conf.DbPort)
		createdbsql = fmt.Sprintf("CREATE DATABASE if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", conf.DbName)
		dropdbsql = fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", conf.DbName)
		if force {
			fmt.Println(dropdbsql)
		}
		fmt.Println(createdbsql)
		break
	default:
		beego.Critical("db driver not support:", conf.DbType)
		return
	}
	db, err := sql.Open(conf.DbType, dns)
	if err != nil {
		panic(err.Error())
	}
	if force {
		_, err = db.Exec(dropdbsql)
	}
	_, err1 := db.Exec(createdbsql)
	if err != nil || err1 != nil {
		beego.Error("db exec error：", err, err1)
		panic(err.Error())
	} else {
		beego.Trace("database ", conf.DbName, " created")
	}

	if conf.Amazon {
		db.Exec("CREATE DATABASE  if not exists smart_backstage CHARSET utf8 COLLATE utf8_general_ci")
		_, derr := db.Exec(`
CREATE TABLE smart_backstage.report (
  id varchar(100) NOT NULL,
  pasin varchar(100) DEFAULT NULL COMMENT 'Asin(父)',
  asin varchar(100) DEFAULT NULL COMMENT 'Asin(子)',
  title varchar(100) DEFAULT NULL COMMENT '商品名称',
  uv int(11) DEFAULT NULL COMMENT '买家访问次数',
  uvb varchar(100) DEFAULT NULL COMMENT '该日买家访问次数百分比',
  pv int(11) DEFAULT NULL COMMENT '页面浏览次数',
  pvb varchar(100) DEFAULT NULL COMMENT '该日页面浏览次数百分比',
  bpvb varchar(100) DEFAULT NULL COMMENT '该日购物车占比',` +
			"`on`" + ` int(11) DEFAULT NULL COMMENT '该日已订购商品数量',
  onr varchar(100) DEFAULT NULL COMMENT '该日订单商品数量转化率',
  v double DEFAULT NULL COMMENT '该日已订购商品销售额',
  c int(11) DEFAULT NULL COMMENT '该日订单数',
  d varchar(100) DEFAULT NULL COMMENT '日期',
  aws varchar(100) DEFAULT NULL COMMENT '店铺名',
  status tinyint(4) DEFAULT NULL COMMENT 'Asin(父)',
  PRIMARY KEY (id)
) ENGINE=InnoDB;
`)
		if derr != nil {
			beego.Trace(derr.Error())
		}
	}
	defer db.Close()
	beego.Trace("create database end")
}

func TruncateRbacTable(table []string) {
	beego.Trace("delete tables start")
	var dns, sqlstring string
	switch conf.DbType {
	case "mysql":
		dns = conf.MYSQLDNS
	default:
		beego.Critical("db driver not support:", conf.DbType)
		return
	}
	db, err := sql.Open(conf.DbType, dns)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	for _, i := range table {
		beego.Trace("table deleting：" + i)
		sqlstring = fmt.Sprintf("TRUNCATE TABLE `%s`", i)
		_, err = db.Exec(sqlstring)
		if err != nil {
			beego.Error("table delete error：" + err.Error())
			panic(err.Error())
		} else {
			beego.Trace("table delete success：" + i)
		}
	}
	beego.Trace("delete table end")
}

func Connect() {
	var dns string
	switch conf.DbType {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = conf.MYSQLDNS
		break
	default:
		beego.Critical("db driver not support:", conf.DbType)
		return
	}

	beego.Trace("database start to connect", dns)
	err := orm.RegisterDataBase("default", conf.DbType, dns)
	if err != nil {
		beego.Error("register data:" + err.Error())
		panic(err.Error())
	}

	if conf.DbLog == "open" {
		beego.Trace("develop mode，debug database: db.log")
		orm.Debug = true
		w, e := os.OpenFile("log/db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if e != nil {
			beego.Error(e.Error())
		}
		orm.DebugLog = orm.NewLog(w)
	}

	RegisterDBModel() // must register
}

func ConnectAmazon() {
	orm.RegisterDataBase("usadatadb", "mysql", conf.AmazonUSA.Data)
	orm.RegisterDataBase("usabasicdb", "mysql", conf.AmazonUSA.Base)
	orm.RegisterDataBase("usahashdb", "mysql", conf.AmazonUSA.Hash)

	orm.RegisterDataBase("jpdatadb", "mysql", conf.AmazonJP.Data)
	orm.RegisterDataBase("jpbasicdb", "mysql", conf.AmazonJP.Base)
	orm.RegisterDataBase("jphashdb", "mysql", conf.AmazonJP.Hash)

	orm.RegisterDataBase("dedatadb", "mysql", conf.AmazonDE.Data)
	orm.RegisterDataBase("debasicdb", "mysql", conf.AmazonDE.Base)
	orm.RegisterDataBase("dehashdb", "mysql", conf.AmazonDE.Hash)

	orm.RegisterDataBase("ukdatadb", "mysql", conf.AmazonUK.Data)
	orm.RegisterDataBase("ukbasicdb", "mysql", conf.AmazonUK.Base)
	orm.RegisterDataBase("ukhashdb", "mysql", conf.AmazonUK.Hash)

	//其他数据库
	dbback := beego.AppConfig.String("dbback")
	orm.RegisterDataBase("dbback", "mysql", dbback)

}
