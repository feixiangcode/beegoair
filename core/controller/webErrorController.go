package controller

import error2 "beegoair/core/error"

type WebErrorController struct {
	Controller
}

func (this *WebErrorController) Error400() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_400)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error401() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_401)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error402() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_402)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error403() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_403)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error404() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_404)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error405() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_405)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error406() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_406)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error407() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_407)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error408() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_408)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error409() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_409)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error410() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_410)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}

func (this *WebErrorController) Error500() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_500)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error501() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_501)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error502() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_502)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error503() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_503)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error504() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_504)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
func (this *WebErrorController) Error505() {
	item := error2.ErrorList.Get(error2.ERR_HTTP_505)
	this.Data["msg"] = &item
	this.TplName = "error.html"
}
