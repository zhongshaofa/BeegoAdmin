package routers

import (
	"BeegoAdmin/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//主页地址
	beego.Router("/", &controllers.IndexController{}, "Get:Index")

	//初始化接口
	beego.Router("/sys_init", &controllers.IndexController{}, "Get:SystemInit")
}
