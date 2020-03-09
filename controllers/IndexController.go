package controllers

import "fmt"

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	fmt.Println("===================")
	fmt.Println(c.cdnUrl)
	c.TplName = "index.html"
}
