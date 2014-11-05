package models

import (
	"github.com/revel/revel"
	"time"
)

type Post struct {
	Id           int64
	Slug         string
	Title        string `xorm:"VARCHAR(300)"`
	Content      string `xorm:"text"`
	Tags         string `xorm:"text"`
	CommentCount int
	Created      time.Time
}

func (post *Post) Validate(v *revel.Validation) {
	v.Required(post.Slug).Message("请填写访问Slug")
	v.Required(post.Title).Message("标题不能为空")
	v.Required(post.Content).Message("文章内容不能为空")
}

func (post *Post) GetCreated() time.Time {
	return post.Created.Add(-time.Hour * 8)
}
