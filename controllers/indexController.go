package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	AdminController
}

func (c *IndexController) Index() {
	c.Data["navnum"] = NAVINDEX
	c.TplName = "admin/index.tpl"
}

func (c *IndexController) ErrorPanic() {
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "admin/error_panic.tpl"
}
