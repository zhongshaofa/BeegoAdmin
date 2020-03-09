package controllers

import (
	"BeegoAdmin/models"
	"fmt"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	c.TplName = "index.html"
}

func (c *IndexController) SystemInit() {

	//menuList := []*models.SystemMenu{}
	//c.o.QueryTable(new(models.SystemMenu).TableName()).Filter("status", 1).All(&menuList)

	menuList := models.SystemMenu{}.MenuList()

	fmt.Println(menuList)

	jsonMap := make(map[string]interface{})
	jsonMap["code"] = "200"
	jsonMap["mag"] = "请求成功"
	//jsonMap["data"] = menuList
	c.Data["json"] = &jsonMap
	fmt.Println("执行到此处")
	c.ServeJSON()
}
