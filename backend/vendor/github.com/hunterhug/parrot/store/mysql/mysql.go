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
package mysql

// 数据库CURD，需要写SQL语句调用，简单就是美
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/hunterhug/parrot/store"
)

// Mysql config
type MysqlConfig struct {
	Username string
	Password string
	Ip       string
	Port     string
	Dbname   string
}

// a client
type Mysql struct {
	Config MysqlConfig
	Client *sql.DB
}

func New(config MysqlConfig) *Mysql {
	return &Mysql{Config: config}
}

//插入数据
//Insert Data
func (db *Mysql) Insert(prestring string, parm ...interface{}) (int64, error) {
	stmt, err := db.Client.Prepare(prestring)
	if err != nil {
		return 0, err
	}
	R, err := stmt.Exec(parm...)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	num, err := R.RowsAffected()
	return num, err

}

// 建表
// Create table
func (db *Mysql) Create(prestring string, parm ...interface{}) (int64, error) {
	stmt, err := db.Client.Prepare(prestring)
	if err != nil {
		return 0, err
	}
	R, err := stmt.Exec(parm...)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	num, err := R.RowsAffected()
	return num, err

}

// 删表
func (db *Mysql) Drop(prestring string, parm ...interface{}) (int64, error) {
	stmt, err := db.Client.Prepare(prestring)
	if err != nil {
		return 0, err
	}
	R, err := stmt.Exec(parm...)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	num, err := R.RowsAffected()
	return num, err

}

// create database
func (dbconfig MysqlConfig) CreateDb() error {
	dbname := dbconfig.Dbname
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;", dbname)
	dbconfig.Dbname = ""
	db := New(dbconfig)
	db.Open(30, 0)
	_, err := db.Create(sql)
	dbconfig.Dbname = dbname
	return err

}

func (dbconfig MysqlConfig) DeleteDb() error {
	dbname := dbconfig.Dbname
	sql := fmt.Sprintf("DROP DATABASE IF EXISTS `%s`;", dbname)
	dbconfig.Dbname = ""
	db := New(dbconfig)
	db.Open(30, 0)
	_, err := db.Create(sql)
	dbconfig.Dbname = dbname
	return err
}

//打开数据库连接 open a connecttion
//username:password@protocol(address)/dbname?param=value
func (db *Mysql) Open(maxopen int, maxidle int) {
	if db.Client != nil {
		return
	}
	dbs, err := sql.Open("mysql", db.Config.Username+":"+db.Config.Password+"@tcp("+db.Config.Ip+":"+db.Config.Port+")/"+db.Config.Dbname+"?charset=utf8")
	if err != nil {
		log.Logger.Fatalf("Open database error: %s", err.Error())
	}
	//defer dbs.Close()
	dbs.SetMaxIdleConns(maxidle)
	dbs.SetMaxOpenConns(maxopen)

	err = dbs.Ping()
	if err != nil {
		log.Logger.Fatalf("Ping err:%s", err.Error())
	}

	db.Client = dbs
}

//查询数据库 Query
func (db *Mysql) Select(prestring string, parm ...interface{}) (returnrows []map[string]interface{}, err error) {
	returnrows = []map[string]interface{}{}
	rows, err := db.Client.Query(prestring, parm...)
	if err != nil {
		return
	}

	defer rows.Close()
	// Get column names
	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		returnrow := map[string]interface{}{}
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			returnrow[columns[i]] = value

		}
		returnrows = append(returnrows, returnrow)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return
}
