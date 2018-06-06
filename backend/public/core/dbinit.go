/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
*/
package core

import (
	"fmt"
)

func InitDB() {
	bdcon := MyConfig.Basicdb
	eee := bdcon.CreateDb()
	if eee != nil {
		fmt.Println(eee.Error())
	}

	ddcon := MyConfig.Datadb
	eee = ddcon.CreateDb()
	if eee != nil {
		fmt.Println(eee.Error())
	}

	hdcon := MyConfig.Hashdb
	eee = hdcon.CreateDb()
	if eee != nil {
		fmt.Println(eee.Error())
	}
	// next
	OpenMysql()
	sql := `
	CREATE TABLE smart_category (
	  id varchar(100) NOT NULL,
	  url varchar(255) DEFAULT NULL COMMENT '类目链接',
	  name varchar(255) DEFAULT NULL COMMENT '类目名字',
	  level tinyint(4) DEFAULT NULL COMMENT '类目级别',
	  pid varchar(100) DEFAULT NULL COMMENT '父类id',
	  createtime datetime DEFAULT NULL COMMENT '创建时间',
	  updatetime datetime DEFAULT NULL COMMENT '更新时间',
	  isvalid tinyint(4) DEFAULT '0' COMMENT '是否有效',
	  page tinyint(4) DEFAULT '5' COMMENT '抓取页数',
	  col1 varchar(255) DEFAULT NULL COMMENT '预留字段',
	  col2 varchar(255) DEFAULT NULL,
	  col3 varchar(255) DEFAULT NULL,
	  bigpname varchar(255) DEFAULT NULL COMMENT '大类名字',
	  bigpid varchar(100) DEFAULT NULL COMMENT '大类ID',
	  ismall tinyint(4) DEFAULT '0' COMMENT '是否最小类',
	  PRIMARY KEY (id),
	  UNIQUE KEY url_UNIQUE (url)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='类目';
	`
	_, e := BasicDb.Create(sql)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(sql)
	}
	monitorsql := `
	CREATE TABLE smart_monitor (
	  id VARCHAR(30) NOT NULL,
	  redistype VARCHAR(50) NOT NULL,
	  doing INT NULL,
	  done INT NULL,
	  updatetime VARCHAR(45) NULL,
	  createtime VARCHAR(45) NULL,
	  PRIMARY KEY (id, redistype)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8;
	`
	_, me := BasicDb.Create(monitorsql)
	if me != nil {
		fmt.Println(me.Error())
	} else {
		fmt.Println(monitorsql)
	}

	asinsql := `
	CREATE TABLE smart_asin (
	  id varchar(100) NOT NULL,
	  createtime varchar(255) DEFAULT NULL COMMENT '添加时间',
	  updatetime varchar(255) DEFAULT NULL COMMENT '更新时间',
	  category varchar(255) DEFAULT NULL COMMENT "which category",
	  times int(11) DEFAULT '0' COMMENT '重复次数',
	  isvalid tinyint(4) DEFAULT '1' COMMENT "valid",
	  PRIMARY KEY (id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='Asin Big Data';
	`
	_, asine := BasicDb.Create(asinsql)
	if asine != nil {
		fmt.Println(asine.Error())
	} else {
		fmt.Println(asinsql)
	}
	for i := 0; i <= MyConfig.Hashnum; i++ {
		a := fmt.Sprintf(`CREATE TABLE %sA%d%s(
	  id varchar(150) NOT NULL,
	  day varchar(150) NOT NULL,
	  bigname varchar(255) DEFAULT NULL COMMENT '大类名',
	  title TEXT COMMENT '商品标题',
      relate TEXT COMMENT 'USA附加ASIN购物篮',
	  other1 TEXT COMMENT '预留',
	  other2 TEXT COMMENT '预留',
	  rank int(11) DEFAULT NULL COMMENT '大类排名',
	  price float DEFAULT NULL,
	  sold varchar(255) DEFAULT NULL COMMENT '自营',
	  ship varchar(255) DEFAULT NULL COMMENT 'FBA',
	  score float DEFAULT NULL COMMENT '打分',
	  reviews int(11) DEFAULT NULL COMMENT '评论数',
	  createtime varchar(255) DEFAULT NULL,
	  img varchar(255) DEFAULT NULL,
	  PRIMARY KEY (id,day)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`, "`", i, "`")
		_, e := HashDb.Create(a)
		if e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Println(i)
		}
	}
}
/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
*/