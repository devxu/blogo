package controllers

import (
	"blogo/app/models"
	"code.google.com/p/go-uuid/uuid"
	"crypto/md5"
	"fmt"
	"github.com/revel/revel"
	"io"
	"strconv"
	"strings"
	"time"
)

type Admin struct {
	AppController
}

/** 后台首页*/
func (c Admin) Index() revel.Result {
	fmt.Println("session = ", c.Session())
	return c.Render()
}

/** 文章列表 */
func (c Admin) ListPost() revel.Result {
	return c.Render()
}

/** 分页查询Posts返回JSON格式*/
func (c Admin) QueryPosts() revel.Result {
	draw, _ := strconv.Atoi(c.Params.Get("draw"))
	start, _ := strconv.Atoi(c.Params.Get("start"))
	length, _ := strconv.Atoi(c.Params.Get("length"))

	var posts []models.Post
	models.Engine.Cols("id", "title", "comment_count", "created").Desc("id").Limit(length, start).Find(&posts)
	total, _ := models.Engine.Count(&models.Post{})

	result := &models.DataTableResult{}
	result.Draw = int64(draw)
	result.Data = &posts
	result.RecordsFiltered = total
	result.RecordsTotal = total
	return c.RenderJson(result)
}

/** 创建文章*/
func (c Admin) CreatePost() revel.Result {
	c.RenderArgs["title"] = "创建文章"
	post := &models.Post{}
	post.Slug = strings.Replace(uuid.NewUUID().String(), "-", "", -1)
	c.RenderArgs["post"] = post
	return c.RenderTemplate("admin/editPost.html")
}

/** 编辑文章*/
func (c Admin) EditPost(id int64) revel.Result {
	c.RenderArgs["title"] = "编辑文章"
	var post models.Post
	succ, _ := models.Engine.Id(id).Get(&post)
	if succ {
		c.RenderArgs["post"] = &post
	}
	return c.Render()
}

/** 删除文章*/
func (c Admin) DeletePost(id int64) revel.Result {
	affects, _ := models.Engine.Id(id).Delete(&models.Post{})
	if affects > 0 {
		models.Engine.Where("post_id = ?", id).Delete(&models.Comment{})
		c.Flash.Success("已成功删除！")
	} else {
		c.Flash.Error("删除失败！")
	}
	return c.Redirect("/admin/posts")
}

/** 保存文章*/
func (c Admin) SavePost(post models.Post) revel.Result {

	c.RenderArgs["post"] = &post

	post.Validate(c.Validation)
	if c.Validation.HasErrors() {
		return c.RenderTemplate("admin/editPost.html")
	}

	existPost := new(models.Post)
	has, _ := models.Engine.Where("Slug = ?", post.Slug).Get(existPost)
	fmt.Println("has == ", has)
	if has && existPost.Id != post.Id {
		c.Validation.Error("Slug已经被使用！")
		return c.RenderTemplate("admin/editPost.html")
	}

	var affects int64
	if post.Id <= 0 {
		post.CommentCount = 0
		post.Created = time.Now()
		affects, _ = models.Engine.InsertOne(&post)
	} else {
		affects, _ = models.Engine.Id(post.Id).Update(&post)
	}

	if affects > 0 {
		c.Flash.Success("保存成功！")
		return c.Redirect("/admin/posts")
	} else {
		c.Validation.Error("保存失败！")
		return c.RenderTemplate("admin/editPost.html")
	}

}

/** 查看评论*/
func (c Admin) ListComment() revel.Result {
	return c.Render()
}

/**
 * 分页查询评论返回JSON结果
 */
func (c Admin) QueryComments() revel.Result {
	draw, _ := strconv.Atoi(c.Params.Get("draw"))
	start, _ := strconv.Atoi(c.Params.Get("start"))
	length, _ := strconv.Atoi(c.Params.Get("length"))

	var comments []models.Comment
	models.Engine.Cols("id", "Name", "Message", "created").Desc("id").Limit(length, start).Find(&comments)
	total, _ := models.Engine.Count(&models.Comment{})

	result := &models.DataTableResult{}
	result.Draw = int64(draw)
	result.Data = &comments
	result.RecordsFiltered = total
	result.RecordsTotal = total
	return c.RenderJson(result)
}

/**
 * 删除评论
 */
func (c Admin) DeleteComment(id int64) revel.Result {
	affects, _ := models.Engine.Id(id).Delete(&models.Comment{})
	if affects > 0 {
		c.Flash.Success("已成功删除！")
	} else {
		c.Flash.Error("删除失败！")
	}
	return c.Redirect("/admin/comments")
}

/**
 * 登录页面
 */
func (c Admin) Login() revel.Result {
	c.RenderArgs["title"] = "管理后台登录"
	return c.Render()
}

/**
 * 提交登录请求
 */
func (c Admin) LoginSubmit(username string, password string) revel.Result {
	if len(username) > 0 && len(password) > 0 {
		hash := md5.New()
		io.WriteString(hash, username+"G$#$%^1352%"+password)
		secret := strings.ToUpper(fmt.Sprintf("%x", hash.Sum(nil)))
		loginSecret, _ := models.MyConfig.String("login", "login.secret")
		if secret == loginSecret {
			c.Session().Set("islogin", true)
			return c.Redirect("/admin/index")
		}

	}
	c.Flash.Error("登录失败！")
	return c.Redirect("/admin/login")
}

/**
 * 退出
 */
func (c Admin) Logout() revel.Result {
	c.Session().Invalidate() //销毁Session
	return c.Redirect("/admin/login")
}
