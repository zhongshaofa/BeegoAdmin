package controllers

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}
