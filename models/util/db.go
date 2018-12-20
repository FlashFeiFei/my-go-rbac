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
