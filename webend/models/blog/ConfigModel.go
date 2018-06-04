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
package blog

import (
	"github.com/astaxie/beego/orm"
)

type Config struct {
	Id      int64
	Title   string `orm:"size(100)"`       //标题
	Content string `orm:"type(text);null"` //网站描述
	Address string `orm:"type(text);null"` //地址，已经改为网站头部
	Phone   string `orm:"null"`            //联系方式
	Webinfo string `orm:"type(text);null"` //备案信息，已经改为首页配置
	Photo   string `orm:"null"`            //logo
	Slogan  string `orm:"type(text);null"` //漂移通知
	Code1   string `orm:"type(text);null"` //跟帖代码
	Code2   string `orm:"type(text);null"` //统计代码
	Code3   string `orm:"type(text);null"` //网站脚部
}

func (m *Config) TableName() string {
	return "config"
}

func (m *Config) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Config) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Config) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Config) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Config) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
