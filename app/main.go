package main

import (
	_ "beegoair/app/init"
	"github.com/astaxie/beego"
	"beegoair/core/application"
)

func main() {
	application.Run()
	beego.Run()
}
