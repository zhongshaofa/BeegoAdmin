package routers

import (
	"BeegoAdmin/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.IndexController{}, "Get:Index")
	beego.Router("/system_init", &controllers.IndexController{}, "Get:SystemInit")
	beego.Router("/clear", &controllers.IndexController{}, "Get:Clear")

}
