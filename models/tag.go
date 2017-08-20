package models

import (
	"errors"
	"net/http"

	"github.com/doorOfChoice/pagination"

	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id    int     `orm:"auto;pk"`
	Name  string  `orm:"unique"`
	Posts []*Post `orm:"reverse(many);on_delete(cascade)"`
}

func (this *Tag) Create() error {
	o := orm.NewOrm()
	b, _, err := o.ReadOrCreate(this, "Name")
	if err == nil {
		if b {
			return nil
		}
		return errors.New("该标签已经存在")
	}

	return err
}

func (this *Tag) Delete() error {
	o := orm.NewOrm()

	_, err := o.Delete(this)

	return err
}

func GetTagsByPagination(r *http.Request, linkCount, per int64) (*pagination.Paginator, []*Tag) {
	o := orm.NewOrm()
	var tags []*Tag
	seter := o.QueryTable(&Tag{}).RelatedSel().OrderBy("-id")
	paginator := pagination.NewPaginator(r, seter, &tags, linkCount, per)

	return paginator, tags
}
