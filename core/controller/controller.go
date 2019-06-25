package controller

import (
	error2 "beegoair/core/error"
	"encoding/json"
	"github.com/astaxie/beego"
)

const CODE_SUCCESS = 0
const CODE_FAIL = -1

type BeforeAction interface {
	BeforeAction() bool
}

type Controller struct {
	beego.Controller
}

func (this *Controller) GetError(code int) (map[string]interface{}, int) {
	item := error2.ErrorList.Get(code)
	return map[string]interface{}{
		"code":  CODE_FAIL,
		"data":  nil,
		"error": map[string]interface{}{"code": item.Code, "msg": item.Msg},
	}, item.Status
}

func (this *Controller) Prepare() {
	if app, ok := this.AppController.(BeforeAction); ok {
		rt := app.BeforeAction()
		if rt == false {
			autoRender,_ := beego.AppConfig.Bool("autorender")
			if autoRender {
				item := error2.ErrorList.Get(error2.ERR_NOAUTH)
				this.Data["msg"] = &item
				this.TplName = "error.html"
				this.Render()
				this.StopRun()
			} else {
				this.Fail(error2.ERR_NOAUTH)
				this.StopRun()
			}
		}
	}
}

func (this *Controller) JsonDecode(input interface{}) (err error) {
	return json.Unmarshal(this.Ctx.Input.RequestBody, input)
}

func (this *Controller) Success(data interface{}) {
	body := map[string]interface{}{
		"code":  CODE_SUCCESS,
		"data":  data,
		"error": map[string]interface{}{},
	}
	this.Data["json"] = body
	this.ServeJSON()
}

func (this *Controller) Fail(code int) {
	errData, status := this.GetError(code)
	this.Data["json"] = errData
	this.Ctx.Output.Status = status
	this.ServeJSON()
}
