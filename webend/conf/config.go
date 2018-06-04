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
package conf

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/hunterhug/parrot/util"
)

type FlagConfig struct {
	User        *string
	DbInit      *bool
	DbInitForce *bool
	Rbac        *bool
}

type AmazonDbConfig struct {
	Base string
	Data string
	Hash string
}

var (
	AuthType      int
	AuthGateWay   string
	AuthAdmin     string
	Cookie7       bool
	Version       string
	HomeTemplate  string
	AdminTemplate string

	DbType   string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DbName   string
	DbLog    string
	MYSQLDNS string

	Amazon    bool
	AmazonUSA AmazonDbConfig
	AmazonJP  AmazonDbConfig
	AmazonUK  AmazonDbConfig
	AmazonDE  AmazonDbConfig
	ConfigDir = util.CurDir()
)

func InitConfig() {
	// version
	Version = beego.AppConfig.DefaultString("version", "version2.0")
	beego.Trace("Version:", Version)

	AuthType, _ = strconv.Atoi(beego.AppConfig.String("user_auth_type"))
	AuthGateWay = beego.AppConfig.DefaultString("rbac_auth_gateway", "/public/login")
	Cookie7, _ = beego.AppConfig.Bool("cookie7")
	AuthAdmin = beego.AppConfig.DefaultString("rbac_admin_user", "admin")
	HomeTemplate = beego.AppConfig.DefaultString("home_template", "default")

	DbType = beego.AppConfig.String("db_type")
	DbHost = beego.AppConfig.String("db_host")
	DbPort = beego.AppConfig.String("db_port")
	DbUser = beego.AppConfig.String("db_user")
	DbPass = beego.AppConfig.String("db_pass")
	DbName = beego.AppConfig.String("db_name")
	DbLog = beego.AppConfig.String("dblog")

	MYSQLDNS = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DbUser, DbPass, DbHost, DbPort, DbName)

	AdminTemplate = beego.AppConfig.DefaultString("admin_template", "default")

	Amazon, _ = beego.AppConfig.Bool("amazon")
	AmazonUSA.Data = beego.AppConfig.String("usadatadb")
	AmazonUSA.Base = beego.AppConfig.String("usabasicdb")
	AmazonUSA.Hash = beego.AppConfig.String("usahashdb")

	AmazonJP.Data = beego.AppConfig.String("jpdatadb")
	AmazonJP.Base = beego.AppConfig.String("jpbasicdb")
	AmazonJP.Hash = beego.AppConfig.String("jphashdb")

	AmazonUK.Data = beego.AppConfig.String("ukdatadb")
	AmazonUK.Base = beego.AppConfig.String("ukbasicdb")
	AmazonUK.Hash = beego.AppConfig.String("ukhashdb")

	AmazonDE.Data = beego.AppConfig.String("dedatadb")
	AmazonDE.Base = beego.AppConfig.String("debasicdb")
	AmazonDE.Hash = beego.AppConfig.String("dehashdb")
}

func ForTestInitConfig() {
	err := beego.LoadAppConfig("ini", filepath.Join(ConfigDir, "app.conf"))
	if err != nil {
		panic(err.Error())
	}
	InitConfig()
}
