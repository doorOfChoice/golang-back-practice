package models

import (
	"errors"
	"log"
	"net/http"

	"github.com/doorOfChoice/pagination"

	"github.com/astaxie/beego/orm"
)

const (
	//POWER_SUPER_ADMIN 超级管理员权限
	//天王老子
	POWER_SUPER_ADMIN = iota

	//POWER_ADMIN 管理员权限
	//包括用户所有的权限
	POWER_ADMIN

	//POWER_USER 用户权限
	//能够自己增删改查自己的文章，只能查看管理员文章
	POWER_USER
)

type User struct {
	Id             int `orm:"auto;pk"`
	Username       string
	Password       string
	Token          string   `orm:"unique"`
	Identification int      `orm:"null;default(0)"`
	Profile        *Profile `orm:"rel(one)"`
	Posts          []*Post  `orm:"reverse(many)"`
}

func (this *User) Create() error {
	o := orm.NewOrm()

	//检测该用户是否是第一位，如果是自动成为超级管理员
	if n, _ := o.QueryTable(&User{}).Count(); n == 0 {
		this.Identification = POWER_SUPER_ADMIN
	}
	//设置默认头像
	profile := &Profile{Head: "/static/img/default/header-custom.svg"}
	o.Insert(profile)

	this.Profile = profile
	b, _, err := o.ReadOrCreate(this, "Username")
	if err == nil {
		if !b {
			return errors.New("用户已经存在")
		}
		return nil
	}

	return err
}

func (this *User) Delete() error {
	o := orm.NewOrm()

	_, err := o.Delete(this)

	return err
}

func (this *User) Update() error {
	o := orm.NewOrm()

	_, err := o.Update(this)

	return err
}

func GetUsersByPagination(r *http.Request, linkCount, per int64) (*pagination.Paginator, []*User) {
	o := orm.NewOrm()
	var users []*User
	seter := o.QueryTable(&User{})

	qType, qValue := r.URL.Query().Get("type"), r.URL.Query().Get("value")

	switch qType {
	case "id":
		seter = seter.Filter("id", qValue).OrderBy("-id")
	case "username":
		seter = seter.Filter("username", qValue).OrderBy("-id")
	}

	paginator := pagination.NewPaginator(r, seter, &users, linkCount, per)

	return paginator, users
}

func GetUser(id int) (*User, error) {
	o := orm.NewOrm()

	user := &User{Id: id}

	err := o.Read(user)

	if err != nil {
		return nil, err
	}
	o.LoadRelated(user, "Profile")
	return user, nil
}

func (this *User) Login() error {
	o := orm.NewOrm()

	err := o.QueryTable(&User{}).Filter("username", this.Username).Filter("password", this.Password).One(this)
	log.Println(err)
	if err != nil {
		return errors.New("请检查账号或者用户名是否正确")
	}

	return nil
}

func FindUserByToken(token string) *User {
	o := orm.NewOrm()
	var user User

	err := o.QueryTable(&User{}).Filter("token", token).RelatedSel().One(&user)
	if err != nil {
		return nil
	}

	return &user
}
