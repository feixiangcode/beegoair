package controller

import error2 "beegoair/core/error"

type ApiErrorController struct {
	Controller
}

func (this *ApiErrorController) Error400() {
	this.Fail(error2.ERR_HTTP_400)
}
func (this *ApiErrorController) Error401() {
	this.Fail(error2.ERR_HTTP_401)
}
func (this *ApiErrorController) Error402() {
	this.Fail(error2.ERR_HTTP_402)
}
func (this *ApiErrorController) Error403() {
	this.Fail(error2.ERR_HTTP_403)
}
func (this *ApiErrorController) Error404() {
	this.Fail(error2.ERR_HTTP_404)
}
func (this *ApiErrorController) Error405() {
	this.Fail(error2.ERR_HTTP_405)
}
func (this *ApiErrorController) Error406() {
	this.Fail(error2.ERR_HTTP_406)
}
func (this *ApiErrorController) Error407() {
	this.Fail(error2.ERR_HTTP_407)
}
func (this *ApiErrorController) Error408() {
	this.Fail(error2.ERR_HTTP_408)
}
func (this *ApiErrorController) Error409() {
	this.Fail(error2.ERR_HTTP_409)
}
func (this *ApiErrorController) Error410() {
	this.Fail(error2.ERR_HTTP_410)
}

func (this *ApiErrorController) Error500() {
	this.Fail(error2.ERR_HTTP_500)
}
func (this *ApiErrorController) Error501() {
	this.Fail(error2.ERR_HTTP_501)
}
func (this *ApiErrorController) Error502() {
	this.Fail(error2.ERR_HTTP_502)
}
func (this *ApiErrorController) Error503() {
	this.Fail(error2.ERR_HTTP_503)
}
func (this *ApiErrorController) Error504() {
	this.Fail(error2.ERR_HTTP_504)
}
func (this *ApiErrorController) Error505() {
	this.Fail(error2.ERR_HTTP_505)
}
