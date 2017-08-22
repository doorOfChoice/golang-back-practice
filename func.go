package main

import (
	"blog/models"
	"blog/tools"
)

type Station struct {
	Name string
	Icon string
	Href string
}

//相加两个数
func add(a, b int) int {
	return a + b
}

//显示职称
func mapPower(power int) string {
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
}

//获取个人站点的名称和链接
func getStations(str string) []Station {
	if !tools.FilterString(`(\[(.+?)\]\((.+?)\);)*`, str) {
		return []Station{}
	}
	params := tools.FindParams(`\[(.+?)\]\((.+?)\);`, str)
	var stations []Station
	for _, v := range params {
		station := Station{Name: v[1], Href: v[2]}
		if tools.FileExist("static/img/icons/" + v[1] + ".png") {
			station.Icon = "/static/img/icons/" + v[1] + ".png"
		} else {
			station.Icon = "/static/img/icons/unknowed.png"
		}
		stations = append(stations, station)
	}

	return stations
}
