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
	"time"

	"github.com/astaxie/beego/orm"
)

type Roll struct {
	Id         int64
	Title      string    `orm:"size(100)"`       //标题
	Content    string    `orm:"type(text);null"` //内容
	Createtime time.Time `orm:"type(datetime);null"`
	Updatetime time.Time `orm:"type(datetime);null"`
	Sort       int64     `orm:"default(0)"` //排序
	Status     int64     `orm:"default(0)"` //0 关闭 1开启
	Photo      string    `orm:"null"`       //图片地址
	View       int64     `orm:"default(0)"` //浏览量
	Url        string    `orm:"null"`
}

func (m *Roll) TableName() string {
	return "roll"
}

func (m *Roll) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Roll) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Roll) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Roll) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Roll) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
