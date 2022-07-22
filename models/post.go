package models

import (
	"time"
)

type Post struct {
	Id           int64  `form:"id"`
	Slug         string `form:"slug"`
	Title        string `xorm:"VARCHAR(300)" form:"title"`
	Content      string `xorm:"text" form:"content"`
	Tags         string `xorm:"text" form:"tags"`
	CommentCount int
	Created      time.Time
}

//func (post *Post) Validate(v *revel.Validation) {
//	v.Required(post.Slug).Message("请填写访问Slug")
//	v.Required(post.Title).Message("标题不能为空")
//	v.Required(post.Content).Message("文章内容不能为空")
//}

func (post *Post) GetCreated() time.Time {
	return post.Created.Add(-time.Hour * 8)
}
