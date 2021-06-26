package controllers

import (
	"blogo/app/models"

	"github.com/revel/revel"

	"strings"
	"time"
)

type Blog struct {
	AppController
}

/**
 * 博客首页
 */
func (c Blog) Index() revel.Result {
	var posts []models.Post
	err := models.Engine.Desc("id").Limit(4, 0).Find(&posts)
	if err != nil {
		panic(err)
	}

	c.ViewArgs["posts"] = posts

	return c.Render()
}

/**
 * 分页查询
 */
func (c Blog) QueryPage(page int) revel.Result {
	return c.Render()
}

/**
 * 文章内容
 */
func (c Blog) ShowPost(slug string) revel.Result {
	post := new(models.Post)
	has, _ := models.Engine.Where("Slug = ? ", slug).Get(post)
	if has {
		var comments []models.Comment
		models.Engine.Where("post_id = ?", post.Id).Find(&comments)
		c.ViewArgs["post"] = post
		c.ViewArgs["comments"] = comments
		return c.Render()
	}
	return c.NotFound("404你懂的！")
}

/**
 * 添加评论
 */
func (c Blog) AddComment(comment models.Comment) revel.Result {

	post := models.Post{}
	has, _ := models.Engine.Id(comment.PostId).Get(&post)
	if !has {
		return c.RenderJSON(&models.AjaxResult{Succ: false, Error: "评论的文章不存在"})
	}

	comment.Validate(c.Validation)
	if c.Validation.HasErrors() {
		var errorMsg string
		for _, validErr := range c.Validation.Errors {
			errorMsg += validErr.Message + "\n"
		}
		return c.RenderJSON(&models.AjaxResult{Succ: false, Error: errorMsg})
	}

	comment.Created = time.Now()
	affects, _ := models.Engine.InsertOne(&comment)
	if affects > 0 {
		sql := "update post set comment_count = comment_count + 1 where id = ?"
		models.Engine.Exec(sql, comment.PostId)
		return c.RenderJSON(&models.AjaxResult{Succ: true})
	}
	return c.RenderJSON(&models.AjaxResult{Succ: false, Error: "添加评论出错！"})
}

/**
 * 文章归档
 */
func (c Blog) Archives() revel.Result {

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

	c.ViewArgs["archiveMonths"] = archiveMonths
	c.ViewArgs["archiveMap"] = archiveMap
	return c.Render()
}

/**
 * 标签
 */
func (c Blog) Tags() revel.Result {

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

	c.ViewArgs["allTags"] = allTags
	return c.Render()
}

/**
 * 作品展示
 */
func (c Blog) Works() revel.Result {
	return c.Render()
}

/**
 * 关于
 */
func (c Blog) About() revel.Result {
	return c.Render()
}
