package controllers

import (
	"blogo/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

// Index The blog home page
func Index(c echo.Context) error {
	var posts []models.Post
	err := models.Engine.Desc("id").Limit(4, 0).Find(&posts)
	if err != nil {
		panic(err)
	}

	data := make(echo.Map)
	data["posts"] = posts
	return c.Render(http.StatusOK, "index.html", data)
}

// QueryPage Paginated query posts
func QueryPage(c echo.Context) error {
	//page := c.Param("page")
	return c.Render(http.StatusOK, "", nil)
}

// ShowPost Show post page
func ShowPost(c echo.Context) error {

	slug := c.Param("slug")
	post := new(models.Post)
	has, _ := models.Engine.Where("Slug = ? ", slug).Get(post)
	if has {
		var comments []models.Comment
		models.Engine.Where("post_id = ?", post.Id).Find(&comments)

		data := make(echo.Map)
		data["post"] = post
		data["comments"] = comments
		return c.Render(http.StatusOK, "showPost.html", data)
	}
	return echo.ErrNotFound
}

// AddComment Add comment
func AddComment(c echo.Context) error {

	var comment models.Comment
	if err := c.Bind(&comment); err != nil {
		return echo.ErrBadRequest
	}

	post := models.Post{}
	has, _ := models.Engine.Id(comment.PostId).Get(&post)
	if !has {
		return c.JSON(http.StatusOK, &models.AjaxResult{Succ: false, Error: "评论的文章不存在"})
	}

	//comment.Validate(c.Validation)
	//if c.Validation.HasErrors() {
	//	var errorMsg string
	//	for _, validErr := range c.Validation.Errors {
	//		errorMsg += validErr.Message + "\n"
	//	}
	//	return c.RenderJSON(&models.AjaxResult{Succ: false, Error: errorMsg})
	//}

	comment.Created = time.Now()
	affects, _ := models.Engine.InsertOne(&comment)
	if affects > 0 {
		sql := "update post set comment_count = comment_count + 1 where id = ?"
		models.Engine.Exec(sql, comment.PostId)
		return c.JSON(http.StatusOK, &models.AjaxResult{Succ: true})
	}
	return c.JSON(http.StatusOK, &models.AjaxResult{Succ: false, Error: "添加评论出错！"})
}

// Archives Show post archives
func Archives(c echo.Context) error {

	var posts []models.Post
	err := models.Engine.Desc("id").Find(&posts)
	if err != nil {
		panic(err)
	}

	archiveMonths := []string{}
	var archiveMap = map[string][]models.Post{}
	for _, p := range posts {
		year_month := p.GetCreated().Format("2006-01")
		if archiveMap[year_month] == nil {
			archiveMonths = append(archiveMonths, year_month)
			archiveMap[year_month] = []models.Post{}
		}
		archiveMap[year_month] = append(archiveMap[year_month], p)
	}

	data := make(echo.Map)
	data["archiveMonths"] = archiveMonths
	data["archiveMap"] = archiveMap
	return c.Render(http.StatusOK, "archives.html", data)
}

// Tags Show all tags
func Tags(c echo.Context) error {

	allTags := map[string]int64{}
	var posts []models.Post
	models.Engine.Distinct("tags").Where("tags is not null").Find(&posts)
	for _, p := range posts {
		tags := strings.Split(p.Tags, ",")
		for _, tag := range tags {
			if tag != "" && len(strings.TrimSpace(tag)) > 0 {
				allTags[tag] = allTags[tag] + 1
			}
		}

	}

	data := make(echo.Map)
	data["allTags"] = allTags
	return c.Render(http.StatusOK, "tags.html", data)
}

// Works Show all works
func Works(c echo.Context) error {
	return c.Render(http.StatusOK, "works.html", nil)
}

// About To about page
func About(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", nil)
}
