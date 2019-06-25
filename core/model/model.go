package model

import (
    "database/sql"
    "github.com/astaxie/beego/orm"
    "fmt"
    "beegoair/core/log"
    error2 "beegoair/core/error"
    "time"
)

type Model struct {
    table  string
    db  *sql.DB
    orm orm.Ormer
}

func (this *Model) InitModel(table string, db *sql.DB) {
    this.table = table
    this.db = db
}

func (this *Model) GetLog() string {
    return fmt.Sprintf("db.%s", this.table)
}

func (this *Model) Orm() orm.Ormer {
    if this.orm != nil {
        return this.orm
    }
    if this.db != nil {
        var err error
        this.orm,err = orm.NewOrmWithDB(DB_MYSQL, alias, this.db)
        if err != nil {
            panic("NewOrmWithDB error:" + err.Error())
        }
    } else {
        this.orm = orm.NewOrm()
    }
    if alias != "default" {
        this.orm.Using(alias)
    }
    return this.orm
}

func (this *Model) GetInfoById(id int64, data interface{}) error2.AppError {
    err := this.Orm().QueryTable(this.table).Filter("id", id).One(data)
    log.Debug(this.GetLog(), "GetInfoById", id, data)
    if err != nil && err != orm.ErrNoRows {
        log.Error(this.GetLog(), "GetInfoById", err.Error(), id, data)
        return error2.ErrorNew(error2.ERR_SELECT, err)
    }
    return nil
}

func (this *Model) GetListByParam(param []Condition,ordeby string,offset int,limit int64, data interface{}) (int64, error2.AppError) {
    query := this.Orm().QueryTable(this.table)
    for _,v := range param {
        query = query.Filter(v.Column,v.Value)
    }
    if len(ordeby) > 0 {
        query = query.OrderBy(ordeby)
    }
    if limit > 0 {
        query = query.Offset(int64(offset)*limit).Limit(limit)
    }
    num,err := query.All(data)
    log.Debug(this.GetLog(),"GetListByParam_info",param,ordeby,limit,data)
    if err != nil && err != orm.ErrNoRows{
        log.Error(this.GetLog(),"GetListByParam fail",err.Error(),param,ordeby,limit,data)
        return  0,error2.ErrorNew(error2.ERR_SELECT, err)
    }
    return num,nil
}

func (this *Model) GetInfoByParam(param []Condition, data interface{}) error2.AppError {
    query := this.Orm().QueryTable(this.table)
    for _,v := range param {
        query = query.Filter(v.Column,v.Value)
    }
    err := query.One(data)
    if err != nil && err != orm.ErrNoRows{
        log.Error(this.GetLog(),"GetInfoFromDB fail", err.Error(),param, data)
        return  error2.ErrorNew(error2.ERR_SELECT, err)
    }
    log.Debug(this.GetLog(),"GetInfoFromDB", param, data)
    return nil
}

func (this *Model) UpdateById(id int64, data map[string]interface{}) (int64, error2.AppError){
    if data["Utime"] == nil {
        data["Utime"] = time.Now().UnixNano()/1e6
    }
    query := this.Orm().QueryTable(this.table).Filter("id", id)
    num,err := query.Update(data)
    if err != nil {
        log.Error(this.GetLog(),"updatebyid fail", err.Error(), id, data)
        return 0,error2.ErrorNew(error2.ERR_UPDATE, err)
    }
    log.Debug(this.GetLog(),"UpdateById_data", id,data,num)
    return num,nil
}

func (this *Model) UpdateByParam(params []Condition,data map[string]interface{}) (int64, error2.AppError){
    query := this.Orm().QueryTable(this.table)
    for _,v := range params {
        query = query.Filter(v.Column,v.Value)
    }
    if data["Utime"] == nil {
        data["Utime"] = time.Now().UnixNano()/1e6
    }
    num,err := query.Update(data)
    if err != nil {
        log.Error(this.GetLog(),"UpdateByParam fail", err.Error(), params, data)
        return 0,error2.ErrorNew(error2.ERR_UPDATE, err)
    }
    log.Debug(this.GetLog(),"UpdateByParam_data", params,data,num)
    return num,nil
}

func (this *Model) Count(params []Condition) (int64,error2.AppError){
    query := this.Orm().QueryTable(this.table)
    for _,v := range params {
        query = query.Filter(v.Column,v.Value)
    }
    num,err := query.Count()
    if err != nil {
        log.Error(this.GetLog(),"getcount fail", err.Error(), params)
        return 0,error2.ErrorNew(error2.ERR_SELECT, err)
    }
    log.Debug(this.GetLog(),"Count_data",params,num)
    return num,nil
}

func (this *Model) Insert(data interface{}) (int64,error2.AppError){
    id,err := this.Orm().Insert(data)
    if err != nil {
        log.Error(this.GetLog(),"Insert fail", err.Error(), data)
        return id,error2.ErrorNew(error2.ERR_UPDATE, err)
    }
    return id,nil
}