package controllers

import (
	"blog/models"
	"blog/tools"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type TagController struct {
	AdminController
}

func (c *TagController) Create() {
	flash := beego.NewFlash()
	defer func() {
		flash.Store(&c.Controller)
		c.Redirect("/admin/manaTags", 302)
	}()

	name := c.GetString("name")
	log.Println(name)
	if !tools.FilterString(`^[\p{Han}|\w]{1,40}$`, name) {
		flash.Error("标签名不合法")
	} else {
		tag := &models.Tag{Name: name}

		if err := tag.Create(); err != nil {
			flash.Error(err.Error())
		} else {
			flash.Success("创建成功")
		}
	}
}

func (c *TagController) Delete() {
	flash := beego.NewFlash()
	defer func() {
		flash.Store(&c.Controller)
		c.Redirect("/admin/manaTags", 302)
	}()

	idStr := c.Ctx.Input.Param(":id")

	if id, err := strconv.Atoi(idStr); err != nil {
		flash.Error("错误的查询参数")
	} else {
		if c.user.Identification >= models.POWER_USER {
			flash.Error("只有管理员以及管理员以上才能删除标签")
		} else {
			tag := &models.Tag{Id: id}
			tag.Delete()

			flash.Success("删除成功")
		}
	}
}

func (c *TagController) ManaPage() {
	beego.ReadFromRequest(&c.Controller)
	paginator, tags := models.GetTagsByPagination(c.Ctx.Request, 6, 10)
	c.Data["navnum"] = NAVMANATAGS
	c.Data["P"] = paginator
	c.Data["Tags"] = tags
	c.TplName = "admin/mana_tags.tpl"
}
