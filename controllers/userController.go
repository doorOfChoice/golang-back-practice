package controllers

import (
	"blog/models"
	"blog/tools"
	"strconv"

	"github.com/astaxie/beego"
)

type UserController struct {
	AdminController
}

func convertPower(powerStr string) int {
	power, err := strconv.Atoi(powerStr)

	if err != nil {
		return models.POWER_USER
	}
	if power < models.POWER_ADMIN || power > models.POWER_USER {
		return models.POWER_USER
	}
	return power
}

//Create 创建用户 【POST】
func (c *UserController) Create() {
	flash := beego.NewFlash()
	crash := true

	username, password, power := c.GetString("username"), c.GetString("password"), c.GetString("power")

	defer func() {
		flash.Data["Username"] = username
		flash.Store(&c.Controller)
		if crash {
			c.Redirect("/admin/user", 302)
		} else {
			c.Redirect("/admin/manaUsers", 302)
		}
	}()

	if !tools.FilterString(`^\w{4,16}$`, username, password) {
		flash.Error("用户名和密码请限制在4-16个字符")
	} else {
		user := &models.User{
			Username:       username,
			Password:       tools.Sha256([]byte(password)),
			Token:          tools.Sha256([]byte(username)),
			Identification: convertPower(power),
		}

		//权限检测
		if c.user.Identification >= user.Identification {
			flash.Error("你不能创建等级大于等于你的用户")
		} else {
			if err := user.Create(); err != nil {
				flash.Error(err.Error())
			} else {
				crash = false
				flash.Success("创建用户:" + username + "成功")
			}
		}

	}
}

func (c *UserController) Get() {
	beego.ReadFromRequest(&c.Controller)
	crash := true
	defer func() {
		if crash {
			c.Redirect("/admin/manaUsers", 302)
		}
	}()

	idStr := c.Ctx.Input.Param(":id")

	if id, err := strconv.Atoi(idStr); err == nil {
		if user, err := models.GetUser(id); err == nil {
			crash = false
			c.Data["navnum"] = NAVMANAUSERS
			c.Data["User"] = user
			c.Data["Powers"] = Powers
			c.TplName = "admin/update_user.tpl"
		}
	}
}

//Delete 删除用户 【Delete】
func (c *UserController) Delete() {
	flash := beego.NewFlash()
	defer func() {
		flash.Store(&c.Controller)
		c.Redirect("/admin/manaUsers", 302)
	}()
	idStr := c.Ctx.Input.Param(":id")

	if id, err := strconv.Atoi(idStr); err != nil {
		flash.Error("错误的查询参数")
	} else {
		if user, err := models.GetUser(id); err != nil {
			flash.Error(err.Error())
		} else {
			//权限检测
			if c.user.Identification >= user.Identification {
				flash.Error("对不起，你没有权限删除此用户")
				return
			}

			if err := user.Delete(); err != nil {
				flash.Error(err.Error())
			} else {
				flash.Success("删除成功")
			}
		}
	}
}

//Update 更新用户密码和权限信息 【Put】
func (c *UserController) Update() {
	flash := beego.NewFlash()
	defer func() {
		flash.Store(&c.Controller)
		c.Redirect(c.Ctx.Request.RequestURI, 302)
	}()
	idStr := c.Ctx.Input.Param(":id")

	if id, err := strconv.Atoi(idStr); err != nil {
		flash.Error("错误的查询参数")
	} else {
		password, power := c.GetString("password"), c.GetString("power")
		if user, err := models.GetUser(id); err != nil {
			flash.Error("用户不存在")
		} else {
			//检测本人是否有修改权限
			if c.user.Identification >= user.Identification {
				flash.Error("对不起，你没有权限更改此用户信息")
				return
			}
			//密码如果不规范，默认不修改
			if tools.FilterString(`^\w{4, 16}$`, password) {
				user.Password = password
			}

			//检测是否能够修改用户的权限
			user.Identification = convertPower(power)
			if c.user.Identification >= user.Identification {
				flash.Error("你只能设置他的权限比你低")
				return
			}

			if err := user.Update(); err != nil {
				flash.Error(err.Error())
			} else {
				flash.Success("更新成功")
			}
		}
	}
}

func (c *UserController) CreatePage() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["navnum"] = NAVMANAUSERS
	c.Data["Powers"] = Powers
	c.TplName = "admin/new_user.tpl"
}

//ManaPage 管理界面
func (c *UserController) ManaPage() {
	beego.ReadFromRequest(&c.Controller)
	paginator, users := models.GetUsersByPagination(c.Ctx.Request, 5, 10)
	c.Data["Users"] = users
	c.Data["P"] = paginator
	c.Data["navnum"] = NAVMANAUSERS
	c.TplName = "admin/mana_users.tpl"
}
