package init

import (
	"beegoair/app/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router
	apiRouter := beego.NewNamespace("/api",
		beego.NSRouter("/test/get", &controllers.TestController{}, "get:Test"),
		beego.NSRouter("/test/db", &controllers.TestController{}, "get:Db"),
		beego.NSRouter("/test/static", &controllers.TestController{}, "get:Static"),
	)
	beego.AddNamespace(apiRouter)
}
