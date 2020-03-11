package controllers

import (
	"BeegoAdmin/models"
)

type IndexController struct {
	BaseController
}

// 主页
func (c *IndexController) Index() {
	c.TplName = "index.html"
}

// 初始化后台框架接口
func (c *IndexController) SystemInit() {
	systemInit := new(models.SystemMenu).GetSystemInit()
	c.Data["json"] = systemInit
	c.ServeJSON()
}

// 清理缓存
func (c *IndexController) Clear() {
	data := Message{Code: 1, Msg: "清理服务端缓存成功"}
	c.Data["json"] = data
	c.ServeJSON()
}
