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
	ViewCount  int       `orm:"null;default(0)"`
	CreateDate time.Time `orm:"null"`
	UpdateDate time.Time `orm:"null"`
	User       *User     `orm:"rel(fk);on_delete(cascade)"`
	Tags       Tags      `orm:"rel(m2m);on_delete(cascade)"`
}

//新建文章
func (this *Post) Create() error {
	o := orm.NewOrm()

	_, err := o.Insert(this)

	return err
}

//删除文章
func (this *Post) Delete() error {
	o := orm.NewOrm()

	_, err := o.Delete(this)

	return err
}

//更新文章
func (this *Post) Update() error {
	o := orm.NewOrm()

	_, err := o.Update(this)
	log.Println(err)
	return err
}

//根据id获取文章
func GetPost(id int) (*Post, error) {
	o := orm.NewOrm()
	var post = &Post{Id: id}
	err := o.Read(post)

	switch err {
	case orm.ErrNoRows:
		return nil, errors.New("该文章不存在")
	case nil:
		o.LoadRelated(post, "Tags")
		o.LoadRelated(post, "User")
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

//删除标签
func (this *Post) DetachTags(tagIds ...int) {
	if len(tagIds) == 0 {
		return
	}

	o := orm.NewOrm()
	m2m := o.QueryM2M(this, "Tags")
	tagIds = tools.UniqIntArray(tagIds)
	var tags Tags
	for _, v := range tagIds {
		tags = append(tags, &Tag{Id: v})
	}
	m2m.Remove(tags)
}

//获取分页数据
func GetPostByPagination(r *http.Request, linkCount, per int64) (*pagination.Paginator, Posts) {
	o := orm.NewOrm()
	seter := o.QueryTable(&Post{}).RelatedSel("User")
	var posts Posts
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
	case "user":
		seter = seter.Filter("user__username__eq", qValue).OrderBy("-id")
	case "tag":
		var tag Tag
		if err := o.QueryTable(&Tag{}).Filter("name", qValue).One(&tag); err != nil {
			log.Println(1, "FAIL")
			return &pagination.Paginator{}, Posts{}
		}

		paginator := pagination.NewPaginatorByFiled(r, &tag, "Posts", linkCount, per)
		log.Println(paginator.OffsetFrom, paginator.OffsetTo)
		return paginator, tag.Posts[paginator.OffsetFrom:paginator.OffsetTo]
	}

	paginator := pagination.NewPaginator(r, seter, &posts, linkCount, per)

	return paginator, posts
}
