package main

import (
	"blog/models"
	_ "blog/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var (
		duser     = beego.AppConfig.String("duser")
		dpassword = beego.AppConfig.String("dpassword")
		database  = beego.AppConfig.String("database")
	)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", duser, dpassword, database))

	orm.RegisterModel(
		new(models.User),
		new(models.Post),
		new(models.Tag),
		new(models.Profile),
	)
	orm.Debug = true

	beego.AddFuncMap("add", add)
	beego.AddFuncMap("whatpower", mapPower)
	beego.AddFuncMap("stations", getStations)

	rebuild, err := beego.AppConfig.Bool("rebuilddatabase")
	if err == nil {
		if rebuild {
			orm.RunSyncdb("default", true, true)
		}
	}

}

func main() {
	beego.Run()
}
