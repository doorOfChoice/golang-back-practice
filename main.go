package main

import (
	"blog/models"
	_ "blog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1997@/dawnblog?charset=utf8")
	orm.RegisterModel(
		new(models.User),
		new(models.Post),
		new(models.Tag),
	)
	beego.SetStaticPath("/static/js", "static/js")
	beego.SetStaticPath("/static/css", "static/css")
	beego.SetStaticPath("/static/image", "static/img")
	beego.SetStaticPath("/static/vedio", "static/vedio")

	orm.Debug = true

	beego.AddFuncMap("add", func(a, b int) int {
		return a + b
	})

	beego.AddFuncMap("whatpower", func(power int) string {
		switch power {
		case models.POWER_ADMIN:
			return "管理员"
		case models.POWER_USER:
			return "普通用户"
		case models.POWER_SUPER_ADMIN:
			return "超级管理员"
		default:
			return "Unkonw"
		}
	})
	//orm.RunSyncdb("default", true, true)
}

func main() {
	beego.Run()
	// o := orm.NewOrm()
	// var ps []*models.Post
	// o.QueryTable(&models.Post{}).Filter("title__icontains", "ew").OrderBy("-id").All(&ps)
	// for _, v := range ps {
	// 	log.Println(v.Title)
	// }
}
