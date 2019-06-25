package filter

import (
	"beegoair/core/log"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("*", beego.BeforeStatic, BeforeStatic)
	beego.InsertFilter("*", beego.BeforeRouter, BeforeRouter)
	beego.InsertFilter("*", beego.BeforeExec, BeforeExec)
	beego.InsertFilter("*", beego.AfterExec, AfterExec)
	beego.InsertFilter("*", beego.FinishRouter, FinishRouter)
}

func BeforeStatic(ctx *context.Context) {
	log.Info("BeforeStatic")
}

func BeforeRouter(ctx *context.Context) {
	log.Info("BeforeRouter")
}

func BeforeExec(ctx *context.Context) {
	log.Info("BeforeExec")
}

func AfterExec(ctx *context.Context) {
	log.Info("AfterExec")
}

func FinishRouter(ctx *context.Context) {
	log.Info("FinishRouter")
}
