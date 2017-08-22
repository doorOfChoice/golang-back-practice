package controllers

import (
	"blog/models"
	"strconv"

	"github.com/astaxie/beego"
)

type ProfileController struct {
	AdminController
}

func (c *ProfileController) Update() {
	flash := beego.NewFlash()
	defer func() {
		flash.Store(&c.Controller)
		rURL := "/admin/user/" + strconv.Itoa(c.user.Id)
		c.Redirect(rURL, 302)
	}()

	idStr := c.Ctx.Input.Param(":id")

	if id, err := strconv.Atoi(idStr); err != nil {
		flash.Error("错误的查询参数")
	} else if profile, err := models.GetProfile(id); err != nil {
		flash.Error("个人资料不存在")
	} else {
		//检测权限
		if c.user.Id != profile.User.Id {
			flash.Error("你没有权限修改这份个人资料")
		} else {
			profile.Name = c.GetString("user-name")
			profile.Hobby = c.GetString("user-hobby")
			profile.Stations = c.GetString("user-stations")
			profile.Introduction = c.GetString("user-introduction")
			profile.Head = c.GetString("user-head")
			if err := profile.Update(); err != nil {
				flash.Error("个人资料更新失败")
			} else {
				flash.Success("个人资料更新成功")
			}
		}
	}
}
