package controllers

import (
	"blog/models"
	"log"

	"github.com/astaxie/beego"
)

const (
	NAVINDEX = iota + 1
	NAVMANAPOSTS
	NAVMANATAGS
	NAVMANAUSERS
)

var Powers = []int{models.POWER_SUPER_ADMIN, models.POWER_ADMIN, models.POWER_USER}

type AdminController struct {
	user    *models.User
	isLogin bool
	beego.Controller
}

func (c *AdminController) Prepare() {
	log.Println("no")
	if str, b := c.Ctx.GetSecureCookie("sha256", "d_user"); !b {
		c.Redirect("/account/login", 302)
	} else if sess := c.Ctx.Input.Session("d_user"); sess == nil {
		c.Redirect("/account/login", 302)
	} else {
		if sess.(string) != str {
			c.Redirect("/account/login", 302)
			return
		}
		c.isLogin = true
		c.user = models.FindUserByToken(str)
		c.Data["LoginUser"] = c.user
		c.Data["Powers"] = Powers
	}
}
