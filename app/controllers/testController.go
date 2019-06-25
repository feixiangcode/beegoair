package controllers

import (
	"beegoair/core/controller"
	"beegoair/core/cache"
	"beegoair/app/models"
	"beegoair/app/constants"
)

type TestController struct {
	controller.Controller
}

/*func (this *TestController) BeforeAction() bool {

	return false
}*/

func (this *TestController) Test() {
	//test redis
	redis := cache.NewAppRedis()
	redis.Set("test","baidu.com")
	data := map[string]interface{}{
		"age":35,
		"name":"kimi",
		"value":cache.ToString(redis.Get("test")),
	}
	this.Success(data)
}

func (this *TestController) Db() {
	model := models.NewTestModel(nil)
	result := new(constants.Test)
	model.GetInfoById(1, result)
	this.Success(result)
}

func (this *TestController) Static() {
	type NameAge struct {
		Name string
		Age int
	}

	a := NameAge{Name:"test",Age:10}
	b := NameAge{Name:"prod",Age:5}
	c := NameAge{Name:"dev",Age:1}

	list := make([]NameAge, 0)
	list = append(list, a,b,c)
	this.Data["list"] = &list

	mapList := make(map[string]interface{})
	mapList["a"] = "aaaa"
	mapList["b"] = "bbbb"
	mapList["c"] = "cccc"
	this.Data["maplist"] = &mapList

	this.TplName = "test.html"
}
