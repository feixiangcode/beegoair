package model

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

const DB_MYSQL = "mysql"

var alias string

func RegisterDataBase(pri string) {
    dns := beego.AppConfig.String(pri + "dns")
    if len(dns) < 1 {
        panic("dns error")
    }
    al := beego.AppConfig.String(pri + "alias")
    if len(al) < 1 {
        al = "default"
    }
    alias = al
    maxIdle,_ := beego.AppConfig.Int(pri + "maxIdle")
    maxConn,_ := beego.AppConfig.Int(pri + "maxConn")
    err := orm.RegisterDataBase(alias, DB_MYSQL, dns,maxIdle,maxConn)
    if err != nil {
        panic("RegisterDataBase error:" + err.Error())
    }
    if beego.AppConfig.String("runmode") != "prod" {
        orm.Debug = true
    }
}

func GetDB() *sql.DB {
    db,err := orm.GetDB(alias)
    if err != nil {
        panic("get db connect err:" + err.Error())
    }
    return db
}


