package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type BaseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
	cdnUrl         string
}

func (c *BaseController) Prepare() {
	controllerName, actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	c.actionName = strings.ToLower(actionName)
	c.o = orm.NewOrm()

	c.cdnUrl = beego.AppConfig.String("cdn_url") // cdn配置，后期用于修改OSS位置的配置信息
	c.Data["cdnUrl"] = c.cdnUrl
}
