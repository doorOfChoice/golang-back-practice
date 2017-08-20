package controllers

import (
	"blog/models"
	"blog/tools"

	"github.com/astaxie/beego"
)

type VerifyController struct {
	beego.Controller
}

func (c *VerifyController) LoginPage() {
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "admin/login.tpl"
}

func (c *VerifyController) Login() {
	flash := beego.NewFlash()
	defer func() { flash.Store(&c.Controller); c.Redirect("/account/login", 302) }()

	username, password := c.GetString("username"), c.GetString("password")
	if !tools.FilterString(`^\w{4,16}$`, username) {
		flash.Error("用户名请保持在4-16个字符且只能是英文和数字")
	} else if !tools.FilterString(`^\w{6,16}$`, password) {
		flash.Error("密码请保持在6-16位字符")
	} else {
		var user = &models.User{Username: username, Password: tools.Sha256([]byte(password))}

		if err := user.Login(); err != nil {
			flash.Error(err.Error())
		} else {
			//设置cookie和session
			c.SetSecureCookie("sha256", "d_user", user.Token, 3600*24*7)
			c.SetSession("d_user", user.Token)
			flash.Success("登陆成功")
			c.Redirect("/admin/index", 302)
		}
	}

}

func (c *VerifyController) RegisterPage() {
	beego.ReadFromRequest(&c.Controller)
	c.TplName = "admin/register.tpl"
}

func (c *VerifyController) Register() {
	flash := beego.NewFlash()
	//验证信息是否有误
	crash := true
	defer func() {
		flash.Store(&c.Controller)
		if crash {
			c.Redirect("/account/register", 302)
		} else {
			c.Redirect("/account/login", 302)
		}
	}()
	username, pw1, pw2 := c.GetString("username"), c.GetString("pass-one"), c.GetString("pass-two")

	if !tools.FilterString(`^\w{4,16}$`, username) {
		flash.Error("用户名请保持在4-16个字符且只能是英文和数字")
	} else if pw1 != pw2 {
		flash.Error("两次密码输入不同")
	} else if !tools.FilterString(`^\w{6,16}$`, pw1, pw2) {
		flash.Error("为了你的安全, 密码请设置在6-16位字符")
	} else {
		var user = &models.User{
			Username:       username,
			Password:       tools.Sha256([]byte(pw1)),
			Token:          tools.Sha256([]byte(username)),
			Identification: models.POWER_USER,
		}

		if err := user.Create(); err != nil {
			flash.Error(err.Error())
		} else {
			crash = false
			flash.Success("注册成功，请登录")
		}

	}
}
