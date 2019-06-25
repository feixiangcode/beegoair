package application

import (
    "beegoair/core/log"
    "github.com/astaxie/beego"
    "beegoair/core/cache"
    "beegoair/core/model"
)

func ConsoleRun() {
    logConfig := new(log.Cfg)
    logConfig.LogPath = beego.AppConfig.String("logpath")
    logConfig.Level = beego.AppConfig.String("loglevel")
    isFile, err := beego.AppConfig.Bool("isfile")
    if err != nil {
        logConfig.IsFileLog = true
    } else {
        logConfig.IsFileLog = isFile
    }

    log.InitLog(logConfig)
    cache.InitRedis("redis")
    model.RegisterDataBase("db")
}
