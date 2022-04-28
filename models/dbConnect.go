package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var dsn string
	db_type := beego.AppConfig.String("db::dbType")
	db_host := beego.AppConfig.String("db::dbHost")
	db_port := beego.AppConfig.String("db::dbPort")
	db_user := beego.AppConfig.String("db::dbUser")
	db_pass := beego.AppConfig.String("db::dbPass")
	db_name := beego.AppConfig.String("db::dbName")

	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	// ORM 使用 golang 自己的连接池
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", db_type, dsn, maxIdle, maxConn)
	////1.注册数据库
	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/im_pro?charset=utf8")
	////2.注册表
	//orm.RegisterModel(new(database.Users), new(database.Articles))
	////3.生成表
	//orm.RunSyncdb("default", false, true) // 如果表存在，则不更新
	// 打开调试模式，开发的时候方便查看orm生成什么样子的sql语句
	orm.Debug = true
}

// 新增数据
