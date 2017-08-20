package models

import (
	"blog/tools"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/doorOfChoice/pagination"
)

type Post struct {
	Id         int       `orm:"auto;pk;"`
	Title      string    `orm:"size(500)"`
	Content    string    `orm:"type(text)"`
	CreateDate time.Time `orm:"null"`
	UpdateDate time.Time `orm:"null"`
	User       *User     `orm:"rel(fk);on_delete(cascade)"`
	Tags       []*Tag    `orm:"rel(m2m);on_delete(cascade)"`
}

//新建文章
func (this *Post) Create() error {
	o := orm.NewOrm()

	_, err := o.Insert(this)

	return err
}

func (this *Post) Delete() error {
	o := orm.NewOrm()

	_, err := o.Delete(this)

	return err
}

func (this *Post) Update() error {
	o := orm.NewOrm()

	_, err := o.Update(this)
	log.Println(err)
	return err
}

func GetPost(id int) (*Post, error) {
	o := orm.NewOrm()
	var post = &Post{Id: id}
	err := o.Read(post)

	switch err {
	case orm.ErrNoRows:
		return nil, errors.New("该文章不存在")
	case nil:
		o.LoadRelated(post, "Tags")
		return post, nil
	default:
		return nil, err
	}
}

//连接标签
func (this *Post) TachTags(tagStrings ...string) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(this, "Tags")
	tagStrings = tools.UniqStringArray(tagStrings)
	for _, v := range tagStrings {
		var tag = Tag{Name: v}
		o.ReadOrCreate(&tag, "Name")
		m2m.Add(&tag)
	}
}

func (this *Post) DetachTags(tagIds ...int) {
	o := orm.NewOrm()
	m2m := o.QueryM2M(this, "Tags")
	tagIds = tools.UniqIntArray(tagIds)
	for _, v := range tagIds {
		m2m.Remove(&Tag{Id: v})
	}
}

//获取分页数据
func GetPostByPagination(r *http.Request, linkCount, per int64) (*pagination.Paginator, []*Post) {
	o := orm.NewOrm()

	seter := o.QueryTable(&Post{})
	var posts []*Post
	//查询类型
	qType := r.URL.Query().Get("type")
	//查询值
	qValue := r.URL.Query().Get("value")

	switch qType {
	default:
		seter = seter.OrderBy("-id")
	case "id":
		seter = seter.Filter("id", qValue).OrderBy("-id")
	case "title":
		seter = seter.Filter("title__icontains", qValue).OrderBy("-id")
	}
	paginator := pagination.NewPaginator(r, seter, &posts, linkCount, per)

	return paginator, posts
}
