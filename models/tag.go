package models

import (
	"errors"
	"net/http"
	"sort"

	"github.com/doorOfChoice/pagination"

	"github.com/astaxie/beego/orm"
)

type Tags []*Tag
type Posts []*Post

type Tag struct {
	Id    int    `orm:"auto;pk"`
	Name  string `orm:"unique"`
	Posts Posts  `orm:"reverse(many);on_delete(cascade)"`
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

func GetTagsByPagination(r *http.Request, linkCount, per int64) (*pagination.Paginator, Tags) {
	o := orm.NewOrm()
	var tags Tags
	//排序参数
	qOrder := r.URL.Query().Get("order")
	seter := o.QueryTable(&Tag{})

	//对id进行排序
	switch qOrder {
	case "id_asc":
		seter = seter.OrderBy("id")
	default:
		seter = seter.OrderBy("-id")
	}

	paginator := pagination.NewPaginator(r, seter, &tags, linkCount, per)

	for _, v := range tags {
		o.LoadRelated(v, "Posts")
	}

	//对标签文章量进行排序
	switch qOrder {
	case "desc":
		sort.Sort(sort.Reverse(tags))
	case "asc":
		sort.Sort(tags)
	}

	return paginator, tags
}

//[]*Tag 继承了Sort的排序接口，默认从小到大
func (c Tags) Len() int {
	return len(c)
}

func (c Tags) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Tags) Less(i, j int) bool {
	return len(c[i].Posts) < len(c[j].Posts)
}
