package application

import (
	"beegoair/core/log"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"runtime/debug"
	"beegoair/core/controller"
	"beegoair/core/cache"
	"beegoair/core/model"
)

func Run() {
	logConfig := new(log.Cfg)
	logConfig.LogPath = beego.AppConfig.String("logpath")
	logConfig.Level = beego.AppConfig.String("loglevel")
	isFile, err := beego.AppConfig.Bool("isfile")
	if err != nil {
		logConfig.IsFileLog = true
	} else {
		logConfig.IsFileLog = isFile
	}
	beego.BConfig.Log.AccessLogs = false
	beego.BConfig.RecoverPanic = true
	beego.BConfig.RecoverFunc = panicHandler

	autoRender,_ := beego.AppConfig.Bool("autorender")
	if autoRender {
		beego.ErrorController(&controller.WebErrorController{})
	} else {
		beego.ErrorController(&controller.ApiErrorController{})
	}
	log.InitLog(logConfig)
	cache.InitRedis("redis")
	model.RegisterDataBase("db")
}

func panicHandler(ctx *context.Context) {
	if err := recover(); err != nil {
		log.Error("error panic", fmt.Sprintf("%+v", err), string(debug.Stack()))
	}
}
