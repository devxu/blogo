package controllers

import (
	"blogo/models"
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Home  Admin Home
func Home(c echo.Context) error {
	fmt.Println("Current session values:", getSession(c).Values)
	return c.Render(http.StatusOK, "home.html", nil)
}

// ListPost  Post list
func ListPost(c echo.Context) error {
	return c.Render(http.StatusOK, "listPost.html", nil)
}

// QueryPosts Paginated query posts
func QueryPosts(c echo.Context) error {
	draw, _ := strconv.Atoi(c.FormValue("draw"))
	start, _ := strconv.Atoi(c.FormValue("start"))
	length, _ := strconv.Atoi(c.FormValue("length"))

	var posts []models.Post
	models.Engine.Cols("id", "Slug", "title", "comment_count", "created").Desc("id").Limit(length, start).Find(&posts)
	total, _ := models.Engine.Count(&models.Post{})

	result := &models.DataTableResult{}
	result.Draw = int64(draw)
	result.Data = &posts
	result.RecordsFiltered = total
	result.RecordsTotal = total
	return c.JSON(http.StatusOK, result)
}

// CreatePost  Create post
func CreatePost(c echo.Context) error {
	post := &models.Post{}
	post.Slug = strings.Replace(uuid.NewString(), "-", "", -1)

	data := make(echo.Map)
	data["title"] = "创建文章"
	data["post"] = post
	return c.Render(http.StatusOK, "editPost.html", data)
}

// EditPost Edit post
func EditPost(c echo.Context) error {
	data := make(echo.Map)
	data["title"] = "编辑文章"

	id := c.Param("id")
	var post models.Post
	succ, _ := models.Engine.Id(id).Get(&post)
	if succ {
		data["post"] = &post
	}
	return c.Render(http.StatusOK, "editPost.html", data)
}

// DeletePost Delete post
func DeletePost(c echo.Context) error {

	id := c.Param("id")
	affects, _ := models.Engine.Id(id).Delete(&models.Post{})
	if affects > 0 {
		models.Engine.Where("post_id = ?", id).Delete(&models.Comment{})
		c.Set("flash.success", "已成功删除！")
	} else {
		c.Set("flash.error", "删除失败！")
	}
	return c.Redirect(http.StatusFound, "/admin/posts")
}

// SavePost Save post
func SavePost(c echo.Context) error {

	var post models.Post
	if err := c.Bind(&post); err != nil {
		return echo.ErrBadRequest
	}

	data := make(echo.Map)
	data["post"] = &post

	//post.Validate(c.Validation)
	//if c.Validation.HasErrors() {
	//	return c.RenderTemplate("admin/editPost.html")
	//}

	existPost := new(models.Post)
	has, _ := models.Engine.Where("Slug = ?", post.Slug).Get(existPost)
	if has && existPost.Id != post.Id {
		//c.Validation.Error("Slug已经被使用！")
		//return c.RenderTemplate("admin/editPost.html")
		return c.Render(http.StatusOK, "editPost.html", data)
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
		c.Set("flash.success", "保存成功！")
		return c.Redirect(http.StatusOK, "/admin/posts")
	} else {
		c.Set("flash.error", "保存失败！")
		return c.Render(http.StatusOK, "editPost.html", nil)
	}

}

// ListComment query comments
func ListComment(c echo.Context) error {
	return c.Render(http.StatusOK, "listComment.html", nil)
}

// QueryComments Paginated query comments
func QueryComments(c echo.Context) error {

	draw, _ := strconv.Atoi(c.FormValue("draw"))
	start, _ := strconv.Atoi(c.FormValue("start"))
	length, _ := strconv.Atoi(c.FormValue("length"))

	var comments []models.Comment
	models.Engine.Cols("id", "Name", "Message", "created").Desc("id").Limit(length, start).Find(&comments)
	total, _ := models.Engine.Count(&models.Comment{})

	result := &models.DataTableResult{}
	result.Draw = int64(draw)
	result.Data = &comments
	result.RecordsFiltered = total
	result.RecordsTotal = total
	return c.JSON(http.StatusOK, result)
}

// DeleteComment Delete comment by id
func DeleteComment(c echo.Context) error {

	id := c.Param("id")
	comment := models.Comment{}
	has, _ := models.Engine.Id(id).Get(&comment)
	if has {
		affects, _ := models.Engine.Id(id).Delete(&comment)
		if affects > 0 {
			sql := "update post set comment_count = comment_count - 1 where id = ?"
			models.Engine.Exec(sql, comment.PostId)
			c.Set("flash.success", "已成功删除！")
		} else {
			c.Set("flash.error", "删除失败！")

		}
	} else {
		c.Set("flash.error", "评论不存在！")
	}
	return c.Redirect(http.StatusFound, "/admin/comments")
}

// Login To login page
func Login(c echo.Context) error {
	data := make(echo.Map)
	data["title"] = "管理后台登录"
	return c.Render(http.StatusOK, "login.html", data)
}

// LoginSubmit Submit to login
func LoginSubmit(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if len(username) > 0 && len(password) > 0 {
		hash := md5.New()
		io.WriteString(hash, username+"^_^"+password)
		secret := strings.ToUpper(fmt.Sprintf("%x", hash.Sum(nil)))
		loginSecret, _ := models.MyConfig.String("login", "login.secret")
		if secret == loginSecret {
			sess := getSession(c)
			sess.Values["loginName"] = username
			sess.Save(c.Request(), c.Response())
			return c.Redirect(http.StatusFound, "/admin/home")
		}

	}
	c.Set(flash_error, "登录失败！")
	return c.Redirect(http.StatusFound, "/admin/login")
}

// Logout do logout
func Logout(c echo.Context) error {
	sess := getSession(c)
	sess.Values = make(map[interface{}]interface{})
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/admin/login")
}
