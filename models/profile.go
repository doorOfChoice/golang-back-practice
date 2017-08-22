package models

import (
	"github.com/astaxie/beego/orm"
)

type Profile struct {
	Id           int    `orm:"pk;auto"`
	Name         string `orm:"null"`
	Head         string `orm:"null;size(1000)"`
	Hobby        string `orm:"null;size(1000)"`
	Stations     string `orm:"null;size(1000)"`
	Introduction string `orm:"null;size(3000)"`
	User         *User  `orm:"reverse(one)"`
}

func (this *Profile) Update() error {
	o := orm.NewOrm()

	if _, err := o.Update(this); err != nil {
		return err
	}

	return nil
}

func GetProfile(id int) (*Profile, error) {
	o := orm.NewOrm()

	profile := &Profile{Id: id}

	if err := o.Read(profile); err != nil {
		return nil, err
	}

	o.LoadRelated(profile, "User")
	return profile, nil
}
