package controllers

import (
	"blog/models"
	"blog/tools"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type PostController struct {
	AdminController
}

// Get 获取单个文章页面
func (c *PostController) Get() {
	beego.ReadFromRequest(&c.Controller)
	idStr := c.Ctx.Input.Param(":id")

	if id, err := strconv.Atoi(idStr); err != nil {
		c.Redirect("/admin/error/panic", 302)
	} else if post, err := models.GetPost(id); err != nil {
		c.Redirect("/admin/error/panic", 302)
	} else {
		c.Data["Post"] = post
		c.Data["navnum"] = NAVMANAPOSTS
		c.TplName = "admin/update_post.tpl"
	}
}

func (c *PostController) Create() {
	flash := beego.NewFlash()
	title, tags, content := c.GetString("title"), c.GetString("tags"), c.GetString("content")
	defer func() {
		flash.Data["Content"] = content
		flash.Data["Tags"] = tags
		flash.Data["Title"] = title
		flash.Store(&c.Controller)
		c.Redirect("/admin/post", 302)
	}()

	if !tools.FilterString(`^.{1,40}$`, title) {
		flash.Error("标题请限制在1-40个字符间")
	} else if !tools.FilterString(`^[\p{Han}|\w]+(,[\p{Han}|\w]+)*$`, tags) {
		flash.Error("请输入正确标签, ex. xxx,xxx,xxx")
	} else {
		t := tools.GetLocalTime()
		post := models.Post{
			Title:      title,
			Content:    content,
			User:       c.user,
			CreateDate: t,
			UpdateDate: t,
		}
		if err := post.Create(); err != nil {
			flash.Error(err.Error())
			return
		}

		tagArrays := tools.UniqStringArray(strings.Split(tags, ","))
		post.TachTags(tagArrays...)

		flash.Success("创建文章成功")
	}
}

func (c *PostController) Update() {
	flash := beego.NewFlash()
	defer func() {
		flash.Store(&c.Controller)
		c.Redirect(c.Ctx.Request.RequestURI, 302)
	}()

	idStr := c.Ctx.Input.Param(":id")

	if id, err := strconv.Atoi(idStr); err != nil {
		flash.Error("错误的查询参数")
	} else if post, err := models.GetPost(id); err != nil {
		flash.Error(err.Error())
	} else {
		//权限检测
		if c.user.Id != post.User.Id && c.user.Identification >= post.User.Identification {
			flash.Error("对不起，你没有权限修改对方的文章")
			return
		}
		//要被删除的标签
		var deletedIds []int
		c.Ctx.Input.Bind(&deletedIds, "deleted-ids")

		title, newTags, content := c.GetString("title"), c.GetString("new-tags"), c.GetString("content")

		if !tools.FilterString(`^.{1,40}$`, title) {
			flash.Error("标题请限制在1-40个字符间")
		} else {
			ok := tools.FilterString(`^[\p{Han}|\w]+?(,[\p{Han}|\w]+?)*$`, newTags)
			blank := strings.TrimSpace(newTags) == ""
			if (!ok && blank) || ok {
				post.Title = title
				post.Content = content
				post.Update()
				post.DetachTags(deletedIds...)
				if !blank {
					newTagArrays := tools.UniqStringArray(strings.Split(newTags, ","))
					post.TachTags(newTagArrays...)
				}
				flash.Success("更新成功")
				return
			}
			flash.Error("请检查标签格式是否正确")
		}
	}
}

func (c *PostController) Delete() {
	flash := beego.NewFlash()
	defer func() {
		flash.Store(&c.Controller)
		c.Redirect("/admin/manaPosts", 302)
	}()

	idStr := c.Ctx.Input.Param(":id")

	if id, err := strconv.Atoi(idStr); err != nil {
		flash.Error("错误的查询参数")
	} else if post, err := models.GetPost(id); err != nil {
		flash.Error(err.Error())
	} else {
		//权限检测
		if c.user.Id != post.User.Id && c.user.Identification >= post.User.Identification {
			flash.Error("对不起，你没有权限删除对方的文章")
			return
		}
		post.Delete()
		flash.Success("删除成功")
	}
}

func (c *PostController) CreatePage() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["navnum"] = NAVMANAPOSTS
	c.TplName = "admin/new_post.tpl"
}

func (c *PostController) ManaPage() {
	beego.ReadFromRequest(&c.Controller)
	paginator, posts := models.GetPostByPagination(c.Ctx.Request, 10, 10)
	c.Data["navnum"] = NAVMANAPOSTS
	c.Data["P"] = paginator
	c.Data["Posts"] = posts
	c.TplName = "admin/mana_posts.tpl"
}
