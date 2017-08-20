package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

//升级方法,使之能使用PUT,DELETE等
func UpgradeMethod(ctx *context.Context) {
	if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
		ctx.Request.Method = ctx.Input.Query("_method")
	}
}

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, UpgradeMethod)
	//beego.InsertFilter("/admin/*", beego.BeforeRouter, VerifyUser)

	beego.Router("/account/login", &controllers.VerifyController{}, "GET:LoginPage;POST:Login")
	beego.Router("/account/register", &controllers.VerifyController{}, "GET:RegisterPage;POST:Register")

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/error/panic", &controllers.IndexController{}, "GET:ErrorPanic"),
		beego.NSRouter("/index", &controllers.IndexController{}, "GET:Index"),

		beego.NSRouter("/post", &controllers.PostController{}, "GET:CreatePage;POST:Create"),
		beego.NSRouter(`/post/:id(\d+)`, &controllers.PostController{}, "GET:Get;PUT:Update;DELETE:Delete"),
		beego.NSRouter("/manaPosts", &controllers.PostController{}, "GET:ManaPage"),

		beego.NSRouter("/tag", &controllers.TagController{}, "POST:Create"),
		beego.NSRouter(`/tag/:id(\d+)`, &controllers.TagController{}, "DELETE:Delete"),
		beego.NSRouter("/manaTags", &controllers.TagController{}, "GET:ManaPage"),

		beego.NSRouter("/user", &controllers.UserController{}, "POST:Create;GET:CreatePage"),
		beego.NSRouter(`/user/:id(\d+)`, &controllers.UserController{}, "DELETE:Delete;Put:Update;GET:Get"),
		beego.NSRouter("/manaUsers", &controllers.UserController{}, "GET:ManaPage"),
	)

	beego.AddNamespace(ns)
}
