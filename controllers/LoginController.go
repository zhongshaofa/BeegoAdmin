package controllers

import (
	"BeegoAdmin/services"
	"fmt"
)

type LoginController struct {
	BaseController
}

type PostLogin struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (c *LoginController) Index() {
	asd := new(services.Helpers).Password("123456")
	fmt.Println(asd)
	if c.IsAjax() {
		post := PostLogin{}
		_ = c.ParseForm(&post)
		c.Data["json"] = post
		c.ServeJSON()
	}
	c.Data["title"] = "BeegoAdmin"
	c.TplName = "login.html"
}
